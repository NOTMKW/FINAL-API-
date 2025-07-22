package handler

import (
	"strconv"

	user "github.com/NOTMKW/API/internal/dto"
	user_service "github.com/NOTMKW/API/internal/service"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService *user_service.UserService
}

func NewUserHandler(UserService *user_service.UserService) *UserHandler {
	return &UserHandler{UserService: UserService}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req user.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid JSON format",
		})
	}
	return c.JSON(fiber.Map{
		"message": "created",
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
	user, err := h.UserService.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	user, err = h.UserService.GetUserByID(int64(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(user)
}
