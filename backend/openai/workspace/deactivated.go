package workspace

import "github.com/gogf/gf/v2/net/ghttp"

func Deactivated(r *ghttp.Request) {
	// r.Session.RemoveAll()
	// 删除 _account cookie
	// r.Cookie.Remove("_account")
	r.Cookie.Set("_account", "personal")
	r.Response.WriteTpl("deactivated.html")
}
