package handlers

import "github.com/gofiber/fiber/v2"

func LoginPage(c *fiber.Ctx) error {
	return c.Render("login", "nil")
}

func ReigsterPage(c *fiber.Ctx) error {
	return c.Render("register", "nil")
}

func RegisterForm(c *fiber.Ctx) error {
	return nil
}
