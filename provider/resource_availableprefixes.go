package provider

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/BESTSELLER/terraform-provider-netbox/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var pathAvailablePrefixes = "/ipam/prefixes/"

type AvailablePrefixes struct {
	PrefixLenght int    `json:"prefix_length"`
	Site         int    `json:"site,omitempty"`
	Tenant       int    `json:"tenant,omitempty"`
	Status       string `json:"status,omitempty"`
	Role         int    `json:"role,omitempty"`
	Description  string `json:"description,omitempty"`
}

type responeListOfPrefixes struct {
	Count    int                        `json:"count"`
	Next     interface{}                `json:"next"`
	Previous interface{}                `json:"previous"`
	Results  []reponseAvailablePrefixes `json:"results"`
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
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
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

	d.SetId(pathAvailablePrefixes + strconv.Itoa(jsonData.ID) + "/")
	return resourceAvailablePrefixRead(d, m)
}

func resourceAvailablePrefixRead(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	resp, err := apiClient.GetAvailablePrefix(d.Id())
	if err != nil {
		return err
	}

	d.Set("cidr_notation", resp.Prefix)
	d.Set("description", resp.Description)
	d.Set("prefix_length", resp.PrefixLength)
	d.Set("prefix_id", resp.ID)
	d.Set("parent_prefix_id", resp.ParentPrefixID)
	d.SetId(resp.ID)

	return nil
}

func resourceAvailablePrefixUpdate(d *schema.ResourceData, m interface{}) error {
	apiClient, body := availablePrefixBody(d, m)
	path := d.Id()

	apiClient.SendRequest("PATCH", path, body, 200)

	return resourceAvailablePrefixRead(d, m)
}

func resourceAvailablePrefixDelete(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	apiClient.SendRequest("DELETE", d.Id(), nil, 204)
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
