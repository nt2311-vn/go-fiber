package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nt2311-vn/go-fiber/views"
)

func HomeHandler(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}

func LoginPage(c *fiber.Ctx) error {
	loginPage := views.Login()
	return Render(loginPage)(c)
}
