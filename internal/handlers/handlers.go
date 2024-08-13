package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nt2311-vn/go-fiber/views"
)

func HomeHandler(c *fiber.Ctx) error {
	isAuthenticated := c.Get("X-Dev-Access") == "true"
	homePage := views.Layout(isAuthenticated)
	return Render(homePage)(c)
}

func LoginPage(c *fiber.Ctx) error {
	isAuthenticated := c.Get("X-Dev-Access") == "true"
	loginPage := views.Layout(isAuthenticated, views.Login())

	return Render(loginPage)(c)
}

func ReigsterPage(c *fiber.Ctx) error  {
	registerPage := views.Layout(false, views.Register())

	return Render(registerPage)(c)
}

func Register(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	confirmPassword := c.FormValue("confirm-password")

	if password != confirmPassword {
		return c.Redirect("/register", fiber.StatusSeeOther)
	}

	payload := map[string]string{
		"email":           email,
		"password":        password,
		"confirmPassword": confirmPassword,
	}

	jsonPayload, err := json.Marshal(payload)
	resp, err := http.Post(
		"http://127.0.0.1:8080/api/collections/users",
		"application/json",
		bytes.NewBuffer(jsonPayload),
	)

	defer resp.Body.Close()

	if err != nil {
		return c.SendString(err.Error())
	}
	return c.Redirect("/login", fiber.StatusSeeOther)
}
