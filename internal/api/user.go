package api

import (
	"JT_CLUB/conf"
	"JT_CLUB/internal/bll"
	"JT_CLUB/internal/code"
	"JT_CLUB/internal/constant"
	"JT_CLUB/internal/models"
	"JT_CLUB/internal/parser/request"
	"JT_CLUB/internal/parser/response"
	"fmt"
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
		response.FailRequest(ctx, err)
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
		response.FailRequest(ctx, err)
		return
	}
	if _, err = bll.CreateUser(&newUser); err != nil {
		response.Fail(ctx, "注册失败", code.SignOnError, err)
		return
	}
	ctx.JSON(http.StatusCreated, response.SuccessMsg(response.SignUp{Email: newUser.Email}))
}

func UserSelect(ctx *gin.Context) {
	var (
		currentUser any
		userSelect  request.UserSelect
		userList    []*response.UserInfo
		err         error
	)
	currentUser, _ = ctx.Get(constant.CurrentUserKey)
	if err = ctx.ShouldBindJSON(&userSelect); err != nil {
		response.FailRequest(ctx, err)
		return
	}
	userList, err = bll.SelectUser(currentUser.(*models.User), userSelect.Query, userSelect.IsContact)
	if err != nil {
		response.Fail(ctx, "查询用户失败", code.ContactListError, err)
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessMsg(userList))
}

func ContactList(ctx *gin.Context) {
	var (
		currentUser any
		contactList []*response.ContactInfo
		err         error
	)
	currentUser, _ = ctx.Get(constant.CurrentUserKey)
	contactList, err = bll.GetContactList(currentUser.(*models.User))
	if err != nil {
		response.Fail(ctx, "获取联系人列表错误", code.ContactListError, err)
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessMsg(contactList))
}

func ContactApplication(ctx *gin.Context) {
	var (
		currentUser any
		application request.ContactApplication
		err         error
	)
	if err = ctx.ShouldBindJSON(&application); err != nil {
		response.FailRequest(ctx, err)
		return
	}
	currentUser, _ = ctx.Get(constant.CurrentUserKey)
	if _, err = bll.CreateContactApplication(currentUser.(*models.User), application); err != nil {
		response.Fail(ctx, "发送请求失败", code.ContactRequestError, err)
		return
	}
	ctx.JSON(http.StatusCreated, response.Success())
}

func ContactConfirm(ctx *gin.Context) {
	var (
		confirmInfo request.ContactConfirm
		err         error
	)
	if err = ctx.ShouldBindJSON(&confirmInfo); err != nil {
		response.FailRequest(ctx, err)
		return
	}
	if confirmInfo.Status != constant.RequestAgreeStatus && confirmInfo.Status != constant.RequestRefuseStatus {
		response.Fail(ctx, "请求信息不合规", code.RequestDataError, fmt.Errorf("status value error"))
		return
	}
	if err = bll.UpdateContactApplicationStatus(confirmInfo); err != nil {
		response.Fail(ctx, "发送失败", code.ContactRequestError, err)
		return
	}
	ctx.JSON(http.StatusCreated, response.Success())
}

func ContactConfirmList(ctx *gin.Context) {
	var (
		confirmList []*response.ApplicationInfo
		currentUser any
		err         error
	)
	currentUser, _ = ctx.Get(constant.CurrentUserKey)
	confirmList, err = bll.GetContactApplicationConfirmList(currentUser.(*models.User))
	if err != nil {
		response.Fail(ctx, "发送请求失败", code.ContactRequestError, err)
		return
	}
	ctx.JSON(http.StatusCreated, response.SuccessMsg(confirmList))
}
