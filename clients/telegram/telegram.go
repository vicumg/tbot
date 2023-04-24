package telegram

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"tbot/lib/reqerr"
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

func (c *Client) Updates(offset int, limit int) ([]Update, err){
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))
	
}

func (c *Client) doRequest(method string, query url.Values) ([]byte, err){

	const errMsg := "cant do request"

	u := url.URL{
		Scheme: "https",
		Host: c.host,
		Path: path.Join(c.basepath , method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil) 

	if err != nil {
		return nil, reqerr.Wrap(errMsg, err)
	}

	req.URL.RawQuery := query.Encode()

	resp, err = c.client.Do(req)

	if err != nil{
		return nil, reqerr.Wrap(errMsg, err)
	}

}

func (c *Client) SendMessages() {

}
