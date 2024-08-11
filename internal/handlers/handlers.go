package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nt2311-vn/go-fiber/views"
)

func HomeHandler(c *fiber.Ctx) error {
	homePage := views.Layout(views.Navbar(false))
	return Render(homePage)(c)
}

func LoginPage(c *fiber.Ctx) error {
	loginPage := views.Layout(views.Navbar(false), views.Login())
	return Render(loginPage)(c)
}
