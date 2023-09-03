// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ContactApplication is the golang structure for table contact_application.
type ContactApplication struct {
	Id          uint64      `json:"id"          ` //
	Ctime       *gtime.Time `json:"ctime"       ` //
	Utime       *gtime.Time `json:"utime"       ` //
	Deleted     uint        `json:"deleted"     ` //
	AppId       string      `json:"appId"       ` //
	Uid         string      `json:"uid"         ` //
	ContactId   string      `json:"contactId"   ` //
	ContactType uint        `json:"contactType" ` //
	Status      uint        `json:"status"      ` //
	Notice      string      `json:"notice"      ` //
}
