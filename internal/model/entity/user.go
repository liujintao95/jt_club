// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        uint64      `json:"id"        ` //
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
	DeletedAt uint        `json:"deletedAt" ` //
	Uid       string      `json:"uid"       ` //
	Name      string      `json:"name"      ` //
	Email     string      `json:"email"     ` //
	Password  string      `json:"password"  ` //
	Avatar    string      `json:"avatar"    ` //
}
