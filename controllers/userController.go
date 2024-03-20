package controllers

import (
	"myhactiv/models"
	"myhactiv/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserCtrlInter interface {
	LoginUser(c *gin.Context)
	CreateUser(c *gin.Context) 
	DeleteUser(c *gin.Context)
}


type UserController struct {
	UserService services.UserSvcInter
	DB *gorm.DB
	Validate *validator.Validate
}

func NewUserController(userService services.UserSvcInter, db *gorm.DB, validate *validator.Validate) UserCtrlInter {
  return &UserController{
		UserService: userService, 
		DB: db, 
		Validate: validate}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User
  if err := c.ShouldBindJSON(&user); err!= nil {
    c.JSON(400, gin.H{
      "error": "bad request",
    })
    return
  }
  if err := uc.Validate.Struct(user); err!= nil {
    c.JSON(400, gin.H{
      "error": "request not validated",
    })
    return
  }
	
	if err :=uc.UserService.CheckUser(user.Username, uc.DB, c); err != nil {
    c.JSON(400, gin.H{
      "error":err.Error(),
    })
    return
  }

	newUser, err := uc.UserService.AddUser(user, uc.DB, c)
  if err!= nil {
    c.JSON(400, gin.H{
      "error": err.Error(),
    })
    return
  }
  c.JSON(200, gin.H{
    "message": "create success",
		"data": newUser,
  })
}

func (uc *UserController) LoginUser(c *gin.Context) {
	var req models.UserLogin
  if err := c.ShouldBindJSON(&req); err!= nil {
    c.JSON(400, gin.H{
      "error": "bad request",
    })
    return
  }
  if err := uc.Validate.Struct(req); err!= nil {
    c.JSON(400, gin.H{
      "error": "unvalid request",
    })
    return
  }

	token,err := uc.UserService.LoginUser(req, uc.DB,c);
  if err!= nil {
    c.JSON(400, gin.H{
      "error": err.Error(),
    })
    return
  }
  c.JSON(200, gin.H{
    "message": "login success",
		"data": token,
  })
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	username := c.Param("id")
  if err := uc.UserService.DeleteUser(username, uc.DB, c); err!= nil {
    c.JSON(400, gin.H{
      "error": err.Error(),
    })
    return
  }
  c.JSON(200, gin.H{
    "message": "delete success",
  })
}