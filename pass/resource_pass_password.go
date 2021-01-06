package pass

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func passPasswordResource() *schema.Resource {
	return &schema.Resource{
		Create: passPasswordResourceWrite,
		Update: passPasswordResourceWrite,
		Delete: passPasswordResourceDelete,
		Read:   passPasswordResourceRead,

		Schema: map[string]*schema.Schema{
			"path": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Full path where the pass data will be written",
			},

			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Secret password",
				Sensitive:   true,
			},

			"data": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Additional key-value data",
				Sensitive:   true,
			},

			"yaml": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "YAML encoded data",
				Sensitive:   true,
			},

			"body": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Body of the secret",
			},

			"full": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Entire secret contents",
			},
		},
	}
}

func passPasswordResourceWrite(d *schema.ResourceData, meta interface{}) error {
	path := d.Get("path").(string)
	log.Printf("writing secret to path %s", path)

	pp := meta.(*passProvider)
	pp.mutex.Lock()
	defer pp.mutex.Unlock()
	st := pp.store

	password := d.Get("password").(string)
	data := d.Get("data").(map[string]interface{})
	yaml := d.Get("yaml").(string)

	if len(data) != 0 && yaml != "" {
		return errors.New("can't set data and yaml at the same time")
	}

	value := password

	if yaml != "" {
		value += "\n---\n" + yaml
	} else {

		elems := make([]string, len(data))
		for k, v := range data {
			elems = append(elems, fmt.Sprintf("%s: %s", k, v))
		}

		value += strings.Join(elems, "\n")
	}

	container := secretContainer{value: value}
	err := st.Set(context.Background(), path, container)

	if err != nil {
		return errors.Wrapf(err, "failed to write secret at %s", path)
	}

	d.SetId(path)
	return nil
}

func passPasswordResourceDelete(d *schema.ResourceData, meta interface{}) error {
	path := d.Id()

	pp := meta.(*passProvider)
	pp.mutex.Lock()
	defer pp.mutex.Unlock()
	st := pp.store
	log.Printf("deleting secret at %s", path)
	err := st.Remove(context.Background(), path)
	if err != nil {
		return errors.Wrapf(err, "failed to delete secret at %s", path)
	}

	return nil
}

func passPasswordResourceRead(d *schema.ResourceData, meta interface{}) error {
	path := d.Id()

	pp := meta.(*passProvider)
	pp.mutex.Lock()
	defer pp.mutex.Unlock()
	return populateResourceData(d, pp, path, false)
}
