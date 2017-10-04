package pass

import (
	"context"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/justwatchcom/gopass/store/root"
	"github.com/justwatchcom/gopass/store/secret"
	"github.com/pkg/errors"
)

func passPasswordResource() *schema.Resource {
	return &schema.Resource{
		Create: passPasswordResourceWrite,
		Update: passPasswordResourceWrite,
		Delete: passPasswordResourceDelete,
		Read:   passPasswordResourceRead,

		Schema: map[string]*schema.Schema{
			"path": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Full path where the pass password will be written.",
			},

			// Data is passed as JSON so that an arbitrary structure is
			// possible, rather than forcing e.g. all values to be strings.
			"data": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "JSON-encoded secret data to write.",
			},
		},
	}
}

func passPasswordResourceWrite(d *schema.ResourceData, meta interface{}) error {
	path := d.Get("path").(string)

	st := meta.(*root.Store)
	data := d.Get("data").(string)
	sec := secret.New("", data)
	err := st.Set(context.Background(), path, sec)
	if err != nil {
		return errors.Wrapf(err, "failed to write secret at %s", path)
	}

	d.SetId(path)

	return nil
}

func passPasswordResourceDelete(d *schema.ResourceData, meta interface{}) error {
	path := d.Id()

	st := meta.(*root.Store)
	log.Printf("[DEBUG] Deleting generic Vault from %s", path)
	err := st.Delete(context.Background(), path)
	if err != nil {
		return errors.Wrapf(err, "failed to delete password at %s", path)
	}

	return nil
}

func passPasswordResourceRead(d *schema.ResourceData, meta interface{}) error {
	path := d.Id()

	st := meta.(*root.Store)
	sec, err := st.Get(context.Background(), path)
	if err != nil {
		errors.Wrapf(err, "failed to retrieve password at %s", path)
	}

	d.Set("data", sec.Body())

	return nil
}
