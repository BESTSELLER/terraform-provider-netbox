package provider

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/BESTSELLER/terraform-provider-netbox/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type reponsePrefix struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		ID     int `json:"id"`
		Family struct {
			Value int    `json:"value"`
			Label string `json:"label"`
		} `json:"family"`
		Prefix string `json:"prefix"`
		Site   struct {
			ID   int    `json:"id"`
			URL  string `json:"url"`
			Name string `json:"name"`
			Slug string `json:"slug"`
		} `json:"site"`
		Vrf    interface{} `json:"vrf"`
		Tenant struct {
			ID   int    `json:"id"`
			URL  string `json:"url"`
			Name string `json:"name"`
			Slug string `json:"slug"`
		} `json:"tenant"`
		Vlan   interface{} `json:"vlan"`
		Status struct {
			Value string `json:"value"`
			Label string `json:"label"`
			ID    int    `json:"id"`
		} `json:"status"`
		Role struct {
			ID   int    `json:"id"`
			URL  string `json:"url"`
			Name string `json:"name"`
			Slug string `json:"slug"`
		} `json:"role"`
		IsPool       bool          `json:"is_pool"`
		Description  string        `json:"description"`
		Tags         []interface{} `json:"tags"`
		CustomFields struct {
		} `json:"custom_fields"`
		Created     string    `json:"created"`
		LastUpdated time.Time `json:"last_updated"`
	} `json:"results"`
}

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
	prefixPath := "/ipam/prefixes/?prefix=" + newCidr

	resp, err := apiClient.SendRequest("GET", prefixPath, nil, 200)
	if err != nil {
		return err
	}

	var jsonData reponsePrefix
	json.Unmarshal([]byte(resp), &jsonData)

	d.Set("prefix_id", jsonData.Results[0].ID)
	d.Set("role_id", jsonData.Results[0].Role.ID)
	d.Set("role_name", jsonData.Results[0].Role.Name)
	d.Set("site_id", jsonData.Results[0].Site.ID)
	d.Set("site_name", jsonData.Results[0].Site.Name)
	d.Set("tenant_id", jsonData.Results[0].Tenant.ID)
	d.Set("tenant_name", jsonData.Results[0].Tenant.ID)

	d.SetId(randomString(15))

	return nil
}
