package accelerate

import (
    "context"

    "github.com/hashicorp/terraform-plugin-framework/datasource"
)

var (
    _ datasource.DataSource              = &buyerKeypairsDataSource{}
    _ datasource.DataSourceWithConfigure = &buyerKeypairsDataSource{}
)

func NewBuyerKeypairsDataSource() datasource.DataSource {
    return &buyerKeypairsDataSource{}
}

type buyerKeypairsDataSource struct {
    client *Client
}

func (d *buyerKeypairsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_buyer_keypairs"
}

func (d *buyerKeypairsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
    resp.Schema = BuyerKeypairsSchema()
}

func (d *buyerKeypairsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
    if req.ProviderData == nil {
        return
    }
    d.client = req.ProviderData.(*Client)
}

func (d *buyerKeypairsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

    buyerKeypairsResponse := ReadBuyerKeypairsResponse{}
    
    err := d.client.GetJSON(ctx, "admin/buyer_keypairs", &buyerKeypairsResponse)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to get buyer keypairs",
            "An unexpected error occurred when calling the network next accelerate API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if buyerKeypairsResponse.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to get buyer keypairs",
            "The network next accelerate API returned an error: "+buyerKeypairsResponse.Error,
        )
        return
    }

    var state BuyerKeypairsModel

    for i := range buyerKeypairsResponse.BuyerKeypairs {
        var buyerKeypairState BuyerKeypairModel
        BuyerKeypairDataToModel(&buyerKeypairsResponse.BuyerKeypairs[i], &buyerKeypairState)
        state.BuyerKeypairs = append(state.BuyerKeypairs, buyerKeypairState)
    }

    diags := resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}
