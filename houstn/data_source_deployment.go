package houstn

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	h "github.com/houstn/terraform-provider-houstn/houstn/client"
)

func dataSourceDeployments() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDeploymentsRead,
		Schema: map[string]*schema.Schema{
			"application_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"deployments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDeploymentsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*h.Client)

	var diags diag.Diagnostics

	application := d.Get("application_id").(string)

	deploys, err := c.GetDeployments(application)

	if err != nil {
		return diag.FromErr(err)
	}

	deployments := make([]map[string]interface{}, 0)

	for _, d := range deploys {
		deployment := make(map[string]interface{})

		deployment["id"] = d.ID
		deployment["url"] = d.URL

		deployments = append(deployments, deployment)
	}

	if err := d.Set("deployments", deployments); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(application)

	return diags
}
