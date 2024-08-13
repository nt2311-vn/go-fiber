package services

import (
	"os"
)

type Client struct {
	BaseURL string
}

func NewClient() *Client {
	return &Client{
		BaseURL: os.Getenv("PB_API_URL"),
	}
}
