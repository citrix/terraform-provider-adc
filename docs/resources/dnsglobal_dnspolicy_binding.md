---
subcategory: "DNS"
---

# Resource: dnsglobal_dnspolicy_binding

The dnsglobal_dnspolicy_binding resource is used to create DNS global transform policy binding.


## Example usage

```hcl
resource "citrixadc_dnspolicy" "dnspolicy" {
  name = "policy_A"
  rule = "CLIENT.IP.SRC.IN_SUBNET(1.1.1.1/24)"
  drop = "YES"
}
resource "citrixadc_dnsglobal_dnspolicy_binding" "dnsglobal_dnspolicy_binding" {
  policyname = citrixadc_dnspolicy.dnspolicy.name
  priority   = 30
  type       = "REQ_DEFAULT"
}
```


## Argument Reference

* `policyname` - (Required) Name of the dns policy.
* `priority` - (Required) Specifies the priority of the policy with which it is bound. Maximum allowed priority should be less than 65535
* `globalbindtype` - (Optional) 0
* `gotopriorityexpression` - (Optional) Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.  Specify one of the following values: * NEXT - Evaluate the policy with the next higher priority number. * END - End policy evaluation. * USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT. * An expression that evaluates to a number. If you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows: * If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next. * If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next. * If the expression evaluates to a priority number that is numerically higher than the highest numbered priority, policy evaluation ends. An UNDEF event is triggered if: * The expression is invalid. * The expression evaluates to a priority number that is numerically lower than the current policy's priority. * The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.
* `invoke` - (Optional) Invoke flag.
* `labelname` - (Optional) Name of the label to invoke if the current policy rule evaluates to TRUE.
* `labeltype` - (Optional) Type of policy label invocation.
* `type` - (Optional) Type of global bind point for which to show bound policies.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnsglobal_dnspolicy_binding. It has the same value as the `policyname` attribute.


## Import

A dnsglobal_dnspolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_dnsglobal_dnspolicy_binding.dnsglobal_dnspolicy_binding policy_A
```
