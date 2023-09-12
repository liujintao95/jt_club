package user

import (
	"context"
	"jt_chat/internal/model"
	"jt_chat/internal/service"

	"jt_chat/api/user/v1"
)

func (c *ControllerV1) DeleteUserGroupMap(ctx context.Context, req *v1.DeleteUserGroupMapReq) (res *v1.DeleteUserGroupMapRes, err error) {
	_, err = service.User().DeleteUserGroupMap(ctx, model.DeleteUserGroupMapInput{
		Gid: req.Gid,
	})
	return res, err
}
