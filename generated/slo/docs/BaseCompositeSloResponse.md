# BaseCompositeSloResponse

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
**CreatedAt** | Pointer to **string** | The creation date | [optional] 
**UpdatedAt** | Pointer to **string** | The last update date | [optional] 

## Methods

### NewBaseCompositeSloResponse

`func NewBaseCompositeSloResponse() *BaseCompositeSloResponse`

NewBaseCompositeSloResponse instantiates a new BaseCompositeSloResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseCompositeSloResponseWithDefaults

`func NewBaseCompositeSloResponseWithDefaults() *BaseCompositeSloResponse`

NewBaseCompositeSloResponseWithDefaults instantiates a new BaseCompositeSloResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseCompositeSloResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseCompositeSloResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseCompositeSloResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BaseCompositeSloResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BaseCompositeSloResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseCompositeSloResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseCompositeSloResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BaseCompositeSloResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTimeWindow

`func (o *BaseCompositeSloResponse) GetTimeWindow() TimeWindowRolling`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *BaseCompositeSloResponse) GetTimeWindowOk() (*TimeWindowRolling, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *BaseCompositeSloResponse) SetTimeWindow(v TimeWindowRolling)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *BaseCompositeSloResponse) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.

### GetBudgetingMethod

`func (o *BaseCompositeSloResponse) GetBudgetingMethod() BudgetingMethod`

GetBudgetingMethod returns the BudgetingMethod field if non-nil, zero value otherwise.

### GetBudgetingMethodOk

`func (o *BaseCompositeSloResponse) GetBudgetingMethodOk() (*BudgetingMethod, bool)`

GetBudgetingMethodOk returns a tuple with the BudgetingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetingMethod

`func (o *BaseCompositeSloResponse) SetBudgetingMethod(v BudgetingMethod)`

SetBudgetingMethod sets BudgetingMethod field to given value.

### HasBudgetingMethod

`func (o *BaseCompositeSloResponse) HasBudgetingMethod() bool`

HasBudgetingMethod returns a boolean if a field has been set.

### GetCompositeMethod

`func (o *BaseCompositeSloResponse) GetCompositeMethod() CompositeMethod`

GetCompositeMethod returns the CompositeMethod field if non-nil, zero value otherwise.

### GetCompositeMethodOk

`func (o *BaseCompositeSloResponse) GetCompositeMethodOk() (*CompositeMethod, bool)`

GetCompositeMethodOk returns a tuple with the CompositeMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeMethod

`func (o *BaseCompositeSloResponse) SetCompositeMethod(v CompositeMethod)`

SetCompositeMethod sets CompositeMethod field to given value.

### HasCompositeMethod

`func (o *BaseCompositeSloResponse) HasCompositeMethod() bool`

HasCompositeMethod returns a boolean if a field has been set.

### GetObjective

`func (o *BaseCompositeSloResponse) GetObjective() Objective`

GetObjective returns the Objective field if non-nil, zero value otherwise.

### GetObjectiveOk

`func (o *BaseCompositeSloResponse) GetObjectiveOk() (*Objective, bool)`

GetObjectiveOk returns a tuple with the Objective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjective

`func (o *BaseCompositeSloResponse) SetObjective(v Objective)`

SetObjective sets Objective field to given value.

### HasObjective

`func (o *BaseCompositeSloResponse) HasObjective() bool`

HasObjective returns a boolean if a field has been set.

### GetSources

`func (o *BaseCompositeSloResponse) GetSources() CompositeSloResponseSources`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *BaseCompositeSloResponse) GetSourcesOk() (*CompositeSloResponseSources, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *BaseCompositeSloResponse) SetSources(v CompositeSloResponseSources)`

SetSources sets Sources field to given value.

### HasSources

`func (o *BaseCompositeSloResponse) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BaseCompositeSloResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BaseCompositeSloResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BaseCompositeSloResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BaseCompositeSloResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *BaseCompositeSloResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BaseCompositeSloResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BaseCompositeSloResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BaseCompositeSloResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


