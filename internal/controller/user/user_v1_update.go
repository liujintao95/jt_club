package user

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"jt_chat/internal/consts"
	"jt_chat/internal/model"
	"jt_chat/internal/service"

	"jt_chat/api/user/v1"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	var (
		out model.UpdateOutput
	)
	out, err = service.User().Update(ctx, model.UpdateInput{
		Name:   req.Name,
		Avatar: req.Avatar,
	})
	if err != nil {
		return nil, err
	}
	res.Uid = out.Uid
	res.Name = out.Name
	res.Avatar = out.Avatar
	res.Email = gconv.String(ctx.Value(consts.CtxUserEmail))
	return res, nil
}
