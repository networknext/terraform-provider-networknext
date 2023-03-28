package networknext

import (
    "context"

    "github.com/hashicorp/terraform-plugin-framework/datasource"
)

var (
    _ datasource.DataSource              = &sellersDataSource{}
    _ datasource.DataSourceWithConfigure = &sellersDataSource{}
)

func NewSellersDataSource() datasource.DataSource {
    return &sellersDataSource{}
}

type sellersDataSource struct {
    client *Client
}

func (d *sellersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_sellers"
}

func (d *sellersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
    resp.Schema = SellersSchema()
}

func (d *sellersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
    if req.ProviderData == nil {
        return
    }
    d.client = req.ProviderData.(*Client)
}

func (d *sellersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

    sellersResponse := ReadSellersResponse{}
    
    err := d.client.GetJSON(ctx, "admin/sellers", &sellersResponse)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to get networknext sellers",
            "An unexpected error occurred when calling the networknext API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if sellersResponse.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to get networknext sellers",
            "The networknext API returned an error: "+sellersResponse.Error,
        )
        return
    }

    var state SellersModel

    for i := range sellersResponse.Sellers {
        var sellerState SellerModel
        SellerDataToModel(&(sellersResponse.Sellers[i]), &sellerState)
        state.Sellers = append(state.Sellers, sellerState)
    }

    diags := resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}
