package airvisual

import "net/http"

type Client struct {
	client *http.Client
	APIUrl string
	APIKey string
}

// func New(apiKey string, opts ...Option) *Client {
// 	client := Client{
// 		client: http.DefaultClient,
// 		APIKey: apiKey,
// 	}

// 	for _, opt := range opts {
// 		opt(&client)
// 	}

// 	return &client
// }

type Option func(*Client)

func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.client = client
	}
}
