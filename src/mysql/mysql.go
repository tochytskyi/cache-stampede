package mysql

import (
	"database/sql"

	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	MysqlDb *sql.DB
)

func Init() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(treatfield_db:3306)/treatfield")

	if err != nil {
		log.Println(err)
	}

	return db
}
