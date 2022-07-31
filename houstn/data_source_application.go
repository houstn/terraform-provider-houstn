package houstn

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	h "github.com/houstn/terraform-provider-houstn/houstn/client"
	"strconv"
	"time"
)

func dataSourceApplications() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceApplicationsRead,
		Schema: map[string]*schema.Schema{
			"applications": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"active": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceApplicationsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*h.Client)

	var diags diag.Diagnostics

	apps, err := c.GetApplications()

	if err != nil {
		return diag.FromErr(err)
	}

	applications := make([]map[string]interface{}, 0)

	for _, a := range apps {
		application := make(map[string]interface{})

		application["id"] = a.ID
		application["name"] = a.Name
		application["active"] = a.Active

		applications = append(applications, application)
	}

	if err := d.Set("applications", applications); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
