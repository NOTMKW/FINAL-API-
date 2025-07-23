package main

import (
	"log"

	"github.com/NOTMKW/API/internal/config"
	handlers "github.com/NOTMKW/API/internal/handler"
	models "github.com/NOTMKW/API/internal/model"
	repository "github.com/NOTMKW/API/internal/repo"
	"github.com/NOTMKW/API/internal/routes"
	user_service "github.com/NOTMKW/API/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})
	app.Use(cors.New())

	config.ConnectDatabase()
	db := config.DB

	if db == nil {
		log.Fatal("database connection is nil")
	}

	db.AutoMigrate(&models.User{})

	userRepo := repository.NewUserRepository(db)
	userService := user_service.NewUserService(userRepo)
	UserHandler := handlers.NewUserHandler(userService)

	routes.SetupRoutes(app, UserHandler)

	log.Fatal(app.Listen(":8080"))
}
