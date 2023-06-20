package provider

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAdd() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAddCalc,
		Schema: map[string]*schema.Schema{
			"a": {
				Type:         schema.TypeInt,
				Required:    true,
			},
			"b": {
				Type:         schema.TypeInt,
				Required:    true,
			},
			"result": {
				Type:         schema.TypeInt,
				Computed:    true,
			},
		},
	}
}

func dataSourceAddCalc(_ context.Context, d *schema.ResourceData, _ interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	a := d.Get("a").(int)
	b := d.Get("b").(int)

	d.SetId(strconv.Itoa(a) + "+" + strconv.Itoa(b))
	d.Set("result", a+b)

	return diags
}
