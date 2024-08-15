package handlers

import "github.com/gofiber/fiber/v2"

func HomePage(c *fiber.Ctx) error {
	return c.Render("home", "nil")
}

func AuthPage(c *fiber.Ctx) error {
	return c.Render("authView", "nil")
}
