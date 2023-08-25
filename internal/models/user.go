package models

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
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

func (user *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}

func (user *User) GetPasswordHash() (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	return string(hash), err
}

type UserGroup struct {
	Id      int64
	Gid     string
	Name    string
	AdminId string
	Notice  string
	Avatar  sql.NullString
	Ctime   time.Time
	Utime   time.Time
	Deleted int
}

type UserContacts struct {
	Id           int64
	Uid          string
	ContactId    string
	ContactType  int
	ContactNotes string
	LastMsg      string
	LastTime     time.Time
	Ctime        time.Time
	Utime        time.Time
	Deleted      int
	User         `db:"user"`
	UserGroup    `db:"user_group"`
}
