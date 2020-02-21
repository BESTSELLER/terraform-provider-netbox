package provider

import (
	"github.com/BESTSELLER/terraform-provider-netbox/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var pathAvailablePrefixes = "api/ipam/prefixes/"

type AvailablePrefixes struct {
	PrefixLenght int    `json:"project_lenght"`
	Site         int    `json:"site"`
	Tenant       int    `json:"tenant"`
	Status       string `json:"status"`
	Role         int    `json:"role"`
	Description  string `json:"description"`
}

func resourceAvailablePrefixes() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"prefix_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"prefix_length": {
				Type:     schema.TypeInt,
				Required: true,
				// ForceNew: true,
			},
			"site": {
				Type: schema.TypeInt,
				// Computed: true,
			},
			"tenant": {
				Type: schema.TypeInt,
				// Optional: true,
				// Default:  false,
			},
			"status": {
				Type: schema.TypeString,
				// Optional: true,
			},
			"role": {
				Type: schema.TypeString,
			},

			"description": {
				Type: schema.TypeString,
			},
		},
		Create: resourceAvailablePrefixCreate,
		Read:   resourceAvailablePrefixRead,
		Update: resourceAvailablePrefixUpdate,
		Delete: resourceAvailablePrefixDelete,
	}
}

func resourceAvailablePrefixCreate(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	id := strcon.Itoa(d.Get("prefix_id").(int))

	path := pathAvailablePrefixes + id + "/available-prefixes/"
	body := AvailablePrefixes{
		PrefixLenght: d.Get("prefix_length").(int),
		Site:         d.Get("site").(int),
		Tenant:       d.Get("tenant").(int),
		Status:       d.Get("status").(string),
		Role:         d.Get("role").(string),
		Description:  d.Get("desscription").(string),
	}

	_, err := apiClient.SendRequest("POST", path, body, 201)
	if err != nil {
		return err
	}

	// d.SetId(randomString(15))
	return resourceAvailablePrefixRead(d, m)
}

func resourceAvailablePrefixRead(d *schema.ResourceData, m interface{}) error {
	// apiClient := m.(*client.Client)
	return nil
}

func resourceAvailablePrefixUpdate(d *schema.ResourceData, m interface{}) error {
	// apiClient := m.(*client.Client)

	return resourceAvailablePrefixRead(d, m)
}

func resourceAvailablePrefixDelete(d *schema.ResourceData, m interface{}) error {
}
