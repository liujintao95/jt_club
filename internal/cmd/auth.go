package cmd

import (
	"context"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"jt_chat/internal/consts"
	"jt_chat/internal/dao"
	"jt_chat/internal/model/entity"
	"jt_chat/utility"
	"jt_chat/utility/response"
)

type LoginRes struct {
	Type     string `json:"type"`
	Token    string `json:"token"`
	ExpireIn int    `json:"expire_in"`
	Uid      string `json:"uid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}

func StartGToken() (gfToken *gtoken.GfToken, err error) {
	gfToken = &gtoken.GfToken{
		CacheMode:       consts.CacheModeRedis,
		ServerName:      consts.BackendServerName,
		LoginPath:       "/login",
		LoginBeforeFunc: loginFunc,
		LoginAfterFunc:  loginAfterFunc,
		LogoutPath:      "/logout",
		//AuthPaths:        g.SliceStr{"/user"},
		AuthExcludePaths: g.SliceStr{"/user/register"},
		AuthAfterFunc:    authAfterFunc,
		MultiLogin:       consts.MultiLogin,
	}
	return
}

func loginFunc(r *ghttp.Request) (string, interface{}) {
	account := r.Get("account").String()
	password := r.Get("password").String()
	ctx := context.TODO()

	if account == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginMsg))
		r.ExitAll()
	}

	userInfo := entity.User{}
	err := dao.User.Ctx(ctx).Where(dao.User.Columns().Email, account).Scan(&userInfo)
	if err != nil {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginMsg))
		r.ExitAll()
	}
	if utility.EncryptPassword(password, consts.Salt) != userInfo.Password {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginMsg))
		r.ExitAll()
	}
	return userInfo.Uid, userInfo
}

func loginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		respData.Code = 0
		r.Response.WriteJson(respData)
		return
	} else {
		respData.Code = 1
		//获得登录用户id
		userId := respData.GetString("userKey")
		//根据id获得登录用户其他信息
		userInfo := entity.User{}
		err := dao.User.Ctx(context.TODO()).Where(dao.User.Columns().Uid, userId).Scan(&userInfo)
		if err != nil {
			return
		}
		data := LoginRes{
			Type:     consts.TokenType,
			Token:    respData.GetString("token"),
			ExpireIn: consts.GTokenExpireIn, //单位秒,
		}
		data.Name = userInfo.Name
		data.Avatar = userInfo.Avatar
		data.Email = userInfo.Email
		data.Uid = userInfo.Uid
		response.JsonExit(r, 0, "", data)
	}
	return
}

func authAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	var userInfo entity.User
	err := gconv.Struct(respData.GetString("data"), &userInfo)
	if err != nil {
		response.Auth(r)
		return
	}
	r.SetCtxVar(consts.CtxUserId, userInfo.Uid)
	r.SetCtxVar(consts.CtxUserName, userInfo.Name)
	r.SetCtxVar(consts.CtxUserAvatar, userInfo.Avatar)
	r.SetCtxVar(consts.CtxUserEmail, userInfo.Email)
	r.Middleware.Next()
}
