package user

import (
	"context"
	"jt_chat/internal/model"
	"jt_chat/internal/service"

	"jt_chat/api/user/v1"
)

func (c *ControllerV1) GetContactApplicationList(ctx context.Context, req *v1.GetContactApplicationListReq) (res *v1.GetContactApplicationListRes, err error) {
	var (
		out model.GetContactApplicationListOutput
	)
	out, err = service.User().GetContactApplicationList(ctx, model.GetContactApplicationListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &v1.GetContactApplicationListRes{
		Applications: out.Applications,
		Page:         req.Page,
		Size:         req.Size,
		Total:        out.Total,
	}, nil
}
