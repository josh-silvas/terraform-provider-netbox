---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netbox_service Resource - terraform-provider-netbox"
subcategory: ""
description: |-
  From the official documentation https://docs.netbox.dev/en/stable/core-functionality/services/#services:
  A service represents a layer four TCP or UDP service available on a device or virtual machine. For example, you might want to document that an HTTP service is running on a device. Each service includes a name, protocol, and port number; for example, "SSH (TCP/22)" or "DNS (UDP/53)."
  A service may optionally be bound to one or more specific IP addresses belonging to its parent device or VM. (If no IP addresses are bound, the service is assumed to be reachable via any assigned IP address.
---

# netbox_service (Resource)

From the [official documentation](https://docs.netbox.dev/en/stable/core-functionality/services/#services):

> A service represents a layer four TCP or UDP service available on a device or virtual machine. For example, you might want to document that an HTTP service is running on a device. Each service includes a name, protocol, and port number; for example, "SSH (TCP/22)" or "DNS (UDP/53)."
>
> A service may optionally be bound to one or more specific IP addresses belonging to its parent device or VM. (If no IP addresses are bound, the service is assumed to be reachable via any assigned IP address.

## Example Usage

```terraform
// Assumes Netbox already has a VM whos name matches 'dc-west-myvm-20'
data "netbox_virtual_machine" "myvm" {
  name_regex = "dc-west-myvm-20"
}

resource "netbox_service" "ssh" {
  name               = "ssh"
  ports              = [22]
  protocol           = "TCP"
  virtual_machine_id = data.netbox_virtual_machine.myvm.id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)
- `protocol` (String)
- `virtual_machine_id` (Number)

### Optional

- `port` (Number, Deprecated)
- `ports` (Set of Number)

### Read-Only

- `id` (String) The ID of this resource.

