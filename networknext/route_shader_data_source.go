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

type routeShadersDataSourceModel struct {
    RouteShaders []RouteShaderModel `tfsdk:"route_shaders"`
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

type RouteShaderModel struct {
    Id                        types.Int64   `tfsdk:"id"`
    Name                      types.String  `tfsdk:"name"`
    ABTest                    types.Bool    `tfsdk:"ab_test"`
    AcceptableLatency         types.Int64   `tfsdk:"acceptable_latency"`
    AcceptablePacketLoss      types.Float64 `tfsdk:"acceptable_packet_loss"`
    PacketLossSustained       types.Float64 `tfsdk:"packet_loss_sustained"`
    AnalysisOnly              types.Bool    `tfsdk:"analysis_only"`
    BandwidthEnvelopeUpKbps   types.Int64   `tfsdk:"bandwidth_envelope_up_kbps"`
    BandwidthEnvelopeDownKbps types.Int64   `tfsdk:"bandwidth_envelope_down_kbps"`
    DisableNetworkNext        types.Bool    `tfsdk:"disable_network_next"`
    LatencyThreshold          types.Int64   `tfsdk:"latency_threshold"`
    Multipath                 types.Bool    `tfsdk:"multipath"`
    ReduceLatency             types.Bool    `tfsdk:"reduce_latency"`
    ReducePacketLoss          types.Bool    `tfsdk:"reduce_packet_loss"`
    SelectionPercent          types.Float64 `tfsdk:"selection_percent"`
    MaxLatencyTradeOff        types.Int64   `tfsdk:"max_latency_trade_off"`
    MaxNextRTT                types.Int64   `tfsdk:"max_next_rtt"`
    RouteSwitchThreshold      types.Int64   `tfsdk:"route_switch_threshold"`
    RouteSelectThreshold      types.Int64   `tfsdk:"route_select_threshold"`
    RTTVeto_Default           types.Int64   `tfsdk:"rtt_veto_default"`
    RTTVeto_Multipath         types.Int64   `tfsdk:"rtt_veto_multipath"`
    RTTVeto_PacketLoss        types.Int64   `tfsdk:"rtt_veto_packetloss"`
    ForceNext                 types.Bool    `tfsdk:"force_next"`
    RouteDiversity            types.Int64   `tfsdk:"route_diversity"`
}

type RouteShaderData struct {
    RouteShaderId             uint64  `json:"route_shader_id"`
    RouteShaderName           string  `json:"route_shader_name"`
    ABTest                    bool    `json:"ab_test"`
    AcceptableLatency         int     `json:"acceptable_latency"`
    AcceptablePacketLoss      float32 `json:"acceptable_packet_loss"`
    PacketLossSustained       float32 `json:"packet_loss_sustained"`
    AnalysisOnly              bool    `json:"analysis_only"`
    BandwidthEnvelopeUpKbps   int     `json:"bandwidth_envelope_up_kbps"`
    BandwidthEnvelopeDownKbps int     `json:"bandwidth_envelope_down_kbps"`
    DisableNetworkNext        bool    `json:"disable_network_next"`
    LatencyThreshold          int     `json:"latency_threshold"`
    Multipath                 bool    `json:"multipath"`
    ReduceLatency             bool    `json:"reduce_latency"`
    ReducePacketLoss          bool    `json:"reduce_packet_loss"`
    SelectionPercent          float32 `json:"selection_percent"`
    MaxLatencyTradeOff        int     `json:"max_latency_trade_off"`
    MaxNextRTT                int     `json:"max_next_rtt"`
    RouteSwitchThreshold      int     `json:"route_switch_threshold"`
    RouteSelectThreshold      int     `json:"route_select_threshold"`
    RTTVeto_Default           int     `json:"rtt_veto_default"`
    RTTVeto_Multipath         int     `json:"rtt_veto_multipath"`
    RTTVeto_PacketLoss        int     `json:"rtt_veto_packetloss"`
    ForceNext                 bool    `json:"force_next"`
    RouteDiversity            int     `json:"route_diversity"`
}

type RouteShadersResponse struct {
    RouteShaders []RouteShaderData `json:"route_shaders"`
    Error        string            `json:"error"`
}

func (d *routeShadersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

    routeShadersResponse := RouteShadersResponse{}
    
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

    var state routeShadersDataSourceModel

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
