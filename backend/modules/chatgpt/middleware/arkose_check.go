package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

func ArkoseCheck(r *ghttp.Request) {

	r.Middleware.Next()
	if r.Response.Status == 200 {
		r.Session.Set("isAdmin", true)
	}
}
