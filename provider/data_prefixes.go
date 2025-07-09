package provider

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/BESTSELLER/terraform-provider-netbox/client"
	"github.com/BESTSELLER/terraform-provider-netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataPrefix() *schema.Resource {
	return &schema.Resource{
		Read: dataPrefixRead,
		Schema: map[string]*schema.Schema{
			"cidr_notation": {
				Type:     schema.TypeString,
				Required: true,
			},
			"prefix_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"role_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"role_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"site_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"site_name": {
				Type:     schema.TypeString,
				Computed: true,
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

func dataPrefixRead(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	cidr := d.Get("cidr_notation").(string)
	newCidr := strings.Replace(cidr, "/", "%2F", 1)
	if newCidr == "" {
		return fmt.Errorf("[ERROR] 'cidr_notation' is empty")
	}

	resp, err := apiClient.GetPrefix(newCidr)
	if err != nil {
		return err
	}

	d.Set("prefix_id", resp.Results[0].ID)
	d.Set("role_id", resp.Results[0].Role.ID)
	d.Set("role_name", resp.Results[0].Role.Name)
	d.Set("site_id", resp.Results[0].Scope.ID)
	d.Set("site_name", resp.Results[0].Scope.Name)
	d.Set("tenant_id", resp.Results[0].Tenant.ID)
	d.Set("tenant_name", resp.Results[0].Tenant.ID)

	id := models.PathAvailablePrefixes + strconv.Itoa(resp.Results[0].ID) + "/"
	d.SetId(id)

	return nil
}
