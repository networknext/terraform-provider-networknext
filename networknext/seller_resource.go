package networknext

import (
    "context"
    "fmt"
    
    "github.com/hashicorp/terraform-plugin-framework/path"
    "github.com/hashicorp/terraform-plugin-framework/resource"
    "github.com/hashicorp/terraform-plugin-framework/types"
)

var (
    _ resource.Resource                = &sellerResource{}
    _ resource.ResourceWithConfigure   = &sellerResource{}
    _ resource.ResourceWithImportState = &sellerResource{}
)

func NewSellerResource() resource.Resource {
    return &sellerResource{}
}

type sellerResource struct {
    client *Client
}

func (r *sellerResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
    if req.ProviderData == nil {
        return
    }
    r.client = req.ProviderData.(*Client)
}

func (r *sellerResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_seller"
}

func (r *sellerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = SellerSchema()
}

func (r *sellerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

    var plan SellerModel
    diags := req.Plan.Get(ctx, &plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    var data SellerData
    SellerModelToData(&plan, &data)

    var response CreateSellerResponse

    err := r.client.Create(ctx, "admin/create_seller", &data, &response)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to create seller",
            "An unexpected error occurred when calling the network next API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to create seller",
            "The network next API returned an error: "+response.Error,
        )
        return
    }

    plan.Id = types.Int64Value(int64(response.Seller.SellerId))

    diags = resp.State.Set(ctx, plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

type ReadSellerResponse struct {
    Seller SellerData `json:"seller"`
    Error  string     `json:"error"`
}

func (r *sellerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

    var state SellerModel
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    response := ReadSellerResponse{}

    err := r.client.GetJSON(ctx, fmt.Sprintf("admin/seller/%d", int64(state.Id.ValueInt64())), &response)

    if err != nil {        
        resp.Diagnostics.AddError(
            "Unable to read seller",
            "An unexpected error occurred when calling the network next API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to read seller",
            "The network next API returned an error while trying to read a seller. "+
                "Network Next Client Error: "+response.Error,
        )
        return
    }

    data := &response.Seller
    SellerDataToModel(data, &state)

    diags = resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

func (r *sellerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

    var plan SellerModel
    diags := req.Plan.Get(ctx, &plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    var data SellerData
    SellerModelToData(&plan, &data)

    var response UpdateSellerResponse
    
    err := r.client.Update(ctx, "admin/update_seller", &data, &response)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to update seller",
            "An unexpected error occurred when calling the network next API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to update seller",
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

func (r *sellerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

    var state SellerModel
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    id := state.Id.ValueInt64()

    var response UpdateSellerResponse

    err := r.client.Delete(ctx, fmt.Sprintf("admin/delete_seller/%d", uint64(id)), &response)

    if err != nil {
        resp.Diagnostics.AddError(
            "Error deleting seller",
            "Could not delete seller, unexpected error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to delete seller",
            "The network next API returned an error: "+response.Error,
        )
        return
    }
}

func (r *sellerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
    resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
