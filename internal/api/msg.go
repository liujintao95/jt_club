package api

import (
	"JT_CLUB/internal/bll"
	"JT_CLUB/internal/code"
	"JT_CLUB/internal/constant"
	"JT_CLUB/internal/models"
	"JT_CLUB/internal/parser/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SocketClient(ctx *gin.Context) {
	var (
		currentUser any
		err         error
	)
	currentUser, _ = ctx.Get(constant.CurrentUserKey)
	err = bll.RunSocketClient(ctx, currentUser.(*models.Users))
	if err != nil {
		response.Fail(ctx, "启动聊天服务失败", code.RunClientError, err)
		return
	}
	ctx.Status(http.StatusCreated)
}
