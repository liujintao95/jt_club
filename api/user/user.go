// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package user

import (
	"context"
	
	"jt_chat/api/user/v1"
)

type IUserV1 interface {
	Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error)
	Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	GetContactList(ctx context.Context, req *v1.GetContactListReq) (res *v1.GetContactListRes, err error)
	DeleteContact(ctx context.Context, req *v1.DeleteContactReq) (res *v1.DeleteContactRes, err error)
	GetContactApplicationList(ctx context.Context, req *v1.GetContactApplicationListReq) (res *v1.GetContactApplicationListRes, err error)
	CreateContactApplication(ctx context.Context, req *v1.CreateContactApplicationReq) (res *v1.CreateContactApplicationRes, err error)
	UpdateContactApplication(ctx context.Context, req *v1.UpdateContactApplicationReq) (res *v1.UpdateContactApplicationRes, err error)
	CreateUserGroup(ctx context.Context, req *v1.CreateUserGroupReq) (res *v1.CreateUserGroupRes, err error)
	UpdateUserGroup(ctx context.Context, req *v1.UpdateUserGroupReq) (res *v1.UpdateUserGroupRes, err error)
	DeleteUserGroupMap(ctx context.Context, req *v1.DeleteUserGroupMapReq) (res *v1.DeleteUserGroupMapRes, err error)
}


