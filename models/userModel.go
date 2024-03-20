package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username 	 string `gorm:"type:varchar(100);uniqueIndex;not null" json:"username" `
	Email      string `gorm:"type:varchar(100);uniqueIndex;not null" json:"email" `
	Name 		   string `gorm:"type:varchar(100);not null" json:"name" ` 
	Role 		   string `gorm:"type:varchar(100);not null" json:"role"` 
	Password   string `gorm:"type:varchar(100);not null" json:"password"`
	Banks      []Bank `gorm:"foreignKey:owner_id"`
	Products 	 []Product `gorm:"foreignKey:owner_id"`
	Orders     []Order `gorm:"foreignKey:user_id"`
}

//req
type UserLogin struct {
	Username  	string `json:"username" validate:"required"`
	Password 		string `json:"password" validate:"required"`
}

//res
type ResUser struct {
	Username string `json:"username"`
	Name string `json:"name"`
	Email string `json:"email"`
	Role string `json:"role"`
}

type ResUserLog struct {
	Username  	string `json:"username"`
	Name 				string `json:"name"`
	AccessToken string `json:"accessToken"`
}


