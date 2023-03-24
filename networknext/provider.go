package networknext

import (
    "context"
    "os"

    "github.com/hashicorp/terraform-plugin-framework/datasource"
    "github.com/hashicorp/terraform-plugin-framework/path"
    "github.com/hashicorp/terraform-plugin-framework/provider"
    "github.com/hashicorp/terraform-plugin-framework/provider/schema"
    "github.com/hashicorp/terraform-plugin-framework/resource"
    "github.com/hashicorp/terraform-plugin-framework/types"
    "github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
    _ provider.Provider = &networknextProvider{}
)

type networknextProvider struct{}

type networknextProviderModel struct {
    HostName types.String `tfsdk:"hostname"`
    ApiKey   types.String `tfsdk:"api_key"`
}

func New() provider.Provider {
    return &networknextProvider{}
}

func (p *networknextProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
    resp.TypeName = "networknext"
}

func (p *networknextProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
    resp.Schema = schema.Schema{
        Attributes: map[string]schema.Attribute{
            "hostname": schema.StringAttribute{
                Optional: true,
            },
            "api_key": schema.StringAttribute{
                Optional:  true,
                Sensitive: true,
            },
        },
    }
}

func (p *networknextProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

    tflog.Info(ctx, "Configuring networknext client")

    // retrieve provider data from configuration

    var config networknextProviderModel
    diags := req.Config.Get(ctx, &config)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    // if practitioner provided a configuration value for any of the attributes, it must be a known value.

    if config.HostName.IsUnknown() {
        resp.Diagnostics.AddAttributeError(
            path.Root("hostname"),
            "Unknown networknext API hostname",
            "The provider cannot create the networknext API client as there is an unknown configuration value for the API hostname. "+
                "Either target apply the source of the value first, set the value statically in the configuration, or use the NETWORKNEXT_HOSTNAME environment variable.",
        )
    }

    if config.ApiKey.IsUnknown() {
        resp.Diagnostics.AddAttributeError(
            path.Root("api_key"),
            "Unknown networknext API key",
            "The provider cannot create the networknext API client as there is an unknown configuration value for the networknext API key. "+
                "Either target apply the source of the value first, set the value statically in the configuration, or use the NETWORKNEXT_API_KEY environment variable.",
        )
    }

    if resp.Diagnostics.HasError() {
        return
    }

    // default values to environment variables, but override with Terraform configuration value if set.

    hostname := os.Getenv("NETWORKNEXT_HOSTNAME")
    api_key := os.Getenv("NETWORKNEXT_API_KEY")

    if !config.HostName.IsNull() {
        hostname = config.HostName.ValueString()
    }

    if !config.ApiKey.IsNull() {
        api_key = config.ApiKey.ValueString()
    }

    // if any of the expected configurations are missing, return errors

    if hostname == "" {
        resp.Diagnostics.AddAttributeError(
            path.Root("hostname"),
            "Missing networknext API hostname",
            "The provider cannot create the networknext API client as there is a missing or empty value for the networknext API hostname. "+
                "Set the hostname value in the configuration or use the NETWORKNEXT_HOSTNAME environment variable. "+
                "If either is already set, ensure the value is not empty.",
        )
    }

    if api_key == "" {
        resp.Diagnostics.AddAttributeError(
            path.Root("api_key"),
            "Missing networknext API key",
            "The provider cannot create the networknext API client as there is a missing or empty value for the networknext API key. "+
                "Set the api_key value in the configuration or use the NETWORKNEXT_API_KEY environment variable. "+
                "If either is already set, ensure the value is not empty.",
        )
    }

    if resp.Diagnostics.HasError() {
        return
    }

    ctx = tflog.SetField(ctx, "networknext_hostname", hostname)
    ctx = tflog.SetField(ctx, "networknext_api_key", api_key)
    ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "networknext_api_key")
    
    tflog.Debug(ctx, "Creating networknext client")

    // todo
    /*
    client, err := hashicups.NewClient(&host, &username, &password)
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to Create HashiCups API Client",
            "An unexpected error occurred when creating the HashiCups API client. "+
                "If the error is not clear, please contact the provider developers.\n\n"+
                "HashiCups Client Error: "+err.Error(),
        )
        return
    }
    */

    /*
    // Make the HashiCups client available during DataSource and Resource type Configure methods.
    resp.DataSourceData = client
    resp.ResourceData = client
    */

    tflog.Info(ctx, "Configured networknext client", map[string]any{"success": true})
}

func (p *networknextProvider) DataSources(_ context.Context) []func() datasource.DataSource {
    return []func() datasource.DataSource {
        NewCustomersDataSource,
    }
}

func (p *networknextProvider) Resources(_ context.Context) []func() resource.Resource {
    return nil
}
