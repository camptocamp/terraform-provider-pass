package pass

import (
	"fmt"
	"testing"

	r "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestResourcePassword(t *testing.T) {
	r.Test(t, r.TestCase{
		Providers: testProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []r.TestStep{
			{
				Config: testResourcePassword_initialConfig,
				Check:  testResourcePassword_initialCheck,
			},
			{
				Config: testResourcePassword_updateConfig,
				Check:  testResourcePassword_updateCheck,
			},
		},
	})
}

var testResourcePassword_initialConfig = `

resource "pass_password" "test" {
    path = "secret/foo"
	password = "0123456789"
    data = {
        zip = "zap"
	}
}

`

func testResourcePassword_initialCheck(s *terraform.State) error {
	resourceState := s.Modules[0].Resources["pass_password.test"]
	if resourceState == nil {
		return fmt.Errorf("resource not found in state")
	}

	instanceState := resourceState.Primary
	if instanceState == nil {
		return fmt.Errorf("resource has no primary instance")
	}

	path := instanceState.ID

	if path != instanceState.Attributes["path"] {
		return fmt.Errorf("id doesn't match path")
	}
	if path != "secret/foo" {
		return fmt.Errorf("unexpected secret path")
	}

	if got, want := instanceState.Attributes["password"], "0123456789"; got != want {
		return fmt.Errorf("data contains %s; want %s", got, want)
	}

	if got, want := instanceState.Attributes["data.zip"], "zap"; got != want {
		return fmt.Errorf("data contains %s; want %s", got, want)
	}

	return nil
}

var testResourcePassword_updateConfig = `

resource "pass_password" "test" {
    path = "secret/foo"
	password = "012345678"
    data = {
        zip = "zoop"
	}
}

`

func testResourcePassword_updateCheck(s *terraform.State) error {
	return nil
}
