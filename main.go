package main

import (
    "context"
    "terraform-provider-networknext/accelerate"

    "github.com/hashicorp/terraform-plugin-framework/providerserver"
)

// Provider documentation generation.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name accelerate

func main() {
    providerserver.Serve(context.Background(), accelerate.New, providerserver.ServeOpts{
        Address: "networknext/accelerate",
    })
}
