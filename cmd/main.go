package main

import (
	"log"

	handlers "github.com/NOTMKW/API/internal/handler"
	models "github.com/NOTMKW/API/internal/model"
	MakeUser "github.com/NOTMKW/API/internal/repository"
	user_service "github.com/NOTMKW/API/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	dsn := "host=localhost user=postgres password=imrankw dbname=clone port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}

	db.AutoMigrate(&models.User{})

	user_service := user_service.NewUserService(db)

	UserHandler := handlers.NewUserHandler(user_service)

	api := app.Group("/api/v1")
	users := api.Group("/users")
	users.Post("/", MakeUser.CreateUser(db))
	users.Get("/:id", UserHandler.GetUser)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "server is running",
		})
	})
	log.Fatal(app.Listen(":8080"))
}
