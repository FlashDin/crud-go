package main

import (
	"crud-go/controller"
	"crud-go/repository"
	"crud-go/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Dependency wiring (like Spring DI)
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	r := gin.Default()

	userController.RegisterRoutes(r)

	r.Run(":8090")
}
