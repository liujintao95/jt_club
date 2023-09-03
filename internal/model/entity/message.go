// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Message is the golang structure for table message.
type Message struct {
	Id           uint64      `json:"id"           ` //
	MessageId    string      `json:"messageId"    ` //
	Ctime        *gtime.Time `json:"ctime"        ` //
	Utime        *gtime.Time `json:"utime"        ` //
	Deleted      int         `json:"deleted"      ` //
	Avatar       string      `json:"avatar"       ` //
	FromUsername string      `json:"fromUsername" ` //
	From         string      `json:"from"         ` //
	To           string      `json:"to"           ` //
	Content      string      `json:"content"      ` //
	ContentType  int         `json:"contentType"  ` //
	Type         string      `json:"type"         ` //
	MessageType  int         `json:"messageType"  ` //
	Url          string      `json:"url"          ` //
	FileSuffix   string      `json:"fileSuffix"   ` //
	FilePath     string      `json:"filePath"     ` //
}
