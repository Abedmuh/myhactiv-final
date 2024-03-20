package routes

import (
	"myhactiv/controllers"
	"myhactiv/middleware"
	"myhactiv/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func BankRoutes(route *gin.RouterGroup, tx *gorm.DB, validate *validator.Validate) {
	service := services.NewBankService()
  controller := controllers.NewBankController(service, tx, validate)

	endpoint := route.Group("/bank")
	endpoint.Use(middleware.Authentication())
	{
		endpoint.POST("/", controller.PostBank)
		endpoint.PUT("/:id", controller.PatchBank)
		endpoint.DELETE("/:id", controller.DeleteBank)
	}
}