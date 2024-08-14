package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type RegisterPayload struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

func (c *Client) EmailExists(email string) (bool, error) {
	escapeEmail := url.QueryEscape(email)

	queryEmailUrl := fmt.Sprintf(
		"%scollections/users/records?filter=(email='%s')",
		c.BaseURL,
		escapeEmail,
	)

	req, err := http.NewRequest(
		"GET",
		queryEmailUrl,
		nil,
	)
	if err != nil {
		return false, err
	}

	authToken, err := getAdminToken()
	if err != nil {
		return false, fmt.Errorf("cannot get admin token: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var records RecordResponse
	err = json.NewDecoder(resp.Body).Decode(&records)
	if err != nil {
		return false, err
	}

	return len(records.Items) > 0, nil
}

func (c *Client) RegisterUser(email, password, passwordConfirm string) error {
	exists, err := c.EmailExists(email)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("email already exists")
	}

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
