package next

import "github.com/gogf/gf/v2/frame/g"

func init() {
	// 注册路由
	s := g.Server()
	nextGroup := s.Group("/_next")
	nextGroup.ALL("/*", ProxyNext)
}
