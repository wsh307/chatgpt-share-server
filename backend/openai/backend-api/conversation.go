package backendapi

import (
	"backend/config"
	"backend/modules/chatgpt/model"
	"backend/utility"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 修改对话标题
func ConversationPATCH(r *ghttp.Request) {
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

	conversationId := r.GetRouter("id").String()
	reqJson, err := r.GetJson()
	if err != nil {
		g.Log().Error(ctx, err)
		r.Response.Status = 400
		r.Response.WriteJson(g.Map{
			"detail": "Unable to parse request body.",
		})
		return
	}
	is_visible := reqJson.Get("is_visible").String()
	// 如果is_visible为false，则删除对话
	if is_visible == "false" {
		cool.DBM(model.NewChatgptConversations()).Where(g.Map{
			"convid":    conversationId,
			"usertoken": usertoken,
			"email":     carinfo.Email,
		}).Delete()
		r.Response.WriteJson(g.Map{
			"succeed": true,
		})
		return
	}
	title := reqJson.Get("title").String()
	// 如果title不为空，则修改对话标题
	AccessToken := carinfo.AccessToken
	originUrl := config.CHATPROXY + "/backend-api/conversation/" + conversationId
	resp, err := g.Client().SetAgent(r.Header.Get("User-Agent")).SetHeaderMap(g.MapStrStr{
		"Authorization": "Bearer " + AccessToken,
		"Content-Type":  "application/json",
	}).Patch(ctx, originUrl, g.MapStrStr{
		"title": title,
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
	_, err = cool.DBM(model.NewChatgptConversations()).Data(g.Map{
		"title": title,
	}).Where("convid", conversationId).Update()
	if err != nil {
		g.Log().Error(ctx, err)
		r.Response.Status = 400
	}
	r.Response.WriteJson(g.Map{
		"succeed": true,
	})
}

// Conversation 会话
func Conversation(r *ghttp.Request) {
	ctx := r.Context()
	// 获取header中的token
	usertoken := r.Session.MustGet("usertoken").String()
	if usertoken == "" {
		r.Response.Status = 401
		r.Response.WriteJson(g.Map{
			"detail": "Could not parse your authentication token. Please try signing in again.",
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
	utility.GetStatsInstance(carid).RecordCall()
	// 读取请求内容
	body, err := r.GetJson()
	if err != nil {
		r.Response.Status = 400
		r.Response.WriteJson(g.Map{
			"detail": "Unable to parse request body.",
		})
		return
	}
	// 如果配置了限制url
	if config.AuditLimitUrl != "" {

		res, err := g.Client().SetHeaderMap(g.MapStrStr{
			"Authorization": "Bearer " + usertoken,
			"Content-Type":  "application/json",
			"Cookie":        r.Header.Get("Cookie"),
			"Referer":       r.Header.Get("Referer"),
			"User-Agent":    r.Header.Get("User-Agent"),
		}).Post(ctx, config.AuditLimitUrl, body)
		if err != nil {
			r.Response.Status = 400
			r.Response.WriteJson(g.Map{
				"detail": "rate server error",
			})
			return
		}
		defer res.Close()
		if res.StatusCode != 200 {
			r.Response.Status = res.StatusCode
			r.Response.WriteJson(gjson.New(res.ReadAllString()))
			return
		}
	}
	AccessToken := carinfo.AccessToken
	u, _ := url.Parse(config.CHATPROXY)
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ErrorHandler = func(writer http.ResponseWriter, request *http.Request, e error) {
		g.Log().Error(ctx, e)
		writer.WriteHeader(http.StatusBadGateway)
	}
	proxy.ModifyResponse = func(proxyResponse *http.Response) error {
		if proxyResponse.StatusCode == 429 {
			body, err := io.ReadAll(proxyResponse.Body)
			if err != nil {
				return err
			}
			bodyJson, err := gjson.DecodeToJson(body)
			if err != nil {
				return err
			}
			clears_in := bodyJson.Get("detail.clears_in").Int()
			if clears_in > 0 && clears_in != 3600 {
				clearsDuration := time.Duration(clears_in) * time.Second
				cool.CacheManager.Set(ctx, "clears_in:"+carid, clears_in, clearsDuration)
			}
			// 将原始内容写入到响应中
			r.Response.Status = 429
			r.Response.Write(body)
		}

		proxyResponse.Header.Del("Set-Cookie")

		return nil
	}
	newreq := r.Request.Clone(ctx)
	newreq.URL.Host = u.Host
	newreq.URL.Scheme = u.Scheme
	newreq.Host = u.Host
	// g.Dump(newreq.Header)
	newreq.Header.Set("authkey", config.AUTHKEY)
	newreq.Header.Set("Authorization", "Bearer "+AccessToken)
	proxy.ServeHTTP(r.Response.Writer.RawWriter(), newreq)
}
