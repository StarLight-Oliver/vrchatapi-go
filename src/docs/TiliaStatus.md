# TiliaStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EconomyOnline** | **bool** |  | 
**EconomyState** | Pointer to **int32** |  | [optional] 
**PlannedOfflineWindowEnd** | Pointer to **time.Time** |  | [optional] 
**PlannedOfflineWindowStart** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTiliaStatus

`func NewTiliaStatus(economyOnline bool, ) *TiliaStatus`

NewTiliaStatus instantiates a new TiliaStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTiliaStatusWithDefaults

`func NewTiliaStatusWithDefaults() *TiliaStatus`

NewTiliaStatusWithDefaults instantiates a new TiliaStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEconomyOnline

`func (o *TiliaStatus) GetEconomyOnline() bool`

GetEconomyOnline returns the EconomyOnline field if non-nil, zero value otherwise.

### GetEconomyOnlineOk

`func (o *TiliaStatus) GetEconomyOnlineOk() (*bool, bool)`

GetEconomyOnlineOk returns a tuple with the EconomyOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEconomyOnline

`func (o *TiliaStatus) SetEconomyOnline(v bool)`

SetEconomyOnline sets EconomyOnline field to given value.


### GetEconomyState

`func (o *TiliaStatus) GetEconomyState() int32`

GetEconomyState returns the EconomyState field if non-nil, zero value otherwise.

### GetEconomyStateOk

`func (o *TiliaStatus) GetEconomyStateOk() (*int32, bool)`

GetEconomyStateOk returns a tuple with the EconomyState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEconomyState

`func (o *TiliaStatus) SetEconomyState(v int32)`

SetEconomyState sets EconomyState field to given value.

### HasEconomyState

`func (o *TiliaStatus) HasEconomyState() bool`

HasEconomyState returns a boolean if a field has been set.

### GetPlannedOfflineWindowEnd

`func (o *TiliaStatus) GetPlannedOfflineWindowEnd() time.Time`

GetPlannedOfflineWindowEnd returns the PlannedOfflineWindowEnd field if non-nil, zero value otherwise.

### GetPlannedOfflineWindowEndOk

`func (o *TiliaStatus) GetPlannedOfflineWindowEndOk() (*time.Time, bool)`

GetPlannedOfflineWindowEndOk returns a tuple with the PlannedOfflineWindowEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedOfflineWindowEnd

`func (o *TiliaStatus) SetPlannedOfflineWindowEnd(v time.Time)`

SetPlannedOfflineWindowEnd sets PlannedOfflineWindowEnd field to given value.

### HasPlannedOfflineWindowEnd

`func (o *TiliaStatus) HasPlannedOfflineWindowEnd() bool`

HasPlannedOfflineWindowEnd returns a boolean if a field has been set.

### GetPlannedOfflineWindowStart

`func (o *TiliaStatus) GetPlannedOfflineWindowStart() time.Time`

GetPlannedOfflineWindowStart returns the PlannedOfflineWindowStart field if non-nil, zero value otherwise.

### GetPlannedOfflineWindowStartOk

`func (o *TiliaStatus) GetPlannedOfflineWindowStartOk() (*time.Time, bool)`

GetPlannedOfflineWindowStartOk returns a tuple with the PlannedOfflineWindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedOfflineWindowStart

`func (o *TiliaStatus) SetPlannedOfflineWindowStart(v time.Time)`

SetPlannedOfflineWindowStart sets PlannedOfflineWindowStart field to given value.

### HasPlannedOfflineWindowStart

`func (o *TiliaStatus) HasPlannedOfflineWindowStart() bool`

HasPlannedOfflineWindowStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


