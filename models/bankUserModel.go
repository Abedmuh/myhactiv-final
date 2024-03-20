package models

import "gorm.io/gorm"

type Bank struct {
	gorm.Model
	Name    string `gorm:"type:varchar(100);not null" json:"name" validate:"required"`
	OwnerID uint   `json:"owner_id"`
	Owner   User   `gorm:"foreignkey:OwnerID"`
	Balance uint64 `json:"balance"`
	Orders []Order `gorm:"foreignkey:bank_id"`
}

//request 
type ReqBank struct {
	Name    string `json:"name" validate:"required"`
	OwnerID uint   `json:"owner_id" validate:"required"`
	Balance uint64 `json:"balance"`
}

