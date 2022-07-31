package houstn

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func providerResource() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"houstn_application": resourceApplication(),
		"houstn_environment": resourceEnvironment(),
		"houstn_deployment":  resourceDeployment(),
	}
}
