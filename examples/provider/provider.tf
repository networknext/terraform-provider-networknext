# Configure local network next with test customer

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

resource "networknext_relay" "test" {
  name = "test.relay"
  datacenter_id = networknext_datacenter.test.id
  public_ip = "127.0.0.1"
  public_key_base64="9SKtwe4Ear59iQyBOggxutzdtVLLc1YQ2qnArgiiz14="
  private_key_base64="lypnDfozGRHepukundjYAF5fKY1Tw2g7Dxh0rAgMCt8="
}

resource "networknext_route_shader" test {
  name = "test"
  force_next = true
}

resource "networknext_buyer" "test" {
  name = "Test Buyer"
  customer_id = networknext_customer.test.id
  route_shader_id = networknext_route_shader.test.id
  public_key_base64 = "leN7D7+9vr24uT4f1Ba8PEEvIQA/UkGZLlT+sdeLRHKsVqaZq723Zw=="
}

resource "networknext_buyer_datacenter_settings" "test" {
  buyer_id = networknext_buyer.test.id
  datacenter_id = networknext_datacenter.test.id
  enable_acceleration = true
}
