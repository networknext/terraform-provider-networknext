# Manage example customer.
resource "networknext_customer" "example" {
  name = "Example Customer"
  code = "example"
  live = true
  debug = true
}
