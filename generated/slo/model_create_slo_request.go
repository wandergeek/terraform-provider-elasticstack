/*
SLOs

OpenAPI schema for SLOs endpoints

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slo

import (
	"encoding/json"
)

// checks if the CreateSloRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSloRequest{}

// CreateSloRequest The create SLO API request body varies depending on the type of indicator, time window and budgeting method.
type CreateSloRequest struct {
	// A unique identifier for the SLO. Must be between 8 and 36 chars
	Id *string `json:"id,omitempty"`
	// A name for the SLO.
	Name string `json:"name"`
	// A description for the SLO.
	Description     string                    `json:"description"`
	Indicator       CreateSloRequestIndicator `json:"indicator"`
	TimeWindow      SloResponseTimeWindow     `json:"timeWindow"`
	BudgetingMethod BudgetingMethod           `json:"budgetingMethod"`
	Objective       Objective                 `json:"objective"`
	Settings        *Settings                 `json:"settings,omitempty"`
}

// NewCreateSloRequest instantiates a new CreateSloRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSloRequest(name string, description string, indicator CreateSloRequestIndicator, timeWindow SloResponseTimeWindow, budgetingMethod BudgetingMethod, objective Objective) *CreateSloRequest {
	this := CreateSloRequest{}
	this.Name = name
	this.Description = description
	this.Indicator = indicator
	this.TimeWindow = timeWindow
	this.BudgetingMethod = budgetingMethod
	this.Objective = objective
	return &this
}

// NewCreateSloRequestWithDefaults instantiates a new CreateSloRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSloRequestWithDefaults() *CreateSloRequest {
	this := CreateSloRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateSloRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSloRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateSloRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateSloRequest) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *CreateSloRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateSloRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateSloRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *CreateSloRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateSloRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateSloRequest) SetDescription(v string) {
	o.Description = v
}

// GetIndicator returns the Indicator field value
func (o *CreateSloRequest) GetIndicator() CreateSloRequestIndicator {
	if o == nil {
		var ret CreateSloRequestIndicator
		return ret
	}

	return o.Indicator
}

// GetIndicatorOk returns a tuple with the Indicator field value
// and a boolean to check if the value has been set.
func (o *CreateSloRequest) GetIndicatorOk() (*CreateSloRequestIndicator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Indicator, true
}

// SetIndicator sets field value
func (o *CreateSloRequest) SetIndicator(v CreateSloRequestIndicator) {
	o.Indicator = v
}

// GetTimeWindow returns the TimeWindow field value
func (o *CreateSloRequest) GetTimeWindow() SloResponseTimeWindow {
	if o == nil {
		var ret SloResponseTimeWindow
		return ret
	}

	return o.TimeWindow
}

// GetTimeWindowOk returns a tuple with the TimeWindow field value
// and a boolean to check if the value has been set.
func (o *CreateSloRequest) GetTimeWindowOk() (*SloResponseTimeWindow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeWindow, true
}

// SetTimeWindow sets field value
func (o *CreateSloRequest) SetTimeWindow(v SloResponseTimeWindow) {
	o.TimeWindow = v
}

// GetBudgetingMethod returns the BudgetingMethod field value
func (o *CreateSloRequest) GetBudgetingMethod() BudgetingMethod {
	if o == nil {
		var ret BudgetingMethod
		return ret
	}

	return o.BudgetingMethod
}

// GetBudgetingMethodOk returns a tuple with the BudgetingMethod field value
// and a boolean to check if the value has been set.
func (o *CreateSloRequest) GetBudgetingMethodOk() (*BudgetingMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BudgetingMethod, true
}

// SetBudgetingMethod sets field value
func (o *CreateSloRequest) SetBudgetingMethod(v BudgetingMethod) {
	o.BudgetingMethod = v
}

// GetObjective returns the Objective field value
func (o *CreateSloRequest) GetObjective() Objective {
	if o == nil {
		var ret Objective
		return ret
	}

	return o.Objective
}

// GetObjectiveOk returns a tuple with the Objective field value
// and a boolean to check if the value has been set.
func (o *CreateSloRequest) GetObjectiveOk() (*Objective, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Objective, true
}

// SetObjective sets field value
func (o *CreateSloRequest) SetObjective(v Objective) {
	o.Objective = v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *CreateSloRequest) GetSettings() Settings {
	if o == nil || IsNil(o.Settings) {
		var ret Settings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSloRequest) GetSettingsOk() (*Settings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *CreateSloRequest) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given Settings and assigns it to the Settings field.
func (o *CreateSloRequest) SetSettings(v Settings) {
	o.Settings = &v
}

func (o CreateSloRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSloRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["indicator"] = o.Indicator
	toSerialize["timeWindow"] = o.TimeWindow
	toSerialize["budgetingMethod"] = o.BudgetingMethod
	toSerialize["objective"] = o.Objective
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	return toSerialize, nil
}

type NullableCreateSloRequest struct {
	value *CreateSloRequest
	isSet bool
}

func (v NullableCreateSloRequest) Get() *CreateSloRequest {
	return v.value
}

func (v *NullableCreateSloRequest) Set(val *CreateSloRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSloRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSloRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSloRequest(val *CreateSloRequest) *NullableCreateSloRequest {
	return &NullableCreateSloRequest{value: val, isSet: true}
}

func (v NullableCreateSloRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSloRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
