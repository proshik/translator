package yandex

import (
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL *url.URL
	*http.Client
}

func NewClient(baseURL string) *Client {
	u, err := url.Parse(baseURL)
	if err != nil {
		panic(err)
	}
	return &Client{u, http.DefaultClient}
}
