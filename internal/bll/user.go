package bll

import (
	"JT_CLUB/conf"
	"JT_CLUB/internal/constant"
	"JT_CLUB/internal/dal"
	"JT_CLUB/internal/models"
	"JT_CLUB/internal/parser/request"
	"JT_CLUB/internal/parser/response"
	"JT_CLUB/pkg/cache"
	"JT_CLUB/pkg/db"
	"fmt"
	"github.com/google/uuid"
)

func Login(login *request.SignInRequest) (string, error) {
	var (
		user models.User
		ok   bool
		err  error
	)
	user, err = dal.GetUserByEmail(db.Conn, login.Account)
	if err != nil {
		return "", fmt.Errorf("select user: %w", err)
	}
	if ok = user.ComparePassword(login.Password); !ok {
		return "", fmt.Errorf("user password error")
	} else {
		token := uuid.New().String()

		cache.Cache.Set(token, user, conf.DefaultDuration)
		return token, nil
	}
}

func CreateUser(user *request.SignUpRequest) (string, error) {
	var (
		currentUser  models.User
		passwordHash string
		err          error
		uid          = uuid.New().String()
		userModel    = &models.User{
			Uid:      uid,
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		}
	)
	currentUser, _ = dal.GetUserByEmail(db.Conn, user.Email)
	if currentUser.Uid != "" {
		return "", fmt.Errorf("%s already exists", user.Email)
	}
	passwordHash, err = userModel.GetPasswordHash()
	if err != nil {
		return "", fmt.Errorf("get password: %w", err)
	}
	err = dal.SetUser(db.Conn, userModel, passwordHash)
	if err != nil {
		return "", fmt.Errorf("create user: %w", err)
	}
	return uid, nil
}

func SelectUser(currentUser *models.User, query string, isContact bool) ([]*response.UserInfo, error) {
	var (
		result []*response.UserInfo
		users  []*models.User
		err    error
	)
	if isContact {
		users, err = dal.SelectContacts(db.Conn, currentUser, query)
	} else {
		users, err = dal.SelectUsers(db.Conn, query)
	}
	if err != nil {
		return nil, fmt.Errorf("select user: %w", err)
	}
	for _, user := range users {
		result = append(result, &response.UserInfo{
			Uid:    user.Uid,
			Name:   user.Name,
			Email:  user.Email,
			Avatar: user.Avatar,
		})
	}
	return result, nil
}

func GetContactList(user *models.User) ([]*response.ContactInfo, error) {
	var (
		result   []*response.ContactInfo
		contacts []*models.UserContacts
		err      error
	)
	contacts, err = dal.GetUserContacts(db.Conn, user.Uid)
	if err != nil {
		return nil, fmt.Errorf("get user contacts: %w", err)
	}
	for _, contact := range contacts {
		contactInfo := &response.ContactInfo{
			Cid:          contact.Cid,
			ContactId:    contact.ContactId,
			ContactType:  contact.ContactType,
			ContactNotes: contact.ContactNotes,
			LastMsg:      contact.LastMsg,
			LastTime:     contact.LastTime,
		}
		if contact.ContactType == constant.ContactsUserType {
			contactInfo.Name = contact.User.Name
			contactInfo.Avatar = contact.User.Avatar
		} else {
			contactInfo.Name = contact.UserGroup.Name
			contactInfo.Avatar = contact.UserGroup.Avatar
		}
		result = append(result, contactInfo)
	}
	return result, nil
}

func CreateContactApplication(currentUser *models.User, application request.ContactApplication) (string, error) {
	var (
		err          error
		appId        = uuid.New().String()
		requestModel = &models.ContactApplication{
			AppId:       appId,
			Uid:         currentUser.Uid,
			ContactId:   application.ContactId,
			ContactType: application.ContactType,
			Status:      constant.RequestWaitStatus,
			Notice:      application.Notice,
		}
	)
	err = dal.SetContactApplication(db.Conn, requestModel)
	if err != nil {
		return "", fmt.Errorf("create contacts_request: %w", err)
	}
	return appId, nil
}

func UpdateContactApplicationStatus(confirmInfo request.ContactConfirm) error {
	var (
		application models.ContactApplication
		err         error
	)
	application, err = dal.GetContactApplicationByAppId(db.Conn, confirmInfo.AppId)
	if err != nil {
		return fmt.Errorf("get application by app_id: %w", err)
	}
	if application.Status != constant.RequestWaitStatus {
		return fmt.Errorf("application(%s) status(%s) error", application.AppId, application.Status)
	}
	err = dal.UpdateContactApplicationStatus(db.Conn, confirmInfo.AppId, confirmInfo.Status)
	if err != nil {
		return fmt.Errorf("update user status: %w", err)
	}
	return nil
}

func GetContactApplicationConfirmList(currentUser *models.User) ([]*response.ApplicationInfo, error) {
	var (
		result       []*response.ApplicationInfo
		applications []*models.ContactApplication
		err          error
	)
	applications, err = dal.GetContactApplicationConfirmList(db.Conn, currentUser.Uid)
	if err != nil {
		return nil, fmt.Errorf("get confirm list: %w", err)
	}
	for _, application := range applications {
		result = append(result, &response.ApplicationInfo{
			AppId:  application.AppId,
			Uid:    application.Uid,
			Name:   application.Applicant.Name,
			Avatar: application.Applicant.Avatar,
			Notice: application.Notice,
			Status: application.Status,
		})
	}
	return result, nil
}
