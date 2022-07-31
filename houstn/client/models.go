package client

type Application struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type Environment struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Order int    `json:"order"`
	Group string `json:"group"`
}

type Deployment struct {
	ID          string `json:"id"`
	Application string `json:"application"`
	Environment string `json:"environment"`
	URL         string `json:"url"`
}
