package provider

import (
	"encoding/json"
	"time"

	"github.com/BESTSELLER/terraform-provider-netbox/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type responseSites struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Slug   string `json:"slug"`
		Status struct {
			Value string `json:"value"`
			Label string `json:"label"`
			ID    int    `json:"id"`
		} `json:"status"`
		Region interface{} `json:"region"`
		Tenant struct {
			ID   int    `json:"id"`
			URL  string `json:"url"`
			Name string `json:"name"`
			Slug string `json:"slug"`
		} `json:"tenant"`
		Facility        string        `json:"facility"`
		Asn             interface{}   `json:"asn"`
		TimeZone        interface{}   `json:"time_zone"`
		Description     string        `json:"description"`
		PhysicalAddress string        `json:"physical_address"`
		ShippingAddress string        `json:"shipping_address"`
		Latitude        interface{}   `json:"latitude"`
		Longitude       interface{}   `json:"longitude"`
		ContactName     string        `json:"contact_name"`
		ContactPhone    string        `json:"contact_phone"`
		ContactEmail    string        `json:"contact_email"`
		Comments        string        `json:"comments"`
		Tags            []interface{} `json:"tags"`
		CustomFields    struct {
		} `json:"custom_fields"`
		Created             string      `json:"created"`
		LastUpdated         time.Time   `json:"last_updated"`
		CircuitCount        interface{} `json:"circuit_count"`
		DeviceCount         interface{} `json:"device_count"`
		PrefixCount         int         `json:"prefix_count"`
		RackCount           interface{} `json:"rack_count"`
		VirtualmachineCount interface{} `json:"virtualmachine_count"`
		VlanCount           interface{} `json:"vlan_count"`
	} `json:"results"`
}

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
	sitesPath := "api/dcim/sites/?name=" + siteName

	resp, err := apiClient.SendRequest("GET", sitesPath, nil, 200)
	if err != nil {
		return err
	}

	var jsonData responseSites
	json.Unmarshal([]byte(resp), &jsonData)

	d.Set("site_id", jsonData.Results[0].ID)
	d.Set("tenant_id", jsonData.Results[0].Tenant.ID)
	d.Set("tenant_name", jsonData.Results[0].Tenant.ID)
	d.SetId(randomString(15))

	return nil
}
