package main

import (
	"fmt"
	"myhactiv/config"
	"myhactiv/routes"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	db, err := config.ConnectDatabase()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	app := gin.Default()
	validate := validator.New()

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := app.Group("/")
	{
		routes.UserRoutes(v1,db,validate)
		routes.BankRoutes(v1,db,validate)
		routes.ProductRoutes(v1,db,validate)
	}
	

	app.Run() // listen and serve on 0.0.0.0:8080
}