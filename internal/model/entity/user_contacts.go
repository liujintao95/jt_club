// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserContacts is the golang structure for table user_contacts.
type UserContacts struct {
	Id             uint64      `json:"id"             ` //
	Cid            string      `json:"cid"            ` //
	Uid            string      `json:"uid"            ` //
	CreatedAt      *gtime.Time `json:"createdAt"      ` //
	UpdatedAt      *gtime.Time `json:"updatedAt"      ` //
	DeletedAt      *gtime.Time `json:"deletedAt"      ` //
	ContactId      string      `json:"contactId"      ` //
	ContactType    uint        `json:"contactType"    ` //
	ContactName    string      `json:"contactName"    ` //
	ContactNotes   string      `json:"contactNotes"   ` //
	LastMsg        string      `json:"lastMsg"        ` //
	LastTime       *gtime.Time `json:"lastTime"       ` //
	LastWatchMsgId string      `json:"lastWatchMsgId" ` //
	NewMsgCount    uint        `json:"newMsgCount"    ` //
}
