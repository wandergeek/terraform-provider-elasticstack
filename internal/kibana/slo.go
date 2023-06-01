package kibana

import (
	"context"

	"github.com/elastic/terraform-provider-elasticstack/generated/slo"
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
			Computed:    true,
		},
		"name": {
			Description: "The name of the SLO.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"description": {
			Description: "A description for the SLO.",
			Type:        schema.TypeString,
			Required:    true,
		},
		"indicator": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"type": {
						Type:         schema.TypeString,
						Required:     true,
						ValidateFunc: validation.StringInSlice([]string{"sli.kql.custom", "sli.apm.transactionErrorRate", "sli.apm.transactionDuration"}, false),
					},
					"params": {
						Type:     schema.TypeList,
						Required: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"index": {
									Type:     schema.TypeString,
									Required: true,
								},
								"filter": {
									Type:     schema.TypeString,
									Optional: true,
								},
								"good": {
									Type:     schema.TypeString,
									Optional: true,
								},
								"service": {
									Type:     schema.TypeString,
									Optional: true,
								},
								"environment": {
									Type:     schema.TypeString,
									Optional: true,
								},
								"transaction_type": {
									Type:     schema.TypeString,
									Optional: true,
								},
								"transaction_name": {
									Type:     schema.TypeString,
									Optional: true,
								},
								"total": {
									Type:     schema.TypeString,
									Optional: true,
								},
								"timestamp_field": {
									Type:     schema.TypeString,
									Optional: true,
								},
								"threshold": {
									Type:     schema.TypeInt,
									Optional: true,
								},
							},
						},
					},
				},
			},
		},
		"time_window": {
			Description: "Currently support calendar aligned and rolling time windows. Any duration greater than 1 day can be used: days, weeks, months, quarters, years. Rolling time window requires a duration, e.g. 1w for one week, and isRolling: true. SLOs defined with such time window, will only consider the SLI data from the last duration period as a moving window. Calendar aligned time window requires a duration, limited to 1M for monthly or 1w for weekly, and isCalendar: true.",
			Type:        schema.TypeList,
			Required:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"duration": {
						Type:     schema.TypeString,
						Required: true,
					},
					"is_rolling": {
						Type:     schema.TypeBool,
						Optional: true,
						Default:  false,
					},
					"is_calendar": {
						Type:     schema.TypeBool,
						Optional: true,
						Default:  false,
					},
				},
			},
		},
		"budgeting_method": {
			Description:  "An occurrences budgeting method uses the number of good and total events during the time window. A timeslices budgeting method uses the number of good slices and total slices during the time window. A slice is an arbitrary time window (smaller than the overall SLO time window) that is either considered good or bad, calculated from the timeslice threshold and the ratio of good over total events that happened during the slice window. A budgeting method is required and must be either occurrences or timeslices.",
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice([]string{"occurrences", "timeslices"}, false),
		},
		"objective": {
			Description: "The target objective is the value the SLO needs to meet during the time window. If a timeslices budgeting method is used, we also need to define the timesliceTarget which can be different than the overall SLO target.",
			Type:        schema.TypeList,
			Required:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"target": {
						Type:     schema.TypeFloat,
						Required: true,
					},
					"timeslices_target": {
						Type:     schema.TypeFloat,
						Optional: true,
					},
					"timeslices_window": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		"settings": {
			Description: "The default settings should be sufficient for most users, but if needed, these properties can be overwritten.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"sync_delay": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"frequency": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
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
		Description: "Creates an SLO.",

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

// surely this kind of thing exists in the SDK?
func getOrNilString(path string, d *schema.ResourceData) *string {
	if v, ok := d.GetOk(path); ok {
		str := v.(string)
		return &str
	}
	return nil
}

func getOrNilFloat32(path string, d *schema.ResourceData) *float32 {
	if v, ok := d.GetOk(path); ok {
		r := v.(float32)
		return &r
	}
	return nil
}

func getSloFromResourceData(d *schema.ResourceData) (models.Slo, diag.Diagnostics) {
	var diags diag.Diagnostics

	var indicator slo.SloResponseIndicator
	indicatorType, ok := d.Get("indicator.type").(string)
	if !ok {
		return models.Slo{}, diag.Errorf("expected indicator.type to be a string, but got %T", d.Get("indicator.type"))
	}
	switch d.Get("indicator.type") {
	case "sli.kql.custom":
		indicator = slo.SloResponseIndicator{
			IndicatorPropertiesCustomKql: &slo.IndicatorPropertiesCustomKql{
				Type: indicatorType,
				Params: slo.IndicatorPropertiesCustomKqlParams{
					Index:          d.Get("indicator.params.index").(string),
					Filter:         getOrNilString("indicator.params.filter", d),
					Good:           getOrNilString("indicator.params.good", d),
					Total:          getOrNilString("indicator.params.total", d),
					TimestampField: d.Get("indicator.params.timestamp_field").(string),
				},
			},
		}

	case "sli.apm.transactionErrorRate":
		indicator = slo.SloResponseIndicator{
			IndicatorPropertiesApmAvailability: &slo.IndicatorPropertiesApmAvailability{
				Type: indicatorType,
				Params: slo.IndicatorPropertiesApmAvailabilityParams{
					Service:         d.Get("indicator.params.service").(string),
					Environment:     d.Get("indicator.params.environment").(string),
					TransactionType: d.Get("indicator.params.transaction_type").(string),
					TransactionName: d.Get("indicator.params.transaction_name").(string),
					Filter:          getOrNilString("indicator.params.filter", d),
					Index:           d.Get("indicator.params.index").(string),
				},
			},
		}

	case "sli.apm.transactionDuration":
		indicator = slo.SloResponseIndicator{
			IndicatorPropertiesApmLatency: &slo.IndicatorPropertiesApmLatency{
				Type: indicatorType,
				Params: slo.IndicatorPropertiesApmLatencyParams{
					Service:         d.Get("indicator.params.service").(string),
					Environment:     d.Get("indicator.params.environment").(string),
					TransactionType: d.Get("indicator.params.transaction_type").(string),
					TransactionName: d.Get("indicator.params.transaction_name").(string),
					Filter:          getOrNilString("indicator.params.filter", d),
					Index:           d.Get("indicator.params.index").(string),
					Threshold:       d.Get("indicator.params.threshold").(float32),
				},
			},
		}
	}

	var timeWindow slo.SloResponseTimeWindow
	if d.Get("time_window.is_rolling").(bool) {
		timeWindow = slo.SloResponseTimeWindow{
			TimeWindowRolling: &slo.TimeWindowRolling{
				IsRolling: true,
				Duration:  d.Get("time_window.duration").(string),
			},
		}
	} else {
		timeWindow = slo.SloResponseTimeWindow{
			TimeWindowCalendarAligned: &slo.TimeWindowCalendarAligned{
				IsCalendar: true,
				Duration:   d.Get("time_window.duration").(string),
			},
		}
	}

	objective := slo.Objective{
		Target:           d.Get("objective.target").(float32),
		TimeslicesTarget: getOrNilFloat32("objective.timeslices_target", d),
		TimeslicesWindow: getOrNilString("objective.timeslices_window", d),
	}

	var settings slo.Settings
	if _, ok := d.GetOk("settings"); ok {
		settings = slo.Settings{
			SyncDelay: getOrNilString("settings.sync_delay", d),
			Frequency: getOrNilString("settings.frequency", d),
		}
	}

	slo := models.Slo{
		Name:            d.Get("name").(string),
		Description:     d.Get("description").(string),
		Indicator:       indicator,
		TimeWindow:      timeWindow,
		BudgetingMethod: d.Get("budgeting_method").(string),
		Objective:       objective,
		Settings:        &settings,
		SpaceID:         d.Get("space_id").(string),
	}

	return slo, diags
}

func resourceSloCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, diags := clients.NewApiClient(d, meta)
	if diags.HasError() {
		return diags
	}

	slo, diags := getSloFromResourceData(d)
	if diags.HasError() {
		return diags
	}

	res, diags := kibana.CreateSlo(ctx, client, slo)

	if diags.HasError() {
		return diags
	}

	id, diags := client.ID(ctx, res.ID)
	if diags.HasError() {
		return diags
	}

	d.SetId(id.String())

	return resourceSloRead(ctx, d, meta)
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

	return resourceSloRead(ctx, d, meta)
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

	s, diags := kibana.GetSlo(ctx, client, id, spaceId)
	if s == nil && diags == nil {
		d.SetId("")
		return diags
	}
	if diags.HasError() {
		return diags
	}

	//I hate this so much
	if s.Indicator.IndicatorPropertiesApmAvailability != nil {
		if err := d.Set("indicator.type", "sli.apm.transactionErrorRate"); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("indicator.params.service", s.Indicator.IndicatorPropertiesApmAvailability.Params.Service); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("indicator.params.environment", s.Indicator.IndicatorPropertiesApmAvailability.Params.Environment); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("indicator.params.transaction_type", s.Indicator.IndicatorPropertiesApmAvailability.Params.TransactionType); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("indicator.params.transaction_name", s.Indicator.IndicatorPropertiesApmAvailability.Params.TransactionName); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("indicator.params.filter", s.Indicator.IndicatorPropertiesApmAvailability.Params.TransactionName); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("indicator.params.index", s.Indicator.IndicatorPropertiesApmAvailability.Params.Index); err != nil {
			return diag.FromErr(err)
		}
	} else if s.Indicator.IndicatorPropertiesApmLatency != nil {
		if err := d.Set("indicator.type", "sli.apm.transactionDuration"); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("indicator.params.service", s.Indicator.IndicatorPropertiesApmLatency.Params.Service); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("indicator.params.environment", s.Indicator.IndicatorPropertiesApmLatency.Params.Environment); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("indicator.params.transaction_type", s.Indicator.IndicatorPropertiesApmLatency.Params.TransactionType); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("indicator.params.transaction_name", s.Indicator.IndicatorPropertiesApmLatency.Params.TransactionName); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("indicator.params.filter", s.Indicator.IndicatorPropertiesApmLatency.Params.Filter); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("indicator.params.index", s.Indicator.IndicatorPropertiesApmLatency.Params.Index); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("indicator.params.threshold", s.Indicator.IndicatorPropertiesApmLatency.Params.Threshold); err != nil {
			return diag.FromErr(err)
		}

	} else if s.Indicator.IndicatorPropertiesCustomKql != nil {
		if err := d.Set("indicator.type", "sli.kql.custom"); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("indicator.params.index", s.Indicator.IndicatorPropertiesCustomKql.Params.Index); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("indicator.params.filter", s.Indicator.IndicatorPropertiesCustomKql.Params.Filter); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("indicator.params.good", s.Indicator.IndicatorPropertiesCustomKql.Params.Good); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("indicator.params.total", s.Indicator.IndicatorPropertiesCustomKql.Params.Total); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("indicator.params.timestamp_field", s.Indicator.IndicatorPropertiesCustomKql.Params.TimestampField); err != nil {
			return diag.FromErr(err)
		}
	}

	if s.TimeWindow.TimeWindowCalendarAligned != nil {
		if err := d.Set("time_window.duration", s.TimeWindow.TimeWindowCalendarAligned.Duration); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("time_window.is_calendar", s.TimeWindow.TimeWindowCalendarAligned.IsCalendar); err != nil {
			return diag.FromErr(err)
		}
	} else if s.TimeWindow.TimeWindowRolling != nil {
		if err := d.Set("time_window.duration", s.TimeWindow.TimeWindowRolling.Duration); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("time_window.is_rolling", s.TimeWindow.TimeWindowRolling.IsRolling); err != nil {
			return diag.FromErr(err)
		}
	}

	if err := d.Set("objective.target", s.Objective.Target); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("objective.timeslices_target", s.Objective.TimeslicesTarget); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("objective.timeslices_window", s.Objective.TimeslicesWindow); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("settings.sync_delay", s.Settings.SyncDelay); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("settings.frequency", s.Settings.Frequency); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("id", s.ID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("space_id", s.SpaceID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", s.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", s.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("budgeting_method", s.BudgetingMethod); err != nil {
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
