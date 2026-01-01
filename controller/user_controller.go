package controller

import (
	"crud-go/dto"
	"crud-go/service"
	"crud-go/util"
	"net/http"
	"strconv"

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
		users.GET("/:id", c.getByID)
		users.PUT("/:id", c.update)
		users.DELETE("/:id", c.delete)
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

func (c *UserController) getByID(ctx *gin.Context) {
	id, _ := parseID(ctx)
	res, err := c.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *UserController) update(ctx *gin.Context) {
	id, _ := parseID(ctx)
	var req dto.UserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	res, err := c.service.Update(id, req)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *UserController) delete(ctx *gin.Context) {
	id, _ := parseID(ctx)
	err := c.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

// helper to parse :id
func parseID(ctx *gin.Context) (uint, error) {
	var id uint64
	var err error
	if id, err = strconv.ParseUint(ctx.Param("id"), 10, 64); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
	}
	return uint(id), err
}
