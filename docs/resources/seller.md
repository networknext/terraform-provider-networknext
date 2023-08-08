---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "networknext_seller Resource - networknext"
subcategory: ""
description: |-
  Manages a seller.
---

# networknext_seller (Resource)

Manages a seller.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the seller. For example, "google", "amazon" or "akamai"

### Optional

- `customer_id` (Number) Optional. The id of the customer that this seller is associated with. Reserved for future functionality where customers can be both buyers and sellers. Defaults to 0.

### Read-Only

- `id` (Number) The id of the seller. Automatically generated when sellers are created.

