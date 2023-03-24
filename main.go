package main

import (
    "context"
    "terraform-provider-networknext/networknext"

    "github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
    providerserver.Serve(context.Background(), networknext.New, providerserver.ServeOpts{
        Address: "hashicorp.com/edu/networknext",
    })
}
