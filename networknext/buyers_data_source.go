package networknext

import (
    "context"

    "github.com/hashicorp/terraform-plugin-framework/datasource"
    "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
    "github.com/hashicorp/terraform-plugin-framework/types"
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

type buyersDataSourceModel struct {
    Buyers []BuyerModel `tfsdk:"buyers"`
}

func (d *buyersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_buyers"
}

func (d *buyersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
    resp.Schema = schema.Schema{
        Attributes: map[string]schema.Attribute{
            "buyers": schema.ListNestedAttribute{
                Computed: true,
                NestedObject: schema.NestedAttributeObject{
                    Attributes: map[string]schema.Attribute{
                        "id": schema.Int64Attribute{
                            Computed: true,
                        },
                        "name": schema.StringAttribute{
                            Computed: true,
                        },
                        "public_key_base64": schema.StringAttribute{
                            Computed: true,
                        },
                        "customer_id": schema.Int64Attribute{
                            Computed: true,
                        },
                        "route_shader_id": schema.Int64Attribute{
                            Computed: true,
                        },
                    },
                },
            },
        },
    }
}

func (d *buyersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
    if req.ProviderData == nil {
        return
    }
    d.client = req.ProviderData.(*Client)
}

type BuyerModel struct {
    Id              types.Int64  `tfsdk:"id"`
    Name            types.String `tfsdk:"name"`
    PublicKeyBase64 types.String `tfsdk:"public_key_base64"`
    CustomerId      types.Int64  `tfsdk:"customer_id"`
    RouteShaderId   types.Int64  `tfsdk:"route_shader_id"`
}

type BuyerData struct {
    BuyerId         uint64 `json:"buyer_id"`
    BuyerName       string `json:"buyer_name"`
    PublicKeyBase64 string `json:"public_key_base64"`
    CustomerId      uint64 `json:"customer_id"`
    RouteShaderId   uint64 `json:"route_shader_id"`
}

type BuyersResponse struct {
    Buyers []BuyerData `json:"buyers"`
    Error  string      `json:"error"`
}

func (d *buyersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

    buyersResponse := BuyersResponse{}
    
    err := d.client.GetJSON(ctx, "admin/buyers", &buyersResponse)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to get networknext buyers",
            "An error occurred when calling the networknext API to get buyers. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if buyersResponse.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to get networknext buyers",
            "An error occurred when calling the networknext API to get buyers. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+buyersResponse.Error,
        )
        return
    }

    var state buyersDataSourceModel

    for i := range buyersResponse.Buyers {

        buyerState := BuyerModel{
            Id:              types.Int64Value(int64(buyersResponse.Buyers[i].BuyerId)),
            Name:            types.StringValue(buyersResponse.Buyers[i].BuyerName),
            PublicKeyBase64: types.StringValue(buyersResponse.Buyers[i].PublicKeyBase64),
            CustomerId:      types.Int64Value(int64(buyersResponse.Buyers[i].CustomerId)),
            RouteShaderId:   types.Int64Value(int64(buyersResponse.Buyers[i].RouteShaderId)),
        }

        state.Buyers = append(state.Buyers, buyerState)
    }

    diags := resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}
