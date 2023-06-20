package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAdh() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAdhCreate,
		ReadContext:   resourceAdhRead,
		DeleteContext: resourceAdhDelete,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"email": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"naissance": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"adresse": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"debut_cotisation": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fin_cotisation": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAdhCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	res, err := post("http://localhost:8000/create", map[string]string{
		"fullname": d.Get("name").(string),
		"email": d.Get("email").(string),
		"naissance": d.Get("naissance").(string),
		"adresse": d.Get("adresse").(string),
		"duree_months": "0",
        "duree_years": "0",
	})
	if err != nil {
		return diag.Errorf("Failed to create a Adh: %s", err.Error())
	}

	d.Set("id", res["status"])

	d.SetId(res["status"])

	return diags
}

func resourceAdhRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceAdhDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}
