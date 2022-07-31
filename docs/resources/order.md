---
page_title: "order Resource - terraform-provider-houstn"
subcategory: ""
description: |-
The application resource allows you to configure a Houstn application.
---

# Resource `houstn_application`

The application resource allows you to configure a Houstn application.

## Example Usage

```terraform
resource "houstn_application" "my-application" {
  application_id = "my-application"
  name           = "My Application"
  active         = true
}
```

## Argument Reference

- `application_id` - (Required) A unique ID used for you application
- `name` - (Optional) A name for your application
- `active` - (Required) Whether your application is active

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

### Application

- `id` - The application id
- `application_id` - The specified application id
- `name` - The application name.
- `active` - The application active status.
