---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cloudbeaver_team Data Source - cloudbeaver"
subcategory: ""
description: |-
  Team data source
---

# cloudbeaver_team (Data Source)

Team data source

## Example Usage

```terraform
data "cloudbeaver_team" "example" {
  id = "some-team"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Team identifier

### Read-Only

- `description` (String) Team description
- `name` (String) Team name