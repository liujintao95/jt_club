package chat

import (
	"context"
	"jt_chat/internal/model"
	"jt_chat/internal/service"

	"jt_chat/api/chat/v1"
)

func (c *ControllerV1) GetNewMessage(ctx context.Context, req *v1.GetNewMessageReq) (res *v1.GetNewMessageRes, err error) {
	var (
		out model.GetNewMessageOutput
	)
	out, err = service.Chat().GetNewMessage(ctx, model.GetNewMessageInput{
		MessageId: req.MessageId,
		ContactId: req.ContactId,
		Page:      req.Page,
		Size:      req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &v1.GetNewMessageRes{
		Messages: out.Messages,
		Page:     req.Page,
		Size:     req.Size,
		Total:    out.Total,
	}, err
}
