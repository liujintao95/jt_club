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
		NameOrUid: req.NameOrId,
		Page:      req.Page,
		Size:      req.Size,
	})
	if err != nil {
		return nil, err
	}
	res.Users = out.Users
	res.Page = out.Page
	res.Size = out.Size
	res.Total = out.Total
	return res, nil
}
