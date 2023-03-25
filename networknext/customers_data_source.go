package networknext

import (
    "context"

    "github.com/hashicorp/terraform-plugin-framework/datasource"
    "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
    "github.com/hashicorp/terraform-plugin-framework/types"
)

var (
    _ datasource.DataSource              = &customersDataSource{}
    _ datasource.DataSourceWithConfigure = &customersDataSource{}
)

func NewCustomersDataSource() datasource.DataSource {
    return &customersDataSource{}
}

type customersDataSource struct {
    client *Client
}

type customersDataSourceModel struct {
    Customers []CustomerModel `tfsdk:"customers"`
}

func (d *customersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_customers"
}

func (d *customersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
    resp.Schema = schema.Schema{
        Attributes: map[string]schema.Attribute{
            "customers": schema.ListNestedAttribute{
                Computed: true,
                NestedObject: schema.NestedAttributeObject{
                    Attributes: map[string]schema.Attribute{
                        "id": schema.Int64Attribute{
                            Computed: true,
                        },
                        "name": schema.StringAttribute{
                            Computed: true,
                        },
                        "code": schema.StringAttribute{
                            Computed: true,
                        },
                        "live": schema.BoolAttribute{
                            Computed: true,
                        },
                        "debug": schema.BoolAttribute{
                            Computed: true,
                        },
                    },
                },
            },
        },
    }
}

func (d *customersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
    if req.ProviderData == nil {
        return
    }
    d.client = req.ProviderData.(*Client)
}

type CustomerModel struct {
    Id    types.Int64  `tfsdk:"id"`
    Name  types.String `tfsdk:"name"`
    Code  types.String `tfsdk:"code"`
    Live  types.Bool   `tfsdk:"live"`
    Debug types.Bool   `tfsdk:"debug"`
}

type CustomerData struct {
    CustomerId   uint64 `json:"customer_id"`
    CustomerName string `json:"customer_name"`
    CustomerCode string `json:"customer_code"`
    Live         bool   `json:"live"`
    Debug        bool   `json:"debug"`
}

type CustomersResponse struct {
    Customers []CustomerData `json:"customers"`
    Error     string         `json:"error"`
}

func (d *customersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

    customersResponse := CustomersResponse{}
    
    err := d.client.GetJSON("admin/customers", &customersResponse)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to get networknext customers",
            "An error occurred when calling the networknext API to get customers. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if customersResponse.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to get networknext customers",
            "An error occurred when calling the networknext API to get customers. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+customersResponse.Error,
        )
        return
    }

    var state customersDataSourceModel

    for i := range customersResponse.Customers {

        customerState := CustomerModel{
            Id:          types.Int64Value(int64(customersResponse.Customers[i].CustomerId)),
            Name:        types.StringValue(customersResponse.Customers[i].CustomerName),
            Code:        types.StringValue(customersResponse.Customers[i].CustomerCode),
            Live:        types.BoolValue(customersResponse.Customers[i].Live),
            Debug:       types.BoolValue(customersResponse.Customers[i].Debug),
        }

        state.Customers = append(state.Customers, customerState)
    }

    diags := resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}