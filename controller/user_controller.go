package controller

import (
	"crud-go/dto"
	"crud-go/service"
	"encoding/json"
	"net/http"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) RegisterRoutes() {
	http.HandleFunc("/users", c.handleUsers)
}

func (c *UserController) handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		users := c.service.GetAll()
		json.NewEncoder(w).Encode(users)

	case http.MethodPost:
		var req dto.UserRequest
		json.NewDecoder(r.Body).Decode(&req)

		res := c.service.Create(req)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(res)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
