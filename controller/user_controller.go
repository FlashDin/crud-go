package controller

import (
	"crud-go/dto"
	"crud-go/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) RegisterRoutes(app *fiber.App) {
	users := app.Group("/users")
	users.Get("/", c.GetAll)
	users.Get("/:id", c.GetByID)
	users.Post("/", c.Create)
	users.Put("/:id", c.Update)
	users.Delete("/:id", c.Delete)
}

func (c *UserController) GetAll(ctx *fiber.Ctx) error {
	return ctx.JSON(c.service.GetAll())
}

func (c *UserController) GetByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Message: "invalid id"})
	}
	res, err := c.service.GetByID(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(res)
}

func (c *UserController) Create(ctx *fiber.Ctx) error {
	var req dto.UserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error()})
	}
	res := c.service.Create(req)
	return ctx.Status(fiber.StatusCreated).JSON(res)
}

func (c *UserController) Update(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Message: "invalid id"})
	}
	var req dto.UserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error()})
	}
	res, err := c.service.Update(uint(id), req)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error()})
	}
	return ctx.JSON(res)
}

func (c *UserController) Delete(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Message: "invalid id"})
	}
	err = c.service.Delete(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Message: err.Error()})
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}
