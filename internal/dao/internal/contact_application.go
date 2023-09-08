// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ContactApplicationDao is the data access object for table contact_application.
type ContactApplicationDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns ContactApplicationColumns // columns contains all the column names of Table for convenient usage.
}

// ContactApplicationColumns defines and stores column names for table contact_application.
type ContactApplicationColumns struct {
	Id          string //
	CreatedAt   string //
	UpdatedAt   string //
	DeletedAt   string //
	AppId       string //
	Uid         string //
	ContactId   string //
	ContactType string //
	Status      string //
	Notice      string //
}

// contactApplicationColumns holds the columns for table contact_application.
var contactApplicationColumns = ContactApplicationColumns{
	Id:          "id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
	AppId:       "app_id",
	Uid:         "uid",
	ContactId:   "contact_id",
	ContactType: "contact_type",
	Status:      "status",
	Notice:      "notice",
}

// NewContactApplicationDao creates and returns a new DAO object for table data access.
func NewContactApplicationDao() *ContactApplicationDao {
	return &ContactApplicationDao{
		group:   "default",
		table:   "contact_application",
		columns: contactApplicationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ContactApplicationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ContactApplicationDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ContactApplicationDao) Columns() ContactApplicationColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ContactApplicationDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ContactApplicationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ContactApplicationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
