package user

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/google/uuid"
	"jt_chat/internal/consts"
	"jt_chat/internal/dao"
	"jt_chat/internal/model"
	"jt_chat/internal/model/entity"
	"jt_chat/internal/service"
	"jt_chat/utility"
	"regexp"
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
		maxUid      float64
		uid         string
		count       int
		newUser     entity.User
		emailRegexp *regexp.Regexp
	)
	emailRegexp = regexp.MustCompile(consts.EmailRegex)
	if !emailRegexp.MatchString(in.Email) {
		err = fmt.Errorf("email format error")
		return
	}
	maxUid, err = dao.User.Ctx(ctx).Max(dao.User.Columns().Uid)
	uid = gconv.String(gconv.Int(maxUid) + 1)
	if err != nil {
		return
	}
	newUser = entity.User{
		Uid:   uid,
		Name:  in.Name,
		Email: in.Email,
	}
	count, err = dao.User.Ctx(ctx).Where(dao.User.Columns().Email, in.Email).Count()
	if err != nil {
		return
	}
	if count != 0 {
		err = fmt.Errorf("email already in use")
		return
	}
	newUser.Password = utility.EncryptPassword(in.Password, consts.Salt)
	_, err = dao.User.Ctx(ctx).Data(newUser).Insert()
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
	).Data(in).Update()
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
	m = dao.User.Ctx(ctx).WhereLike(
		dao.User.Columns().Uid, "%"+in.NameOrId+"%",
	).WhereOrLike(
		dao.User.Columns().Name, "%"+in.NameOrId+"%",
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
		uid    string
		m      *gdb.Model
		userM  *gdb.Model
		groupM *gdb.Model
	)
	out.Size = in.Size
	out.Page = in.Page
	uid = gconv.String(ctx.Value(consts.CtxUserId))
	m = dao.UserContacts.Ctx(ctx).WithAll().Where(
		dao.UserContacts.Columns().Uid, uid,
	)
	if in.NameOrId != "" {
		userM = dao.User.Ctx(ctx).Fields(dao.User.Columns().Uid).WhereLike(
			dao.User.Columns().Uid, "%"+in.NameOrId+"%",
		).WhereOrLike(
			dao.User.Columns().Name, "%"+in.NameOrId+"%",
		)
		groupM = dao.UserGroup.Ctx(ctx).Fields(dao.UserGroup.Columns().Gid).WhereLike(
			dao.UserGroup.Columns().Gid, "%"+in.NameOrId+"%",
		).WhereOrLike(
			dao.UserGroup.Columns().Name, "%"+in.NameOrId+"%",
		)
		m = m.WhereIn(
			dao.UserContacts.Columns().ContactId, userM,
		).WhereOrIn(
			dao.UserContacts.Columns().ContactId, groupM,
		)
	}
	m = m.OrderDesc(
		dao.UserContacts.Columns().UpdatedAt,
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
	m = dao.ContactApplication.Ctx(ctx).WithAll().Where(
		dao.ContactApplication.Columns().ContactId, uid,
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
		count          int
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
	count, err = dao.ContactApplication.Ctx(ctx).Where(
		dao.ContactApplication.Columns().Uid, newApplication.Uid,
	).Where(
		dao.ContactApplication.Columns().ContactId, newApplication.ContactId,
	).Where(
		dao.ContactApplication.Columns().Status, consts.ApplicationWaitStatus,
	).Count()
	if err != nil {
		return
	}
	if count != 0 {
		err = fmt.Errorf("the request has been submitted, please do not resubmit")
		return
	}
	_, err = dao.ContactApplication.Ctx(ctx).Data(newApplication).Insert()
	if err != nil {
		return
	}
	out.AppID = appId
	return
}

func (s *sUser) UpdateContactApplication(ctx context.Context, in model.UpdateContactApplicationInput) (out model.UpdateContactApplicationOutput, err error) {
	var (
		tx          gdb.TX
		application entity.ContactApplication
	)
	tx, err = g.DB().Begin(ctx)
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	_, err = dao.ContactApplication.Ctx(ctx).TX(tx).Data(
		dao.ContactApplication.Columns().Status, in.Status,
	).Where(
		dao.ContactApplication.Columns().AppId, in.AppId,
	).Update()
	if in.Status == consts.ApplicationAgreeStatus {
		err = dao.ContactApplication.Ctx(ctx).TX(tx).Where(
			dao.ContactApplication.Columns().AppId, in.AppId,
		).Scan(&application)
		if err != nil {
			return
		}
		_, err = dao.UserContacts.Ctx(ctx).TX(tx).Data(entity.UserContacts{
			Cid:         uuid.New().String(),
			Uid:         application.Uid,
			ContactId:   application.ContactId,
			ContactType: application.ContactType,
		}).Insert()
		_, err = dao.UserContacts.Ctx(ctx).TX(tx).Data(entity.UserContacts{
			Cid:         uuid.New().String(),
			Uid:         application.ContactId,
			ContactId:   application.Uid,
			ContactType: application.ContactType,
		}).Insert()
	}

	return
}