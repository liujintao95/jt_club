package api

import (
	"JT_CLUB/conf"
	"JT_CLUB/internal/bll"
	"JT_CLUB/internal/code"
	"JT_CLUB/internal/constant"
	"JT_CLUB/internal/parser/request"
	"JT_CLUB/internal/parser/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignIn(ctx *gin.Context) {
	var (
		login request.SignInRequest
		token string
		err   error
	)
	if err = ctx.ShouldBindJSON(&login); err != nil {
		response.Fail(ctx, "登录信息不合规", code.RequestDataError, err)
		return
	}
	token, err = bll.Login(&login)
	if err != nil {
		response.Fail(ctx, "登录失败", code.SignInError, err)
		return
	}
	ctx.SetCookie(constant.TokenKey, token, conf.DefaultDuration, "/", "", false, true)
	ctx.Status(http.StatusOK)
}

func SignUp(ctx *gin.Context) {
	var (
		newUser request.SignUpRequest
		err     error
	)
	if err = ctx.ShouldBindJSON(&newUser); err != nil {
		response.Fail(ctx, "注册信息不合规", code.RequestDataError, err)
		return
	}
	if _, err = bll.CreateUser(&newUser); err != nil {
		response.Fail(ctx, "注册失败", code.SignOnError, err)
		return
	}
	ctx.JSON(http.StatusCreated, response.SuccessMsg(response.SignUpResponse{Email: newUser.Email}))
}
