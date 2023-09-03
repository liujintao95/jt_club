package user

import (
	"context"
	"jt_chat/internal/model"
	"jt_chat/internal/service"

	"jt_chat/api/user/v1"
)

func (c *ControllerV1) GetContactList(ctx context.Context, req *v1.GetContactListReq) (res *v1.GetContactListRes, err error) {
	var (
		out model.GetContactListOutput
	)
	out, err = service.User().GetContactList(ctx, model.GetContactListInput{
		NameOrUid: req.NameOrId,
		Page:      req.Page,
		Size:      req.Size,
	})
	if err != nil {
		return nil, err
	}
	res.Page = out.Page
	res.Size = out.Size
	res.Total = out.Total
	res.Contacts = out.Contacts
	return res, nil
}
