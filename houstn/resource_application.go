package houstn

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	h "github.com/houstn/terraform-provider-houstn/houstn/client"
)

func resourceApplication() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceApplicationCreate,
		ReadContext:   resourceApplicationRead,
		UpdateContext: resourceApplicationUpdate,
		DeleteContext: resourceApplicationDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"application_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"active": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceApplicationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*h.Client)

	var diags diag.Diagnostics

	id := d.Get("application_id").(string)
	name := d.Get("name").(string)
	active := d.Get("active").(bool)

	application := h.Application{
		ID:     id,
		Name:   name,
		Active: active,
	}

	a, err := c.CreateApplication(id, application)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(a.ID)

	return diags
}

func resourceApplicationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*h.Client)

	var diags diag.Diagnostics

	id := d.Id()

	a, err := c.GetApplication(id)

	if err != nil {
		return diag.FromErr(err)
	}

	properties := map[string]interface{}{
		"id":     a.ID,
		"name":   a.Name,
		"active": a.Active,
	}

	for key, value := range properties {
		if err := d.Set(key, value); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

func resourceApplicationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*h.Client)

	id := d.Get("application_id").(string)
	name := d.Get("name").(string)
	active := d.Get("active").(bool)

	application := h.Application{
		ID:     id,
		Name:   name,
		Active: active,
	}

	_, err := c.UpdateApplication(id, application)

	if err != nil {
		return diag.FromErr(err)
	}

	return resourceApplicationRead(ctx, d, m)
}

func resourceApplicationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*h.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	_, err := c.DeleteApplication(id)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
