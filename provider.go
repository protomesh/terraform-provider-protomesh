package main

import (
	"context"
	"errors"
	"os"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/protomesh/go-app"
	"github.com/protomesh/protomesh/pkg/client"
	"github.com/protomesh/protomesh/pkg/config"
	"github.com/protomesh/protomesh/provider/tls"
	servicesv1 "github.com/protomesh/terraform-provider-protomesh/proto/api/services/v1"
	"github.com/protomesh/terraform-provider-protomesh/resources"
)

func kvIntoMap(kv []interface{}) map[string]interface{} {

	m := make(map[string]interface{})

	var key string

	for i, v := range kv {

		if i%2 == 0 {
			key = v.(string)
			continue
		}

		m[key] = v

	}

	return m

}

type providerLogger struct {
	ctx  context.Context
	name string
	kv   []interface{}
}

func (p *providerLogger) addNameToKv(kv []interface{}) []interface{} {
	kv = append(kv, p.kv...)
	if len(p.name) > 0 {
		kv = append(kv, "name", p.name)
	}
	return kv
}

func (p *providerLogger) Debug(message string, kv ...interface{}) {
	tflog.Debug(p.ctx, message, kvIntoMap(p.addNameToKv(kv)))
}

func (p *providerLogger) Info(message string, kv ...interface{}) {
	tflog.Info(p.ctx, message, kvIntoMap(p.addNameToKv(kv)))
}

func (p *providerLogger) Warn(message string, kv ...interface{}) {
	tflog.Warn(p.ctx, message, kvIntoMap(p.addNameToKv(kv)))
}

func (p *providerLogger) Error(message string, kv ...interface{}) {
	tflog.Error(p.ctx, message, kvIntoMap(p.addNameToKv(kv)))
}

func (p *providerLogger) Panic(message string, kv ...interface{}) {
	tflog.Error(p.ctx, message, kvIntoMap(p.addNameToKv(kv)))
	os.Exit(1)
}

func (p *providerLogger) With(kv ...interface{}) app.Logger {
	return &providerLogger{
		ctx:  p.ctx,
		name: p.name,
		kv:   kv,
	}
}

type providerApp struct {
	log *providerLogger
}

func (a *providerApp) Log() app.Logger {
	return a.log
}

func (a *providerApp) Close() {
}

type providerDeps struct {
	*app.Injector[*providerDeps]

	GrpcClient *client.GrpcClient[any]
}

func (d *providerDeps) Dependency() *providerDeps {
	return d
}

func (d *providerDeps) GetResourceStoreClient() servicesv1.ResourceStoreClient {
	return servicesv1.NewResourceStoreClient(d.GrpcClient.ClientConn)
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": {
				Description: "Address of the Protomesh resource store gRPC server to manage resources. Default to localhost:6680.",
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PROTOMESH_STORE_ADDRESS", "localhost:6680"),
			},
			"enable_tls": {
				Description: "Enable in the Protomesh resource store connection, default **false**.",
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PROTOMESH_ENABLE_TLS", false),
			},
			"server_name_override": {
				Description: "Override server name in the TLS handshake, default nil.",
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PROTOMESH_STORE_SERVER_NAME_OVERRIDE", nil),
			},
			"tls_certificate_path": {
				Description: "Client TLS certificate path default nil.",
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PROTOMESH_TLS_CERTIFICATE_PATH", nil),
			},
			"tls_private_key_path": {
				Description: "Client TLS private key path default nil.",
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PROTOMESH_TLS_PRIVATE_KEY_PATH", nil),
			},
		},
		ConfigureContextFunc: func(ctx context.Context, rd *schema.ResourceData) (interface{}, diag.Diagnostics) {

			tflog.Info(ctx, "Configuring Protomesh resource store client")

			address, ok := rd.Get("address").(string)
			if !ok {
				return nil, diag.FromErr(errors.New("Invalid address type provided"))
			}

			enableTls, ok := rd.Get("enable_tls").(bool)
			if !ok {
				return nil, diag.FromErr(errors.New("Invalid enable_tls type provided."))
			}

			deps := &providerDeps{
				GrpcClient: &client.GrpcClient[any]{
					EnableTls: config.NewConfig(enableTls),
					TlsBuilder: &tls.CertificateLoader[any]{
						CertificatePath: config.EmptyConfig(),
						PrivateKey: &tls.KeyLoader[any]{
							KeysPath: config.EmptyConfig(),
						},
					},
					ServerAddress:      config.NewConfig(address),
					ServerNameOverride: config.EmptyConfig(),
				},
			}

			serverNameOverride := rd.Get("server_name_override")
			if serverNameOverride != nil {
				deps.GrpcClient.ServerNameOverride = config.NewConfig(serverNameOverride)
			}

			tlsCertPath := rd.Get("tls_certificate_path")
			if tlsCertPath != nil {
				deps.GrpcClient.TlsBuilder.CertificatePath = config.NewConfig(tlsCertPath)
			}

			tlsPrivPath := rd.Get("tls_private_key_path")
			if tlsPrivPath != nil {
				deps.GrpcClient.TlsBuilder.PrivateKey.KeysPath = config.NewConfig(tlsPrivPath)
			}

			provApp := &providerApp{
				log: &providerLogger{
					ctx:  ctx,
					name: "protomesh",
					kv:   make([]interface{}, 0),
				},
			}

			app.Inject(provApp, deps)

			deps.GrpcClient.Start()

			return deps, nil
		},
		ResourcesMap: map[string]*schema.Resource{
			"protomesh_gateway_policy":  resources.ResourceGatewayPolicy(),
			"protomesh_service":         resources.ResourceService(),
			"protomesh_http_ingress":    resources.ResourceHttpIngress(),
			"protomesh_instance_set":    resources.ResourceInstanceSet(),
			"protomesh_routing_policy":  resources.ResourceRoutingPolicy(),
			"protomesh_process_trigger": resources.ResourceProcessTrigger(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"protomesh_gateway_policy":  resources.DataSourceGatewayPolicy(),
			"protomesh_service":         resources.DataSourceService(),
			"protomesh_http_ingress":    resources.DataSourceHttpIngress(),
			"protomesh_instance_set":    resources.DataSourceInstanceSet(),
			"protomesh_routing_policy":  resources.DataSourceRoutingPolicy(),
			"protomesh_process_trigger": resources.DataSourceProcessTrigger(),
		},
	}
}
