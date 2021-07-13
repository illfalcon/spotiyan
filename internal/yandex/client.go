package yandex

import "net/http"

const (
	host   = "api.music.yandex.net"
	port   = 443
	scheme = "https"
)

type Client struct {
	httpClient *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	return &Client{httpClient: httpClient}
}
