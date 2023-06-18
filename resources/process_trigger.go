package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	typesv1 "github.com/protomesh/terraform-provider-protomesh/proto/api/types/v1"
)

func DataSourceProcessTrigger() *schema.Resource {
	return &schema.Resource{
		ReadContext: getProcessTrigger,

		Schema: makeDataSourceSchema("process", typesv1.NewTriggerSchema()),
	}
}

func ResourceProcessTrigger() *schema.Resource {
	return &schema.Resource{
		Description:   "Process trigger is an instruction for Protomesh worker to keep a Temporal Workflow running.",
		CreateContext: createProcessTrigger,
		ReadContext:   readProcessTrigger,
		UpdateContext: updateProcessTrigger,
		DeleteContext: deleteProcessTrigger,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: makeResourceSchema("process", typesv1.NewTriggerSchema()),
	}
}

func createProcessTrigger(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {

	triggerData, diagErr := itemDataFromResourceData("process", rd)
	if diagErr != nil {
		return diagErr
	}

	trigger := &typesv1.Trigger{}

	if err := typesv1.UnmarshalTriggerProto(triggerData, trigger); err != nil {
		return diag.FromErr(err)
	}

	processTrigger := &typesv1.Process{
		Process: &typesv1.Process_Trigger{
			Trigger: trigger,
		},
	}

	res, diagErr := resourceFromResourceData(rd, processTrigger)
	if diagErr != nil {
		return diagErr
	}

	return putResource(ctx, rd, i, res)

}

func updateProcessTrigger(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {

	return createProcessTrigger(ctx, rd, i)

}

func deleteProcessTrigger(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {

	return dropResource(ctx, rd, i)

}

func readProcessTrigger(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {

	res, diagErr := getResource(ctx, rd, i)
	if res == nil || diagErr != nil {
		return diagErr
	}

	process := &typesv1.Process{}

	if err := res.Spec.UnmarshalTo(process); err != nil {
		return diag.FromErr(err)
	}

	if triggerSpec, ok := process.Process.(*typesv1.Process_Trigger); ok {

		trigger, err := typesv1.MarshalTriggerProto(triggerSpec.Trigger)
		if err != nil {
			return diag.FromErr(err)
		}

		rd.Set("process", []interface{}{trigger})

	}

	return nil

}

func getProcessTrigger(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {

	rd.SetId(rd.Get("resource_id").(string))

	return readProcessTrigger(ctx, rd, i)

}
