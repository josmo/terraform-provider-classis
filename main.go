package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/classis/terraform-provider-classis/classis"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: classis.Provider})
}
