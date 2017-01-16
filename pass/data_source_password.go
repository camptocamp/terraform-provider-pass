package pass

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"

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

			"data_row": &schema.Schema{
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

	log.Printf("[DEBUG] Using PASSWORD_STORE_DIR=%v", os.Getenv("PASSWORD_STORE_DIR"))
	log.Printf("[DEBUG] Reading %s from Pass", path)
	output, err := exec.Command("pass", path).Output()
	if err != nil {
		return fmt.Errorf("error reading from Pass: %s", err)
	}
	data_raw := string(output)

	d.Set("data_row", data_raw)

	var data map[string]string

	if err := json.Unmarshal(output, &data); err != nil {
		return fmt.Errorf("error unmarshaling data_row")
	}
	d.Set("data", data)
	log.Printf("[DEBUG] data = %v", d.Get("data"))

	return nil
}
