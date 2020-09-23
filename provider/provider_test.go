package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/BESTSELLER/terraform-provider-netbox/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"netbox": testAccProvider,
	}
}

func testProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {

	if v := os.Getenv("NETBOX_ENDPOINT"); v == "" {
		t.Fatal("NETBOX_ENDPOINT must be set for acceptance tests")
	}
	if v := os.Getenv("NETBOX_API_TOKEN"); v == "" {
		t.Fatal("NETNOX_API_TOKEN must be set for acceptance tests")
	}

}

func testAccCheckResourceExists(resource string) resource.TestCheckFunc {

	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("Not found: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}
		name := rs.Primary.ID

		apiClient := testAccProvider.Meta().(*client.Client)
		_, _, err := apiClient.SendRequest("GET", name, nil, 200)
		if err != nil {
			return fmt.Errorf("error fetching item with resource %s. %s", resource, err)
		}
		return nil
	}
}
