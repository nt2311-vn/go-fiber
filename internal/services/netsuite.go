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
	BaseURL    string
	Scope      string
	ClientID   string
	TokenID    string
	CertID     string
	AccountID  string
	GrantType  string
	AssertType string
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

func newNSClient() *NSClient {
	return &NSClient{
		BaseURL:    os.Getenv("NS_BASE_URL"),
		Scope:      "rest_webservices",
		ClientID:   os.Getenv("NS_CLIENT_ID"),
		TokenID:    os.Getenv("NS_TOKEN_ID"),
		CertID:     os.Getenv("NS_CERT_ID"),
		AccountID:  os.Getenv("NS_ACCOUNT_ID"),
		GrantType:  "client_credentials",
		AssertType: "urn:ietf:params:oauth:client-assertion-type:jwt-bearer",
	}
}

func (ns *NSClient) signedJWT() string {
	privateKeyData, err := os.ReadFile("finopsx_pri.pem")
	if err != nil {
		return ""
	}

	privateKey, err := jwt.ParseECPrivateKeyFromPEM(privateKeyData)
	if err != nil {
		return ""
	}

	claims := jwt.MapClaims{
		"iss":   ns.ClientID,
		"scope": ns.Scope,
		"aud":   ns.BaseURL + "/auth/oauth2/v1/token",
		"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(),
		"iat":   time.Now().Unix(),
		"jti":   ns.TokenID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token.Header["typ"] = "JWT"
	token.Header["alg"] = "ES256"
	token.Header["kid"] = ns.CertID
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return ""
	}

	return signedToken
}

func RequestToken() *TokenResponse {
	nsClient := newNSClient()
	jwt := nsClient.signedJWT()

	formData := url.Values{}
	formData.Set("grant_type", nsClient.GrantType)
	formData.Set("client_assertion_type", nsClient.AssertType)
	formData.Set("client_assertion", jwt)

	resp, err := http.Post(
		nsClient.BaseURL+"/auth/oauth2/v1/token",
		"application/x-www-form-urlencoded",
		strings.NewReader(formData.Encode()),
	)
	if err != nil {
		fmt.Println("Reponse error:", err)
		return nil
	}

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
