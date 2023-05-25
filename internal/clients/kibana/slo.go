package kibana

import (
	"context"
	"net/http"

	"github.com/elastic/terraform-provider-elasticstack/generated/slo"
	"github.com/elastic/terraform-provider-elasticstack/internal/clients"
	"github.com/elastic/terraform-provider-elasticstack/internal/models"
	"github.com/elastic/terraform-provider-elasticstack/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func GetSlo(ctx context.Context, apiClient *clients.ApiClient, id, spaceID string) (*models.Slo, diag.Diagnostics) {
	client, err := apiClient.GetSloClient()
	if err != nil {
		return nil, diag.FromErr(err)
	}

	ctxWithAuth := apiClient.SetGeneratedClientAuthContext(ctx)
	req := client.GetSlo(ctxWithAuth, id, spaceID)
	sloRes, res, err := req.Execute()
	if err != nil && res == nil {
		return nil, diag.FromErr(err)
	}

	defer res.Body.Close()
	if res.StatusCode == http.StatusNotFound {
		return nil, nil
	}
	return sloResponseToModel(spaceID, sloRes), utils.CheckHttpError(res, "Unable to get slo with ID "+string(id))
}

func DeleteSlo(ctx context.Context, apiClient *clients.ApiClient, sloId string, spaceId string) diag.Diagnostics {
	client, err := apiClient.GetSloClient()
	if err != nil {
		return diag.FromErr(err)
	}

	ctxWithAuth := apiClient.SetGeneratedClientAuthContext(ctx)
	req := client.DeleteSlo(ctxWithAuth, sloId, spaceId).KbnXsrf("true")
	res, err := req.Execute()
	if err != nil && res == nil {
		return diag.FromErr(err)
	}

	defer res.Body.Close()
	return utils.CheckHttpError(res, "Unabled to delete slo")
}

func UpdateSlo(ctx context.Context, apiClient *clients.ApiClient, s models.Slo) (*models.Slo, diag.Diagnostics) {
	client, err := apiClient.GetSloClient()
	if err != nil {
		return nil, diag.FromErr(err)
	}

	ctxWithAuth := apiClient.SetGeneratedClientAuthContext(ctx)

	//fark this is so verbose
	indicator := slo.CreateSloRequestIndicator{}
	switch s.Indicator.Type {
	case "sli.kql.custom":
		params := slo.IndicatorPropertiesCustomKqlParams{
			Index:          s.Indicator.Params.Index,
			Filter:         &s.Indicator.Params.Filter,
			Good:           &s.Indicator.Params.Good,
			Total:          &s.Indicator.Params.Total,
			TimestampField: s.Indicator.Params.TimestampField,
		}
		i := slo.IndicatorPropertiesCustomKql{
			Params: params,
			Type:   "sli.kql.custom",
		}
		indicator.IndicatorPropertiesCustomKql = &i

	case "sli.apm.transactionErrorRate":
		params := slo.IndicatorPropertiesApmAvailabilityParams{
			Service:         s.Indicator.Params.Service,
			Environment:     s.Indicator.Params.Environment,
			TransactionType: s.Indicator.Params.TransactionType,
			TransactionName: s.Indicator.Params.TransactionName,
			Filter:          &s.Indicator.Params.Filter,
			Index:           s.Indicator.Params.Index,
		}
		i := slo.IndicatorPropertiesApmAvailability{
			Params: params,
			Type:   "sli.apm.transactionErrorRate",
		}
		indicator.IndicatorPropertiesApmAvailability = &i

	case "sli.apm.transactionDuration":
		params := slo.IndicatorPropertiesApmLatencyParams{
			Service:         s.Indicator.Params.Service,
			Environment:     s.Indicator.Params.Environment,
			TransactionType: s.Indicator.Params.TransactionType,
			TransactionName: s.Indicator.Params.TransactionName,
			Filter:          &s.Indicator.Params.Filter,
			Index:           s.Indicator.Params.Index,
			Threshold:       s.Indicator.Params.Threshold,
		}
		i := slo.IndicatorPropertiesApmLatency{
			Params: params,
			Type:   "sli.apm.transactionDuration",
		}
		indicator.IndicatorPropertiesApmLatency = &i
	}

	window := slo.SloResponseTimeWindow{}
	switch s.TimeWindow.IsCalendar {
	case true:
		w := slo.TimeWindowCalendarAligned{
			Duration:   s.TimeWindow.Duration,
			IsCalendar: true,
		}
		window.TimeWindowCalendarAligned = &w
	case false:
		w := slo.TimeWindowRolling{
			Duration:  s.TimeWindow.Duration,
			IsRolling: true,
		}
		window.TimeWindowRolling = &w
	}

	reqModel := slo.UpdateSloRequest{
		Name:        &s.Name,
		Description: &s.Description,
		Indicator:   &indicator,
		TimeWindow:  &window,
	}
	req := client.UpdateRule(ctxWithAuth, rule.ID, rule.SpaceID).KbnXsrf("true").UpdateRuleRequest(reqModel)
	ruleRes, res, err := req.Execute()
	if err != nil && res == nil {
		return nil, diag.FromErr(err)
	}

	defer res.Body.Close()
	if diags := utils.CheckHttpError(res, "Unable to update alerting rule"); diags.HasError() {
		return nil, diags
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

func sloResponseToModel(spaceID string, res *slo.SloResponse) *models.Slo {
	if res == nil {
		return nil
	}

	window := windowFromSloResponse(res)
	indicator := indicatorFromSloResponse(res)

	return &models.Slo{
		ID:              *res.Id,
		SpaceID:         spaceID,
		Name:            *res.Name,
		Description:     *res.Description,
		BudgetingMethod: string(*res.BudgetingMethod),
		Indicator:       indicator,
		TimeWindow:      window,
		Objective: models.Objective{
			Target:           float64(res.Objective.Target),
			TimeslicesTarget: float64(unwrapOptionalField(res.Objective.TimeslicesTarget)),
			TimeslicesWindow: string(unwrapOptionalField(res.Objective.TimeslicesWindow)),
		},
		Settings: models.Settings{
			SyncDelay: string(unwrapOptionalField(res.Settings.SyncDelay)),
			Frequency: string(unwrapOptionalField(res.Settings.Frequency)),
		},
	}
}

func indicatorFromSloResponse(res *slo.SloResponse) models.Indicator {
	indicator := models.Indicator{}

	if res.Indicator.IndicatorPropertiesCustomKql != nil {
		i := res.Indicator.IndicatorPropertiesCustomKql
		indicator = models.Indicator{
			Type: "sli.kql.custom",
			Params: models.Params{
				Index:          string(i.Params.Index),
				Filter:         string(unwrapOptionalField(i.Params.Filter)),
				Good:           string(unwrapOptionalField(i.Params.Good)),
				Total:          string(unwrapOptionalField(i.Params.Total)),
				TimestampField: string(unwrapOptionalField(&i.Params.TimestampField)),
			},
		}

	} else if res.Indicator.IndicatorPropertiesApmLatency != nil {
		i := res.Indicator.IndicatorPropertiesApmLatency
		indicator = models.Indicator{
			Type: "sli.apm.transactionDuration",
			Params: models.Params{
				Index:           string(i.Params.Index),
				Service:         string(unwrapOptionalField(&i.Params.Service)),
				Environment:     string(unwrapOptionalField(&i.Params.Environment)),
				TransactionType: string(unwrapOptionalField(&i.Params.TransactionType)),
				TransactionName: string(unwrapOptionalField(&i.Params.TransactionName)),
				Filter:          string(unwrapOptionalField(i.Params.Filter)),
				Threshold:       i.Params.Threshold, //I can't unwrap this. Will this cause issues?
			},
		}

	} else if res.Indicator.IndicatorPropertiesApmAvailability != nil {
		i := res.Indicator.IndicatorPropertiesApmAvailability
		indicator = models.Indicator{
			Type: "sli.apm.transactionErrorRate",
			Params: models.Params{
				Index:           string(i.Params.Index),
				Service:         string(unwrapOptionalField(&i.Params.Service)),
				Environment:     string(unwrapOptionalField(&i.Params.Environment)),
				TransactionType: string(unwrapOptionalField(&i.Params.TransactionType)),
				TransactionName: string(unwrapOptionalField(&i.Params.TransactionName)),
				Filter:          string(unwrapOptionalField(i.Params.Filter)),
			},
		}
	}

	return indicator

}

func windowFromSloResponse(res *slo.SloResponse) models.TimeWindow {
	window := models.TimeWindow{}
	if res.TimeWindow.TimeWindowCalendarAligned != nil {
		timeWindowCalendarAligned := res.TimeWindow.TimeWindowCalendarAligned
		window = models.TimeWindow{
			Duration:   string(timeWindowCalendarAligned.Duration),
			IsRolling:  false,
			IsCalendar: true,
		}
	} else if res.TimeWindow.TimeWindowRolling != nil {
		timeWindowRolling := res.TimeWindow.TimeWindowRolling
		window = models.TimeWindow{
			Duration:   string(timeWindowRolling.Duration),
			IsRolling:  true,
			IsCalendar: false,
		}
	}
	return window
}
