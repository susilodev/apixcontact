package routes

import (
	"apixcontact/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Tambahkan route untuk user
	api.Get("/users", handlers.GetUsers)
	api.Post("/users", handlers.CreateUser)
}
