// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserGroup is the golang structure for table user_group.
type UserGroup struct {
	Id        uint64      `json:"id"        ` //
	Gid       string      `json:"gid"       ` //
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
	DeletedAt *gtime.Time `json:"deletedAt" ` //
	Name      string      `json:"name"      ` //
	AdminId   string      `json:"adminId"   ` //
	Notice    string      `json:"notice"    ` //
	Avatar    string      `json:"avatar"    ` //
}
