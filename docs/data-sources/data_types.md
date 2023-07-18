# Data Source - netbox_device_type

Get a specific device type.

## Example Usage

```hcl
data "netbox_device_type" "main" {
  id = 1
}
```

## Argument Reference

* `id` - (Required) The ID of the device type.

## Attribute Reference

The following attributes are exported:

* `displayname` - Display name of the device type.
* `manufacturer` - The manufacturer object of the device type.
  * `id` - A unique integer value identifying this manufacturer.
  * `displayname` - The display name of the manufacturer.
  * `name` - The name of the manufacturer name.
  * `slug` - The slug of the manufacturer.
* `model` - The model of the device type.
* `slug` - The slug of the device type.
* `part_number` - The part number of the device type.
* `description` - The description of the device type.
* `custom_fields` - The custom fields of the device type.