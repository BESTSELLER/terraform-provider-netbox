package provider

import (
	"fmt"
	"strconv"

	"github.com/BESTSELLER/terraform-provider-netbox/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataDeviceType() *schema.Resource {
	return &schema.Resource{
		Read: dataDeviceTypeRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "A unique integer value identifying this device type.",
			},
			"displayname": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The display name of the device type.",
			},
			"manufacturer": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "The manufacturer object of the device type.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "A unique integer value identifying this manufacturer.",
						},
						"displayname": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The display name of the manufacturer.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the manufacturer name.",
						},
						"slug": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The slug of the manufacturer.",
						},
					},
				},
			},
			"model": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The model of the device type.",
			},
			"slug": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The slug of the device type.",
			},
			"part_number": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The part number of the device type.",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The description of the device type.",
			},
			"custom_fields": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "The custom fields of the device type.",
				Default:     nil,
				Elem: &schema.Schema{
					Type:    schema.TypeString,
					Default: nil,
				},
			},
		},
	}
}

func dataDeviceTypeRead(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)

	id := d.Get("id").(int)
	if id < 0 {
		return fmt.Errorf("[ERROR] 'id' cannot be less than 0")
	}

	resp, err := apiClient.GetDeviceType(id)
	if err != nil {
		return err
	}

	cf := getCustomFields(resp.CustomFields)
	if cf != nil {
		d.Set("custom_fields", cf)
	}

	d.Set("displayname", resp.DisplayName)
	d.Set("manufacturer", resp.Manufacturer)
	d.Set("model", resp.Model)
	d.Set("slug", resp.Slug)
	d.Set("part_number", resp.PartNumber)
	d.Set("description", resp.Descrption)
	d.SetId(strconv.Itoa(resp.ID))
	return nil
}

func getCustomFields(cf interface{}) map[string]interface{} {
	cfm, ok := cf.(map[string]interface{})
	if !ok || len(cfm) == 0 {
		return nil
	}
	return cfm
}
