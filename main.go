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

	// Auto create table (like Hibernate ddl-auto)
	db.AutoMigrate(&model.User{})

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	r := gin.Default()
	userController.RegisterRoutes(r)

	r.Run(":8090")
}
