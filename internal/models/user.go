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

type UserGroupMap struct {
	Id        int64
	MapId     string
	Gid       string
	Uid       string
	Ctime     time.Time
	Utime     time.Time
	Deleted   int
	User      User      `db:"user"`
	UserGroup UserGroup `db:"user_group"`
}

type UserContacts struct {
	Id           int64
	Cid          string
	Uid          string
	ContactId    string
	ContactType  int
	ContactNotes string
	LastMsg      string
	LastTime     time.Time
	Ctime        time.Time
	Utime        time.Time
	Deleted      int
	User         User      `db:"user"`
	UserGroup    UserGroup `db:"user_group"`
}

type ContactApplication struct {
	Id          int64
	AppId       string
	Uid         string
	ContactId   string
	ContactType int
	Status      int
	Notice      string
	Ctime       time.Time
	Utime       time.Time
	Deleted     int
	Applicant   User `db:"user"`
}
