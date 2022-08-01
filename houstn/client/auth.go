package client

import "fmt"

type AuthResponse struct {
	Token string `json:"token"`
}

func (c *Client) Login(organisation, token string) (string, error) {
	var authResponse AuthResponse

	err := c.
		SetBasicAuth(organisation, token).
		Post("login", "", &authResponse)

	if err != nil {
		return "", err
	}

	if authResponse.Token == "" {
		return "", fmt.Errorf("invalid user credentials")
	}

	return authResponse.Token, nil
}
