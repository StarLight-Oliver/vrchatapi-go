# CalendarEventDiscovery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextCursor** | **string** | Base64-encoded JSON:   type: object   properties:     dataSource:       type: string       enum:         - featured         - personalized     dataIndex:       type: integer       format: int32     phase:       type: string       enum:         - all         - live         - upcoming       description: see CalendarEventDiscoveryScope     asOf:       type: integer       format: int64       description: milliseconds since Unix epoch     paramHash:       type: string       format: string       description: Base64-encoded 256-bit hash of the original query parameters | 
**Results** | [**[]CalendarEvent**](CalendarEvent.md) |  | 

## Methods

### NewCalendarEventDiscovery

`func NewCalendarEventDiscovery(nextCursor string, results []CalendarEvent, ) *CalendarEventDiscovery`

NewCalendarEventDiscovery instantiates a new CalendarEventDiscovery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalendarEventDiscoveryWithDefaults

`func NewCalendarEventDiscoveryWithDefaults() *CalendarEventDiscovery`

NewCalendarEventDiscoveryWithDefaults instantiates a new CalendarEventDiscovery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextCursor

`func (o *CalendarEventDiscovery) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *CalendarEventDiscovery) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *CalendarEventDiscovery) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### GetResults

`func (o *CalendarEventDiscovery) GetResults() []CalendarEvent`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CalendarEventDiscovery) GetResultsOk() (*[]CalendarEvent, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CalendarEventDiscovery) SetResults(v []CalendarEvent)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


