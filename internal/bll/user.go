package bll

import (
	"JT_CLUB/conf"
	"JT_CLUB/internal/dal"
	"JT_CLUB/internal/models"
	"JT_CLUB/internal/parser/request"
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
	err = dal.SaveUser(db.Conn, userModel, passwordHash)
	if err != nil {
		return "", fmt.Errorf("insert user: %w", err)
	}
	return uid, err
}

func GetContactList(user *models.User) {

}
