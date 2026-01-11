# Prop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**AuthorId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**AuthorName** | **string** |  | 
**Description** | **string** |  | 
**Id** | **string** |  | 
**ImageUrl** | **string** |  | 
**MaxCountPerUser** | **int32** |  | [default to 1]
**Name** | **string** |  | 
**ReleaseStatus** | [**ReleaseStatus**](ReleaseStatus.md) |  | [default to PUBLIC]
**SpawnType** | **int32** | How a prop is summoned and interacted with. 0: the prop fixed to some surface in the world 1: the prop is a pickup and may be held by users 2: ??? | [default to 1]
**Tags** | **[]string** |  | 
**ThumbnailImageUrl** | **string** |  | 
**UnityPackageUrl** | **NullableString** |  | 
**UnityPackages** | [**[]PropUnityPackage**](PropUnityPackage.md) |  | 
**WorldPlacementMask** | **int32** | Bitmask for restrictions on what world surfaces a prop may be summoned. 0: no restrictions 1: floors 2: walls 4: ceilings | [default to 1]

## Methods

### NewProp

`func NewProp(createdAt time.Time, updatedAt time.Time, authorId string, authorName string, description string, id string, imageUrl string, maxCountPerUser int32, name string, releaseStatus ReleaseStatus, spawnType int32, tags []string, thumbnailImageUrl string, unityPackageUrl NullableString, unityPackages []PropUnityPackage, worldPlacementMask int32, ) *Prop`

NewProp instantiates a new Prop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropWithDefaults

`func NewPropWithDefaults() *Prop`

NewPropWithDefaults instantiates a new Prop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Prop) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Prop) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Prop) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Prop) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Prop) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Prop) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetAuthorId

`func (o *Prop) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *Prop) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *Prop) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.


### GetAuthorName

`func (o *Prop) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *Prop) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *Prop) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.


### GetDescription

`func (o *Prop) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Prop) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Prop) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *Prop) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Prop) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Prop) SetId(v string)`

SetId sets Id field to given value.


### GetImageUrl

`func (o *Prop) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *Prop) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *Prop) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.


### GetMaxCountPerUser

`func (o *Prop) GetMaxCountPerUser() int32`

GetMaxCountPerUser returns the MaxCountPerUser field if non-nil, zero value otherwise.

### GetMaxCountPerUserOk

`func (o *Prop) GetMaxCountPerUserOk() (*int32, bool)`

GetMaxCountPerUserOk returns a tuple with the MaxCountPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCountPerUser

`func (o *Prop) SetMaxCountPerUser(v int32)`

SetMaxCountPerUser sets MaxCountPerUser field to given value.


### GetName

`func (o *Prop) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Prop) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Prop) SetName(v string)`

SetName sets Name field to given value.


### GetReleaseStatus

`func (o *Prop) GetReleaseStatus() ReleaseStatus`

GetReleaseStatus returns the ReleaseStatus field if non-nil, zero value otherwise.

### GetReleaseStatusOk

`func (o *Prop) GetReleaseStatusOk() (*ReleaseStatus, bool)`

GetReleaseStatusOk returns a tuple with the ReleaseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseStatus

`func (o *Prop) SetReleaseStatus(v ReleaseStatus)`

SetReleaseStatus sets ReleaseStatus field to given value.


### GetSpawnType

`func (o *Prop) GetSpawnType() int32`

GetSpawnType returns the SpawnType field if non-nil, zero value otherwise.

### GetSpawnTypeOk

`func (o *Prop) GetSpawnTypeOk() (*int32, bool)`

GetSpawnTypeOk returns a tuple with the SpawnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpawnType

`func (o *Prop) SetSpawnType(v int32)`

SetSpawnType sets SpawnType field to given value.


### GetTags

`func (o *Prop) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Prop) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Prop) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetThumbnailImageUrl

`func (o *Prop) GetThumbnailImageUrl() string`

GetThumbnailImageUrl returns the ThumbnailImageUrl field if non-nil, zero value otherwise.

### GetThumbnailImageUrlOk

`func (o *Prop) GetThumbnailImageUrlOk() (*string, bool)`

GetThumbnailImageUrlOk returns a tuple with the ThumbnailImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailImageUrl

`func (o *Prop) SetThumbnailImageUrl(v string)`

SetThumbnailImageUrl sets ThumbnailImageUrl field to given value.


### GetUnityPackageUrl

`func (o *Prop) GetUnityPackageUrl() string`

GetUnityPackageUrl returns the UnityPackageUrl field if non-nil, zero value otherwise.

### GetUnityPackageUrlOk

`func (o *Prop) GetUnityPackageUrlOk() (*string, bool)`

GetUnityPackageUrlOk returns a tuple with the UnityPackageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnityPackageUrl

`func (o *Prop) SetUnityPackageUrl(v string)`

SetUnityPackageUrl sets UnityPackageUrl field to given value.


### SetUnityPackageUrlNil

`func (o *Prop) SetUnityPackageUrlNil(b bool)`

 SetUnityPackageUrlNil sets the value for UnityPackageUrl to be an explicit nil

### UnsetUnityPackageUrl
`func (o *Prop) UnsetUnityPackageUrl()`

UnsetUnityPackageUrl ensures that no value is present for UnityPackageUrl, not even an explicit nil
### GetUnityPackages

`func (o *Prop) GetUnityPackages() []PropUnityPackage`

GetUnityPackages returns the UnityPackages field if non-nil, zero value otherwise.

### GetUnityPackagesOk

`func (o *Prop) GetUnityPackagesOk() (*[]PropUnityPackage, bool)`

GetUnityPackagesOk returns a tuple with the UnityPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnityPackages

`func (o *Prop) SetUnityPackages(v []PropUnityPackage)`

SetUnityPackages sets UnityPackages field to given value.


### GetWorldPlacementMask

`func (o *Prop) GetWorldPlacementMask() int32`

GetWorldPlacementMask returns the WorldPlacementMask field if non-nil, zero value otherwise.

### GetWorldPlacementMaskOk

`func (o *Prop) GetWorldPlacementMaskOk() (*int32, bool)`

GetWorldPlacementMaskOk returns a tuple with the WorldPlacementMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorldPlacementMask

`func (o *Prop) SetWorldPlacementMask(v int32)`

SetWorldPlacementMask sets WorldPlacementMask field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


