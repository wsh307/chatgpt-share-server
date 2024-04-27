package next

import (
	"backend/config"
	"backend/utility"
	"bytes"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

func ProxyNext(r *ghttp.Request) {
	ctx := r.GetCtx()

	carid := r.Session.MustGet("carid").String()
	carinfo, err := utility.CheckCar(ctx, carid)
	if err != nil {
		g.Log().Error(ctx, err)
		r.Response.Status = 401
		r.Response.WriteJson(g.Map{
			"detail": "Authentication credentials were not provided.",
		})
		return
	}
	u, _ := url.Parse(config.CHATPROXY)
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ErrorHandler = func(writer http.ResponseWriter, request *http.Request, e error) {
		g.Log().Error(ctx, e)
		writer.WriteHeader(http.StatusBadGateway)
	}
	proxy.ModifyResponse = func(response *http.Response) error {
		response.Header.Del("Set-Cookie")
		if response.StatusCode != 200 {
			return nil
		}
		if response.Header.Get("Content-Type") != "application/json" {
			return nil
		}
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}
		// 修改响应体
		bodyJson := gjson.New(body)
		bodyJson.Set("pageProps.user.email", "admin@openai.com")
		bodyJson.Set("pageProps.user.name", "admin")
		bodyJson.Set("pageProps.user.image", "/avatars.png")
		bodyJson.Set("pageProps.user.picture", "/avatars.png")
		bodyJson.Set("pageProps.user.id", "user-xadmin")
		bodyStr := bodyJson.String()
		// 替换 body 中的邮件
		bodyStr = gstr.Replace(bodyStr, carinfo.Email, "share@closeai.biz", -1)
		bodyJson = gjson.New(bodyStr)

		// 写入响应体
		response.Body = io.NopCloser(bytes.NewReader(gconv.Bytes(bodyJson)))
		// 重写响应头大小
		response.ContentLength = int64(len(body))

		return nil
	}
	newreq := r.Request.Clone(ctx)
	newreq.URL.Host = u.Host
	newreq.URL.Scheme = u.Scheme
	newreq.Host = u.Host
	newreq.Header.Set("Cookie", "__Secure-next-auth.session-token="+carinfo.RefreshCookie)
	// 去除header 中的 压缩
	newreq.Header.Del("Accept-Encoding")
	proxy.ServeHTTP(r.Response.Writer.RawWriter(), newreq)

}
