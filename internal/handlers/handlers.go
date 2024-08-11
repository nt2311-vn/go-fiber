package handlers

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/nt2311-vn/go-fiber/views"
	"golang.org/x/oauth2"
)

var jiraOAuth2Config = oauth2.Config{
	ClientID:     os.Getenv("JIRA_CLIENT_ID"),
	ClientSecret: os.Getenv("JIRA_CLIENT_SECRET"),
	RedirectURL:  os.Getenv("JIRA_REDIRECT_URI"),
	Scopes: []string{
		"read:jira-user",
		"read:jira-work",
		"write:jira-work",
		"offline_access",
		"read:me",
	},
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://auth.atlassian.com/authorize",
		TokenURL: "https://auth.atlassian.com/oauth/token",
	},
}

func HomeHandler(c *fiber.Ctx) error {
	isAuthenticated := c.Get("X-Dev-Access") == "true"
	homePage := views.Layout(views.Navbar(isAuthenticated))
	return Render(homePage)(c)
}

func LoginPage(c *fiber.Ctx) error {
	isAuthenticated := c.Get("X-Dev-Access") == "true"

	loginPage := views.Layout(views.Navbar(isAuthenticated), views.Login())

	return Render(loginPage)(c)
}

func JiraLogin(c *fiber.Ctx) error {
	url := jiraOAuth2Config.AuthCodeURL("state", oauth2.AccessTypeOnline)
	return c.Redirect(url)
}

func JiraCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Authorization code not found")
	}

	token, err := jiraOAuth2Config.Exchange(c.Context(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to exchange token")
	}

	client := jiraOAuth2Config.Client(c.Context(), token)
	resp, err := client.Get("https://api.atlassian.com/me")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get user info")
	}
	defer resp.Body.Close()

	return c.Status(fiber.StatusOK).SendString("User info: " + resp.Status)
}
