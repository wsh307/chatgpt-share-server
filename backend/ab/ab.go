package ab

import "github.com/gogf/gf/v2/frame/g"

func Init(ctx g.Ctx) {
	s := g.Server()
	abGroup := s.Group("/v1")
	abGroup.POST("/initialize", Initialize)
	abGroup.POST("/rgstr", Rgstr)

}
