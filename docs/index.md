---
page_title: "Provider: Houstn"
subcategory: ""
description: |-
Terraform provider for interacting with Houstn API.
---

# Houstn Provider

The Houstn provider is used to interact with https://houstn.io.

Use the navigation to the left to read about the available resources.

## Example Usage

Do not keep your authentication password in HCL for production environments, use Terraform environment variables.

```terraform
provider "houstn" {
  organisation = "your-org"
  token        = "your-token"
}
```

## Schema

### Optional

- **organisation** (String, Required) Organisation ID used when signing up to houstn
- **token** (String, Optional) Token to authenticate to Houstn API
- **host** (String, Optional) Houstn API address (defaults to `houstn.io`)