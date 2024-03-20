package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null" json:"name" validate:"required"`
	Price    uint64 `json:"price" validate:"required"`
	OwnerID  uint   `json:"owner_id" validate:"required"`
	Owner    User   `gorm:"foreignkey:OwnerID"`
	ImageUrl string `json:"imageUrl" validate:"required"`
	Order    []Order `gorm:"foreignkey:product_id"`
}