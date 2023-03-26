package networknext

import (
    "context"

    "github.com/hashicorp/terraform-plugin-framework/datasource"
    "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
    "github.com/hashicorp/terraform-plugin-framework/types"
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
    resp.Schema = schema.Schema{
        Attributes: map[string]schema.Attribute{
            "datacenters": schema.ListNestedAttribute{
                Computed: true,
                NestedObject: schema.NestedAttributeObject{
                    Attributes: map[string]schema.Attribute{
                        "id": schema.Int64Attribute{
                            Computed: true,
                        },
                        "name": schema.StringAttribute{
                            Computed: true,
                        },
                        "native_name": schema.StringAttribute{
                            Computed: true,
                        },
                        "longitude": schema.Float64Attribute{
                            Computed: true,
                        },
                        "latitude": schema.Float64Attribute{
                            Computed: true,
                        },
                        "seller_id": schema.Int64Attribute{
                            Computed: true,
                        },
                        "notes": schema.StringAttribute{
                            Computed: true,
                        },
                    },
                },
            },
        },
    }
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
            "Unable to get networknext datacenters",
            "An error occurred when calling the networknext API to get datacenters. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if datacentersResponse.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to get networknext datacenters",
            "An error occurred when calling the networknext API to get datacenters. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+datacentersResponse.Error,
        )
        return
    }

    var state DatacentersModel

    for i := range datacentersResponse.Datacenters {

        datacenterState := DatacenterModel{
            Id:         types.Int64Value(int64(datacentersResponse.Datacenters[i].DatacenterId)),
            Name:       types.StringValue(datacentersResponse.Datacenters[i].DatacenterName),
            NativeName: types.StringValue(datacentersResponse.Datacenters[i].NativeName),
            Latitude:   types.Float64Value(float64(datacentersResponse.Datacenters[i].Latitude)),
            Longitude:  types.Float64Value(float64(datacentersResponse.Datacenters[i].Longitude)),
            SellerId: types.Int64Value(int64(datacentersResponse.Datacenters[i].SellerId)),
        }

        state.Datacenters = append(state.Datacenters, datacenterState)
    }

    diags := resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}
