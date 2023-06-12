---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "protomesh_http_ingress Resource - protomesh"
subcategory: ""
description: |-
  Expose a AWS Lambda through a gRPC interface
---

# protomesh_http_ingress (Resource)

Expose a AWS Lambda through a gRPC interface



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the resource (display name)
- `namespace` (String) Namespace in the resource store
- `node` (Block List, Min: 1, Max: 1) Node specification. (see [below for nested schema](#nestedblock--node))
- `resource_id` (String) A valid UUID for the resource

### Read-Only

- `id` (String) The ID of this resource.
- `version_index` (Number) Resource version index (unix timestamp in seconds)
- `version_sha256_sum` (String) Current version sha256 sum

<a id="nestedblock--node"></a>
### Nested Schema for `node`

Optional:

- `http_filters` (Block List) Http filters to apply to the ingress listener (see [below for nested schema](#nestedblock--node--http_filters))
- `ingress_name` (String) Ingress name (used as route config name for the route specifier)
- `listen_port` (Number) Ingress port to listen for incoming requests
- `xds_cluster_name` (String) XDS cluster name: must be the same of the envoy config to be matched by  xDS server

<a id="nestedblock--node--http_filters"></a>
### Nested Schema for `node.http_filters`

Optional:

- `filter` (Block List, Max: 1) (see [below for nested schema](#nestedblock--node--http_filters--filter))

<a id="nestedblock--node--http_filters--filter"></a>
### Nested Schema for `node.http_filters.filter`

Optional:

- `cors` (Block List, Max: 1) (see [below for nested schema](#nestedblock--node--http_filters--filter--cors))
- `grpc_web` (Block List, Max: 1) (see [below for nested schema](#nestedblock--node--http_filters--filter--grpc_web))
- `health_check` (Block List, Max: 1) (see [below for nested schema](#nestedblock--node--http_filters--filter--health_check))
- `jwt_authn` (Block List, Max: 1) (see [below for nested schema](#nestedblock--node--http_filters--filter--jwt_authn))

<a id="nestedblock--node--http_filters--filter--cors"></a>
### Nested Schema for `node.http_filters.filter.cors`


<a id="nestedblock--node--http_filters--filter--grpc_web"></a>
### Nested Schema for `node.http_filters.filter.grpc_web`


<a id="nestedblock--node--http_filters--filter--health_check"></a>
### Nested Schema for `node.http_filters.filter.health_check`

Required:

- `path` (String)


<a id="nestedblock--node--http_filters--filter--jwt_authn"></a>
### Nested Schema for `node.http_filters.filter.jwt_authn`

Optional:

- `providers` (Block List) (see [below for nested schema](#nestedblock--node--http_filters--filter--jwt_authn--providers))
- `rules` (Block List) (see [below for nested schema](#nestedblock--node--http_filters--filter--jwt_authn--rules))

<a id="nestedblock--node--http_filters--filter--jwt_authn--providers"></a>
### Nested Schema for `node.http_filters.filter.jwt_authn.providers`

Required:

- `forward` (Boolean)
- `issuer` (String)
- `provider_name` (String)

Optional:

- `audiences` (List of String)
- `claim_to_headers` (Block List) (see [below for nested schema](#nestedblock--node--http_filters--filter--jwt_authn--providers--claim_to_headers))
- `from_headers` (Block List) (see [below for nested schema](#nestedblock--node--http_filters--filter--jwt_authn--providers--from_headers))
- `remote_jwks` (Block List, Max: 1) (see [below for nested schema](#nestedblock--node--http_filters--filter--jwt_authn--providers--remote_jwks))

<a id="nestedblock--node--http_filters--filter--jwt_authn--providers--claim_to_headers"></a>
### Nested Schema for `node.http_filters.filter.jwt_authn.providers.claim_to_headers`

Optional:

- `claim_name` (String)
- `header_name` (String)


<a id="nestedblock--node--http_filters--filter--jwt_authn--providers--from_headers"></a>
### Nested Schema for `node.http_filters.filter.jwt_authn.providers.from_headers`

Required:

- `header_name` (String)
- `value_prefix` (String)


<a id="nestedblock--node--http_filters--filter--jwt_authn--providers--remote_jwks"></a>
### Nested Schema for `node.http_filters.filter.jwt_authn.providers.remote_jwks`

Required:

- `cluster_name` (String)
- `http_uri` (String)

Optional:

- `timeout` (Number)



<a id="nestedblock--node--http_filters--filter--jwt_authn--rules"></a>
### Nested Schema for `node.http_filters.filter.jwt_authn.rules`

Optional:

- `match_prefix` (String)
- `required_providers_names` (List of String)

