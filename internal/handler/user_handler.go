package handler

import (
	"strconv"

	models "github.com/NOTMKW/API/internal/model"
	user_service "github.com/NOTMKW/API/internal/service"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService user_service.UserService
}

func NewUserHandler(userService user_service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid JSON format",
		})
	}

	user, err := h.UserService.CreateUser(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "created",
		"user":    user,
	})
}
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID Format",
		})
	}
	user, err := h.UserService.GetUserByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(user)
}
