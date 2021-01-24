package main

import (
  "github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
	"provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: Provider,
	})
}
