package initDB

import (
	"GinHello/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open(setting.AppConfig.Database.Type, setting.AppConfig.Database.Url)
	if err != nil {
		log.Panicln("err:", err.Error())
	}
}
