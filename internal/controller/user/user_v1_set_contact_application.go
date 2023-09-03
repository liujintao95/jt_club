package user

import (
	"context"
	"jt_chat/internal/model"
	"jt_chat/internal/service"

	"jt_chat/api/user/v1"
)

func (c *ControllerV1) SetContactApplication(ctx context.Context, req *v1.SetContactApplicationReq) (res *v1.SetContactApplicationRes, err error) {
	_, err = service.User().SetContactApplication(ctx, model.SetContactApplicationInput{
		ContactId:   req.ContactId,
		ContactType: req.ContactType,
		Notice:      req.Notice,
	})
	return res, err
}
