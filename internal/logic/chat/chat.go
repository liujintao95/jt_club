package chat

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"
	"jt_chat/internal/consts"
	"jt_chat/internal/dao"
	"jt_chat/internal/model"
	"jt_chat/internal/model/entity"
	"jt_chat/internal/service"
)

type sChat struct{}

func init() {
	service.RegisterChat(New())
}
func New() *sChat {
	return &sChat{}
}

func (s *sChat) GetHistoryMessage(ctx context.Context, in model.GetHistoryMessageInput) (out model.GetHistoryMessageOutput, err error) {
	var (
		msg entity.Message
		uid string
		m   *gdb.Model
	)
	uid = gconv.String(ctx.Value(consts.CtxUserId))
	if in.MessageId != "" {
		err = dao.Message.Ctx(ctx).Where(dao.Message.Columns().MessageId, in.MessageId).Scan(&msg)
		if err != nil {
			return out, err
		}
		m = dao.Message.Ctx(ctx).WhereLT(dao.Message.Columns().Id, msg.Id)
	} else {
		m = dao.Message.Ctx(ctx)
	}
	err = m.Where(
		m.Builder().Where(
			dao.Message.Columns().From, uid,
		).Where(
			dao.Message.Columns().To, in.ContactId,
		),
	).WhereOr(
		m.Builder().Where(
			dao.Message.Columns().From, in.ContactId,
		).Where(
			dao.Message.Columns().To, uid,
		),
	).OrderDesc(
		dao.Message.Columns().CreatedAt,
	).Limit(
		in.Page, in.Size,
	).ScanAndCount(
		&out.Messages, &out.Total, true,
	)
	return out, err
}

func (s *sChat) GetNewMessage(ctx context.Context, in model.GetNewMessageInput) (out model.GetNewMessageOutput, err error) {
	var (
		msg entity.Message
		uid string
		m   *gdb.Model
	)
	uid = gconv.String(ctx.Value(consts.CtxUserId))
	err = dao.Message.Ctx(ctx).Where(dao.Message.Columns().MessageId, in.MessageId).Scan(&msg)
	if err != nil {
		return out, err
	}
	m = dao.Message.Ctx(ctx).WhereGT(dao.Message.Columns().Id, msg.Id)
	err = m.Where(
		m.Builder().Where(
			dao.Message.Columns().From, uid,
		).Where(
			dao.Message.Columns().To, in.ContactId,
		),
	).WhereOr(
		m.Builder().Where(
			dao.Message.Columns().From, in.ContactId,
		).Where(
			dao.Message.Columns().To, uid,
		),
	).OrderAsc(
		dao.Message.Columns().CreatedAt,
	).Limit(
		in.Page, in.Size,
	).ScanAndCount(
		&out.Messages, &out.Total, true,
	)
	return out, err
}
