# ReportReason

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | The label or name of the report reason | 
**Tooltip** | **string** | A brief explanation of what this reason entails | 

## Methods

### NewReportReason

`func NewReportReason(text string, tooltip string, ) *ReportReason`

NewReportReason instantiates a new ReportReason object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportReasonWithDefaults

`func NewReportReasonWithDefaults() *ReportReason`

NewReportReasonWithDefaults instantiates a new ReportReason object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *ReportReason) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ReportReason) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ReportReason) SetText(v string)`

SetText sets Text field to given value.


### GetTooltip

`func (o *ReportReason) GetTooltip() string`

GetTooltip returns the Tooltip field if non-nil, zero value otherwise.

### GetTooltipOk

`func (o *ReportReason) GetTooltipOk() (*string, bool)`

GetTooltipOk returns a tuple with the Tooltip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTooltip

`func (o *ReportReason) SetTooltip(v string)`

SetTooltip sets Tooltip field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


