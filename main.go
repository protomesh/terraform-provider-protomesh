package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

//go:generate go install github.com/protomesh/protoc-gen-terraform@7b91d2ba39fd819bc8791c5b8b0bccdeb55b3d64
//go:generate buf generate buf.build/protomesh/protomesh
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name protomesh

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: Provider,
	})
}
