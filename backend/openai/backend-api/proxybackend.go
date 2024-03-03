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
	"github.com/gogf/gf/v2/text/gstr"
)

func ProxyBackend(r *ghttp.Request) {
	ctx := r.GetCtx()
	// usertoken := r.Session.MustGet("usertoken").String()
	carid := r.Session.MustGet("carid").String()
	conv := r.GetRouter("convid").String()
	fileid := r.GetRouter("fileid").String()
	if fileid != "" {
		g.Log().Info(ctx, "fileid:", fileid)
		refer := r.Referer()
		g.Log().Info(ctx, "refer:", refer)
		// http://localhost:8001/c/98b65c1a-27e1-40d3-b045-49c11b34d768
		// 从refer中获取convid /c/ 后面的内容
		referArr := gstr.Split(refer, "/c/")
		if len(referArr) > 1 {
			conv = referArr[1]
			g.Log().Info(ctx, "conv:", conv)
		}
	}
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
		caridVar, err := cool.CacheManager.Get(ctx, "email:"+result["email"].String())
		if err != nil {
			g.Log().Error(ctx, err)
			r.Response.Status = 500
			r.Response.WriteJson(g.Map{
				"detail": "Internal Server Error",
			})
			return
		}
		carid = caridVar.String()
		if carid == "" {
			r.Response.Status = 404
			r.Response.WriteJson(g.Map{
				"detail": "The car " + conv + " belongs to is unavailable",
			})
			return
		}
		// r.Session.Set("carid", carid)
		chatgptaccountid := result["chatgptaccountid"].String()
		if chatgptaccountid != "" {
			r.Header.Set("ChatGPT-Account-ID", chatgptaccountid)
		} else {
			r.Header.Del("ChatGPT-Account-ID")
		}
		// r.Session.Set("carid", carid)
		// r.Session.Set("chatgptaccountid", chatgptaccountid)
		r.Session.Set("convcarid", carid)
		r.Session.Set("convchatgptaccountid", chatgptaccountid)
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
