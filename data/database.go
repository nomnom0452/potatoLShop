package data

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func Connect() {
	var err error

	Db, err = sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/absurd?parseTime=true")

	if err != nil {
		panic(err.Error())
	}
}
