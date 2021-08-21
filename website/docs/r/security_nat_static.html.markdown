---
layout: "junos"
page_title: "Junos: junos_security_nat_static"
sidebar_current: "docs-junos-resource-security-nat-static"
description: |-
  Create a security nat static (when Junos device supports it)
---

# junos_security_nat_static

Provides a security static nat resource.

## Example Usage

```hcl
# Add a static nat
resource junos_security_nat_static "demo_nat" {
  name = "nat_from_trust"
  from {
    type  = "zone"
    value = ["trust"]
  }
  rule {
    name                = "nat_192_0_2_0_25"
    destination_address = "192.0.2.0/25"
    then {
      type   = "prefix"
      prefix = "192.0.2.128/25"
    }
  }
}
```

## Argument Reference

The following arguments are supported:

- **name** (Required, String, Forces new resource)  
  The name of static nat.
- **from** (Required, Block)  
  Declare `from` configuration.
  - **type** (Required, String)  
    Type of from options.  
    Need to be `interface`, `routing-instance` or `zone`.
  - **value** (Required, Set of String)  
    Name of interface, routing-instance or zone for from options.
- **rule** (Required, Block List)  
  For each name of rule to declare.  
  See [below for nested schema](#rule-arguments).

---

### rule arguments

- **name** (Required, String)  
  Name of rule.
- **destination_address** (Required, String)  
  CIDR of destination address for rule.
- **then** (Required, Block)  
  Declare `then` configuration.
  - **type** (Required, String)  
    Type of static nat.  
    Need to be `inet` or `prefix`.
  - **prefix** (Optional, String)  
    CIDR for prefix static nat.
  - **routing_instance** (Optional, String)  
    Change routing_instance with nat.

## Attributes Reference

The following attributes are exported:

- **id** (String)  
  An identifier for the resource with format `<name>`.

## Import

Junos security nat static can be imported using an id made up of `<name>`, e.g.

```shell
$ terraform import junos_security_nat_static.demo_nat nat_from_trust
```
