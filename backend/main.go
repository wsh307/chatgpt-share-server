package main

import (
	_ "backend/internal/packed"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	_ "github.com/cool-team-official/cool-admin-go/contrib/drivers/mysql"
	_ "github.com/cool-team-official/cool-admin-go/contrib/drivers/sqlite"
	_ "github.com/cool-team-official/cool-admin-go/contrib/files/local"

	_ "backend/arkose"
	_ "backend/modules"
	_ "backend/openai"

	"backend/ab"

	"github.com/gogf/gf/v2/os/gctx"

	"backend/internal/cmd"
)

func main() {
	// gres.Dump()
	ctx := gctx.New()
	ab.Init(ctx)
	cmd.Main.Run(gctx.New())
}
