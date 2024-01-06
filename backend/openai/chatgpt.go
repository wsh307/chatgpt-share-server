package openai

import (
	_ "backend/openai/api"
	_ "backend/openai/auth"
	_ "backend/openai/backend-api"
	_ "backend/openai/next"

	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	s := g.Server()
	// 根路由
	group := s.Group("/")
	group.GET("/", Index)
	group.GET("/c/:convId", C)
	group.GET("/g/:gizmoId", G)
	group.GET("/gpts/discovery", Discovery)
	group.GET("/gpts/editor", Editor)
	group.GET("/gpts/editor/:slug", Slug)
	group.GET("/g/:gizmoId/c/:convId", GC)
	group.GET(("/gpts/mine"), Mine)

	// 状态相关
	group.GET("/status", Status)
	group.GET("/endpoint", EndPoint)
	group.POST("/carpage", CarPage)

}
