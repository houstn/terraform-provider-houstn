package houstn

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"organisation": {
			Type:        schema.TypeString,
			Required:    true,
			DefaultFunc: schema.EnvDefaultFunc("HOUSTN_ORGANISATION", nil),
		},
		"host": {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("HOUSTN_HOST", "http://localhost:7070"),
		},
		"token": {
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("HOUSTN_API_TOKEN", nil),
		},
	}
}
