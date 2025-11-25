package provider

import (
	"context"
	"strconv"

	"github.com/BESTSELLER/terraform-provider-netbox/client"
	"github.com/BESTSELLER/terraform-provider-netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
				Type:     schema.TypeString,
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
		CreateContext: resourceAvailablePrefixCreate,
		ReadContext:   resourceAvailablePrefixRead,
		UpdateContext: resourceAvailablePrefixUpdate,
		DeleteContext: resourceAvailablePrefixDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceAvailablePrefixCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	body := client.AvailablePrefixBody(d)
	parentID := d.Get("parent_prefix_id").(int)

	resp, err := apiClient.CreatePrefix(&body, parentID)
	if err != nil {
		return diag.FromErr(err)
	}

	id := models.PathAvailablePrefixes + strconv.Itoa(resp.ID) + "/"
	d.SetId(id)
	return resourceAvailablePrefixRead(ctx, d, m)
}

func resourceAvailablePrefixRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	resp, err := apiClient.GetAvailablePrefix(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("cidr_notation", resp.Prefix)
	d.Set("description", resp.Description)
	d.Set("prefix_length", resp.PrefixLength)
	d.Set("prefix_id", resp.ID)
	d.Set("parent_prefix_id", d.Get("parent_prefix_id").(int))

	return nil
}

func resourceAvailablePrefixUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	err := apiClient.UpdatePrefix(d)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceAvailablePrefixRead(ctx, d, m)
}

func resourceAvailablePrefixDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)
	err := apiClient.DeletePrefix(d)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
