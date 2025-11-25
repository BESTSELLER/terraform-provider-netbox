package provider

import (
	"fmt"
	"testing"

	"github.com/BESTSELLER/terraform-provider-netbox/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccRegistryBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRegistryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAvailablePrefix(),
				Check: resource.ComposeTestCheckFunc(

					testAccCheckResourceExists("netbox_available_prefix.main"),
					resource.TestCheckResourceAttr(
						"netbox_available_prefix.main", "description", "acc-testing-terraform"),
				),
			},
		},
	})
}

func testAccCheckRegistryDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_available_prefix" {
			continue
		}

		resp, err := apiClient.SendRequest("GET", rs.Primary.ID, nil, 404)
		if err != nil {
			return fmt.Errorf("Resouse was not delete \n %s", resp)
		}

	}

	return nil
}

func testAccCheckAvailablePrefix() string {

	return (`

	data "netbox_prefix" "prefix" {
		cidr_notation = "172.24.0.0/13"
	}

	resource "netbox_available_prefix" "main" {
		parent_prefix_id = data.netbox_prefix.prefix.prefix_id
		prefix_length = 19
		description = "acc-testing-terraform"
		status        = "active"
		site          = data.netbox_prefix.prefix.site_id
		tenant        = data.netbox_prefix.prefix.tenant_id
		role          = data.netbox_prefix.prefix.role_id
	}
	resource "netbox_available_prefix" "main2" {
		parent_prefix_id = data.netbox_prefix.prefix.prefix_id
		prefix_length = 20
		description = "acc-testing-terraform"
		status        = "active"
		site          = data.netbox_prefix.prefix.site_id
		tenant        = data.netbox_prefix.prefix.tenant_id
		role          = data.netbox_prefix.prefix.role_id
	}
	resource "netbox_available_prefix" "main3" {
		parent_prefix_id = data.netbox_prefix.prefix.prefix_id
		prefix_length = 28
		description = "acc-testing-terraform"
		status        = "active"
		site          = data.netbox_prefix.prefix.site_id
		tenant        = data.netbox_prefix.prefix.tenant_id
		role          = data.netbox_prefix.prefix.role_id
	}

	`)
}
