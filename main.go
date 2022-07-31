package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/houstn/terraform-provider-houstn/houstn"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: houstn.Provider,
	})
}
