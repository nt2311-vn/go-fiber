package services

import (
	"crypto/rand"
	"encoding/base64"
	"net/url"
	"os"

	"github.com/gofiber/fiber/v2"
)

type NSClient struct {
	BaseURL      string
	Scope        string
	ClientID     string
	State        string
	RedirectURI  string
	ResponseType string
	AuthURL      string
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type TokenRequest struct {
	GrantType           string `json:"grant_type"`
	ClientAssertionType string `json:"client_assertion_type"`
	ClientAssertion     string `json:"client_assertion"`
}

func generateState() (string, error) {
	randomBytes := make([]byte, 32)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", nil
	}

	state := base64.URLEncoding.EncodeToString(randomBytes)
	return state, nil
}

func newNSClient() (*NSClient, error) {
	state, err := generateState()
	if err != nil {
		return nil, err
	}

	return &NSClient{
		ResponseType: "code",
		BaseURL:      os.Getenv("NS_BASE_URL"),
		Scope:        "rest_webservices",
		ClientID:     os.Getenv("NS_CONSUMER_KEY"),
		State:        state,
		RedirectURI:  os.Getenv("NS_REDIRECT_URI"),
		AuthURL:      "https://5574610.app.netsuite.com/app/login/oauth2/authorize.nl",
	}, nil
}

func GetToken(c *fiber.Ctx) error {
	ns, err := newNSClient()
	if err != nil {
		return err
	}

	params := url.Values{}

	params.Add("response_type", ns.ResponseType)
	params.Add("client_id", ns.ClientID)
	params.Add("scope", ns.Scope)
	params.Add("state", ns.State)
	params.Add("redirect_uri", ns.RedirectURI)

	redirectURL := ns.AuthURL + "?" + params.Encode()

	return c.Redirect(redirectURL, fiber.StatusTemporaryRedirect)
}
