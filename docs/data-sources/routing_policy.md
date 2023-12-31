---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "protomesh_routing_policy Data Source - protomesh"
subcategory: ""
description: |-
  
---

# protomesh_routing_policy (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `namespace` (String) Namespace in the resource store

### Optional

- `resource_id` (String) A valid UUID for the resource

### Read-Only

- `id` (String) The ID of this resource.
- `name` (String) Name of the resource (display name)
- `node` (List of Object) Node specification. (see [below for nested schema](#nestedatt--node))
- `version_index` (Number) Resource version index (unix timestamp in seconds)
- `version_sha256_sum` (String) Current version sha256 sum

<a id="nestedatt--node"></a>
### Nested Schema for `node`

Read-Only:

- `cors` (List of Object) (see [below for nested schema](#nestedobjatt--node--cors))
- `domain` (String)
- `ingress_name` (String)
- `routes` (List of Object) (see [below for nested schema](#nestedobjatt--node--routes))
- `xds_cluster_name` (String)

<a id="nestedobjatt--node--cors"></a>
### Nested Schema for `node.cors`

Read-Only:

- `allow_headers` (List of String)
- `allow_methods` (List of String)
- `allow_origin_string_match_prefix` (List of String)
- `expose_headers` (List of String)
- `max_age` (String)


<a id="nestedobjatt--node--routes"></a>
### Nested Schema for `node.routes`

Read-Only:

- `match_prefix` (String)
- `prefix_rewrite` (String)
- `target_service` (String)
- `timeout` (String)


