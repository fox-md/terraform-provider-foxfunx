// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestToCidrHappyFlow(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::foxfunx::tocidr("10.10.10.0","255.255.255.0")
                }
                `,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownOutputValue("test", knownvalue.StringExact("10.10.10.0/24")),
					},
				},
			},
		},
	})
}

func TestToCidr0000address(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::foxfunx::tocidr("0.0.0.0","255.255.255.255")
                }
                `,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownOutputValue("test", knownvalue.StringExact("0.0.0.0/32")),
					},
				},
			},
		},
	})
}

func TestToCidr0000mask(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::foxfunx::tocidr("0.0.0.0","0.0.0.0")
                }
                `,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownOutputValue("test", knownvalue.StringExact("0.0.0.0/0")),
					},
				},
			},
		},
	})
}

func TestToCidrWrongIp(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::foxfunx::tocidr("10.10.1000.0","255.255.255.0")
                }
                `,
				ExpectError: regexp.MustCompile(`not an IPv4 address/invalid address`),
			},
		},
	})
}

func TestToCidrWrongNetmask(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::foxfunx::tocidr("10.10.10.0","255.255.2555.0")
                }
                `,
				ExpectError: regexp.MustCompile(`not an IPv4 netmask/invalid netmask`),
			},
		},
	})
}
