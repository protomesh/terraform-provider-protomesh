package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	typesv1 "github.com/protomesh/terraform-provider-protomesh/proto/api/types/v1"
)

func DataSourceGatewayPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: getGatewayPolicy,

		Schema: makeDataSourceSchema("policy", typesv1.NewGatewayPolicySchema()),
	}
}

func ResourceGatewayPolicy() *schema.Resource {
	return &schema.Resource{
		Description:   "Expose a AWS Lambda through a gRPC interface. It uses the Protomesh Gateway to expose the gRPC method.",
		CreateContext: createGatewayPolicy,
		ReadContext:   readGatewayPolicy,
		UpdateContext: updateGatewayPolicy,
		DeleteContext: deleteGatewayPolicy,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: makeResourceSchema("policy", typesv1.NewGatewayPolicySchema()),
	}
}

func createGatewayPolicy(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {

	policyData, diagErr := itemDataFromResourceData("policy", rd)
	if diagErr != nil {
		return diagErr
	}

	protoJsonMap, err := typesv1.UnmarshalGatewayPolicy(policyData)
	if err != nil {
		return diag.FromErr(err)
	}

	policy := &typesv1.GatewayPolicy{}

	if err := typesv1.UnmarshalGatewayPolicyProto(protoJsonMap, policy); err != nil {
		return diag.FromErr(err)
	}

	res, diagErr := resourceFromResourceData(rd, policy)
	if diagErr != nil {
		return diagErr
	}

	return putResource(ctx, rd, i, res)

}

func updateGatewayPolicy(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {

	return createGatewayPolicy(ctx, rd, i)

}

func deleteGatewayPolicy(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {

	return dropResource(ctx, rd, i)

}

func readGatewayPolicy(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {

	res, diagErr := getResource(ctx, rd, i)
	if res == nil || diagErr != nil {
		return diagErr
	}

	policy := &typesv1.GatewayPolicy{}

	if err := res.Spec.UnmarshalTo(policy); err != nil {
		return diag.FromErr(err)
	}

	policyData, err := typesv1.MarshalGatewayPolicyProto(policy)
	if err != nil {
		return diag.FromErr(err)
	}

	rd.Set("policy", []interface{}{policyData})

	return nil

}

func getGatewayPolicy(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {

	rd.SetId(rd.Get("resource_id").(string))

	return readGatewayPolicy(ctx, rd, i)

}
