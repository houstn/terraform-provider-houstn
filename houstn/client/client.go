package client

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	BaseURL      string
	Organisation string
	Token        string

	HTTPClient *resty.Client
}

const HostUrl string = "https://houstn.io"

func NewClient(host, organisation, token *string) (*Client, error) {
	c := Client{
		HTTPClient:   resty.New(),
		BaseURL:      fmt.Sprintf("%s/api", HostUrl),
		Organisation: *organisation,
	}

	if host != nil {
		c.BaseURL = fmt.Sprintf("%s/api", *host)
	}

	jwt, err := c.Login(*organisation, *token)

	if err != nil {
		return nil, err
	}

	c.Token = jwt
	c.SetToken(c.Token)

	return &c, nil
}

func (c *Client) SetBasicAuth(user, password string) *Client {
	c.HTTPClient.SetBasicAuth(user, password)

	return c
}

func (c *Client) SetToken(token string) *Client {
	c.HTTPClient.SetAuthScheme("Bearer").SetAuthToken(token)

	return c
}

func (c *Client) Get(path string, res interface{}) error {
	_, err := c.HTTPClient.R().
		EnableTrace().
		SetResult(&res).
		Get(fmt.Sprintf("%s/%s", c.BaseURL, path))

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Post(path string, body interface{}, res interface{}) error {
	_, err := c.HTTPClient.R().
		EnableTrace().
		SetBody(body).
		SetResult(&res).
		Post(fmt.Sprintf("%s/%s", c.BaseURL, path))

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Put(path string, body interface{}, res interface{}) error {
	_, err := c.HTTPClient.R().
		EnableTrace().
		SetBody(body).
		SetResult(&res).
		Put(fmt.Sprintf("%s/%s", c.BaseURL, path))

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Delete(path string, res interface{}) error {
	_, err := c.HTTPClient.R().
		EnableTrace().
		SetResult(&res).
		Delete(fmt.Sprintf("%s/%s", c.BaseURL, path))

	if err != nil {
		return err
	}

	return nil
}
