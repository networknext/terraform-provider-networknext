package networknext

import (
    "github.com/hashicorp/terraform-plugin-framework/types"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
    datasource_schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// -------------------------------------------------------------------

type CustomerData struct {
    CustomerId   uint64 `json:"customer_id"`
    CustomerName string `json:"customer_name"`
    CustomerCode string `json:"customer_code"`
    Live         bool   `json:"live"`
    Debug        bool   `json:"debug"`
}

type ReadCustomersResponse struct {
    Customers []CustomerData `json:"customers"`
    Error     string         `json:"error"`
}

type CustomersModel struct {
    Customers []CustomerModel `tfsdk:"customers"`
}

type CustomerModel struct {
    Id    types.Int64  `tfsdk:"id"`
    Name  types.String `tfsdk:"name"`
    Code  types.String `tfsdk:"code"`
    Live  types.Bool   `tfsdk:"live"`
    Debug types.Bool   `tfsdk:"debug"`
}

type ReadCustomerResponse struct {
    Customer CustomerData `json:"customer"`
    Error    string       `json:"error"`
}

func CustomerModelToData(model *CustomerModel, data *CustomerData) {
    model.Id = types.Int64Value(int64(data.CustomerId))
    model.Name = types.StringValue(data.CustomerName)
    model.Code = types.StringValue(data.CustomerCode)
    model.Live = types.BoolValue(data.Live)
    model.Debug = types.BoolValue(data.Debug)
}

func CustomerDataToModel(data *CustomerData, model *CustomerModel) {
    data.CustomerName = model.Name.ValueString()
    data.CustomerCode = model.Code.ValueString()
    data.Live = model.Live.ValueBool()
    data.Debug = model.Debug.ValueBool()    
}

func CustomerSchema() schema.Schema {
    return schema.Schema{
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
            "code": schema.StringAttribute{
                Required: true,
            },
            "live": schema.BoolAttribute{
                Optional: true,
            },
            "debug": schema.BoolAttribute{
                Optional: true,
            },
        },
    }
}

func CustomersSchema() datasource_schema.Schema {
    return datasource_schema.Schema{
        Attributes: map[string]datasource_schema.Attribute{
            "customers": datasource_schema.ListNestedAttribute{
                Computed: true,
                NestedObject: datasource_schema.NestedAttributeObject{
                    Attributes: map[string]datasource_schema.Attribute{
                        "id": datasource_schema.Int64Attribute{
                            Computed: true,
                        },
                        "name": datasource_schema.StringAttribute{
                            Computed: true,
                        },
                        "code": datasource_schema.StringAttribute{
                            Computed: true,
                        },
                        "live": datasource_schema.BoolAttribute{
                            Computed: true,
                        },
                        "debug": datasource_schema.BoolAttribute{
                            Computed: true,
                        },
                    },
                },
            },
        },
    }
}

// -------------------------------------------------------------------

type BuyerData struct {
    BuyerId         uint64 `json:"buyer_id"`
    BuyerName       string `json:"buyer_name"`
    PublicKeyBase64 string `json:"public_key_base64"`
    CustomerId      uint64 `json:"customer_id"`
    RouteShaderId   uint64 `json:"route_shader_id"`
}

type ReadBuyerResponse struct {
    Buyer  BuyerData `json:"buyer"`
    Error  string    `json:"error"`
}

type ReadBuyersResponse struct {
    Buyers []BuyerData `json:"buyers"`
    Error  string      `json:"error"`
}

type BuyerModel struct {
    Id              types.Int64  `tfsdk:"id"`
    Name            types.String `tfsdk:"name"`
    PublicKeyBase64 types.String `tfsdk:"public_key_base64"`
    CustomerId      types.Int64  `tfsdk:"customer_id"`
    RouteShaderId   types.Int64  `tfsdk:"route_shader_id"`
}

type BuyersModel struct {
    Buyers []BuyerModel `tfsdk:"buyers"`
}

func BuyerDataToModel(data *BuyerData, model *BuyerModel) {
    model.Id = types.Int64Value(int64(data.BuyerId))
    model.Name = types.StringValue(data.BuyerName)
    model.PublicKeyBase64 = types.StringValue(data.PublicKeyBase64)
    model.CustomerId = types.Int64Value(int64(data.CustomerId))
    model.RouteShaderId = types.Int64Value(int64(data.RouteShaderId))
}   

func BuyerModelToData(model *BuyerModel, data *BuyerData) {
    data.BuyerId = uint64(model.Id.ValueInt64())
    data.BuyerName = model.Name.ValueString()
    data.PublicKeyBase64 = model.PublicKeyBase64.ValueString()
    data.CustomerId = uint64(model.CustomerId.ValueInt64())
    data.RouteShaderId = uint64(model.RouteShaderId.ValueInt64())
}

func BuyerSchema() schema.Schema {
    return schema.Schema{
        Attributes: map[string]schema.Attribute{
            "buyers": schema.ListNestedAttribute{
                Computed: true,
                NestedObject: schema.NestedAttributeObject{
                    Attributes: map[string]schema.Attribute{
                        "id": schema.Int64Attribute{
                            Computed: true,
                        },
                        "name": schema.StringAttribute{
                            Computed: true,
                        },
                        "public_key_base64": schema.StringAttribute{
                            Computed: true,
                        },
                        "customer_id": schema.Int64Attribute{
                            Computed: true,
                        },
                        "route_shader_id": schema.Int64Attribute{
                            Computed: true,
                        },
                    },
                },
            },
        },
    }    
}

func BuyersSchema() datasource_schema.Schema {
    return datasource_schema.Schema{
        Attributes: map[string]datasource_schema.Attribute{
            "buyers": schema.ListNestedAttribute{
                Computed: true,
                NestedObject: schema.NestedAttributeObject{
                    Attributes: map[string]schema.Attribute{
                        "id": schema.Int64Attribute{
                            Computed: true,
                        },
                        "name": schema.StringAttribute{
                            Computed: true,
                        },
                        "public_key_base64": schema.StringAttribute{
                            Computed: true,
                        },
                        "customer_id": schema.Int64Attribute{
                            Computed: true,
                        },
                        "route_shader_id": schema.Int64Attribute{
                            Computed: true,
                        },
                    },
                },
            },
        },
    }    
}

// -------------------------------------------------------------------

type SellerData struct {
    SellerId         uint64 `json:"seller_id"`
    SellerName       string `json:"seller_name"`
    CustomerId      uint64 `json:"customer_id"`
}

type ReadSellersResponse struct {
    Sellers []SellerData `json:"sellers"`
    Error  string        `json:"error"`
}

type SellerModel struct {
    Id              types.Int64  `tfsdk:"id"`
    Name            types.String `tfsdk:"name"`
    CustomerId      types.Int64  `tfsdk:"customer_id"`
}

type SellersModel struct {
    Sellers []SellerModel `tfsdk:"sellers"`
}

func SellerDataToModel(data *SellerData, model *SellerModel) {
    model.Id = types.Int64Value(int64(data.SellerId))
    model.Name = types.StringValue(data.SellerName)
    model.CustomerId = types.Int64Value(int64(data.CustomerId))
}

func SellerModelToData(model *SellerModel, data *SellerData) {
    data.SellerId = uint64(model.Id.ValueInt64())
    data.SellerName = model.Name.ValueString()
    data.CustomerId = uint64(model.CustomerId.ValueInt64())    
}

func SellerSchema() schema.Schema {
    return schema.Schema{
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
            "customer_id": schema.Int64Attribute{
                Required: true,
                // Optional: true,
                // Default: int64default.StaticValue(0),
            },
        },
    }
}

func SellersSchema() datasource_schema.Schema {
    return datasource_schema.Schema{
        Attributes: map[string]datasource_schema.Attribute{
            "sellers": schema.ListNestedAttribute{
                Computed: true,
                NestedObject: schema.NestedAttributeObject{
                    Attributes: map[string]schema.Attribute{
                        "id": schema.Int64Attribute{
                            Computed: true,
                        },
                        "name": schema.StringAttribute{
                            Computed: true,
                        },
                        "customer_id": schema.Int64Attribute{
                            Computed: true,
                        },
                    },
                },
            },
        },
    }
}

// -------------------------------------------------------------------

type DatacenterData struct {
    DatacenterId   uint64  `json:"datacenter_id"`
    DatacenterName string  `json:"datacenter_name"`
    NativeName     string  `json:"native_name"`
    Latitude       float32 `json:"latitude"`
    Longitude      float32 `json:"longitude"`
    SellerId       uint64  `json:"seller_id"`
    Notes          string  `json:"notes"`
}

type ReadDatacentersResponse struct {
    Datacenters []DatacenterData `json:"datacenters"`
    Error       string           `json:"error"`
}

type ReadDatacenterResponse struct {
    Datacenter DatacenterData `json:"datacenter"`
    Error      string         `json:"error"`
}

type DatacenterModel struct {
    Id              types.Int64   `tfsdk:"id"`
    Name            types.String  `tfsdk:"name"`
    NativeName      types.String  `tfsdk:"native_name"`
    Latitude        types.Float64 `tfsdk:"latitude"`
    Longitude       types.Float64 `tfsdk:"longitude"`
    SellerId        types.Int64   `tfsdk:"seller_id"`
    Notes           types.String  `tfsdk:"notes"`
}

type DatacentersModel struct {
    Datacenters []DatacenterModel `tfsdk:"datacenters"`
}

// -------------------------------------------------------------------

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

type ReadRelaysResponse struct {
    Relays []RelayData `json:"relays"`
    Error  string      `json:"error"`
}

type ReadRelayResponse struct {
    Relay   RelayData `json:"relay"`
    Error   string    `json:"error"`
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

type RelaysModel struct {
    Relays []RelayModel `tfsdk:"relays"`
}

// -------------------------------------------------------------------

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

type ReadRouteShadersResponse struct {
    RouteShaders []RouteShaderData `json:"route_shaders"`
    Error        string            `json:"error"`
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

type RouteShadersModel struct {
    RouteShaders []RouteShaderModel `tfsdk:"route_shaders"`
}

// -------------------------------------------------------------------