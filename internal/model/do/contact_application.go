// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ContactApplication is the golang structure of table contact_application for DAO operations like Where/Data.
type ContactApplication struct {
	g.Meta      `orm:"table:contact_application, do:true"`
	Id          interface{} //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	DeletedAt   interface{} //
	AppId       interface{} //
	Uid         interface{} //
	ContactId   interface{} //
	ContactType interface{} //
	Status      interface{} //
	Notice      interface{} //
}
