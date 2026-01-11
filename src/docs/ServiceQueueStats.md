# ServiceQueueStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EstimatedServiceDurationSeconds** | **int32** |  | 

## Methods

### NewServiceQueueStats

`func NewServiceQueueStats(estimatedServiceDurationSeconds int32, ) *ServiceQueueStats`

NewServiceQueueStats instantiates a new ServiceQueueStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceQueueStatsWithDefaults

`func NewServiceQueueStatsWithDefaults() *ServiceQueueStats`

NewServiceQueueStatsWithDefaults instantiates a new ServiceQueueStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEstimatedServiceDurationSeconds

`func (o *ServiceQueueStats) GetEstimatedServiceDurationSeconds() int32`

GetEstimatedServiceDurationSeconds returns the EstimatedServiceDurationSeconds field if non-nil, zero value otherwise.

### GetEstimatedServiceDurationSecondsOk

`func (o *ServiceQueueStats) GetEstimatedServiceDurationSecondsOk() (*int32, bool)`

GetEstimatedServiceDurationSecondsOk returns a tuple with the EstimatedServiceDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedServiceDurationSeconds

`func (o *ServiceQueueStats) SetEstimatedServiceDurationSeconds(v int32)`

SetEstimatedServiceDurationSeconds sets EstimatedServiceDurationSeconds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


