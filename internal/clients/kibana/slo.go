package kibana

import (
	"github.com/elastic/terraform-provider-elasticstack/generated/slo"
	"github.com/elastic/terraform-provider-elasticstack/internal/models"
)

func sloResponseToModel(spaceID string, res *slo.SloResponse) *models.Slo {
	if res == nil {
		return nil
	}

	if res.TimeWindow.TimeWindowCalendarAligned != nil {
		w := res.TimeWindow.TimeWindowCalendarAligned
		window := models.TimeWindow{}
		window.Duration = unwrapOptionalField(*w.Duration).(string)
		window.IsRolling = bool(unwrapOptionalField(w.IsRolling))

	} else if res.TimeWindow.TimeWindowRolling != nil {
		window := res.TimeWindow.TimeWindowRolling
	}

	return &models.Slo{
		ID:              *res.Id,
		SpaceID:         spaceID,
		Name:            *res.Name,
		Description:     *res.Description,
		BudgetingMethod: string(*res.BudgetingMethod),
		Indicator: models.Indicator{
			Params: models.Params{
				Index:           string(unwrapOptionalField(indicator.Params.Index)),
				Service:         string(unwrapOptionalField(indicator.Params.Service)),
				Environment:     string(unwrapOptionalField(indicator.Params.Environment)),
				TransactionType: string(unwrapOptionalField(indicator.Params.TransactionType)),
				TransactionName: string(unwrapOptionalField(indicator.Params.TransactionName)),
				//does GoodStatusCodes need to be casted?
				GoodStatusCodes: unwrapOptionalField(indicator.Params.GoodStatusCodes),
				Filter:          string(unwrapOptionalField(indicator.Params.Filter)),
				Good:            string(unwrapOptionalField(indicator.Params.Good)),
				Total:           string(unwrapOptionalField(indicator.Params.Total)),
				TimestampField:  string(unwrapOptionalField(indicator.Params.TimestampField)),
			},
			Type: string(indicator.Type),
		},
		TimeWindow: models.TimeWindow{
			Duration:   string(unwrapOptionalField(window.Duration)),
			IsRolling:  bool(unwrapOptionalField(window.IsRolling)),
			IsCalendar: bool(unwrapOptionalField(window.IsCalendar)),
		},
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
