// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider              = &foxfunxProvider{}
	_ provider.ProviderWithFunctions = &foxfunxProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &foxfunxProvider{
			version: version,
		}
	}
}

// foxfunxProvider is the provider implementation.
type foxfunxProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// Metadata returns the provider type name.
func (p *foxfunxProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "foxfunx"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *foxfunxProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "`foxfunx` is a function-only provider.\n\n" +
			"`foxfunx` includes below functions:" +
			"- `direxists` given a path, return boolean depending on directory existence. Fails for files." +
			"- `tocidr` given subnet and netmask, return subnet and netmask in the cidr format. Fails for invalid subnet or netmasks.",
	}
}

// Configure prepares an API client for data sources and resources.
func (p *foxfunxProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

// DataSources defines the data sources implemented in the provider.
func (p *foxfunxProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

// Resources defines the resources implemented in the provider.
func (p *foxfunxProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}

// Functions defines the functions implemented in the provider.
func (p *foxfunxProvider) Functions(_ context.Context) []func() function.Function {
	return []func() function.Function{
		NewDirExistsFunction,
		NewToCidrFunction,
	}
}
