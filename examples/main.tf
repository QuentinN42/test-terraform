# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

terraform {
  required_providers {
    provider = {
      version = "0.1"
      source = "hashicorp.com/QuentinN42/provider"
    }
  }
}

provider "provider" { }


data provider_add calc {
  a = 1
  b = 2
}

output "add" {
  value = data.provider_add.calc.result
}
