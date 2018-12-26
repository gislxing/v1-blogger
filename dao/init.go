package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {

	db, err := sqlx.Open("mysql", "root:1qaz@tcp(127.0.0.1:3306)/blogger?parseTime=true")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(20)
}
