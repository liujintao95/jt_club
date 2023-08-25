package db

import (
	"JT_CLUB/conf"
	"JT_CLUB/pkg/log"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

type Session interface {
	Get(any, string, ...any) error
	Select(any, string, ...any) error
	MustExec(string, ...any) sql.Result
	Exec(string, ...any) (sql.Result, error)
	SelectContext(context.Context, any, string, ...any) error
	ExecContext(context.Context, string, ...any) (sql.Result, error)
	MustExecContext(context.Context, string, ...any) sql.Result
}

var Conn *sqlx.DB

func InitDataBase() {
	var (
		err error
		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			conf.Config.MySQL.User,
			conf.Config.MySQL.Password,
			conf.Config.MySQL.Host,
			conf.Config.MySQL.Port,
			conf.Config.MySQL.Name,
		)
	)
	// 也可以使用MustConnect连接不成功就panic
	Conn, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}
	Conn.SetMaxOpenConns(20)
	Conn.SetMaxIdleConns(10)
	return
}

func Run(txFunc func(*sqlx.Tx) error) error {
	tx, err := Conn.Beginx()
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			log.Logger.Error("run transaction:" + err.Error())
			_ = tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
			log.Logger.Error("commit transaction:" + err.Error())
		}
	}()
	err = txFunc(tx)
	return err
}
