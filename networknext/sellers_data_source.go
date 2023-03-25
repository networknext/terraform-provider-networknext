package networknext

import (
    "context"

    "github.com/hashicorp/terraform-plugin-framework/datasource"
    "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
    "github.com/hashicorp/terraform-plugin-framework/types"
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

type sellersDataSourceModel struct {
    Sellers []SellerModel `tfsdk:"sellers"`
}

func (d *sellersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_sellers"
}

func (d *sellersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
    resp.Schema = schema.Schema{
        Attributes: map[string]schema.Attribute{
            "sellers": schema.ListNestedAttribute{
                Computed: true,
                NestedObject: schema.NestedAttributeObject{
                    Attributes: map[string]schema.Attribute{
                        "id": schema.Int64Attribute{
                            Computed: true,
                        },
                        "name": schema.StringAttribute{
                            Computed: true,
                        },
                        "customer_id": schema.Int64Attribute{
                            Computed: true,
                        },
                    },
                },
            },
        },
    }
}

func (d *sellersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
    if req.ProviderData == nil {
        return
    }
    d.client = req.ProviderData.(*Client)
}

type SellerModel struct {
    Id              types.Int64  `tfsdk:"id"`
    Name            types.String `tfsdk:"name"`
    CustomerId      types.Int64  `tfsdk:"customer_id"`
}

type SellerData struct {
    SellerId         uint64 `json:"seller_id"`
    SellerName       string `json:"seller_name"`
    CustomerId      uint64 `json:"customer_id"`
}

type SellersResponse struct {
    Sellers []SellerData `json:"sellers"`
    Error  string        `json:"error"`
}

func (d *sellersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

    sellersResponse := SellersResponse{}
    
    err := d.client.GetJSON("admin/sellers", &sellersResponse)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to get networknext sellers",
            "An error occurred when calling the networknext API to get sellers. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if sellersResponse.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to get networknext sellers",
            "An error occurred when calling the networknext API to get sellers. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+sellersResponse.Error,
        )
        return
    }

    var state sellersDataSourceModel

    for i := range sellersResponse.Sellers {

        sellerState := SellerModel{
            Id:              types.Int64Value(int64(sellersResponse.Sellers[i].SellerId)),
            Name:            types.StringValue(sellersResponse.Sellers[i].SellerName),
            CustomerId:      types.Int64Value(int64(sellersResponse.Sellers[i].CustomerId)),
        }

        state.Sellers = append(state.Sellers, sellerState)
    }

    diags := resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}
