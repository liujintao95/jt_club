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
	GetContactApplicationList(ctx context.Context, req *v1.GetContactApplicationListReq) (res *v1.GetContactApplicationListRes, err error)
	SetContactApplication(ctx context.Context, req *v1.SetContactApplicationReq) (res *v1.SetContactApplicationRes, err error)
	UpdateContactApplication(ctx context.Context, req *v1.UpdateContactApplicationReq) (res *v1.UpdateContactApplicationRes, err error)
}


