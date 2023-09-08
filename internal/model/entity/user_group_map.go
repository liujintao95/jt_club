// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserGroupMap is the golang structure for table user_group_map.
type UserGroupMap struct {
	Id        uint64      `json:"id"        ` //
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
	DeletedAt uint        `json:"deletedAt" ` //
	MapId     string      `json:"mapId"     ` //
	Gid       string      `json:"gid"       ` //
	Uid       string      `json:"uid"       ` //
}
