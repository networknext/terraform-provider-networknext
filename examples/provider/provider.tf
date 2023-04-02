# Configure local network next with example buyer

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

resource "networknext_customer" "example" {
  name = "Example Customer"
  code = "example"
}

resource "networknext_seller" "example" {
  name = "example"
}

resource "networknext_datacenter" "example" {
  name = "example"
  seller_id = networknext_seller.example.id
  latitude = 100
  longitude = 50
}

resource "networknext_relay" "example" {
  name = "example.relay"
  datacenter_id = networknext_datacenter.example.id
  public_ip = "127.0.0.1"
  public_key_base64="9SKtwe4Ear59iQyBOggxutzdtVLLc1YQ2qnArgiiz14="
  private_key_base64="lypnDfozGRHepukundjYAF5fKY1Tw2g7Dxh0rAgMCt8="
}

resource "networknext_route_shader" example {
  name = "example"
  force_next = true
}

resource "networknext_buyer" "example" {
  name = "Example Buyer"
  customer_id = networknext_customer.example.id
  route_shader_id = networknext_route_shader.example.id
  public_key_base64 = "leN7D7+9vr24uT4f1Ba8PEEvIQA/UkGZLlT+sdeLRHKsVqaZq723Zw=="
}

resource "networknext_buyer_datacenter_settings" "example" {
  buyer_id = networknext_buyer.example.id
  datacenter_id = networknext_datacenter.example.id
  enable_acceleration = true
}
