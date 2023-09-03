package user

import (
	"context"
	"jt_chat/internal/model"
	"jt_chat/internal/service"

	"jt_chat/api/user/v1"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	_, err = service.User().Register(ctx, model.RegisterInput{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &v1.RegisterRes{Email: req.Email}, nil
}
