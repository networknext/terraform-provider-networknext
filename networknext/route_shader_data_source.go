package networknext

import (
    "context"

    "github.com/hashicorp/terraform-plugin-framework/datasource"
    "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
    "github.com/hashicorp/terraform-plugin-framework/types"
)

var (
    _ datasource.DataSource              = &routeShadersDataSource{}
    _ datasource.DataSourceWithConfigure = &routeShadersDataSource{}
)

func NewRouteShadersDataSource() datasource.DataSource {
    return &routeShadersDataSource{}
}

type routeShadersDataSource struct {
    client *Client
}

func (d *routeShadersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_route_shaders"
}

func (d *routeShadersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
    resp.Schema = schema.Schema{
        Attributes: map[string]schema.Attribute{
            "route_shaders": schema.ListNestedAttribute{
                Computed: true,
                NestedObject: schema.NestedAttributeObject{
                    Attributes: map[string]schema.Attribute{
                        "id": schema.Int64Attribute{
                            Computed: true,
                        },
                        "name": schema.StringAttribute{
                            Computed: true,
                        },
                        "ab_test": schema.BoolAttribute{
                            Computed: true,
                        },
                        "acceptable_latency": schema.Int64Attribute{
                            Computed: true,
                        },
                        "acceptable_packet_loss": schema.Float64Attribute{
                            Computed: true,
                        },
                        "packet_loss_sustained": schema.Float64Attribute{
                            Computed: true,
                        },
                        "analysis_only": schema.BoolAttribute{
                            Computed: true,
                        },
                        "bandwidth_envelope_up_kbps": schema.Int64Attribute{
                            Computed: true,
                        },
                        "bandwidth_envelope_down_kbps": schema.Int64Attribute{
                            Computed: true,
                        },
                        "disable_network_next": schema.BoolAttribute{
                            Computed: true,
                        },
                        "latency_threshold": schema.Int64Attribute{
                            Computed: true,
                        },
                        "multipath": schema.BoolAttribute{
                            Computed: true,
                        },
                        "reduce_latency": schema.BoolAttribute{
                            Computed: true,
                        },
                        "reduce_packet_loss": schema.BoolAttribute{
                            Computed: true,
                        },
                        "selection_percent": schema.Float64Attribute{
                            Computed: true,
                        },
                        "max_latency_trade_off": schema.Int64Attribute{
                            Computed: true,
                        },
                        "max_next_rtt": schema.Int64Attribute{
                            Computed: true,
                        },
                        "route_switch_threshold": schema.Int64Attribute{
                            Computed: true,
                        },
                        "route_select_threshold": schema.Int64Attribute{
                            Computed: true,
                        },
                        "rtt_veto_default": schema.Int64Attribute{
                            Computed: true,
                        },
                        "rtt_veto_multipath": schema.Int64Attribute{
                            Computed: true,
                        },
                        "rtt_veto_packetloss": schema.Int64Attribute{
                            Computed: true,
                        },
                        "force_next": schema.BoolAttribute{
                            Computed: true,
                        },
                        "route_diversity": schema.Int64Attribute{
                            Computed: true,
                        },
                    },
                },
            },
        },
    }
}

func (d *routeShadersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
    if req.ProviderData == nil {
        return
    }
    d.client = req.ProviderData.(*Client)
}

func (d *routeShadersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

    routeShadersResponse := ReadRouteShadersResponse{}
    
    err := d.client.GetJSON(ctx, "admin/route_shaders", &routeShadersResponse)
    
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to get networknext route shaders",
            "An error occurred when calling the networknext API to get route shaders. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+err.Error(),
        )
        return
    }

    if routeShadersResponse.Error != "" {
        resp.Diagnostics.AddError(
            "Unable to get networknext route shaders",
            "An error occurred when calling the networknext API to get route shaders. "+
                "Please check that your network next instance is running and properly configured.\n\n"+
                "Network Next Client Error: "+routeShadersResponse.Error,
        )
        return
    }

    var state RouteShadersModel

    for i := range routeShadersResponse.RouteShaders {

        routeShaderState := RouteShaderModel{
            Id:                        types.Int64Value(int64(routeShadersResponse.RouteShaders[i].RouteShaderId)),
            Name:                      types.StringValue(routeShadersResponse.RouteShaders[i].RouteShaderName),
            ABTest:                    types.BoolValue(routeShadersResponse.RouteShaders[i].ABTest),
            AcceptableLatency:         types.Int64Value(int64(routeShadersResponse.RouteShaders[i].AcceptableLatency)),
            AcceptablePacketLoss:      types.Float64Value(float64(routeShadersResponse.RouteShaders[i].AcceptablePacketLoss)),
            PacketLossSustained:       types.Float64Value(float64(routeShadersResponse.RouteShaders[i].PacketLossSustained)),
            AnalysisOnly:              types.BoolValue(routeShadersResponse.RouteShaders[i].AnalysisOnly),
            BandwidthEnvelopeUpKbps:   types.Int64Value(int64(routeShadersResponse.RouteShaders[i].BandwidthEnvelopeUpKbps)),
            BandwidthEnvelopeDownKbps: types.Int64Value(int64(routeShadersResponse.RouteShaders[i].BandwidthEnvelopeDownKbps)),
            DisableNetworkNext:        types.BoolValue(routeShadersResponse.RouteShaders[i].DisableNetworkNext),
            LatencyThreshold:          types.Int64Value(int64(routeShadersResponse.RouteShaders[i].LatencyThreshold)),
            Multipath:                 types.BoolValue(routeShadersResponse.RouteShaders[i].Multipath),
            ReduceLatency:             types.BoolValue(routeShadersResponse.RouteShaders[i].ReduceLatency),
            ReducePacketLoss:          types.BoolValue(routeShadersResponse.RouteShaders[i].ReducePacketLoss),
            SelectionPercent:          types.Float64Value(float64(routeShadersResponse.RouteShaders[i].SelectionPercent)),
            MaxLatencyTradeOff:        types.Int64Value(int64(routeShadersResponse.RouteShaders[i].MaxLatencyTradeOff)),
            MaxNextRTT:                types.Int64Value(int64(routeShadersResponse.RouteShaders[i].MaxNextRTT)),
            RouteSwitchThreshold:      types.Int64Value(int64(routeShadersResponse.RouteShaders[i].RouteSwitchThreshold)),
            RouteSelectThreshold:      types.Int64Value(int64(routeShadersResponse.RouteShaders[i].RouteSelectThreshold)),
            RTTVeto_Default:           types.Int64Value(int64(routeShadersResponse.RouteShaders[i].RTTVeto_Default)),
            RTTVeto_Multipath:         types.Int64Value(int64(routeShadersResponse.RouteShaders[i].RTTVeto_Multipath)),
            RTTVeto_PacketLoss:        types.Int64Value(int64(routeShadersResponse.RouteShaders[i].RTTVeto_PacketLoss)),
            ForceNext:                 types.BoolValue(routeShadersResponse.RouteShaders[i].ForceNext),
            RouteDiversity:            types.Int64Value(int64(routeShadersResponse.RouteShaders[i].RouteDiversity)),
        }

        state.RouteShaders = append(state.RouteShaders, routeShaderState)
    }

    diags := resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}
