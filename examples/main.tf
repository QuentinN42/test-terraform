# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

terraform {
  required_providers {
    provider = {
      version = "0.3.1"
      source = "hashicorp.com/QuentinN42/provider"
    }
  }
}

provider "provider" {
}
