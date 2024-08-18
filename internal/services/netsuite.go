package services

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type NSClient struct {
	BaseURL     string
	AccessToken string
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    string `json:"expires_in"`
	TokenType    string `json:"token_type"`
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

func OAuthURL() string {
	state, err := generateState()
	if err != nil {
		return ""
	}

	params := url.Values{}
	params.Set("response_type", "code")
	params.Set("client_id", os.Getenv("NS_CONSUMER_KEY"))
	params.Set("redirect_uri", os.Getenv("NS_REDIRECT_URI"))
	params.Set("scope", "rest_webservices")
	params.Set("state", state)

	redirectURL := os.Getenv("NS_AUTH_ENDPOINT") + "?" + params.Encode()
	return redirectURL
}

func ExchangeToken(code string) (*TokenResponse, error) {
	data := url.Values{}
	data.Set("code", code)
	data.Set("redirect_uri", os.Getenv("NS_REDIRECT_URI"))
	data.Set("grant_type", "authorization_code")

	req, err := http.NewRequest(
		"POST",
		os.Getenv("NS_BASE_URL")+"/auth/oauth2/v1/token",
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	auth := base64.StdEncoding.EncodeToString(
		[]byte(os.Getenv("NS_CONSUMER_KEY") + ":" + os.Getenv("NS_CONSUMER_SECRET")),
	)
	req.Header.Set("Authorization", "Basic "+auth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to exchange token: %s", resp.Status)
	}

	token := &TokenResponse{}
	err = json.NewDecoder(resp.Body).Decode(token)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func NewNSClient(c *fiber.Ctx) (*NSClient, error) {
	code := c.Query("code")
	if code == "" {
		return nil, fmt.Errorf("authorization code not found")
	}
	fmt.Println(code)
	token, err := ExchangeToken(code)
	if err != nil {
		return nil, err
	}

	return &NSClient{
		BaseURL:     os.Getenv("NS_BASE_URL"),
		AccessToken: token.AccessToken,
	}, nil
}
