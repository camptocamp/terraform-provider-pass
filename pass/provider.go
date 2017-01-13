package pass

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"store_dir": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PASSWORD_STORE_DIR", nil),
				Description: "Password storage directory to use.",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"pass_password": passwordDataSource(),
		},
	}
}
