package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

//go:generate go install github.com/protomesh/protoc-gen-terraform@0ee4386b0252eb9a1aceaf0bffdd946db8044171
//go:generate buf generate buf.build/protomesh/protomesh
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name protomesh

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: Provider,
	})
}
