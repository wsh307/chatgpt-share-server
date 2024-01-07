package middleware

import (
	"backend/config"

	"github.com/gogf/gf/v2/net/ghttp"
)

func ChatGPTAdminApiAuth(r *ghttp.Request) {
	// TODO
	apiauth := r.Header.Get("apiauth")
	if apiauth == "" {
		r.Response.WriteStatusExit(403, "apiauth is empty")
	}
	if apiauth != config.APIAUTH {
		r.Response.WriteStatusExit(403, "apiauth is error")
	}
	r.Middleware.Next()
}
