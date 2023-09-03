// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"jt_chat/internal/model"
)

type (
	IUser interface {
		Register(ctx context.Context, in model.RegisterInput) (out model.RegisterOutput, err error)
		Update(ctx context.Context, in model.UpdateInput) (out model.UpdateOutput, err error)
		GetList(ctx context.Context, in model.GetListInput) (out model.GetListOutput, err error)
		GetContactList(ctx context.Context, in model.GetContactListInput) (out model.GetContactListOutput, err error)
		GetContactApplicationList(ctx context.Context, in model.GetContactApplicationListInput) (out model.GetContactApplicationListOutput, err error)
		SetContactApplication(ctx context.Context, in model.SetContactApplicationInput) (out model.SetContactApplicationOutput, err error)
		UpdateContactApplication(ctx context.Context, in model.UpdateContactApplicationInput) (out model.UpdateContactApplicationOutput, err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
