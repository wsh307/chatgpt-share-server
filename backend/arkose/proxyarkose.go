package arkose

import (
	"backend/config"
	"backend/ratelimit"
	"backend/utility"
	"net/http/httputil"
	"net/url"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	UpStream    = config.CHATPROXY
	proxy       *httputil.ReverseProxy
	Remote      *url.URL
	arkoselimit = ratelimit.NewRateLimiter(1, 10)
)

func init() {
	remote, _ := url.Parse(UpStream)
	Remote = remote
	proxy = httputil.NewSingleHostReverseProxy(remote)
}

func ProxyArkose(r *ghttp.Request) {
	ctx := r.GetCtx()
	path := r.RequestURI
	// g.Log().Info(ctx, "ProxyArkose", path)
	isAdmin := r.Session.MustGet("isAdmin").Bool()
	if !isAdmin {
		usertoken := r.Session.MustGet("usertoken").String()
		if usertoken == "" {
			g.Log().Error(ctx, "usertoken is empty")
			r.Response.Status = 401
			r.Response.WriteJson(g.Map{
				"detail": "Authentication credentials were not provided.",
			})
			return
		}
		if path == "/fc/gt2/public_key/35536E1E-65B4-4D96-9D97-6ADB7EFF8147" && !arkoselimit.Allow(usertoken) {
			r.Response.Status = 429
			r.Response.WriteJson(g.Map{
				"detail": "Too Many Requests",
			})
			return
		}

		carid := r.Session.MustGet("carid").String()

		_, err := utility.CheckCar(ctx, carid)
		if err != nil {
			g.Log().Error(ctx, err)
			r.Response.Status = 401
			r.Response.WriteJson(g.Map{
				"detail": "Authentication credentials were not provided.",
			})
			return
		}
	}

	newreq := r.Request.Clone(ctx)
	newreq.URL.Host = Remote.Host
	newreq.URL.Scheme = Remote.Scheme
	newreq.Host = Remote.Host
	g.Log().Info(ctx, "ProxyArkose", path, "--->", newreq.URL.String())

	proxy.ServeHTTP(r.Response.Writer.RawWriter(), newreq)

}
