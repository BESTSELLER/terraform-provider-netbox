# Netbox Provider

Summary of what the provider is for, including use cases and links to
app/service documentation.

## Example Usage

```hcl
provider "netbox" {
  endpoint  = data.vault_generic_secret.netbox_auth.data["endpoint"]
  api_token = data.vault_generic_secret.netbox_auth.data["token"]
}
```

## Argument Reference

* `endpoint` - (Required) Base url of your netbox. Fx. https://netbox.example.com/ - Trailing "/" is important!
* `api_token` - (Optional) API token used to communicate with netbox.
