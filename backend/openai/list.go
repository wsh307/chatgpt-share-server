package openai

import "github.com/gogf/gf/v2/net/ghttp"

func List(r *ghttp.Request) {
	r.Response.WriteTpl("list.html")
}
