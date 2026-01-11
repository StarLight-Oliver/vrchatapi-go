# ModerationReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **string** | Valid values are the keys of the object &#x60;$.reportOptions[type]&#x60; from &#x60;GET /config&#x60;. Descriptions of these are found at &#x60;$.reportCategories[type]&#x60;. | 
**ContentId** | **string** |  | 
**ContentName** | **string** |  | 
**ContentThumbnailImageUrl** | **NullableString** |  | 
**Description** | **string** | The subjective reason for the report | 
**EvidenceRequired** | **bool** |  | 
**Id** | **string** |  | 
**Reason** | **string** | Valid values are the strings in the array &#x60;$.reportOptions[type][category]&#x60; from &#x60;GET /config&#x60;. Descriptions of these are found at &#x60;$.reportReasons[type]&#x60;. | 
**SupportRequired** | **bool** |  | 
**Type** | **string** | Valid values are the keys of the object &#x60;$.reportOptions&#x60; from &#x60;GET /config&#x60;. | 

## Methods

### NewModerationReport

`func NewModerationReport(category string, contentId string, contentName string, contentThumbnailImageUrl NullableString, description string, evidenceRequired bool, id string, reason string, supportRequired bool, type_ string, ) *ModerationReport`

NewModerationReport instantiates a new ModerationReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModerationReportWithDefaults

`func NewModerationReportWithDefaults() *ModerationReport`

NewModerationReportWithDefaults instantiates a new ModerationReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *ModerationReport) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ModerationReport) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ModerationReport) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetContentId

`func (o *ModerationReport) GetContentId() string`

GetContentId returns the ContentId field if non-nil, zero value otherwise.

### GetContentIdOk

`func (o *ModerationReport) GetContentIdOk() (*string, bool)`

GetContentIdOk returns a tuple with the ContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentId

`func (o *ModerationReport) SetContentId(v string)`

SetContentId sets ContentId field to given value.


### GetContentName

`func (o *ModerationReport) GetContentName() string`

GetContentName returns the ContentName field if non-nil, zero value otherwise.

### GetContentNameOk

`func (o *ModerationReport) GetContentNameOk() (*string, bool)`

GetContentNameOk returns a tuple with the ContentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentName

`func (o *ModerationReport) SetContentName(v string)`

SetContentName sets ContentName field to given value.


### GetContentThumbnailImageUrl

`func (o *ModerationReport) GetContentThumbnailImageUrl() string`

GetContentThumbnailImageUrl returns the ContentThumbnailImageUrl field if non-nil, zero value otherwise.

### GetContentThumbnailImageUrlOk

`func (o *ModerationReport) GetContentThumbnailImageUrlOk() (*string, bool)`

GetContentThumbnailImageUrlOk returns a tuple with the ContentThumbnailImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentThumbnailImageUrl

`func (o *ModerationReport) SetContentThumbnailImageUrl(v string)`

SetContentThumbnailImageUrl sets ContentThumbnailImageUrl field to given value.


### SetContentThumbnailImageUrlNil

`func (o *ModerationReport) SetContentThumbnailImageUrlNil(b bool)`

 SetContentThumbnailImageUrlNil sets the value for ContentThumbnailImageUrl to be an explicit nil

### UnsetContentThumbnailImageUrl
`func (o *ModerationReport) UnsetContentThumbnailImageUrl()`

UnsetContentThumbnailImageUrl ensures that no value is present for ContentThumbnailImageUrl, not even an explicit nil
### GetDescription

`func (o *ModerationReport) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModerationReport) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModerationReport) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEvidenceRequired

`func (o *ModerationReport) GetEvidenceRequired() bool`

GetEvidenceRequired returns the EvidenceRequired field if non-nil, zero value otherwise.

### GetEvidenceRequiredOk

`func (o *ModerationReport) GetEvidenceRequiredOk() (*bool, bool)`

GetEvidenceRequiredOk returns a tuple with the EvidenceRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceRequired

`func (o *ModerationReport) SetEvidenceRequired(v bool)`

SetEvidenceRequired sets EvidenceRequired field to given value.


### GetId

`func (o *ModerationReport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModerationReport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModerationReport) SetId(v string)`

SetId sets Id field to given value.


### GetReason

`func (o *ModerationReport) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ModerationReport) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ModerationReport) SetReason(v string)`

SetReason sets Reason field to given value.


### GetSupportRequired

`func (o *ModerationReport) GetSupportRequired() bool`

GetSupportRequired returns the SupportRequired field if non-nil, zero value otherwise.

### GetSupportRequiredOk

`func (o *ModerationReport) GetSupportRequiredOk() (*bool, bool)`

GetSupportRequiredOk returns a tuple with the SupportRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportRequired

`func (o *ModerationReport) SetSupportRequired(v bool)`

SetSupportRequired sets SupportRequired field to given value.


### GetType

`func (o *ModerationReport) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModerationReport) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModerationReport) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


