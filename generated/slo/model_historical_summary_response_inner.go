/*
SLOs

OpenAPI schema for SLOs endpoints

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slo

import (
	"encoding/json"
)

// checks if the HistoricalSummaryResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoricalSummaryResponseInner{}

// HistoricalSummaryResponseInner struct for HistoricalSummaryResponseInner
type HistoricalSummaryResponseInner struct {
	Date        *string        `json:"date,omitempty"`
	Status      *SummaryStatus `json:"status,omitempty"`
	SliValue    *float64       `json:"sliValue,omitempty"`
	ErrorBudget *ErrorBudget   `json:"errorBudget,omitempty"`
}

// NewHistoricalSummaryResponseInner instantiates a new HistoricalSummaryResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricalSummaryResponseInner() *HistoricalSummaryResponseInner {
	this := HistoricalSummaryResponseInner{}
	return &this
}

// NewHistoricalSummaryResponseInnerWithDefaults instantiates a new HistoricalSummaryResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricalSummaryResponseInnerWithDefaults() *HistoricalSummaryResponseInner {
	this := HistoricalSummaryResponseInner{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *HistoricalSummaryResponseInner) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalSummaryResponseInner) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *HistoricalSummaryResponseInner) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *HistoricalSummaryResponseInner) SetDate(v string) {
	o.Date = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HistoricalSummaryResponseInner) GetStatus() SummaryStatus {
	if o == nil || IsNil(o.Status) {
		var ret SummaryStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalSummaryResponseInner) GetStatusOk() (*SummaryStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HistoricalSummaryResponseInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given SummaryStatus and assigns it to the Status field.
func (o *HistoricalSummaryResponseInner) SetStatus(v SummaryStatus) {
	o.Status = &v
}

// GetSliValue returns the SliValue field value if set, zero value otherwise.
func (o *HistoricalSummaryResponseInner) GetSliValue() float64 {
	if o == nil || IsNil(o.SliValue) {
		var ret float64
		return ret
	}
	return *o.SliValue
}

// GetSliValueOk returns a tuple with the SliValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalSummaryResponseInner) GetSliValueOk() (*float64, bool) {
	if o == nil || IsNil(o.SliValue) {
		return nil, false
	}
	return o.SliValue, true
}

// HasSliValue returns a boolean if a field has been set.
func (o *HistoricalSummaryResponseInner) HasSliValue() bool {
	if o != nil && !IsNil(o.SliValue) {
		return true
	}

	return false
}

// SetSliValue gets a reference to the given float64 and assigns it to the SliValue field.
func (o *HistoricalSummaryResponseInner) SetSliValue(v float64) {
	o.SliValue = &v
}

// GetErrorBudget returns the ErrorBudget field value if set, zero value otherwise.
func (o *HistoricalSummaryResponseInner) GetErrorBudget() ErrorBudget {
	if o == nil || IsNil(o.ErrorBudget) {
		var ret ErrorBudget
		return ret
	}
	return *o.ErrorBudget
}

// GetErrorBudgetOk returns a tuple with the ErrorBudget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalSummaryResponseInner) GetErrorBudgetOk() (*ErrorBudget, bool) {
	if o == nil || IsNil(o.ErrorBudget) {
		return nil, false
	}
	return o.ErrorBudget, true
}

// HasErrorBudget returns a boolean if a field has been set.
func (o *HistoricalSummaryResponseInner) HasErrorBudget() bool {
	if o != nil && !IsNil(o.ErrorBudget) {
		return true
	}

	return false
}

// SetErrorBudget gets a reference to the given ErrorBudget and assigns it to the ErrorBudget field.
func (o *HistoricalSummaryResponseInner) SetErrorBudget(v ErrorBudget) {
	o.ErrorBudget = &v
}

func (o HistoricalSummaryResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoricalSummaryResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SliValue) {
		toSerialize["sliValue"] = o.SliValue
	}
	if !IsNil(o.ErrorBudget) {
		toSerialize["errorBudget"] = o.ErrorBudget
	}
	return toSerialize, nil
}

type NullableHistoricalSummaryResponseInner struct {
	value *HistoricalSummaryResponseInner
	isSet bool
}

func (v NullableHistoricalSummaryResponseInner) Get() *HistoricalSummaryResponseInner {
	return v.value
}

func (v *NullableHistoricalSummaryResponseInner) Set(val *HistoricalSummaryResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoricalSummaryResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoricalSummaryResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoricalSummaryResponseInner(val *HistoricalSummaryResponseInner) *NullableHistoricalSummaryResponseInner {
	return &NullableHistoricalSummaryResponseInner{value: val, isSet: true}
}

func (v NullableHistoricalSummaryResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoricalSummaryResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
