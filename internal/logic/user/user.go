package user

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/google/uuid"
	"jt_chat/internal/consts"
	"jt_chat/internal/dao"
	"jt_chat/internal/model"
	"jt_chat/internal/model/entity"
	"jt_chat/internal/service"
	"jt_chat/utility"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}
func New() *sUser {
	return &sUser{}
}

func (s *sUser) Register(ctx context.Context, in model.RegisterInput) (out model.RegisterOutput, err error) {
	var (
		currentUser entity.User
		uid         string
		newUser     entity.User
	)
	uid = uuid.New().String()
	newUser = entity.User{
		Uid:   uid,
		Name:  in.Name,
		Email: in.Email,
	}
	err = dao.User.Ctx(ctx).Where(dao.User.Columns().Email, in.Email).Scan(&currentUser)
	if err != nil {
		return
	}
	newUser.Password = utility.EncryptPassword(in.Password, consts.Salt)
	_, err = dao.User.Ctx(ctx).Data(newUser).InsertAndGetId()
	if err != nil {
		return
	}
	out.Id = uid
	return
}

func (s *sUser) Update(ctx context.Context, in model.UpdateInput) (out model.UpdateOutput, err error) {
	var (
		uid string
	)
	uid = gconv.String(ctx.Value(consts.CtxUserId))
	_, err = dao.User.Ctx(ctx).Where(
		dao.User.Columns().Uid, uid,
	).Where(
		dao.User.Columns().Deleted, false,
	).Update(in)
	if err != nil {
		return
	}
	return model.UpdateOutput{
		Uid:    uid,
		Name:   in.Name,
		Avatar: in.Avatar,
	}, err
}

func (s *sUser) GetList(ctx context.Context, in model.GetListInput) (out model.GetListOutput, err error) {
	var (
		m *gdb.Model
	)
	out.Size = in.Size
	out.Page = in.Page
	m = dao.User.Ctx(ctx).Where(
		dao.User.Columns().Deleted, false,
	).WhereLike(
		dao.User.Columns().Uid, "%"+in.NameOrUid+"%",
	).WhereOrLike(
		dao.User.Columns().Name, "%"+in.NameOrUid+"%",
	).Page(in.Page, in.Size)
	out.Total, err = m.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}
	out.Users = make([]model.UserInfoItem, 0, in.Size)
	err = m.Scan(&out.Users)
	return out, err
}

func (s *sUser) GetContactList(ctx context.Context, in model.GetContactListInput) (out model.GetContactListOutput, err error) {
	var (
		uid string
		m   *gdb.Model
	)
	out.Size = in.Size
	out.Page = in.Page
	uid = gconv.String(ctx.Value(consts.CtxUserId))
	m = dao.UserContacts.Ctx(ctx).Where(
		dao.UserContacts.Columns().Deleted, false,
	).With(model.UserInfoItem{}).With(model.UserGroupItem{}).Where(
		dao.UserContacts.Columns().Uid, uid,
	).Where(
		dao.User.Columns().Deleted, false,
	).Where(
		dao.UserGroup.Columns().Deleted, false,
	)
	if in.NameOrUid != "" {
		m = m.Where(
			dao.User.Columns().Uid, in.NameOrUid,
		).WhereOr(
			dao.User.Columns().Name, in.NameOrUid,
		).WhereOr(
			dao.UserGroup.Columns().Name, in.NameOrUid,
		).WhereOr(
			dao.UserGroup.Columns().Gid, in.NameOrUid,
		)
	}
	m = m.OrderDesc(
		dao.UserContacts.Columns().Utime,
	).Page(in.Page, in.Size)
	out.Total, err = m.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}
	out.Contacts = make([]model.ContactInfoItem, 0, in.Size)
	err = m.Scan(&out.Contacts)
	return out, err
}

func (s *sUser) GetContactApplicationList(ctx context.Context, in model.GetContactApplicationListInput) (out model.GetContactApplicationListOutput, err error) {
	var (
		uid string
		m   *gdb.Model
	)
	out.Size = in.Size
	out.Page = in.Page
	uid = gconv.String(ctx.Value(consts.CtxUserId))
	m = dao.ContactApplication.Ctx(ctx).Where(
		dao.ContactApplication.Columns().ContactId, uid,
	).Where(
		dao.ContactApplication.Columns().Deleted, false,
	).Page(in.Page, in.Size)
	out.Total, err = m.Count()
	if err != nil || out.Total == 0 {
		return out, err
	}
	out.Applications = make([]model.ContactApplicationItem, 0, in.Size)
	err = m.Scan(&out.Applications)
	return out, err
}

func (s *sUser) SetContactApplication(ctx context.Context, in model.SetContactApplicationInput) (out model.SetContactApplicationOutput, err error) {
	var (
		appId          string
		newApplication entity.ContactApplication
	)
	appId = uuid.New().String()
	newApplication = entity.ContactApplication{
		AppId:       appId,
		Uid:         gconv.String(ctx.Value(consts.CtxUserId)),
		ContactId:   in.ContactId,
		ContactType: in.ContactType,
		Notice:      in.Notice,
		Status:      consts.ApplicationWaitStatus,
	}
	_, err = dao.ContactApplication.Ctx(ctx).Data(newApplication).InsertAndGetId()
	if err != nil {
		return
	}
	out.AppID = appId
	return
}

func (s *sUser) UpdateContactApplication(ctx context.Context, in model.UpdateContactApplicationInput) (out model.UpdateContactApplicationOutput, err error) {
	_, err = dao.ContactApplication.Ctx(ctx).Where(
		dao.ContactApplication.Columns().AppId, in.AppId,
	).Where(
		dao.ContactApplication.Columns().Deleted, false,
	).Update(dao.ContactApplication.Columns().Status, in.Status)
	return
}
