package pass

import (
	"context"
	"encoding/json"
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

			"data_raw": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "String read from Pass.",
			},

			"data": &schema.Schema{
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Map of strings read from Pass.",
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

	data_raw := sec.String()

	d.SetId(path)

	d.Set("data_raw", data_raw)
	log.Printf("[DEBUG] data_raw (id=%s) = %v", d.Id(), d.Get("data_raw"))

	var data map[string]string

	if err := json.Unmarshal([]byte(data_raw), &data); err != nil {
		log.Printf("[WARNING] error unmarshaling data_raw")
		d.Set("data", d.Get("data_raw"))
	} else {
		d.Set("data", data)
	}
	log.Printf("[DEBUG] data (id=%s) = %v", d.Id(), d.Get("data"))

	return nil
}
