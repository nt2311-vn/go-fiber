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
	certFile := filepath.Join("localhost.pem")
	keyFile := filepath.Join("localhost-key.pem")

	app := fiber.New()
	app.Static("/static", filepath.Join("static"))
	app.Use(middleware.AuthMiddleWare)

	routes.Setup(app)

	if err := app.ListenTLS(listenAddr, certFile, keyFile); err != nil {
		slog.Error("Failed to start server", "error", err)
	}

	slog.Info("Serving web app on", "address", listenAddr)
}
