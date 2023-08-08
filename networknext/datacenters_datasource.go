package networknext

import (
    "context"

    "github.com/hashicorp/terraform-plugin-framework/datasource"
)

var (
    _ datasource.DataSource              = &datacentersDataSource{}
    _ datasource.DataSourceWithConfigure = &datacentersDataSource{}
)

func NewDatacentersDataSource() datasource.DataSource {
    return &datacentersDataSource{}
}

type datacentersDataSource struct {
    client *Client
}

func (d *datacentersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_datacenters"
}

func (d *datacentersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
    resp.Schema = DatacentersSchema()
}

func (d *datacentersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
    if req.ProviderData == nil {
        return
    }
    d.client = req.ProviderData.(*Client)
}

func (d *datacentersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

    datacentersResponse := ReadDatacentersResponse{}
    
    err := d.client.GetJSON(ctx, "admin/datacenters", &datacentersResponse)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to get datacenters",
            "An unexpected error occurred when calling the network next API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if datacentersResponse.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to get datacenters",
            "An error occurred when calling the network next API to get datacenters. "+
                "Network Next Client Error: "+datacentersResponse.Error,
        )
        return
    }

    var state DatacentersModel

    for i := range datacentersResponse.Datacenters {
        var datacenterState DatacenterModel
        DatacenterDataToModel(&datacentersResponse.Datacenters[i], &datacenterState)
        state.Datacenters = append(state.Datacenters, datacenterState)
    }

    diags := resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}
