// Copyright (c) HopopOps
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccTeamDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccTeamDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.cloudbeaver_team.test", "id", "infra"),
					resource.TestCheckResourceAttr("data.cloudbeaver_team.test", "name", "infra"),
					resource.TestCheckResourceAttr("data.cloudbeaver_team.test", "description", ""),
				),
			},
		},
	})
}

const testAccTeamDataSourceConfig = `
data "cloudbeaver_team" "test" {
  id = "infra"
}
`
