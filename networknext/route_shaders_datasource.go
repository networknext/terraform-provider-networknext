package networknext

import (
    "context"

    "github.com/hashicorp/terraform-plugin-framework/datasource"
)

var (
    _ datasource.DataSource              = &routeShadersDataSource{}
    _ datasource.DataSourceWithConfigure = &routeShadersDataSource{}
)

func NewRouteShadersDataSource() datasource.DataSource {
    return &routeShadersDataSource{}
}

type routeShadersDataSource struct {
    client *Client
}

func (d *routeShadersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_route_shaders"
}

func (d *routeShadersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
    resp.Schema = RouteShadersSchema()
}

func (d *routeShadersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
    if req.ProviderData == nil {
        return
    }
    d.client = req.ProviderData.(*Client)
}

func (d *routeShadersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

    routeShadersResponse := ReadRouteShadersResponse{}
    
    err := d.client.GetJSON(ctx, "admin/route_shaders", &routeShadersResponse)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to get route shaders",
            "An unexpected error occurred when calling the network next API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if routeShadersResponse.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to get route shaders",
            "The network next API returned an error: "+routeShadersResponse.Error,
        )
        return
    }

    var state RouteShadersModel

    for i := range routeShadersResponse.RouteShaders {
        var routeShaderState RouteShaderModel
        RouteShaderDataToModel(&routeShadersResponse.RouteShaders[i], &routeShaderState)
        state.RouteShaders = append(state.RouteShaders, routeShaderState)
    }

    diags := resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}
