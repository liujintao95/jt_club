package chat

import (
	"context"
	"jt_chat/internal/model"
	"jt_chat/internal/service"

	"jt_chat/api/chat/v1"
)

func (c *ControllerV1) GetHistoryMessage(ctx context.Context, req *v1.GetHistoryMessageReq) (res *v1.GetHistoryMessageRes, err error) {
	var (
		out model.GetHistoryMessageOutput
	)
	out, err = service.Chat().GetHistoryMessage(ctx, model.GetHistoryMessageInput{
		MessageId: req.MessageId,
		ContactId: req.ContactId,
		Page:      req.Page,
		Size:      req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &v1.GetHistoryMessageRes{
		Messages: out.Messages,
		Page:     req.Page,
		Size:     req.Size,
		Total:    out.Total,
	}, err
}
