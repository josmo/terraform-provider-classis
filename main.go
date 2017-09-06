package main

import (
	"github.com/classis/terraform-provider-classis/classis"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: classis.Provider})
}
