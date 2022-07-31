package houstn

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	h "github.com/houstn/terraform-provider-houstn/houstn/client"
)

func resourceEnvironment() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceEnvironmentCreate,
		ReadContext:   resourceEnvironmentRead,
		UpdateContext: resourceEnvironmentUpdate,
		DeleteContext: resourceEnvironmentDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"environment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"order": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"group": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceEnvironmentCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*h.Client)

	var diags diag.Diagnostics

	id := d.Get("environment_id").(string)

	environment := h.Environment{
		ID: id,
	}

	name, ok := d.GetOk("name")
	if ok {
		environment.Name = name.(string)
	}

	group, ok := d.GetOk("group")
	if ok {
		environment.Group = group.(string)
	}

	order, ok := d.GetOk("order")
	if ok {
		environment.Order = order.(int)
	}

	e, err := c.CreateEnvironment(id, environment)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(e.ID)

	return diags
}

func resourceEnvironmentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*h.Client)

	var diags diag.Diagnostics

	id := d.Id()

	e, err := c.GetEnvironment(id)

	if err != nil {
		return diag.FromErr(err)
	}

	properties := map[string]interface{}{
		"id":    e.ID,
		"name":  e.Name,
		"order": e.Order,
		"group": e.Group,
	}

	for key, value := range properties {
		if err := d.Set(key, value); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

func resourceEnvironmentUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*h.Client)

	id := d.Get("environment_id").(string)

	environment := h.Environment{
		ID: id,
	}

	name, ok := d.GetOk("name")
	if ok {
		environment.Name = name.(string)
	}

	group, ok := d.GetOk("group")
	if ok {
		environment.Group = group.(string)
	}

	order, ok := d.GetOk("order")
	if ok {
		environment.Order = order.(int)
	}

	_, err := c.UpdateEnvironment(id, environment)

	if err != nil {
		return diag.FromErr(err)
	}

	return resourceEnvironmentRead(ctx, d, m)
}

func resourceEnvironmentDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*h.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	_, err := c.DeleteEnvironment(id)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
