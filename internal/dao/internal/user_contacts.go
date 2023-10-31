// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserContactsDao is the data access object for table user_contacts.
type UserContactsDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns UserContactsColumns // columns contains all the column names of Table for convenient usage.
}

// UserContactsColumns defines and stores column names for table user_contacts.
type UserContactsColumns struct {
	Id           string //
	Cid          string //
	Uid          string //
	CreatedAt    string //
	UpdatedAt    string //
	DeletedAt    string //
	ContactId    string //
	ContactType  string //
	ContactName  string //
	ContactNotes string //
	LastMsg      string //
	LastTime     string //
	LastMsgId    string //
	NewMsgCount  string //
}

// userContactsColumns holds the columns for table user_contacts.
var userContactsColumns = UserContactsColumns{
	Id:           "id",
	Cid:          "cid",
	Uid:          "uid",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
	ContactId:    "contact_id",
	ContactType:  "contact_type",
	ContactName:  "contact_name",
	ContactNotes: "contact_notes",
	LastMsg:      "last_msg",
	LastTime:     "last_time",
	LastMsgId:    "last_msg_id",
	NewMsgCount:  "new_msg_count",
}

// NewUserContactsDao creates and returns a new DAO object for table data access.
func NewUserContactsDao() *UserContactsDao {
	return &UserContactsDao{
		group:   "default",
		table:   "user_contacts",
		columns: userContactsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserContactsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserContactsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserContactsDao) Columns() UserContactsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserContactsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserContactsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserContactsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
