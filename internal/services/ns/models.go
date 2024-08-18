package ns

type NSClient struct {
	AccessToken string
	BaseURL     string
}

type AppID struct {
	ClientID   string
	ClientSec  string
	CertID     string
	GrantType  string
	Scope      string
	AssertType string
	BaseURL    string
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
