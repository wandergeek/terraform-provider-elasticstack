# CompositeSloResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The identifier of the composite SLO. | [optional] 
**Name** | Pointer to **string** | The name of the composite SLO. | [optional] 
**TimeWindow** | Pointer to [**TimeWindowRolling**](TimeWindowRolling.md) |  | [optional] 
**BudgetingMethod** | Pointer to [**BudgetingMethod**](BudgetingMethod.md) |  | [optional] 
**CompositeMethod** | Pointer to [**CompositeMethod**](CompositeMethod.md) |  | [optional] 
**Objective** | Pointer to [**Objective**](Objective.md) |  | [optional] 
**Sources** | Pointer to [**CompositeSloResponseSources**](CompositeSloResponseSources.md) |  | [optional] 
**Summary** | Pointer to [**Summary**](Summary.md) |  | [optional] 
**CreatedAt** | Pointer to **string** | The creation date | [optional] 
**UpdatedAt** | Pointer to **string** | The last update date | [optional] 

## Methods

### NewCompositeSloResponse

`func NewCompositeSloResponse() *CompositeSloResponse`

NewCompositeSloResponse instantiates a new CompositeSloResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompositeSloResponseWithDefaults

`func NewCompositeSloResponseWithDefaults() *CompositeSloResponse`

NewCompositeSloResponseWithDefaults instantiates a new CompositeSloResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompositeSloResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompositeSloResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompositeSloResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CompositeSloResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CompositeSloResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompositeSloResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompositeSloResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompositeSloResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTimeWindow

`func (o *CompositeSloResponse) GetTimeWindow() TimeWindowRolling`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *CompositeSloResponse) GetTimeWindowOk() (*TimeWindowRolling, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *CompositeSloResponse) SetTimeWindow(v TimeWindowRolling)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *CompositeSloResponse) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.

### GetBudgetingMethod

`func (o *CompositeSloResponse) GetBudgetingMethod() BudgetingMethod`

GetBudgetingMethod returns the BudgetingMethod field if non-nil, zero value otherwise.

### GetBudgetingMethodOk

`func (o *CompositeSloResponse) GetBudgetingMethodOk() (*BudgetingMethod, bool)`

GetBudgetingMethodOk returns a tuple with the BudgetingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetingMethod

`func (o *CompositeSloResponse) SetBudgetingMethod(v BudgetingMethod)`

SetBudgetingMethod sets BudgetingMethod field to given value.

### HasBudgetingMethod

`func (o *CompositeSloResponse) HasBudgetingMethod() bool`

HasBudgetingMethod returns a boolean if a field has been set.

### GetCompositeMethod

`func (o *CompositeSloResponse) GetCompositeMethod() CompositeMethod`

GetCompositeMethod returns the CompositeMethod field if non-nil, zero value otherwise.

### GetCompositeMethodOk

`func (o *CompositeSloResponse) GetCompositeMethodOk() (*CompositeMethod, bool)`

GetCompositeMethodOk returns a tuple with the CompositeMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeMethod

`func (o *CompositeSloResponse) SetCompositeMethod(v CompositeMethod)`

SetCompositeMethod sets CompositeMethod field to given value.

### HasCompositeMethod

`func (o *CompositeSloResponse) HasCompositeMethod() bool`

HasCompositeMethod returns a boolean if a field has been set.

### GetObjective

`func (o *CompositeSloResponse) GetObjective() Objective`

GetObjective returns the Objective field if non-nil, zero value otherwise.

### GetObjectiveOk

`func (o *CompositeSloResponse) GetObjectiveOk() (*Objective, bool)`

GetObjectiveOk returns a tuple with the Objective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjective

`func (o *CompositeSloResponse) SetObjective(v Objective)`

SetObjective sets Objective field to given value.

### HasObjective

`func (o *CompositeSloResponse) HasObjective() bool`

HasObjective returns a boolean if a field has been set.

### GetSources

`func (o *CompositeSloResponse) GetSources() CompositeSloResponseSources`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *CompositeSloResponse) GetSourcesOk() (*CompositeSloResponseSources, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *CompositeSloResponse) SetSources(v CompositeSloResponseSources)`

SetSources sets Sources field to given value.

### HasSources

`func (o *CompositeSloResponse) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetSummary

`func (o *CompositeSloResponse) GetSummary() Summary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *CompositeSloResponse) GetSummaryOk() (*Summary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *CompositeSloResponse) SetSummary(v Summary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *CompositeSloResponse) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CompositeSloResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CompositeSloResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CompositeSloResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CompositeSloResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CompositeSloResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CompositeSloResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CompositeSloResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CompositeSloResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


