package networknext

import (
    "context"
    "fmt"
    
    "github.com/hashicorp/terraform-plugin-framework/path"
    "github.com/hashicorp/terraform-plugin-framework/resource"
)

var (
    _ resource.Resource                = &buyerKeypairResource{}
    _ resource.ResourceWithConfigure   = &buyerKeypairResource{}
    _ resource.ResourceWithImportState = &buyerKeypairResource{}
)

func NewBuyerKeypairResource() resource.Resource {
    return &buyerKeypairResource{}
}

type buyerKeypairResource struct {
    client *Client
}

func (r *buyerKeypairResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
    if req.ProviderData == nil {
        return
    }
    r.client = req.ProviderData.(*Client)
}

func (r *buyerKeypairResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_buyer_keypair"
}

func (r *buyerKeypairResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = BuyerKeypairSchema()
}

func (r *buyerKeypairResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

    var plan BuyerKeypairModel
    diags := req.Plan.Get(ctx, &plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    var data BuyerKeypairData
    BuyerKeypairModelToData(&plan, &data)

    var response CreateBuyerKeypairResponse
    
    err := r.client.Create(ctx, "admin/create_buyer_keypair", &data, &response)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to create buyer keypair",
            "An unexpected error occurred when calling the network next API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to create buyer keypair",
            "The network next API returned an error: "+response.Error,
        )
        return
    }

    BuyerKeypairDataToModel(&response.BuyerKeypair, &plan)

    diags = resp.State.Set(ctx, plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

func (r *buyerKeypairResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

    var state BuyerKeypairModel
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    response := ReadBuyerKeypairResponse{}

    err := r.client.GetJSON(ctx, fmt.Sprintf("admin/buyer_keypair/%x", int64(state.Id.ValueInt64())), &response)

    if err != nil {        
        resp.Diagnostics.AddError(
            "Unable to read buyer keypair",
            "An unexpected error occurred when calling the network next API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to read buyer keypair",
            "The network next API returned an error: "+response.Error,
        )
        return
    }

    data := &response.BuyerKeypair
    BuyerKeypairDataToModel(data, &state)

    diags = resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

func (r *buyerKeypairResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

    var plan BuyerKeypairModel
    diags := req.Plan.Get(ctx, &plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    var data BuyerKeypairData
    BuyerKeypairModelToData(&plan, &data)

    var response UpdateBuyerKeypairResponse
    
    err := r.client.Update(ctx, "admin/update_buyer_keypair", &data, &response)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to update buyer keypair",
            "An unexpected error occurred when calling the network next API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to update buyer keypair",
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

func (r *buyerKeypairResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

    var state BuyerKeypairModel
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    id := state.Id.ValueInt64()

    var response UpdateBuyerKeypairResponse

    err := r.client.Delete(ctx, fmt.Sprintf("admin/delete_buyer_keypair/%x", uint64(id)), &response)

    if err != nil {
        resp.Diagnostics.AddError(
            "Error deleting buyer keypair",
            "An unexpected error occurred when calling the network next API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to delete buyer keypair",
            "The network next API returned an error: "+response.Error,
        )
        return
    }
}

func (r *buyerKeypairResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
    resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
