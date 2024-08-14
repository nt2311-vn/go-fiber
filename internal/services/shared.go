package services

type RecordResponse struct {
	Page       int `json:"page"`
	PerPage    int `json:"perPage"`
	TotalPages int `json:"totalPages"`
	TotalItems int `json:"totalItems"`
	Items      []struct {
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
	} `json:"items"`
}
