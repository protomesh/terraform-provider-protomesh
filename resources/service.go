package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	typesv1 "github.com/protomesh/terraform-provider-protomesh/proto/api/types/v1"
)

func DataSourceService() *schema.Resource {
	return &schema.Resource{
		ReadContext: getService,

		Schema: makeDataSourceSchema("node", typesv1.NewServiceSchema()),
	}
}

func ResourceService() *schema.Resource {
	return &schema.Resource{
		Description:   "Service defines an upstream cluster to receive routed requests from HTTP Ingress.",
		CreateContext: createService,
		ReadContext:   readService,
		UpdateContext: updateService,
		DeleteContext: deleteService,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: makeResourceSchema("node", typesv1.NewServiceSchema()),
	}
}

func createService(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {

	nodeData, diagErr := itemDataFromResourceData("node", rd)
	if diagErr != nil {
		return diagErr
	}

	node := &typesv1.Service{}

	if err := typesv1.UnmarshalServiceProto(nodeData, node); err != nil {
		return diag.FromErr(err)
	}

	netNode := &typesv1.NetworkingNode{
		NetworkingNode: &typesv1.NetworkingNode_Service{
			Service: node,
		},
	}

	res, diagErr := resourceFromResourceData(rd, netNode)
	if diagErr != nil {
		return diagErr
	}

	return putResource(ctx, rd, i, res)

}

func updateService(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {

	return createService(ctx, rd, i)

}

func deleteService(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {

	return dropResource(ctx, rd, i)

}

func readService(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {

	res, diagErr := getResource(ctx, rd, i)
	if res == nil || diagErr != nil {
		return diagErr
	}

	netNode := &typesv1.NetworkingNode{}

	if err := res.Spec.UnmarshalTo(netNode); err != nil {
		return diag.FromErr(err)
	}

	if nodeSpec, ok := netNode.NetworkingNode.(*typesv1.NetworkingNode_Service); ok {

		node, err := typesv1.MarshalServiceProto(nodeSpec.Service)
		if err != nil {
			return diag.FromErr(err)
		}
		// panic(fmt.Sprintf("%+v", node))

		rd.Set("node", []interface{}{node})

	}

	return nil

}

func getService(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {

	rd.SetId(rd.Get("resource_id").(string))

	return readService(ctx, rd, i)

}
