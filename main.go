package main

import (
	"apixcontact/config"
	"apixcontact/routes"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	cfg := config.LoadConfig()

	// Setup routes
	routes.SetupRoutes(app)

	// Jalankan server
	log.Fatal(app.Listen(fmt.Sprintf(":%s", cfg.ServerPort)))
}
