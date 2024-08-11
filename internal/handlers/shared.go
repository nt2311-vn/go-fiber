package handlers

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func Render(t templ.Component) fiber.Handler {
	return func(c *fiber.Ctx) error {
		handler := adaptor.HTTPHandler(templ.Handler(t))
		return handler(c)
	}
}
