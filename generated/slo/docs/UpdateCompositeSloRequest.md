# UpdateCompositeSloRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier for the composite SLO. Must be between 8 and 36 chars | [optional] 
**Name** | Pointer to **string** | A name for the composite SLO. | [optional] 
**TimeWindow** | Pointer to [**TimeWindowRolling**](TimeWindowRolling.md) |  | [optional] 
**BudgetingMethod** | Pointer to [**BudgetingMethod**](BudgetingMethod.md) |  | [optional] 
**CompositeMethod** | Pointer to [**CompositeMethod**](CompositeMethod.md) |  | [optional] 
**Objective** | Pointer to [**Objective**](Objective.md) |  | [optional] 
**Sources** | Pointer to [**CompositeSloResponseSources**](CompositeSloResponseSources.md) |  | [optional] 

## Methods

### NewUpdateCompositeSloRequest

`func NewUpdateCompositeSloRequest() *UpdateCompositeSloRequest`

NewUpdateCompositeSloRequest instantiates a new UpdateCompositeSloRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCompositeSloRequestWithDefaults

`func NewUpdateCompositeSloRequestWithDefaults() *UpdateCompositeSloRequest`

NewUpdateCompositeSloRequestWithDefaults instantiates a new UpdateCompositeSloRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateCompositeSloRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateCompositeSloRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateCompositeSloRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateCompositeSloRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateCompositeSloRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCompositeSloRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCompositeSloRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCompositeSloRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTimeWindow

`func (o *UpdateCompositeSloRequest) GetTimeWindow() TimeWindowRolling`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *UpdateCompositeSloRequest) GetTimeWindowOk() (*TimeWindowRolling, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *UpdateCompositeSloRequest) SetTimeWindow(v TimeWindowRolling)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *UpdateCompositeSloRequest) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.

### GetBudgetingMethod

`func (o *UpdateCompositeSloRequest) GetBudgetingMethod() BudgetingMethod`

GetBudgetingMethod returns the BudgetingMethod field if non-nil, zero value otherwise.

### GetBudgetingMethodOk

`func (o *UpdateCompositeSloRequest) GetBudgetingMethodOk() (*BudgetingMethod, bool)`

GetBudgetingMethodOk returns a tuple with the BudgetingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetingMethod

`func (o *UpdateCompositeSloRequest) SetBudgetingMethod(v BudgetingMethod)`

SetBudgetingMethod sets BudgetingMethod field to given value.

### HasBudgetingMethod

`func (o *UpdateCompositeSloRequest) HasBudgetingMethod() bool`

HasBudgetingMethod returns a boolean if a field has been set.

### GetCompositeMethod

`func (o *UpdateCompositeSloRequest) GetCompositeMethod() CompositeMethod`

GetCompositeMethod returns the CompositeMethod field if non-nil, zero value otherwise.

### GetCompositeMethodOk

`func (o *UpdateCompositeSloRequest) GetCompositeMethodOk() (*CompositeMethod, bool)`

GetCompositeMethodOk returns a tuple with the CompositeMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeMethod

`func (o *UpdateCompositeSloRequest) SetCompositeMethod(v CompositeMethod)`

SetCompositeMethod sets CompositeMethod field to given value.

### HasCompositeMethod

`func (o *UpdateCompositeSloRequest) HasCompositeMethod() bool`

HasCompositeMethod returns a boolean if a field has been set.

### GetObjective

`func (o *UpdateCompositeSloRequest) GetObjective() Objective`

GetObjective returns the Objective field if non-nil, zero value otherwise.

### GetObjectiveOk

`func (o *UpdateCompositeSloRequest) GetObjectiveOk() (*Objective, bool)`

GetObjectiveOk returns a tuple with the Objective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjective

`func (o *UpdateCompositeSloRequest) SetObjective(v Objective)`

SetObjective sets Objective field to given value.

### HasObjective

`func (o *UpdateCompositeSloRequest) HasObjective() bool`

HasObjective returns a boolean if a field has been set.

### GetSources

`func (o *UpdateCompositeSloRequest) GetSources() CompositeSloResponseSources`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *UpdateCompositeSloRequest) GetSourcesOk() (*CompositeSloResponseSources, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *UpdateCompositeSloRequest) SetSources(v CompositeSloResponseSources)`

SetSources sets Sources field to given value.

### HasSources

`func (o *UpdateCompositeSloRequest) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


