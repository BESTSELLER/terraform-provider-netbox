package provider

import (
	"github.com/BESTSELLER/terraform-provider-netbox/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},
			"api_token": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"netbox_available_prefix": resourceAvailablePrefixes(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"netbox_prefix": dataPrefix(),
			"netbox_sites":  dataSites(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	endpoint := d.Get("endpoint").(string)
	apiToken := d.Get("api_token").(string)

	return client.NewClient(endpoint, apiToken), nil
}
