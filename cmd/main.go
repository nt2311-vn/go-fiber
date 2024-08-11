package main

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	listenAddr := os.Getenv("LISTEN_ADDR")

	app := fiber.New()
	app.Static("/static", filepath.Join("static"))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	app.Listen(listenAddr)
	slog.Info("Serving web app on", "address", listenAddr)
}
