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
	out.Size = in.Size
	out.Page = in.Page
	out.Users = make([]model.UserInfoItem, 0, in.Size)
	err = dao.User.Ctx(ctx).WhereLike(
		dao.User.Columns().Uid, "%"+in.NameOrId+"%",
	).WhereOrLike(
		dao.User.Columns().Name, "%"+in.NameOrId+"%",
	).Page(in.Page, in.Size).ScanAndCount(&out.Users, &out.Total, true)
	return out, err
}

func (s *sUser) GetContactList(ctx context.Context, in model.GetContactListInput) (out model.GetContactListOutput, err error) {
	var (
		uid string
		m   *gdb.Model
	)
	out.Size = in.Size
	out.Page = in.Page
	out.Contacts = make([]model.ContactInfoItem, 0, in.Size)
	uid = gconv.String(ctx.Value(consts.CtxUserId))
	m = dao.UserContacts.Ctx(ctx).WithAll().Where(
		dao.UserContacts.Columns().Uid, uid,
	)
	if in.NameOrId != "" {
		m = m.Where(
			m.Builder().WhereLike(
				dao.UserContacts.Columns().ContactId, "%"+in.NameOrId+"%",
			).WhereOrLike(
				dao.UserContacts.Columns().ContactName, "%"+in.NameOrId+"%",
			).WhereOrLike(
				dao.UserContacts.Columns().ContactNotes, "%"+in.NameOrId+"%",
			),
		)
	}
	err = m.OrderDesc(
		dao.UserContacts.Columns().UpdatedAt,
	).Page(in.Page, in.Size).ScanAndCount(&out.Contacts, &out.Total, true)
	return out, err
}

func (s *sUser) GetContactApplicationList(ctx context.Context, in model.GetContactApplicationListInput) (out model.GetContactApplicationListOutput, err error) {
	var (
		uid  string
		m    *gdb.Model
		subM *gdb.Model
	)
	out.Applications = make([]model.ContactApplicationItem, 0, in.Size)
	out.Size = in.Size
	out.Page = in.Page
	uid = gconv.String(ctx.Value(consts.CtxUserId))
	m = dao.ContactApplication.Ctx(ctx).WithAll().Where(
		dao.ContactApplication.Columns().ContactType, in.ContactType,
	)
	if in.ContactType == consts.ContactsUserType {
		m = m.WhereIn(dao.ContactApplication.Columns().ContactId, uid)
	} else if in.ContactType == consts.ContactsGroupType {
		subM = dao.UserGroup.Ctx(ctx).Fields(dao.UserGroup.Columns().Gid).Where(
			dao.UserGroup.Columns().AdminId, uid,
		)
		m = m.Where(dao.ContactApplication.Columns().ContactId, subM)
	}
	err = m.Page(in.Page, in.Size).ScanAndCount(&out.Applications, &out.Total, true)
	if err != nil || out.Total == 0 {
		return out, err
	}
	return out, err
}

func (s *sUser) CreateContactApplication(ctx context.Context, in model.CreateContactApplicationInput) (out model.CreateContactApplicationOutput, err error) {
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
		ctxUserName  string
		tx           gdb.TX
		application  model.ContactApplicationItem
		contactUser  model.UserInfoItem
		contactGroup model.UserGroupItem
	)
	ctxUserName = gconv.String(ctx.Value(consts.CtxUserName))
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
		err = dao.ContactApplication.Ctx(ctx).WithAll().TX(tx).Where(
			dao.ContactApplication.Columns().AppId, in.AppId,
		).Scan(&application)
		if err != nil {
			return
		}
		if application.ContactType == consts.ContactsUserType {
			err = dao.User.Ctx(ctx).Where(
				dao.User.Columns().Uid, application.ContactId,
			).Scan(&contactUser)
			if err != nil {
				return
			}
			_, err = dao.UserContacts.Ctx(ctx).TX(tx).Data(
				entity.UserContacts{
					Cid:          uuid.New().String(),
					Uid:          application.Uid,
					ContactId:    application.ContactId,
					ContactType:  application.ContactType,
					ContactNotes: application.Notice,
					ContactName:  ctxUserName,
				}, entity.UserContacts{
					Cid:          uuid.New().String(),
					Uid:          application.ContactId,
					ContactId:    application.Uid,
					ContactType:  application.ContactType,
					ContactNotes: in.Notice,
					ContactName:  application.User.Name,
				},
			).Insert()
		} else if application.ContactType == consts.ContactsGroupType {
			err = dao.UserGroup.Ctx(ctx).Where(
				dao.UserGroup.Columns().Gid, application.ContactId,
			).Scan(&contactGroup)
			if err != nil {
				return
			}
			_, err = dao.UserContacts.Ctx(ctx).TX(tx).Data(entity.UserContacts{
				Cid:          uuid.New().String(),
				Uid:          application.Uid,
				ContactId:    contactGroup.Gid,
				ContactType:  application.ContactType,
				ContactNotes: application.Notice,
				ContactName:  contactGroup.Name,
			}).Insert()
			_, err = dao.UserGroupMap.Ctx(ctx).TX(tx).Data(entity.UserGroupMap{
				MapId: uuid.New().String(),
				Uid:   contactUser.Uid,
				Gid:   contactGroup.Gid,
			}).Insert()
		}
	}
	return
}

func (s *sUser) CreateUserGroup(ctx context.Context, in model.CreateUserGroupInput) (out model.CreateUserGroupOutput, err error) {
	var (
		newGroup entity.UserGroup
	)
	newGroup = entity.UserGroup{
		Gid:     uuid.New().String(),
		Name:    in.Name,
		Avatar:  in.Avatar,
		Notice:  in.Notice,
		AdminId: gconv.String(ctx.Value(consts.CtxUserId)),
	}
	_, err = dao.UserGroup.Ctx(ctx).Data(newGroup).Insert()
	if err != nil {
		return
	}
	out.Gid = newGroup.Gid
	return
}

func (s *sUser) UpdateUserGroup(ctx context.Context, in model.UpdateUserGroupInput) (out model.UpdateUserGroupOutput, err error) {
	var (
		mapCount int
		group    entity.UserGroup
		m        *gdb.Model
	)
	err = dao.UserGroup.Ctx(ctx).Where(dao.UserGroup.Columns().Gid, in.Gid).Scan(&group)
	if err != nil {
		return
	}
	if group.AdminId != gconv.String(ctx.Value(consts.CtxUserId)) {
		err = fmt.Errorf("only administrators can modify")
		return
	}

	m = dao.UserGroup.Ctx(ctx)
	if in.Name != "" {
		m = m.Data(dao.UserGroup.Columns().Name, in.Name)
	}
	if in.Avatar != "" {
		m = m.Data(dao.UserGroup.Columns().Avatar, in.Avatar)
	}
	if in.Notice != "" {
		m = m.Data(dao.UserGroup.Columns().Notice, in.Notice)
	}
	if in.AdminId != "" {
		mapCount, err = dao.UserGroupMap.Ctx(ctx).Where(
			dao.UserGroupMap.Columns().Gid, in.Gid,
		).Where(
			dao.UserGroupMap.Columns().Uid, in.AdminId,
		).Count()
		if err != nil {
			return
		}
		if mapCount == 0 {
			err = fmt.Errorf("administrators can only be group members")
			return
		}
		m = m.Data(dao.UserGroup.Columns().AdminId, in.AdminId)
	}
	_, err = m.Where(dao.UserGroup.Columns().Gid, in.Gid).Update()
	return
}
