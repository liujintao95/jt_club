package user

import (
	"context"
	"jt_chat/internal/model"
	"jt_chat/internal/service"

	"jt_chat/api/user/v1"
)

func (c *ControllerV1) CreateUserGroup(ctx context.Context, req *v1.CreateUserGroupReq) (res *v1.CreateUserGroupRes, err error) {
	_, err = service.User().CreateUserGroup(ctx, model.CreateUserGroupInput{
		Name:   req.Name,
		Avatar: req.Avatar,
		Notice: req.Notice,
	})
	return res, err
}
