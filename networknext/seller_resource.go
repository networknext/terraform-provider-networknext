package networknext

import (
    "context"
    "fmt"
    
    "github.com/hashicorp/terraform-plugin-framework/path"
    "github.com/hashicorp/terraform-plugin-framework/resource"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
    // "github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
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

type sellerResourceModel struct {
    Id          types.Int64     `tfsdk:"id"`
    Name        types.String    `tfsdk:"name"`
    CustomerId  types.Int64     `tfsdk:"customer_id"`
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
    resp.Schema = schema.Schema{
        Attributes: map[string]schema.Attribute{
            "id": schema.Int64Attribute{
                Computed: true,
                PlanModifiers: []planmodifier.Int64{
                    int64planmodifier.UseStateForUnknown(),
                },
            },
            "name": schema.StringAttribute{
                Required: true,
            },
            "customer_id": schema.Int64Attribute{
                Required: true,
                // Optional: true,
                // Default: int64default.StaticValue(0),
            },
        },
    }
}

func (r *sellerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

    var plan sellerResourceModel
    diags := req.Plan.Get(ctx, &plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    var data SellerData
    data.SellerName = plan.Name.ValueString()
    data.CustomerId = uint64(plan.CustomerId.ValueInt64())

    id, err := r.client.Create("admin/create_seller", &data)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to create networknext seller",
            "An error occurred when calling the networknext API to create a seller. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    plan.Id = types.Int64Value(int64(id))

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

    var state sellerResourceModel
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    response := ReadSellerResponse{}

    err := r.client.GetJSON(ctx, fmt.Sprintf("admin/seller/%x", int64(state.Id.ValueInt64())), &response)

    if err != nil {        
        resp.Diagnostics.AddError(
            "Unable to read networknext seller",
            "An unexpected error occurred when calling the networknext API to read a seller. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to read networknext seller",
            "The networknext API returned an error while trying to read a seller. "+
                "Network Next Client Error: "+response.Error,
        )
        return
    }

    data := &response.Seller

    state.Id = types.Int64Value(int64(data.SellerId))
    state.Name = types.StringValue(data.SellerName)
    state.CustomerId = types.Int64Value(int64(data.CustomerId))

    diags = resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

func (r *sellerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

    var plan sellerResourceModel
    diags := req.Plan.Get(ctx, &plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    var data SellerData
    data.SellerId = uint64(plan.Id.ValueInt64())
    data.SellerName = plan.Name.ValueString()
    data.CustomerId = uint64(plan.CustomerId.ValueInt64())

    err := r.client.Update(ctx, "admin/update_seller", &data)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to update networknext seller",
            "An error occurred when calling the networknext API to update a seller. "+
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

func (r *sellerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

    var state sellerResourceModel
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    id := state.Id.ValueInt64()

    err := r.client.Delete(ctx, "admin/delete_seller", uint64(id))

    if err != nil {
        resp.Diagnostics.AddError(
            "Error deleting networknext seller",
            "Could not delete seller, unexpected error: "+err.Error(),
        )
        return
    }
}

func (r *sellerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
    resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
