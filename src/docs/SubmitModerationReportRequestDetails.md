# SubmitModerationReportRequestDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | Pointer to **string** |  | [optional] 
**HolderId** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 
**ImageType** | Pointer to **string** | Relevant detail for reports about image content, such as emoji. | [optional] 
**InstanceAgeGated** | Pointer to **bool** | Relevant detail for reports taking place from within an instance. | [optional] 
**InstanceType** | Pointer to **string** | Relevant detail for reports taking place from within an instance. | [optional] 
**SuggestedWarnings** | Pointer to [**[]ContentFilter**](ContentFilter.md) | Relevant detail for reports about content that might not be tagged properly. | [optional] 
**UserInSameInstance** | Pointer to **bool** | Relevant detail for reports involving another user in the same instance world. | [optional] 

## Methods

### NewSubmitModerationReportRequestDetails

`func NewSubmitModerationReportRequestDetails() *SubmitModerationReportRequestDetails`

NewSubmitModerationReportRequestDetails instantiates a new SubmitModerationReportRequestDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitModerationReportRequestDetailsWithDefaults

`func NewSubmitModerationReportRequestDetailsWithDefaults() *SubmitModerationReportRequestDetails`

NewSubmitModerationReportRequestDetailsWithDefaults instantiates a new SubmitModerationReportRequestDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *SubmitModerationReportRequestDetails) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *SubmitModerationReportRequestDetails) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *SubmitModerationReportRequestDetails) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *SubmitModerationReportRequestDetails) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetHolderId

`func (o *SubmitModerationReportRequestDetails) GetHolderId() string`

GetHolderId returns the HolderId field if non-nil, zero value otherwise.

### GetHolderIdOk

`func (o *SubmitModerationReportRequestDetails) GetHolderIdOk() (*string, bool)`

GetHolderIdOk returns a tuple with the HolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderId

`func (o *SubmitModerationReportRequestDetails) SetHolderId(v string)`

SetHolderId sets HolderId field to given value.

### HasHolderId

`func (o *SubmitModerationReportRequestDetails) HasHolderId() bool`

HasHolderId returns a boolean if a field has been set.

### GetImageType

`func (o *SubmitModerationReportRequestDetails) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *SubmitModerationReportRequestDetails) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *SubmitModerationReportRequestDetails) SetImageType(v string)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *SubmitModerationReportRequestDetails) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### GetInstanceAgeGated

`func (o *SubmitModerationReportRequestDetails) GetInstanceAgeGated() bool`

GetInstanceAgeGated returns the InstanceAgeGated field if non-nil, zero value otherwise.

### GetInstanceAgeGatedOk

`func (o *SubmitModerationReportRequestDetails) GetInstanceAgeGatedOk() (*bool, bool)`

GetInstanceAgeGatedOk returns a tuple with the InstanceAgeGated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceAgeGated

`func (o *SubmitModerationReportRequestDetails) SetInstanceAgeGated(v bool)`

SetInstanceAgeGated sets InstanceAgeGated field to given value.

### HasInstanceAgeGated

`func (o *SubmitModerationReportRequestDetails) HasInstanceAgeGated() bool`

HasInstanceAgeGated returns a boolean if a field has been set.

### GetInstanceType

`func (o *SubmitModerationReportRequestDetails) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *SubmitModerationReportRequestDetails) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *SubmitModerationReportRequestDetails) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *SubmitModerationReportRequestDetails) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetSuggestedWarnings

`func (o *SubmitModerationReportRequestDetails) GetSuggestedWarnings() []ContentFilter`

GetSuggestedWarnings returns the SuggestedWarnings field if non-nil, zero value otherwise.

### GetSuggestedWarningsOk

`func (o *SubmitModerationReportRequestDetails) GetSuggestedWarningsOk() (*[]ContentFilter, bool)`

GetSuggestedWarningsOk returns a tuple with the SuggestedWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedWarnings

`func (o *SubmitModerationReportRequestDetails) SetSuggestedWarnings(v []ContentFilter)`

SetSuggestedWarnings sets SuggestedWarnings field to given value.

### HasSuggestedWarnings

`func (o *SubmitModerationReportRequestDetails) HasSuggestedWarnings() bool`

HasSuggestedWarnings returns a boolean if a field has been set.

### GetUserInSameInstance

`func (o *SubmitModerationReportRequestDetails) GetUserInSameInstance() bool`

GetUserInSameInstance returns the UserInSameInstance field if non-nil, zero value otherwise.

### GetUserInSameInstanceOk

`func (o *SubmitModerationReportRequestDetails) GetUserInSameInstanceOk() (*bool, bool)`

GetUserInSameInstanceOk returns a tuple with the UserInSameInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInSameInstance

`func (o *SubmitModerationReportRequestDetails) SetUserInSameInstance(v bool)`

SetUserInSameInstance sets UserInSameInstance field to given value.

### HasUserInSameInstance

`func (o *SubmitModerationReportRequestDetails) HasUserInSameInstance() bool`

HasUserInSameInstance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


