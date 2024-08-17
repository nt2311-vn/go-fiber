package services

import (
	"bytes"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
		fmt.Println("Error reading private key:", err)
		return ""
	}
	block, _ := pem.Decode(privateKeyData)

	if block == nil {
		fmt.Println("Error decoding private key")
		return ""
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("Error parsing private key:", err)
		return ""
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"iss": ns.ClientID,
		"aud": ns.BaseURL + "/auth/oauth2/v1/token",
		"exp": time.Now().Add(5 * time.Minute).Unix(),
		"iat": time.Now().Unix(),
	})

	token.Header["kid"] = ns.CertID
	token.Header["alg"] = "ES256"

	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		fmt.Println("Error signing token:", err)
		return ""
	}

	return signedToken
}

func RequestToken() *TokenResponse {
	nsClient := newNSClient()
	jwt := nsClient.signedJWT()

	form := url.Values{}
	form.Add("grant_type", nsClient.GrantType)
	form.Add("client_assertion_type", nsClient.AssertType)
	form.Add("client_assertion", jwt)

	resp, err := http.Post(
		nsClient.BaseURL+"/auth/oauth2/v1/token",
		"application/x-www-form-urlencoded",
		bytes.NewBufferString(form.Encode()),
	)
	if err != nil {
		fmt.Println("Reponse error:", err)
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
