# CreateCompositeSloRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier for the composite SLO. Must be between 8 and 36 chars | [optional] 
**Name** | **string** | A name for the composite SLO. | 
**TimeWindow** | [**TimeWindowRolling**](TimeWindowRolling.md) |  | 
**BudgetingMethod** | [**BudgetingMethod**](BudgetingMethod.md) |  | 
**CompositeMethod** | [**CompositeMethod**](CompositeMethod.md) |  | 
**Objective** | [**Objective**](Objective.md) |  | 
**Sources** | [**CompositeSloResponseSources**](CompositeSloResponseSources.md) |  | 

## Methods

### NewCreateCompositeSloRequest

`func NewCreateCompositeSloRequest(name string, timeWindow TimeWindowRolling, budgetingMethod BudgetingMethod, compositeMethod CompositeMethod, objective Objective, sources CompositeSloResponseSources, ) *CreateCompositeSloRequest`

NewCreateCompositeSloRequest instantiates a new CreateCompositeSloRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCompositeSloRequestWithDefaults

`func NewCreateCompositeSloRequestWithDefaults() *CreateCompositeSloRequest`

NewCreateCompositeSloRequestWithDefaults instantiates a new CreateCompositeSloRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateCompositeSloRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateCompositeSloRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateCompositeSloRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateCompositeSloRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CreateCompositeSloRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCompositeSloRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCompositeSloRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTimeWindow

`func (o *CreateCompositeSloRequest) GetTimeWindow() TimeWindowRolling`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *CreateCompositeSloRequest) GetTimeWindowOk() (*TimeWindowRolling, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *CreateCompositeSloRequest) SetTimeWindow(v TimeWindowRolling)`

SetTimeWindow sets TimeWindow field to given value.


### GetBudgetingMethod

`func (o *CreateCompositeSloRequest) GetBudgetingMethod() BudgetingMethod`

GetBudgetingMethod returns the BudgetingMethod field if non-nil, zero value otherwise.

### GetBudgetingMethodOk

`func (o *CreateCompositeSloRequest) GetBudgetingMethodOk() (*BudgetingMethod, bool)`

GetBudgetingMethodOk returns a tuple with the BudgetingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetingMethod

`func (o *CreateCompositeSloRequest) SetBudgetingMethod(v BudgetingMethod)`

SetBudgetingMethod sets BudgetingMethod field to given value.


### GetCompositeMethod

`func (o *CreateCompositeSloRequest) GetCompositeMethod() CompositeMethod`

GetCompositeMethod returns the CompositeMethod field if non-nil, zero value otherwise.

### GetCompositeMethodOk

`func (o *CreateCompositeSloRequest) GetCompositeMethodOk() (*CompositeMethod, bool)`

GetCompositeMethodOk returns a tuple with the CompositeMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeMethod

`func (o *CreateCompositeSloRequest) SetCompositeMethod(v CompositeMethod)`

SetCompositeMethod sets CompositeMethod field to given value.


### GetObjective

`func (o *CreateCompositeSloRequest) GetObjective() Objective`

GetObjective returns the Objective field if non-nil, zero value otherwise.

### GetObjectiveOk

`func (o *CreateCompositeSloRequest) GetObjectiveOk() (*Objective, bool)`

GetObjectiveOk returns a tuple with the Objective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjective

`func (o *CreateCompositeSloRequest) SetObjective(v Objective)`

SetObjective sets Objective field to given value.


### GetSources

`func (o *CreateCompositeSloRequest) GetSources() CompositeSloResponseSources`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *CreateCompositeSloRequest) GetSourcesOk() (*CompositeSloResponseSources, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *CreateCompositeSloRequest) SetSources(v CompositeSloResponseSources)`

SetSources sets Sources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


