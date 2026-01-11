# Avatar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acknowledgements** | Pointer to **string** |  | [optional] 
**ActiveAssetReviewId** | Pointer to **string** | Only present for the avatar author on avatars under active review. | [optional] 
**AssetUrl** | Pointer to **string** | Not present from general search &#x60;/avatars&#x60;, only on specific requests &#x60;/avatars/{avatarId}&#x60;. | [optional] 
**AssetUrlObject** | Pointer to **map[string]interface{}** | Not present from general search &#x60;/avatars&#x60;, only on specific requests &#x60;/avatars/{avatarId}&#x60;. **Deprecation:** &#x60;Object&#x60; has unknown usage/fields, and is always empty. Use normal &#x60;Url&#x60; field instead. | [optional] 
**AuthorId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**AuthorName** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Description** | **string** |  | 
**Featured** | **bool** |  | [default to false]
**HighestPrice** | Pointer to **int32** |  | [optional] 
**Id** | **string** |  | 
**ImageUrl** | **string** |  | 
**ListingDate** | **NullableString** |  | 
**Lock** | Pointer to **bool** |  | [optional] 
**LowestPrice** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**PendingUpload** | Pointer to **bool** |  | [optional] [default to false]
**Performance** | [**AvatarPerformance**](AvatarPerformance.md) |  | 
**ProductId** | Pointer to **string** |  | [optional] 
**PublishedListings** | Pointer to [**[]AvatarPublishedListingsInner**](AvatarPublishedListingsInner.md) |  | [optional] 
**ReleaseStatus** | [**ReleaseStatus**](ReleaseStatus.md) |  | [default to PUBLIC]
**Searchable** | Pointer to **bool** |  | [optional] [default to false]
**Styles** | [**AvatarStyles**](AvatarStyles.md) |  | 
**Tags** | **[]string** |   | 
**ThumbnailImageUrl** | **string** |  | 
**UnityPackageUrl** | **string** |  | 
**UnityPackageUrlObject** | [**AvatarUnityPackageUrlObject**](AvatarUnityPackageUrlObject.md) |  | 
**UnityPackages** | [**[]UnityPackage**](UnityPackage.md) |  | 
**UpdatedAt** | **time.Time** |  | 
**Version** | **int32** |  | [default to 0]

## Methods

### NewAvatar

`func NewAvatar(authorId string, authorName string, createdAt time.Time, description string, featured bool, id string, imageUrl string, listingDate NullableString, name string, performance AvatarPerformance, releaseStatus ReleaseStatus, styles AvatarStyles, tags []string, thumbnailImageUrl string, unityPackageUrl string, unityPackageUrlObject AvatarUnityPackageUrlObject, unityPackages []UnityPackage, updatedAt time.Time, version int32, ) *Avatar`

NewAvatar instantiates a new Avatar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvatarWithDefaults

`func NewAvatarWithDefaults() *Avatar`

NewAvatarWithDefaults instantiates a new Avatar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcknowledgements

`func (o *Avatar) GetAcknowledgements() string`

GetAcknowledgements returns the Acknowledgements field if non-nil, zero value otherwise.

### GetAcknowledgementsOk

`func (o *Avatar) GetAcknowledgementsOk() (*string, bool)`

GetAcknowledgementsOk returns a tuple with the Acknowledgements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgements

`func (o *Avatar) SetAcknowledgements(v string)`

SetAcknowledgements sets Acknowledgements field to given value.

### HasAcknowledgements

`func (o *Avatar) HasAcknowledgements() bool`

HasAcknowledgements returns a boolean if a field has been set.

### GetActiveAssetReviewId

`func (o *Avatar) GetActiveAssetReviewId() string`

GetActiveAssetReviewId returns the ActiveAssetReviewId field if non-nil, zero value otherwise.

### GetActiveAssetReviewIdOk

`func (o *Avatar) GetActiveAssetReviewIdOk() (*string, bool)`

GetActiveAssetReviewIdOk returns a tuple with the ActiveAssetReviewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAssetReviewId

`func (o *Avatar) SetActiveAssetReviewId(v string)`

SetActiveAssetReviewId sets ActiveAssetReviewId field to given value.

### HasActiveAssetReviewId

`func (o *Avatar) HasActiveAssetReviewId() bool`

HasActiveAssetReviewId returns a boolean if a field has been set.

### GetAssetUrl

`func (o *Avatar) GetAssetUrl() string`

GetAssetUrl returns the AssetUrl field if non-nil, zero value otherwise.

### GetAssetUrlOk

`func (o *Avatar) GetAssetUrlOk() (*string, bool)`

GetAssetUrlOk returns a tuple with the AssetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetUrl

`func (o *Avatar) SetAssetUrl(v string)`

SetAssetUrl sets AssetUrl field to given value.

### HasAssetUrl

`func (o *Avatar) HasAssetUrl() bool`

HasAssetUrl returns a boolean if a field has been set.

### GetAssetUrlObject

`func (o *Avatar) GetAssetUrlObject() map[string]interface{}`

GetAssetUrlObject returns the AssetUrlObject field if non-nil, zero value otherwise.

### GetAssetUrlObjectOk

`func (o *Avatar) GetAssetUrlObjectOk() (*map[string]interface{}, bool)`

GetAssetUrlObjectOk returns a tuple with the AssetUrlObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetUrlObject

`func (o *Avatar) SetAssetUrlObject(v map[string]interface{})`

SetAssetUrlObject sets AssetUrlObject field to given value.

### HasAssetUrlObject

`func (o *Avatar) HasAssetUrlObject() bool`

HasAssetUrlObject returns a boolean if a field has been set.

### GetAuthorId

`func (o *Avatar) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *Avatar) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *Avatar) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.


### GetAuthorName

`func (o *Avatar) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *Avatar) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *Avatar) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.


### GetCreatedAt

`func (o *Avatar) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Avatar) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Avatar) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *Avatar) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Avatar) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Avatar) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFeatured

`func (o *Avatar) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *Avatar) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *Avatar) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.


### GetHighestPrice

`func (o *Avatar) GetHighestPrice() int32`

GetHighestPrice returns the HighestPrice field if non-nil, zero value otherwise.

### GetHighestPriceOk

`func (o *Avatar) GetHighestPriceOk() (*int32, bool)`

GetHighestPriceOk returns a tuple with the HighestPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestPrice

`func (o *Avatar) SetHighestPrice(v int32)`

SetHighestPrice sets HighestPrice field to given value.

### HasHighestPrice

`func (o *Avatar) HasHighestPrice() bool`

HasHighestPrice returns a boolean if a field has been set.

### GetId

`func (o *Avatar) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Avatar) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Avatar) SetId(v string)`

SetId sets Id field to given value.


### GetImageUrl

`func (o *Avatar) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *Avatar) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *Avatar) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.


### GetListingDate

`func (o *Avatar) GetListingDate() string`

GetListingDate returns the ListingDate field if non-nil, zero value otherwise.

### GetListingDateOk

`func (o *Avatar) GetListingDateOk() (*string, bool)`

GetListingDateOk returns a tuple with the ListingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingDate

`func (o *Avatar) SetListingDate(v string)`

SetListingDate sets ListingDate field to given value.


### SetListingDateNil

`func (o *Avatar) SetListingDateNil(b bool)`

 SetListingDateNil sets the value for ListingDate to be an explicit nil

### UnsetListingDate
`func (o *Avatar) UnsetListingDate()`

UnsetListingDate ensures that no value is present for ListingDate, not even an explicit nil
### GetLock

`func (o *Avatar) GetLock() bool`

GetLock returns the Lock field if non-nil, zero value otherwise.

### GetLockOk

`func (o *Avatar) GetLockOk() (*bool, bool)`

GetLockOk returns a tuple with the Lock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLock

`func (o *Avatar) SetLock(v bool)`

SetLock sets Lock field to given value.

### HasLock

`func (o *Avatar) HasLock() bool`

HasLock returns a boolean if a field has been set.

### GetLowestPrice

`func (o *Avatar) GetLowestPrice() int32`

GetLowestPrice returns the LowestPrice field if non-nil, zero value otherwise.

### GetLowestPriceOk

`func (o *Avatar) GetLowestPriceOk() (*int32, bool)`

GetLowestPriceOk returns a tuple with the LowestPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowestPrice

`func (o *Avatar) SetLowestPrice(v int32)`

SetLowestPrice sets LowestPrice field to given value.

### HasLowestPrice

`func (o *Avatar) HasLowestPrice() bool`

HasLowestPrice returns a boolean if a field has been set.

### GetName

`func (o *Avatar) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Avatar) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Avatar) SetName(v string)`

SetName sets Name field to given value.


### GetPendingUpload

`func (o *Avatar) GetPendingUpload() bool`

GetPendingUpload returns the PendingUpload field if non-nil, zero value otherwise.

### GetPendingUploadOk

`func (o *Avatar) GetPendingUploadOk() (*bool, bool)`

GetPendingUploadOk returns a tuple with the PendingUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingUpload

`func (o *Avatar) SetPendingUpload(v bool)`

SetPendingUpload sets PendingUpload field to given value.

### HasPendingUpload

`func (o *Avatar) HasPendingUpload() bool`

HasPendingUpload returns a boolean if a field has been set.

### GetPerformance

`func (o *Avatar) GetPerformance() AvatarPerformance`

GetPerformance returns the Performance field if non-nil, zero value otherwise.

### GetPerformanceOk

`func (o *Avatar) GetPerformanceOk() (*AvatarPerformance, bool)`

GetPerformanceOk returns a tuple with the Performance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformance

`func (o *Avatar) SetPerformance(v AvatarPerformance)`

SetPerformance sets Performance field to given value.


### GetProductId

`func (o *Avatar) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *Avatar) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *Avatar) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *Avatar) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetPublishedListings

`func (o *Avatar) GetPublishedListings() []AvatarPublishedListingsInner`

GetPublishedListings returns the PublishedListings field if non-nil, zero value otherwise.

### GetPublishedListingsOk

`func (o *Avatar) GetPublishedListingsOk() (*[]AvatarPublishedListingsInner, bool)`

GetPublishedListingsOk returns a tuple with the PublishedListings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedListings

`func (o *Avatar) SetPublishedListings(v []AvatarPublishedListingsInner)`

SetPublishedListings sets PublishedListings field to given value.

### HasPublishedListings

`func (o *Avatar) HasPublishedListings() bool`

HasPublishedListings returns a boolean if a field has been set.

### GetReleaseStatus

`func (o *Avatar) GetReleaseStatus() ReleaseStatus`

GetReleaseStatus returns the ReleaseStatus field if non-nil, zero value otherwise.

### GetReleaseStatusOk

`func (o *Avatar) GetReleaseStatusOk() (*ReleaseStatus, bool)`

GetReleaseStatusOk returns a tuple with the ReleaseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseStatus

`func (o *Avatar) SetReleaseStatus(v ReleaseStatus)`

SetReleaseStatus sets ReleaseStatus field to given value.


### GetSearchable

`func (o *Avatar) GetSearchable() bool`

GetSearchable returns the Searchable field if non-nil, zero value otherwise.

### GetSearchableOk

`func (o *Avatar) GetSearchableOk() (*bool, bool)`

GetSearchableOk returns a tuple with the Searchable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchable

`func (o *Avatar) SetSearchable(v bool)`

SetSearchable sets Searchable field to given value.

### HasSearchable

`func (o *Avatar) HasSearchable() bool`

HasSearchable returns a boolean if a field has been set.

### GetStyles

`func (o *Avatar) GetStyles() AvatarStyles`

GetStyles returns the Styles field if non-nil, zero value otherwise.

### GetStylesOk

`func (o *Avatar) GetStylesOk() (*AvatarStyles, bool)`

GetStylesOk returns a tuple with the Styles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyles

`func (o *Avatar) SetStyles(v AvatarStyles)`

SetStyles sets Styles field to given value.


### GetTags

`func (o *Avatar) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Avatar) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Avatar) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetThumbnailImageUrl

`func (o *Avatar) GetThumbnailImageUrl() string`

GetThumbnailImageUrl returns the ThumbnailImageUrl field if non-nil, zero value otherwise.

### GetThumbnailImageUrlOk

`func (o *Avatar) GetThumbnailImageUrlOk() (*string, bool)`

GetThumbnailImageUrlOk returns a tuple with the ThumbnailImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailImageUrl

`func (o *Avatar) SetThumbnailImageUrl(v string)`

SetThumbnailImageUrl sets ThumbnailImageUrl field to given value.


### GetUnityPackageUrl

`func (o *Avatar) GetUnityPackageUrl() string`

GetUnityPackageUrl returns the UnityPackageUrl field if non-nil, zero value otherwise.

### GetUnityPackageUrlOk

`func (o *Avatar) GetUnityPackageUrlOk() (*string, bool)`

GetUnityPackageUrlOk returns a tuple with the UnityPackageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnityPackageUrl

`func (o *Avatar) SetUnityPackageUrl(v string)`

SetUnityPackageUrl sets UnityPackageUrl field to given value.


### GetUnityPackageUrlObject

`func (o *Avatar) GetUnityPackageUrlObject() AvatarUnityPackageUrlObject`

GetUnityPackageUrlObject returns the UnityPackageUrlObject field if non-nil, zero value otherwise.

### GetUnityPackageUrlObjectOk

`func (o *Avatar) GetUnityPackageUrlObjectOk() (*AvatarUnityPackageUrlObject, bool)`

GetUnityPackageUrlObjectOk returns a tuple with the UnityPackageUrlObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnityPackageUrlObject

`func (o *Avatar) SetUnityPackageUrlObject(v AvatarUnityPackageUrlObject)`

SetUnityPackageUrlObject sets UnityPackageUrlObject field to given value.


### GetUnityPackages

`func (o *Avatar) GetUnityPackages() []UnityPackage`

GetUnityPackages returns the UnityPackages field if non-nil, zero value otherwise.

### GetUnityPackagesOk

`func (o *Avatar) GetUnityPackagesOk() (*[]UnityPackage, bool)`

GetUnityPackagesOk returns a tuple with the UnityPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnityPackages

`func (o *Avatar) SetUnityPackages(v []UnityPackage)`

SetUnityPackages sets UnityPackages field to given value.


### GetUpdatedAt

`func (o *Avatar) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Avatar) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Avatar) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVersion

`func (o *Avatar) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Avatar) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Avatar) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


