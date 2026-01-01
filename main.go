package main

import (
	"crud-go/config"
	"crud-go/controller"
	"crud-go/model"
	"crud-go/repository"
	"crud-go/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := config.ConnectDatabase()
	db.AutoMigrate(&model.User{})

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	app := fiber.New()

	userController.RegisterRoutes(app)

	app.Listen(":8090")
}
