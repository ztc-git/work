package initDB

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Token string	`gorm:"not null;primary_key;"`
	Username string `gorm:"not null;"`
	Sex string `gorm:"not null"`
}

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("mysql", "root:210377091ztc@(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	table := Db.HasTable(User{})
	if !table {
		Db.CreateTable(User{})
	}
}

func TokenExist(signedToken string) bool {
	t := new(User)
	Db.Find(&t, "Token=?", signedToken)
	if t.Token == ""{
		return false
	}else {
		return true
	}
}