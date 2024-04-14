package auth

import "github.com/gogf/gf/v2/net/ghttp"

func Csrf(r *ghttp.Request) {
	resStr := `{"csrfToken":"csrf-token"}`
	r.Response.WriteJsonExit(resStr)
}
