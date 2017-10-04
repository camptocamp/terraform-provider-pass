package pass

import (
	"context"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/justwatchcom/gopass/store/root"
	"github.com/pkg/errors"
)

func passwordDataSource() *schema.Resource {
	return &schema.Resource{
		Read: passwordDataSourceRead,

		Schema: map[string]*schema.Schema{
			"path": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Full path from which a password will be read.",
			},

			"password": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "secret password.",
			},

			"data": &schema.Schema{
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "additional secret data.",
			},
		},
	}
}

func passwordDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	path := d.Get("path").(string)

	st := meta.(*root.Store)
	log.Printf("[DEBUG] Reading %s from Pass", path)

	sec, err := st.Get(context.Background(), path)
	if err != nil {
		return errors.Wrapf(err, "failed to read password at %s", path)
	}

	d.SetId(path)

	d.Set("password", sec.Password())
	d.Set("data", sec.Data())

	return nil
}
