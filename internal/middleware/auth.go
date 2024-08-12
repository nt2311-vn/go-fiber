package middleware

import "github.com/gofiber/fiber/v2"

var PublicPaths = map[string]bool{
	"/":         true,
	"/login":    true,
	"/register": true,
}

func AuthMiddleWare(c *fiber.Ctx) error {
	isAuthenticated := c.Get("X-Dev-Access") == "true"

	if isPublicPath(c.Path()) {
		return c.Next()
	}

	if !isAuthenticated {
		return c.Redirect("/login", fiber.StatusSeeOther)
	}

	return c.Next()
}

func IsAuthenticated(c *fiber.Ctx) bool {
	return c.Get("X-Dev-Access") == "true"
}

func isPublicPath(path string) bool {
	return PublicPaths[path]
}
