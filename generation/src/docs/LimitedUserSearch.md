# LimitedUserSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bio** | Pointer to **string** |  | [optional] 
**BioLinks** | Pointer to **[]string** |   | [optional] 
**CurrentAvatarImageUrl** | **string** | When profilePicOverride is not empty, use it instead. | 
**CurrentAvatarTags** | **[]string** |  | 
**CurrentAvatarThumbnailImageUrl** | **string** | When profilePicOverride is not empty, use it instead. | 
**DeveloperType** | [**DeveloperType**](DeveloperType.md) |  | [default to NONE]
**DisplayName** | **string** |  | 
**Id** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**IsFriend** | **bool** |  | 
**LastPlatform** | **string** | This is normally &#x60;android&#x60;, &#x60;ios&#x60;, &#x60;standalonewindows&#x60;, &#x60;web&#x60;, or the empty value &#x60;&#x60;, but also supposedly can be any random Unity version such as &#x60;2019.2.4-801-Release&#x60; or &#x60;2019.2.2-772-Release&#x60; or even &#x60;unknownplatform&#x60;. | 
**ProfilePicOverride** | Pointer to **string** |  | [optional] 
**Pronouns** | Pointer to **string** |  | [optional] 
**Status** | [**UserStatus**](UserStatus.md) |  | [default to OFFLINE]
**StatusDescription** | **string** |  | 
**Tags** | **[]string** | &lt;- Always empty. | 
**UserIcon** | Pointer to **string** |  | [optional] 

## Methods

### NewLimitedUserSearch

`func NewLimitedUserSearch(currentAvatarImageUrl string, currentAvatarTags []string, currentAvatarThumbnailImageUrl string, developerType DeveloperType, displayName string, id string, isFriend bool, lastPlatform string, status UserStatus, statusDescription string, tags []string, ) *LimitedUserSearch`

NewLimitedUserSearch instantiates a new LimitedUserSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLimitedUserSearchWithDefaults

`func NewLimitedUserSearchWithDefaults() *LimitedUserSearch`

NewLimitedUserSearchWithDefaults instantiates a new LimitedUserSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBio

`func (o *LimitedUserSearch) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *LimitedUserSearch) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *LimitedUserSearch) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *LimitedUserSearch) HasBio() bool`

HasBio returns a boolean if a field has been set.

### GetBioLinks

`func (o *LimitedUserSearch) GetBioLinks() []string`

GetBioLinks returns the BioLinks field if non-nil, zero value otherwise.

### GetBioLinksOk

`func (o *LimitedUserSearch) GetBioLinksOk() (*[]string, bool)`

GetBioLinksOk returns a tuple with the BioLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBioLinks

`func (o *LimitedUserSearch) SetBioLinks(v []string)`

SetBioLinks sets BioLinks field to given value.

### HasBioLinks

`func (o *LimitedUserSearch) HasBioLinks() bool`

HasBioLinks returns a boolean if a field has been set.

### GetCurrentAvatarImageUrl

`func (o *LimitedUserSearch) GetCurrentAvatarImageUrl() string`

GetCurrentAvatarImageUrl returns the CurrentAvatarImageUrl field if non-nil, zero value otherwise.

### GetCurrentAvatarImageUrlOk

`func (o *LimitedUserSearch) GetCurrentAvatarImageUrlOk() (*string, bool)`

GetCurrentAvatarImageUrlOk returns a tuple with the CurrentAvatarImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAvatarImageUrl

`func (o *LimitedUserSearch) SetCurrentAvatarImageUrl(v string)`

SetCurrentAvatarImageUrl sets CurrentAvatarImageUrl field to given value.


### GetCurrentAvatarTags

`func (o *LimitedUserSearch) GetCurrentAvatarTags() []string`

GetCurrentAvatarTags returns the CurrentAvatarTags field if non-nil, zero value otherwise.

### GetCurrentAvatarTagsOk

`func (o *LimitedUserSearch) GetCurrentAvatarTagsOk() (*[]string, bool)`

GetCurrentAvatarTagsOk returns a tuple with the CurrentAvatarTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAvatarTags

`func (o *LimitedUserSearch) SetCurrentAvatarTags(v []string)`

SetCurrentAvatarTags sets CurrentAvatarTags field to given value.


### GetCurrentAvatarThumbnailImageUrl

`func (o *LimitedUserSearch) GetCurrentAvatarThumbnailImageUrl() string`

GetCurrentAvatarThumbnailImageUrl returns the CurrentAvatarThumbnailImageUrl field if non-nil, zero value otherwise.

### GetCurrentAvatarThumbnailImageUrlOk

`func (o *LimitedUserSearch) GetCurrentAvatarThumbnailImageUrlOk() (*string, bool)`

GetCurrentAvatarThumbnailImageUrlOk returns a tuple with the CurrentAvatarThumbnailImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAvatarThumbnailImageUrl

`func (o *LimitedUserSearch) SetCurrentAvatarThumbnailImageUrl(v string)`

SetCurrentAvatarThumbnailImageUrl sets CurrentAvatarThumbnailImageUrl field to given value.


### GetDeveloperType

`func (o *LimitedUserSearch) GetDeveloperType() DeveloperType`

GetDeveloperType returns the DeveloperType field if non-nil, zero value otherwise.

### GetDeveloperTypeOk

`func (o *LimitedUserSearch) GetDeveloperTypeOk() (*DeveloperType, bool)`

GetDeveloperTypeOk returns a tuple with the DeveloperType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperType

`func (o *LimitedUserSearch) SetDeveloperType(v DeveloperType)`

SetDeveloperType sets DeveloperType field to given value.


### GetDisplayName

`func (o *LimitedUserSearch) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *LimitedUserSearch) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *LimitedUserSearch) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetId

`func (o *LimitedUserSearch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LimitedUserSearch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LimitedUserSearch) SetId(v string)`

SetId sets Id field to given value.


### GetIsFriend

`func (o *LimitedUserSearch) GetIsFriend() bool`

GetIsFriend returns the IsFriend field if non-nil, zero value otherwise.

### GetIsFriendOk

`func (o *LimitedUserSearch) GetIsFriendOk() (*bool, bool)`

GetIsFriendOk returns a tuple with the IsFriend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFriend

`func (o *LimitedUserSearch) SetIsFriend(v bool)`

SetIsFriend sets IsFriend field to given value.


### GetLastPlatform

`func (o *LimitedUserSearch) GetLastPlatform() string`

GetLastPlatform returns the LastPlatform field if non-nil, zero value otherwise.

### GetLastPlatformOk

`func (o *LimitedUserSearch) GetLastPlatformOk() (*string, bool)`

GetLastPlatformOk returns a tuple with the LastPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPlatform

`func (o *LimitedUserSearch) SetLastPlatform(v string)`

SetLastPlatform sets LastPlatform field to given value.


### GetProfilePicOverride

`func (o *LimitedUserSearch) GetProfilePicOverride() string`

GetProfilePicOverride returns the ProfilePicOverride field if non-nil, zero value otherwise.

### GetProfilePicOverrideOk

`func (o *LimitedUserSearch) GetProfilePicOverrideOk() (*string, bool)`

GetProfilePicOverrideOk returns a tuple with the ProfilePicOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePicOverride

`func (o *LimitedUserSearch) SetProfilePicOverride(v string)`

SetProfilePicOverride sets ProfilePicOverride field to given value.

### HasProfilePicOverride

`func (o *LimitedUserSearch) HasProfilePicOverride() bool`

HasProfilePicOverride returns a boolean if a field has been set.

### GetPronouns

`func (o *LimitedUserSearch) GetPronouns() string`

GetPronouns returns the Pronouns field if non-nil, zero value otherwise.

### GetPronounsOk

`func (o *LimitedUserSearch) GetPronounsOk() (*string, bool)`

GetPronounsOk returns a tuple with the Pronouns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPronouns

`func (o *LimitedUserSearch) SetPronouns(v string)`

SetPronouns sets Pronouns field to given value.

### HasPronouns

`func (o *LimitedUserSearch) HasPronouns() bool`

HasPronouns returns a boolean if a field has been set.

### GetStatus

`func (o *LimitedUserSearch) GetStatus() UserStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LimitedUserSearch) GetStatusOk() (*UserStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LimitedUserSearch) SetStatus(v UserStatus)`

SetStatus sets Status field to given value.


### GetStatusDescription

`func (o *LimitedUserSearch) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *LimitedUserSearch) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *LimitedUserSearch) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.


### GetTags

`func (o *LimitedUserSearch) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LimitedUserSearch) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LimitedUserSearch) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetUserIcon

`func (o *LimitedUserSearch) GetUserIcon() string`

GetUserIcon returns the UserIcon field if non-nil, zero value otherwise.

### GetUserIconOk

`func (o *LimitedUserSearch) GetUserIconOk() (*string, bool)`

GetUserIconOk returns a tuple with the UserIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIcon

`func (o *LimitedUserSearch) SetUserIcon(v string)`

SetUserIcon sets UserIcon field to given value.

### HasUserIcon

`func (o *LimitedUserSearch) HasUserIcon() bool`

HasUserIcon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


