// Copyright (c) HopopOps
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccTeamResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccTeamResourceConfig("one"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("cloudbeaver_team.test", "id", "one"),
					resource.TestCheckResourceAttr("cloudbeaver_team.test", "name", "one"),
					resource.TestCheckResourceAttr("cloudbeaver_team.test", "description", ""),
				),
			},
			// ImportState testing
			{
				ResourceName:      "cloudbeaver_team.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: testAccTeamResourceConfigWithDescription("two", "Some description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("cloudbeaver_team.test", "id", "two"),
					resource.TestCheckResourceAttr("cloudbeaver_team.test", "name", "two"),
					resource.TestCheckResourceAttr("cloudbeaver_team.test", "description", "Some description"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccTeamResourceConfig(id string) string {
	return fmt.Sprintf(`
resource "cloudbeaver_team" "test" {
  id = %[1]q
}
`, id)
}

func testAccTeamResourceConfigWithDescription(id, description string) string {
	return fmt.Sprintf(`
resource "cloudbeaver_team" "test" {
  id = %[1]q
  description = %[2]q
}
`, id, description)
}
