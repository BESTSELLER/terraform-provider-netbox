# Terraform Provider Netbox

This is a Terraform provider for [Netbox](https://github.com/netbox-community/netbox).

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.12.x+
- [Go](https://golang.org/doc/install) 1.25.x (to build the provider plugin)

## Building The Provider

```sh
git clone git@github.com:BESTSELLER/terraform-provider-netbox.git
cd terraform-provider-netbox
go build -o terraform-provider-netbox
```

## Installation

To install the provider, you need to copy the binary to your Terraform plugins directory.

**Linux/macOS:**

```sh
mkdir -p ~/.terraform.d/plugins
mv terraform-provider-netbox ~/.terraform.d/plugins/
```

**Windows:**

```powershell
mkdir %APPDATA%\terraform.d\plugins
move terraform-provider-netbox.exe %APPDATA%\terraform.d\plugins\
```

## Usage

### Provider Configuration

```hcl
provider "netbox" {
  endpoint  = "https://netbox.example.com"
  api_token = "your-api-token"
}
```

You can also set the configuration via environment variables:

- `NETBOX_ENDPOINT`
- `NETBOX_API_TOKEN`

### Resources

#### netbox_available_prefix

Allocates the next available prefix within a parent prefix.

```hcl
resource "netbox_available_prefix" "example" {
  parent_prefix_id = 123
  prefix_length    = 24

  # Optional
  site        = 1
  tenant      = 2
  status      = "active"
  role        = 3
  description = "My new prefix"
}
```

**Arguments:**

- `parent_prefix_id` (Required, Int) The ID of the parent prefix to allocate from.
- `prefix_length` (Required, Int) The length of the prefix to allocate.
- `site` (Optional, Int) The ID of the site.
- `tenant` (Optional, Int) The ID of the tenant.
- `status` (Optional, String) The status of the prefix (e.g., "active", "reserved").
- `role` (Optional, Int) The ID of the role.
- `description` (Optional, String) A description for the prefix.

**Attributes:**

- `prefix_id` (Int) The ID of the allocated prefix.
- `cidr_notation` (String) The allocated prefix in CIDR notation.

### Data Sources

#### netbox_prefix

Retrieves information about a specific prefix.

```hcl
data "netbox_prefix" "example" {
  cidr_notation = "10.0.0.0/24"
}
```

**Arguments:**

- `cidr_notation` (Required, String) The CIDR notation of the prefix.

**Attributes:**

- `prefix_id` (Int) The ID of the prefix.
- `role_id` (Int) The ID of the role.
- `role_name` (String) The name of the role.
- `site_id` (Int) The ID of the site.
- `site_name` (String) The name of the site.
- `tenant_id` (Int) The ID of the tenant.
- `tenant_name` (String) The name of the tenant.

#### netbox_sites

Retrieves information about a specific site.

```hcl
data "netbox_sites" "example" {
  name = "MySite"
}
```

**Arguments:**

- `name` (Required, String) The name of the site.

**Attributes:**

- `site_id` (Int) The ID of the site.
- `tenant_id` (Int) The ID of the tenant.
- `tenant_name` (String) The name of the tenant.

#### netbox_device_type

Retrieves information about a specific device type.

```hcl
data "netbox_device_type" "example" {
  id = 123
}
```

**Arguments:**

- `id` (Required, Int) The ID of the device type.

**Attributes:**

- `displayname` (String) The display name of the device type.
- `manufacturer` (Map) The manufacturer object.
- `model` (String) The model of the device type.
- `slug` (String) The slug of the device type.
- `part_number` (String) The part number.
- `description` (String) The description.
- `custom_fields` (Map) Custom fields associated with the device type.

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.25+ is recommended).

To compile the provider, run `go build`. This will build the provider and put the provider binary in the current directory.

```sh
go build
```

## Debugging

### VS Code

You can debug the acceptance tests directly in VS Code.

1.  Create a `.env` file in the root of the project with your Netbox configuration:
    ```env
    NETBOX_ENDPOINT=https://netbox.example.com
    NETBOX_API_TOKEN=your-api-token
    ```
2.  Open the "Run and Debug" view in VS Code.
3.  Select "Debug TestAccRegistryBasic" from the dropdown.
4.  Press F5 or click the green play button.

This configuration runs the `TestAccRegistryBasic` test with `TF_ACC=true` enabled.

### Terminal

To run acceptance tests from the terminal:

```sh
export NETBOX_ENDPOINT="https://netbox.example.com"
export NETBOX_API_TOKEN="your-api-token"
export TF_ACC=true
go test -v ./...
```

### Troubleshooting

- **TF_LOG**: Set `TF_LOG=DEBUG` or `TF_LOG=TRACE` to see detailed Terraform logs.
- **Test Data**: The acceptance tests may expect specific data to exist in your Netbox instance (e.g., specific prefixes or sites). Check the test files (e.g., `provider/resource_availableprefixes_test.go`) to see what data is required.
