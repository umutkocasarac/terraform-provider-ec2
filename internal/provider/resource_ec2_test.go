package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceScaffolding(t *testing.T) {

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceScaffolding,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"scaffolding_resource.ec2", "ami", regexp.MustCompile("^ami-e7527ed7")),
				),
			},
		},
	})
}

const testAccResourceScaffolding = `
resource "scaffolding_resource" "ec2" {
  amount =1
  ami = "ami-e7527ed7"
}
`
