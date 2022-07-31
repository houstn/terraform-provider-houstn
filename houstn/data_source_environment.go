package houstn

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	h "github.com/houstn/terraform-provider-houstn/houstn/client"
	"strconv"
	"time"
)

func dataSourceEnvironments() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceEnvironmentsRead,
		Schema: map[string]*schema.Schema{
			"environments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"order": {
							Type:     schema.TypeInt,
							Computed: true,
							Optional: true,
						},
						"group": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceEnvironmentsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*h.Client)

	var diags diag.Diagnostics

	apps, err := c.GetEnvironments()

	if err != nil {
		return diag.FromErr(err)
	}

	environments := make([]map[string]interface{}, 0)

	for _, a := range apps {
		environment := make(map[string]interface{})

		environment["id"] = a.ID
		environment["name"] = a.Name
		environment["order"] = a.Order
		environment["group"] = a.Group

		environments = append(environments, environment)
	}

	if err := d.Set("environments", environments); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
