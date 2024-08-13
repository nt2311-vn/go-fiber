package handlers

import (
	"html/template"
	"regexp"

	"github.com/gofiber/fiber/v2"
)

func renderTemplate(c *fiber.Ctx, templ string, data interface{}) error {
	t, err := template.ParseFiles("views/"+templ, "views/navbar.html", "views/footer.html")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return t.Execute(c.Response().BodyWriter(), data)
}

func isValidEmail(email string) bool {
	// Regex validation for email input

	const emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(emailRegexPattern).MatchString(email)
}
