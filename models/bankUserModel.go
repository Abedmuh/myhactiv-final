package models

import "gorm.io/gorm"

type Bank struct {
	gorm.Model
	Name    string `gorm:"type:varchar(100);not null" json:"name" validate:"required"`
	OwnerID uint   `json:"owner_id" validate:"required"`
	Owner   User   `gorm:"foreignkey:OwnerID"`
	Balance uint64 `json:"balance"`
}

