package main

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/nt2311-vn/go-fiber/internal/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	listenAddr := os.Getenv("LISTEN_ADDR")

	app := fiber.New()
	app.Static("/static", filepath.Join("static"))

	routes.Setup(app)

	app.Listen(listenAddr)
	slog.Info("Serving web app on", "address", listenAddr)
}
