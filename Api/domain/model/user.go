package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserId   string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Mail     string
}
