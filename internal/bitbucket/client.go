package bitbucket

import (
	"io"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	BaseURL string
	Token   string
}

var client = &http.Client{
	Timeout: 5 * time.Second,
}

func (b *Client) Do(method string, path string, body *string) (*http.Response, error) {
	var bytesBody io.Reader
	if body != nil && *body != "" {
		bytesBody = strings.NewReader(*body)
	}
	req, err := http.NewRequest(method, b.BaseURL+path, bytesBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+b.Token)
	req.Header.Set("Accept", "application/json")

	return client.Do(req)
}
