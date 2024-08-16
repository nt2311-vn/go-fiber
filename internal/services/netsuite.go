package services

type NSClient struct {
	BaseURL string
}

func NewNSClient() *NSClient {
	return &NSClient{
		BaseURL: "http://localhost:8080",
	}
}
