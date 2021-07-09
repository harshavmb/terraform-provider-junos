---
layout: "junos"
page_title: "Junos: junos_security_policy"
sidebar_current: "docs-junos-resource-security-policy"
description: |-
  Create a security policy (when Junos device supports it)
---

# junos_security_policy

Provides a security policy resource.

## Example Usage

```hcl
# Add a security policy
resource junos_security_policy "demo_policy" {
  from_zone = "trust"
  to_zone   = "untrust"
  policy {
    name                      = "allow_trust"
    match_source_address      = ["any"]
    match_destination_address = ["any"]
    match_application         = ["any"]
  }
}
```

## Argument Reference

The following arguments are supported:

* `from_zone` - (Required, Forces new resource)(`String`) The name of source zone.
* `to_zone` - (Required, Forces new resource)(`String`) The name of destination zone.
* `policy` - (Required)([attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html)) List of policy with options. Can be specified multiple times for each policy.
  * `name`  - (Required)(`String`) The name of policy.
  * `match_source_address` - (Required)(`ListOfString`) List of source address match.
  * `match_destination_address` - (Required)(`ListOfString`) List of destination address match.
  * `then` - (Optional)(`String`) Action of policy. Defaults to `permit`.
  * `count` - (Optional)(`Bool`) Enable count.
  * `log_init` - (Optional)(`Bool`) Log at session init time.
  * `log_close` - (Optional)(`Bool`) Log at session close time.
  * `match_application` - (Optional)(`ListOfString`) List of applications match.
  * `match_destination_address_excluded` - (Optional)(`Bool`) Exclude destination addresses.
  * `match_dynamic_application` - (Optional)(`ListOfString`) List of dynamic application or group match.
  * `match_source_address_excluded` - (Optional)(`Bool`) Exclude source addresses.
  * `match_source_end_user_profile` - (Optional)(`String`) Match source end user profile (device identity profile).
  * `permit_tunnel_ipsec_vpn` - (Optional)(`String`) Name of vpn to permit with a tunnel ipsec.
  * `permit_application_services` - (Optional)([attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html) Define application services for permit. See the [`permit_application_services` arguments for policy](#permit_application_services-arguments-for-policy) block. Max of 1.

---

### permit_application_services arguments for policy

* `advanced_anti_malware_policy` - (Optional)(`String`) Specify advanced-anti-malware policy name.
* `application_firewall_rule_set` - (Optional)(`String`) Service rule-set name for Application firewall.
* `application_traffic_control_rule_set` - (Optional)(`String`) Service rule-set name Application traffic control.
* `gprs_gtp_profile` - (Optional)(`String`) Specify GPRS Tunneling Protocol profile name.
* `gprs_sctp_profile` - (Optional)(`String`) Specify GPRS stream control protocol profile name.
* `idp` - (Optional)(`Bool`) Enable Intrusion detection and prevention.
* `idp_policy` - (Optional)(`String`) Specify idp policy name.
* `redirect_wx` - (Optional)(`Bool`) Set WX redirection.
* `reverse_redirect_wx` - (Optional)(`Bool`) Set WX reverse redirection.
* `security_intelligence_policy` - (Optional)(`String`) Specify security-intelligence policy name.
* `ssl_proxy` - (Optional)([attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html) Enable SSL Proxy. Max of 1.
  * `profile_name` - (Optional)(`String`) Specify SSL proxy service profile name.
* `uac_policy` - (Optional)([attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html) Enable unified access control enforcement. Max of 1.
  * `captive_portal` - (Optional)(`String`) Specify captive portal.
* `utm_policy` - (Optional)(`String`) Specify utm policy name.

## Import

Junos security policy can be imported using an id made up of `<from_zone>_-_<to_zone>`, e.g.

```shell
$ terraform import junos_security_zone.demo_policy trust_-_untrust
```
