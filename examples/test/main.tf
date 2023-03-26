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

resource "networknext_customer" "test" {
  name = "Test Customer"
  code = "test"
  live = true
  debug = false
}

resource "networknext_seller" "test" {
  name = "test"
  customer_id = 0
}

resource "networknext_datacenter" "test" {
  name = "test.datacenter"
  native_name = "test native name"
  seller_id = networknext_seller.test.id
  latitude = 100
  longitude = 50
  notes = ""
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
