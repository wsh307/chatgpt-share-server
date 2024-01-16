package auth

import "github.com/gogf/gf/v2/frame/g"

func init() {
	s := g.Server()
	// auth路由组
	authGroup := s.Group("/auth")
	authGroup.GET("/login", Login)
	authGroup.POST("/login", Login)
	authGroup.POST("/oauth", Oauth)
	authGroup.POST("/oauthfree", OauthFree)
	authGroup.GET("/logout", Logout)
	authGroup.ALL("/logintoken", LoginToken)
}
