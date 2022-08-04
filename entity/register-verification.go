package entity

import "gorm.io/gorm"

type RegisterVerification struct {
	gorm.Model
	UserName string `gorm:"type:varchar(255);UNIQUE" json:"username"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Password string `gorm:"type:varchar(255)" json:"-"`
}
