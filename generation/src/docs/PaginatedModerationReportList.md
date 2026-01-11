# PaginatedModerationReportList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasNext** | Pointer to **bool** | Whether there are more results after this page. | [optional] 
**Results** | Pointer to [**[]ModerationReport**](ModerationReport.md) | The list of moderation reports. | [optional] 
**TotalCount** | Pointer to **int32** | The total number of results that the query would return if there were no pagination. | [optional] 

## Methods

### NewPaginatedModerationReportList

`func NewPaginatedModerationReportList() *PaginatedModerationReportList`

NewPaginatedModerationReportList instantiates a new PaginatedModerationReportList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedModerationReportListWithDefaults

`func NewPaginatedModerationReportListWithDefaults() *PaginatedModerationReportList`

NewPaginatedModerationReportListWithDefaults instantiates a new PaginatedModerationReportList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasNext

`func (o *PaginatedModerationReportList) GetHasNext() bool`

GetHasNext returns the HasNext field if non-nil, zero value otherwise.

### GetHasNextOk

`func (o *PaginatedModerationReportList) GetHasNextOk() (*bool, bool)`

GetHasNextOk returns a tuple with the HasNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNext

`func (o *PaginatedModerationReportList) SetHasNext(v bool)`

SetHasNext sets HasNext field to given value.

### HasHasNext

`func (o *PaginatedModerationReportList) HasHasNext() bool`

HasHasNext returns a boolean if a field has been set.

### GetResults

`func (o *PaginatedModerationReportList) GetResults() []ModerationReport`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedModerationReportList) GetResultsOk() (*[]ModerationReport, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedModerationReportList) SetResults(v []ModerationReport)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedModerationReportList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotalCount

`func (o *PaginatedModerationReportList) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PaginatedModerationReportList) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PaginatedModerationReportList) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PaginatedModerationReportList) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


