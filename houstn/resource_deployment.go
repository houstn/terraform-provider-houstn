package houstn

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	h "github.com/houstn/terraform-provider-houstn/houstn/client"
)

func resourceDeployment() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDeploymentCreate,
		ReadContext:   resourceDeploymentRead,
		UpdateContext: resourceDeploymentUpdate,
		DeleteContext: resourceDeploymentDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"environment": {
				Type:     schema.TypeString,
				Required: true,
			},
			"application": {
				Type:     schema.TypeString,
				Required: true,
			},
			"url": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceDeploymentCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*h.Client)

	var diags diag.Diagnostics

	id := d.Get("environment").(string)
	environment := id
	application := d.Get("application").(string)
	url := d.Get("url").(string)

	deployment := h.Deployment{
		ID:          id,
		Environment: environment,
		Application: application,
		URL:         url,
	}

	de, err := c.CreateDeployment(application, deployment)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(de.ID)

	return diags
}

func resourceDeploymentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*h.Client)

	var diags diag.Diagnostics

	id := d.Id()
	application := d.Get("application").(string)

	de, err := c.GetDeployment(application, id)

	if err != nil {
		return diag.FromErr(err)
	}

	properties := map[string]interface{}{
		"id":          de.ID,
		"environment": de.Environment,
		"application": de.Application,
		"url":         de.URL,
	}

	for key, value := range properties {
		if err := d.Set(key, value); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

func resourceDeploymentUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*h.Client)

	id := d.Get("environment").(string)
	environment := id
	application := d.Get("application").(string)
	url := d.Get("url").(string)

	deployment := h.Deployment{
		ID:          id,
		Environment: environment,
		Application: application,
		URL:         url,
	}

	_, err := c.UpdateDeployment(application, deployment)

	if err != nil {
		return diag.FromErr(err)
	}

	return resourceDeploymentRead(ctx, d, m)
}

func resourceDeploymentDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*h.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()
	application := d.Get("application").(string)

	_, err := c.DeleteDeployment(application, id)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
