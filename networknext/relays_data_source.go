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

func (d *relaysDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

    relaysResponse := ReadRelaysResponse{}
    
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

    var state RelaysModel

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
