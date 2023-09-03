package user

import (
	"context"
	"jt_chat/internal/model"
	"jt_chat/internal/service"

	"jt_chat/api/user/v1"
)

func (c *ControllerV1) UpdateContactApplication(ctx context.Context, req *v1.UpdateContactApplicationReq) (res *v1.UpdateContactApplicationRes, err error) {
	_, err = service.User().UpdateContactApplication(ctx, model.UpdateContactApplicationInput{
		AppId:  req.AppId,
		Status: req.Status,
	})
	return res, err
}
