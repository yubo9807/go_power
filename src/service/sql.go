package service

import (
	"server/configs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 连接数据库
func DBConnect() *sqlx.DB {
	db, err := sqlx.Open("mysql", configs.SqlSecret)
	if err != nil {
		panic(err.Error())
	}
	return db
}
