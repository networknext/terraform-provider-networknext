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

    id, err := r.client.Create("admin/create_customer", &data)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to create networknext customer",
            "An error occurred when calling the networknext API to create a customer. "+
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
            "An unexpected error occurred when calling the networknext API to read a customer. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to read networknext customer",
            "The networknext API returned an error while trying to read a customer. "+
                "Network Next Client Error: "+response.Error,
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
    
    err := r.client.Update(ctx, "admin/update_customer", &data)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to update networknext customer",
            "An error occurred when calling the networknext API to update a customer. "+
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

func (r *customerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

    var state CustomerModel
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    id := state.Id.ValueInt64()

    err := r.client.Delete(ctx, "admin/delete_customer", uint64(id))

    if err != nil {
        resp.Diagnostics.AddError(
            "Error deleting networknext customer",
            "Could not delete customer, unexpected error: "+err.Error(),
        )
        return
    }
}

func (r *customerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
    resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
