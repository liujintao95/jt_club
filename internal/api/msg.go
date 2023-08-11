package api

import (
	"JT_CLUB/internal/bll"
	"JT_CLUB/internal/constant"
	"JT_CLUB/internal/models"
	"JT_CLUB/internal/parser/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SocketClient(ctx *gin.Context) {
	currentUser, ok := ctx.Get(constant.CurrentUserKey)
	if !ok {
		_ = ctx.Error(fmt.Errorf("ctx get current user is nil"))
		ctx.JSON(http.StatusUnauthorized, response.FailMsg("启动聊天服务失败"))
		return
	}
	err := bll.RunSocketClient(ctx, currentUser.(*models.Users))
	if err != nil {
		_ = ctx.Error(fmt.Errorf("run socket client:" + err.Error()))
		ctx.JSON(http.StatusInternalServerError, response.FailMsg("启动聊天服务失败"))
		return
	}
	ctx.Status(http.StatusCreated)
}
