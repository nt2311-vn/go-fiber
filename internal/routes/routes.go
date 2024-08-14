package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nt2311-vn/go-fiber/internal/handlers"
)

func Setup(app *fiber.App) {
	app.Get("/", handlers.HomePage)
	app.Get("/login", handlers.LoginPage)
	app.Get("/register", handlers.ReigsterPage)

	app.Post("/validate-field", handlers.ValidateFields)

	app.Post("/register", handlers.RegisterForm)
}
