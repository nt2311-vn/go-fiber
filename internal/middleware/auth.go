package middleware

import (
	"github.com/gofiber/fiber/v2"
)

var PublicPaths = map[string]bool{
	"/":               true,
	"/login":          true,
	"/register":       true,
	"/validate-field": true,
	"/auth/callback":  true,
}

func AuthMiddleWare(c *fiber.Ctx) error {
	if isPublicPath(c.Path()) {
		return c.Next()
	}
	token := c.Cookies("auth_token")

	if token == "" {
		return c.Redirect("/login")
	}

	return c.Next()
}

func isPublicPath(path string) bool {
	return PublicPaths[path]
}
