package networknext

import (
    "context"

    "github.com/hashicorp/terraform-plugin-framework/datasource"
    "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
    "github.com/hashicorp/terraform-plugin-framework/types"
)

var (
    _ datasource.DataSource              = &customersDataSource{}
    _ datasource.DataSourceWithConfigure = &customersDataSource{}
)

func NewCustomersDataSource() datasource.DataSource {
    return &customersDataSource{}
}

type customersDataSource struct {
    client *Client
}

type customersDataSourceModel struct {
    Customers []customersModel `tfsdk:"customers"`
}

type customersModel struct {
    ID          types.Int64               `tfsdk:"id"`
    Name        types.String              `tfsdk:"name"`
    Code        types.String              `tfsdk:"code"`
    Live        types.Bool                `tfsdk:"live"`
    Debug       types.Bool                `tfsdk:"debug"`
}

func (d *customersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_customers"
}

func (d *customersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
    resp.Schema = schema.Schema{
        Attributes: map[string]schema.Attribute{
            "customers": schema.ListNestedAttribute{
                Computed: true,
                NestedObject: schema.NestedAttributeObject{
                    Attributes: map[string]schema.Attribute{
                        "id": schema.Int64Attribute{
                            Computed: true,
                        },
                        "name": schema.StringAttribute{
                            Computed: true,
                        },
                        "code": schema.StringAttribute{
                            Computed: true,
                        },
                        "live": schema.BoolAttribute{
                            Computed: true,
                        },
                        "debug": schema.BoolAttribute{
                            Computed: true,
                        },
                    },
                },
            },
        },
    }
}

func (d *customersDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
    
    if req.ProviderData == nil {
        return
    }

    d.client = req.ProviderData.(*Client)
}

func (d *customersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

    /*
    var state customersDataSourceModel

    customers, err := d.client.GetCustomers()
    if err != nil {
        resp.Diagnostics.AddError(
            "Unable to read networknext customers",
            err.Error(),
        )
        return
    }

    // Map response body to model
    for _, coffee := range coffees {
        coffeeState := coffeesModel{
            ID:          types.Int64Value(int64(coffee.ID)),
            Name:        types.StringValue(coffee.Name),
            Teaser:      types.StringValue(coffee.Teaser),
            Description: types.StringValue(coffee.Description),
            Price:       types.Float64Value(coffee.Price),
            Image:       types.StringValue(coffee.Image),
        }

        for _, ingredient := range coffee.Ingredient {
            coffeeState.Ingredients = append(coffeeState.Ingredients, coffeesIngredientsModel{
                ID: types.Int64Value(int64(ingredient.ID)),
            })
        }

        state.Coffees = append(state.Coffees, coffeeState)
    }

    // Set state
    diags := resp.State.Set(ctx, &state)
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
    */
}
