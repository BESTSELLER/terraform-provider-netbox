package provider

import (
	"bitbucket.org/bestsellerit/terraform-provider-harbor/client"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},
			"api_token": {
				Type:     schema.TypeString,
				Optional: true,
			}
		},

		ResourcesMap: map[string]*schema.Resource{
			"harbor_config_auth":    resourceConfigAuth(),
			"harbor_config_email":   resourceConfigEmail(),
			"harbor_config_system":  resourceConfigSystem(),
			"harbor_project":        resourceProject(),
			"harbor_project_member": resourceMembers(),
			"harbor_tasks":          resourceTasks(),
			"harbor_robot_account":  resourceRobotAccount(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	endpoint := d.Get("endpoint").(string)
	api_token := d.Get("api_token").(string)

	return client.NewClient(endpoint, token), nil
}