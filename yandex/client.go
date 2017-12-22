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
	url, err := url.Parse(baseURL)
	if err != nil {
		panic(err)
	}
	return &Client{url, http.DefaultClient}
}
