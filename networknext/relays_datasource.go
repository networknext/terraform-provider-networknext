package networknext

import (
    "context"

    "github.com/hashicorp/terraform-plugin-framework/datasource"
)

var (
    _ datasource.DataSource              = &relaysDataSource{}
    _ datasource.DataSourceWithConfigure = &relaysDataSource{}
)

func NewRelaysDataSource() datasource.DataSource {
    return &relaysDataSource{}
}

type relaysDataSource struct {
    client *Client
}

func (d *relaysDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_relays"
}

func (d *relaysDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
    resp.Schema = RelaysSchema()
}

func (d *relaysDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
    if req.ProviderData == nil {
        return
    }
    d.client = req.ProviderData.(*Client)
}

func (d *relaysDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

    relaysResponse := ReadRelaysResponse{}
    
    err := d.client.GetJSON(ctx, "admin/relays", &relaysResponse)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to get relays",
            "An unexpected error occurred when calling the network next API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if relaysResponse.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to get relays",
            "An error occurred when calling the network next API to get relays. "+
                "Network Next Client Error: "+relaysResponse.Error,
        )
        return
    }

    var state RelaysModel

    for i := range relaysResponse.Relays {
        var relayState RelayModel
        RelayDataToModel(&relaysResponse.Relays[i], &relayState)
        state.Relays = append(state.Relays, relayState)
    }

    diags := resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}
