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

    id, err := r.client.Create("admin/create_relay", &data)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to create networknext relay",
            "An unexpected error occurred when calling the networknext API to create a relay. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    // todo: we need a real error here

    plan.Id = types.Int64Value(int64(id))

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

    err := r.client.GetJSON(ctx, fmt.Sprintf("admin/relay/%x", int64(state.Id.ValueInt64())), &response)

    if err != nil {        
        resp.Diagnostics.AddError(
            "Unable to read networknext relay",
            "An unexpected error occurred when calling the networknext API to read a relay. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to read networknext relay",
            "The networknext API returned an error while trying to read a relay. "+
                "Network Next Client Error: "+response.Error,
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

    err := r.client.Update(ctx, "admin/update_relay", &data)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to update networknext relay",
            "An unexpected error occurred when calling the networknext API to update a relay. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    // todo: we need a real error message here

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

    err := r.client.Delete(ctx, "admin/delete_relay", uint64(id))

    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to delete networknext relay",
            "An unexpected error occurred when calling the networknext API to delete a relay. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }
}

func (r *relayResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
    resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
