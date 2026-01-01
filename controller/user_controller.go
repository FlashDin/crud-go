package controller

import (
	"crud-go/dto"
	"crud-go/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{service: service}
}

// Like @RequestMapping("/users")
func (c *UserController) RegisterRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.GET("", c.getAll)
		users.POST("", c.create)
	}
}

// GET /users
func (c *UserController) getAll(ctx *gin.Context) {
	users := c.service.GetAll()
	ctx.JSON(http.StatusOK, users)
}

// POST /users
func (c *UserController) create(ctx *gin.Context) {
	var req dto.UserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := c.service.Create(req)
	ctx.JSON(http.StatusCreated, res)
}
