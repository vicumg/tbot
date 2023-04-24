package telegram

import (
	"net/http"
)

type Client struct {
	host     string
	basepath string
	client   http.Client
}

func New(host string, token string) Client {
	return Client{
		host:     host,
		basepath: newBathPath(token),
		client:   http.Client{},
	}
}

func newBathPath(token string) string {
	return "bot" + token
}

func (c *Client) Updates() {

}

func (c *Client) SendMessages() {

}
