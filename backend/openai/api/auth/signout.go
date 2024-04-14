package auth

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func SignOut(r *ghttp.Request) {
	r.Session.RemoveAll()
	// r.Response.RedirectTo("/list")
	r.Response.WriteJsonExit(g.Map{
		"url": "/",
	})
}
