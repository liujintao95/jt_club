package api

import (
	"JT_CLUB/conf"
	"JT_CLUB/internal/bll"
	"JT_CLUB/internal/constant"
	"JT_CLUB/internal/parser/request"
	"JT_CLUB/internal/parser/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignIn(ctx *gin.Context) {
	var login *request.SignInRequest
	if err := ctx.ShouldBindJSON(login); err != nil {
		_ = ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, response.FailMsg("登录信息不合规"))
		return
	}
	token, err := bll.Login(login)
	if err != nil {
		_ = ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, response.FailMsg("登录失败"))
		return
	}
	ctx.SetCookie(constant.TokenKey, token, conf.DefaultDuration, "/", "", false, true)
	ctx.Status(http.StatusOK)
}

func SignUp(ctx *gin.Context) {
	var newUser *request.SignUpRequest
	if err := ctx.ShouldBindJSON(newUser); err != nil {
		_ = ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, response.FailMsg("注册信息不合规"))
		return
	}
	if _, err := bll.CreateUser(newUser); err != nil {
		_ = ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, response.FailMsg("注册失败"))
		return
	}
	ctx.JSON(http.StatusCreated, response.SuccessMsg(response.SignUpResponse{Email: newUser.Email}))
}
