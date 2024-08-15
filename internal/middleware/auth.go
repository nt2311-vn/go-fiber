package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nt2311-vn/go-fiber/internal/services"
)

var PublicPaths = map[string]bool{
	"/":               true,
	"/login":          true,
	"/register":       true,
	"/validate-field": true,
	"/simple":         true,
	"/dashboard":      true,
	"/sales":          true,
	"/stock":          true,
	"/pnp":            true,
	"/budget":         true,
}

func AuthMiddleWare(c *fiber.Ctx) error {
	sessionToken := c.Get("Authorization")

	pbClient := services.NewClient()

	_, err := pbClient.VerifyUserToken(sessionToken)
	if err != nil {
		return c.Redirect("/login")
	}

	return c.Next()
}

func isPublicPath(path string) bool {
	return PublicPaths[path]
}
