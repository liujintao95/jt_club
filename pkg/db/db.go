package db

import (
	"JT_CLUB/conf"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Conn *sqlx.DB

func InitDataBase() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Config.MySQL.User,
		conf.Config.MySQL.Password,
		conf.Config.MySQL.Host,
		conf.Config.MySQL.Port,
		conf.Config.MySQL.Name,
	)
	// 也可以使用MustConnect连接不成功就panic
	Conn, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}
	Conn.SetMaxOpenConns(20)
	Conn.SetMaxIdleConns(10)
	return
}
