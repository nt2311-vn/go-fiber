package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nt2311-vn/go-fiber/internal/services"
)

func LoginPage(c *fiber.Ctx) error {
	return c.Render("login", "nil")
}

func ReigsterPage(c *fiber.Ctx) error {
	return c.Render("register", "nil")
}

func RegisterForm(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	confirmPassword := c.FormValue("confirm-password")

	pbClient := services.NewClient()

	err := pbClient.RegisterUser(email, password, confirmPassword)
	if err != nil {
		return c.SendString(err.Error())
	}

	c.Set("HX-Redirect", "/login")
	return c.Status(fiber.StatusCreated).
		SendString("User created successfully! Please login to continue.")
}
