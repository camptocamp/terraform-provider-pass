package pass

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"store_dir": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PASSWORD_STORE_DIR", ""),
				Description: "Password storage directory to use.",
			},
			"refresh_store": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether or not call `pass git pull`.",
			},
		},

		ConfigureFunc: providerConfigure,

		DataSourcesMap: map[string]*schema.Resource{
			"pass_password": passwordDataSource(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"pass_password": passPasswordResource(),
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	os.Setenv("PASSWORD_STORE_DIR", d.Get("store_dir").(string))

	if d.Get("refresh_store").(bool) {
		log.Printf("[DEBUG] Pull pass repository")
		output, err := exec.Command("pass", "git", "pull").Output()
		if err != nil {
			return nil, fmt.Errorf("error refreshing password store: %s", err)
		}
		log.Printf("[DEBUG] output: %s", string(output))
	}

	return nil, nil
}
