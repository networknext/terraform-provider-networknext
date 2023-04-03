terraform {
  required_providers {
    networknext = {
      source = "networknext.com/networknext/networknext"
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
}

resource "networknext_seller" "test" {
  name = "test"
}

resource "networknext_datacenter" "test" {
  name = "test"
  seller_id = networknext_seller.test.id
  latitude = 100
  longitude = 50
}

resource "networknext_relay_keypair" "test" {}

resource "networknext_relay" "test" {
  name = "test.relay"
  datacenter_id = networknext_datacenter.test.id
  public_ip = "127.0.0.1"
  public_key_base64=networknext_relay_keypair.test.public_key_base64
  private_key_base64=networknext_relay_keypair.test.private_key_base64
}

resource "networknext_route_shader" test {
  name = "test"
}

resource "networknext_buyer_keypair" "test" {}

resource "networknext_buyer" "test" {
  name = "Test Buyer"
  customer_id = networknext_customer.test.id
  route_shader_id = networknext_route_shader.test.id
  public_key_base64 = networknext_buyer_keypair.test.public_key_base64
}

resource "networknext_buyer_datacenter_settings" "test" {
  buyer_id = networknext_buyer.test.id
  datacenter_id = networknext_datacenter.test.id
  enable_acceleration = true
}

data "networknext_customers" "test" {
  depends_on = [
    networknext_customer.test,
  ]
}

data "networknext_sellers" "test" {
  depends_on = [
    networknext_seller.test,
  ]
}

data "networknext_datacenters" "test" {
  depends_on = [
    networknext_datacenter.test,
  ]
}

data "networknext_relays" "test" {
  depends_on = [
    networknext_relay.test,
  ]
}

data "networknext_route_shaders" "test" {
  depends_on = [
    networknext_route_shader.test,
  ]
}

data "networknext_buyers" "test" {
  depends_on = [
    networknext_buyer.test,
  ]
}

data "networknext_buyer_datacenter_settings" "test" {
  depends_on = [
    networknext_buyer_datacenter_settings.test,
  ]
}

data "networknext_buyer_keypairs" "test" {
  depends_on = [
    networknext_buyer_keypairs.test,
  ]
}

data "networknext_relay_keypairs" "test" {
  depends_on = [
    networknext_relay_keypairs.test,
  ]
}

output "customers" {
  value = data.networknext_customers.test
}

output "sellers" {
  value = data.networknext_sellers.test
}

output "datacenters" {
  value = data.networknext_datacenters.test
}

output "relays" {
  value = data.networknext_relays.test
}

output "route_shaders" {
  value = data.networknext_route_shaders.test
}

output "buyers" {
  value = data.networknext_buyers.test
}

output "buyer_datacenter_settings" {
  value = data.networknext_buyer_datacenter_settings.test
}

output "buyer_keypairs" {
  value = data.networknext_buyer_keypairs.test
}

output "relay_keypairs" {
  value = data.networknext_relay_keypairs.test
}
