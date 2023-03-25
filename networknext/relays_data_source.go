package networknext

import (
    "context"

    "github.com/hashicorp/terraform-plugin-framework/datasource"
    "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
    "github.com/hashicorp/terraform-plugin-framework/types"
)

var (
    _ datasource.DataSource              = &relaysDataSource{}
    _ datasource.DataSourceWithConfigure = &relaysDataSource{}
)

func NewRelaysDataSource() datasource.DataSource {
    return &relaysDataSource{}
}

type relaysDataSource struct {
    client *Client
}

type relaysDataSourceModel struct {
    Relays []RelayModel `tfsdk:"relays"`
}

func (d *relaysDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_relays"
}

func (d *relaysDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
    resp.Schema = schema.Schema{
        Attributes: map[string]schema.Attribute{
            "relays": schema.ListNestedAttribute{
                Computed: true,
                NestedObject: schema.NestedAttributeObject{
                    Attributes: map[string]schema.Attribute{
                        "id": schema.Int64Attribute{
                            Computed: true,
                        },
                        "name": schema.StringAttribute{
                            Computed: true,
                        },
                        "datacenter_id": schema.Int64Attribute{
                            Computed: true,
                        },
                        "public_ip": schema.StringAttribute{
                            Computed: true,
                        },
                        "public_port": schema.Int64Attribute{
                            Computed: true,
                        },
                        "internal_ip": schema.StringAttribute{
                            Computed: true,
                        },
                        "internal_port": schema.Int64Attribute{
                            Computed: true,
                        },
                        "internal_group": schema.StringAttribute{
                            Computed: true,
                        },
                        "ssh_ip": schema.StringAttribute{
                            Computed: true,
                        },
                        "ssh_port": schema.Int64Attribute{
                            Computed: true,
                        },
                        "ssh_user": schema.StringAttribute{
                            Computed: true,
                        },
                        "public_key_base64": schema.StringAttribute{
                            Computed: true,
                        },
                        "private_key_base64": schema.StringAttribute{
                            Computed: true,
                        },
                        "version": schema.StringAttribute{
                            Computed: true,
                        },
                        "mrc": schema.Int64Attribute{
                            Computed: true,
                        },
                        "port_speed": schema.Int64Attribute{
                            Computed: true,
                        },
                        "max_sessions": schema.Int64Attribute{
                            Computed: true,
                        },
                        "notes": schema.StringAttribute{
                            Computed: true,
                        },
                    },
                },
            },
        },
    }
}

func (d *relaysDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
    if req.ProviderData == nil {
        return
    }
    d.client = req.ProviderData.(*Client)
}

type RelayModel struct {
    Id               types.Int64   `tfsdk:"id"`
    Name             types.String  `tfsdk:"name"`
    DatacenterId     types.Int64   `tfsdk:"datacenter_id"`
    PublicIP         types.String  `tfsdk:"public_ip"`
    PublicPort       types.Int64   `tfsdk:"public_port"`
    InternalIP       types.String  `tfsdk:"internal_ip"`
    InternalPort     types.Int64   `tfsdk:"internal_port"`
    InternalGroup    types.String  `tfsdk:"internal_group"`
    SSH_IP           types.String  `tfsdk:"ssh_ip"`
    SSH_Port         types.Int64   `tfsdk:"ssh_port"`
    SSH_User         types.String  `tfsdk:"ssh_user"`
    PublicKeyBase64  types.String  `tfsdk:"public_key_base64"`
    PrivateKeyBase64 types.String  `tfsdk:"private_key_base64"`
    Version          types.String  `tfsdk:"version"`
    MRC              types.Int64   `tfsdk:"mrc"`
    PortSpeed        types.Int64   `tfsdk:"port_speed"`
    MaxSessions      types.Int64   `tfsdk:"max_sessions"`
    Notes            types.String  `tfsdk:"notes"`
}

type RelayData struct {
    RelayId          uint64 `json:"relay_id"`
    RelayName        string `json:"relay_name"`
    DatacenterId     uint64 `json:"datacenter_id"`
    PublicIP         string `json:"public_ip"`
    PublicPort       int    `json:"public_port"`
    InternalIP       string `json:"internal_ip"`
    InternalPort     int    `json:"internal_port`
    InternalGroup    string `json:"internal_group`
    SSH_IP           string `json:"ssh_ip"`
    SSH_Port         int    `json:"ssh_port`
    SSH_User         string `json:"ssh_user`
    PublicKeyBase64  string `json:"public_key_base64"`
    PrivateKeyBase64 string `json:"private_key_base64"`
    Version          string `json:"version"`
    MRC              int    `json:"mrc"`
    PortSpeed        int    `json:"port_speed"`
    MaxSessions      int    `json:"max_sessions"`
    Notes            string `json:"notes"`
}

type RelaysResponse struct {
    Relays []RelayData `json:"relays"`
    Error  string      `json:"error"`
}

func (d *relaysDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

    relaysResponse := RelaysResponse{}
    
    err := d.client.GetJSON(ctx, "admin/relays", &relaysResponse)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to get networknext relays",
            "An error occurred when calling the networknext API to get relays. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if relaysResponse.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to get networknext relays",
            "An error occurred when calling the networknext API to get relays. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+relaysResponse.Error,
        )
        return
    }

    var state relaysDataSourceModel

    for i := range relaysResponse.Relays {

        relayState := RelayModel{
            Id:               types.Int64Value(int64(relaysResponse.Relays[i].RelayId)),
            Name:             types.StringValue(relaysResponse.Relays[i].RelayName),
            DatacenterId:     types.Int64Value(int64(relaysResponse.Relays[i].DatacenterId)),
            PublicIP:         types.StringValue(relaysResponse.Relays[i].PublicIP),
            PublicPort:       types.Int64Value(int64(relaysResponse.Relays[i].PublicPort)),
            InternalIP:       types.StringValue(relaysResponse.Relays[i].InternalIP),
            InternalPort:     types.Int64Value(int64(relaysResponse.Relays[i].InternalPort)),
            InternalGroup:    types.StringValue(relaysResponse.Relays[i].InternalGroup),
            SSH_IP:           types.StringValue(relaysResponse.Relays[i].SSH_IP),
            SSH_Port:         types.Int64Value(int64(relaysResponse.Relays[i].SSH_Port)),
            SSH_User:         types.StringValue(relaysResponse.Relays[i].SSH_User),
            PublicKeyBase64:  types.StringValue(relaysResponse.Relays[i].PublicKeyBase64),
            PrivateKeyBase64: types.StringValue(relaysResponse.Relays[i].PrivateKeyBase64),
            Version:          types.StringValue(relaysResponse.Relays[i].Version),
            MRC:              types.Int64Value(int64(relaysResponse.Relays[i].MRC)),
            PortSpeed:        types.Int64Value(int64(relaysResponse.Relays[i].PortSpeed)),
            MaxSessions:      types.Int64Value(int64(relaysResponse.Relays[i].MaxSessions)),
        }

        state.Relays = append(state.Relays, relayState)
    }

    diags := resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}
