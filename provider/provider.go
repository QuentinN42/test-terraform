package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
		},
		ResourcesMap: map[string]*schema.Resource{
			"provider_adh": resourceAdh(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"provider_add": dataSourceAdd(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}
