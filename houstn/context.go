package houstn

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/houstn/terraform-provider-houstn/houstn/client"
)

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	organisation := d.Get("organisation").(string)
	token := d.Get("token").(string)

	var host *string

	hVal, ok := d.GetOk("host")

	if ok {
		tempHost := hVal.(string)
		host = &tempHost
	}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	c, err := client.NewClient(host, &organisation, &token)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create houstn client",
			Detail:   fmt.Sprintf("Unable to authenticate user for authenticated houstn client: %s", err),
		})

		return nil, diags
	}

	return c, diags
}
