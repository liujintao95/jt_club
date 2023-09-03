// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MessageDao is the data access object for table message.
type MessageDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns MessageColumns // columns contains all the column names of Table for convenient usage.
}

// MessageColumns defines and stores column names for table message.
type MessageColumns struct {
	Id           string //
	MessageId    string //
	Ctime        string //
	Utime        string //
	Deleted      string //
	Avatar       string //
	FromUsername string //
	From         string //
	To           string //
	Content      string //
	ContentType  string //
	Type         string //
	MessageType  string //
	Url          string //
	FileSuffix   string //
	FilePath     string //
}

// messageColumns holds the columns for table message.
var messageColumns = MessageColumns{
	Id:           "id",
	MessageId:    "message_id",
	Ctime:        "ctime",
	Utime:        "utime",
	Deleted:      "deleted",
	Avatar:       "avatar",
	FromUsername: "from_username",
	From:         "from",
	To:           "to",
	Content:      "content",
	ContentType:  "content_type",
	Type:         "type",
	MessageType:  "message_type",
	Url:          "url",
	FileSuffix:   "file_suffix",
	FilePath:     "file_path",
}

// NewMessageDao creates and returns a new DAO object for table data access.
func NewMessageDao() *MessageDao {
	return &MessageDao{
		group:   "default",
		table:   "message",
		columns: messageColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MessageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MessageDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MessageDao) Columns() MessageColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MessageDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MessageDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MessageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
