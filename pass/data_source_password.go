package pass

import (
	"context"
	"log"

	"github.com/gopasspw/gopass/pkg/store/root"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pkg/errors"
)

func passwordDataSource() *schema.Resource {
	return &schema.Resource{
		Read: passwordDataSourceRead,

		Schema: map[string]*schema.Schema{
			"path": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Full path from which a password will be read.",
			},

			"password": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "secret password.",
			},

			"data": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "additional secret data.",
			},

			"body": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "raw secret data if not YAML.",
			},

			"full": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "entire secret contents",
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

	if err := d.Set("password", sec.Password()); err != nil {
		log.Printf("[ERROR] Error when setting password: %v", err)
		return err
	}
	if err := d.Set("data", sec.Data()); err != nil {
		log.Printf("[ERROR] Error when setting data: %v", err)
		return err
	}
	if err := d.Set("body", sec.Body()); err != nil {
		log.Printf("[ERROR] Error when setting body: %v", err)
		return err
	}
	if err := d.Set("full", sec.String()); err != nil {
		log.Printf("[ERROR] Error when setting full: %v", err)
		return err
	}

	return nil
}
