package provider

import (
	"context"

	"github.com/BESTSELLER/terraform-provider-netbox/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider returns a *schema.Provider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETBOX_ENDPOINT", ""),
			},
			"api_token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETBOX_API_TOKEN", ""),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"netbox_available_prefix": resourceAvailablePrefixes(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"netbox_prefix":      dataPrefix(),
			"netbox_sites":       dataSites(),
			"netbox_device_type": dataDeviceType(),
		},

		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	endpoint := d.Get("endpoint").(string)
	apiToken := d.Get("api_token").(string)

	return client.NewClient(endpoint, apiToken), nil
}
