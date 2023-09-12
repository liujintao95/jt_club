package user

import (
	"context"
	"jt_chat/internal/model"
	"jt_chat/internal/service"

	"jt_chat/api/user/v1"
)

func (c *ControllerV1) UpdateUserGroup(ctx context.Context, req *v1.UpdateUserGroupReq) (res *v1.UpdateUserGroupRes, err error) {
	_, err = service.User().UpdateUserGroup(ctx, model.UpdateUserGroupInput{
		Gid:     req.Gid,
		Name:    req.Name,
		Avatar:  req.Avatar,
		Notice:  req.Notice,
		AdminId: req.AdminId,
	})
	return res, err
}
