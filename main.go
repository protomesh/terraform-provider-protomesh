package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

//go:generate go install github.com/protomesh/protoc-gen-terraform@8e31675a2cb0afe68925785b125bbe1ea10ac303
//go:generate buf generate buf.build/protomesh/protomesh
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name protomesh

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: Provider,
	})
}
