package utility

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type CarInfo struct {
	Carid         string
	Email         string
	IsPlus        bool
	IsPlusStr     string
	RefreshCookie string
	AccessToken   string
}

func CheckCar(ctx g.Ctx, carid string) (carInfo *CarInfo, err error) {
	sessionVar, err := cool.CacheManager.Get(ctx, "session:"+carid)
	if err != nil {
		return
	}
	sessionJson := gjson.New(sessionVar)
	carInfo = &CarInfo{}
	carInfo.Carid = carid
	email := sessionJson.Get("user.email").String()
	if email == "" {
		err = gerror.New("email is empty")
		return
	}
	carInfo.Email = email
	refreshCookie := sessionJson.Get("refreshCookie").String()
	if refreshCookie == "" {
		err = gerror.New("refreshCookie is empty")
		return
	}
	carInfo.RefreshCookie = refreshCookie
	accessToken := sessionJson.Get("accessToken").String()
	if accessToken == "" {
		err = gerror.New("accessToken is empty")
		return
	}
	carInfo.AccessToken = accessToken
	models := sessionJson.Get("models").Array()
	if len(models) == 0 {
		err = gerror.New("models is empty")
		return
	}
	carInfo.IsPlus = len(models) > 1
	if carInfo.IsPlus {
		carInfo.IsPlusStr = "PLUS"
	} else {
		carInfo.IsPlusStr = "3.5"
	}
	return
}
