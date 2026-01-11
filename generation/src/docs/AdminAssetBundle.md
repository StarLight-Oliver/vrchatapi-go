# AdminAssetBundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**AssetType** | **string** |  | 
**AuthorId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**AuthorName** | **string** |  | 
**Description** | **string** |  | 
**ImageUrl** | **string** |  | 
**Name** | **string** |  | 
**ReleaseStatus** | [**ReleaseStatus**](ReleaseStatus.md) |  | [default to PUBLIC]
**Tags** | **[]string** |  | 
**ThumbnailImageUrl** | **string** |  | 
**UnityPackageUrl** | **NullableString** |  | 
**UnityPackages** | [**[]AdminUnityPackage**](AdminUnityPackage.md) |  | 

## Methods

### NewAdminAssetBundle

`func NewAdminAssetBundle(createdAt time.Time, updatedAt time.Time, assetType string, authorId string, authorName string, description string, imageUrl string, name string, releaseStatus ReleaseStatus, tags []string, thumbnailImageUrl string, unityPackageUrl NullableString, unityPackages []AdminUnityPackage, ) *AdminAssetBundle`

NewAdminAssetBundle instantiates a new AdminAssetBundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminAssetBundleWithDefaults

`func NewAdminAssetBundleWithDefaults() *AdminAssetBundle`

NewAdminAssetBundleWithDefaults instantiates a new AdminAssetBundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *AdminAssetBundle) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdminAssetBundle) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdminAssetBundle) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AdminAssetBundle) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdminAssetBundle) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdminAssetBundle) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetAssetType

`func (o *AdminAssetBundle) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *AdminAssetBundle) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *AdminAssetBundle) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.


### GetAuthorId

`func (o *AdminAssetBundle) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *AdminAssetBundle) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *AdminAssetBundle) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.


### GetAuthorName

`func (o *AdminAssetBundle) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *AdminAssetBundle) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *AdminAssetBundle) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.


### GetDescription

`func (o *AdminAssetBundle) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdminAssetBundle) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdminAssetBundle) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetImageUrl

`func (o *AdminAssetBundle) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *AdminAssetBundle) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *AdminAssetBundle) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.


### GetName

`func (o *AdminAssetBundle) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdminAssetBundle) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdminAssetBundle) SetName(v string)`

SetName sets Name field to given value.


### GetReleaseStatus

`func (o *AdminAssetBundle) GetReleaseStatus() ReleaseStatus`

GetReleaseStatus returns the ReleaseStatus field if non-nil, zero value otherwise.

### GetReleaseStatusOk

`func (o *AdminAssetBundle) GetReleaseStatusOk() (*ReleaseStatus, bool)`

GetReleaseStatusOk returns a tuple with the ReleaseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseStatus

`func (o *AdminAssetBundle) SetReleaseStatus(v ReleaseStatus)`

SetReleaseStatus sets ReleaseStatus field to given value.


### GetTags

`func (o *AdminAssetBundle) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AdminAssetBundle) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AdminAssetBundle) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetThumbnailImageUrl

`func (o *AdminAssetBundle) GetThumbnailImageUrl() string`

GetThumbnailImageUrl returns the ThumbnailImageUrl field if non-nil, zero value otherwise.

### GetThumbnailImageUrlOk

`func (o *AdminAssetBundle) GetThumbnailImageUrlOk() (*string, bool)`

GetThumbnailImageUrlOk returns a tuple with the ThumbnailImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailImageUrl

`func (o *AdminAssetBundle) SetThumbnailImageUrl(v string)`

SetThumbnailImageUrl sets ThumbnailImageUrl field to given value.


### GetUnityPackageUrl

`func (o *AdminAssetBundle) GetUnityPackageUrl() string`

GetUnityPackageUrl returns the UnityPackageUrl field if non-nil, zero value otherwise.

### GetUnityPackageUrlOk

`func (o *AdminAssetBundle) GetUnityPackageUrlOk() (*string, bool)`

GetUnityPackageUrlOk returns a tuple with the UnityPackageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnityPackageUrl

`func (o *AdminAssetBundle) SetUnityPackageUrl(v string)`

SetUnityPackageUrl sets UnityPackageUrl field to given value.


### SetUnityPackageUrlNil

`func (o *AdminAssetBundle) SetUnityPackageUrlNil(b bool)`

 SetUnityPackageUrlNil sets the value for UnityPackageUrl to be an explicit nil

### UnsetUnityPackageUrl
`func (o *AdminAssetBundle) UnsetUnityPackageUrl()`

UnsetUnityPackageUrl ensures that no value is present for UnityPackageUrl, not even an explicit nil
### GetUnityPackages

`func (o *AdminAssetBundle) GetUnityPackages() []AdminUnityPackage`

GetUnityPackages returns the UnityPackages field if non-nil, zero value otherwise.

### GetUnityPackagesOk

`func (o *AdminAssetBundle) GetUnityPackagesOk() (*[]AdminUnityPackage, bool)`

GetUnityPackagesOk returns a tuple with the UnityPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnityPackages

`func (o *AdminAssetBundle) SetUnityPackages(v []AdminUnityPackage)`

SetUnityPackages sets UnityPackages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


