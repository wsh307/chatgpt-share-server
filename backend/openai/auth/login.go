package auth

import (
	"backend/config"
	"backend/utility"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/narqo/go-badge"
)

func Login(r *ghttp.Request) {
	ctx := r.GetCtx()
	method := r.Method
	if method == "GET" {
		req := r.GetMapStrStr()

		carid := req["carid"]
		if carid == "" {
			r.Response.WriteTpl("login.html")
			return
		}
		carInfo, err := utility.CheckCar(ctx, carid)
		if err != nil {
			g.Log().Error(ctx, err)
			badge, err := badge.RenderBytes("😭", "      翻车|不可用", "red")
			if err != nil {
				g.Log().Error(ctx, err)
				r.Response.WriteTpl("login.html")
			}
			r.Response.WriteTpl("login.html", g.Map{"badge": string(badge)})

			return
		}

		badge, err := badge.RenderBytes(carInfo.IsPlusStr, "    😊空闲|推荐", "green")
		if err != nil {
			g.Log().Error(ctx, err)
			r.Response.WriteTpl("login.html")
		}
		// fmt.Printf("%s", badge)

		r.Response.WriteTpl("login.html", g.Map{"badge": string(badge)})
		return
	} else {
		req := r.GetMapStrStr()
		loginVar := g.Client().PostVar(ctx, config.OauthUrl, req)
		loginJson := gjson.New(loginVar)
		loginJson.Dump()
		code := loginJson.Get("code").Int()
		if code != 1 {
			msg := loginJson.Get("msg").String()
			r.Response.WriteTpl("login.html", g.Map{
				"error": msg,
				"carid": req["carid"],
			})
			return
		}
	}
}
