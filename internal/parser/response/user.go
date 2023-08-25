package response

import (
	"database/sql"
	"time"
)

type SignUp struct {
	Email string
}

type SignIn struct {
	Token string
}

type UserInfo struct {
	Uid    string
	Name   string
	Email  string
	Avatar sql.NullString
}

type ContactInfo struct {
	Cid          string
	ContactId    string
	ContactType  int
	ContactNotes string
	LastMsg      string
	LastTime     time.Time
	Name         string
	Avatar       sql.NullString
}
