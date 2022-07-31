---
page_title: "deployment Resource - terraform-provider-houstn"
subcategory: ""
description: |-
The deployment resource allows you to configure a Houstn deployment.
---

# Resource `houstn_deployment`

The deployment resource allows you to configure a Houstn deployment.

## Example Usage

```terraform
resource "houstn_deployment" "my-deployment" {
  environment = "my-environment"
  application = "my-application"
  url         = "https://example.com"
}
```

## Argument Reference

- `environment` - (String, Required) An environment ID
- `application` - (String, Required) An application ID
- `url` - (String, Required) The URL for this deployment

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

### Deployment

- `id` - The deployment id
- `environment` - The deployment environment
- `application` - The deployment application
- `url` - The deployment url
