package kibana

import (
	"context"
	"net/http"

	"github.com/elastic/terraform-provider-elasticstack/generated/alerting"
	"github.com/elastic/terraform-provider-elasticstack/generated/slo"
	"github.com/elastic/terraform-provider-elasticstack/internal/clients"
	"github.com/elastic/terraform-provider-elasticstack/internal/models"
	"github.com/elastic/terraform-provider-elasticstack/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func sloResponseToModel(spaceID string, res *slo.SloResponse) *models.Slo {
	if res == nil {
		return nil
	}

	indicator := res.Indicator.GetActualInstance()
	window := res.TimeWindow.GetActualInstance()

	return &models.Slo{
		ID:              *res.Id,
		SpaceID:         spaceID,
		Name:            *res.Name,
		Description:     *res.Description,
		BudgetingMethod: string(*res.BudgetingMethod),
		Indicator: models.Indicator{
			Params: models.Params{
				Index:          string(unwrapOptionalField(indicator.Params.Index)),
				Service: string(unwrapOptionalField(indicator.Params.Service)),
				Environment: string(unwrapOptionalField(indicator.Params.Environment)),
				TransactionType: string(unwrapOptionalField(indicator.Params.TransactionType)),
				TransactionName: string(unwrapOptionalField(indicator.Params.TransactionName)),
				//does GoodStatusCodes need to be casted?
				GoodStatusCodes: unwrapOptionalField(indicator.Params.GoodStatusCodes), 
				Filter: string(unwrapOptionalField(indicator.Params.Filter)),
				Good: string(unwrapOptionalField(indicator.Params.Good)),
				Total: string(unwrapOptionalField(indicator.Params.Total)),
				TimestampField: string(unwrapOptionalField(indicator.Params.TimestampField)),
			},
			Type: string(indicator.Type),
		},
		TimeWindow: models.TimeWindow{
			Duration:  string(window.Duration),
			IsRolling: bool(window.IsRolling),
		},
		Objective: models.Objective{
			Target:           float64(*res.Objective.Target),
			TimeslicesTarget: float64(unwrapOptionalField(*res.Objective.TimeslicesTarget)),
			TimeslicesWindow: string(unwrapOptionalField(*res.Objective.TimeslicesWindow)),
		},
	},
}

// Maps the rule actions to the struct required by the request model (ActionsInner)
func ruleActionsToActionsInner(ruleActions []models.AlertingRuleAction) []alerting.ActionsInner {
	actions := []alerting.ActionsInner{}
	for index := range ruleActions {
		action := ruleActions[index]
		actions = append(actions, alerting.ActionsInner{
			Group:  &action.Group,
			Id:     &action.ID,
			Params: action.Params,
		})
	}
	return actions
}

func CreateAlertingRule(ctx context.Context, apiClient *clients.ApiClient, rule models.AlertingRule) (*models.AlertingRule, diag.Diagnostics) {
	client, err := apiClient.GetAlertingClient()
	if err != nil {
		return nil, diag.FromErr(err)
	}

	ctxWithAuth := apiClient.SetGeneratedClientAuthContext(ctx)

	reqModel := alerting.CreateRuleRequest{
		Consumer:   rule.Consumer,
		Actions:    ruleActionsToActionsInner(rule.Actions),
		Enabled:    rule.Enabled,
		Name:       rule.Name,
		NotifyWhen: (*alerting.NotifyWhen)(&rule.NotifyWhen),
		Params:     rule.Params,
		RuleTypeId: rule.RuleTypeID,
		Schedule: alerting.Schedule{
			Interval: &rule.Schedule.Interval,
		},
		Tags:     rule.Tags,
		Throttle: *alerting.NewNullableString(rule.Throttle),
	}
	req := client.CreateRule(ctxWithAuth, rule.SpaceID, "").KbnXsrf("true").CreateRuleRequest(reqModel)
	ruleRes, res, err := req.Execute()
	if err != nil && res == nil {
		return nil, diag.FromErr(err)
	}

	defer res.Body.Close()
	return ruleResponseToModel(rule.SpaceID, ruleRes), utils.CheckHttpError(res, "Unabled to create alerting rule")
}

func UpdateAlertingRule(ctx context.Context, apiClient *clients.ApiClient, rule models.AlertingRule) (*models.AlertingRule, diag.Diagnostics) {
	client, err := apiClient.GetAlertingClient()
	if err != nil {
		return nil, diag.FromErr(err)
	}

	ctxWithAuth := apiClient.SetGeneratedClientAuthContext(ctx)

	reqModel := alerting.UpdateRuleRequest{
		Actions:    ruleActionsToActionsInner((rule.Actions)),
		Name:       rule.Name,
		NotifyWhen: (*alerting.NotifyWhen)(&rule.NotifyWhen),
		Params:     rule.Params,
		Schedule: alerting.Schedule{
			Interval: &rule.Schedule.Interval,
		},
		Tags:     rule.Tags,
		Throttle: *alerting.NewNullableString(rule.Throttle),
	}
	req := client.UpdateRule(ctxWithAuth, rule.ID, rule.SpaceID).KbnXsrf("true").UpdateRuleRequest(reqModel)
	ruleRes, res, err := req.Execute()
	if err != nil && res == nil {
		return nil, diag.FromErr(err)
	}

	defer res.Body.Close()
	if diags := utils.CheckHttpError(res, "Unable to update alerting rule"); diags.HasError() {
		return nil, diag.FromErr(err)
	}

	shouldBeEnabled := rule.Enabled != nil && *rule.Enabled
	if shouldBeEnabled && !ruleRes.Enabled {
		res, err := client.EnableRule(ctxWithAuth, rule.ID, rule.SpaceID).KbnXsrf("true").Execute()
		if err != nil && res == nil {
			return nil, diag.FromErr(err)
		}

		if diags := utils.CheckHttpError(res, "Unable to enable alerting rule"); diags.HasError() {
			return nil, diag.FromErr(err)
		}
	}

	if !shouldBeEnabled && ruleRes.Enabled {
		res, err := client.DisableRule(ctxWithAuth, rule.ID, rule.SpaceID).KbnXsrf("true").Execute()
		if err != nil && res == nil {
			return nil, diag.FromErr(err)
		}

		if diags := utils.CheckHttpError(res, "Unable to disable alerting rule"); diags.HasError() {
			return nil, diag.FromErr(err)
		}
	}

	return ruleResponseToModel(rule.SpaceID, ruleRes), diag.Diagnostics{}
}

func GetAlertingRule(ctx context.Context, apiClient *clients.ApiClient, id, spaceID string) (*models.AlertingRule, diag.Diagnostics) {
	client, err := apiClient.GetAlertingClient()
	if err != nil {
		return nil, diag.FromErr(err)
	}

	ctxWithAuth := apiClient.SetGeneratedClientAuthContext(ctx)
	req := client.GetRule(ctxWithAuth, id, spaceID)
	ruleRes, res, err := req.Execute()
	if err != nil && res == nil {
		return nil, diag.FromErr(err)
	}

	defer res.Body.Close()
	if res.StatusCode == http.StatusNotFound {
		return nil, nil
	}
	return ruleResponseToModel(spaceID, ruleRes), utils.CheckHttpError(res, "Unabled to get alerting rule")
}

func DeleteAlertingRule(ctx context.Context, apiClient *clients.ApiClient, ruleId string, spaceId string) diag.Diagnostics {
	client, err := apiClient.GetAlertingClient()
	if err != nil {
		return diag.FromErr(err)
	}

	ctxWithAuth := apiClient.SetGeneratedClientAuthContext(ctx)
	req := client.DeleteRule(ctxWithAuth, ruleId, spaceId).KbnXsrf("true")
	res, err := req.Execute()
	if err != nil && res == nil {
		return diag.FromErr(err)
	}

	defer res.Body.Close()
	return utils.CheckHttpError(res, "Unabled to delete alerting rule")
}
