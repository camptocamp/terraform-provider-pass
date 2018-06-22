package pass

import (
	"context"
	"log"
	"os"

	"github.com/blang/semver"
	"github.com/gopasspw/gopass/pkg/action"
	"github.com/gopasspw/gopass/pkg/config"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/pkg/errors"
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

	ctx := context.Background()

	act, err := action.New(ctx, config.Load(), semver.Version{})
	if err != nil {
		return nil, errors.Wrap(err, "error instantiating password store")
	}

	if !act.Store.Initialized(ctx) {
		return nil, errors.Wrap(err, "password-store not initialized")
	}
	st := act.Store

	if d.Get("refresh_store").(bool) {
		log.Printf("[DEBUG] Pull pass repository")
		err := st.GitPull(ctx, "", "origin", "master")

		if err != nil {
			return nil, errors.Wrap(err, "error refreshing password store")
		}
	}

	return st, nil
}
