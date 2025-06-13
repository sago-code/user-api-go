package models

import (
	"gorm.io/gorm"
)

type User struct {
	FirstName string `gorm:"type:varchar(50);not null"`
	LastName  string `gorm:"type:varchar(50);not null"`
	Age       int    `gorm:"type:tinyint;not null"`
	Email     string `gorm:"type:varchar(100);not null;unique"`
	Phone     string `gorm:"type:varchar(20);not null;unique"`

	gorm.Model
}
