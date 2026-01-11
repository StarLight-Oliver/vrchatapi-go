# FavoritedWorld

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorId** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 
**AuthorName** | **string** |  | 
**Capacity** | **int32** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DefaultContentSettings** | Pointer to [**InstanceContentSettings**](InstanceContentSettings.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FavoriteGroup** | **string** |  | 
**FavoriteId** | **string** |  | 
**Favorites** | Pointer to **int32** |  | [optional] [default to 0]
**Featured** | Pointer to **bool** |  | [optional] [default to false]
**Heat** | Pointer to **int32** |  | [optional] [default to 0]
**Id** | **string** | WorldID be \&quot;offline\&quot; on User profiles if you are not friends with that user. | 
**ImageUrl** | **string** |  | 
**LabsPublicationDate** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Occupants** | **int32** |  | [default to 0]
**Organization** | Pointer to **string** |  | [optional] [default to "vrchat"]
**Popularity** | Pointer to **int32** |  | [optional] [default to 0]
**PreviewYoutubeId** | Pointer to **NullableString** |  | [optional] 
**PublicationDate** | Pointer to **string** |  | [optional] 
**RecommendedCapacity** | Pointer to **int32** |  | [optional] 
**ReleaseStatus** | [**ReleaseStatus**](ReleaseStatus.md) |  | [default to PUBLIC]
**Tags** | Pointer to **[]string** |   | [optional] 
**ThumbnailImageUrl** | **string** |  | 
**UdonProducts** | Pointer to **[]string** |  | [optional] 
**UnityPackages** | Pointer to [**[]UnityPackage**](UnityPackage.md) |   | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UrlList** | Pointer to **[]string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**Visits** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewFavoritedWorld

`func NewFavoritedWorld(authorName string, capacity int32, favoriteGroup string, favoriteId string, id string, imageUrl string, name string, occupants int32, releaseStatus ReleaseStatus, thumbnailImageUrl string, ) *FavoritedWorld`

NewFavoritedWorld instantiates a new FavoritedWorld object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFavoritedWorldWithDefaults

`func NewFavoritedWorldWithDefaults() *FavoritedWorld`

NewFavoritedWorldWithDefaults instantiates a new FavoritedWorld object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorId

`func (o *FavoritedWorld) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *FavoritedWorld) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *FavoritedWorld) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *FavoritedWorld) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetAuthorName

`func (o *FavoritedWorld) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *FavoritedWorld) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *FavoritedWorld) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.


### GetCapacity

`func (o *FavoritedWorld) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *FavoritedWorld) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *FavoritedWorld) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.


### GetCreatedAt

`func (o *FavoritedWorld) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FavoritedWorld) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FavoritedWorld) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FavoritedWorld) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefaultContentSettings

`func (o *FavoritedWorld) GetDefaultContentSettings() InstanceContentSettings`

GetDefaultContentSettings returns the DefaultContentSettings field if non-nil, zero value otherwise.

### GetDefaultContentSettingsOk

`func (o *FavoritedWorld) GetDefaultContentSettingsOk() (*InstanceContentSettings, bool)`

GetDefaultContentSettingsOk returns a tuple with the DefaultContentSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultContentSettings

`func (o *FavoritedWorld) SetDefaultContentSettings(v InstanceContentSettings)`

SetDefaultContentSettings sets DefaultContentSettings field to given value.

### HasDefaultContentSettings

`func (o *FavoritedWorld) HasDefaultContentSettings() bool`

HasDefaultContentSettings returns a boolean if a field has been set.

### GetDescription

`func (o *FavoritedWorld) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FavoritedWorld) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FavoritedWorld) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FavoritedWorld) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFavoriteGroup

`func (o *FavoritedWorld) GetFavoriteGroup() string`

GetFavoriteGroup returns the FavoriteGroup field if non-nil, zero value otherwise.

### GetFavoriteGroupOk

`func (o *FavoritedWorld) GetFavoriteGroupOk() (*string, bool)`

GetFavoriteGroupOk returns a tuple with the FavoriteGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavoriteGroup

`func (o *FavoritedWorld) SetFavoriteGroup(v string)`

SetFavoriteGroup sets FavoriteGroup field to given value.


### GetFavoriteId

`func (o *FavoritedWorld) GetFavoriteId() string`

GetFavoriteId returns the FavoriteId field if non-nil, zero value otherwise.

### GetFavoriteIdOk

`func (o *FavoritedWorld) GetFavoriteIdOk() (*string, bool)`

GetFavoriteIdOk returns a tuple with the FavoriteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavoriteId

`func (o *FavoritedWorld) SetFavoriteId(v string)`

SetFavoriteId sets FavoriteId field to given value.


### GetFavorites

`func (o *FavoritedWorld) GetFavorites() int32`

GetFavorites returns the Favorites field if non-nil, zero value otherwise.

### GetFavoritesOk

`func (o *FavoritedWorld) GetFavoritesOk() (*int32, bool)`

GetFavoritesOk returns a tuple with the Favorites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorites

`func (o *FavoritedWorld) SetFavorites(v int32)`

SetFavorites sets Favorites field to given value.

### HasFavorites

`func (o *FavoritedWorld) HasFavorites() bool`

HasFavorites returns a boolean if a field has been set.

### GetFeatured

`func (o *FavoritedWorld) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *FavoritedWorld) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *FavoritedWorld) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *FavoritedWorld) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetHeat

`func (o *FavoritedWorld) GetHeat() int32`

GetHeat returns the Heat field if non-nil, zero value otherwise.

### GetHeatOk

`func (o *FavoritedWorld) GetHeatOk() (*int32, bool)`

GetHeatOk returns a tuple with the Heat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeat

`func (o *FavoritedWorld) SetHeat(v int32)`

SetHeat sets Heat field to given value.

### HasHeat

`func (o *FavoritedWorld) HasHeat() bool`

HasHeat returns a boolean if a field has been set.

### GetId

`func (o *FavoritedWorld) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FavoritedWorld) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FavoritedWorld) SetId(v string)`

SetId sets Id field to given value.


### GetImageUrl

`func (o *FavoritedWorld) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *FavoritedWorld) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *FavoritedWorld) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.


### GetLabsPublicationDate

`func (o *FavoritedWorld) GetLabsPublicationDate() string`

GetLabsPublicationDate returns the LabsPublicationDate field if non-nil, zero value otherwise.

### GetLabsPublicationDateOk

`func (o *FavoritedWorld) GetLabsPublicationDateOk() (*string, bool)`

GetLabsPublicationDateOk returns a tuple with the LabsPublicationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabsPublicationDate

`func (o *FavoritedWorld) SetLabsPublicationDate(v string)`

SetLabsPublicationDate sets LabsPublicationDate field to given value.

### HasLabsPublicationDate

`func (o *FavoritedWorld) HasLabsPublicationDate() bool`

HasLabsPublicationDate returns a boolean if a field has been set.

### GetName

`func (o *FavoritedWorld) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FavoritedWorld) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FavoritedWorld) SetName(v string)`

SetName sets Name field to given value.


### GetOccupants

`func (o *FavoritedWorld) GetOccupants() int32`

GetOccupants returns the Occupants field if non-nil, zero value otherwise.

### GetOccupantsOk

`func (o *FavoritedWorld) GetOccupantsOk() (*int32, bool)`

GetOccupantsOk returns a tuple with the Occupants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupants

`func (o *FavoritedWorld) SetOccupants(v int32)`

SetOccupants sets Occupants field to given value.


### GetOrganization

`func (o *FavoritedWorld) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FavoritedWorld) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FavoritedWorld) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FavoritedWorld) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetPopularity

`func (o *FavoritedWorld) GetPopularity() int32`

GetPopularity returns the Popularity field if non-nil, zero value otherwise.

### GetPopularityOk

`func (o *FavoritedWorld) GetPopularityOk() (*int32, bool)`

GetPopularityOk returns a tuple with the Popularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopularity

`func (o *FavoritedWorld) SetPopularity(v int32)`

SetPopularity sets Popularity field to given value.

### HasPopularity

`func (o *FavoritedWorld) HasPopularity() bool`

HasPopularity returns a boolean if a field has been set.

### GetPreviewYoutubeId

`func (o *FavoritedWorld) GetPreviewYoutubeId() string`

GetPreviewYoutubeId returns the PreviewYoutubeId field if non-nil, zero value otherwise.

### GetPreviewYoutubeIdOk

`func (o *FavoritedWorld) GetPreviewYoutubeIdOk() (*string, bool)`

GetPreviewYoutubeIdOk returns a tuple with the PreviewYoutubeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewYoutubeId

`func (o *FavoritedWorld) SetPreviewYoutubeId(v string)`

SetPreviewYoutubeId sets PreviewYoutubeId field to given value.

### HasPreviewYoutubeId

`func (o *FavoritedWorld) HasPreviewYoutubeId() bool`

HasPreviewYoutubeId returns a boolean if a field has been set.

### SetPreviewYoutubeIdNil

`func (o *FavoritedWorld) SetPreviewYoutubeIdNil(b bool)`

 SetPreviewYoutubeIdNil sets the value for PreviewYoutubeId to be an explicit nil

### UnsetPreviewYoutubeId
`func (o *FavoritedWorld) UnsetPreviewYoutubeId()`

UnsetPreviewYoutubeId ensures that no value is present for PreviewYoutubeId, not even an explicit nil
### GetPublicationDate

`func (o *FavoritedWorld) GetPublicationDate() string`

GetPublicationDate returns the PublicationDate field if non-nil, zero value otherwise.

### GetPublicationDateOk

`func (o *FavoritedWorld) GetPublicationDateOk() (*string, bool)`

GetPublicationDateOk returns a tuple with the PublicationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationDate

`func (o *FavoritedWorld) SetPublicationDate(v string)`

SetPublicationDate sets PublicationDate field to given value.

### HasPublicationDate

`func (o *FavoritedWorld) HasPublicationDate() bool`

HasPublicationDate returns a boolean if a field has been set.

### GetRecommendedCapacity

`func (o *FavoritedWorld) GetRecommendedCapacity() int32`

GetRecommendedCapacity returns the RecommendedCapacity field if non-nil, zero value otherwise.

### GetRecommendedCapacityOk

`func (o *FavoritedWorld) GetRecommendedCapacityOk() (*int32, bool)`

GetRecommendedCapacityOk returns a tuple with the RecommendedCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedCapacity

`func (o *FavoritedWorld) SetRecommendedCapacity(v int32)`

SetRecommendedCapacity sets RecommendedCapacity field to given value.

### HasRecommendedCapacity

`func (o *FavoritedWorld) HasRecommendedCapacity() bool`

HasRecommendedCapacity returns a boolean if a field has been set.

### GetReleaseStatus

`func (o *FavoritedWorld) GetReleaseStatus() ReleaseStatus`

GetReleaseStatus returns the ReleaseStatus field if non-nil, zero value otherwise.

### GetReleaseStatusOk

`func (o *FavoritedWorld) GetReleaseStatusOk() (*ReleaseStatus, bool)`

GetReleaseStatusOk returns a tuple with the ReleaseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseStatus

`func (o *FavoritedWorld) SetReleaseStatus(v ReleaseStatus)`

SetReleaseStatus sets ReleaseStatus field to given value.


### GetTags

`func (o *FavoritedWorld) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FavoritedWorld) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FavoritedWorld) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FavoritedWorld) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetThumbnailImageUrl

`func (o *FavoritedWorld) GetThumbnailImageUrl() string`

GetThumbnailImageUrl returns the ThumbnailImageUrl field if non-nil, zero value otherwise.

### GetThumbnailImageUrlOk

`func (o *FavoritedWorld) GetThumbnailImageUrlOk() (*string, bool)`

GetThumbnailImageUrlOk returns a tuple with the ThumbnailImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailImageUrl

`func (o *FavoritedWorld) SetThumbnailImageUrl(v string)`

SetThumbnailImageUrl sets ThumbnailImageUrl field to given value.


### GetUdonProducts

`func (o *FavoritedWorld) GetUdonProducts() []string`

GetUdonProducts returns the UdonProducts field if non-nil, zero value otherwise.

### GetUdonProductsOk

`func (o *FavoritedWorld) GetUdonProductsOk() (*[]string, bool)`

GetUdonProductsOk returns a tuple with the UdonProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdonProducts

`func (o *FavoritedWorld) SetUdonProducts(v []string)`

SetUdonProducts sets UdonProducts field to given value.

### HasUdonProducts

`func (o *FavoritedWorld) HasUdonProducts() bool`

HasUdonProducts returns a boolean if a field has been set.

### GetUnityPackages

`func (o *FavoritedWorld) GetUnityPackages() []UnityPackage`

GetUnityPackages returns the UnityPackages field if non-nil, zero value otherwise.

### GetUnityPackagesOk

`func (o *FavoritedWorld) GetUnityPackagesOk() (*[]UnityPackage, bool)`

GetUnityPackagesOk returns a tuple with the UnityPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnityPackages

`func (o *FavoritedWorld) SetUnityPackages(v []UnityPackage)`

SetUnityPackages sets UnityPackages field to given value.

### HasUnityPackages

`func (o *FavoritedWorld) HasUnityPackages() bool`

HasUnityPackages returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FavoritedWorld) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FavoritedWorld) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FavoritedWorld) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FavoritedWorld) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrlList

`func (o *FavoritedWorld) GetUrlList() []string`

GetUrlList returns the UrlList field if non-nil, zero value otherwise.

### GetUrlListOk

`func (o *FavoritedWorld) GetUrlListOk() (*[]string, bool)`

GetUrlListOk returns a tuple with the UrlList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlList

`func (o *FavoritedWorld) SetUrlList(v []string)`

SetUrlList sets UrlList field to given value.

### HasUrlList

`func (o *FavoritedWorld) HasUrlList() bool`

HasUrlList returns a boolean if a field has been set.

### GetVersion

`func (o *FavoritedWorld) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FavoritedWorld) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FavoritedWorld) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FavoritedWorld) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVisits

`func (o *FavoritedWorld) GetVisits() int32`

GetVisits returns the Visits field if non-nil, zero value otherwise.

### GetVisitsOk

`func (o *FavoritedWorld) GetVisitsOk() (*int32, bool)`

GetVisitsOk returns a tuple with the Visits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisits

`func (o *FavoritedWorld) SetVisits(v int32)`

SetVisits sets Visits field to given value.

### HasVisits

`func (o *FavoritedWorld) HasVisits() bool`

HasVisits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


