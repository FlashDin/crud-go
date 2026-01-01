package controller

import (
	"crud-go/dto"
	"crud-go/service"
	"crud-go/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) RegisterRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.GET("", c.getAll)
		users.POST("", c.create)
	}
}

func (c *UserController) getAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.service.GetAll())
}

func (c *UserController) create(ctx *gin.Context) {
	var req dto.UserRequest

	// @Valid equivalent
	//if err := ctx.ShouldBindJSON(&req); err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		"message": "Validation failed",
	//		"error":   err.Error(),
	//	})
	//	return
	//}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": util.ParseValidationError(err),
		})
		return
	}

	res := c.service.Create(req)
	ctx.JSON(http.StatusCreated, res)
}
