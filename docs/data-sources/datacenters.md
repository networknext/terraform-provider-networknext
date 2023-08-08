---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "networknext_datacenters Data Source - networknext"
subcategory: ""
description: |-
  Fetches the list of datacenters.
---

# networknext_datacenters (Data Source)

Fetches the list of datacenters.



<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `datacenters` (Attributes List) (see [below for nested schema](#nestedatt--datacenters))

<a id="nestedatt--datacenters"></a>
### Nested Schema for `datacenters`

Read-Only:

- `id` (Number) The id of the datacenter. Automatically generated when datacenters are created.
- `latitude` (Number) The approximate longitude of the datacenter.
- `longitude` (Number) The approximate latitude of the datacenter.
- `name` (String) The name of the datacenter. Must be in the format [seller].[location] with optional datacenter number. For example: google.losangeles.1, vultr.chicago, amazon.virginia.2
- `native_name` (String) The native datacenter name. Used to associate the network next name of a datacenter with the native name of the datacenter on that platform. For example, 'google.taiwan.1' has a native name of 'asia-east1-a'
- `notes` (String) Optional notes about this datacenter.
- `seller_id` (Number) The id of the seller this relay belongs to.

