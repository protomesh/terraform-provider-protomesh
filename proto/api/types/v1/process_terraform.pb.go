// Code generated by protoc-gen-terraform. DO NOT EDIT.

package typesv1

import (
	"time"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"encoding/json"
	protomeshpb "github.com/protomesh/protoc-gen-terraform/pkg/protobuf"
)

func NewProcessSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"process": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"trigger": {
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Description: "Trigger process.",
						Elem: &schema.Resource{
							Schema: NewTriggerSchema(),
						},
					},
				},
			},
		},
	}
}

func UnmarshalProcess(obj map[string]interface{}) (map[string]interface{}, error) {
	p := map[string]interface{}{}
	if valueProcess, okProcess := obj["process"].([]interface{}); okProcess && len(valueProcess) > 0 {
		o := valueProcess[0].(map[string]interface{})
		if oneOfVal, ok := o["trigger"]; ok {
			if valueTriggerCollection, okTrigger := oneOfVal.([]interface{}); okTrigger && len(valueTriggerCollection) > 0 {
				if valueTrigger, okTrigger := valueTriggerCollection[0].(map[string]interface{}); okTrigger {
					msg, err := UnmarshalTrigger(valueTrigger)
					if err != nil {
						return nil, err
					}
					p["trigger"] = msg
				}
			}
		}
	}
	return p, nil
}

func UnmarshalProcessProto(obj map[string]interface{}, m proto.Message) error {
	d, err := UnmarshalProcess(obj)
	if err != nil {
		return err
	}
	b, err := json.Marshal(d)
	if err != nil {
		return err
	}
	if err := protojson.Unmarshal(b, m); err != nil {
		return err
	}
	return nil
}

func MarshalProcess(obj map[string]interface{}) (map[string]interface{}, error) {
	p := map[string]interface{}{}
	p["process"] = []interface{}{}
	if _, ok := obj["trigger"]; ok {
		p["process"] = append(p["process"].([]interface{}), map[string]interface{}{})
		if m, ok := obj["trigger"].(map[string]interface{}); ok {
			d, err := MarshalTrigger(m)
			if err != nil {
				return nil, err
			}
			p["process"].([]interface{})[0].(map[string]interface{})["trigger"] = []interface{}{d}
		}
	}
	return p, nil
}

func MarshalProcessProto(m proto.Message) (map[string]interface{}, error) {
	obj := map[string]interface{}{}
	b, err := protojson.MarshalOptions{UseProtoNames: true}.Marshal(m)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &obj)
	if err != nil {
		return nil, err
	}
	return MarshalProcess(obj)
}

func NewTriggerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the workflow to trigger.",
		},
		"task_queue": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Task queue on temporal to send workflow tasks.",
		},
		"id_prefix": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Static ID prefix.",
		},
		"cron_schedule": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Optional cron schedule for workflow. If a cron schedule is specified, the  workflow will run as a cron based on the schedule. The scheduling will be  based on UTC time. Schedule for next run only happen after the current run  is completed/failed/timeout. If a RetryPolicy is also supplied, and the  workflow failed or timeout, the workflow will be retried based on the retry  policy. While the workflow is retrying, it won't schedule its next run. If  next schedule is due while workflow is running (or retrying), then it will  skip that schedule. Cron workflow will not stop until it is terminated or  canceled (by returning temporal.CanceledError). The cron spec is as  following: ┌───────────── minute (0 - 59) │ ┌───────────── hour (0 - 23) │  │ ┌───────────── day of the month (1 - 31) │ │ │ ┌───────────── month (1 -  12) │ │ │ │ ┌───────────── day of the week (0 - 6) (Sunday to Saturday) │ │  │ │ │ │ │ │ │ │  * * * * *",
		},
		"execution_timeout": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "The end to end timeout for the child workflow execution including retries  and continue as new.  Optional: defaults to unlimited.",
		},
		"run_timeout": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "The timeout for a single run of the child workflow execution. Each retry or  continue as new should obey this timeout. Use WorkflowExecutionTimeout to  specify how long the parent is willing to wait for the child completion.  Optional: defaults to WorkflowExecutionTimeout",
		},
		"task_timeout": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Maximum execution time of a single Workflow Task. In the majority of cases  there is no need to change this timeout. Note that this timeout is not  related to the overall Workflow duration in any way. It defines for how  long the Workflow can get blocked in the case of a Workflow Worker crash.  Default is 10 seconds. Maximum value allowed by the Temporal Server is 1  minute.",
		},
		"arguments": {
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Description: "A list of arguments to pass to workflow.",
			Elem: &schema.Resource{
				Schema: protomeshpb.NewValueSchema(),
			},
		},
		"retry_policy": {
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Description: "RetryPolicy specify how to retry child workflow if error happens.  Optional: default is no retry",
			Elem: &schema.Resource{
				Schema: NewTriggerRetryPolicySchema(),
			},
		},
		"id_suffix": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"exact_id_suffix": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Use this exact id for the workflow id.",
					},
					"id_suffix_builder": {
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"ID_BUILDER_ONLY_PREFIX", "ID_BUILDER_RANDOM", "ID_BUILDER_UNIQUE"}, false),
						Optional:     true,
						Description:  "ID builder to use to generate the suffix.",
					},
				},
			},
		},
		"if_running": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"if_running_action": {
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"IF_RUNNING_ABORT", "IF_RUNNING_OVERLAP"}, false),
						Optional:     true,
					},
				},
			},
		},
		"on_drop": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"on_drop_action": {
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice([]string{"ON_DROP_DO_NOTHING", "ON_DROP_CANCEL", "ON_DROP_TERMINATE"}, false),
						Optional:     true,
						Description:  "On drop action.",
					},
				},
			},
		},
	}
}

func UnmarshalTrigger(obj map[string]interface{}) (map[string]interface{}, error) {
	p := map[string]interface{}{}
	if valueName, okName := obj["name"].(string); okName {
		p["name"] = valueName
	}
	if valueTaskQueue, okTaskQueue := obj["task_queue"].(string); okTaskQueue {
		p["task_queue"] = valueTaskQueue
	}
	if valueIdPrefix, okIdPrefix := obj["id_prefix"].(string); okIdPrefix {
		p["id_prefix"] = valueIdPrefix
	}
	if valueCronSchedule, okCronSchedule := obj["cron_schedule"].(string); okCronSchedule {
		p["cron_schedule"] = valueCronSchedule
	}
	if valueArgumentsCollection, okArguments := obj["arguments"].([]interface{}); okArguments && len(valueArgumentsCollection) > 0 {
		if valueArguments, okArguments := valueArgumentsCollection[0].(map[string]interface{}); okArguments {
			msg, err := protomeshpb.UnmarshalValue(valueArguments)
			if err != nil {
				return nil, err
			}
			p["arguments"] = msg
		}
	}
	if valueRetryPolicyCollection, okRetryPolicy := obj["retry_policy"].([]interface{}); okRetryPolicy && len(valueRetryPolicyCollection) > 0 {
		if valueRetryPolicy, okRetryPolicy := valueRetryPolicyCollection[0].(map[string]interface{}); okRetryPolicy {
			msg, err := UnmarshalTriggerRetryPolicy(valueRetryPolicy)
			if err != nil {
				return nil, err
			}
			p["retry_policy"] = msg
		}
	}
	if valueIdSuffix, okIdSuffix := obj["id_suffix"].([]interface{}); okIdSuffix && len(valueIdSuffix) > 0 {
		o := valueIdSuffix[0].(map[string]interface{})
		if oneOfVal, ok := o["exact_id_suffix"]; ok {
			if valueExactIdSuffix, okExactIdSuffix := oneOfVal.(string); okExactIdSuffix {
				p["exact_id_suffix"] = valueExactIdSuffix
			}
		}
		if oneOfVal, ok := o["id_suffix_builder"]; ok {
			if valueIdSuffixBuilder, okIdSuffixBuilder := oneOfVal.(string); okIdSuffixBuilder {
				p["id_suffix_builder"] = valueIdSuffixBuilder
			}
		}
	}
	if valueIfRunning, okIfRunning := obj["if_running"].([]interface{}); okIfRunning && len(valueIfRunning) > 0 {
		o := valueIfRunning[0].(map[string]interface{})
		if oneOfVal, ok := o["if_running_action"]; ok {
			if valueIfRunningAction, okIfRunningAction := oneOfVal.(string); okIfRunningAction {
				p["if_running_action"] = valueIfRunningAction
			}
		}
	}
	if valueOnDrop, okOnDrop := obj["on_drop"].([]interface{}); okOnDrop && len(valueOnDrop) > 0 {
		o := valueOnDrop[0].(map[string]interface{})
		if oneOfVal, ok := o["on_drop_action"]; ok {
			if valueOnDropAction, okOnDropAction := oneOfVal.(string); okOnDropAction {
				p["on_drop_action"] = valueOnDropAction
			}
		}
	}
	return p, nil
}

func UnmarshalTriggerProto(obj map[string]interface{}, m proto.Message) error {
	d, err := UnmarshalTrigger(obj)
	if err != nil {
		return err
	}
	b, err := json.Marshal(d)
	if err != nil {
		return err
	}
	if err := protojson.Unmarshal(b, m); err != nil {
		return err
	}
	return nil
}

func MarshalTrigger(obj map[string]interface{}) (map[string]interface{}, error) {
	p := map[string]interface{}{}
	p["name"], _ = obj["name"].(string)
	p["task_queue"], _ = obj["task_queue"].(string)
	p["id_prefix"], _ = obj["id_prefix"].(string)
	p["cron_schedule"], _ = obj["cron_schedule"].(string)
	if m, ok := obj["arguments"].(map[string]interface{}); ok {
		d, err := protomeshpb.MarshalValue(m)
		if err != nil {
			return nil, err
		}
		p["arguments"] = []interface{}{d}
	}
	if m, ok := obj["retry_policy"].(map[string]interface{}); ok {
		d, err := MarshalTriggerRetryPolicy(m)
		if err != nil {
			return nil, err
		}
		p["retry_policy"] = []interface{}{d}
	}
	p["id_suffix"] = []interface{}{}
	if _, ok := obj["exact_id_suffix"]; ok {
		p["id_suffix"] = append(p["id_suffix"].([]interface{}), map[string]interface{}{})
		p["id_suffix"].([]interface{})[0].(map[string]interface{})["exact_id_suffix"], _ = obj["exact_id_suffix"].(string)
	}
	if _, ok := obj["id_suffix_builder"]; ok {
		p["id_suffix"] = append(p["id_suffix"].([]interface{}), map[string]interface{}{})
		p["id_suffix"].([]interface{})[0].(map[string]interface{})["id_suffix_builder"], _ = obj["id_suffix_builder"].(string)
	}
	p["if_running"] = []interface{}{}
	if _, ok := obj["if_running_action"]; ok {
		p["if_running"] = append(p["if_running"].([]interface{}), map[string]interface{}{})
		p["if_running"].([]interface{})[0].(map[string]interface{})["if_running_action"], _ = obj["if_running_action"].(string)
	}
	p["on_drop"] = []interface{}{}
	if _, ok := obj["on_drop_action"]; ok {
		p["on_drop"] = append(p["on_drop"].([]interface{}), map[string]interface{}{})
		p["on_drop"].([]interface{})[0].(map[string]interface{})["on_drop_action"], _ = obj["on_drop_action"].(string)
	}
	return p, nil
}

func MarshalTriggerProto(m proto.Message) (map[string]interface{}, error) {
	obj := map[string]interface{}{}
	b, err := protojson.MarshalOptions{UseProtoNames: true}.Marshal(m)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &obj)
	if err != nil {
		return nil, err
	}
	return MarshalTrigger(obj)
}

func NewTriggerRetryPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"initial_interval": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  time.Duration(30000000000),
		},
		"maximum_backoff": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Maximum backoff interval between retries. Exponential backoff leads to  interval increase. This value is the cap of the interval. Default is 100x  of initial interval.",
		},
		"maximum_attempts": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Maximum number of attempts. When exceeded the retries stop even if not  expired yet. If not set or set to 0, it means unlimited, and rely on  activity ScheduleToCloseTimeout to stop.",
		},
		"non_retryable_errors": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Non-Retriable errors. This is optional. Temporal server will stop retry  if error type matches this list. Note:   - cancellation is not a failure, so it won't be retried,   - only StartToClose or Heartbeat timeouts are retryable.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}

func UnmarshalTriggerRetryPolicy(obj map[string]interface{}) (map[string]interface{}, error) {
	p := map[string]interface{}{}
	if valueMaximumAttempts, okMaximumAttempts := obj["maximum_attempts"].(int32); okMaximumAttempts {
		p["maximum_attempts"] = valueMaximumAttempts
	}
	if valueNonRetryableErrors, okNonRetryableErrors := obj["non_retryable_errors"].([]interface{}); okNonRetryableErrors {
		list := valueNonRetryableErrors
		r := []string{}
		for _, val := range list {
			r = append(r, val.(string))
		}
		p["non_retryable_errors"] = r
	}
	return p, nil
}

func UnmarshalTriggerRetryPolicyProto(obj map[string]interface{}, m proto.Message) error {
	d, err := UnmarshalTriggerRetryPolicy(obj)
	if err != nil {
		return err
	}
	b, err := json.Marshal(d)
	if err != nil {
		return err
	}
	if err := protojson.Unmarshal(b, m); err != nil {
		return err
	}
	return nil
}

func MarshalTriggerRetryPolicy(obj map[string]interface{}) (map[string]interface{}, error) {
	p := map[string]interface{}{}
	p["maximum_attempts"], _ = obj["maximum_attempts"].(int32)
	if l, ok := obj["non_retryable_errors"].([]interface{}); ok {
		p["non_retryable_errors"] = []interface{}{}
		for _, i := range l {
			d := i.(string)
			p["non_retryable_errors"] = append(p["non_retryable_errors"].([]interface{}), d)
		}
	}
	return p, nil
}

func MarshalTriggerRetryPolicyProto(m proto.Message) (map[string]interface{}, error) {
	obj := map[string]interface{}{}
	b, err := protojson.MarshalOptions{UseProtoNames: true}.Marshal(m)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &obj)
	if err != nil {
		return nil, err
	}
	return MarshalTriggerRetryPolicy(obj)
}
