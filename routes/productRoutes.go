package routes

import (
	"myhactiv/controllers"
	"myhactiv/middleware"
	"myhactiv/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func ProductRoutes(route *gin.RouterGroup, tx *gorm.DB, validate *validator.Validate) {
	service := services.NewProductService()
  controller := controllers.NewProductController(service, tx, validate)

	endpoint := route.Group("/product")
	{
		endpoint.GET("/", )

		endpoint.POST("/", middleware.Authentication(),controller.PostProduct)
		endpoint.PUT("/:id", middleware.Authentication(),controller.PatchProduct)
		endpoint.DELETE("/:id", middleware.Authentication(),controller.DeleteProduct)
	}
  
}