package pass

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/blang/semver"
	"github.com/gopasspw/gopass/pkg/action"
	"github.com/gopasspw/gopass/pkg/backend/crypto/gpg"
	_ "github.com/gopasspw/gopass/pkg/backend/storage"
	"github.com/gopasspw/gopass/pkg/config"
	"github.com/gopasspw/gopass/pkg/store/root"
	"github.com/gopasspw/gopass/pkg/store/sub"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/pkg/errors"
)

type passProvider struct {
	st  *root.Store
	ctx context.Context
}

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"store_dir": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PASSWORD_STORE_DIR", ""),
				Description: "Password storage directory to use.",
			},
			"refresh_store": {
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
	ctx = gpg.WithAlwaysTrust(ctx, true)
	ctx = sub.WithCheckRecipients(ctx, false)

	act, err := action.New(ctx, config.Load(), semver.Version{})
	if err != nil {
		return nil, errors.Wrap(err, "error instantiating password store")
	}

	if ok, err := act.Store.Initialized(ctx); !ok || err != nil {
		return nil, errors.New(fmt.Sprintf("password-store not initialized: %s", err))
	}
	st := act.Store

	if d.Get("refresh_store").(bool) {
		log.Printf("[DEBUG] Pull pass repository")
		err := st.GitPull(ctx, "", "origin", "master")

		if err != nil {
			return nil, errors.Wrap(err, "error refreshing password store")
		}
	}

	return &passProvider{st, ctx}, nil
}
