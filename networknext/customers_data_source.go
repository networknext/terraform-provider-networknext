package networknext

import (
    "context"

    "github.com/hashicorp/terraform-plugin-framework/datasource"
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

func (d *customersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_customers"
}

func (d *customersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
    resp.Schema = CustomersSchema()
}

func (d *customersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
    if req.ProviderData == nil {
        return
    }
    d.client = req.ProviderData.(*Client)
}

func (d *customersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

    customersResponse := ReadCustomersResponse{}
    
    err := d.client.GetJSON(ctx, "admin/customers", &customersResponse)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to get networknext customers",
            "An unexpected error occurred when calling the networknext API. "+
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

    var state CustomersModel

    for i := range customersResponse.Customers {
        var customerState CustomerModel
        CustomerDataToModel(&customersResponse.Customers[i], &customerState)
        state.Customers = append(state.Customers, customerState)
    }

    diags := resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}
