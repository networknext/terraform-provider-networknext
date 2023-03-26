package networknext

import (
    "context"
    "fmt"
    
    "github.com/hashicorp/terraform-plugin-framework/path"
    "github.com/hashicorp/terraform-plugin-framework/resource"
    "github.com/hashicorp/terraform-plugin-framework/types"
)

var (
    _ resource.Resource                = &buyerResource{}
    _ resource.ResourceWithConfigure   = &buyerResource{}
    _ resource.ResourceWithImportState = &buyerResource{}
)

func NewBuyerResource() resource.Resource {
    return &buyerResource{}
}

type buyerResource struct {
    client *Client
}

func (r *buyerResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
    if req.ProviderData == nil {
        return
    }
    r.client = req.ProviderData.(*Client)
}

func (r *buyerResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_buyer"
}

func (r *buyerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = BuyerSchema()
}

func (r *buyerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

    var plan BuyerModel
    diags := req.Plan.Get(ctx, &plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    var data BuyerData
    BuyerModelToData(&plan, &data)

    id, err := r.client.Create("admin/create_buyer", &data)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to create networknext buyer",
            "An error occurred when calling the networknext API to create a buyer. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    // todo: we need to return a better error here, not just an id value

    plan.Id = types.Int64Value(int64(id))

    diags = resp.State.Set(ctx, plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

func (r *buyerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

    var state BuyerModel
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    response := ReadBuyerResponse{}

    err := r.client.GetJSON(ctx, fmt.Sprintf("admin/buyer/%x", int64(state.Id.ValueInt64())), &response)

    if err != nil {        
        resp.Diagnostics.AddError(
            "Unable to read networknext buyer",
            "An unexpected error occurred when calling the networknext API to read a buyer. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to read networknext buyer",
            "The networknext API returned an error while trying to read a buyer. "+
                "Network Next Client Error: "+response.Error,
        )
        return
    }

    data := &response.Buyer
    BuyerDataToModel(data, &state)

    diags = resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

func (r *buyerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

    var plan BuyerModel
    diags := req.Plan.Get(ctx, &plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    var data BuyerData
    BuyerModelToData(&plan, &data)
    
    err := r.client.Update(ctx, "admin/update_buyer", &data)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to update networknext buyer",
            "An error occurred when calling the networknext API to update a buyer. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    diags = resp.State.Set(ctx, plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

func (r *buyerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

    var state BuyerModel
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    id := state.Id.ValueInt64()

    err := r.client.Delete(ctx, "admin/delete_buyer", uint64(id))

    if err != nil {
        resp.Diagnostics.AddError(
            "Error deleting networknext buyer",
            "Could not delete buyer, unexpected error: "+err.Error(),
        )
        return
    }
}

func (r *buyerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
    resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
