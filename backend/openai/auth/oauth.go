package auth

import (
	"backend/modules/chatgpt/model"
	"backend/utility"

	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/util/gconv"
)

// 登陆接口演示
func Oauth(r *ghttp.Request) {
	ctx := r.GetCtx()
	usertoken := r.Get("usertoken").String()
	carid := r.Get("carid").String()
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
	// g.Dump(record)

	//
	carinfo, err := utility.CheckCar(ctx, carid)
	if err != nil {
		g.Log().Error(ctx, err)
		r.Response.WriteJson(g.Map{
			"code": 0,
			"msg":  "服务器错误",
		})
		return
	}
	// PLUS用户可以直接登陆任意类型的车
	if record["isPlus"].Bool() {
		r.Response.WriteJson(g.Map{
			"code": 1,
			"msg":  "登陆成功",
		})
		return
	} else {
		// 非PLUS用户只能登陆同类型的车
		if record["isPlus"].Bool() == carinfo.IsPlus {
			r.Response.WriteJson(g.Map{
				"code": 1,
				"msg":  "登陆成功",
			})
			return
		} else {
			r.Response.WriteJson(g.Map{
				"code": 0,
				"msg":  "您不是PLUS用户，无法登陆该车",
			})
			return
		}
	}

}

// OauthFree 免费登陆接口演示
func OauthFree(r *ghttp.Request) {
	ctx := r.GetCtx()
	usertoken := r.Get("usertoken").String()
	// 检查usertoken是否6-18位字母数字组合

	carid := r.Get("carid").String()

	//
	_, err := utility.CheckCar(ctx, carid)
	if err != nil {
		g.Log().Error(ctx, err)
		r.Response.WriteJson(g.Map{
			"code": 0,
			"msg":  "服务器错误",
		})
		return
	}
	if !gregex.IsMatchString(`^[a-zA-Z0-9]{6,18}$`, usertoken) {
		r.Response.WriteJson(g.Map{
			"code": 0,
			"msg":  "请输入6-18位字母数字组合",
		})
		return
	}

	r.Response.WriteJson(g.Map{
		"code": 1,
		"msg":  "登陆成功",
	})
}
