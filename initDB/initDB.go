package initDB

import (
	"app/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Token  string `gorm:"not null;unique;"`
	UserID string `gorm:"not null;primary_key"`
	Roles  string `gorm:"not null;primary_key"`
}

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("mysql", "root:*****@(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	table := Db.HasTable(User{})
	if !table {
		Db.CreateTable(User{})
	}
	if !Db.HasTable(model.Student{}) {
		Db.CreateTable(model.Student{})
	}
	if !Db.HasTable(model.Teacher{}) {
		Db.CreateTable(model.Teacher{})
	}
	if !Db.HasTable(model.Admin{}) {
		Db.CreateTable(model.Admin{})
	}
	if !Db.HasTable(model.Family{}) {
		Db.CreateTable(model.Family{})
	}
}

func TokenExist(signedToken string) bool {
	t := new(User)
	Db.Find(&t, "Token=?", signedToken)
	if t.Token == "" {
		return false
	} else {
		return true
	}
}
