package model

import (
	"github.com/jinzhu/gorm"
)

type Project struct {
	gorm.Model
	ProjectName string `gorm:"not null"`
}
