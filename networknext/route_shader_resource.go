package networknext

import (
    "context"
    "fmt"
    
    "github.com/hashicorp/terraform-plugin-framework/path"
    "github.com/hashicorp/terraform-plugin-framework/resource"
)

var (
    _ resource.Resource                = &routeShaderResource{}
    _ resource.ResourceWithConfigure   = &routeShaderResource{}
    _ resource.ResourceWithImportState = &routeShaderResource{}
)

func NewRouteShaderResource() resource.Resource {
    return &routeShaderResource{}
}

type routeShaderResource struct {
    client *Client
}

func (r *routeShaderResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
    if req.ProviderData == nil {
        return
    }
    r.client = req.ProviderData.(*Client)
}

func (r *routeShaderResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_route_shader"
}

func (r *routeShaderResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = RouteShaderSchema()
}

func (r *routeShaderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

    var plan RouteShaderModel
    diags := req.Plan.Get(ctx, &plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    var data RouteShaderData
    RouteShaderModelToData(&plan, &data)

    var response CreateRouteShaderResponse
    
    err := r.client.Create(ctx, "admin/create_route_shader", &data, &response)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to create route shader",
            "An unexpected error occurred when calling the network next API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to create route shader",
            "The network next API returned an error: "+response.Error,
        )
        return
    }

    RouteShaderDataToModel(&response.RouteShader, &plan)

    diags = resp.State.Set(ctx, plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

func (r *routeShaderResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

    var state RouteShaderModel
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    response := ReadRouteShaderResponse{}

    err := r.client.GetJSON(ctx, fmt.Sprintf("admin/route_shader/%d", int64(state.Id.ValueInt64())), &response)

    if err != nil {        
        resp.Diagnostics.AddError(
            "Unable to read route shader",
            "An unexpected error occurred when calling the network next API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to read route shader",
            "The network next API returned an error: "+response.Error,
        )
        return
    }

    data := &response.RouteShader
    RouteShaderDataToModel(data, &state)

    diags = resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

func (r *routeShaderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

    var plan RouteShaderModel
    diags := req.Plan.Get(ctx, &plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    var data RouteShaderData
    RouteShaderModelToData(&plan, &data)

    var response UpdateRouteShaderResponse
    
    err := r.client.Update(ctx, "admin/update_route_shader", &data, &response)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to update route shader",
            "An unexpected error occurred when calling the network next API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to update route shader",
            "The network next API returned an error: "+response.Error,
        )
        return
    }

    diags = resp.State.Set(ctx, plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

func (r *routeShaderResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

    var state RouteShaderModel
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    id := state.Id.ValueInt64()

    var response UpdateRouteShaderResponse

    err := r.client.Delete(ctx, fmt.Sprintf("admin/delete_route_shader/%d", uint64(id)), &response)

    if err != nil {
        resp.Diagnostics.AddError(
            "Error deleting route shader",
            "Could not delete route shader, unexpected error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to delete route shader",
            "The network next API returned an error: "+response.Error,
        )
        return
    }
}

func (r *routeShaderResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
    resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
