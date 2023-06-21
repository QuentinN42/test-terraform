package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFile(typ ResourceType) *schema.Resource {
	schema := schema.Resource{
		Schema: map[string]*schema.Schema{
			"filename": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"content": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "",
				ExactlyOneOf: []string{"content", "present"},
			},
			"present": {
				Type:         schema.TypeBool,
				Optional:     true,
				Default:      true,
				ExactlyOneOf: []string{"content", "present"},
			},
		},
	}
	schema.ReadContext = resourceFileRead

	if typ == Resource {
		schema.CreateContext = resourceFileCreate
		schema.UpdateContext = resourceFileCreate
		schema.DeleteContext = resourceFileDelete
	}

	return &schema
}

func resourceFileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	path := m.(*Context).BasePath + d.Get("filename").(string)
	d.SetId(path)

	present := d.Get("present").(bool)
	// RM file
	if !present {
		d.Set("content", "")

		err := os.Remove(path)
		if err != nil {
			if os.IsNotExist(err) {
				return diags
			}
			return diag.FromErr(err)
		}
		return diags
	}

	// Create file
	content := d.Get("content").(string)

	err := os.WriteFile(path, []byte(content), 0600)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceFileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	path := m.(*Context).BasePath + d.Get("filename").(string)
	d.SetId(path)
	d.Set("present", true)

	res, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			d.Set("present", false)
			d.Set("content", "")
			return diags
		}
		return diag.FromErr(err)
	}

	d.Set("content", string(res))

	return diags
}

func resourceFileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	path := m.(*Context).BasePath + d.Get("filename").(string)

	err := os.Remove(path)
	if err != nil {
		if os.IsNotExist(err) {
			return diags
		}
		return diag.FromErr(err)
	}

	return diags
}
