// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserContacts is the golang structure of table user_contacts for DAO operations like Where/Data.
type UserContacts struct {
	g.Meta         `orm:"table:user_contacts, do:true"`
	Id             interface{} //
	Cid            interface{} //
	Uid            interface{} //
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
	DeletedAt      *gtime.Time //
	ContactId      interface{} //
	ContactType    interface{} //
	ContactName    interface{} //
	ContactNotes   interface{} //
	LastMsg        interface{} //
	LastTime       *gtime.Time //
	LastWatchMsgId interface{} //
	NewMsgCount    interface{} //
}
