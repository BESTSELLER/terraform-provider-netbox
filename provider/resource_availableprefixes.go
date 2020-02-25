package provider

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/BESTSELLER/terraform-provider-netbox/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var pathAvailablePrefixes = "api/ipam/prefixes/"

type AvailablePrefixes struct {
	PrefixLenght int    `json:"prefix_length"`
	Site         int    `json:"site,omitempty"`
	Tenant       int    `json:"tenant,omitempty"`
	Status       string `json:"status,omitempty"`
	Role         int    `json:"role,omitempty"`
	Description  string `json:"description,omitempty"`
}

type reponseAvailablePrefixes struct {
	ID     int `json:"id"`
	Family struct {
		Value int    `json:"value"`
		Label string `json:"label"`
	} `json:"family"`
	Prefix string      `json:"prefix"`
	Site   interface{} `json:"site"`
	Vrf    interface{} `json:"vrf"`
	Tenant interface{} `json:"tenant"`
	Vlan   interface{} `json:"vlan"`
	Status struct {
		Value string `json:"value"`
		Label string `json:"label"`
		ID    int    `json:"id"`
	} `json:"status"`
	Role        interface{}   `json:"role"`
	IsPool      bool          `json:"is_pool"`
	Description string        `json:"description"`
	Tags        []interface{} `json:"tags"`
	Created     string        `json:"created"`
	LastUpdated time.Time     `json:"last_updated"`
}

func resourceAvailablePrefixes() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"parent_prefix_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"prefix_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"prefix_length": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"site": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tenant": {
				Type:     schema.TypeInt,
				Optional: true,
				// Default:  false,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"role": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cidr_notation": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		Create: resourceAvailablePrefixCreate,
		Read:   resourceAvailablePrefixRead,
		Update: resourceAvailablePrefixUpdate,
		Delete: resourceAvailablePrefixDelete,
	}
}

func resourceAvailablePrefixCreate(d *schema.ResourceData, m interface{}) error {
	apiClient, body := availablePrefixBody(d, m)
	id := strconv.Itoa(d.Get("parent_prefix_id").(int))

	path := pathAvailablePrefixes + id + "/available-prefixes/"

	resp, err := apiClient.SendRequest("POST", path, body, 201)
	if err != nil {
		return err
	}
	var jsonData reponseAvailablePrefixes
	json.Unmarshal([]byte(resp), &jsonData)

	d.Set("prefix_id", jsonData.ID)
	d.Set("cidr_notation", jsonData.Prefix)
	d.SetId(randomString(15))
	return resourceAvailablePrefixRead(d, m)
}

func resourceAvailablePrefixRead(d *schema.ResourceData, m interface{}) error {
	// apiClient := m.(*client.Client)
	return nil
}

func resourceAvailablePrefixUpdate(d *schema.ResourceData, m interface{}) error {
	apiClient, body := availablePrefixBody(d, m)
	id := strconv.Itoa(d.Get("prefix_id").(int))

	path := pathAvailablePrefixes + id + "/"

	apiClient.SendRequest("PATCH", path, body, 200)

	return resourceAvailablePrefixRead(d, m)
}

func resourceAvailablePrefixDelete(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	id := strconv.Itoa(d.Get("prefix_id").(int))
	prefixPath := pathAvailablePrefixes + id + "/"
	apiClient.SendRequest("DELETE", prefixPath, nil, 204)
	return nil
}

func availablePrefixBody(d *schema.ResourceData, m interface{}) (*client.Client, AvailablePrefixes) {
	apiClient := m.(*client.Client)

	body := AvailablePrefixes{
		PrefixLenght: d.Get("prefix_length").(int),
		Site:         d.Get("site").(int),
		Tenant:       d.Get("tenant").(int),
		Status:       d.Get("status").(string),
		Role:         d.Get("role").(int),
		Description:  d.Get("description").(string),
	}
	return apiClient, body
}
