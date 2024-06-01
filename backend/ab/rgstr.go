package ab

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func Rgstr(r *ghttp.Request) {
	r.Response.WriteJson(g.Map{
		"success": true,
	})
}
