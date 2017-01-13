package pass

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
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

			"data": &schema.Schema{
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Map of strings read from Pass.",
			},
		},
	}
}

func passwordDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*api.Client)

	path := d.Get("path").(string)

	log.Printf("[DEBUG] Reading %s from Pass", path)
	password, err := os.exec.Command("pass", path)
	if err != nil {
		return fmt.Errorf("error reading from Pass: %s", err)
	}

	d.Set("data", password)

	return nil
}
