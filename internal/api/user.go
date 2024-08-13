package api

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type RegisterPayload struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

func (c *Client) RegisterUser(email, password, passwordConfirm string) error {
	payload := RegisterPayload{
		Email:           email,
		Password:        password,
		PasswordConfirm: passwordConfirm,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(
		c.BaseURL+"collections/users/records",
		"application/json",
		bytes.NewBuffer(jsonPayload),
	)

	if err != nil || resp.StatusCode != http.StatusOK {
		return err
	}

	defer resp.Body.Close()

	return nil
}
