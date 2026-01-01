package main

import (
	"crud-go/controller"
	"crud-go/repository"
	"crud-go/service"
	"log"
	"net/http"
)

func main() {
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	userController.RegisterRoutes()

	log.Println("Server started on :8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
