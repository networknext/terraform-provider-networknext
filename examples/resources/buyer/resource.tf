# Manage example buyer.

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
