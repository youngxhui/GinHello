package model

import (
	"GinHello/initDB"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Roles    []Role `json:"role" gorm:"many2many:roles"`
}

func init() {
	if !initDB.Db.HasTable(User{}) {
		initDB.Db.CreateTable(User{})
	}
}

func (user User) QueryByUsername() User {
	initDB.Db.First(&user, user.Username)
	return user
}

func (user User) Insert() bool {
	initDB.Db.Create(&user)

	if initDB.Db.Error == nil {
		return true
	}
	return false
}
