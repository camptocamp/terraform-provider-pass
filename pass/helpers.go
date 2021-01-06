package pass

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

// Used as interface to write gopass secrets
type secretContainer struct {
	value string
}

func (c secretContainer) Bytes() []byte {
	return []byte(c.value)
}

// Used as context for secret template
type templateData struct {
	Data map[string]interface{}
}

// generic function to be used in resource and data source
func populateResourceData(d *schema.ResourceData, provider *passProvider, path string, readData bool) error {
	st := provider.store
	log.Printf("reading %s from gopass", path)

	sec, err := st.Get(context.Background(), path, "") //TODO: support getting a revison via terraform?
	if err != nil {
		return errors.Wrapf(err, "failed to read password at %s", path)
	}

	// Retrieve all data items if keys exist
	if readData {
		var keys = sec.Keys()
		if len(keys) != 0 {
			log.Printf("populating data with keys")
			var data = make(map[string]interface{})
			for _, key := range keys {
				data[key] = sec.Get(key)
			}
			d.Set("data", data)
		}
	}

	if err := d.Set("password", sec.Get("Password")); err != nil {
		log.Printf("error when setting password: %v", err)
		return err
	}

	if err := d.Set("body", sec.GetBody()); err != nil {
		log.Printf("error when setting body: %v", err)
		return err
	}

	if err := d.Set("full", string(sec.Bytes())); err != nil {
		log.Printf("error when setting full: %v", err)
		return err
	}

	return nil

}
