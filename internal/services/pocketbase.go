package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Client struct {
	BaseURL string
}

type AuthClient struct {
	AuthToken string
}

type Admin struct {
	ID      string `json:"id"`
	Created string `json:"created"`
	Updated string `json:"updated"`
	Email   string `json:"email"`
	Avatar  int    `json:"avatar"`
}

type AdminLogin struct {
	Email    string `json:"identity"`
	Password string `json:"password"`
}

type AdminAuthResponse struct {
	Token string `json:"token"`
	Admin Admin  `json:"admin"`
}

func getAdminToken() (string, error) {
	adminEmail := os.Getenv("PB_ADMIN_EMAIL")
	adminPassword := os.Getenv("PB_ADMIN_PASSWORD")

	adminLogin := AdminLogin{
		Email:    adminEmail,
		Password: adminPassword,
	}

	jsonVal, err := json.Marshal(adminLogin)
	if err != nil {
		return "", fmt.Errorf("cannot login to bytes: %v", err)
	}

	resp, err := http.Post(
		os.Getenv("PB_API_URL")+"admins/auth-with-password",
		"application/json",
		bytes.NewBuffer(jsonVal),
	)
	if err != nil {
		return "", fmt.Errorf("cannot login as admin: %v", err)
	}

	defer resp.Body.Close()

	var adminResponse AdminAuthResponse

	err = json.NewDecoder(resp.Body).Decode(&adminResponse)
	if err != nil {
		return "", fmt.Errorf("cannot decode admin response: %v", err)
	}

	return adminResponse.Token, nil
}

func NewClient() *Client {
	return &Client{
		BaseURL: os.Getenv("PB_API_URL"),
	}
}
