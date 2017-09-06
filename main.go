package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/classis/terraform-classis-provider/classis"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: classis.Provider})
}
