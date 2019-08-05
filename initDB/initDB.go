package initDB

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("mysql", "root:1234@tcp(localhost:3306)/ginhello")
	if err != nil {
		log.Panicln("err:", err.Error())
	}
}
