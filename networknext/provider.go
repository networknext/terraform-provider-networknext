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
    APIKey   types.String `tfsdk:"api_key"`
}

func New() provider.Provider {
    return &networknextProvider{}
}

func (p *networknextProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
    resp.TypeName = "networknext"
}

func (p *networknextProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
    resp.Schema = schema.Schema{
        Description: "Configure Network Next.",
        Attributes: map[string]schema.Attribute{
            "hostname": schema.StringAttribute{
                Description: "The URI for the Network Next API. May also be provided via NETWORK_NEXT_HOSTNAME environment variable.",
                Optional: true,
            },
            "api_key": schema.StringAttribute{
                Description: "The API Key that allows interaction with the Network Next API. May also be provided via NETWORK_NEXT_API_KEY environment variable.",
                Optional:  true,
                Sensitive: true,
            },
        },
    }
}

func (p *networknextProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

    tflog.Info(ctx, "Configuring network next client")

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
            "Unknown network next API hostname",
            "The provider cannot create the network next API client as there is an unknown configuration value for the network next API hostname. "+
                "Either target apply the source of the value first, set the value statically in the configuration, or use the NETWORK_NEXT_HOSTNAME environment variable.",
        )
    }

    if config.APIKey.IsUnknown() {
        resp.Diagnostics.AddAttributeError(
            path.Root("api_key"),
            "Unknown network next API key",
            "The provider cannot create the network next API client as there is an unknown configuration value for the network next API key. "+
                "Either target apply the source of the value first, set the value statically in the configuration, or use the NETWORK_NEXT_API_KEY environment variable.",
        )
    }

    if resp.Diagnostics.HasError() {
        return
    }

    // default values to environment variables, but override with Terraform configuration value if set.

    hostname := os.Getenv("NETWORK_NEXT_HOSTNAME")
    api_key := os.Getenv("NETWORK_NEXT_API_KEY")

    if !config.HostName.IsNull() {
        hostname = config.HostName.ValueString()
    }

    if !config.APIKey.IsNull() {
        api_key = config.APIKey.ValueString()
    }

    // if any of the expected configurations are missing, return errors

    if hostname == "" {
        resp.Diagnostics.AddAttributeError(
            path.Root("hostname"),
            "Missing network next API hostname",
            "The provider cannot create the network next API client as there is a missing or empty value for the network next API hostname. "+
                "Set the hostname value in the configuration or use the NETWORK_NEXT_HOSTNAME environment variable. "+
                "If either is already set, ensure the value is not empty.",
        )
    }

    if api_key == "" {
        resp.Diagnostics.AddAttributeError(
            path.Root("api_key"),
            "Missing network next API key",
            "The provider cannot create the network next API client as there is a missing or empty value for the network next API key. "+
                "Set the api_key value in the configuration or use the NETWORK_NEXT_API_KEY environment variable. "+
                "If either is already set, ensure the value is not empty.",
        )
    }

    if resp.Diagnostics.HasError() {
        return
    }

    ctx = tflog.SetField(ctx, "networknext_hostname", hostname)
    ctx = tflog.SetField(ctx, "networknext_api_key", api_key)
    ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "networknext_api_key")
    
    tflog.Debug(ctx, "Creating network next client")

    client, err := NewClient(ctx, hostname, api_key)
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to create network next API client",
            "An error occurred when creating the network next API client. "+
                "Please check that the hostname is correct and your api key is valid.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    resp.DataSourceData = client

    resp.ResourceData = client

    tflog.Info(ctx, "Configured network next client", map[string]any{"success": true})
}

func (p *networknextProvider) DataSources(_ context.Context) []func() datasource.DataSource {
    return []func() datasource.DataSource {
        NewCustomersDataSource,
        NewBuyersDataSource,
        NewSellersDataSource,
        NewDatacentersDataSource,
        NewRelaysDataSource,
        NewRouteShadersDataSource,
        NewBuyerDatacenterSettingsDataSource,
        NewRelayKeypairsDataSource,
    }
}

func (p *networknextProvider) Resources(_ context.Context) []func() resource.Resource {
    return []func() resource.Resource {
        NewCustomerResource,
        NewSellerResource,
        NewBuyerResource,
        NewDatacenterResource,
        NewRelayResource,
        NewRouteShaderResource,
        NewBuyerDatacenterSettingsResource,
        NewRelayKeypairResource,
    }
}
