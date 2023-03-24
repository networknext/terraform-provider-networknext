terraform {
  required_providers {
    networknext = {
      source = "hashicorp.com/edu/networknext"
    }
  }
}

provider "networknext" {
  hostname = "http://localhost:50000"
  api_key  = "test123"
}

data "networknext_customer" "example" {}

output "example_customer" {
  value = data.networknext_customer.example
}
