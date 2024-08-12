package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nt2311-vn/go-fiber/internal/handlers"
)

func Setup(app *fiber.App) {
	app.Get("/", handlers.HomeHandler)
	app.Get("/login", handlers.LoginPage)
	app.Post("/register", handlers.Register)
}
