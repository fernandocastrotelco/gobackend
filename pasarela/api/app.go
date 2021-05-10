package main

import (
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var msg string
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
				msg = e.Message
			}

			if msg == "" {
				msg = "No se pudo procesar el llamado a la api"
			}

			err = ctx.Status(code).JSON(internalError{
				Message: msg,
			})

			return nil
		},
	})

	app.Use(recover.New())
	app.Use(cors.New())
	frontend := filepath.Join("api", "frontend")
	app.Static("/", frontend)

	_ = app.Listen(":3390")
}

type internalError struct {
	Message string `json:"message"`
}
