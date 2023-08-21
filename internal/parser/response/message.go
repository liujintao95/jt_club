package response

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type MsgResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessMsg(data any) *MsgResponse {
	msg := &MsgResponse{
		Code: 0,
		Msg:  "SUCCESS",
		Data: data,
	}
	return msg
}

func FailMsg(msg string, code int) *MsgResponse {
	msgObj := &MsgResponse{
		Code: code,
		Msg:  msg,
	}
	return msgObj
}

func Fail(ctx *gin.Context, respMsg string, respCode int, errMsg error) {
	var (
		strCode  string
		httpCode int
	)
	strCode = string(rune(respCode))[0:3]
	httpCode, _ = strconv.Atoi(strCode)
	_ = ctx.Error(errMsg)
	ctx.JSON(httpCode, FailMsg(respMsg, respCode))
}
