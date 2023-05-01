
terraform {
  required_providers {
    accelerate = {
      source = "networknext.com/networknext/accelerate"
    }
  }
}

provider "accelerate" {
  hostname = "http://localhost:50000"
  api_key  = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZGF0YWJhc2UiOnRydWUsInBvcnRhbCI6dHJ1ZX0.QFPdb-RcP8wyoaOIBYeB_X6uA7jefGPVxm2VevJvpwU"
}

resource "accelerate_customer" "test" {
  name = "Test Customer"
  code = "test"
}

resource "accelerate_seller" "test" {
  name = "test"
}

resource "accelerate_datacenter" "test" {
  name = "test"
  seller_id = accelerate_seller.test.id
  latitude = 100
  longitude = 50
}

resource "accelerate_relay_keypair" "test" {}

resource "accelerate_relay" "test" {
  name = "test.relay"
  datacenter_id = accelerate_datacenter.test.id
  public_ip = "127.0.0.1"
  public_key_base64=accelerate_relay_keypair.test.public_key_base64
  private_key_base64=accelerate_relay_keypair.test.private_key_base64
}

resource "accelerate_route_shader" test {
  name = "test"
}

resource "accelerate_buyer_keypair" "test" {}

resource "accelerate_buyer" "test" {
  name = "Test Buyer"
  customer_id = accelerate_customer.test.id
  route_shader_id = accelerate_route_shader.test.id
  public_key_base64 = accelerate_buyer_keypair.test.public_key_base64
}

resource "accelerate_buyer_datacenter_settings" "test" {
  buyer_id = accelerate_buyer.test.id
  datacenter_id = accelerate_datacenter.test.id
  enable_acceleration = true
}

data "accelerate_customers" "test" {
  depends_on = [
    accelerate_customer.test,
  ]
}

data "accelerate_sellers" "test" {
  depends_on = [
    accelerate_seller.test,
  ]
}

data "accelerate_datacenters" "test" {
  depends_on = [
    accelerate_datacenter.test,
  ]
}

data "accelerate_relays" "test" {
  depends_on = [
    accelerate_relay.test,
  ]
}

data "accelerate_route_shaders" "test" {
  depends_on = [
    accelerate_route_shader.test,
  ]
}

data "accelerate_buyers" "test" {
  depends_on = [
    accelerate_buyer.test,
  ]
}

data "accelerate_buyer_datacenter_settings" "test" {
  depends_on = [
    accelerate_buyer_datacenter_settings.test,
  ]
}

data "accelerate_buyer_keypairs" "test" {
  depends_on = [
    accelerate_buyer_keypairs.test,
  ]
}

data "accelerate_relay_keypairs" "test" {
  depends_on = [
    accelerate_relay_keypairs.test,
  ]
}

output "customers" {
  value = data.accelerate_customers.test
}

output "sellers" {
  value = data.accelerate_sellers.test
}

output "datacenters" {
  value = data.accelerate_datacenters.test
}

output "relays" {
  value = data.accelerate_relays.test
}

output "route_shaders" {
  value = data.accelerate_route_shaders.test
}

output "buyers" {
  value = data.accelerate_buyers.test
}

output "buyer_datacenter_settings" {
  value = data.accelerate_buyer_datacenter_settings.test
}

output "buyer_keypairs" {
  value = data.accelerate_buyer_keypairs.test
}

output "relay_keypairs" {
  value = data.accelerate_relay_keypairs.test
}
