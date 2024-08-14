package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

type Client struct {
	BaseURL   string
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

func NewClient() *Client {
	adminEmail := os.Getenv("PB_ADMIN_EMAIL")
	adminPassword := os.Getenv("PB_ADMIN_PASSWORD")

	adminLogin := AdminLogin{
		Email:    adminEmail,
		Password: adminPassword,
	}

	jsonValue, _ := json.Marshal(adminLogin)

	resp, err := http.Post(
		os.Getenv("PB_API_URL")+"admins/auth-with-password",
		"application/json",
		bytes.NewBuffer(jsonValue),
	)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var adminResponse AdminAuthResponse

	err = json.NewDecoder(resp.Body).Decode(&adminResponse)
	if err != nil {
		panic(err)
	}

	return &Client{
		BaseURL:   os.Getenv("PB_API_URL"),
		AuthToken: adminResponse.Token,
	}
}
