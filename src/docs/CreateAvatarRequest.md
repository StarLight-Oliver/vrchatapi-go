# CreateAvatarRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetUrl** | Pointer to **string** |  | [optional] 
**AssetVersion** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** | A date and time of the pattern &#x60;M/d/yyyy h:mm:ss tt&#x60; (see C Sharp &#x60;System.DateTime&#x60;) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ImageUrl** | **string** |  | 
**Name** | **string** |  | 
**Platform** | Pointer to **string** | This is normally &#x60;android&#x60;, &#x60;ios&#x60;, &#x60;standalonewindows&#x60;, &#x60;web&#x60;, or the empty value &#x60;&#x60;, but also supposedly can be any random Unity version such as &#x60;2019.2.4-801-Release&#x60; or &#x60;2019.2.2-772-Release&#x60; or even &#x60;unknownplatform&#x60;. | [optional] 
**ReleaseStatus** | Pointer to [**ReleaseStatus**](ReleaseStatus.md) |  | [optional] [default to PUBLIC]
**Tags** | Pointer to **[]string** |   | [optional] 
**ThumbnailImageUrl** | Pointer to **string** |  | [optional] 
**UnityPackageUrl** | Pointer to **string** |  | [optional] 
**UnityVersion** | Pointer to **string** |  | [optional] [default to "5.3.4p1"]
**UpdatedAt** | Pointer to **string** | A date and time of the pattern &#x60;M/d/yyyy h:mm:ss tt&#x60; (see C Sharp &#x60;System.DateTime&#x60;) | [optional] 
**Version** | Pointer to **int32** |  | [optional] [default to 1]

## Methods

### NewCreateAvatarRequest

`func NewCreateAvatarRequest(imageUrl string, name string, ) *CreateAvatarRequest`

NewCreateAvatarRequest instantiates a new CreateAvatarRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAvatarRequestWithDefaults

`func NewCreateAvatarRequestWithDefaults() *CreateAvatarRequest`

NewCreateAvatarRequestWithDefaults instantiates a new CreateAvatarRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetUrl

`func (o *CreateAvatarRequest) GetAssetUrl() string`

GetAssetUrl returns the AssetUrl field if non-nil, zero value otherwise.

### GetAssetUrlOk

`func (o *CreateAvatarRequest) GetAssetUrlOk() (*string, bool)`

GetAssetUrlOk returns a tuple with the AssetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetUrl

`func (o *CreateAvatarRequest) SetAssetUrl(v string)`

SetAssetUrl sets AssetUrl field to given value.

### HasAssetUrl

`func (o *CreateAvatarRequest) HasAssetUrl() bool`

HasAssetUrl returns a boolean if a field has been set.

### GetAssetVersion

`func (o *CreateAvatarRequest) GetAssetVersion() string`

GetAssetVersion returns the AssetVersion field if non-nil, zero value otherwise.

### GetAssetVersionOk

`func (o *CreateAvatarRequest) GetAssetVersionOk() (*string, bool)`

GetAssetVersionOk returns a tuple with the AssetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetVersion

`func (o *CreateAvatarRequest) SetAssetVersion(v string)`

SetAssetVersion sets AssetVersion field to given value.

### HasAssetVersion

`func (o *CreateAvatarRequest) HasAssetVersion() bool`

HasAssetVersion returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CreateAvatarRequest) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateAvatarRequest) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateAvatarRequest) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateAvatarRequest) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *CreateAvatarRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAvatarRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAvatarRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAvatarRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *CreateAvatarRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateAvatarRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateAvatarRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateAvatarRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageUrl

`func (o *CreateAvatarRequest) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *CreateAvatarRequest) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *CreateAvatarRequest) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.


### GetName

`func (o *CreateAvatarRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAvatarRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAvatarRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPlatform

`func (o *CreateAvatarRequest) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CreateAvatarRequest) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CreateAvatarRequest) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CreateAvatarRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetReleaseStatus

`func (o *CreateAvatarRequest) GetReleaseStatus() ReleaseStatus`

GetReleaseStatus returns the ReleaseStatus field if non-nil, zero value otherwise.

### GetReleaseStatusOk

`func (o *CreateAvatarRequest) GetReleaseStatusOk() (*ReleaseStatus, bool)`

GetReleaseStatusOk returns a tuple with the ReleaseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseStatus

`func (o *CreateAvatarRequest) SetReleaseStatus(v ReleaseStatus)`

SetReleaseStatus sets ReleaseStatus field to given value.

### HasReleaseStatus

`func (o *CreateAvatarRequest) HasReleaseStatus() bool`

HasReleaseStatus returns a boolean if a field has been set.

### GetTags

`func (o *CreateAvatarRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateAvatarRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateAvatarRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateAvatarRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThumbnailImageUrl

`func (o *CreateAvatarRequest) GetThumbnailImageUrl() string`

GetThumbnailImageUrl returns the ThumbnailImageUrl field if non-nil, zero value otherwise.

### GetThumbnailImageUrlOk

`func (o *CreateAvatarRequest) GetThumbnailImageUrlOk() (*string, bool)`

GetThumbnailImageUrlOk returns a tuple with the ThumbnailImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailImageUrl

`func (o *CreateAvatarRequest) SetThumbnailImageUrl(v string)`

SetThumbnailImageUrl sets ThumbnailImageUrl field to given value.

### HasThumbnailImageUrl

`func (o *CreateAvatarRequest) HasThumbnailImageUrl() bool`

HasThumbnailImageUrl returns a boolean if a field has been set.

### GetUnityPackageUrl

`func (o *CreateAvatarRequest) GetUnityPackageUrl() string`

GetUnityPackageUrl returns the UnityPackageUrl field if non-nil, zero value otherwise.

### GetUnityPackageUrlOk

`func (o *CreateAvatarRequest) GetUnityPackageUrlOk() (*string, bool)`

GetUnityPackageUrlOk returns a tuple with the UnityPackageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnityPackageUrl

`func (o *CreateAvatarRequest) SetUnityPackageUrl(v string)`

SetUnityPackageUrl sets UnityPackageUrl field to given value.

### HasUnityPackageUrl

`func (o *CreateAvatarRequest) HasUnityPackageUrl() bool`

HasUnityPackageUrl returns a boolean if a field has been set.

### GetUnityVersion

`func (o *CreateAvatarRequest) GetUnityVersion() string`

GetUnityVersion returns the UnityVersion field if non-nil, zero value otherwise.

### GetUnityVersionOk

`func (o *CreateAvatarRequest) GetUnityVersionOk() (*string, bool)`

GetUnityVersionOk returns a tuple with the UnityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnityVersion

`func (o *CreateAvatarRequest) SetUnityVersion(v string)`

SetUnityVersion sets UnityVersion field to given value.

### HasUnityVersion

`func (o *CreateAvatarRequest) HasUnityVersion() bool`

HasUnityVersion returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CreateAvatarRequest) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateAvatarRequest) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateAvatarRequest) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CreateAvatarRequest) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *CreateAvatarRequest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreateAvatarRequest) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreateAvatarRequest) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CreateAvatarRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


