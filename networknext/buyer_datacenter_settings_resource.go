package networknext

import (
    "context"
    "fmt"
    
    "github.com/hashicorp/terraform-plugin-framework/path"
    "github.com/hashicorp/terraform-plugin-framework/resource"
)

var (
    _ resource.Resource                = &buyerDatacenterSettingsResource{}
    _ resource.ResourceWithConfigure   = &buyerDatacenterSettingsResource{}
    _ resource.ResourceWithImportState = &buyerDatacenterSettingsResource{}
)

func NewBuyerDatacenterSettingsResource() resource.Resource {
    return &buyerDatacenterSettingsResource{}
}

type buyerDatacenterSettingsResource struct {
    client *Client
}

func (r *buyerDatacenterSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
    if req.ProviderData == nil {
        return
    }
    r.client = req.ProviderData.(*Client)
}

func (r *buyerDatacenterSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_buyer_datacenter_settings"
}

func (r *buyerDatacenterSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = BuyerDatacenterSettingsSchema()
}

func (r *buyerDatacenterSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

    var plan BuyerDatacenterSettingsModel
    diags := req.Plan.Get(ctx, &plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    var data BuyerDatacenterSettingsData
    BuyerDatacenterSettingsModelToData(&plan, &data)

    // todo
    /*
    err := r.client.Create("admin/create_buyer_datacenter_settings", &data, &response)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to create networknext buyer datacenter settings",
            "An unexpected error occurred when calling the networknext API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }
    */

    diags = resp.State.Set(ctx, plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

func (r *buyerDatacenterSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

    var state BuyerDatacenterSettingsModel
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    response := ReadBuyerDatacenterSettingsResponse{}

    err := r.client.GetJSON(ctx, fmt.Sprintf("admin/buyer_datacenter_settings/%x/%x", int64(state.BuyerId.ValueInt64()), int64(state.DatacenterId.ValueInt64())), &response)

    if err != nil {        
        resp.Diagnostics.AddError(
            "Unable to read networknext buyer datacenter settings",
            "An unexpected error occurred when calling the networknext API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to read networknext buyer datacenter settings",
            "The networknext API returned an error while trying to read a buyer datacenter settings. "+
                "Network Next Client Error: "+response.Error,
        )
        return
    }

    data := &response.Settings
    BuyerDatacenterSettingsDataToModel(data, &state)

    diags = resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

func (r *buyerDatacenterSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

    var plan BuyerDatacenterSettingsModel
    diags := req.Plan.Get(ctx, &plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    var data BuyerDatacenterSettingsData
    BuyerDatacenterSettingsModelToData(&plan, &data)
    
    // todo
    /*
    err := r.client.Update(ctx, "admin/update_buyer_datacenter_settings", &data)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to update networknext buyer datacenter settings",
            "An unexpected error occurred when calling the networknext API. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }
    */

    // todo: we need a real error message here

    diags = resp.State.Set(ctx, plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

func (r *buyerDatacenterSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

    var state BuyerDatacenterSettingsModel
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    // todo: wait, we need an id...

    /*
    id := state.Id.ValueInt64()

    err := r.client.Delete(ctx, "admin/delete_buyer_datacenter_settings", uint64(id))

    if err != nil {
        resp.Diagnostics.AddError(
            "Error deleting networknext buyer datacenter settings",
            "Could not delete buyer datacenter settings, unexpected error: "+err.Error(),
        )
        return
    }
    */
}

func (r *buyerDatacenterSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
    resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
