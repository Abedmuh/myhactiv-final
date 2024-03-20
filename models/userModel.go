package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username 	 string `gorm:"type:varchar(100);uniqueIndex;not null" json:"username" validate:"required"`
	Email      string `gorm:"type:varchar(100);uniqueIndex;not null" json:"email" validate:"required,email"`
	Name 		   string `gorm:"type:varchar(100);not null" json:"name" validate:"required"` 
	Role 		   string `gorm:"type:varchar(100);not null" json:"role" validate:"required"` 
	Password   string `gorm:"type:varchar(100);not null" json:"password" validate:"required,min=6"`
	Banks      []Bank `gorm:"foreignKey:OwnerID"`
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


