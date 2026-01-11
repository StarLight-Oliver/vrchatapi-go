# LimitedWorld

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**AuthorName** | **string** |  | 
**Capacity** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**DefaultContentSettings** | Pointer to [**InstanceContentSettings**](InstanceContentSettings.md) |  | [optional] 
**Favorites** | **int32** |  | [default to 0]
**Heat** | **int32** |  | [default to 0]
**Id** | **string** | WorldID be \&quot;offline\&quot; on User profiles if you are not friends with that user. | 
**ImageUrl** | **string** |  | 
**LabsPublicationDate** | **string** |  | 
**Name** | **string** |  | 
**Occupants** | **int32** |  | [default to 0]
**Organization** | **string** |  | [default to "vrchat"]
**Popularity** | **int32** |  | [default to 0]
**PreviewYoutubeId** | Pointer to **NullableString** |  | [optional] 
**PublicationDate** | **string** |  | 
**RecommendedCapacity** | Pointer to **int32** |  | [optional] 
**ReleaseStatus** | [**ReleaseStatus**](ReleaseStatus.md) |  | [default to PUBLIC]
**StoreId** | Pointer to **string** |  | [optional] 
**Tags** | **[]string** |   | 
**ThumbnailImageUrl** | **string** |  | 
**UdonProducts** | Pointer to **[]string** |  | [optional] 
**UnityPackages** | [**[]LimitedUnityPackage**](LimitedUnityPackage.md) |   | 
**UpdatedAt** | **time.Time** |  | 
**Visits** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewLimitedWorld

`func NewLimitedWorld(authorId string, authorName string, capacity int32, createdAt time.Time, favorites int32, heat int32, id string, imageUrl string, labsPublicationDate string, name string, occupants int32, organization string, popularity int32, publicationDate string, releaseStatus ReleaseStatus, tags []string, thumbnailImageUrl string, unityPackages []LimitedUnityPackage, updatedAt time.Time, ) *LimitedWorld`

NewLimitedWorld instantiates a new LimitedWorld object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLimitedWorldWithDefaults

`func NewLimitedWorldWithDefaults() *LimitedWorld`

NewLimitedWorldWithDefaults instantiates a new LimitedWorld object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorId

`func (o *LimitedWorld) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *LimitedWorld) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *LimitedWorld) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.


### GetAuthorName

`func (o *LimitedWorld) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *LimitedWorld) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *LimitedWorld) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.


### GetCapacity

`func (o *LimitedWorld) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *LimitedWorld) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *LimitedWorld) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.


### GetCreatedAt

`func (o *LimitedWorld) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LimitedWorld) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LimitedWorld) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDefaultContentSettings

`func (o *LimitedWorld) GetDefaultContentSettings() InstanceContentSettings`

GetDefaultContentSettings returns the DefaultContentSettings field if non-nil, zero value otherwise.

### GetDefaultContentSettingsOk

`func (o *LimitedWorld) GetDefaultContentSettingsOk() (*InstanceContentSettings, bool)`

GetDefaultContentSettingsOk returns a tuple with the DefaultContentSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultContentSettings

`func (o *LimitedWorld) SetDefaultContentSettings(v InstanceContentSettings)`

SetDefaultContentSettings sets DefaultContentSettings field to given value.

### HasDefaultContentSettings

`func (o *LimitedWorld) HasDefaultContentSettings() bool`

HasDefaultContentSettings returns a boolean if a field has been set.

### GetFavorites

`func (o *LimitedWorld) GetFavorites() int32`

GetFavorites returns the Favorites field if non-nil, zero value otherwise.

### GetFavoritesOk

`func (o *LimitedWorld) GetFavoritesOk() (*int32, bool)`

GetFavoritesOk returns a tuple with the Favorites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorites

`func (o *LimitedWorld) SetFavorites(v int32)`

SetFavorites sets Favorites field to given value.


### GetHeat

`func (o *LimitedWorld) GetHeat() int32`

GetHeat returns the Heat field if non-nil, zero value otherwise.

### GetHeatOk

`func (o *LimitedWorld) GetHeatOk() (*int32, bool)`

GetHeatOk returns a tuple with the Heat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeat

`func (o *LimitedWorld) SetHeat(v int32)`

SetHeat sets Heat field to given value.


### GetId

`func (o *LimitedWorld) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LimitedWorld) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LimitedWorld) SetId(v string)`

SetId sets Id field to given value.


### GetImageUrl

`func (o *LimitedWorld) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *LimitedWorld) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *LimitedWorld) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.


### GetLabsPublicationDate

`func (o *LimitedWorld) GetLabsPublicationDate() string`

GetLabsPublicationDate returns the LabsPublicationDate field if non-nil, zero value otherwise.

### GetLabsPublicationDateOk

`func (o *LimitedWorld) GetLabsPublicationDateOk() (*string, bool)`

GetLabsPublicationDateOk returns a tuple with the LabsPublicationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabsPublicationDate

`func (o *LimitedWorld) SetLabsPublicationDate(v string)`

SetLabsPublicationDate sets LabsPublicationDate field to given value.


### GetName

`func (o *LimitedWorld) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LimitedWorld) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LimitedWorld) SetName(v string)`

SetName sets Name field to given value.


### GetOccupants

`func (o *LimitedWorld) GetOccupants() int32`

GetOccupants returns the Occupants field if non-nil, zero value otherwise.

### GetOccupantsOk

`func (o *LimitedWorld) GetOccupantsOk() (*int32, bool)`

GetOccupantsOk returns a tuple with the Occupants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupants

`func (o *LimitedWorld) SetOccupants(v int32)`

SetOccupants sets Occupants field to given value.


### GetOrganization

`func (o *LimitedWorld) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *LimitedWorld) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *LimitedWorld) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetPopularity

`func (o *LimitedWorld) GetPopularity() int32`

GetPopularity returns the Popularity field if non-nil, zero value otherwise.

### GetPopularityOk

`func (o *LimitedWorld) GetPopularityOk() (*int32, bool)`

GetPopularityOk returns a tuple with the Popularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopularity

`func (o *LimitedWorld) SetPopularity(v int32)`

SetPopularity sets Popularity field to given value.


### GetPreviewYoutubeId

`func (o *LimitedWorld) GetPreviewYoutubeId() string`

GetPreviewYoutubeId returns the PreviewYoutubeId field if non-nil, zero value otherwise.

### GetPreviewYoutubeIdOk

`func (o *LimitedWorld) GetPreviewYoutubeIdOk() (*string, bool)`

GetPreviewYoutubeIdOk returns a tuple with the PreviewYoutubeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewYoutubeId

`func (o *LimitedWorld) SetPreviewYoutubeId(v string)`

SetPreviewYoutubeId sets PreviewYoutubeId field to given value.

### HasPreviewYoutubeId

`func (o *LimitedWorld) HasPreviewYoutubeId() bool`

HasPreviewYoutubeId returns a boolean if a field has been set.

### SetPreviewYoutubeIdNil

`func (o *LimitedWorld) SetPreviewYoutubeIdNil(b bool)`

 SetPreviewYoutubeIdNil sets the value for PreviewYoutubeId to be an explicit nil

### UnsetPreviewYoutubeId
`func (o *LimitedWorld) UnsetPreviewYoutubeId()`

UnsetPreviewYoutubeId ensures that no value is present for PreviewYoutubeId, not even an explicit nil
### GetPublicationDate

`func (o *LimitedWorld) GetPublicationDate() string`

GetPublicationDate returns the PublicationDate field if non-nil, zero value otherwise.

### GetPublicationDateOk

`func (o *LimitedWorld) GetPublicationDateOk() (*string, bool)`

GetPublicationDateOk returns a tuple with the PublicationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationDate

`func (o *LimitedWorld) SetPublicationDate(v string)`

SetPublicationDate sets PublicationDate field to given value.


### GetRecommendedCapacity

`func (o *LimitedWorld) GetRecommendedCapacity() int32`

GetRecommendedCapacity returns the RecommendedCapacity field if non-nil, zero value otherwise.

### GetRecommendedCapacityOk

`func (o *LimitedWorld) GetRecommendedCapacityOk() (*int32, bool)`

GetRecommendedCapacityOk returns a tuple with the RecommendedCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedCapacity

`func (o *LimitedWorld) SetRecommendedCapacity(v int32)`

SetRecommendedCapacity sets RecommendedCapacity field to given value.

### HasRecommendedCapacity

`func (o *LimitedWorld) HasRecommendedCapacity() bool`

HasRecommendedCapacity returns a boolean if a field has been set.

### GetReleaseStatus

`func (o *LimitedWorld) GetReleaseStatus() ReleaseStatus`

GetReleaseStatus returns the ReleaseStatus field if non-nil, zero value otherwise.

### GetReleaseStatusOk

`func (o *LimitedWorld) GetReleaseStatusOk() (*ReleaseStatus, bool)`

GetReleaseStatusOk returns a tuple with the ReleaseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseStatus

`func (o *LimitedWorld) SetReleaseStatus(v ReleaseStatus)`

SetReleaseStatus sets ReleaseStatus field to given value.


### GetStoreId

`func (o *LimitedWorld) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *LimitedWorld) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *LimitedWorld) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *LimitedWorld) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetTags

`func (o *LimitedWorld) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LimitedWorld) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LimitedWorld) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetThumbnailImageUrl

`func (o *LimitedWorld) GetThumbnailImageUrl() string`

GetThumbnailImageUrl returns the ThumbnailImageUrl field if non-nil, zero value otherwise.

### GetThumbnailImageUrlOk

`func (o *LimitedWorld) GetThumbnailImageUrlOk() (*string, bool)`

GetThumbnailImageUrlOk returns a tuple with the ThumbnailImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailImageUrl

`func (o *LimitedWorld) SetThumbnailImageUrl(v string)`

SetThumbnailImageUrl sets ThumbnailImageUrl field to given value.


### GetUdonProducts

`func (o *LimitedWorld) GetUdonProducts() []string`

GetUdonProducts returns the UdonProducts field if non-nil, zero value otherwise.

### GetUdonProductsOk

`func (o *LimitedWorld) GetUdonProductsOk() (*[]string, bool)`

GetUdonProductsOk returns a tuple with the UdonProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdonProducts

`func (o *LimitedWorld) SetUdonProducts(v []string)`

SetUdonProducts sets UdonProducts field to given value.

### HasUdonProducts

`func (o *LimitedWorld) HasUdonProducts() bool`

HasUdonProducts returns a boolean if a field has been set.

### GetUnityPackages

`func (o *LimitedWorld) GetUnityPackages() []LimitedUnityPackage`

GetUnityPackages returns the UnityPackages field if non-nil, zero value otherwise.

### GetUnityPackagesOk

`func (o *LimitedWorld) GetUnityPackagesOk() (*[]LimitedUnityPackage, bool)`

GetUnityPackagesOk returns a tuple with the UnityPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnityPackages

`func (o *LimitedWorld) SetUnityPackages(v []LimitedUnityPackage)`

SetUnityPackages sets UnityPackages field to given value.


### GetUpdatedAt

`func (o *LimitedWorld) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LimitedWorld) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LimitedWorld) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVisits

`func (o *LimitedWorld) GetVisits() int32`

GetVisits returns the Visits field if non-nil, zero value otherwise.

### GetVisitsOk

`func (o *LimitedWorld) GetVisitsOk() (*int32, bool)`

GetVisitsOk returns a tuple with the Visits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisits

`func (o *LimitedWorld) SetVisits(v int32)`

SetVisits sets Visits field to given value.

### HasVisits

`func (o *LimitedWorld) HasVisits() bool`

HasVisits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


