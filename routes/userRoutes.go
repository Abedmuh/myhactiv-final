package routes

import (
	"myhactiv/controllers"
	"myhactiv/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func UserRoutes(route *gin.RouterGroup, db *gorm.DB, validate *validator.Validate) {
	service := services.NewUserService()
	controller := controllers.NewUserController(service, db, validate)

	path:= route.Group("/user")
	{
    path.POST("/register", controller.CreateUser)
    path.POST("/login", controller.LoginUser)
    path.DELETE("/:id", controller.DeleteUser)
	}
}