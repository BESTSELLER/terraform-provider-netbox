package provider

import (
	"context"
	"strconv"

	"github.com/BESTSELLER/terraform-provider-netbox/client"
	"github.com/BESTSELLER/terraform-provider-netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSites() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSitesRead,
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

func dataSitesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	siteName := d.Get("name").(string)

	resp, err := apiClient.GetSite(siteName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("site_id", resp.Results[0].ID)
	d.Set("tenant_id", resp.Results[0].Tenant.ID)
	d.Set("tenant_name", resp.Results[0].Tenant.ID)

	id := models.PathAvailablePrefixes + strconv.Itoa(resp.Results[0].ID) + "/"
	d.SetId(id)

	return nil
}
