package provider

import (
	"strconv"

	"github.com/BESTSELLER/terraform-provider-netbox/client"
	"github.com/BESTSELLER/terraform-provider-netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSites() *schema.Resource {
	return &schema.Resource{
		Read: dataSitesRead,
		Schema: map[string]*schema.Schema{
			"site_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"tenant_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSitesRead(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	siteName := d.Get("name").(string)

	resp, err := apiClient.GetSite(siteName)
	if err != nil {
		return err
	}

	d.Set("site_id", resp.Results[0].ID)
	d.Set("tenant_id", resp.Results[0].Tenant.ID)
	d.Set("tenant_name", resp.Results[0].Tenant.ID)

	id := models.PathAvailablePrefixes + strconv.Itoa(resp.Results[0].ID) + "/"
	d.SetId(id)

	return nil
}
