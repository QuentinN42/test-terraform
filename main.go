// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"provider/provider"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	var res, err = provider.Post("http://localhost:8000/create",  map[string]string{
		"fullname": "name",
		"email": "email",
		"naissance": "naissance",
		"adresse": "adresse",
		"duree_months": "0",
        "duree_years": "0",
	})
	println(fmt.Sprintf("%v", res))
	println(fmt.Sprintf("%v", err))

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return provider.Provider()
		},
	})
}
