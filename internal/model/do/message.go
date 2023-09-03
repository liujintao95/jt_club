// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Message is the golang structure of table message for DAO operations like Where/Data.
type Message struct {
	g.Meta       `orm:"table:message, do:true"`
	Id           interface{} //
	MessageId    interface{} //
	Ctime        *gtime.Time //
	Utime        *gtime.Time //
	Deleted      interface{} //
	Avatar       interface{} //
	FromUsername interface{} //
	From         interface{} //
	To           interface{} //
	Content      interface{} //
	ContentType  interface{} //
	Type         interface{} //
	MessageType  interface{} //
	Url          interface{} //
	FileSuffix   interface{} //
	FilePath     interface{} //
}
