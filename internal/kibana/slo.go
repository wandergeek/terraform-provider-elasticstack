package kibana

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/elastic/terraform-provider-elasticstack/internal/clients"
	"github.com/elastic/terraform-provider-elasticstack/internal/clients/kibana"
	"github.com/elastic/terraform-provider-elasticstack/internal/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceSlo() *schema.Resource {
	sloSchema := map[string]*schema.Schema{
		"id": {
			Description: "An ID (8 and 36 characters). If omitted, a UUIDv1 will be generated server-side.",
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
		},
		"name": {
			Description: "The name of the rule. While this name does not have to be unique, a distinctive name can help you identify an slo .",
			Type:        schema.TypeString,
			Required:    true,
		},
		"description": {
			Description: "A description for the SLO.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"indicator": &schema.Schema{
			Type:     schema.TypeMap,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"type": &schema.Schema{
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.StringInSlice([]string{"sli.kql.custom", "sli.apm.transactionErrorRate", "sli.apm.transactionDuration"}, false),
					},
					"params": &schema.Schema{
						Type:     schema.TypeMap,
						Required: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"index": {
									Type:     schema.TypeString,
									Required: true,
								},
								"filter": {
									Type:     schema.TypeString,
									Required: false,
								},
								"good": {
									Type:     schema.TypeString,
									Required: false,
								},
								"total": {
									Type:     schema.TypeString,
									Required: false,
								},
								"timestampField": {
									Type:     schema.TypeString,
									Required: false,
								},
							},
						},
					},
				},
			},
		},
		"time_window": {
			Description: "Currently support calendar aligned and rolling time windows. Any duration greater than 1 day can be used: days, weeks, months, quarters, years. Rolling time window requires a duration, e.g. 1w for one week, and isRolling: true. SLOs defined with such time window, will only consider the SLI data from the last duration period as a moving window. Calendar aligned time window requires a duration, limited to 1M for monthly or 1w for weekly, and isCalendar: true.",
		},
		"space_id": {
			Description: "An identifier for the space. If space_id is not provided, the default space is used.",
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "default",
			ForceNew:    true,
		},
	}

	return &schema.Resource{
		Description: "Creates a Kibana rule. See https://www.elastic.co/guide/en/kibana/master/create-rule-api.html",

		CreateContext: resourceSloCreate,
		UpdateContext: resourceSloUpdate,
		ReadContext:   resourceSloRead,
		DeleteContext: resourceSloDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: sloSchema,
	}
}

func getSloFromResourceData(d *schema.ResourceData) (models.AlertingRule, diag.Diagnostics) {
	var diags diag.Diagnostics
	rule := models.AlertingRule{
		SpaceID:    d.Get("space_id").(string),
		Name:       d.Get("name").(string),
		Consumer:   d.Get("consumer").(string),
		NotifyWhen: d.Get("notify_when").(string),
		RuleTypeID: d.Get("rule_type_id").(string),
		Schedule: models.AlertingRuleSchedule{
			Interval: d.Get("interval").(string),
		},
	}

	paramsStr := d.Get("params")
	params := map[string]interface{}{}
	if err := json.NewDecoder(strings.NewReader(paramsStr.(string))).Decode(&params); err != nil {
		return models.AlertingRule{}, diag.FromErr(err)
	}
	rule.Params = params

	if v, ok := d.GetOk("enabled"); ok {
		e := v.(bool)
		rule.Enabled = &e
	}

	if v, ok := d.GetOk("throttle"); ok {
		t := v.(string)
		rule.Throttle = &t
	}

	actions, diags := getActionsFromResourceData(d)
	if diags.HasError() {
		return models.AlertingRule{}, diags
	}
	rule.Actions = actions

	if tags, ok := d.GetOk("tags"); ok {
		for _, t := range tags.([]interface{}) {
			rule.Tags = append(rule.Tags, t.(string))
		}
	}

	return rule, diags
}

func getSlosFromResourceData(d *schema.ResourceData) ([]models.AlertingRuleAction, diag.Diagnostics) {
	actions := []models.AlertingRuleAction{}
	if v, ok := d.GetOk("actions"); ok {
		resourceActions := v.([]interface{})
		for _, a := range resourceActions {
			action := a.(map[string]interface{})
			paramsStr := action["params"].(string)
			var params map[string]interface{}
			err := json.Unmarshal([]byte(paramsStr), &params)
			if err != nil {
				return []models.AlertingRuleAction{}, diag.FromErr(err)
			}

			actions = append(actions, models.AlertingRuleAction{
				Group:  action["group"].(string),
				ID:     action["id"].(string),
				Params: params,
			})
		}
	}

	return actions, nil
}

func resourceSloCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := clients.NewApiClient(d, meta)
	if diags.HasError() {
		return diags
	}

	rule, diags := getAlertingRuleFromResourceData(d)
	if diags.HasError() {
		return diags
	}

	res, diags := kibana.CreateAlertingRule(ctx, client, rule)

	if diags.HasError() {
		return diags
	}

	id := &clients.CompositeId{ClusterId: rule.SpaceID, ResourceId: res.ID}
	d.SetId(id.String())

	return resourceRuleRead(ctx, d, meta)
}

func resourceSloUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := clients.NewApiClient(d, meta)
	if diags.HasError() {
		return diags
	}

	rule, diags := getAlertingRuleFromResourceData(d)
	if diags.HasError() {
		return diags
	}

	compId, diags := clients.CompositeIdFromStr(d.Id())
	if diags.HasError() {
		return diags
	}
	rule.ID = compId.ResourceId

	res, diags := kibana.UpdateAlertingRule(ctx, client, rule)

	if diags.HasError() {
		return diags
	}

	id := &clients.CompositeId{ClusterId: rule.SpaceID, ResourceId: res.ID}
	d.SetId(id.String())

	return resourceRuleRead(ctx, d, meta)
}

func resourceSloRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := clients.NewApiClient(d, meta)
	if diags.HasError() {
		return diags
	}
	compId, diags := clients.CompositeIdFromStr(d.Id())
	if diags.HasError() {
		return diags
	}
	id := compId.ResourceId
	spaceId := compId.ClusterId

	rule, diags := kibana.GetAlertingRule(ctx, client, id, spaceId)
	if rule == nil && diags == nil {
		d.SetId("")
		return diags
	}
	if diags.HasError() {
		return diags
	}

	// set the fields
	if err := d.Set("rule_id", rule.ID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("space_id", rule.SpaceID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", rule.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("consumer", rule.Consumer); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("notify_when", rule.NotifyWhen); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("rule_type_id", rule.RuleTypeID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("interval", rule.Schedule.Interval); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("enabled", rule.Enabled); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("tags", rule.Tags); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("throttle", rule.Throttle); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("scheduled_task_id", rule.ScheduledTaskID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("last_execution_status", rule.ExecutionStatus.Status); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("last_execution_date", rule.ExecutionStatus.LastExecutionDate.Format("2006-01-02 15:04:05.999 -0700 MST")); err != nil {
		return diag.FromErr(err)
	}

	params, err := json.Marshal(rule.Params)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("params", string(params)); err != nil {
		return diag.FromErr(err)
	}

	actions := []interface{}{}
	for _, action := range rule.Actions {
		params, err := json.Marshal(action.Params)
		if err != nil {
			return diag.FromErr(err)
		}
		actions = append(actions, map[string]interface{}{
			"group":  action.Group,
			"id":     action.ID,
			"params": string(params),
		})
	}
	if err := d.Set("actions", actions); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceSloDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := clients.NewApiClient(d, meta)
	if diags.HasError() {
		return diags
	}
	compId, diags := clients.CompositeIdFromStr(d.Id())
	if diags.HasError() {
		return diags
	}

	spaceId := d.Get("space_id").(string)

	if diags = kibana.DeleteAlertingRule(ctx, client, compId.ResourceId, spaceId); diags.HasError() {
		return diags
	}

	d.SetId("")
	return diags
}
