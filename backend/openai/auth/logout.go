package auth

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

func Logout(r *ghttp.Request) {
	// 退出登录
	r.Session.RemoveAll()
	r.Response.RedirectTo("/list")
}
