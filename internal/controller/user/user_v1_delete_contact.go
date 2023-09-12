package user

import (
	"context"
	"jt_chat/internal/model"
	"jt_chat/internal/service"

	"jt_chat/api/user/v1"
)

func (c *ControllerV1) DeleteContact(ctx context.Context, req *v1.DeleteContactReq) (res *v1.DeleteContactRes, err error) {
	_, err = service.User().DeleteContact(ctx, model.DeleteContactInput{
		Uid: req.Uid,
	})
	return res, err
}
