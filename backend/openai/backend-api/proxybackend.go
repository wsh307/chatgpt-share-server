package backendapi

import (
	"backend/config"
	"backend/modules/chatgpt/model"
	"backend/utility"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func ProxyBackend(r *ghttp.Request) {
	ctx := r.GetCtx()
	// usertoken := r.Session.MustGet("usertoken").String()
	carid := r.Session.MustGet("carid").String()
	conv := r.GetRouter("convid").String()
	// chatgptaccountid := r.Header.Get("ChatGPT-Account-ID")
	if conv != "" {
		g.Log().Info(ctx, "conv:", conv)
		// 查询会话
		result, err := cool.DBM(model.NewChatgptConversations()).Where(g.Map{
			"convid": conv,
		}).One()
		if err != nil {
			g.Log().Error(ctx, err)
			r.Response.Status = 500
			r.Response.WriteJson(g.Map{
				"detail": "Internal Server Error",
			})
			return
		}
		if result == nil {
			r.Response.Status = 404
			r.Response.WriteJson(g.Map{
				"detail": "Can't load conversation " + conv,
			})
			return
		}
		carid = cool.CacheManager.MustGet(ctx, "email:"+result["email"].String()).String()
		if carid == "" {
			r.Response.Status = 404
			r.Response.WriteJson(g.Map{
				"detail": "Can't load conversation " + conv,
			})
			return
		}
		r.Session.Set("carid", carid)
		chatgptaccountid := result["chatgptaccountid"].String()
		if chatgptaccountid != "" {
			r.Header.Set("ChatGPT-Account-ID", chatgptaccountid)
		} else {
			r.Header.Del("ChatGPT-Account-ID")
		}

	}

	carinfo, err := utility.CheckCar(ctx, carid)
	if err != nil {
		g.Log().Error(ctx, err)
		r.Response.Status = 401
		r.Response.WriteJson(g.Map{
			"detail": "Authentication credentials were not provided.",
		})
		return
	}

	Authorization := r.Header.Get("Authorization")
	if Authorization != "" {
		r.Header.Set("Authorization", "Bearer "+carinfo.AccessToken)
	}

	u, _ := url.Parse(config.CHATPROXY)
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ErrorHandler = func(writer http.ResponseWriter, request *http.Request, e error) {
		g.Log().Error(ctx, e)
		writer.WriteHeader(http.StatusBadGateway)
	}
	newreq := r.Request.Clone(ctx)
	newreq.URL.Host = u.Host
	newreq.URL.Scheme = u.Scheme
	newreq.Host = u.Host
	// g.Dump(newreq.Header)
	newreq.Header.Set("authkey", config.AUTHKEY)
	proxy.ServeHTTP(r.Response.Writer.RawWriter(), newreq)
}
