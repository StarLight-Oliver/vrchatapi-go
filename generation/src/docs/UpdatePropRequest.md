# UpdatePropRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetUrl** | Pointer to **string** |  | [optional] 
**AssetVersion** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ImageUrl** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** | This is normally &#x60;android&#x60;, &#x60;ios&#x60;, &#x60;standalonewindows&#x60;, &#x60;web&#x60;, or the empty value &#x60;&#x60;, but also supposedly can be any random Unity version such as &#x60;2019.2.4-801-Release&#x60; or &#x60;2019.2.2-772-Release&#x60; or even &#x60;unknownplatform&#x60;. | [optional] 
**PropSignature** | Pointer to **string** |  | [optional] 
**SpawnType** | Pointer to **int32** | How a prop is summoned and interacted with. 0: the prop fixed to some surface in the world 1: the prop is a pickup and may be held by users 2: ??? | [optional] [default to 1]
**Tags** | Pointer to **[]string** |  | [optional] 
**UnityVersion** | Pointer to **string** |  | [optional] 
**WorldPlacementMask** | Pointer to **int32** | Bitmask for restrictions on what world surfaces a prop may be summoned. 0: no restrictions 1: floors 2: walls 4: ceilings | [optional] [default to 1]

## Methods

### NewUpdatePropRequest

`func NewUpdatePropRequest() *UpdatePropRequest`

NewUpdatePropRequest instantiates a new UpdatePropRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePropRequestWithDefaults

`func NewUpdatePropRequestWithDefaults() *UpdatePropRequest`

NewUpdatePropRequestWithDefaults instantiates a new UpdatePropRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetUrl

`func (o *UpdatePropRequest) GetAssetUrl() string`

GetAssetUrl returns the AssetUrl field if non-nil, zero value otherwise.

### GetAssetUrlOk

`func (o *UpdatePropRequest) GetAssetUrlOk() (*string, bool)`

GetAssetUrlOk returns a tuple with the AssetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetUrl

`func (o *UpdatePropRequest) SetAssetUrl(v string)`

SetAssetUrl sets AssetUrl field to given value.

### HasAssetUrl

`func (o *UpdatePropRequest) HasAssetUrl() bool`

HasAssetUrl returns a boolean if a field has been set.

### GetAssetVersion

`func (o *UpdatePropRequest) GetAssetVersion() int32`

GetAssetVersion returns the AssetVersion field if non-nil, zero value otherwise.

### GetAssetVersionOk

`func (o *UpdatePropRequest) GetAssetVersionOk() (*int32, bool)`

GetAssetVersionOk returns a tuple with the AssetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetVersion

`func (o *UpdatePropRequest) SetAssetVersion(v int32)`

SetAssetVersion sets AssetVersion field to given value.

### HasAssetVersion

`func (o *UpdatePropRequest) HasAssetVersion() bool`

HasAssetVersion returns a boolean if a field has been set.

### GetDescription

`func (o *UpdatePropRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdatePropRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdatePropRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdatePropRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImageUrl

`func (o *UpdatePropRequest) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *UpdatePropRequest) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *UpdatePropRequest) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *UpdatePropRequest) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetName

`func (o *UpdatePropRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePropRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePropRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdatePropRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatform

`func (o *UpdatePropRequest) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *UpdatePropRequest) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *UpdatePropRequest) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *UpdatePropRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPropSignature

`func (o *UpdatePropRequest) GetPropSignature() string`

GetPropSignature returns the PropSignature field if non-nil, zero value otherwise.

### GetPropSignatureOk

`func (o *UpdatePropRequest) GetPropSignatureOk() (*string, bool)`

GetPropSignatureOk returns a tuple with the PropSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropSignature

`func (o *UpdatePropRequest) SetPropSignature(v string)`

SetPropSignature sets PropSignature field to given value.

### HasPropSignature

`func (o *UpdatePropRequest) HasPropSignature() bool`

HasPropSignature returns a boolean if a field has been set.

### GetSpawnType

`func (o *UpdatePropRequest) GetSpawnType() int32`

GetSpawnType returns the SpawnType field if non-nil, zero value otherwise.

### GetSpawnTypeOk

`func (o *UpdatePropRequest) GetSpawnTypeOk() (*int32, bool)`

GetSpawnTypeOk returns a tuple with the SpawnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpawnType

`func (o *UpdatePropRequest) SetSpawnType(v int32)`

SetSpawnType sets SpawnType field to given value.

### HasSpawnType

`func (o *UpdatePropRequest) HasSpawnType() bool`

HasSpawnType returns a boolean if a field has been set.

### GetTags

`func (o *UpdatePropRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdatePropRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdatePropRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdatePropRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUnityVersion

`func (o *UpdatePropRequest) GetUnityVersion() string`

GetUnityVersion returns the UnityVersion field if non-nil, zero value otherwise.

### GetUnityVersionOk

`func (o *UpdatePropRequest) GetUnityVersionOk() (*string, bool)`

GetUnityVersionOk returns a tuple with the UnityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnityVersion

`func (o *UpdatePropRequest) SetUnityVersion(v string)`

SetUnityVersion sets UnityVersion field to given value.

### HasUnityVersion

`func (o *UpdatePropRequest) HasUnityVersion() bool`

HasUnityVersion returns a boolean if a field has been set.

### GetWorldPlacementMask

`func (o *UpdatePropRequest) GetWorldPlacementMask() int32`

GetWorldPlacementMask returns the WorldPlacementMask field if non-nil, zero value otherwise.

### GetWorldPlacementMaskOk

`func (o *UpdatePropRequest) GetWorldPlacementMaskOk() (*int32, bool)`

GetWorldPlacementMaskOk returns a tuple with the WorldPlacementMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorldPlacementMask

`func (o *UpdatePropRequest) SetWorldPlacementMask(v int32)`

SetWorldPlacementMask sets WorldPlacementMask field to given value.

### HasWorldPlacementMask

`func (o *UpdatePropRequest) HasWorldPlacementMask() bool`

HasWorldPlacementMask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


