---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "networknext_relays Data Source - networknext"
subcategory: ""
description: |-
  Fetches the list of relays.
---

# networknext_relays (Data Source)

Fetches the list of relays.



<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `relays` (Attributes List) (see [below for nested schema](#nestedatt--relays))

<a id="nestedatt--relays"></a>
### Nested Schema for `relays`

Read-Only:

- `datacenter_id` (Number) The id of the datacenter this relay is in.
- `id` (Number) The id of the relay. Automatically generated when relays are created.
- `internal_group` (String) The internal group of the relay. Relays only communicate with internal addresses when they are in the same internal group. Use this when sellers (like amazon) can only use private addresses between a subset of relays. For amazon, set this string to the relay region, then relays will only use the internal addresses for other relays in the same region. Default value is "" which lets all relays for a seller communicate with each other via internal address when provided.
- `internal_ip` (String) The internal IP address of the relay. Use this only when a seller has an internal network between multiple relays that provide cost or performance benefits. Default value is "0.0.0.0" which disables the internal address.
- `internal_port` (Number) The internal UDP port of the relay. Default value is 0, which uses the same port as the public port. Only set if you need to override to a different port.
- `max_sessions` (Number) The maximum number of sessions allowed across this relay at any time. Once this number is exceeded, the network next backend will not send additional sessions across this relay. Default is zero which means unlimited.
- `mrc` (Number) Monthly recurring cost for this relay in USD. Useful for keeping track of how much relays cost. Default value is zero.
- `name` (String) The name of the relay. The name must be in the form of [datacenter]<.variant>. For example, "google.losangeles.1" (same name as datacenter for one relay in the datacenter), or "amazon.virginia.2.a", "amazon.virginia.2.b" (for two relays in the same datacenter).
- `notes` (String) Optional free form notes to store information about this relay.
- `port_speed` (Number) The speed of the network port in megabits per-second (mbps). Useful for keeping track of relays with 1gbps (1,000) vs. 10gbps (10,000) speeds.
- `private_key_base64` (String) The relay private key as base64.
- `public_ip` (String) The public IP address of the relay. For example, "45.23.66.10".
- `public_key_base64` (String) The relay public key as base64.
- `public_port` (Number) The public UDP port of the relay. By default it is 40000. Make sure this port is open on the firewall to receive UDP packets.
- `ssh_ip` (String) The SSH address of the relay. The default value is "0.0.0.0" which uses the public address to SSH into. Set this only if you have a specific management IP address that you need to use when SSH'ing into a relay.
- `ssh_port` (Number) The TCP port to SSH into. Set to 22 by default.
- `ssh_user` (String) The username to SSH in as. Defaults to root. Other common usernames include "ubuntu". Seller specific.
- `version` (String) The version of the relay. Default value is "". Leave as empty, and you can manage relay versions manually with the next tool. Set to a specific versioen, eg. "1.0.19" and the relay will automatically upgrade itself to the relay binary in google cloud storage corresponding to the version you set.


