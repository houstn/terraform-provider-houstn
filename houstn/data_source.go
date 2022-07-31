package houstn

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func providerDataSource() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"houstn_applications": dataSourceApplications(),
		"houstn_deployments":  dataSourceDeployments(),
		"houstn_environments": dataSourceEnvironments(),
	}
}
