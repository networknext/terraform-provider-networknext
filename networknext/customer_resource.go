package networknext

import (
    "context"
    "fmt"
    
    "github.com/hashicorp/terraform-plugin-framework/path"
    "github.com/hashicorp/terraform-plugin-framework/resource"
    "github.com/hashicorp/terraform-plugin-framework/types"
)

var (
    _ resource.Resource                = &customerResource{}
    _ resource.ResourceWithConfigure   = &customerResource{}
    _ resource.ResourceWithImportState = &customerResource{}
)

func NewCustomerResource() resource.Resource {
    return &customerResource{}
}

type customerResource struct {
    client *Client
}

func (r *customerResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
    if req.ProviderData == nil {
        return
    }
    r.client = req.ProviderData.(*Client)
}

func (r *customerResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_customer"
}

func (r *customerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = CustomerSchema()
}

func (r *customerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

    var plan CustomerModel
    diags := req.Plan.Get(ctx, &plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    var data CustomerData
    CustomerModelToData(&plan, &data)

    var response CreateCustomerResponse
    
    err := r.client.Create(ctx, "admin/create_customer", &data, &response)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to create networknext customer",
            "An unexpected error occurred when calling the networknext API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to create networknext customer",
            "The networknext API returned an error: "+response.Error,
        )
        return
    }

    plan.Id = types.Int64Value(int64(response.Customer.CustomerId))

    diags = resp.State.Set(ctx, plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

func (r *customerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

    var state CustomerModel
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    response := ReadCustomerResponse{}

    err := r.client.GetJSON(ctx, fmt.Sprintf("admin/customer/%x", int64(state.Id.ValueInt64())), &response)

    if err != nil {        
        resp.Diagnostics.AddError(
            "Unable to read networknext customer",
            "An unexpected error occurred when calling the networknext API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to read networknext customer",
            "The networknext API returned an error: "+response.Error,
        )
        return
    }

    data := &response.Customer
    CustomerDataToModel(data, &state)

    diags = resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

func (r *customerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

    var plan CustomerModel
    diags := req.Plan.Get(ctx, &plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    var data CustomerData
    CustomerModelToData(&plan, &data)

    var response UpdateCustomerResponse
    
    err := r.client.Update(ctx, "admin/update_customer", &data, &response)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to update networknext customer",
            "An unexpected error occurred when calling the networknext API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to update networknext customer",
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

func (r *customerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

    var state CustomerModel
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    id := state.Id.ValueInt64()

    var response UpdateCustomerResponse

    err := r.client.Delete(ctx, "admin/delete_customer", uint64(id), &response)

    if err != nil {
        resp.Diagnostics.AddError(
            "Error deleting networknext customer",
            "Could not delete customer, unexpected error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to delete networknext customer",
            "The networknext API returned an error: "+response.Error,
        )
        return
    }
}

func (r *customerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
    resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
