package entity

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Photo   string `gorm:"type:varchar(255)" json:"photo"`
	Balance uint64 `json:"balance"`
	User    User   `gorm:"foreignKey: ID;Constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
