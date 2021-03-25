package main

/* Bootstrap the plugin for Pass */

import (
	"github.com/camptocamp/terraform-provider-pass/pass"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: pass.Provider,
	})
}
