package client

import (
	"fmt"
)

func (c *Client) GetDeployments(application string) ([]Deployment, error) {
	var deployments []Deployment

	err := c.Get(fmt.Sprintf("%s/applications/%s/deployments", c.Organisation, application), &deployments)

	if err != nil {
		return nil, err
	}

	return deployments, nil
}

func (c *Client) CreateDeployment(application string, deployment Deployment) (*Deployment, error) {
	var deploy Deployment

	err := c.Post(fmt.Sprintf("%s/applications/%s/deployments/%s", c.Organisation, application, deployment.ID), deployment, &deploy)

	if err != nil {
		return nil, err
	}

	return &deploy, nil
}

func (c *Client) UpdateDeployment(application string, deployment Deployment) (*Deployment, error) {
	var deploy Deployment

	err := c.Put(fmt.Sprintf("%s/applications/%s/deployments/%s", c.Organisation, application, deployment.ID), deployment, &deploy)

	if err != nil {
		return nil, err
	}

	return &deploy, nil
}

func (c *Client) GetDeployment(application, id string) (*Deployment, error) {
	var deploy Deployment

	err := c.Get(fmt.Sprintf("%s/applications/%s/deployments/%s", c.Organisation, application, id), &deploy)

	if err != nil {
		return nil, err
	}

	return &deploy, nil
}

func (c *Client) DeleteDeployment(application, id string) (interface{}, error) {
	var result interface{}

	err := c.Delete(fmt.Sprintf("%s/applications/%s/deployments/%s", c.Organisation, application, id), &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
