package auth

import "github.com/gogf/gf/v2/net/ghttp"

func Login(r *ghttp.Request) {
	r.Response.WriteTpl("login.html")
}
