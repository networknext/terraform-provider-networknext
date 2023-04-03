# Manage example relay.

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
