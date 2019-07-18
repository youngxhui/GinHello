package initDB

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/ginhello")
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	Db.SetMaxOpenConns(20)
	Db.SetMaxIdleConns(20)
}
