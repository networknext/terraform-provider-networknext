package accelerate

import (
    "context"

    "github.com/hashicorp/terraform-plugin-framework/datasource"
)

var (
    _ datasource.DataSource              = &relayKeypairsDataSource{}
    _ datasource.DataSourceWithConfigure = &relayKeypairsDataSource{}
)

func NewRelayKeypairsDataSource() datasource.DataSource {
    return &relayKeypairsDataSource{}
}

type relayKeypairsDataSource struct {
    client *Client
}

func (d *relayKeypairsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_relay_keypairs"
}

func (d *relayKeypairsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
    resp.Schema = RelayKeypairsSchema()
}

func (d *relayKeypairsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
    if req.ProviderData == nil {
        return
    }
    d.client = req.ProviderData.(*Client)
}

func (d *relayKeypairsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

    relayKeypairsResponse := ReadRelayKeypairsResponse{}
    
    err := d.client.GetJSON(ctx, "admin/relay_keypairs", &relayKeypairsResponse)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to get relay keypairs",
            "An unexpected error occurred when calling the network next accelerate API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if relayKeypairsResponse.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to get relay keypairs",
            "The network next accelerate API returned an error: "+relayKeypairsResponse.Error,
        )
        return
    }

    var state RelayKeypairsModel

    for i := range relayKeypairsResponse.RelayKeypairs {
        var relayKeypairState RelayKeypairModel
        RelayKeypairDataToModel(&relayKeypairsResponse.RelayKeypairs[i], &relayKeypairState)
        state.RelayKeypairs = append(state.RelayKeypairs, relayKeypairState)
    }

    diags := resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}
