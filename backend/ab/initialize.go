package ab

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

func Initialize(r *ghttp.Request) {
	// ctx := r.GetCtx()
	r.Response.Status = 204
	return
	// initialize := gfile.GetContents("initialize.json")
	// r.Response.WriteJson(gjson.New(initialize))
	// return
	// header := r.Header
	// utility.RemoveCfHeaders(header)
	// header.Set("Host", "ab.chatgpt.com")
	// header.Set("Origin", "http://ab.chatgpt.com")
	// header.Set("Referer", "http://chatgpt.com/")
	// g.Dump(header)
	// body := r.GetBody()
	// g.Dump(body)
	// // https://httpbin.org/post
	// // http://ab.chatgpt.com/v1/initialize
	// res, err := requests.Post(ctx, "http://ab.chatgpt.com/v1/initialize", requests.RequestOption{
	// 	Headers: header,
	// 	Body:    body,
	// })
	// if err != nil {
	// 	g.Log().Error(ctx, err)
	// 	r.Response.Status = 500
	// 	r.Response.WriteJson(g.Map{
	// 		"detail": err.Error(),
	// 	})
	// 	return
	// }
	// r.Response.WriteJson(gjson.New(res.Text()))
}
