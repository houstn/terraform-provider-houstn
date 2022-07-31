package client

import (
	"fmt"
)

func (c *Client) GetEnvironments() ([]Environment, error) {
	var environments []Environment

	err := c.Get(fmt.Sprintf("%s/environments", c.Organisation), &environments)

	if err != nil {
		return nil, err
	}

	return environments, nil
}

func (c *Client) CreateEnvironment(id string, environment Environment) (*Environment, error) {
	var env Environment

	err := c.Post(fmt.Sprintf("%s/environments/%s", c.Organisation, id), environment, &env)

	if err != nil {
		return nil, err
	}

	return &env, nil
}

func (c *Client) UpdateEnvironment(id string, environment Environment) (*Environment, error) {
	var env Environment

	err := c.Put(fmt.Sprintf("%s/environments/%s", c.Organisation, id), environment, &env)

	if err != nil {
		return nil, err
	}

	return &env, nil
}

func (c *Client) GetEnvironment(id string) (*Environment, error) {
	var env Environment

	err := c.Get(fmt.Sprintf("%s/environments/%s", c.Organisation, id), &env)

	if err != nil {
		return nil, err
	}

	return &env, nil
}

func (c *Client) DeleteEnvironment(id string) (interface{}, error) {
	var result interface{}

	err := c.Delete(fmt.Sprintf("%s/environments/%s", c.Organisation, id), &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
