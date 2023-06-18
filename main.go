package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

//go:generate go install github.com/protomesh/protoc-gen-terraform@3789b8fe76911d0f8155642c09e048eeef6efda0
//go:generate buf generate buf.build/protomesh/protomesh
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name protomesh

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: Provider,
	})
}
