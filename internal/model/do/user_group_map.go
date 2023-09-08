// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserGroupMap is the golang structure of table user_group_map for DAO operations like Where/Data.
type UserGroupMap struct {
	g.Meta    `orm:"table:user_group_map, do:true"`
	Id        interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt interface{} //
	MapId     interface{} //
	Gid       interface{} //
	Uid       interface{} //
}
