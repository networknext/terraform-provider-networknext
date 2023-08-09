package networknext

import (
    "context"
    "fmt"
    
    "github.com/hashicorp/terraform-plugin-framework/path"
    "github.com/hashicorp/terraform-plugin-framework/resource"
)

var (
    _ resource.Resource                = &relayKeypairResource{}
    _ resource.ResourceWithConfigure   = &relayKeypairResource{}
    _ resource.ResourceWithImportState = &relayKeypairResource{}
)

func NewRelayKeypairResource() resource.Resource {
    return &relayKeypairResource{}
}

type relayKeypairResource struct {
    client *Client
}

func (r *relayKeypairResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
    if req.ProviderData == nil {
        return
    }
    r.client = req.ProviderData.(*Client)
}

func (r *relayKeypairResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_relay_keypair"
}

func (r *relayKeypairResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = RelayKeypairSchema()
}

func (r *relayKeypairResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

    var plan RelayKeypairModel
    diags := req.Plan.Get(ctx, &plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    var data RelayKeypairData
    RelayKeypairModelToData(&plan, &data)

    var response CreateRelayKeypairResponse
    
    err := r.client.Create(ctx, "admin/create_relay_keypair", &data, &response)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to create relay keypair",
            "An unexpected error occurred when calling the network next API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to create relay keypair",
            "The network next API returned an error: "+response.Error,
        )
        return
    }

    RelayKeypairDataToModel(&response.RelayKeypair, &plan)

    diags = resp.State.Set(ctx, plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

func (r *relayKeypairResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

    var state RelayKeypairModel
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    response := ReadRelayKeypairResponse{}

    err := r.client.GetJSON(ctx, fmt.Sprintf("admin/relay_keypair/%x", int64(state.Id.ValueInt64())), &response)

    if err != nil {        
        resp.Diagnostics.AddError(
            "Unable to read relay keypair",
            "An unexpected error occurred when calling the network next API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to read relay keypair",
            "The network next API returned an error: "+response.Error,
        )
        return
    }

    data := &response.RelayKeypair
    RelayKeypairDataToModel(data, &state)

    diags = resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

func (r *relayKeypairResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

    var plan RelayKeypairModel
    diags := req.Plan.Get(ctx, &plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    var data RelayKeypairData
    RelayKeypairModelToData(&plan, &data)

    var response UpdateRelayKeypairResponse
    
    err := r.client.Update(ctx, "admin/update_relay_keypair", &data, &response)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to update relay keypair",
            "An unexpected error occurred when calling the network next API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to update relay keypair",
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

func (r *relayKeypairResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

    var state RelayKeypairModel
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    id := state.Id.ValueInt64()

    var response UpdateRelayKeypairResponse

    err := r.client.Delete(ctx, fmt.Sprintf("admin/delete_relay_keypair/%x", uint64(id)), &response)

    if err != nil {
        resp.Diagnostics.AddError(
            "Error deleting relay keypair",
            "Could not delete relay keypair, unexpected error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to delete relay keypair",
            "The network next API returned an error: "+response.Error,
        )
        return
    }
}

func (r *relayKeypairResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
    resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
