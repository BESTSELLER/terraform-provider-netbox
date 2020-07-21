# <resource name> Resource/Data Source

Description of what this resource does, with links to official
app/service documentation.

## Example Usage

```hcl
data "netbox_prefix" "main" {
  cidr_notation = "172.24.0.0/13"
}
resource "netbox_available_prefix" "main" {
  parent_prefix_id = data.netbox_prefix.main.prefix_id

  prefix_length = 19
  status        = "active"
  site          = data.netbox_prefix.main.site_id
  tenant        = data.netbox_prefix.main.tenant_id
  role          = data.netbox_prefix.main.role_id

  description = "Some description here"
}
```

## Argument Reference

* `parent_prefix_id` - (Required) ID of the parent prefix.
* `prefix_length` - (Required) Length of the prefix you wanna create.
* `site` - (Optional) ID of the site.
* `tenant` - (Optional) ID of the tenant.
* `status` - (Optional) Status of the prefix.
* `role` - (Optional) ID of the role.
* `description` - (Optional) Description of the prefix

## Attribute Reference

* `prefix_id` - ID of the prefix.
* `cidr_notation` - List attributes that this resource exports.
