# Manage example datacenter.

resource "networknext_customer" "example" {
  name = "Example Customer"
  code = "example"
  live = true
  debug = true
}

resource "networknext_route_shader" example {
  name = "example"
  force_next = true
}

resource "networknext_buyer" "example" {
  name = "example"
  customer = networknext_customer.example.id
  route_shader = networknext_route_shader.example.id
}

resource "networknext_datacenter" "example" {
  name = "example"
  seller_id = networknext_seller.example.id
  latitude = 100
  longitude = 50
}
