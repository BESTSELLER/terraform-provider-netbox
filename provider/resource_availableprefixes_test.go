package provider

import (
	"fmt"
	"testing"

	"github.com/BESTSELLER/terraform-provider-netbox/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
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
						"netbox_available_prefix.main", "description", "bw-testing-terraform"),
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

		resp, _, err := apiClient.SendRequest("GET", rs.Primary.ID, nil, 404)
		if err != nil {
			return fmt.Errorf("Resouse was not delete \n %s", resp)
		}

	}

	return nil
}

func testAccCheckAvailablePrefix() string {

	return fmt.Sprintf(`

	resource "netbox_available_prefix" "main" {
		parent_prefix_id = "758"
		prefix_length = 26
		description = "acc-testing-terraform"
	}

	`)
}
