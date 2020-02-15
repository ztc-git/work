package model

import (
	"github.com/jinzhu/gorm"
)

type Student struct {
	StudentID string `gorm:"primary_key";"'`
	Name      string `gorm:"size:32;index;not null"`
	ClassID   string   `gorm:"not null"`
	QQ        string
	Address   string `gorm:"not null"`
	Phone     string
	//Families  []*Family `gorm:"many2many:par_child"`
	Major     string    `gorm:"not null;unique"`
	College   string    `gorm:"not null;unique"`
}

type Family struct {
	gorm.Model
	Name     string     `gorm:"size:32;index;not null"`
	Relation string     `gorm:"not null"`
	//Student  []*Student `gorm:"many2many:par_child"`
}

type Admin struct {
	AdminID string   `gorm:"primary_key"`
	Name    string `gorm:"size:32;index;not null"`
}

type Teacher struct {
	TeacherID string `gorm:"primary_key"`
	Name      string `gorm:"size:32;index;not null"`
	ClassID   string   `gorm:"not null"`
	QQ        string
	Phone     string
	Major     string `gorm:"not null;unique"`
}
