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
  name = "test"
  native_name = "test native name"
  seller_id = networknext_seller.test.id
  latitude = 100
  longitude = 50
  notes = ""
}

resource "networknext_relay" "test" {
  name = "test.relay"
  datacenter_id = networknext_datacenter.test.id
  public_ip = "127.0.0.1"
  public_port = 40000
  internal_ip = "0.0.0.0"
  internal_port = 0
  internal_group = ""
  ssh_ip = "127.0.0.1"
  ssh_port = 22
  ssh_user = "ubuntu"
  public_key_base64="9SKtwe4Ear59iQyBOggxutzdtVLLc1YQ2qnArgiiz14="
  private_key_base64="lypnDfozGRHepukundjYAF5fKY1Tw2g7Dxh0rAgMCt8="
  version = "1.0.19"
  mrc = 0
  port_speed = 1000
  max_sessions = 100
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
