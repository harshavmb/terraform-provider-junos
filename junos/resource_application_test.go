package junos_test

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccJunosApplication_basic(t *testing.T) {
	if os.Getenv("TESTACC_SWITCH") == "" && os.Getenv("TESTACC_ROUTER") == "" {
		resource.Test(t, resource.TestCase{
			PreCheck:  func() { testAccPreCheck(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: testAccJunosApplicationConfigCreate(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("junos_application.testacc_app", "protocol", "tcp"),
						resource.TestCheckResourceAttr("junos_application.testacc_app", "destination_port", "22"),
					),
				},
				{
					Config: testAccJunosApplicationConfigUpdate(),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("junos_application.testacc_app", "protocol", "tcp"),
						resource.TestCheckResourceAttr("junos_application.testacc_app", "destination_port", "22"),
						resource.TestCheckResourceAttr("junos_application.testacc_app", "source_port", "1024-65535"),
					),
				},
				{
					ResourceName:      "junos_application.testacc_app",
					ImportState:       true,
					ImportStateVerify: true,
				},
				{
					ResourceName:      "junos_application.testacc_app2",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	}
}

func testAccJunosApplicationConfigCreate() string {
	return `
resource "junos_application" "testacc_app" {
  name             = "testacc_app"
  protocol         = "tcp"
  destination_port = 22
}
`
}

func testAccJunosApplicationConfigUpdate() string {
	return `
resource "junos_application" "testacc_app" {
  name                 = "testacc_app"
  protocol             = "tcp"
  destination_port     = "22"
  application_protocol = "ssh"
  description          = "ssh protocol"
  inactivity_timeout   = 900
  source_port          = "1024-65535"
}
resource "junos_application" "testacc_app2" {
  name               = "testacc_app2"
  protocol           = "tcp"
  ether_type         = "0x0800"
  rpc_program_number = "0-0"
  uuid               = "AAAAA0AA-B9B0-CCcc-DDDD-EEEffFFFAAAA"
}
`
}
