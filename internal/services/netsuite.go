package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type NSClient struct {
	BaseURL string
	Token   string
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
	TokenType   string `json:"token_type"`
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
		"iss": os.Getenv("NS_CONSUMER_KEY"),
		"sub": os.Getenv("NS_ACCOUNT_ID"),
		"aud": os.Getenv("NS_BASE_URL"),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return ""
	}

	return signedToken
}

func requestToken() *TokenResponse {
	jwt := signedJWT()
	reqBody := map[string]string{
		"grant_type":            "client_credentials",
		"client_assertion_type": "urn:ietf:params:oauth:client-assertion-type:jwt-bearer",
		"client_assertion":      jwt,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	resp, err := http.Post(
		os.Getenv("NS_BASE_URL")+"/services/rest/auth/oauth2/v1/token",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer resp.Body.Close()

	fmt.Println(resp)

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
		BaseURL: os.Getenv("NS_BASE_URL") + "/services/rest",
		Token:   tokenResponse.AccessToken,
	}, nil
}
