package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	// Logika untuk mengambil data user
	users := []string{"User1", "User2"}
	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	// Logika untuk membuat user baru
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})
}
