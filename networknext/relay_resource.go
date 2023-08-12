package networknext

import (
    "context"
    "fmt"
    
    "github.com/hashicorp/terraform-plugin-framework/path"
    "github.com/hashicorp/terraform-plugin-framework/resource"
    "github.com/hashicorp/terraform-plugin-framework/types"
)

var (
    _ resource.Resource                = &relayResource{}
    _ resource.ResourceWithConfigure   = &relayResource{}
    _ resource.ResourceWithImportState = &relayResource{}
)

func NewRelayResource() resource.Resource {
    return &relayResource{}
}

type relayResource struct {
    client *Client
}

func (r *relayResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
    if req.ProviderData == nil {
        return
    }
    r.client = req.ProviderData.(*Client)
}

func (r *relayResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_relay"
}

func (r *relayResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = RelaySchema()
}

func (r *relayResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

    var plan RelayModel
    diags := req.Plan.Get(ctx, &plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    var data RelayData
    RelayModelToData(&plan, &data)

    var response CreateRelayResponse
    
    err := r.client.Create(ctx, "admin/create_relay", &data, &response)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to create relay",
            "An unexpected error occurred when calling the network next API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to create relay",
            "The network next API returned an error: "+response.Error,
        )
        return
    }

    plan.Id = types.Int64Value(int64(response.Relay.RelayId))

    diags = resp.State.Set(ctx, plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

func (r *relayResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

    var state RelayModel
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    response := ReadRelayResponse{}

    err := r.client.GetJSON(ctx, fmt.Sprintf("admin/relay/%d", int64(state.Id.ValueInt64())), &response)

    if err != nil {        
        resp.Diagnostics.AddError(
            "Unable to read relay",
            "An unexpected error occurred when calling the network next API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to read relay",
            "The network next API returned an error: "+response.Error,
        )
        return
    }

    data := &response.Relay
    RelayDataToModel(data, &state)

    diags = resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

func (r *relayResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

    var plan RelayModel
    diags := req.Plan.Get(ctx, &plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    var data RelayData
    RelayModelToData(&plan, &data)

    var response UpdateRelayResponse
    
    err := r.client.Update(ctx, "admin/update_relay", &data, &response)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to update relay",
            "An unexpected error occurred when calling the network next API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to update relay",
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

func (r *relayResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

    var state RelayModel
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    id := state.Id.ValueInt64()

    var response UpdateRelayResponse

    err := r.client.Delete(ctx, fmt.Sprintf("admin/delete_relay/%d", uint64(id)), &response)

    if err != nil {
        resp.Diagnostics.AddError(
            "Error deleting relay",
            "An unexpected error occurred when calling the network next API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to delete relay",
            "The network next API returned an error: "+response.Error,
        )
        return
    }
}

func (r *relayResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
    resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
