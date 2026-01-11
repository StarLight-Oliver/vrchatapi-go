# PaginatedCalendarEventList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNext** | Pointer to **bool** | Whether there are more results after this page. | [optional] 
**Results** | Pointer to [**[]CalendarEvent**](CalendarEvent.md) |   | [optional] 
**TotalCount** | Pointer to **int32** | The total number of results that the query would return if there were no pagination. | [optional] 

## Methods

### NewPaginatedCalendarEventList

`func NewPaginatedCalendarEventList() *PaginatedCalendarEventList`

NewPaginatedCalendarEventList instantiates a new PaginatedCalendarEventList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedCalendarEventListWithDefaults

`func NewPaginatedCalendarEventListWithDefaults() *PaginatedCalendarEventList`

NewPaginatedCalendarEventListWithDefaults instantiates a new PaginatedCalendarEventList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNext

`func (o *PaginatedCalendarEventList) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *PaginatedCalendarEventList) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *PaginatedCalendarEventList) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *PaginatedCalendarEventList) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.

### GetResults

`func (o *PaginatedCalendarEventList) GetResults() []CalendarEvent`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedCalendarEventList) GetResultsOk() (*[]CalendarEvent, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedCalendarEventList) SetResults(v []CalendarEvent)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedCalendarEventList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotalCount

`func (o *PaginatedCalendarEventList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PaginatedCalendarEventList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PaginatedCalendarEventList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PaginatedCalendarEventList) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


