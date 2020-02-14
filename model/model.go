package model

import "github.com/jinzhu/gorm"

type Student struct {
	StudentID uint64 `gorm:"primary_key"`
	Name string `gorm:"size:32;index;not null"`
	ClassID uint `gorm:"not null"`
	QQ uint64
	Adress string `gorm:"not null"`
	Phone uint64
	Families []*Family `gorm:"many2many:par_child"`
}

type Family struct {
	gorm.Model
	Name string `gorm:"size:32;index;not null"`
	Relation string `gorm:"not null"`
	Student []*Student `gorm:"many2many:par_child"`
}

