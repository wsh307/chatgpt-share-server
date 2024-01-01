package openai

import (
	"backend/openai/auth"

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

	// auth路由组
	authGroup := s.Group("/auth")
	authGroup.GET("/login", auth.Login)

}
