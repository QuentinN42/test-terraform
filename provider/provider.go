package provider

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ResourceType int64

const (
	Undefined ResourceType = iota
	Resource
	DataSource
)

func Provider(name string) *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"base_path": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			fmt.Sprintf("%s_file", name): resourceFile(Resource),
		},
		DataSourcesMap: map[string]*schema.Resource{
			fmt.Sprintf("%s_file", name): resourceFile(DataSource),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

type Context struct {
	BasePath string
}

func providerConfigure(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	p := d.Get("base_path").(string)
	res, err := os.Stat(p)
	if err != nil {
		return nil, diag.Errorf("base_path must be a directory")
	}
	if !res.IsDir() {
		return nil, diag.Errorf("base_path must be a directory")
	}
	if p[len(p)-1] != '/' {
		p += "/"
	}

	return &Context{
		BasePath: p,
	}, diags
}
