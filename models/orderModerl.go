package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID    uint    `json:"user_id" validate:"required"`
	User      User    `gorm:"foreignkey:UserID"`
	BankID    uint    `json:"bank_id" validate:"required"`
	Bank      Bank    `gorm:"foreignkey:BankID"`
	ProductID uint    `json:"product_id" validate:"required"`
	Product   Product `gorm:"foreignkey:ProductID"`
	Amount    uint64  `json:"amount" validate:"required"`
	Review    string 	`json:"review"`
}