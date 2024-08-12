package main

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/nt2311-vn/go-fiber/internal/middleware"
	"github.com/nt2311-vn/go-fiber/internal/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	listenAddr := os.Getenv("LISTEN_ADDR")

	app := fiber.New()
	app.Static("/static", filepath.Join("static"))

	app.Use(middleware.AuthMiddleWare)

	routes.Setup(app)
	slog.Info("Serving web app on", "address", listenAddr)

	app.Listen(listenAddr)
}