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
		CreateContactApplication(ctx context.Context, in model.CreateContactApplicationInput) (out model.CreateContactApplicationOutput, err error)
		UpdateContactApplication(ctx context.Context, in model.UpdateContactApplicationInput) (out model.UpdateContactApplicationOutput, err error)
		CreateUserGroup(ctx context.Context, in model.CreateUserGroupInput) (out model.CreateUserGroupOutput, err error)
		UpdateUserGroup(ctx context.Context, in model.UpdateUserGroupInput) (out model.UpdateUserGroupOutput, err error)
		DeleteContact(ctx context.Context, in model.DeleteContactInput) (out model.DeleteContactOutput, err error)
		DeleteUserGroupMap(ctx context.Context, in model.DeleteUserGroupMapInput) (out model.DeleteUserGroupMapOutput, err error)
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
