package pass

import (
	"context"
	"sync"

	"github.com/gopasspw/gopass/pkg/gopass/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

type passProvider struct {
	store *api.Gopass
	mutex *sync.Mutex
}

func Provider() *schema.Provider {
	return &schema.Provider{
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
	ctx := context.Background()

	store, err := api.New(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error instantiating password store")
	}

	pp := &passProvider{
		store: store,
		mutex: &sync.Mutex{},
	}

	return pp, nil
}
