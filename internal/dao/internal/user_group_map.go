// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserGroupMapDao is the data access object for table user_group_map.
type UserGroupMapDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns UserGroupMapColumns // columns contains all the column names of Table for convenient usage.
}

// UserGroupMapColumns defines and stores column names for table user_group_map.
type UserGroupMapColumns struct {
	Id      string //
	Ctime   string //
	Utime   string //
	Deleted string //
	MapId   string //
	Gid     string //
	Uid     string //
}

// userGroupMapColumns holds the columns for table user_group_map.
var userGroupMapColumns = UserGroupMapColumns{
	Id:      "id",
	Ctime:   "ctime",
	Utime:   "utime",
	Deleted: "deleted",
	MapId:   "map_id",
	Gid:     "gid",
	Uid:     "uid",
}

// NewUserGroupMapDao creates and returns a new DAO object for table data access.
func NewUserGroupMapDao() *UserGroupMapDao {
	return &UserGroupMapDao{
		group:   "default",
		table:   "user_group_map",
		columns: userGroupMapColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserGroupMapDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserGroupMapDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserGroupMapDao) Columns() UserGroupMapColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserGroupMapDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserGroupMapDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserGroupMapDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
