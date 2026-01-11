# CreatePropRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetUrl** | **string** |  | 
**AssetVersion** | **int32** |  | 
**Description** | **string** |  | 
**Id** | **string** |  | 
**ImageUrl** | **string** |  | 
**Name** | **string** |  | 
**Platform** | **string** | This is normally &#x60;android&#x60;, &#x60;ios&#x60;, &#x60;standalonewindows&#x60;, &#x60;web&#x60;, or the empty value &#x60;&#x60;, but also supposedly can be any random Unity version such as &#x60;2019.2.4-801-Release&#x60; or &#x60;2019.2.2-772-Release&#x60; or even &#x60;unknownplatform&#x60;. | 
**PropSignature** | Pointer to **string** |  | [optional] 
**SpawnType** | **int32** | How a prop is summoned and interacted with. 0: the prop fixed to some surface in the world 1: the prop is a pickup and may be held by users 2: ??? | [default to 1]
**Tags** | **[]string** |  | 
**UnityVersion** | **string** |  | 
**WorldPlacementMask** | **int32** | Bitmask for restrictions on what world surfaces a prop may be summoned. 0: no restrictions 1: floors 2: walls 4: ceilings | [default to 1]

## Methods

### NewCreatePropRequest

`func NewCreatePropRequest(assetUrl string, assetVersion int32, description string, id string, imageUrl string, name string, platform string, spawnType int32, tags []string, unityVersion string, worldPlacementMask int32, ) *CreatePropRequest`

NewCreatePropRequest instantiates a new CreatePropRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePropRequestWithDefaults

`func NewCreatePropRequestWithDefaults() *CreatePropRequest`

NewCreatePropRequestWithDefaults instantiates a new CreatePropRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetUrl

`func (o *CreatePropRequest) GetAssetUrl() string`

GetAssetUrl returns the AssetUrl field if non-nil, zero value otherwise.

### GetAssetUrlOk

`func (o *CreatePropRequest) GetAssetUrlOk() (*string, bool)`

GetAssetUrlOk returns a tuple with the AssetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetUrl

`func (o *CreatePropRequest) SetAssetUrl(v string)`

SetAssetUrl sets AssetUrl field to given value.


### GetAssetVersion

`func (o *CreatePropRequest) GetAssetVersion() int32`

GetAssetVersion returns the AssetVersion field if non-nil, zero value otherwise.

### GetAssetVersionOk

`func (o *CreatePropRequest) GetAssetVersionOk() (*int32, bool)`

GetAssetVersionOk returns a tuple with the AssetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetVersion

`func (o *CreatePropRequest) SetAssetVersion(v int32)`

SetAssetVersion sets AssetVersion field to given value.


### GetDescription

`func (o *CreatePropRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatePropRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatePropRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *CreatePropRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreatePropRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreatePropRequest) SetId(v string)`

SetId sets Id field to given value.


### GetImageUrl

`func (o *CreatePropRequest) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *CreatePropRequest) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *CreatePropRequest) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.


### GetName

`func (o *CreatePropRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePropRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePropRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPlatform

`func (o *CreatePropRequest) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CreatePropRequest) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CreatePropRequest) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetPropSignature

`func (o *CreatePropRequest) GetPropSignature() string`

GetPropSignature returns the PropSignature field if non-nil, zero value otherwise.

### GetPropSignatureOk

`func (o *CreatePropRequest) GetPropSignatureOk() (*string, bool)`

GetPropSignatureOk returns a tuple with the PropSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropSignature

`func (o *CreatePropRequest) SetPropSignature(v string)`

SetPropSignature sets PropSignature field to given value.

### HasPropSignature

`func (o *CreatePropRequest) HasPropSignature() bool`

HasPropSignature returns a boolean if a field has been set.

### GetSpawnType

`func (o *CreatePropRequest) GetSpawnType() int32`

GetSpawnType returns the SpawnType field if non-nil, zero value otherwise.

### GetSpawnTypeOk

`func (o *CreatePropRequest) GetSpawnTypeOk() (*int32, bool)`

GetSpawnTypeOk returns a tuple with the SpawnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpawnType

`func (o *CreatePropRequest) SetSpawnType(v int32)`

SetSpawnType sets SpawnType field to given value.


### GetTags

`func (o *CreatePropRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreatePropRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreatePropRequest) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetUnityVersion

`func (o *CreatePropRequest) GetUnityVersion() string`

GetUnityVersion returns the UnityVersion field if non-nil, zero value otherwise.

### GetUnityVersionOk

`func (o *CreatePropRequest) GetUnityVersionOk() (*string, bool)`

GetUnityVersionOk returns a tuple with the UnityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnityVersion

`func (o *CreatePropRequest) SetUnityVersion(v string)`

SetUnityVersion sets UnityVersion field to given value.


### GetWorldPlacementMask

`func (o *CreatePropRequest) GetWorldPlacementMask() int32`

GetWorldPlacementMask returns the WorldPlacementMask field if non-nil, zero value otherwise.

### GetWorldPlacementMaskOk

`func (o *CreatePropRequest) GetWorldPlacementMaskOk() (*int32, bool)`

GetWorldPlacementMaskOk returns a tuple with the WorldPlacementMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorldPlacementMask

`func (o *CreatePropRequest) SetWorldPlacementMask(v int32)`

SetWorldPlacementMask sets WorldPlacementMask field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


