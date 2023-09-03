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
	res.Applications = out.Applications
	res.Page = out.Page
	res.Size = out.Size
	res.Total = out.Total
	return res, nil
}
