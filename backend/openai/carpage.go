package openai

import (
	"backend/modules/chatgpt/model"

	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 分页返回车辆列表
func CarPage(r *ghttp.Request) {
	// ctx := r.GetCtx()
	reqJson, err := r.GetJson()
	if err != nil {
		r.Response.WriteJson(g.Map{
			"detail": "Unable to parse request body.",
		})
		return
	}
	page := reqJson.Get("page").Int()
	size := reqJson.Get("size").Int()
	if page == 0 {
		page = 1
	}
	if size == 0 {
		size = 10
	}
	record, count, err := cool.DBM(model.NewChatgptSession()).Fields("carID", "status", "isPlus").Limit(size).Offset((page - 1) * size).AllAndCount(false)
	if err != nil {
		r.Response.WriteJson(g.Map{
			"detail": err.Error(),
		})
		return
	}
	r.Response.WriteJson(g.Map{
		"code":     1000,
		"messages": "success",
		"data": g.Map{
			"list": record,
			"pagination": g.Map{
				"page":  page,
				"size":  size,
				"total": count,
			},
		},
	})
}
