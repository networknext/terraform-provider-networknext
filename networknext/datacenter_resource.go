package networknext

import (
    "context"
    "fmt"
    
    "github.com/hashicorp/terraform-plugin-framework/path"
    "github.com/hashicorp/terraform-plugin-framework/resource"
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
    resp.Schema = DatacenterSchema()
}

func (r *datacenterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

    var plan DatacenterModel
    diags := req.Plan.Get(ctx, &plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    var data DatacenterData
    DatacenterModelToData(&plan, &data)

    var response CreateDatacenterResponse
    
    err := r.client.Create(ctx, "admin/create_datacenter", &data, &response)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to create networknext datacenter",
            "An unexpected error occurred when calling the networknext API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to create networknext datacenter",
            "The networknext API returned an error: "+response.Error,
        )
        return
    }

    plan.Id = types.Int64Value(int64(response.Datacenter.DatacenterId))

    diags = resp.State.Set(ctx, plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

func (r *datacenterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

    var state DatacenterModel
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
            "An unexpected error occurred when calling the networknext API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to read networknext datacenter",
            "The networknext API returned an error: "+response.Error,
        )
        return
    }

    data := &response.Datacenter
    DatacenterDataToModel(data, &state)

    diags = resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

func (r *datacenterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

    var plan DatacenterModel
    diags := req.Plan.Get(ctx, &plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    var data DatacenterData
    DatacenterModelToData(&plan, &data)

    var response UpdateDatacenterResponse
    
    err := r.client.Update(ctx, "admin/update_datacenter", &data, &response)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to update networknext datacenter",
            "An unexpected error occurred when calling the networknext API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to update networknext datacenter",
            "The networknext API returned an error: "+response.Error,
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

    var state DatacenterModel
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    id := state.Id.ValueInt64()

    var response UpdateDatacenterResponse

    err := r.client.Delete(ctx, "admin/delete_datacenter", uint64(id), &response)

    if err != nil {
        resp.Diagnostics.AddError(
            "Error deleting networknext datacenter",
            "Could not delete datacenter, unexpected error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to delete networknext datacenter",
            "The networknext API returned an error: "+response.Error,
        )
        return
    }
}

func (r *datacenterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
    resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
