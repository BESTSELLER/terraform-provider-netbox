# <resource name> Resource/Data Source

Description of what this resource does, with links to official
app/service documentation.

## Example Usage

```hcl
data "netbox_prefix" "main" {
  cidr_notation = "172.24.0.0/13"
}
```

## Argument Reference

* `cidr_notation` - (Required) The CIDR prefix you wanna use.

## Attribute Reference

* `prefix_id` - ID of the prefix.
* `role_id` - ID of the role the prefix has.
* `role_name` - The name of that role.
* `site_id` - ID of the site.
* `site_name` - Name of that site.
* `tenant_id` - ID of the tenant
* `tenant_name` - Name of that tenant
