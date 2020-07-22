package provider

import (
	"strconv"

	"github.com/BESTSELLER/terraform-provider-netbox/client"
	"github.com/BESTSELLER/terraform-provider-netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

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
	apiClient := m.(*client.Client)

	body := client.AvailablePrefixBody(d)
	parentID := d.Get("parent_prefix_id").(int)

	resp, err := apiClient.CreatePrefix(&body, parentID)
	if err != nil {
		return err
	}

	id := models.PathAvailablePrefixes + strconv.Itoa(resp.ID) + "/"
	d.SetId(id)
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
	d.Set("parent_prefix_id", d.Get("parent_prefix_id").(int))

	return nil
}

func resourceAvailablePrefixUpdate(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	err := apiClient.UpdatePrefix(d)
	if err != nil {
		return err
	}

	return resourceAvailablePrefixRead(d, m)
}

func resourceAvailablePrefixDelete(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	err := apiClient.DeletePrefix(d)
	if err != nil {
		return err
	}

	return nil
}
