package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nt2311-vn/go-fiber/internal/services"
)

func LoginPage(c *fiber.Ctx) error {
	return c.Render("login", "nil")
}

func ReigsterPage(c *fiber.Ctx) error {
	return c.Render("register", "nil")
}

func RegisterForm(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	confirmPassword := c.FormValue("confirm-password")

	pbClient := services.NewClient()

	err := pbClient.RegisterUser(email, password, confirmPassword)
	if err != nil {
		return c.SendString(err.Error())
	}

	c.Set("HX-Redirect", "/login")
	return c.Status(fiber.StatusCreated).
		SendString("User created successfully! Please login to continue.")
}

func LoginForm(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	pbClient := services.NewClient()

	user, err := pbClient.LoginUser(email, password)
	if err != nil {
		return c.SendString(err.Error())
	}

	c.Set("HX-Redirect", "start-oauth")

	if c.Query("code") == "" {
		return fmt.Errorf("authorization code not found")
	}

	nsClient, err := services.NewNSClient(c)
	if err != nil {
		return c.SendString(err.Error())
	}

	fmt.Println(nsClient.AccessToken)

	c.Cookie(&fiber.Cookie{
		Name:     "auth_token",
		Value:    user.Token,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		HTTPOnly: true,
		SameSite: "None",
	})

	c.Set("HX-Redirect", "/app/dashboard")

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "auth_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HTTPOnly: true,
		SameSite: "None",
	})

	c.Set("HX-Redirect", "/login")
	return c.SendStatus(fiber.StatusNoContent)
}

func StartOAuth(c *fiber.Ctx) error {
	fmt.Println(services.OAuthURL())
	return c.Redirect(services.OAuthURL())
}

func AuthCallBack(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Authorization code not found")
	}

	_, err := services.ExchangeToken(code)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Redirect("/app/dashboard")
}
