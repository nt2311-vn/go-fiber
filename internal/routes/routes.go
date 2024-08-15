package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nt2311-vn/go-fiber/internal/handlers"
	"github.com/nt2311-vn/go-fiber/internal/middleware"
)

func Setup(app *fiber.App) {
	app.Get("/", handlers.HomePage)
	app.Get("/login", handlers.LoginPage)
	app.Get("/register", handlers.ReigsterPage)
	app.Post("/validate-field", handlers.ValidateFields)
	app.Post("/register", handlers.RegisterForm)
	app.Post("/login", handlers.LoginForm)

	authGroup := app.Group("/app", middleware.AuthMiddleWare)
	authGroup.Get("/dashboard", handlers.DashboardPage)
	authGroup.Get("/stock", handlers.StockPage)
	authGroup.Get("/sales", handlers.SalesPage)
	authGroup.Get("/pnp", handlers.PnPPage)
	authGroup.Get("/budget", handlers.BudgetPage)
}
