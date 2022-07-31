---
page_title: "environment Resource - terraform-provider-houstn"
subcategory: ""
description: |-
The environment resource allows you to configure a Houstn environment.
---

# Resource `houstn_environment`

The environment resource allows you to configure a Houstn environment.

## Example Usage

```terraform
resource "houstn_environment" "my-environment" {
  environment_id = "my-environment"
  name           = "My Environment"
  active         = true
}
```

## Argument Reference

- `environment_id` - (String, Required) A unique ID used for you environment
- `name` - (String, Optional) A name for your environment
- `group` - (String, Optional) A group key to allow the UI to group with other environments
- `order` - (Integer, Optional) An order number to specify the order in which each environment appears

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

### Environment

- `id` - The environment id
- `environment_id` - The specified environment id
- `name` - The environment name
- `group` - The environment group
- `order` - The environment order
