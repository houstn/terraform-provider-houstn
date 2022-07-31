package client

import (
	"fmt"
)

func (c *Client) GetApplications() ([]Application, error) {
	var applications []Application

	err := c.Get(fmt.Sprintf("%s/applications", c.Organisation), &applications)

	if err != nil {
		return nil, err
	}

	return applications, nil
}

func (c *Client) CreateApplication(id string, application Application) (*Application, error) {
	var app Application

	err := c.Post(fmt.Sprintf("%s/applications/%s", c.Organisation, id), application, &app)

	if err != nil {
		return nil, err
	}

	return &app, nil
}

func (c *Client) UpdateApplication(id string, application Application) (*Application, error) {
	var app Application

	err := c.Put(fmt.Sprintf("%s/applications/%s", c.Organisation, id), application, &app)

	if err != nil {
		return nil, err
	}

	return &app, nil
}

func (c *Client) GetApplication(id string) (*Application, error) {
	var app Application

	err := c.Get(fmt.Sprintf("%s/applications/%s", c.Organisation, id), &app)

	if err != nil {
		return nil, err
	}

	return &app, nil
}

func (c *Client) DeleteApplication(id string) (interface{}, error) {
	var result interface{}

	err := c.Delete(fmt.Sprintf("%s/applications/%s", c.Organisation, id), &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
