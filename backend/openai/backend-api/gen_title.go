package backendapi

import (
	"backend/config"
	"backend/modules/chatgpt/model"
	"backend/utility"

	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// GenTitle 生成会话标题
func GenTitle(r *ghttp.Request) {
	ctx := r.Context()
	// 获取当前登录用户信息
	usertoken := r.Session.MustGet("usertoken").String()
	if usertoken == "" {
		r.Response.Status = 401
		r.Response.WriteJson(g.Map{
			"detail": "Authentication credentials were not provided.",
		})
		return
	}
	carid := r.Session.MustGet("carid").String()
	carinfo, err := utility.CheckCar(ctx, carid)

	if err != nil {
		g.Log().Error(ctx, err)
		r.Response.Status = 401
		r.Response.WriteJson(g.Map{
			"detail": "Authentication credentials were not provided.",
		})
		return
	}

	reqJson, err := r.GetJson()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	message_id := reqJson.Get("message_id").String()
	AccessToken := carinfo.AccessToken
	id := r.GetRouter("id").String()
	originUrl := config.CHATPROXY + "/backend-api/conversation/gen_title/" + id
	resp, err := g.Client().SetAgent(r.Header.Get("User-Agent")).SetHeaderMap(g.MapStrStr{
		"Authorization":      "Bearer " + AccessToken,
		"Content-Type":       "application/json",
		"ChatGPT-Account-ID": r.Header.Get("ChatGPT-Account-ID"),
	}).Post(ctx, originUrl, g.MapStrStr{
		"message_id": message_id,
	})
	if err != nil {
		g.Log().Error(ctx, err)
		r.Response.Status = 500
		r.Response.WriteJson(g.Map{
			"detail": err.Error(),
		})
		return
	}
	defer resp.Close()
	r.Response.Status = resp.StatusCode
	if resp.StatusCode != 200 {
		r.Response.Write(resp.ReadAllString())
		return
	}
	respBody := resp.ReadAllString()
	respJson := gjson.New(respBody)
	title := respJson.Get("title").String()
	if title != "" {
		cool.DBM(model.NewChatgptConversations()).Save(g.Map{
			"convid":           id,
			"title":            title,
			"usertoken":        usertoken,
			"email":            carinfo.Email,
			"chatgptaccountid": r.Header.Get("ChatGPT-Account-ID"),
		})
	}
	r.Response.WriteJsonExit(respBody)

}
