package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Username string `gorm:"type:varchar(255);UNIQUE" json:"username"`
	Password string `gorm:"->;<-not null" json:"-"`
	Role     string `gorm:"type:varchar(255)" json:"role"`
	Token    string `gorm:"-" json:"token,omitempty"`
}
