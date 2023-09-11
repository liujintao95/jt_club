package user

import (
	"context"
	"jt_chat/internal/model"
	"jt_chat/internal/service"

	"jt_chat/api/user/v1"
)

func (c *ControllerV1) CreateContactApplication(ctx context.Context, req *v1.CreateContactApplicationReq) (res *v1.CreateContactApplicationRes, err error) {
	_, err = service.User().CreateContactApplication(ctx, model.CreateContactApplicationInput{
		ContactId:   req.ContactId,
		ContactType: req.ContactType,
		Notice:      req.Notice,
	})
	return res, err
}
