package entity

import "gorm.io/gorm"

type History struct {
	gorm.Model
	Type       string   `gorm:"type:varchar(255)" json:"type"`
	Amount     uint64   `json:"amount"`
	CustomerID uint64   `gorm:"not null" json:"customer_id"`
	Customer   Customer `gorm:"foreignkey:CustomerID;Constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"customer"`
}
