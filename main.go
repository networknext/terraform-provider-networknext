package main

import (
    "context"
    "terraform-provider-networknext/networknext"

    "github.com/hashicorp/terraform-plugin-framework/providerserver"
)

// Provider documentation generation.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name networknext

func main() {
    providerserver.Serve(context.Background(), networknext.New, providerserver.ServeOpts{
        Address: "networknext/networknext",
    })
}
