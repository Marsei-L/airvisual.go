package airvisual

import (
	"net/http"
)

type Api struct {
	APIKey string
	url    string
	client *http.Client
}

type errMsg struct {
	Message string
}

type baseResponse struct {
	Status string `json:"status"`
}
