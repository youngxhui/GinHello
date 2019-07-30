package model

import (
	"GinHello/initDB"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Content string
}

func init() {
	table := initDB.Db.HasTable(Comment{})
	if !table {
		initDB.Db.CreateTable(Comment{})
	}
}
