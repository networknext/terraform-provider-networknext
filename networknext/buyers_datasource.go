package networknext

import (
    "context"

    "github.com/hashicorp/terraform-plugin-framework/datasource"
)

var (
    _ datasource.DataSource              = &buyersDataSource{}
    _ datasource.DataSourceWithConfigure = &buyersDataSource{}
)

func NewBuyersDataSource() datasource.DataSource {
    return &buyersDataSource{}
}

type buyersDataSource struct {
    client *Client
}

func (d *buyersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_buyers"
}

func (d *buyersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
    resp.Schema = BuyersSchema()
}

func (d *buyersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
    if req.ProviderData == nil {
        return
    }
    d.client = req.ProviderData.(*Client)
}

func (d *buyersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

    buyersResponse := ReadBuyersResponse{}
    
    err := d.client.GetJSON(ctx, "admin/buyers", &buyersResponse)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to get buyers",
            "An error occurred when calling the network next API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if buyersResponse.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to get buyers",
            "An error occurred when calling the network next API to get buyers. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+buyersResponse.Error,
        )
        return
    }

    var state BuyersModel

    for i := range buyersResponse.Buyers {
        var buyerState BuyerModel
        BuyerDataToModel(&buyersResponse.Buyers[i], &buyerState)
        state.Buyers = append(state.Buyers, buyerState)
    }

    diags := resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}
