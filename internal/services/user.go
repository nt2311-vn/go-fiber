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

type UserResponse struct {
	ID              string `json:"id"`
	CollectionID    string `json:"collectionId"`
	CollectionName  string `json:"collectionName"`
	Username        string `json:"username"`
	Verified        bool   `json:"verified"`
	EmailVisibility bool   `json:"emailVisibility"`
	Email           string `json:"email"`
	Created         string `json:"created"`
	Updated         string `json:"updated"`
	Name            string `json:"name"`
	Avatar          string `json:"avatar"`
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
		return false, fmt.Errorf("internal error cannot check your email account")
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

func (c *Client) LoginUser(email, password string) (*UserAuthReponse, error) {
	loginPayload := LoginPayload{
		Email:    email,
		Password: password,
	}

	jsonPayload, err := json.Marshal(loginPayload)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal payload")
	}

	resp, err := http.Post(
		c.BaseURL+"collections/users/auth-with-password",
		"application/json",
		bytes.NewBuffer(jsonPayload),
	)
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid email or password")
	}

	var loginResponse UserAuthReponse
	err = json.NewDecoder(resp.Body).Decode(&loginResponse)
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	return &loginResponse, nil
}
