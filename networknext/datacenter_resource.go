package networknext

import (
    "context"
    "fmt"
    
    "github.com/hashicorp/terraform-plugin-framework/path"
    "github.com/hashicorp/terraform-plugin-framework/resource"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
    "github.com/hashicorp/terraform-plugin-framework/types"
)

var (
    _ resource.Resource                = &datacenterResource{}
    _ resource.ResourceWithConfigure   = &datacenterResource{}
    _ resource.ResourceWithImportState = &datacenterResource{}
)

func NewDatacenterResource() resource.Resource {
    return &datacenterResource{}
}

type datacenterResource struct {
    client *Client
}

type datacenterResourceModel struct {
    Id          types.Int64     `tfsdk:"id"`
    Name        types.String    `tfsdk:"name"`
    NativeName  types.String    `tfsdk:"native_name"`
    Latitude    types.Float64   `tfsdk:"latitude"`
    Longitude   types.Float64   `tfsdk:"longitude"`
    SellerId    types.Int64     `tfsdk:"seller_id"`
    Notes       types.String    `tfsdk:"notes"`
}

func (r *datacenterResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
    if req.ProviderData == nil {
        return
    }
    r.client = req.ProviderData.(*Client)
}

func (r *datacenterResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_datacenter"
}

func (r *datacenterResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
            "native_name": schema.StringAttribute{
                Required: true,
            },
            "latitude": schema.Float64Attribute{
                Required: true,
            },
            "longitude": schema.Float64Attribute{
                Required: true,
            },
            "seller_id": schema.Int64Attribute{
                Required: true,
            },
            "notes": schema.StringAttribute{
                Required: true,
            },
        },
    }
}

func (r *datacenterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

    var plan datacenterResourceModel
    diags := req.Plan.Get(ctx, &plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    var data DatacenterData
    data.DatacenterName = plan.Name.ValueString()
    data.NativeName = plan.NativeName.ValueString()
    data.Latitude = float32(plan.Latitude.ValueFloat64())
    data.Longitude = float32(plan.Longitude.ValueFloat64())
    data.SellerId = uint64(plan.SellerId.ValueInt64())
    data.Notes = plan.Notes.ValueString()

    id, err := r.client.Create("admin/create_datacenter", &data)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to create networknext datacenter",
            "An error occurred when calling the networknext API to create a datacenter. "+
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

type ReadDatacenterResponse struct {
    Datacenter DatacenterData `json:"datacenter"`
    Error      string         `json:"error"`
}

func (r *datacenterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

    var state datacenterResourceModel
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    response := ReadDatacenterResponse{}

    err := r.client.GetJSON(ctx, fmt.Sprintf("admin/datacenter/%x", int64(state.Id.ValueInt64())), &response)

    if err != nil {        
        resp.Diagnostics.AddError(
            "Unable to read networknext datacenter",
            "An unexpected error occurred when calling the networknext API to read a datacenter. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to read networknext datacenter",
            "The networknext API returned an error while trying to read a datacenter. "+
                "Network Next Client Error: "+response.Error,
        )
        return
    }

    data := &response.Datacenter

    state.Id = types.Int64Value(int64(data.DatacenterId))
    state.Name = types.StringValue(data.DatacenterName)
    state.NativeName = types.StringValue(data.NativeName)
    state.Latitude = types.Float64Value(float64(data.Latitude))
    state.Longitude = types.Float64Value(float64(data.Longitude))
    state.SellerId = types.Int64Value(int64(data.SellerId))
    state.Notes = types.StringValue(data.Notes)

    diags = resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

func (r *datacenterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

    var plan datacenterResourceModel
    diags := req.Plan.Get(ctx, &plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    var data DatacenterData
    data.DatacenterId = uint64(plan.Id.ValueInt64())
    data.DatacenterName = plan.Name.ValueString()
    data.NativeName = plan.NativeName.ValueString()
    data.Latitude = float32(plan.Latitude.ValueFloat64())
    data.Longitude = float32(plan.Longitude.ValueFloat64())
    data.SellerId = uint64(plan.SellerId.ValueInt64())
    data.Notes = plan.Notes.ValueString()

    err := r.client.Update(ctx, "admin/update_datacenter", &data)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to update networknext datacenter",
            "An error occurred when calling the networknext API to update a datacenter. "+
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

func (r *datacenterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

    var state datacenterResourceModel
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    id := state.Id.ValueInt64()

    err := r.client.Delete(ctx, "admin/delete_datacenter", uint64(id))

    if err != nil {
        resp.Diagnostics.AddError(
            "Error deleting networknext datacenter",
            "Could not delete datacenter, unexpected error: "+err.Error(),
        )
        return
    }
}

func (r *datacenterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
    resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

