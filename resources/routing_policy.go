package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	typesv1 "github.com/protomesh/terraform-provider-protomesh/proto/api/types/v1"
)

func DataSourceRoutingPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: getRoutingPolicy,

		Schema: makeDataSourceSchema("node", typesv1.NewRoutingPolicySchema()),
	}
}

func ResourceRoutingPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "Defines how the HTTP Ingress routes requests to services.",
		CreateContext: createRoutingPolicy,
		ReadContext:   readRoutingPolicy,
		UpdateContext: updateRoutingPolicy,
		DeleteContext: deleteRoutingPolicy,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: makeResourceSchema("node", typesv1.NewRoutingPolicySchema()),
	}
}

func createRoutingPolicy(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {

	nodeData, diagErr := itemDataFromResourceData("node", rd)
	if diagErr != nil {
		return diagErr
	}

	protoJsonMap, err := typesv1.UnmarshalRoutingPolicy(nodeData)
	if err != nil {
		return diag.FromErr(err)
	}

	node := &typesv1.RoutingPolicy{}

	if err := typesv1.UnmarshalRoutingPolicyProto(protoJsonMap, node); err != nil {
		return diag.FromErr(err)
	}

	netNode := &typesv1.NetworkingNode{
		NetworkingNode: &typesv1.NetworkingNode_RoutingPolicy{
			RoutingPolicy: node,
		},
	}

	res, diagErr := resourceFromResourceData(rd, netNode)
	if diagErr != nil {
		return diagErr
	}

	return putResource(ctx, rd, i, res)

}

func updateRoutingPolicy(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {

	return createRoutingPolicy(ctx, rd, i)

}

func deleteRoutingPolicy(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {

	return dropResource(ctx, rd, i)

}

func readRoutingPolicy(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {

	res, diagErr := getResource(ctx, rd, i)
	if res == nil || diagErr != nil {
		return diagErr
	}

	netNode := &typesv1.NetworkingNode{}

	if err := res.Spec.UnmarshalTo(netNode); err != nil {
		return diag.FromErr(err)
	}

	if nodeSpec, ok := netNode.NetworkingNode.(*typesv1.NetworkingNode_RoutingPolicy); ok {

		node, err := typesv1.MarshalRoutingPolicyProto(nodeSpec.RoutingPolicy)
		if err != nil {
			return diag.FromErr(err)
		}
		// panic(fmt.Sprintf("%+v", node))

		rd.Set("node", []interface{}{node})

	}

	return nil

}

func getRoutingPolicy(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {

	rd.SetId(rd.Get("resource_id").(string))

	return readRoutingPolicy(ctx, rd, i)

}
