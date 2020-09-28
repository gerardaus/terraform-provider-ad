package main

import (
	"github.com/gerardaus/terraform-provider-ad/ad"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ad.Provider})
}
