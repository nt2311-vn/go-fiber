package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type NSClient struct {
	BaseURL        string
	Scope          string
	ClientID       string
	AuthorizeToken string
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

func signedJWT() string {
	privateKeyData, err := os.ReadFile("finopsx_pri.pem")
	if err != nil {
		return ""
	}

	privateKey, err := jwt.ParseECPrivateKeyFromPEM(privateKeyData)
	if err != nil {
		return ""
	}

	claims := jwt.MapClaims{
		"iss":   os.Getenv("NS_CONSUMER_KEY"),
		"scope": "restlets,rest_webservices",
		"aud":   os.Getenv("NS_BASE_URL") + "/auth/oauth2/v1/token",
		"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(),
		"iat":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token.Header["typ"] = "JWT"
	token.Header["alg"] = "EC"
	token.Header["kid"] = os.Getenv("NS_CERTIFICATE_ID")
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return ""
	}

	return signedToken
}

func requestToken() *TokenResponse {
	jwt := signedJWT()

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_assertion_type", "urn:ietf:params:oauth:client-assertion-type:jwt-bearer")

	data.Set("client_assertion", jwt)

	req, err := http.NewRequest(
		"POST",
		os.Getenv("NS_BASE_URL")+"/auth/oauth2/v1/token",
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return nil
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)

	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Headers:", resp.Header)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}
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

func NewNSClient() (*NSClient, error) {
	tokenResponse := requestToken()
	if tokenResponse == nil {
		return nil, fmt.Errorf("cannot get ns token")
	}

	return &NSClient{
		BaseURL:        os.Getenv("NS_BASE_URL"),
		AuthorizeToken: tokenResponse.AccessToken,
	}, nil
}
