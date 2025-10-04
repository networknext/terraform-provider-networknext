package networknext

import (
    "github.com/hashicorp/terraform-plugin-framework/types"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
    datasource_schema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// -------------------------------------------------------------------

type SellerData struct {
    SellerId         uint64 `json:"seller_id"`
    SellerName       string `json:"seller_name"`
    SellerCode       string `json:"seller_code"`
}

type SellerModel struct {
    Id              types.Int64  `tfsdk:"id"`
    Name            types.String `tfsdk:"name"`
    Code            types.String `tfsdk:"code"`
}

type SellersModel struct {
    Sellers []SellerModel `tfsdk:"sellers"`
}

type CreateSellerResponse struct {
    Seller   SellerData   `json:"seller"`
    Error    string       `json:"error"`
}

type ReadSellersResponse struct {
    Sellers []SellerData `json:"sellers"`
    Error  string        `json:"error"`
}

type UpdateSellerResponse struct {
    Seller   SellerData   `json:"seller"`
    Error    string       `json:"error"`
}

func SellerDataToModel(data *SellerData, model *SellerModel) {
    model.Id = types.Int64Value(int64(data.SellerId))
    model.Name = types.StringValue(data.SellerName)
    model.Code = types.StringValue(data.SellerCode)
}

func SellerModelToData(model *SellerModel, data *SellerData) {
    data.SellerId = uint64(model.Id.ValueInt64())
    data.SellerName = model.Name.ValueString()
    data.SellerCode = model.Code.ValueString()
}

func SellerSchema() schema.Schema {
    return schema.Schema{
        Description: "Manages a seller.",
        Attributes: map[string]schema.Attribute{
            "id": schema.Int64Attribute{
                Description: "The id of the seller. Automatically generated when sellers are created.",
                Computed: true,
                PlanModifiers: []planmodifier.Int64{
                    int64planmodifier.UseStateForUnknown(),
                },
            },
            "name": schema.StringAttribute{
                Description: "The name of the seller. For example, \"Google\", \"Amazon\" or \"Akamai\"", 
                Required: true,
            },
            "code": schema.StringAttribute{
                Description: "Short seller code. For example, \"google\", \"amazon\" or \"akamai\"", 
                Required: true,
            },
        },
    }
}

func SellersSchema() datasource_schema.Schema {
    return datasource_schema.Schema{
        Description: "Fetches the list of sellers.",
        Attributes: map[string]datasource_schema.Attribute{
            "sellers": schema.ListNestedAttribute{
                Computed: true,
                NestedObject: schema.NestedAttributeObject{
                    Attributes: map[string]schema.Attribute{
                        "id": schema.Int64Attribute{
                            Description: "The id of the seller. Automatically generated when sellers are created.",
                            Computed: true,
                        },
                        "name": schema.StringAttribute{
                            Description: "The name of the seller. For example, \"Google\", \"Amazon\" or \"Akamai\"", 
                            Computed: true,
                        },
                        "code": schema.StringAttribute{
                            Description: "A short code for the seller. For example, \"google\", \"amazon\" or \"akamai\"", 
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
    BuyerCode       string `json:"buyer_code"`
    PublicKeyBase64 string `json:"public_key_base64"`
    RouteShaderId   uint64 `json:"route_shader_id"`
    Live            bool   `json:"live"`
    Debug           bool   `json:"debug"`
}

type BuyerModel struct {
    Id              types.Int64  `tfsdk:"id"`
    Name            types.String `tfsdk:"name"`
    Code            types.String `tfsdk:"code"`
    PublicKeyBase64 types.String `tfsdk:"public_key_base64"`
    RouteShaderId   types.Int64  `tfsdk:"route_shader_id"`
    Live            types.Bool   `tfsdk:"live"`
    Debug           types.Bool   `tfsdk:"debug"`
}

type BuyersModel struct {
    Buyers []BuyerModel `tfsdk:"buyers"`
}

type CreateBuyerResponse struct {
    Buyer    BuyerData    `json:"buyer"`
    Error    string       `json:"error"`
}

type ReadBuyerResponse struct {
    Buyer    BuyerData    `json:"buyer"`
    Error    string       `json:"error"`
}

type ReadBuyersResponse struct {
    Buyers    []BuyerData    `json:"buyers"`
    Error     string         `json:"error"`
}

type UpdateBuyerResponse struct {
    Buyer    BuyerData    `json:"buyer"`
    Error    string       `json:"error"`
}

type DeleteBuyerResponse struct {
    Error    string       `json:"error"`
}

func BuyerDataToModel(data *BuyerData, model *BuyerModel) {
    model.Id = types.Int64Value(int64(data.BuyerId))
    model.Name = types.StringValue(data.BuyerName)
    model.Code = types.StringValue(data.BuyerCode)
    model.PublicKeyBase64 = types.StringValue(data.PublicKeyBase64)
    model.RouteShaderId = types.Int64Value(int64(data.RouteShaderId))
    model.Live = types.BoolValue(data.Live)
    model.Debug = types.BoolValue(data.Debug)
}   

func BuyerModelToData(model *BuyerModel, data *BuyerData) {
    data.BuyerId = uint64(model.Id.ValueInt64())
    data.BuyerName = model.Name.ValueString()
    data.BuyerCode = model.Code.ValueString()
    data.PublicKeyBase64 = model.PublicKeyBase64.ValueString()
    data.RouteShaderId = uint64(model.RouteShaderId.ValueInt64())
    data.Live = model.Live.ValueBool()
    data.Debug = model.Debug.ValueBool()
}

func BuyerSchema() schema.Schema {
    return schema.Schema{
        Description: "Manages a buyer.",
        Attributes: map[string]schema.Attribute{
            "id": schema.Int64Attribute{
                Description: "The id of the buyer. Automatically generated when buyers are created.",
                Computed: true,
                PlanModifiers: []planmodifier.Int64{
                    int64planmodifier.UseStateForUnknown(),
                },
            },
            "name": schema.StringAttribute{
                Description: "The name of the buyer. For example, \"Riot Games\", \"Valve\" or \"Respawn Entertainment\"", 
                Required: true,
            },
            "code": schema.StringAttribute{
                Description: "A short buyer code. For example, \"riot\", \"valve\" or \"respawn\"", 
                Required: true,
            },
            "public_key_base64": schema.StringAttribute{
                Description: "The buyer public key base64 string. To generate a keypair run 'keygen' in the SDK. Keep the private portion secret, and paste the public key into this field for the buyer.", 
                Required: true,
            },
            "route_shader_id": schema.Int64Attribute{
                Description: "The id of the route shader for this buyer. The route shader configures when to accelerate traffic for this buyer.", 
                Required: true,
            },
            "live": schema.BoolAttribute{
                Description: "If true then the buyer is live and can use network next.",
                Optional: true,
                Computed: true,
                Default: booldefault.StaticBool(true),
            },
            "debug": schema.BoolAttribute{
                Description: "If true then additional debug information is displayed in the network next client to assist with debugging.",
                Optional: true,
                Computed: true,
                Default: booldefault.StaticBool(false),
            },    
        },
    }    
}

func BuyersSchema() datasource_schema.Schema {
    return datasource_schema.Schema{
        Description: "Fetches the list of buyers.",
        Attributes: map[string]datasource_schema.Attribute{
            "buyers": schema.ListNestedAttribute{
                Computed: true,
                NestedObject: schema.NestedAttributeObject{
                    Attributes: map[string]schema.Attribute{
                        "id": schema.Int64Attribute{
                            Description: "The id of the buyer. Automatically generated when buyers are created.",
                            Computed: true,
                        },
                        "name": schema.StringAttribute{
                            Description: "The name of the buyer. For example, \"Riot Games\", \"Valve\" or \"Respawn Entertainment\"", 
                            Computed: true,
                        },
                        "code": schema.StringAttribute{
                            Description: "Short buyer code. For example, \"riot\", \"valve\" or \"respawn\"", 
                            Computed: true,
                        },
                        "public_key_base64": schema.StringAttribute{
                            Description: "The buyer public key base64 string. To generate a keypair run 'keygen' in the SDK. Keep the private portion secret, and paste the public key into this field for the buyer.", 
                            Computed: true,
                        },
                        "route_shader_id": schema.Int64Attribute{
                            Description: "The id of the route shader for this buyer. The route shader configures when to accelerate traffic for this buyer.", 
                            Computed: true,
                        },
                        "live": schema.BoolAttribute{
                            Description: "If true then the buyer is live and can use network next.",
                            Optional: true,
                            Computed: true,
                            Default: booldefault.StaticBool(true),
                        },
                        "debug": schema.BoolAttribute{
                            Description: "If true then additional debug information is displayed in the network next client to assist with debugging.",
                            Optional: true,
                            Computed: true,
                            Default: booldefault.StaticBool(false),
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
    Latitude       float64 `json:"latitude"`
    Longitude      float64 `json:"longitude"`
    SellerId       uint64  `json:"seller_id"`
    Notes          string  `json:"notes"`
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

type CreateDatacenterResponse struct {
    Datacenter  DatacenterData  `json:"datacenter"`
    Error       string          `json:"error"`
}

type ReadDatacentersResponse struct {
    Datacenters []DatacenterData `json:"datacenters"`
    Error       string           `json:"error"`
}

type UpdateDatacenterResponse struct {
    Datacenter  DatacenterData  `json:"datacenter"`
    Error       string          `json:"error"`
}

type ReadDatacenterResponse struct {
    Datacenter DatacenterData `json:"datacenter"`
    Error      string         `json:"error"`
}

type DeleteDatacenterResponse struct {
    Error    string       `json:"error"`
}

func DatacenterDataToModel(data *DatacenterData, model *DatacenterModel) {
    model.Id = types.Int64Value(int64(data.DatacenterId))
    model.Name = types.StringValue(data.DatacenterName)
    model.NativeName = types.StringValue(data.NativeName)
    model.Latitude = types.Float64Value(float64(data.Latitude))
    model.Longitude = types.Float64Value(float64(data.Longitude))
    model.SellerId = types.Int64Value(int64(data.SellerId))
}

func DatacenterModelToData(model *DatacenterModel, data *DatacenterData) {
    data.DatacenterId = uint64(model.Id.ValueInt64())
    data.DatacenterName = model.Name.ValueString()
    data.NativeName = model.NativeName.ValueString()
    data.Latitude = model.Latitude.ValueFloat64()
    data.Longitude = model.Longitude.ValueFloat64()
    data.SellerId = uint64(model.SellerId.ValueInt64())
    data.Notes = model.Notes.ValueString()
}

func DatacenterSchema() schema.Schema {
    return schema.Schema{
        Description: "Manages datacenters.",
        Attributes: map[string]schema.Attribute{
            "id": schema.Int64Attribute{
                Description: "The id of the datacenter. Automatically generated when datacenters are created.",
                Computed: true,
                PlanModifiers: []planmodifier.Int64{
                    int64planmodifier.UseStateForUnknown(),
                },
            },
            "name": schema.StringAttribute{
                Description: "The name of the datacenter. Must be in the format [seller].[location] with optional datacenter number. For example: google.losangeles.1, vultr.chicago, amazon.virginia.2",
                Required: true,
            },
            "native_name": schema.StringAttribute{
                Description: "The native datacenter name. Used to associate the network next name of a datacenter with the native name of the datacenter on that platform. For example, 'google.taiwan.1' has a native name of 'asia-east1-a'",
                Optional: true,
                Computed: true,
                Default: stringdefault.StaticString(""),
            },
            "latitude": schema.Float64Attribute{
                Description: "The approximate latitude of the datacenter.",
                Required: true,
            },
            "longitude": schema.Float64Attribute{
                Description: "The approximate longitude of the datacenter.",
                Required: true,
            },
            "seller_id": schema.Int64Attribute{
                Description: "The id of the seller this relay belongs to.",
                Required: true,
            },
            "notes": schema.StringAttribute{
                Description: "Optional notes about this datacenter.",
                Optional: true,
                Computed: true,
                Default: stringdefault.StaticString(""),
            },
        },
    }
}

func DatacentersSchema() datasource_schema.Schema {
    return datasource_schema.Schema{
        Description: "Fetches the list of datacenters.",
        Attributes: map[string]datasource_schema.Attribute{
            "datacenters": datasource_schema.ListNestedAttribute{
                Computed: true,
                NestedObject: datasource_schema.NestedAttributeObject{
                    Attributes: map[string]datasource_schema.Attribute{
                        "id": schema.Int64Attribute{
                            Description: "The id of the datacenter. Automatically generated when datacenters are created.",
                            Computed: true,
                        },
                        "name": schema.StringAttribute{
                            Description: "The name of the datacenter. Must be in the format [seller].[location] with optional datacenter number. For example: google.losangeles.1, vultr.chicago, amazon.virginia.2",
                            Computed: true,
                        },
                        "native_name": schema.StringAttribute{
                            Description: "The native datacenter name. Used to associate the network next name of a datacenter with the native name of the datacenter on that platform. For example, 'google.taiwan.1' has a native name of 'asia-east1-a'",
                            Computed: true,
                        },
                        "longitude": schema.Float64Attribute{
                            Description: "The approximate latitude of the datacenter.",
                            Computed: true,
                        },
                        "latitude": schema.Float64Attribute{
                            Description: "The approximate longitude of the datacenter.",
                            Computed: true,
                        },
                        "seller_id": schema.Int64Attribute{
                            Description: "The id of the seller this relay belongs to.",
                            Computed: true,
                        },
                        "notes": schema.StringAttribute{
                            Description: "Optional notes about this datacenter.",
                            Computed: true,
                        },
                    },
                },
            },
        },
    }
}

// -------------------------------------------------------------------

type RelayData struct {
    RelayId          uint64 `json:"relay_id"`
    RelayName        string `json:"relay_name"`
    DatacenterId     uint64 `json:"datacenter_id"`
    PublicIP         string `json:"public_ip"`
    PublicPort       int    `json:"public_port"`
    InternalIP       string `json:"internal_ip"`
    InternalPort     int    `json:"internal_port"`
    InternalGroup    string `json:"internal_group"`
    SSH_IP           string `json:"ssh_ip"`
    SSH_Port         int    `json:"ssh_port"`
    SSH_User         string `json:"ssh_user"`
    PublicKeyBase64  string `json:"public_key_base64"`
    PrivateKeyBase64 string `json:"private_key_base64"`
    Version          string `json:"version"`
    MRC              int    `json:"mrc"`
    PortSpeed        int    `json:"port_speed"`
    MaxSessions      int    `json:"max_sessions"`
    BandwidthPrice   int    `json:"bandwidth_price"`
    Notes            string `json:"notes"`
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
    BandwidthPrice   types.Int64   `tfsdk:"bandwidth_price"`
    Notes            types.String  `tfsdk:"notes"`
}

type RelaysModel struct {
    Relays []RelayModel `tfsdk:"relays"`
}

type CreateRelayResponse struct {
    Relay    RelayData    `json:"relay"`
    Error    string       `json:"error"`
}

type ReadRelaysResponse struct {
    Relays []RelayData `json:"relays"`
    Error  string      `json:"error"`
}

type ReadRelayResponse struct {
    Relay   RelayData `json:"relay"`
    Error   string    `json:"error"`
}

type UpdateRelayResponse struct {
    Relay    RelayData    `json:"relay"`
    Error    string       `json:"error"`
}

type DeleteRelayResponse struct {
    Error    string       `json:"error"`
}

func RelayModelToData(model *RelayModel, data *RelayData) {
    data.RelayId = uint64(model.Id.ValueInt64())
    data.RelayName = model.Name.ValueString()
    data.DatacenterId = uint64(model.DatacenterId.ValueInt64())
    data.PublicIP = model.PublicIP.ValueString()
    data.PublicPort = int(model.PublicPort.ValueInt64())
    data.InternalIP = model.InternalIP.ValueString()
    data.InternalPort = int(model.InternalPort.ValueInt64())
    data.InternalGroup = model.InternalGroup.ValueString()
    data.SSH_IP = model.SSH_IP.ValueString()
    data.SSH_Port = int(model.SSH_Port.ValueInt64())
    data.SSH_User = model.SSH_User.ValueString()
    data.PublicKeyBase64 = model.PublicKeyBase64.ValueString()
    data.PrivateKeyBase64 = model.PrivateKeyBase64.ValueString()
    data.Version = model.Version.ValueString()
    data.MRC = int(model.MRC.ValueInt64())
    data.PortSpeed = int(model.PortSpeed.ValueInt64())
    data.MaxSessions = int(model.MaxSessions.ValueInt64())
    data.BandwidthPrice = int(model.BandwidthPrice.ValueInt64())
    data.Notes = model.Notes.ValueString()
}

func RelayDataToModel(data *RelayData, model *RelayModel) {
    model.Id = types.Int64Value(int64(data.RelayId))
    model.Name = types.StringValue(data.RelayName)
    model.DatacenterId = types.Int64Value(int64(data.DatacenterId))
    model.PublicIP = types.StringValue(data.PublicIP)
    model.PublicPort = types.Int64Value(int64(data.PublicPort))
    model.InternalIP = types.StringValue(data.InternalIP)
    model.InternalPort = types.Int64Value(int64(data.InternalPort))
    model.InternalGroup = types.StringValue(data.InternalGroup)
    model.SSH_IP = types.StringValue(data.SSH_IP)
    model.SSH_Port = types.Int64Value(int64(data.SSH_Port))
    model.SSH_User = types.StringValue(data.SSH_User)
    model.PublicKeyBase64 = types.StringValue(data.PublicKeyBase64)
    model.PrivateKeyBase64 = types.StringValue(data.PrivateKeyBase64)
    model.Version = types.StringValue(data.Version)
    model.MRC = types.Int64Value(int64(data.MRC))
    model.PortSpeed = types.Int64Value(int64(data.PortSpeed))
    model.MaxSessions = types.Int64Value(int64(data.MaxSessions))
    model.BandwidthPrice = types.Int64Value(int64(data.BandwidthPrice))
    model.Notes = types.StringValue(data.Notes)
}

func RelaySchema() schema.Schema {
    return schema.Schema{
        Description: "Manages a relay.",
        Attributes: map[string]schema.Attribute{
            "id": schema.Int64Attribute{
                Description: "The id of the relay. Automatically generated when relays are created.",
                Computed: true,
                PlanModifiers: []planmodifier.Int64{
                    int64planmodifier.UseStateForUnknown(),
                },
            },
            "name": schema.StringAttribute{
                Description: "The name of the relay. The name must be in the form of [datacenter]<.variant>. For example, \"google.losangeles.1\" (same name as datacenter for one relay in the datacenter), or \"amazon.virginia.2.a\", \"amazon.virginia.2.b\" (for two relays in the same datacenter).",
                Required: true,
            },
            "datacenter_id": schema.Int64Attribute{
                Description: "The id of the datacenter this relay is in.",
                Required: true,
            },
            "public_ip": schema.StringAttribute{
                Description: "The public IP address of the relay. For example, \"45.23.66.10\".",
                Required: true,
            },
            "public_port": schema.Int64Attribute{
                Description: "The public UDP port of the relay. By default it is 40000. Make sure this port is open on the firewall to receive UDP packets.",
                Optional: true,
                Computed: true,
                Default: int64default.StaticInt64(40000),
            },
            "internal_ip": schema.StringAttribute{
                Description: "The internal IP address of the relay. Use this only when a seller has an internal network between multiple relays that provide cost or performance benefits. Default value is \"0.0.0.0\" which disables the internal address.",
                Optional: true,
                Computed: true,
                Default: stringdefault.StaticString("0.0.0.0"),
            },
            "internal_port": schema.Int64Attribute{
                Description: "The internal UDP port of the relay. Default value is 0, which uses the same port as the public port. Only set if you need to override to a different port.",
                Optional: true,
                Computed: true,
                Default: int64default.StaticInt64(0),
            },
            "internal_group": schema.StringAttribute{
                Description: "The internal group of the relay. Relays only communicate with internal addresses when they are in the same internal group. Use this when sellers (like amazon) can only use private addresses between a subset of relays. For amazon, set this string to the relay region, then relays will only use the internal addresses for other relays in the same region. Default value is \"\" which lets all relays for a seller communicate with each other via internal address when provided.",
                Optional: true,
                Computed: true,
                Default: stringdefault.StaticString(""),
            },
            "ssh_ip": schema.StringAttribute{
                Description: "The SSH address of the relay. The default value is \"0.0.0.0\" which uses the public address to SSH into. Set this only if you have a specific management IP address that you need to use when SSH'ing into a relay.",
                Optional: true,
                Computed: true,
                Default: stringdefault.StaticString("0.0.0.0"),
            },
            "ssh_port": schema.Int64Attribute{
                Description: "The TCP port to SSH into. Set to 22 by default.",
                Optional: true,
                Computed: true,
                Default: int64default.StaticInt64(22),
            },
            "ssh_user": schema.StringAttribute{
                Description: "The username to SSH in as. Defaults to root. Other common usernames include \"ubuntu\". Seller specific.",
                Optional: true,
                Computed: true,
                Default: stringdefault.StaticString("root"),
            },
            "private_key_base64": schema.StringAttribute{
                Description: "The relay private key as base64.",
                Required: true,
            },
            "public_key_base64": schema.StringAttribute{
                Description: "The relay public key as base64.",
                Required: true,
            },
            "version": schema.StringAttribute{
                Description: "The version of the relay. Default value is \"\". Leave as empty, and you can manage relay versions manually with the next tool. Set to a specific versioen, eg. \"1.0.19\" and the relay will automatically upgrade itself to the relay binary in google cloud storage corresponding to the version you set.",
                Optional: true,
                Computed: true,
                Default: stringdefault.StaticString(""),
            },
            "mrc": schema.Int64Attribute{
                Description: "Monthly recurring cost for this relay in USD. Useful for keeping track of how much relays cost. Default value is zero.",
                Optional: true,
                Computed: true,
                Default: int64default.StaticInt64(0),
            },
            "port_speed": schema.Int64Attribute{
                Description: "The speed of the network port in megabits per-second (mbps). Useful for keeping track of relays with 1gbps (1,000) vs. 10gbps (10,000) speeds.",
                Optional: true,
                Computed: true,
                Default: int64default.StaticInt64(0),
            },
            "max_sessions": schema.Int64Attribute{
                Description: "The maximum number of sessions allowed across this relay at any time. Once this number is exceeded, the network next backend will not send additional sessions across this relay. Default is zero which means unlimited.",
                Optional: true,
                Computed: true,
                Default: int64default.StaticInt64(0),
            },
            "bandwidth_price": schema.Int64Attribute{
                Description: "An integer value representing the bandwidth price for this relay. Used to prefer cheaper bare metal routes over more expensive cloud routes.",
                Optional: true,
                Computed: true,
                Default: int64default.StaticInt64(0),
            },
            "notes": schema.StringAttribute{
                Description: "Optional free form notes to store information about this relay.",
                Optional: true,
                Computed: true,
                Default: stringdefault.StaticString(""),
            },
        },
    }
}

func RelaysSchema() datasource_schema.Schema {
    return datasource_schema.Schema{
        Description: "Fetches the list of relays.",
        Attributes: map[string]datasource_schema.Attribute{
            "relays": datasource_schema.ListNestedAttribute{
                Computed: true,
                NestedObject: datasource_schema.NestedAttributeObject{
                    Attributes: map[string]datasource_schema.Attribute{
                        "id": datasource_schema.Int64Attribute{
                            Description: "The id of the relay. Automatically generated when relays are created.",
                            Computed: true,
                        },
                        "name": datasource_schema.StringAttribute{
                            Description: "The name of the relay. The name must be in the form of [datacenter]<.variant>. For example, \"google.losangeles.1\" (same name as datacenter for one relay in the datacenter), or \"amazon.virginia.2.a\", \"amazon.virginia.2.b\" (for two relays in the same datacenter).",
                            Computed: true,
                        },
                        "datacenter_id": datasource_schema.Int64Attribute{
                            Description: "The id of the datacenter this relay is in.",
                            Computed: true,
                        },
                        "public_ip": datasource_schema.StringAttribute{
                            Description: "The public IP address of the relay. For example, \"45.23.66.10\".",
                            Computed: true,
                        },
                        "public_port": datasource_schema.Int64Attribute{
                            Description: "The public UDP port of the relay. By default it is 40000. Make sure this port is open on the firewall to receive UDP packets.",
                            Computed: true,
                        },
                        "internal_ip": datasource_schema.StringAttribute{
                            Description: "The internal IP address of the relay. Use this only when a seller has an internal network between multiple relays that provide cost or performance benefits. Default value is \"0.0.0.0\" which disables the internal address.",
                            Computed: true,
                        },
                        "internal_port": datasource_schema.Int64Attribute{
                            Description: "The internal UDP port of the relay. Default value is 0, which uses the same port as the public port. Only set if you need to override to a different port.",
                            Computed: true,
                        },
                        "internal_group": datasource_schema.StringAttribute{
                            Description: "The internal group of the relay. Relays only communicate with internal addresses when they are in the same internal group. Use this when sellers (like amazon) can only use private addresses between a subset of relays. For amazon, set this string to the relay region, then relays will only use the internal addresses for other relays in the same region. Default value is \"\" which lets all relays for a seller communicate with each other via internal address when provided.",
                            Computed: true,
                        },
                        "ssh_ip": datasource_schema.StringAttribute{
                            Description: "The SSH address of the relay. The default value is \"0.0.0.0\" which uses the public address to SSH into. Set this only if you have a specific management IP address that you need to use when SSH'ing into a relay.",
                            Computed: true,
                        },
                        "ssh_port": datasource_schema.Int64Attribute{
                            Description: "The TCP port to SSH into. Set to 22 by default.",
                            Computed: true,
                        },
                        "ssh_user": datasource_schema.StringAttribute{
                            Description: "The username to SSH in as. Defaults to root. Other common usernames include \"ubuntu\". Seller specific.",
                            Computed: true,
                        },
                        "private_key_base64": datasource_schema.StringAttribute{
                            Description: "The relay private key as base64.",
                            Computed: true,
                        },
                        "public_key_base64": datasource_schema.StringAttribute{
                            Description: "The relay public key as base64.",
                            Computed: true,
                        },
                        "version": datasource_schema.StringAttribute{
                            Description: "The version of the relay. Default value is \"\". Leave as empty, and you can manage relay versions manually with the next tool. Set to a specific versioen, eg. \"1.0.19\" and the relay will automatically upgrade itself to the relay binary in google cloud storage corresponding to the version you set.",
                            Computed: true,
                        },
                        "mrc": datasource_schema.Int64Attribute{
                            Description: "Monthly recurring cost for this relay in USD. Useful for keeping track of how much relays cost. Default value is zero.",
                            Computed: true,
                        },
                        "port_speed": datasource_schema.Int64Attribute{
                            Description: "The speed of the network port in megabits per-second (mbps). Useful for keeping track of relays with 1gbps (1,000) vs. 10gbps (10,000) speeds.",
                            Computed: true,
                        },
                        "max_sessions": datasource_schema.Int64Attribute{
                            Description: "The maximum number of sessions allowed across this relay at any time. Once this number is exceeded, the network next backend will not send additional sessions across this relay. Default is zero which means unlimited.",
                            Computed: true,
                        },
                        "bandwidth_price": datasource_schema.Int64Attribute{
                            Description: "An integer value representing the bandwidth price for this relay. Used to prefer cheaper bare metal routes over more expensive cloud routes.",
                            Computed: true,
                        },
                        "notes": datasource_schema.StringAttribute{
                            Description: "Optional free form notes to store information about this relay.",
                            Computed: true,
                        },
                    },
                },
            },
        },
    }
}

// -------------------------------------------------------------------

type RouteShaderData struct {
    RouteShaderId                   uint64      `json:"route_shader_id"`
    RouteShaderName                 string      `json:"route_shader_name"`
    ABTest                          bool        `json:"ab_test"`
    AcceptableLatency               int         `json:"acceptable_latency"`
    AcceptablePacketLossInstant     float64     `json:"acceptable_packet_loss_instant"`
    AcceptablePacketLossSustained   float64     `json:"acceptable_packet_loss_sustained"`
    AnalysisOnly                    bool        `json:"analysis_only"`
    BandwidthEnvelopeUpKbps         int         `json:"bandwidth_envelope_up_kbps"`
    BandwidthEnvelopeDownKbps       int         `json:"bandwidth_envelope_down_kbps"`
    DisableNetworkNext              bool        `json:"disable_network_next"`
    LatencyReductionThreshold       int         `json:"latency_reduction_threshold"`
    Multipath                       bool        `json:"multipath"`
    SelectionPercent                float64     `json:"selection_percent"`
    MaxLatencyTradeOff              int         `json:"max_latency_trade_off"`
    MaxNextRTT                      int         `json:"max_next_rtt"`
    RouteSwitchThreshold            int         `json:"route_switch_threshold"`
    RouteSelectThreshold            int         `json:"route_select_threshold"`
    RTTVeto                         int         `json:"rtt_veto"`
    ForceNext                       bool        `json:"force_next"`
    RouteDiversity                  int         `json:"route_diversity"`
}

type RouteShaderModel struct {
    Id                              types.Int64   `tfsdk:"id"`
    Name                            types.String  `tfsdk:"name"`
    ABTest                          types.Bool    `tfsdk:"ab_test"`
    AcceptableLatency               types.Int64   `tfsdk:"acceptable_latency"`
    AcceptablePacketLossInstant     types.Float64 `tfsdk:"acceptable_packet_loss_instant"`
    AcceptablePacketLossSustained   types.Float64 `tfsdk:"acceptable_packet_loss_sustained"`
    AnalysisOnly                    types.Bool    `tfsdk:"analysis_only"`
    BandwidthEnvelopeUpKbps         types.Int64   `tfsdk:"bandwidth_envelope_up_kbps"`
    BandwidthEnvelopeDownKbps       types.Int64   `tfsdk:"bandwidth_envelope_down_kbps"`
    DisableNetworkNext              types.Bool    `tfsdk:"disable_network_next"`
    LatencyReductionThreshold       types.Int64   `tfsdk:"latency_reduction_threshold"`
    Multipath                       types.Bool    `tfsdk:"multipath"`
    SelectionPercent                types.Float64 `tfsdk:"selection_percent"`
    MaxLatencyTradeOff              types.Int64   `tfsdk:"max_latency_trade_off"`
    MaxNextRTT                      types.Int64   `tfsdk:"max_next_rtt"`
    RouteSwitchThreshold            types.Int64   `tfsdk:"route_switch_threshold"`
    RouteSelectThreshold            types.Int64   `tfsdk:"route_select_threshold"`
    RTTVeto                         types.Int64   `tfsdk:"rtt_veto"`
    ForceNext                       types.Bool    `tfsdk:"force_next"`
    RouteDiversity                  types.Int64   `tfsdk:"route_diversity"`
}

type RouteShadersModel struct {
    RouteShaders []RouteShaderModel `tfsdk:"route_shaders"`
}

type CreateRouteShaderResponse struct {
    RouteShader RouteShaderData `json:"route_shader"`
    Error       string          `json:"error"`
}

type ReadRouteShaderResponse struct {
    RouteShader RouteShaderData `json:"route_shader"`
    Error       string          `json:"error"`
}

type ReadRouteShadersResponse struct {
    RouteShaders []RouteShaderData `json:"route_shaders"`
    Error        string            `json:"error"`
}

type UpdateRouteShaderResponse struct {
    RouteShader RouteShaderData `json:"route_shader"`
    Error       string          `json:"error"`
}

type DeleteRouteShaderResponse struct {
    Error    string       `json:"error"`
}

func RouteShaderDataToModel(data *RouteShaderData, model *RouteShaderModel) {
    model.Id = types.Int64Value(int64(data.RouteShaderId))
    model.Name = types.StringValue(data.RouteShaderName)
    model.ABTest = types.BoolValue(data.ABTest)
    model.AcceptableLatency = types.Int64Value(int64(data.AcceptableLatency))
    model.AcceptablePacketLossInstant = types.Float64Value(float64(data.AcceptablePacketLossInstant))
    model.AcceptablePacketLossSustained = types.Float64Value(float64(data.AcceptablePacketLossSustained))
    model.AnalysisOnly = types.BoolValue(data.AnalysisOnly)
    model.BandwidthEnvelopeUpKbps = types.Int64Value(int64(data.BandwidthEnvelopeUpKbps))
    model.BandwidthEnvelopeDownKbps = types.Int64Value(int64(data.BandwidthEnvelopeDownKbps))
    model.DisableNetworkNext = types.BoolValue(data.DisableNetworkNext)
    model.LatencyReductionThreshold = types.Int64Value(int64(data.LatencyReductionThreshold))
    model.Multipath = types.BoolValue(data.Multipath)
    model.SelectionPercent = types.Float64Value(float64(data.SelectionPercent))
    model.MaxLatencyTradeOff = types.Int64Value(int64(data.MaxLatencyTradeOff))
    model.MaxNextRTT = types.Int64Value(int64(data.MaxNextRTT))
    model.RouteSwitchThreshold = types.Int64Value(int64(data.RouteSwitchThreshold))
    model.RouteSelectThreshold = types.Int64Value(int64(data.RouteSelectThreshold))
    model.RTTVeto = types.Int64Value(int64(data.RTTVeto))
    model.ForceNext = types.BoolValue(data.ForceNext)
    model.RouteDiversity = types.Int64Value(int64(data.RouteDiversity))
}

func RouteShaderModelToData(model *RouteShaderModel, data *RouteShaderData) {
    data.RouteShaderId = uint64(model.Id.ValueInt64())
    data.RouteShaderName = model.Name.ValueString()
    data.ABTest = model.ABTest.ValueBool()
    data.AcceptableLatency = int(model.AcceptableLatency.ValueInt64())
    data.AcceptablePacketLossInstant = model.AcceptablePacketLossInstant.ValueFloat64()
    data.AcceptablePacketLossSustained = model.AcceptablePacketLossSustained.ValueFloat64()
    data.AnalysisOnly = model.AnalysisOnly.ValueBool()
    data.BandwidthEnvelopeUpKbps = int(model.BandwidthEnvelopeUpKbps.ValueInt64())
    data.BandwidthEnvelopeDownKbps = int(model.BandwidthEnvelopeDownKbps.ValueInt64())
    data.DisableNetworkNext = model.DisableNetworkNext.ValueBool()
    data.LatencyReductionThreshold = int(model.LatencyReductionThreshold.ValueInt64())
    data.Multipath = model.Multipath.ValueBool()
    data.SelectionPercent = model.SelectionPercent.ValueFloat64()
    data.MaxLatencyTradeOff = int(model.MaxLatencyTradeOff.ValueInt64())
    data.MaxNextRTT = int(model.MaxNextRTT.ValueInt64())
    data.RouteSwitchThreshold = int(model.RouteSwitchThreshold.ValueInt64())
    data.RouteSelectThreshold = int(model.RouteSelectThreshold.ValueInt64())
    data.RTTVeto = int(model.RTTVeto.ValueInt64())
    data.ForceNext = model.ForceNext.ValueBool()
    data.RouteDiversity = int(model.RouteDiversity.ValueInt64())
}

func RouteShaderSchema() schema.Schema {
    return schema.Schema{
        Description: "Manages a route shader.",
        Attributes: map[string]schema.Attribute{
            "id": schema.Int64Attribute{
                Description: "The id of the route shader. Automatically generated when route shaders are created.",
                Computed: true,
                PlanModifiers: []planmodifier.Int64{
                    int64planmodifier.UseStateForUnknown(),
                },
            },
            "name": schema.StringAttribute{
                Description: "The name of the route shader. Generally, it's best to set it to the same name as the buyer that will use it. Typically, there is one route shader per-buyer.",
                Required: true,
            },
            "ab_test": schema.BoolAttribute{
                Description: "If true then AB test mode is enabled. 50%% of users will be eligible for acceleration, and 50%% will not be eligible for acceleration. In the accelerated group, players are only accelerated if the remaining route shader parameters determine they should be accelerated.",
                Optional: true,
                Computed: true,
                Default: booldefault.StaticBool(false),
            },
            "acceptable_latency": schema.Int64Attribute{
                Description: "The amount of latency that is acceptable in milliseconds. Any latency above this is eligible for acceleration. For example, setting this value to 50ms would exclude any players with latency < 50ms from being accelerated. Default is 20ms",
                Optional: true,
                Computed: true,
                Default: int64default.StaticInt64(20),
            },
            "acceptable_packet_loss_instant": schema.Float64Attribute{
                Description: "The instantaneous packet loss that is acceptable. For example, setting to 1%% will allow packet loss up to 1%% in a 10 second period before enabling acceleration. Default is 0.25%%",
                Optional: true,
                Computed: true,
                Default: float64default.StaticFloat64(0.25),
            },
            "acceptable_packet_loss_sustained": schema.Float64Attribute{
                Description: "The sustained packet loss that is acceptable. For example, setting to 0.1%% will allow packet loss up to 0.1%% in a 30 second period before enabling acceleration. Default is 0.1%%",
                Optional: true,
                Computed: true,
                Default: float64default.StaticFloat64(0.1),
            },
            "analysis_only": schema.BoolAttribute{
                Description: "Set this to true and acceleration is disabled. Analytics data will still be gathered from players of any buyer who use this route shader. Use this when you want to gather network perforance data for players, but you want accelerated turned off. Default is false.",
                Optional: true,
                Computed: true,
                Default: booldefault.StaticBool(false),
            },
            "bandwidth_envelope_up_kbps": schema.Int64Attribute{
                Description: "The maximum amount of acceleration bandwidth up in the client to server direction in kilobits per-second (kbps). Exceeding this amount of bandwidth for a session will result in packets above the limit not being accelerated. Default value is 1024 (1mbps).",
                Optional: true,
                Computed: true,
                Default: int64default.StaticInt64(1024),
            },
            "bandwidth_envelope_down_kbps": schema.Int64Attribute{
                Description: "The maximum amount of acceleration bandwidth down in the server to client direction in kilobits per-second (kbps). Exceeding this amount of bandwidth for a session will result in packets above the limit not being accelerated. Default value is 1024 (1mbps).",
                Optional: true,
                Computed: true,
                Default: int64default.StaticInt64(1024),
            },
            "disable_network_next": schema.BoolAttribute{
                Description: "Set this to true and network next is completely disabled for players belonging to any buyer using this route shader. No acceleration will be applied and no data will be collected.",
                Optional: true,
                Computed: true,
                Default: booldefault.StaticBool(false),
            },
            "latency_reduction_threshold": schema.Int64Attribute{
                Description: "The minimum latency reduction threshold in milliseconds. A latency reduction greater than or equal to this number in milliseconds must be predicted in order to accelerate the player. For example, 10ms would only accelerate players if a latency reduction >= 10ms is predicted. Default is 10ms.",
                Optional: true,
                Computed: true,
                Default: int64default.StaticInt64(10),
            },
            "multipath": schema.BoolAttribute{
                Description: "If this is set to true then packets will be sent across the direct route (default internet path), as well as the network next path when a player is accelerated. Recommend setting this to true always. Default value is true.",
                Optional: true,
                Computed: true,
                Default: booldefault.StaticBool(true),
            },
            "selection_percent": schema.Float64Attribute{
                Description: "The percentage of players eligible for acceleration. For example, setting this to 10%% would only accelerate 1 in 10 players. Default value is 100%% (all players eligible for acceleration).",
                Optional: true,
                Computed: true,
                Default: float64default.StaticFloat64(100.0),
            },
            "max_latency_trade_off": schema.Int64Attribute{
                Description: "The maximum amount of latency to trade for reduced packet loss. Packet loss above a certain amount is generally considered worse than latency. Trade up to this amount of latency to take a route with reduced packet loss. IMPORTANT: Network Next will still prefer the lowest latency route available. Default value is 20ms.",
                Optional: true,
                Computed: true,
                Default: int64default.StaticInt64(20),
            },
            "max_next_rtt": schema.Int64Attribute{
                Description: "The maximum network next latency allowed. There is no point patting yourself on the back when you reduce 500ms latency to 450ms. 450ms is still completely unplayable! Set this value to the maximum reasonable accelerated latency that makes sense for your game. Network Next will not accelerate a player if their predicted post-acceleration latency is higher than this value. Default value is 250ms.",
                Optional: true,
                Computed: true,
                Default: int64default.StaticInt64(250),
            },
            "route_switch_threshold": schema.Int64Attribute{
                Description: "If a player is already being accelerated, and a better route is available with a reduction of at least this much milliseconds from the current route, then switch to it. Don't set this too low or sessions will switch routes every 10 seconds due to natural fluctuation. Default value is 10ms.",
                Optional: true,
                Computed: true,
                Default: int64default.StaticInt64(10),
            },
            "route_select_threshold": schema.Int64Attribute{
                Description: "When initially selecting a route, find routes within this amount of milliseconds of the best available route and consider them to be equal, randomly selecting between them. Don't set this too low, or you won't get effective load balancing across relays. Default value is 5ms.",
                Optional: true,
                Computed: true,
                Default: int64default.StaticInt64(5),
            },
            "rtt_veto": schema.Int64Attribute{
                Description: "If the accelerated latency becomes worse than the direct latency (default internet route) by this amount of milliseconds, then veto the session and stop accelerating the player. This is to avoid players hopping on/off of network next when latency fluctuation occurs. Don't set this too low, or players with fluctuating latency (due to edge issues or wifi) will fall off network next due to temporary high latency conditions. Default value is 20ms.",
                Optional: true,
                Computed: true,
                Default: int64default.StaticInt64(20),
            },
            "force_next": schema.BoolAttribute{
                Description: "Force all players to be accelerated. This is useful for testing, or if you have a small professional player base for whom you want to enable acceleration BEFORE anything goes wrong, instead of being reactive and only accelerating them after something goes wrong. Default is false.",
                Optional: true,
                Computed: true,
                Default: booldefault.StaticBool(false),
            },
            "route_diversity": schema.Int64Attribute{
                Description: "The minimum amount of distinct viable routes that must be available for a player to be accelerated. This setting can be useful to limit players in remote regions from taking network next when there is only one relay available to them. Only players with mulitple distinct paths will be accelerated, so there are backups if one relay becomse unroutable for this player. Don't set this too high or players with bad edge network conditions won't get accelerated. Default value is 0.",
                Optional: true,
                Computed: true,
                Default: int64default.StaticInt64(0),
            },
        },
    }
}

func RouteShadersSchema() datasource_schema.Schema {
    return datasource_schema.Schema{
        Description: "Fetches the list of route shaders.",
        Attributes: map[string]datasource_schema.Attribute{
            "route_shaders": datasource_schema.ListNestedAttribute{
                Computed: true,
                NestedObject: datasource_schema.NestedAttributeObject{
                    Attributes: map[string]datasource_schema.Attribute{
                        "id": datasource_schema.Int64Attribute{
                            Description: "The id of the route shader. Automatically generated when route shaders are created.",
                            Computed: true,
                        },
                        "name": datasource_schema.StringAttribute{
                            Description: "The name of the route shader. Generally, it's best to set it to the same name as the buyer that will use it. Typically, there is one route shader per-buyer.",
                            Required: true,
                        },
                        "ab_test": datasource_schema.BoolAttribute{
                            Description: "If true then AB test mode is enabled. 50%% of users will be eligible for acceleration, and 50%% will not be eligible for acceleration. In the accelerated group, players are only accelerated if the remaining route shader parameters determine they should be accelerated.",
                            Required: true,
                        },
                        "acceptable_latency": datasource_schema.Int64Attribute{
                            Description: "The amount of latency that is acceptable in milliseconds. Any latency above this is eligible for acceleration. For example, setting this value to 50ms would exclude any players with latency < 50ms from being accelerated. Default is 20ms",
                            Required: true,
                        },
                        "acceptable_packet_loss_instant": datasource_schema.Float64Attribute{
                            Description: "The instantaneous packet loss that is acceptable. For example, setting to 1%% will allow packet loss up to 1%% in a 10 second period before enabling acceleration. Default is 0.25%%",
                            Required: true,
                        },
                        "acceptable_packet_loss_sustained": datasource_schema.Float64Attribute{
                            Description: "The sustained packet loss that is acceptable. For example, setting to 0.1%% will allow packet loss up to 0.1%% in a 30 second period before enabling acceleration. Default is 0.1%%",
                            Required: true,
                        },
                        "analysis_only": datasource_schema.BoolAttribute{
                            Description: "Set this to true and acceleration is disabled. Analytics data will still be gathered from players of any buyer who use this route shader. Use this when you want to gather network perforance data for players, but you want accelerated turned off. Default is false.",
                            Required: true,
                        },
                        "bandwidth_envelope_up_kbps": datasource_schema.Int64Attribute{
                            Description: "The maximum amount of acceleration bandwidth up in the client to server direction in kilobits per-second (kbps). Exceeding this amount of bandwidth for a session will result in packets above the limit not being accelerated. Default value is 1024 (1mbps).",
                            Required: true,
                        },
                        "bandwidth_envelope_down_kbps": datasource_schema.Int64Attribute{
                            Description: "The maximum amount of acceleration bandwidth down in the server to client direction in kilobits per-second (kbps). Exceeding this amount of bandwidth for a session will result in packets above the limit not being accelerated. Default value is 1024 (1mbps).",
                            Required: true,
                        },
                        "disable_network_next": datasource_schema.BoolAttribute{
                            Description: "Set this to true and network next is completely disabled for players belonging to any buyer using this route shader. No acceleration will be applied and no data will be collected.",
                            Required: true,
                        },
                        "latency_reduction_threshold": datasource_schema.Int64Attribute{
                            Description: "The minimum latency reduction threshold in milliseconds. A latency reduction greater than or equal to this number in milliseconds must be predicted in order to accelerate the player. For example, 10ms would only accelerate players if a latency reduction >= 10ms is predicted. Default is 10ms.",
                            Required: true,
                        },
                        "multipath": datasource_schema.BoolAttribute{
                            Description: "If this is set to true then packets will be sent across the direct route (default internet path), as well as the network next path when a player is accelerated. Recommend setting this to true always. Default value is true.",
                            Required: true,
                        },
                        "selection_percent": datasource_schema.Float64Attribute{
                            Description: "The percentage of players eligible for acceleration. For example, setting this to 10%% would only accelerate 1 in 10 players. Default value is 100%% (all players eligible for acceleration).",
                            Required: true,
                        },
                        "max_latency_trade_off": datasource_schema.Int64Attribute{
                            Description: "The maximum amount of latency to trade for reduced packet loss. Packet loss above a certain amount is generally considered worse than latency. Trade up to this amount of latency to take a route with reduced packet loss. IMPORTANT: Network Next will still prefer the lowest latency route available. Default value is 20ms.",
                            Required: true,
                        },
                        "max_next_rtt": datasource_schema.Int64Attribute{
                            Description: "The maximum network next latency allowed. There is no point patting yourself on the back when you reduce 500ms latency to 450ms. 450ms is still completely unplayable! Set this value to the maximum reasonable accelerated latency that makes sense for your game. Network Next will not accelerate a player if their predicted post-acceleration latency is higher than this value. Default value is 250ms.",
                            Required: true,
                        },
                        "route_switch_threshold": datasource_schema.Int64Attribute{
                            Description: "If a player is already being accelerated, and a better route is available with a reduction of at least this much milliseconds from the current route, then switch to it. Don't set this too low or sessions will switch routes every 10 seconds due to natural fluctuation. Default value is 10ms.",
                            Required: true,
                        },
                        "route_select_threshold": datasource_schema.Int64Attribute{
                            Description: "When initially selecting a route, find routes within this amount of milliseconds of the best available route and consider them to be equal, randomly selecting between them. Don't set this too low, or you won't get effective load balancing across relays. Default value is 5ms.",
                            Required: true,
                        },
                        "rtt_veto": datasource_schema.Int64Attribute{
                            Description: "If the accelerated latency becomes worse than the direct latency (default internet route) by this amount of milliseconds, then veto the session and stop accelerating the player. This is to avoid players hopping on/off of network next when latency fluctuation occurs. Don't set this too low, or players with fluctuating latency (due to edge issues or wifi) will fall off network next due to temporary high latency conditions. Default value is 20ms.",
                            Required: true,
                        },
                        "force_next": datasource_schema.BoolAttribute{
                            Description: "Force all players to be accelerated. This is useful for testing, or if you have a small professional player base for whom you want to enable acceleration BEFORE anything goes wrong, instead of being reactive and only accelerating them after something goes wrong. Default is false.",
                            Required: true,
                        },
                        "route_diversity": datasource_schema.Int64Attribute{
                            Description: "The minimum amount of distinct viable routes that must be available for a player to be accelerated. This setting can be useful to limit players in remote regions from taking network next when there is only one relay available to them. Only players with mulitple distinct paths will be accelerated, so there are backups if one relay becomse unroutable for this player. Don't set this too high or players with bad edge network conditions won't get accelerated. Default value is 0.",
                            Required: true,
                        },
                    },
                },
            },
        },
    }
}

// -------------------------------------------------------------------

type BuyerDatacenterSettingsData struct {
    BuyerId            uint64 `json:"buyer_id"`
    DatacenterId       uint64 `json:"datacenter_id"`
    EnableAcceleration bool   `json:"enable_acceleration"`
}

type BuyerDatacenterSettingsModel struct {
    BuyerId                     types.Int64  `tfsdk:"buyer_id"`
    DatacenterId                types.Int64  `tfsdk:"datacenter_id"`
    EnableAcceleration          types.Bool   `tfsdk:"enable_acceleration"`
}

type BuyerDatacenterSettingsListModel struct {
    Settings []BuyerDatacenterSettingsModel `tfsdk:"settings"`
}

type CreateBuyerDatacenterSettingsResponse struct {
    Settings BuyerDatacenterSettingsData    `json:"settings"`
    Error    string                         `json:"error"`
}

type ReadBuyerDatacenterSettingsResponse struct {
    Settings BuyerDatacenterSettingsData    `json:"settings"`
    Error    string                         `json:"error"`
}

type ReadBuyerDatacenterSettingsListResponse struct {
    Settings []BuyerDatacenterSettingsData  `json:"settings"`
    Error    string                         `json:"error"`
}

type UpdateBuyerDatacenterSettingsResponse struct {
    Settings BuyerDatacenterSettingsData    `json:"settings"`
    Error    string                         `json:"error"`
}

type DeleteBuyerDatacenterSettingsResponse struct {
    Error    string                         `json:"error"`
}

func BuyerDatacenterSettingsModelToData(model *BuyerDatacenterSettingsModel, data *BuyerDatacenterSettingsData) {
    data.BuyerId = uint64(model.BuyerId.ValueInt64())
    data.DatacenterId = uint64(model.DatacenterId.ValueInt64())
    data.EnableAcceleration = model.EnableAcceleration.ValueBool()    
}

func BuyerDatacenterSettingsDataToModel(data *BuyerDatacenterSettingsData, model *BuyerDatacenterSettingsModel) {
    model.BuyerId = types.Int64Value(int64(data.BuyerId))
    model.DatacenterId = types.Int64Value(int64(data.DatacenterId))
    model.EnableAcceleration = types.BoolValue(data.EnableAcceleration)
}

func BuyerDatacenterSettingsSchema() schema.Schema {
    return schema.Schema{
        Description: "Manages buyer datacenter settings.",
        Attributes: map[string]schema.Attribute{
            "buyer_id": schema.Int64Attribute{
                Description: "The buyer id this setting belongs to.",
                Required: true,
            },
            "datacenter_id": schema.Int64Attribute{
                Description: "The datacenter id this setting applies to.",
                Required: true,
            },
            "enable_acceleration": schema.BoolAttribute{
                Description: "Set to true to enable acceleration to this datacenter from the buyer",
                Required: true,
            },
        },
    }
}

func BuyerDatacenterSettingsListSchema() datasource_schema.Schema {
    return datasource_schema.Schema{
        Description: "Fetches the list of buyer datacenter settings.",
        Attributes: map[string]datasource_schema.Attribute{
            "settings": datasource_schema.ListNestedAttribute{
                Computed: true,
                NestedObject: datasource_schema.NestedAttributeObject{
                    Attributes: map[string]datasource_schema.Attribute{
                        "buyer_id": datasource_schema.Int64Attribute{
                            Description: "The buyer id this setting belongs to.",
                            Required: true,
                        },
                        "datacenter_id": datasource_schema.Int64Attribute{
                            Description: "The datacenter id this setting applies to.",
                            Required: true,
                        },
                        "enable_acceleration": datasource_schema.BoolAttribute{
                            Description: "Set to true to enable acceleration to this datacenter from the buyer",
                            Required: true,
                        },
                    },
                },
            },
        },
    }
}

// -------------------------------------------------------------------

type RelayKeypairData struct {
    RelayKeypairId      uint64                  `json:"relay_keypair_id"`
    PublicKeyBase64     string                  `json:"public_key_base64"`
    PrivateKeyBase64    string                  `json:"private_key_base64"`
}

type RelayKeypairModel struct {
    Id                  types.Int64             `tfsdk:"id"`
    PublicKeyBase64     types.String            `tfsdk:"public_key_base64"`
    PrivateKeyBase64    types.String            `tfsdk:"private_key_base64"`
}

type RelayKeypairsModel struct {
    RelayKeypairs       []RelayKeypairModel     `tfsdk:"relay_keypairs"`
}

type CreateRelayKeypairResponse struct {
    RelayKeypair        RelayKeypairData        `json:"relay_keypair"`
    Error               string                  `json:"error"`
}

type ReadRelayKeypairResponse struct {
    RelayKeypair        RelayKeypairData        `json:"relay_keypair"`
    Error               string                  `json:"error"`
}

type ReadRelayKeypairsResponse struct {
    RelayKeypairs       []RelayKeypairData      `json:"relay_keypairs"`
    Error               string                  `json:"error"`
}

type UpdateRelayKeypairResponse struct {
    RelayKeypair        RelayKeypairData        `json:"relay_keypair"`
    Error               string                  `json:"error"`
}

type DeleteRelayKeypairResponse struct {
    Error               string                  `json:"error"`
}

func RelayKeypairModelToData(model *RelayKeypairModel, data *RelayKeypairData) {
    data.RelayKeypairId = uint64(model.Id.ValueInt64())
    data.PublicKeyBase64 = model.PublicKeyBase64.ValueString()
    data.PrivateKeyBase64 = model.PrivateKeyBase64.ValueString()
}

func RelayKeypairDataToModel(data *RelayKeypairData, model *RelayKeypairModel) {
    model.Id = types.Int64Value(int64(data.RelayKeypairId))
    model.PublicKeyBase64 = types.StringValue(data.PublicKeyBase64)
    model.PrivateKeyBase64 = types.StringValue(data.PrivateKeyBase64)
}

func RelayKeypairSchema() schema.Schema {
    return schema.Schema{
        Description: "Manages a relay keypair.",
        Attributes: map[string]schema.Attribute{
            "id": schema.Int64Attribute{
                Description: "The id of the relay keypair. Automatically generated when relay keypairs are created.",
                Computed: true,
                PlanModifiers: []planmodifier.Int64{
                    int64planmodifier.UseStateForUnknown(),
                },
            },
            "public_key_base64": schema.StringAttribute{
                Description: "The generated public key as a base64 encoded string.",
                Computed: true,
            },
            "private_key_base64": schema.StringAttribute{
                Description: "The generated private key as a base64 encoded string.",
                Computed: true,
            },
        },
    }
}

func RelayKeypairsSchema() datasource_schema.Schema {
    return datasource_schema.Schema{
        Description: "Fetches the list of relay keypairs.",
        Attributes: map[string]datasource_schema.Attribute{
            "relay_keypairs": datasource_schema.ListNestedAttribute{
                Computed: true,
                NestedObject: datasource_schema.NestedAttributeObject{
                    Attributes: map[string]datasource_schema.Attribute{
                        "id": datasource_schema.Int64Attribute{
                            Description: "The id of the relay keypair. Automatically generated when relay keypairs are created.",
                            Computed: true,
                        },
                        "public_key_base64": schema.StringAttribute{
                            Description: "The generated public key as a base64 encoded string.",
                            Computed: true,
                        },
                        "private_key_base64": schema.StringAttribute{
                            Description: "The generated private key as a base64 encoded string.",
                            Computed: true,
                        },
                    },
                },
            },
        },
    }
}

// -------------------------------------------------------------------
