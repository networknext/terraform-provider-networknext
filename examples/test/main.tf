terraform {
  required_providers {
    networknext = {
      source = "hashicorp.com/edu/networknext"
    }
  }
}

provider "networknext" {
  hostname = "http://localhost:50000"
  api_key  = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZGF0YWJhc2UiOnRydWUsInBvcnRhbCI6dHJ1ZX0.QFPdb-RcP8wyoaOIBYeB_X6uA7jefGPVxm2VevJvpwU"
}

data "networknext_customers" "example" {}

data "networknext_buyers" "example" {}

data "networknext_sellers" "example" {}

data "networknext_datacenters" "example" {}

data "networknext_relays" "example" {}

data "networknext_route_shaders" "example" {}

output "customers" {
  value = data.networknext_customers.example
}

output "buyers" {
  value = data.networknext_buyers.example
}

output "sellers" {
  value = data.networknext_sellers.example
}

output "datacenters" {
  value = data.networknext_datacenters.example
}

output "relays" {
  value = data.networknext_relays.example
}

output "route_shaders" {
  value = data.networknext_route_shaders.example
}
