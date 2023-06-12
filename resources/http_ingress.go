package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	typesv1 "github.com/protomesh/terraform-provider-protomesh/proto/api/types/v1"
)

func DataSourceHttpIngress() *schema.Resource {
	return &schema.Resource{
		ReadContext: getHttpIngress,
	}
}

func ResourceHttpIngress() *schema.Resource {
	return &schema.Resource{
		Description:   "Expose a AWS Lambda through a gRPC interface",
		CreateContext: createHttpIngress,
		ReadContext:   readHttpIngress,
		UpdateContext: updateHttpIngress,
		DeleteContext: deleteHttpIngress,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: makeResourceSchema(typesv1.NewHttpIngressSchema()),
	}
}

func createHttpIngress(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {

	nodeData, diagErr := nodeDataFromResourceData(rd)
	if diagErr != nil {
		return diagErr
	}

	protoJsonMap, err := typesv1.UnmarshalHttpIngress(nodeData)
	if err != nil {
		return diag.FromErr(err)
	}

	node := &typesv1.HttpIngress{}

	if err := typesv1.UnmarshalHttpIngressProto(protoJsonMap, node); err != nil {
		return diag.FromErr(err)
	}

	netNode := &typesv1.NetworkingNode{
		NetworkingNode: &typesv1.NetworkingNode_HttpIngress{
			HttpIngress: node,
		},
	}

	res, diagErr := resourceFromResourceData(rd, netNode)
	if diagErr != nil {
		return diagErr
	}

	return putResource(ctx, rd, i, res)

}

func updateHttpIngress(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {

	return createHttpIngress(ctx, rd, i)

}

func deleteHttpIngress(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {

	return dropResource(ctx, rd, i)

}

func readHttpIngress(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {

	res, diagErr := getResource(ctx, rd, i)
	if res == nil || diagErr != nil {
		return diagErr
	}

	netNode := &typesv1.NetworkingNode{}

	if err := res.Spec.UnmarshalTo(netNode); err != nil {
		return diag.FromErr(err)
	}

	if nodeSpec, ok := netNode.NetworkingNode.(*typesv1.NetworkingNode_HttpIngress); ok {

		node, err := typesv1.MarshalHttpIngressProto(nodeSpec.HttpIngress)
		if err != nil {
			return diag.FromErr(err)
		}
		// panic(fmt.Sprintf("%+v", node))

		rd.Set("node", []interface{}{node})

	}

	return nil

}

func getHttpIngress(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {

	rd.SetId(rd.Get("resource_id").(string))

	return readHttpIngress(ctx, rd, i)

}