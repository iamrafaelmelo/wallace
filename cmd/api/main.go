package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	postgres, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatalf("unable to connect to database: %v\n", err)
	}

	defer postgres.Close(context.Background())

	app := fiber.New()
	app.Use(recoverer.New())
	app.Use(logger.New())

	app.Get("/healthz", func(ctx fiber.Ctx) error {
		err := postgres.Ping(ctx.Context())

		if err != nil {
			return ctx.SendStatus(http.StatusInternalServerError)
		}

		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"status": "healthy",
		})
	})

	app.Get("/", func(ctx fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Hello, world!",
		})
	})

	app.Listen(":8000")
}
