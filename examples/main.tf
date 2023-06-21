terraform {
  required_providers {
    provider = {
      version = "0.1"
      source = "hashicorp.com/QuentinN42/provider"
    }
  }
}

provider "provider" { }

resource "provider_adh" "john" {
  name = "john"
  naissance = "15 juin 2001"
  email = "john@rezel.net"
  adresse = "1 rue de la paix"
}
