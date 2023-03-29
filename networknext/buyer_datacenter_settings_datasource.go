package networknext

import (
    "context"

    "github.com/hashicorp/terraform-plugin-framework/datasource"
)

var (
    _ datasource.DataSource              = &buyerDatacenterSettingsDataSource{}
    _ datasource.DataSourceWithConfigure = &buyerDatacenterSettingsDataSource{}
)

func NewBuyerDatacenterSettingsDataSource() datasource.DataSource {
    return &buyerDatacenterSettingsDataSource{}
}

type buyerDatacenterSettingsDataSource struct {
    client *Client
}

func (d *buyerDatacenterSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_buyer_datacenter_settings"
}

func (d *buyerDatacenterSettingsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
    resp.Schema = BuyerDatacenterSettingsListSchema()
}

func (d *buyerDatacenterSettingsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
    if req.ProviderData == nil {
        return
    }
    d.client = req.ProviderData.(*Client)
}

func (d *buyerDatacenterSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

    buyerDatacenterSettingsListResponse := ReadBuyerDatacenterSettingsListResponse{}
    
    err := d.client.GetJSON(ctx, "admin/buyer_datacenter_settings", &buyerDatacenterSettingsListResponse)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to get networknext buyer datacenter settings",
            "An unexpected error occurred when calling the networknext API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if buyerDatacenterSettingsListResponse.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to get networknext buyer datacenter settings",
            "The networknext API returned an error: "+buyerDatacenterSettingsListResponse.Error,
        )
        return
    }

    var state BuyerDatacenterSettingsListModel

    for i := range buyerDatacenterSettingsListResponse.Settings {
        var settingsState BuyerDatacenterSettingsModel
        BuyerDatacenterSettingsDataToModel(&buyerDatacenterSettingsListResponse.Settings[i], &settingsState)
        state.Settings = append(state.Settings, settingsState)
    }

    diags := resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}
