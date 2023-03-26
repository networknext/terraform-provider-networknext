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
    _ resource.Resource                = &relayResource{}
    _ resource.ResourceWithConfigure   = &relayResource{}
    _ resource.ResourceWithImportState = &relayResource{}
)

func NewRelayResource() resource.Resource {
    return &relayResource{}
}

type relayResource struct {
    client *Client
}

func (r *relayResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
    if req.ProviderData == nil {
        return
    }
    r.client = req.ProviderData.(*Client)
}

func (r *relayResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_relay"
}

func (r *relayResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
            "datacenter_id": schema.Int64Attribute{
                Required: true,
            },
            "public_ip": schema.StringAttribute{
                Required: true,
            },
            "public_port": schema.Int64Attribute{
                Required: true,
            },
            "internal_ip": schema.StringAttribute{
                Required: true,
            },
            "internal_port": schema.Int64Attribute{
                Required: true,
            },
            "internal_group": schema.StringAttribute{
                Required: true,
            },
            "ssh_ip": schema.StringAttribute{
                Required: true,
            },
            "ssh_port": schema.Int64Attribute{
                Required: true,
            },
            "ssh_user": schema.StringAttribute{
                Required: true,
            },
            "private_key_base64": schema.StringAttribute{
                Required: true,
            },
            "public_key_base64": schema.StringAttribute{
                Required: true,
            },
            "version": schema.StringAttribute{
                Required: true,
            },
            "mrc": schema.Int64Attribute{
                Required: true,
            },
            "port_speed": schema.Int64Attribute{
                Required: true,
            },
            "max_sessions": schema.Int64Attribute{
                Required: true,
            },
            "notes": schema.StringAttribute{
                Required: true,
            },
        },
    }
}

func (r *relayResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

    var plan RelayModel
    diags := req.Plan.Get(ctx, &plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    var data RelayData
    data.RelayName = plan.Name.ValueString()
    data.DatacenterId = uint64(plan.DatacenterId.ValueInt64())
    data.PublicIP = plan.PublicIP.ValueString()
    data.PublicPort = int(plan.PublicPort.ValueInt64())
    data.InternalIP = plan.InternalIP.ValueString()
    data.InternalPort = int(plan.InternalPort.ValueInt64())
    data.InternalGroup = plan.InternalGroup.ValueString()
    data.SSH_IP = plan.SSH_IP.ValueString()
    data.SSH_Port = int(plan.SSH_Port.ValueInt64())
    data.SSH_User = plan.SSH_User.ValueString()
    data.PublicKeyBase64 = plan.PublicKeyBase64.ValueString()
    data.PrivateKeyBase64 = plan.PrivateKeyBase64.ValueString()
    data.Version = plan.Version.ValueString()
    data.MRC = int(plan.MRC.ValueInt64())
    data.PortSpeed = int(plan.PortSpeed.ValueInt64())
    data.MaxSessions = int(plan.MaxSessions.ValueInt64())
    data.Notes = plan.Notes.ValueString()

    id, err := r.client.Create("admin/create_relay", &data)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to create networknext relay",
            "An error occurred when calling the networknext API to create a relay. "+
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

func (r *relayResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

    var state RelayModel
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    response := ReadRelayResponse{}

    err := r.client.GetJSON(ctx, fmt.Sprintf("admin/relay/%x", int64(state.Id.ValueInt64())), &response)

    if err != nil {        
        resp.Diagnostics.AddError(
            "Unable to read networknext relay",
            "An unexpected error occurred when calling the networknext API to read a relay. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if response.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to read networknext relay",
            "The networknext API returned an error while trying to read a relay. "+
                "Network Next Client Error: "+response.Error,
        )
        return
    }

    data := &response.Relay

    state.Id = types.Int64Value(int64(data.RelayId))
    state.Name = types.StringValue(data.RelayName)
    state.DatacenterId = types.Int64Value(int64(data.DatacenterId))
    state.PublicIP = types.StringValue(data.PublicIP)
    state.PublicPort = types.Int64Value(int64(data.PublicPort))
    state.InternalIP = types.StringValue(data.InternalIP)
    state.InternalPort = types.Int64Value(int64(data.InternalPort))
    state.InternalGroup = types.StringValue(data.InternalGroup)
    state.SSH_IP = types.StringValue(data.SSH_IP)
    state.SSH_Port = types.Int64Value(int64(data.SSH_Port))
    state.SSH_User = types.StringValue(data.SSH_User)
    state.PublicKeyBase64 = types.StringValue(data.PublicKeyBase64)
    state.PrivateKeyBase64 = types.StringValue(data.PrivateKeyBase64)
    state.Version = types.StringValue(data.Version)
    state.MRC = types.Int64Value(int64(data.MRC))
    state.PortSpeed = types.Int64Value(int64(data.PortSpeed))
    state.MaxSessions = types.Int64Value(int64(data.MaxSessions))
    state.Notes = types.StringValue(data.Notes)

    diags = resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

func (r *relayResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

    var plan RelayModel
    diags := req.Plan.Get(ctx, &plan)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    var data RelayData
    data.RelayName = plan.Name.ValueString()
    data.DatacenterId = uint64(plan.DatacenterId.ValueInt64())
    data.PublicIP = plan.PublicIP.ValueString()
    data.PublicPort = int(plan.PublicPort.ValueInt64())
    data.InternalIP = plan.InternalIP.ValueString()
    data.InternalPort = int(plan.InternalPort.ValueInt64())
    data.InternalGroup = plan.InternalGroup.ValueString()
    data.SSH_IP = plan.SSH_IP.ValueString()
    data.SSH_Port = int(plan.SSH_Port.ValueInt64())
    data.SSH_User = plan.SSH_User.ValueString()
    data.PublicKeyBase64 = plan.PublicKeyBase64.ValueString()
    data.PrivateKeyBase64 = plan.PrivateKeyBase64.ValueString()
    data.Version = plan.Version.ValueString()
    data.MRC = int(plan.MRC.ValueInt64())
    data.PortSpeed = int(plan.PortSpeed.ValueInt64())
    data.MaxSessions = int(plan.MaxSessions.ValueInt64())
    data.Notes = plan.Notes.ValueString()

    err := r.client.Update(ctx, "admin/update_relay", &data)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to update networknext relay",
            "An error occurred when calling the networknext API to update a relay. "+
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

func (r *relayResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

    var state RelayModel
    diags := req.State.Get(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }

    id := state.Id.ValueInt64()

    err := r.client.Delete(ctx, "admin/delete_relay", uint64(id))

    if err != nil {
        resp.Diagnostics.AddError(
            "Error deleting networknext relay",
            "Could not delete relay, unexpected error: "+err.Error(),
        )
        return
    }
}

func (r *relayResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
    resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
