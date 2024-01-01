package auth

import (
	"backend/modules/chatgpt/model"

	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 登陆接口演示
func Oauth(r *ghttp.Request) {
	ctx := r.GetCtx()
	usertoken := r.Get("usertoken").String()
	record, err := cool.DBM(model.NewChatgptUser()).Where("usertoken", usertoken).Where("expireTime>?", gconv.Time(gtime.Now())).One()
	if err != nil {
		g.Log().Error(ctx, err)
		r.Response.WriteJson(g.Map{
			"code": 0,
			"msg":  "服务器错误",
		})
		return
	}
	if record == nil {
		r.Response.WriteJson(g.Map{
			"code": 0,
			"msg":  "用户不存在或已过期",
		})
		return
	}
	var user model.ChatgptUser
	err = gconv.Struct(record, &user)
	if err != nil {
		g.Log().Error(ctx, err)
		r.Response.WriteJson(g.Map{
			"code": 0,
			"msg":  "服务器错误",
		})
		return
	}

}
