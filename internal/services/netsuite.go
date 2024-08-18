package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type NSClient struct {
	AccessToken string
	BaseURL     string
}

type AppID struct {
	ClientID   string
	ClientSec  string
	CertID     string
	GrantType  string
	Scope      string
	AssertType string
	BaseURL    string
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

func createApp() *AppID {
	return &AppID{
		ClientID:   os.Getenv("NS_CONSUMER_KEY"),
		ClientSec:  os.Getenv("NS_CONSUMER_SECRET"),
		CertID:     os.Getenv("NS_CERT_ID"),
		GrantType:  "client_credentials",
		Scope:      "rest_webservices",
		AssertType: "urn:ietf:params:oauth:client-assertion-type:jwt-bearer",
		BaseURL:    os.Getenv("NS_BASE_URL"),
	}
}

func (app *AppID) signedToken() (string, error) {
	byteKey, err := os.ReadFile("private.pem")
	if err != nil {
		return "", fmt.Errorf("error reading private key: %v", err)
	}

	priKey, err := jwt.ParseRSAPrivateKeyFromPEM(byteKey)
	if err != nil {
		return "", fmt.Errorf("error parsing private key: %v", err)
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodPS256, jwt.MapClaims{
		"iss":   app.ClientID,
		"scope": app.Scope,
		"aud":   app.BaseURL + "/auth/oauth2/v1/token",
		"exp":   time.Now().Add(time.Minute * 5).Unix(),
		"iat":   time.Now().Unix(),
	})

	jwtToken.Header["alg"] = "PS256"
	jwtToken.Header["typ"] = "JWT"
	jwtToken.Header["kid"] = app.CertID

	tokenString, err := jwtToken.SignedString(priKey)
	if err != nil {
		return "", fmt.Errorf("error signing token: %v", err)
	}

	return tokenString, nil
}

func NewNSClient() (*NSClient, error) {
	app := createApp()
	token, err := app.signedToken()
	if err != nil {
		return nil, fmt.Errorf("error signed token:%v", err)
	}

	form := url.Values{}
	form.Add("grant_type", app.GrantType)
	form.Add("client_assertion_type", app.AssertType)
	form.Add("client_assertion", token)

	req, err := http.NewRequest(
		"POST",
		app.BaseURL+"/auth/oauth2/v1/token",
		bytes.NewBufferString(form.Encode()),
	)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}

	defer resp.Body.Close()

	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %v, token: %v", err, tokenResp)
	}

	return &NSClient{AccessToken: tokenResp.AccessToken, BaseURL: app.BaseURL}, nil
}

func (ns *NSClient) GetUser(email string) {
}
