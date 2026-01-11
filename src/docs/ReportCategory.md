# ReportCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the report category | [optional] 
**Text** | **string** | The label of the report category | 
**Title** | Pointer to **string** | The title of the report category | [optional] 
**Tooltip** | **string** | The tooltip that describes the category | 

## Methods

### NewReportCategory

`func NewReportCategory(text string, tooltip string, ) *ReportCategory`

NewReportCategory instantiates a new ReportCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportCategoryWithDefaults

`func NewReportCategoryWithDefaults() *ReportCategory`

NewReportCategoryWithDefaults instantiates a new ReportCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ReportCategory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReportCategory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReportCategory) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReportCategory) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetText

`func (o *ReportCategory) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ReportCategory) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ReportCategory) SetText(v string)`

SetText sets Text field to given value.


### GetTitle

`func (o *ReportCategory) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ReportCategory) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ReportCategory) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ReportCategory) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTooltip

`func (o *ReportCategory) GetTooltip() string`

GetTooltip returns the Tooltip field if non-nil, zero value otherwise.

### GetTooltipOk

`func (o *ReportCategory) GetTooltipOk() (*string, bool)`

GetTooltipOk returns a tuple with the Tooltip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTooltip

`func (o *ReportCategory) SetTooltip(v string)`

SetTooltip sets Tooltip field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


