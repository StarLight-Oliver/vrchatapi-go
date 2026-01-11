# SubmitModerationReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **string** | Valid values are the keys of the object &#x60;$.reportOptions[type]&#x60; from &#x60;GET /config&#x60;. Descriptions of these are found at &#x60;$.reportCategories[type]&#x60;. | 
**ContentId** | **string** | The id of the user, group, world, avatar, inventory item, print, etc. being reported. | 
**Description** | Pointer to **string** | The subjective reason for the report | [optional] 
**Details** | Pointer to [**SubmitModerationReportRequestDetails**](SubmitModerationReportRequestDetails.md) |  | [optional] 
**Reason** | **string** | Valid values are the strings in the array &#x60;$.reportOptions[type][category]&#x60; from &#x60;GET /config&#x60;. Descriptions of these are found at &#x60;$.reportReasons[type]&#x60;. | 
**Type** | **string** | Valid values are the keys of the object &#x60;$.reportOptions&#x60; from &#x60;GET /config&#x60;. | 

## Methods

### NewSubmitModerationReportRequest

`func NewSubmitModerationReportRequest(category string, contentId string, reason string, type_ string, ) *SubmitModerationReportRequest`

NewSubmitModerationReportRequest instantiates a new SubmitModerationReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitModerationReportRequestWithDefaults

`func NewSubmitModerationReportRequestWithDefaults() *SubmitModerationReportRequest`

NewSubmitModerationReportRequestWithDefaults instantiates a new SubmitModerationReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *SubmitModerationReportRequest) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SubmitModerationReportRequest) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SubmitModerationReportRequest) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetContentId

`func (o *SubmitModerationReportRequest) GetContentId() string`

GetContentId returns the ContentId field if non-nil, zero value otherwise.

### GetContentIdOk

`func (o *SubmitModerationReportRequest) GetContentIdOk() (*string, bool)`

GetContentIdOk returns a tuple with the ContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentId

`func (o *SubmitModerationReportRequest) SetContentId(v string)`

SetContentId sets ContentId field to given value.


### GetDescription

`func (o *SubmitModerationReportRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubmitModerationReportRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubmitModerationReportRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SubmitModerationReportRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDetails

`func (o *SubmitModerationReportRequest) GetDetails() SubmitModerationReportRequestDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *SubmitModerationReportRequest) GetDetailsOk() (*SubmitModerationReportRequestDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *SubmitModerationReportRequest) SetDetails(v SubmitModerationReportRequestDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *SubmitModerationReportRequest) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetReason

`func (o *SubmitModerationReportRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SubmitModerationReportRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SubmitModerationReportRequest) SetReason(v string)`

SetReason sets Reason field to given value.


### GetType

`func (o *SubmitModerationReportRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubmitModerationReportRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubmitModerationReportRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


