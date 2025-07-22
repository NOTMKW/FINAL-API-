package routes

import (
	handlers "github.com/NOTMKW/API/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userHandler *handlers.UserHandler) {
	api := app.Group("/api/v1")
	users := api.Group("/users")

	users.Post("/", userHandler.CreateUser)
	users.Get("/:id", userHandler.GetUser)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "server is running",
		})
	})
}
