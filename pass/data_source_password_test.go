package pass

import (
	"fmt"
	"testing"

	r "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestDataSourcePassword(t *testing.T) {
	r.Test(t, r.TestCase{
		Providers: testAccProviders,
		Steps: []r.TestStep{
			{
				Config: testDataSourcePasswordConfig,
				Check:  testDataSourcePasswordCheck,
			},
		},
	})
}

var testDataSourcePasswordConfig = `

resource "pass_password" "test" {
    path = "secret/foo"
	password = "0123456789"
    data = {
	  zip = "zap"
    }
}

data "pass_password" "test" {
    path = "${pass_password.test.path}"
}

`

func testDataSourcePasswordCheck(s *terraform.State) error {
	resourceState := s.Modules[0].Resources["data.pass_password.test"]
	if resourceState == nil {
		return fmt.Errorf("resource not found in state %v", s.Modules[0].Resources)
	}

	iState := resourceState.Primary
	if iState == nil {
		return fmt.Errorf("resource has no primary instance")
	}

	if got, want := iState.Attributes["password"], "0123456789"; got != want {
		return fmt.Errorf("data contains %s; want %s", got, want)
	}

	if got, want := iState.Attributes["data.zip"], "zap"; got != want {
		return fmt.Errorf("data contains %s; want %s", got, want)
	}

	return nil
}
