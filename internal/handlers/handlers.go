package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nt2311-vn/go-fiber/internal/api"
	"github.com/nt2311-vn/go-fiber/views"
)

func HomeHandler(c *fiber.Ctx) error {
	isAuthenticated := c.Get("X-Dev-Access") == "true"
	homePage := views.Layout(isAuthenticated)
	return Render(homePage)(c)
}

func LoginPage(c *fiber.Ctx) error {
	isAuthenticated := c.Get("X-Dev-Access") == "true"
	loginPage := views.Layout(isAuthenticated, views.Login())

	return Render(loginPage)(c)
}

func ReigsterPage(c *fiber.Ctx) error {
	registerPage := views.Layout(false, views.Register())

	return Render(registerPage)(c)
}

func RegisterForm(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	confirmPassword := c.FormValue("confirm-password")

	if password != confirmPassword {
		return c.Status(fiber.StatusBadRequest).
			SendString(`<div class="text-red-500 dark:text-red-100">Passwords do not match</div>`)
	}

	dbClient := api.NewClient()

	err := dbClient.RegisterUser(email, password, confirmPassword)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString(`<div class="text-red-500 dark:text-red-100">Error registering user</div>`)
	}
	return c.SendString(`<script>window.location.href = "/login";</script>`)
}
