// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestFileContainsCaseSensitive(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.txt")
	fileText := "HELLO world\n"
	searchText := "hello WORLD"
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					err := os.WriteFile(filePath, []byte(fileText), 0644)
					if err != nil {
						t.Fatal(err)
					}
				},
				Config: `
                output "test" {
                    value = provider::foxfunx::filecontains("` + filePath + `","` + searchText + `")
                }
                `,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownOutputValue("test", knownvalue.Bool(false)),
					},
				},
			},
		},
	})
}

func TestFileContainsCaseInsensitive(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.txt")
	fileText := "HELLO world\nWoRld HellO"
	searchText := "hello WORLD"
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					err := os.WriteFile(filePath, []byte(fileText), 0644)
					if err != nil {
						t.Fatal(err)
					}
				},
				Config: `
                output "test" {
                    value = provider::foxfunx::filecontains("` + filePath + `","` + searchText + `", false)
                }
                `,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownOutputValue("test", knownvalue.Bool(true)),
					},
				},
			},
		},
	})
}

func TestFileContainsEnTrue(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.txt")
	searchText := "hello world"
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					err := os.WriteFile(filePath, []byte(searchText), 0644)
					if err != nil {
						t.Fatal(err)
					}
				},
				Config: `
                output "test" {
                    value = provider::foxfunx::filecontains("` + filePath + `","` + searchText + `")
                }
                `,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownOutputValue("test", knownvalue.Bool(true)),
					},
				},
			},
		},
	})
}

func TestFileContainsEnFalse(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.txt")
	searchText := "hello world"
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					err := os.WriteFile(filePath, []byte("some dummy text"), 0644)
					if err != nil {
						t.Fatal(err)
					}
				},
				Config: `
                output "test" {
                    value = provider::foxfunx::filecontains("` + filePath + `","` + searchText + `")
                }
                `,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownOutputValue("test", knownvalue.Bool(false)),
					},
				},
			},
		},
	})
}

func TestFileContainsRuTrue(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.txt")
	searchText := "привет"
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					err := os.WriteFile(filePath, []byte(searchText), 0644)
					if err != nil {
						t.Fatal(err)
					}
				},
				Config: `
                output "test" {
                    value = provider::foxfunx::filecontains("` + filePath + `","` + searchText + `")
                }
                `,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownOutputValue("test", knownvalue.Bool(true)),
					},
				},
			},
		},
	})
}

func TestFileContainsRuFalse(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.txt")
	searchText := "привет"
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					err := os.WriteFile(filePath, []byte("превет"), 0644)
					if err != nil {
						t.Fatal(err)
					}
				},
				Config: `
                output "test" {
                    value = provider::foxfunx::filecontains("` + filePath + `","` + searchText + `")
                }
                `,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownOutputValue("test", knownvalue.Bool(false)),
					},
				},
			},
		},
	})
}

func TestFileContainsFrTrue(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.txt")
	searchText := "café"
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					err := os.WriteFile(filePath, []byte(searchText), 0644)
					if err != nil {
						t.Fatal(err)
					}
				},
				Config: `
                output "test" {
                    value = provider::foxfunx::filecontains("` + filePath + `","` + searchText + `")
                }
                `,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownOutputValue("test", knownvalue.Bool(true)),
					},
				},
			},
		},
	})
}

func TestFileContainsFrFalse(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.txt")
	searchText := "café"
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					err := os.WriteFile(filePath, []byte("cafe"), 0644)
					if err != nil {
						t.Fatal(err)
					}
				},
				Config: `
                output "test" {
                    value = provider::foxfunx::filecontains("` + filePath + `","` + searchText + `")
                }
                `,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectKnownOutputValue("test", knownvalue.Bool(false)),
					},
				},
			},
		},
	})
}

func TestFileContainsFailEmptySearchText(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.txt")
	searchText := "hello world"
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					err := os.WriteFile(filePath, []byte(searchText), 0644)
					if err != nil {
						t.Fatal(err)
					}
				},
				Config: `
                output "test" {
                    value = provider::foxfunx::filecontains("` + filePath + `","")
                }
                `,
				ExpectError: regexp.MustCompile(`Invalid value for "search_text" parameter: Invalid Parameter Value Length:
string length must be at least 1`),
			},
		},
	})
}

func TestFileContainsFailEmptyFilePath(t *testing.T) {
	searchText := "hello world"
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                output "test" {
                    value = provider::foxfunx::filecontains("","` + searchText + `")
                }
                `,
				ExpectError: regexp.MustCompile(`Invalid value for "file_path" parameter: Invalid Parameter Value Length:
string length must be at least 1`),
			},
		},
	})
}
