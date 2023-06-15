package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	typesv1 "github.com/protomesh/terraform-provider-protomesh/proto/api/types/v1"
)

func DataSourceInstanceSet() *schema.Resource {
	return &schema.Resource{
		ReadContext: getInstanceSet,

		Schema: makeDataSourceSchema("node", typesv1.NewInstanceSetSchema()),
	}
}

func ResourceInstanceSet() *schema.Resource {
	return &schema.Resource{
		Description:   "Instance set is a group of instances serving the same service through a well known port.",
		CreateContext: createInstanceSet,
		ReadContext:   readInstanceSet,
		UpdateContext: updateInstanceSet,
		DeleteContext: deleteInstanceSet,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: makeResourceSchema("node", typesv1.NewInstanceSetSchema()),
	}
}

func createInstanceSet(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {

	nodeData, diagErr := itemDataFromResourceData("node", rd)
	if diagErr != nil {
		return diagErr
	}

	protoJsonMap, err := typesv1.UnmarshalInstanceSet(nodeData)
	if err != nil {
		return diag.FromErr(err)
	}

	node := &typesv1.InstanceSet{}

	if err := typesv1.UnmarshalInstanceSetProto(protoJsonMap, node); err != nil {
		return diag.FromErr(err)
	}

	netNode := &typesv1.NetworkingNode{
		NetworkingNode: &typesv1.NetworkingNode_InstanceSet{
			InstanceSet: node,
		},
	}

	res, diagErr := resourceFromResourceData(rd, netNode)
	if diagErr != nil {
		return diagErr
	}

	return putResource(ctx, rd, i, res)

}

func updateInstanceSet(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {

	return createInstanceSet(ctx, rd, i)

}

func deleteInstanceSet(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {

	return dropResource(ctx, rd, i)

}

func readInstanceSet(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {

	res, diagErr := getResource(ctx, rd, i)
	if res == nil || diagErr != nil {
		return diagErr
	}

	netNode := &typesv1.NetworkingNode{}

	if err := res.Spec.UnmarshalTo(netNode); err != nil {
		return diag.FromErr(err)
	}

	if nodeSpec, ok := netNode.NetworkingNode.(*typesv1.NetworkingNode_InstanceSet); ok {

		node, err := typesv1.MarshalInstanceSetProto(nodeSpec.InstanceSet)
		if err != nil {
			return diag.FromErr(err)
		}
		// panic(fmt.Sprintf("%+v", node))

		rd.Set("node", []interface{}{node})

	}

	return nil

}

func getInstanceSet(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {

	rd.SetId(rd.Get("resource_id").(string))

	return readInstanceSet(ctx, rd, i)

}
