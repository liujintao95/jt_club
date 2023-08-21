package models

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Users struct {
	Id       int64
	Uid      string
	Name     string
	Email    string
	Password string
	Avatar   sql.NullString
	Ctime    time.Time
	Utime    time.Time
	Deleted  int
}

func (user *Users) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}

func (user *Users) GetPasswordHash() (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	return string(hash), err
}
