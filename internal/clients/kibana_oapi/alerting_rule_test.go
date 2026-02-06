package kibana_oapi_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/elastic/terraform-provider-elasticstack/internal/clients/kibana_oapi"
	"github.com/elastic/terraform-provider-elasticstack/internal/models"
	"github.com/elastic/terraform-provider-elasticstack/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/stretchr/testify/require"
)

func Test_ruleResponseToModel(t *testing.T) {
	// This test verifies the conversion logic from kbapi responses to models.AlertingRule
	now := time.Now()
	nowFormatted := now.Format(time.RFC3339)

	tests := []struct {
		name          string
		spaceId       string
		response      interface{}
		expectedModel *models.AlertingRule
	}{
		{
			name:          "nil response should return a nil model",
			spaceId:       "space-id",
			response:      nil,
			expectedModel: nil,
		},
		{
			name:    "minimal response with only required fields",
			spaceId: "space-id",
			response: map[string]interface{}{
				"id":           "id",
				"name":         "name",
				"consumer":     "consumer",
				"params":       map[string]interface{}{},
				"rule_type_id": "rule-type-id",
				"enabled":      true,
				"tags":         []string{},
				"schedule": map[string]interface{}{
					"interval": "1m",
				},
				"execution_status": map[string]interface{}{
					"status":              "ok",
					"last_execution_date": "",
				},
				"actions": []interface{}{},
			},
			expectedModel: &models.AlertingRule{
				RuleID:     "id",
				SpaceID:    "space-id",
				Name:       "name",
				Consumer:   "consumer",
				Params:     map[string]interface{}{},
				RuleTypeID: "rule-type-id",
				Enabled:    utils.Pointer(true),
				Tags:       []string{},
				Schedule:   models.AlertingRuleSchedule{Interval: "1m"},
				ExecutionStatus: models.AlertingRuleExecutionStatus{
					LastExecutionDate: nil,
					Status:            utils.Pointer("ok"),
				},
				Actions: []models.AlertingRuleAction{},
			},
		},
		{
			name:    "response with nil optional fields",
			spaceId: "space-id",
			response: map[string]interface{}{
				"id":           "id",
				"name":         "name",
				"consumer":     "consumer",
				"params":       map[string]interface{}{},
				"rule_type_id": "rule-type-id",
				"enabled":      false,
				"tags":         []string{"tag1"},
				"schedule": map[string]interface{}{
					"interval": "5m",
				},
				"execution_status": map[string]interface{}{
					"status":              "error",
					"last_execution_date": "",
				},
				"actions": []interface{}{},
			},
			expectedModel: &models.AlertingRule{
				RuleID:     "id",
				SpaceID:    "space-id",
				Name:       "name",
				Consumer:   "consumer",
				Params:     map[string]interface{}{},
				RuleTypeID: "rule-type-id",
				Enabled:    utils.Pointer(false),
				Tags:       []string{"tag1"},
				Schedule:   models.AlertingRuleSchedule{Interval: "5m"},
				ExecutionStatus: models.AlertingRuleExecutionStatus{
					LastExecutionDate: nil,
					Status:            utils.Pointer("error"),
				},
				Actions: []models.AlertingRuleAction{},
			},
		},
		{
			name:    "full response with all fields populated",
			spaceId: "space-id",
			response: map[string]interface{}{
				"id":           "id",
				"name":         "name",
				"consumer":     "consumer",
				"params":       map[string]interface{}{"key": "value"},
				"rule_type_id": "rule-type-id",
				"enabled":      true,
				"tags":         []string{"hello", "world"},
				"notify_when":  "broken",
				"throttle":     "throttle",
				"schedule": map[string]interface{}{
					"interval": "1m",
				},
				"scheduled_task_id": "scheduled-task-id",
				"execution_status": map[string]interface{}{
					"status":              "firing",
					"last_execution_date": nowFormatted,
				},
				"alert_delay": map[string]interface{}{
					"active": float64(4),
				},
				"actions": []interface{}{
					map[string]interface{}{
						"group":  "group-1",
						"id":     "action-id",
						"params": map[string]interface{}{"message": "alert"},
						"frequency": map[string]interface{}{
							"summary":     true,
							"notify_when": "onThrottleInterval",
							"throttle":    "10s",
						},
						"alerts_filter": map[string]interface{}{
							"query": map[string]interface{}{
								"kql": "foobar",
							},
							"timeframe": map[string]interface{}{
								"days": []interface{}{float64(3), float64(5), float64(7)},
								"hours": map[string]interface{}{
									"start": "00:00",
									"end":   "08:00",
								},
								"timezone": "UTC+1",
							},
						},
					},
				},
			},
			expectedModel: &models.AlertingRule{
				RuleID:          "id",
				SpaceID:         "space-id",
				Name:            "name",
				Consumer:        "consumer",
				Params:          map[string]interface{}{"key": "value"},
				RuleTypeID:      "rule-type-id",
				Enabled:         utils.Pointer(true),
				Tags:            []string{"hello", "world"},
				NotifyWhen:      utils.Pointer("broken"),
				Schedule:        models.AlertingRuleSchedule{Interval: "1m"},
				Throttle:        utils.Pointer("throttle"),
				ScheduledTaskID: utils.Pointer("scheduled-task-id"),
				ExecutionStatus: models.AlertingRuleExecutionStatus{
					LastExecutionDate: &now,
					Status:            utils.Pointer("firing"),
				},
				Actions: []models.AlertingRuleAction{
					{
						Group:  "group-1",
						ID:     "action-id",
						Params: map[string]interface{}{"message": "alert"},
						Frequency: &models.ActionFrequency{
							Summary:    true,
							NotifyWhen: "onThrottleInterval",
							Throttle:   utils.Pointer("10s"),
						},
						AlertsFilter: &models.ActionAlertsFilter{
							Kql: utils.Pointer("foobar"),
							Timeframe: &models.AlertsFilterTimeframe{
								Days:       []int32{3, 5, 7},
								Timezone:   "UTC+1",
								HoursStart: "00:00",
								HoursEnd:   "08:00",
							},
						},
					},
				},
				AlertDelay: utils.Pointer(float32(4)),
			},
		},
		{
			name:    "multiple actions with different configurations",
			spaceId: "space-id",
			response: map[string]interface{}{
				"id":           "id",
				"name":         "name",
				"consumer":     "consumer",
				"params":       map[string]interface{}{},
				"rule_type_id": "rule-type-id",
				"enabled":      true,
				"tags":         []string{},
				"schedule": map[string]interface{}{
					"interval": "1m",
				},
				"execution_status": map[string]interface{}{
					"status":              "ok",
					"last_execution_date": "",
				},
				"actions": []interface{}{
					map[string]interface{}{
						"group":  "group-1",
						"id":     "action-1",
						"params": map[string]interface{}{},
						"frequency": map[string]interface{}{
							"summary":     true,
							"notify_when": "onActiveAlert",
						},
					},
					map[string]interface{}{
						"group":  "group-2",
						"id":     "action-2",
						"params": map[string]interface{}{},
						"alerts_filter": map[string]interface{}{
							"query": map[string]interface{}{
								"kql": "test",
							},
						},
					},
					map[string]interface{}{
						"group":  "group-3",
						"id":     "action-3",
						"params": map[string]interface{}{},
					},
				},
			},
			expectedModel: &models.AlertingRule{
				RuleID:     "id",
				SpaceID:    "space-id",
				Name:       "name",
				Consumer:   "consumer",
				Params:     map[string]interface{}{},
				RuleTypeID: "rule-type-id",
				Enabled:    utils.Pointer(true),
				Tags:       []string{},
				Schedule:   models.AlertingRuleSchedule{Interval: "1m"},
				ExecutionStatus: models.AlertingRuleExecutionStatus{
					LastExecutionDate: nil,
					Status:            utils.Pointer("ok"),
				},
				Actions: []models.AlertingRuleAction{
					{
						Group:  "group-1",
						ID:     "action-1",
						Params: map[string]interface{}{},
						Frequency: &models.ActionFrequency{
							Summary:    true,
							NotifyWhen: "onActiveAlert",
						},
					},
					{
						Group:  "group-2",
						ID:     "action-2",
						Params: map[string]interface{}{},
						AlertsFilter: &models.ActionAlertsFilter{
							Kql: utils.Pointer("test"),
						},
					},
					{
						Group:  "group-3",
						ID:     "action-3",
						Params: map[string]interface{}{},
					},
				},
			},
		},
		{
			name:    "action with nil group defaults to 'default'",
			spaceId: "space-id",
			response: map[string]interface{}{
				"id":           "id",
				"name":         "name",
				"consumer":     "consumer",
				"params":       map[string]interface{}{},
				"rule_type_id": "rule-type-id",
				"enabled":      true,
				"tags":         []string{},
				"schedule": map[string]interface{}{
					"interval": "1m",
				},
				"execution_status": map[string]interface{}{
					"status":              "ok",
					"last_execution_date": "",
				},
				"actions": []interface{}{
					map[string]interface{}{
						"group":  nil,
						"id":     "action-1",
						"params": map[string]interface{}{},
					},
				},
			},
			expectedModel: &models.AlertingRule{
				RuleID:     "id",
				SpaceID:    "space-id",
				Name:       "name",
				Consumer:   "consumer",
				Params:     map[string]interface{}{},
				RuleTypeID: "rule-type-id",
				Enabled:    utils.Pointer(true),
				Tags:       []string{},
				Schedule:   models.AlertingRuleSchedule{Interval: "1m"},
				ExecutionStatus: models.AlertingRuleExecutionStatus{
					LastExecutionDate: nil,
					Status:            utils.Pointer("ok"),
				},
				Actions: []models.AlertingRuleAction{
					{
						Group:  "default",
						ID:     "action-1",
						Params: map[string]interface{}{},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, diags := kibana_oapi.ConvertResponseToModel(tt.spaceId, tt.response)
			require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)

			if tt.expectedModel == nil {
				require.Nil(t, result)
				return
			}

			require.NotNil(t, result)
			require.Equal(t, tt.expectedModel.RuleID, result.RuleID)
			require.Equal(t, tt.expectedModel.SpaceID, result.SpaceID)
			require.Equal(t, tt.expectedModel.Name, result.Name)
			require.Equal(t, tt.expectedModel.Consumer, result.Consumer)
			require.Equal(t, tt.expectedModel.RuleTypeID, result.RuleTypeID)
			require.Equal(t, tt.expectedModel.Enabled, result.Enabled)
			require.Equal(t, tt.expectedModel.Tags, result.Tags)
			require.Equal(t, tt.expectedModel.NotifyWhen, result.NotifyWhen)
			require.Equal(t, tt.expectedModel.Schedule.Interval, result.Schedule.Interval)
			require.Equal(t, tt.expectedModel.Throttle, result.Throttle)
			require.Equal(t, tt.expectedModel.ScheduledTaskID, result.ScheduledTaskID)
			require.Equal(t, tt.expectedModel.AlertDelay, result.AlertDelay)
			require.Equal(t, tt.expectedModel.Params, result.Params)

			// Check execution status
			require.Equal(t, tt.expectedModel.ExecutionStatus.Status, result.ExecutionStatus.Status)
			if tt.expectedModel.ExecutionStatus.LastExecutionDate != nil {
				require.NotNil(t, result.ExecutionStatus.LastExecutionDate)
				// Allow small time difference due to RFC3339 parsing precision
				timeDiff := result.ExecutionStatus.LastExecutionDate.Sub(*tt.expectedModel.ExecutionStatus.LastExecutionDate)
				require.Less(t, timeDiff.Abs().Seconds(), 1.0, "execution date differs by more than 1 second")
			} else {
				require.Nil(t, result.ExecutionStatus.LastExecutionDate)
			}

			// Check actions
			require.Equal(t, len(tt.expectedModel.Actions), len(result.Actions))
			for i, expectedAction := range tt.expectedModel.Actions {
				actualAction := result.Actions[i]
				require.Equal(t, expectedAction.Group, actualAction.Group)
				require.Equal(t, expectedAction.ID, actualAction.ID)
				require.Equal(t, expectedAction.Params, actualAction.Params)

				if expectedAction.Frequency != nil {
					require.NotNil(t, actualAction.Frequency)
					require.Equal(t, expectedAction.Frequency.Summary, actualAction.Frequency.Summary)
					require.Equal(t, expectedAction.Frequency.NotifyWhen, actualAction.Frequency.NotifyWhen)
					require.Equal(t, expectedAction.Frequency.Throttle, actualAction.Frequency.Throttle)
				} else {
					require.Nil(t, actualAction.Frequency)
				}

				if expectedAction.AlertsFilter != nil {
					require.NotNil(t, actualAction.AlertsFilter)
					require.Equal(t, expectedAction.AlertsFilter.Kql, actualAction.AlertsFilter.Kql)
					if expectedAction.AlertsFilter.Timeframe != nil {
						require.NotNil(t, actualAction.AlertsFilter.Timeframe)
						require.Equal(t, expectedAction.AlertsFilter.Timeframe.Days, actualAction.AlertsFilter.Timeframe.Days)
						require.Equal(t, expectedAction.AlertsFilter.Timeframe.Timezone, actualAction.AlertsFilter.Timeframe.Timezone)
						require.Equal(t, expectedAction.AlertsFilter.Timeframe.HoursStart, actualAction.AlertsFilter.Timeframe.HoursStart)
						require.Equal(t, expectedAction.AlertsFilter.Timeframe.HoursEnd, actualAction.AlertsFilter.Timeframe.HoursEnd)
					} else {
						require.Nil(t, actualAction.AlertsFilter.Timeframe)
					}
				} else {
					require.Nil(t, actualAction.AlertsFilter)
				}
			}
		})
	}
}

func Test_CreateUpdateAlertingRule_ErrorHandling(t *testing.T) {
	// Error handling tests using httptest to mock kbapi responses
	tests := []struct {
		name             string
		statusCode       int
		responseBody     string
		expectedErrorMsg string
		testCreate       bool // if true, test Create; if false, test Update
	}{
		{
			name:             "Create with 4xx returns error diagnostic",
			statusCode:       400,
			responseBody:     `{"statusCode":400,"error":"Bad Request","message":"Invalid rule configuration"}`,
			expectedErrorMsg: "Unexpected status code from server: got HTTP 400",
			testCreate:       true,
		},
		{
			name:             "Create with 409 conflict returns specific error",
			statusCode:       409,
			responseBody:     `{"statusCode":409,"error":"Conflict"}`,
			expectedErrorMsg: "Rule ID conflict",
			testCreate:       true,
		},
		{
			name:             "Update with 4xx returns error diagnostic",
			statusCode:       401,
			responseBody:     `{"statusCode":401,"error":"Unauthorized"}`,
			expectedErrorMsg: "Unexpected status code from server: got HTTP 401",
			testCreate:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test server that returns our mock response
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
				w.Write([]byte(tt.responseBody))
			}))
			defer server.Close()

			// Create a kbapi client pointing to our test server
			client, err := kibana_oapi.NewClient(kibana_oapi.Config{
				URL: server.URL,
			})
			require.NoError(t, err)

			ctx := context.Background()
			testRule := models.AlertingRule{
				RuleID:     "test-rule-id",
				SpaceID:    "default",
				Name:       "Test Rule",
				Consumer:   "alerts",
				RuleTypeID: ".index-threshold",
				Schedule: models.AlertingRuleSchedule{
					Interval: "1m",
				},
				Params:  map[string]interface{}{"test": "value"},
				Enabled: utils.Pointer(true),
				Tags:    []string{},
				Actions: []models.AlertingRuleAction{},
			}

			var result *models.AlertingRule
			var diags diag.Diagnostics

			if tt.testCreate {
				result, diags = kibana_oapi.CreateAlertingRule(ctx, client, testRule.SpaceID, testRule)
			} else {
				result, diags = kibana_oapi.UpdateAlertingRule(ctx, client, testRule.SpaceID, testRule)
			}

			require.Nil(t, result)
			require.True(t, diags.HasError())
			require.Contains(t, diags[0].Summary(), tt.expectedErrorMsg)
		})
	}
}

func Test_Delete_404_NotAnError(t *testing.T) {
	// Test that delete treats 404 as success (idempotent delete)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"statusCode":404,"error":"Not Found"}`))
	}))
	defer server.Close()

	client, err := kibana_oapi.NewClient(kibana_oapi.Config{
		URL: server.URL,
	})
	require.NoError(t, err)

	ctx := context.Background()
	diags := kibana_oapi.DeleteAlertingRule(ctx, client, "default", "test-rule-id")

	require.False(t, diags.HasError(), "Delete should treat 404 as success")
}

func Test_Delete_4xx_ReturnsError(t *testing.T) {
	// Test that delete returns error for non-404 errors
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"statusCode":401,"error":"Unauthorized"}`))
	}))
	defer server.Close()

	client, err := kibana_oapi.NewClient(kibana_oapi.Config{
		URL: server.URL,
	})
	require.NoError(t, err)

	ctx := context.Background()
	diags := kibana_oapi.DeleteAlertingRule(ctx, client, "default", "test-rule-id")

	require.True(t, diags.HasError())
	require.Contains(t, diags[0].Summary(), "Unexpected status code from server: got HTTP 401")
}

func Test_BuildCreateRequestBody(t *testing.T) {
	tests := []struct {
		name          string
		rule          models.AlertingRule
		checkFunction func(t *testing.T, body interface{})
	}{
		{
			name: "basic rule with only required fields",
			rule: models.AlertingRule{
				RuleID:     "test-rule",
				SpaceID:    "default",
				Name:       "Test Rule",
				Consumer:   "alerts",
				RuleTypeID: ".index-threshold",
				Schedule: models.AlertingRuleSchedule{
					Interval: "1m",
				},
				Params: map[string]interface{}{"threshold": 100},
			},
			checkFunction: func(t *testing.T, body interface{}) {
				data, err := json.Marshal(body)
				require.NoError(t, err)

				var parsed map[string]interface{}
				err = json.Unmarshal(data, &parsed)
				require.NoError(t, err)

				require.Equal(t, "Test Rule", parsed["name"])
				require.Equal(t, "alerts", parsed["consumer"])
				require.Equal(t, ".index-threshold", parsed["rule_type_id"])

				schedule := parsed["schedule"].(map[string]interface{})
				require.Equal(t, "1m", schedule["interval"])

				// Params should be present
				require.NotNil(t, parsed["params"])
			},
		},
		{
			name: "rule with enabled=true",
			rule: models.AlertingRule{
				Name:       "Test Rule",
				Consumer:   "alerts",
				RuleTypeID: ".index-threshold",
				Schedule:   models.AlertingRuleSchedule{Interval: "5m"},
				Params:     map[string]interface{}{},
				Enabled:    utils.Pointer(true),
			},
			checkFunction: func(t *testing.T, body interface{}) {
				data, err := json.Marshal(body)
				require.NoError(t, err)

				var parsed map[string]interface{}
				err = json.Unmarshal(data, &parsed)
				require.NoError(t, err)

				require.Equal(t, true, parsed["enabled"])
			},
		},
		{
			name: "rule with enabled=false",
			rule: models.AlertingRule{
				Name:       "Test Rule",
				Consumer:   "alerts",
				RuleTypeID: ".index-threshold",
				Schedule:   models.AlertingRuleSchedule{Interval: "5m"},
				Params:     map[string]interface{}{},
				Enabled:    utils.Pointer(false),
			},
			checkFunction: func(t *testing.T, body interface{}) {
				data, err := json.Marshal(body)
				require.NoError(t, err)

				var parsed map[string]interface{}
				err = json.Unmarshal(data, &parsed)
				require.NoError(t, err)

				require.Equal(t, false, parsed["enabled"])
			},
		},
		{
			name: "rule with tags and throttle",
			rule: models.AlertingRule{
				Name:       "Test Rule",
				Consumer:   "alerts",
				RuleTypeID: ".index-threshold",
				Schedule:   models.AlertingRuleSchedule{Interval: "5m"},
				Params:     map[string]interface{}{},
				Tags:       []string{"tag1", "tag2"},
				Throttle:   utils.Pointer("10m"),
				NotifyWhen: utils.Pointer("onActiveAlert"),
			},
			checkFunction: func(t *testing.T, body interface{}) {
				data, err := json.Marshal(body)
				require.NoError(t, err)

				var parsed map[string]interface{}
				err = json.Unmarshal(data, &parsed)
				require.NoError(t, err)

				tags := parsed["tags"].([]interface{})
				require.Len(t, tags, 2)
				require.Equal(t, "tag1", tags[0])
				require.Equal(t, "tag2", tags[1])

				require.Equal(t, "10m", parsed["throttle"])
				require.Equal(t, "onActiveAlert", parsed["notify_when"])
			},
		},
		{
			name: "rule with alert_delay",
			rule: models.AlertingRule{
				Name:       "Test Rule",
				Consumer:   "alerts",
				RuleTypeID: ".index-threshold",
				Schedule:   models.AlertingRuleSchedule{Interval: "5m"},
				Params:     map[string]interface{}{},
				AlertDelay: utils.Pointer(float32(5)),
			},
			checkFunction: func(t *testing.T, body interface{}) {
				data, err := json.Marshal(body)
				require.NoError(t, err)

				var parsed map[string]interface{}
				err = json.Unmarshal(data, &parsed)
				require.NoError(t, err)

				alertDelay := parsed["alert_delay"].(map[string]interface{})
				require.Equal(t, float64(5), alertDelay["active"])
			},
		},
		{
			name: "rule with actions including frequency",
			rule: models.AlertingRule{
				Name:       "Test Rule",
				Consumer:   "alerts",
				RuleTypeID: ".index-threshold",
				Schedule:   models.AlertingRuleSchedule{Interval: "5m"},
				Params:     map[string]interface{}{},
				Actions: []models.AlertingRuleAction{
					{
						Group:  "threshold met",
						ID:     "action-id",
						Params: map[string]interface{}{"message": "alert"},
						Frequency: &models.ActionFrequency{
							Summary:    true,
							NotifyWhen: "onActiveAlert",
							Throttle:   utils.Pointer("15m"),
						},
					},
				},
			},
			checkFunction: func(t *testing.T, body interface{}) {
				data, err := json.Marshal(body)
				require.NoError(t, err)

				var parsed map[string]interface{}
				err = json.Unmarshal(data, &parsed)
				require.NoError(t, err)

				actions := parsed["actions"].([]interface{})
				require.Len(t, actions, 1)

				action := actions[0].(map[string]interface{})
				require.Equal(t, "threshold met", action["group"])
				require.Equal(t, "action-id", action["id"])

				frequency := action["frequency"].(map[string]interface{})
				require.Equal(t, true, frequency["summary"])
				require.Equal(t, "onActiveAlert", frequency["notify_when"])
				require.Equal(t, "15m", frequency["throttle"])
			},
		},
		{
			name: "rule with actions including alerts_filter with kql",
			rule: models.AlertingRule{
				Name:       "Test Rule",
				Consumer:   "alerts",
				RuleTypeID: ".index-threshold",
				Schedule:   models.AlertingRuleSchedule{Interval: "5m"},
				Params:     map[string]interface{}{},
				Actions: []models.AlertingRuleAction{
					{
						Group:  "default",
						ID:     "action-id",
						Params: map[string]interface{}{},
						AlertsFilter: &models.ActionAlertsFilter{
							Kql: utils.Pointer("kibana.alert.status: active"),
						},
					},
				},
			},
			checkFunction: func(t *testing.T, body interface{}) {
				data, err := json.Marshal(body)
				require.NoError(t, err)

				var parsed map[string]interface{}
				err = json.Unmarshal(data, &parsed)
				require.NoError(t, err)

				actions := parsed["actions"].([]interface{})
				action := actions[0].(map[string]interface{})

				alertsFilter := action["alerts_filter"].(map[string]interface{})
				query := alertsFilter["query"].(map[string]interface{})
				require.Equal(t, "kibana.alert.status: active", query["kql"])
			},
		},
		{
			name: "rule with actions including alerts_filter with timeframe",
			rule: models.AlertingRule{
				Name:       "Test Rule",
				Consumer:   "alerts",
				RuleTypeID: ".index-threshold",
				Schedule:   models.AlertingRuleSchedule{Interval: "5m"},
				Params:     map[string]interface{}{},
				Actions: []models.AlertingRuleAction{
					{
						Group:  "default",
						ID:     "action-id",
						Params: map[string]interface{}{},
						AlertsFilter: &models.ActionAlertsFilter{
							Timeframe: &models.AlertsFilterTimeframe{
								Days:       []int32{1, 2, 3},
								Timezone:   "UTC",
								HoursStart: "09:00",
								HoursEnd:   "17:00",
							},
						},
					},
				},
			},
			checkFunction: func(t *testing.T, body interface{}) {
				data, err := json.Marshal(body)
				require.NoError(t, err)

				var parsed map[string]interface{}
				err = json.Unmarshal(data, &parsed)
				require.NoError(t, err)

				actions := parsed["actions"].([]interface{})
				action := actions[0].(map[string]interface{})

				alertsFilter := action["alerts_filter"].(map[string]interface{})
				timeframe := alertsFilter["timeframe"].(map[string]interface{})

				days := timeframe["days"].([]interface{})
				require.Len(t, days, 3)
				require.Equal(t, float64(1), days[0])

				require.Equal(t, "UTC", timeframe["timezone"])

				hours := timeframe["hours"].(map[string]interface{})
				require.Equal(t, "09:00", hours["start"])
				require.Equal(t, "17:00", hours["end"])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := kibana_oapi.BuildCreateRequestBody(tt.rule)
			tt.checkFunction(t, body)
		})
	}
}

func Test_BuildUpdateRequestBody(t *testing.T) {
	tests := []struct {
		name          string
		rule          models.AlertingRule
		checkFunction func(t *testing.T, body interface{})
	}{
		{
			name: "basic update with required fields",
			rule: models.AlertingRule{
				Name: "Updated Rule",
				Schedule: models.AlertingRuleSchedule{
					Interval: "10m",
				},
				Params: map[string]interface{}{"threshold": 200},
			},
			checkFunction: func(t *testing.T, body interface{}) {
				data, err := json.Marshal(body)
				require.NoError(t, err)

				var parsed map[string]interface{}
				err = json.Unmarshal(data, &parsed)
				require.NoError(t, err)

				require.Equal(t, "Updated Rule", parsed["name"])

				schedule := parsed["schedule"].(map[string]interface{})
				require.Equal(t, "10m", schedule["interval"])

				// Params should be present
				require.NotNil(t, parsed["params"])
			},
		},
		{
			name: "update with actions",
			rule: models.AlertingRule{
				Name:     "Updated Rule",
				Schedule: models.AlertingRuleSchedule{Interval: "10m"},
				Params:   map[string]interface{}{},
				Actions: []models.AlertingRuleAction{
					{
						Group:  "default",
						ID:     "action-1",
						Params: map[string]interface{}{"msg": "test"},
					},
				},
			},
			checkFunction: func(t *testing.T, body interface{}) {
				data, err := json.Marshal(body)
				require.NoError(t, err)

				var parsed map[string]interface{}
				err = json.Unmarshal(data, &parsed)
				require.NoError(t, err)

				actions := parsed["actions"].([]interface{})
				require.Len(t, actions, 1)

				action := actions[0].(map[string]interface{})
				require.Equal(t, "default", action["group"])
				require.Equal(t, "action-1", action["id"])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := kibana_oapi.BuildUpdateRequestBody(tt.rule)
			tt.checkFunction(t, body)
		})
	}
}
