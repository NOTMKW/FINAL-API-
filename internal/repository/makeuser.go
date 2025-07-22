package MakeUser

import (
	"strings"

	user "github.com/NOTMKW/API/internal/dto"
	models "github.com/NOTMKW/API/internal/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req user.CreateUserRequest

		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid JSON format",
			})
		}

		user := models.User{
			Firstname: req.Firstname,
			Lastname:  req.Lastname,
			Email:     req.Email,
			Password:  req.Password,
		}

		if err := db.Create(&user).Error; err != nil {
			errStr := strings.ToLower(err.Error())
			if strings.Contains(errStr, "unique constraint failed") ||
				strings.Contains(errStr, "duplicate entry") ||
				strings.Contains(errStr, "duplicate key value") {
				return c.Status(fiber.StatusConflict).JSON(fiber.Map{
					"error": "user with this email exists",
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to create user",
			})
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "user created",
			"user_id": user.ID,
		})
	}
}
