// Copyright (c) HopopOps
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"os"

	cloudbeaver "github.com/hopopops/cloudbeaver-client-go"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure CloudbeaverProvider satisfies various provider interfaces.
var _ provider.Provider = &CloudbeaverProvider{}

// CloudbeaverProvider defines the provider implementation.
type CloudbeaverProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// CloudbeaverProviderModel describes the provider data model.
type CloudbeaverProviderModel struct {
	Host     types.String `tfsdk:"host"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

func (p *CloudbeaverProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "cloudbeaver"
	resp.Version = p.version
}

func (p *CloudbeaverProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				MarkdownDescription: "Cloudbeaver's host",
				Optional:            true,
			},
			"username": schema.StringAttribute{
				MarkdownDescription: "Username of an account with administration privileges",
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Password for the account with administration privileges",
				Optional:            true,
				Sensitive:           true,
			},
		},
	}
}

func (p *CloudbeaverProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data CloudbeaverProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.

	if data.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown CloudBeaver Host",
			"The provider cannot create the CloudBeaver client as there is an unknown configuration value for the CloudBeaver host. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the CLOUDBEAVER_HOST environment variable.",
		)
	}

	if data.Username.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Unknown CloudBeaver Username",
			"The provider cannot create the CloudBeaver client as there is an unknown configuration value for the CloudBeaver username. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the CLOUDBEAVER_USERNAME environment variable.",
		)
	}

	if data.Password.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Unknown CloudBeaver Password",
			"The provider cannot create the CloudBeaver client as there is an unknown configuration value for the CloudBeaver password. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the CLOUDBEAVER_PASSWORD environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	host := os.Getenv("CLOUDBEAVER_HOST")
	username := os.Getenv("CLOUDBEAVER_USERNAME")
	password := os.Getenv("CLOUDBEAVER_PASSWORD")

	if !data.Host.IsNull() {
		host = data.Host.ValueString()
	}

	if !data.Username.IsNull() {
		username = data.Username.ValueString()
	}

	if !data.Password.IsNull() {
		password = data.Password.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing CloudBeaver Host",
			"The provider cannot create the CloudBeaver client as there is a missing or empty value for the CloudBeaver host. "+
				"Set the host value in the configuration or use the CLOUDBEAVER_HOST environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if username == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Missing CloudBeaver Username",
			"The provider cannot create the CloudBeaver client as there is a missing or empty value for the CloudBeaver username. "+
				"Set the username value in the configuration or use the CLOUDBEAVER_USERNAME environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if password == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing CloudBeaver Password",
			"The provider cannot create the CloudBeaver client as there is a missing or empty value for the CloudBeaver password. "+
				"Set the password value in the configuration or use the CLOUDBEAVER_PASSWORD environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Create a new CloudBeaver client using the configuration values
	client, err := cloudbeaver.NewClient(&host, &username, &password)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create CloudBeaver Client",
			"An unexpected error occurred when creating the CloudBeaver client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"CloudBeaver Client Error: "+err.Error(),
		)
		return
	}

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *CloudbeaverProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewTeamResource,
	}
}

func (p *CloudbeaverProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewTeamDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CloudbeaverProvider{
			version: version,
		}
	}
}
