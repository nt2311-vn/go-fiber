package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nt2311-vn/go-fiber/views"
)

func HomeHandler(c *fiber.Ctx) error {
	isAuthenticated := c.Get("X-Dev-Access") == "true"
	homePage := views.Layout(views.Navbar(isAuthenticated))
	return Render(homePage)(c)
}

func LoginPage(c *fiber.Ctx) error {
	isAuthenticated := c.Get("X-Dev-Access") == "true"

	loginPage := views.Layout(views.Navbar(isAuthenticated), views.Login())

	return Render(loginPage)(c)
}
