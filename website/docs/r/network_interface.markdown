---
subcategory: "VPC"
layout: "aws"
page_title: "AWS: aws_network_interface"
description: |-
  Provides an Elastic network interface (ENI) resource.
---

# Resource: aws_network_interface

Provides an Elastic network interface (ENI) resource.

## Example Usage

```terraform
resource "aws_network_interface" "test" {
  subnet_id       = aws_subnet.public_a.id
  private_ips     = ["10.0.0.50"]
  security_groups = [aws_security_group.web.id]

  attachment {
    instance     = aws_instance.test.id
    device_index = 1
  }
}
```

## Argument Reference

The following arguments are supported:

* `subnet_id` - (Required) Subnet ID to create the ENI in.
* `description` - (Optional) A description for the network interface.
* `private_ips` - (Optional) List of private IPs to assign to the ENI.
* `private_ips_count` - (Optional) Number of secondary private IPs to assign to the ENI. The total number of private IPs will be 1 + private_ips_count, as a primary private IP will be assiged to an ENI by default.
* `ipv6_addresses` - (Optional) One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. You can't use this option if you're specifying `ipv6_address_count`.
* `ipv6_address_count` - (Optional) The number of IPv6 addresses to assign to a network interface. You can't use this option if specifying specific `ipv6_addresses`. If your subnet has the AssignIpv6AddressOnCreation attribute set to `true`, you can specify `0` to override this setting.
* `security_groups` - (Optional) List of security group IDs to assign to the ENI.
* `attachment` - (Optional) Block to define the attachment of the ENI. Documented below.
* `source_dest_check` - (Optional) Whether to enable source destination checking for the ENI. Default true.

-> **NOTE:** Changing `interface_type` will cause the resource to be destroyed and re-created.

* `interface_type` - (Optional) Type of network interface to create. Set to `efa` for Elastic Fabric Adapter.
* `tags` - (Optional) A map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

The `attachment` block supports:

* `instance` - (Required) ID of the instance to attach to.
* `device_index` - (Required) Integer to define the devices index.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the network interface.
* `mac_address` - The MAC address of the network interface.
* `private_dns_name` - The private DNS name of the network interface (IPv4).
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](/docs/providers/aws/index.html#default_tags-configuration-block).

## Import

Network Interfaces can be imported using the `id`, e.g.,

```
$ terraform import aws_network_interface.test eni-e5aa89a3
```
