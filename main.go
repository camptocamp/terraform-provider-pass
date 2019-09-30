package main

/* Bootstrap the plugin for Pass */

import (
	"github.com/camptocamp/terraform-provider-pass/pass"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: pass.Provider,
	})
}
