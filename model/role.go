package model

import (
	"GinHello/initDB"
	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	Name string `json:"name"`
}

func init() {
	if !initDB.Db.HasTable(Role{}) {
		initDB.Db.CreateTable(Role{})
	}
}
