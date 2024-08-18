package services

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
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
		ClientID:     os.Getenv("NS_CLIENT_ID"),
		State:        state,
		RedirectURI:  os.Getenv("NS_REDIRECT_URI"),
		AuthURL:      "https://5574610.app.netsuite.com/app/login/oauth2/authorize.nl",
	}, nil
}

func RequestToken() *TokenResponse {
	nsClient, err := newNSClient()
	if err != nil {
		fmt.Println("Error creating NSClient:", err)
		return nil
	}

	form := url.Values{}
	form.Add("response_type", nsClient.ResponseType)
	form.Add("client_id", nsClient.ClientID)
	form.Add("redirect_uri", nsClient.RedirectURI)
	form.Add("scope", nsClient.Scope)
	form.Add("state", nsClient.State)

	req, err := http.NewRequest(
		"GET",
		nsClient.AuthURL+"?"+form.Encode(), nil,
	)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Headers:", resp.Header)
	fmt.Println("Response Body:", string(body))

	var tokenResponse TokenResponse
	err = json.NewDecoder(resp.Body).Decode(&tokenResponse)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println(tokenResponse.AccessToken)

	return &tokenResponse
}
