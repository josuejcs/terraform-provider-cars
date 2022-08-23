package main

import (
	"terra-cars/myprovider"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: myprovider.Provider,
	})
}
