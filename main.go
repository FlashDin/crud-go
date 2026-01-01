package main

import (
	"crud-go/config"
	"crud-go/controller"
	"crud-go/model"
	"crud-go/repository"
	"crud-go/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDatabase()
	db.AutoMigrate(&model.User{})

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	r := gin.Default()

	// === Global error middleware ===
	r.Use(func(c *gin.Context) {
		c.Next() // process request

		if len(c.Errors) > 0 {
			// pick the first error
			c.JSON(-1, gin.H{"message": c.Errors[0].Error()})
		}
	})

	// Register routes
	userController.RegisterRoutes(r)

	r.Run(":8090")
}
