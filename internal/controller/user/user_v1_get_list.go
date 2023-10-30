package user

import (
	"context"
	"jt_chat/internal/model"
	"jt_chat/internal/service"

	"jt_chat/api/user/v1"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	var (
		out model.GetListOutput
	)
	out, err = service.User().GetList(ctx, model.GetListInput{
		NameOrId: req.NameOrId,
		Page:     req.Page,
		Size:     req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &v1.GetListRes{
		Users: out.Users,
		Page:  req.Page,
		Size:  req.Size,
		Total: out.Total,
	}, nil
}
