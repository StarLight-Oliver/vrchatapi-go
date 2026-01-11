# LimitedUserFriend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bio** | Pointer to **string** |  | [optional] 
**BioLinks** | Pointer to **[]string** |   | [optional] 
**CurrentAvatarImageUrl** | Pointer to **string** | When profilePicOverride is not empty, use it instead. | [optional] 
**CurrentAvatarTags** | Pointer to **[]string** |  | [optional] 
**CurrentAvatarThumbnailImageUrl** | Pointer to **string** | When profilePicOverride is not empty, use it instead. | [optional] 
**DeveloperType** | [**DeveloperType**](DeveloperType.md) |  | [default to NONE]
**DisplayName** | **string** |  | 
**FriendKey** | **string** |  | 
**Id** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**ImageUrl** | **string** |  | 
**IsFriend** | **bool** |  | 
**LastActivity** | **NullableTime** |  | 
**LastLogin** | **NullableTime** |  | 
**LastMobile** | **NullableTime** |  | 
**LastPlatform** | **string** | This is normally &#x60;android&#x60;, &#x60;ios&#x60;, &#x60;standalonewindows&#x60;, &#x60;web&#x60;, or the empty value &#x60;&#x60;, but also supposedly can be any random Unity version such as &#x60;2019.2.4-801-Release&#x60; or &#x60;2019.2.2-772-Release&#x60; or even &#x60;unknownplatform&#x60;. | 
**Location** | **string** |  | 
**Platform** | **string** |  | 
**ProfilePicOverride** | Pointer to **string** |  | [optional] 
**ProfilePicOverrideThumbnail** | Pointer to **string** |  | [optional] 
**Status** | [**UserStatus**](UserStatus.md) |  | [default to OFFLINE]
**StatusDescription** | **string** |  | 
**Tags** | **[]string** | &lt;- Always empty. | 
**UserIcon** | Pointer to **string** |  | [optional] 

## Methods

### NewLimitedUserFriend

`func NewLimitedUserFriend(developerType DeveloperType, displayName string, friendKey string, id string, imageUrl string, isFriend bool, lastActivity NullableTime, lastLogin NullableTime, lastMobile NullableTime, lastPlatform string, location string, platform string, status UserStatus, statusDescription string, tags []string, ) *LimitedUserFriend`

NewLimitedUserFriend instantiates a new LimitedUserFriend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLimitedUserFriendWithDefaults

`func NewLimitedUserFriendWithDefaults() *LimitedUserFriend`

NewLimitedUserFriendWithDefaults instantiates a new LimitedUserFriend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBio

`func (o *LimitedUserFriend) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *LimitedUserFriend) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *LimitedUserFriend) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *LimitedUserFriend) HasBio() bool`

HasBio returns a boolean if a field has been set.

### GetBioLinks

`func (o *LimitedUserFriend) GetBioLinks() []string`

GetBioLinks returns the BioLinks field if non-nil, zero value otherwise.

### GetBioLinksOk

`func (o *LimitedUserFriend) GetBioLinksOk() (*[]string, bool)`

GetBioLinksOk returns a tuple with the BioLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBioLinks

`func (o *LimitedUserFriend) SetBioLinks(v []string)`

SetBioLinks sets BioLinks field to given value.

### HasBioLinks

`func (o *LimitedUserFriend) HasBioLinks() bool`

HasBioLinks returns a boolean if a field has been set.

### GetCurrentAvatarImageUrl

`func (o *LimitedUserFriend) GetCurrentAvatarImageUrl() string`

GetCurrentAvatarImageUrl returns the CurrentAvatarImageUrl field if non-nil, zero value otherwise.

### GetCurrentAvatarImageUrlOk

`func (o *LimitedUserFriend) GetCurrentAvatarImageUrlOk() (*string, bool)`

GetCurrentAvatarImageUrlOk returns a tuple with the CurrentAvatarImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAvatarImageUrl

`func (o *LimitedUserFriend) SetCurrentAvatarImageUrl(v string)`

SetCurrentAvatarImageUrl sets CurrentAvatarImageUrl field to given value.

### HasCurrentAvatarImageUrl

`func (o *LimitedUserFriend) HasCurrentAvatarImageUrl() bool`

HasCurrentAvatarImageUrl returns a boolean if a field has been set.

### GetCurrentAvatarTags

`func (o *LimitedUserFriend) GetCurrentAvatarTags() []string`

GetCurrentAvatarTags returns the CurrentAvatarTags field if non-nil, zero value otherwise.

### GetCurrentAvatarTagsOk

`func (o *LimitedUserFriend) GetCurrentAvatarTagsOk() (*[]string, bool)`

GetCurrentAvatarTagsOk returns a tuple with the CurrentAvatarTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAvatarTags

`func (o *LimitedUserFriend) SetCurrentAvatarTags(v []string)`

SetCurrentAvatarTags sets CurrentAvatarTags field to given value.

### HasCurrentAvatarTags

`func (o *LimitedUserFriend) HasCurrentAvatarTags() bool`

HasCurrentAvatarTags returns a boolean if a field has been set.

### GetCurrentAvatarThumbnailImageUrl

`func (o *LimitedUserFriend) GetCurrentAvatarThumbnailImageUrl() string`

GetCurrentAvatarThumbnailImageUrl returns the CurrentAvatarThumbnailImageUrl field if non-nil, zero value otherwise.

### GetCurrentAvatarThumbnailImageUrlOk

`func (o *LimitedUserFriend) GetCurrentAvatarThumbnailImageUrlOk() (*string, bool)`

GetCurrentAvatarThumbnailImageUrlOk returns a tuple with the CurrentAvatarThumbnailImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAvatarThumbnailImageUrl

`func (o *LimitedUserFriend) SetCurrentAvatarThumbnailImageUrl(v string)`

SetCurrentAvatarThumbnailImageUrl sets CurrentAvatarThumbnailImageUrl field to given value.

### HasCurrentAvatarThumbnailImageUrl

`func (o *LimitedUserFriend) HasCurrentAvatarThumbnailImageUrl() bool`

HasCurrentAvatarThumbnailImageUrl returns a boolean if a field has been set.

### GetDeveloperType

`func (o *LimitedUserFriend) GetDeveloperType() DeveloperType`

GetDeveloperType returns the DeveloperType field if non-nil, zero value otherwise.

### GetDeveloperTypeOk

`func (o *LimitedUserFriend) GetDeveloperTypeOk() (*DeveloperType, bool)`

GetDeveloperTypeOk returns a tuple with the DeveloperType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperType

`func (o *LimitedUserFriend) SetDeveloperType(v DeveloperType)`

SetDeveloperType sets DeveloperType field to given value.


### GetDisplayName

`func (o *LimitedUserFriend) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *LimitedUserFriend) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *LimitedUserFriend) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetFriendKey

`func (o *LimitedUserFriend) GetFriendKey() string`

GetFriendKey returns the FriendKey field if non-nil, zero value otherwise.

### GetFriendKeyOk

`func (o *LimitedUserFriend) GetFriendKeyOk() (*string, bool)`

GetFriendKeyOk returns a tuple with the FriendKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendKey

`func (o *LimitedUserFriend) SetFriendKey(v string)`

SetFriendKey sets FriendKey field to given value.


### GetId

`func (o *LimitedUserFriend) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LimitedUserFriend) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LimitedUserFriend) SetId(v string)`

SetId sets Id field to given value.


### GetImageUrl

`func (o *LimitedUserFriend) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *LimitedUserFriend) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *LimitedUserFriend) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.


### GetIsFriend

`func (o *LimitedUserFriend) GetIsFriend() bool`

GetIsFriend returns the IsFriend field if non-nil, zero value otherwise.

### GetIsFriendOk

`func (o *LimitedUserFriend) GetIsFriendOk() (*bool, bool)`

GetIsFriendOk returns a tuple with the IsFriend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFriend

`func (o *LimitedUserFriend) SetIsFriend(v bool)`

SetIsFriend sets IsFriend field to given value.


### GetLastActivity

`func (o *LimitedUserFriend) GetLastActivity() time.Time`

GetLastActivity returns the LastActivity field if non-nil, zero value otherwise.

### GetLastActivityOk

`func (o *LimitedUserFriend) GetLastActivityOk() (*time.Time, bool)`

GetLastActivityOk returns a tuple with the LastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivity

`func (o *LimitedUserFriend) SetLastActivity(v time.Time)`

SetLastActivity sets LastActivity field to given value.


### SetLastActivityNil

`func (o *LimitedUserFriend) SetLastActivityNil(b bool)`

 SetLastActivityNil sets the value for LastActivity to be an explicit nil

### UnsetLastActivity
`func (o *LimitedUserFriend) UnsetLastActivity()`

UnsetLastActivity ensures that no value is present for LastActivity, not even an explicit nil
### GetLastLogin

`func (o *LimitedUserFriend) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *LimitedUserFriend) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *LimitedUserFriend) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.


### SetLastLoginNil

`func (o *LimitedUserFriend) SetLastLoginNil(b bool)`

 SetLastLoginNil sets the value for LastLogin to be an explicit nil

### UnsetLastLogin
`func (o *LimitedUserFriend) UnsetLastLogin()`

UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
### GetLastMobile

`func (o *LimitedUserFriend) GetLastMobile() time.Time`

GetLastMobile returns the LastMobile field if non-nil, zero value otherwise.

### GetLastMobileOk

`func (o *LimitedUserFriend) GetLastMobileOk() (*time.Time, bool)`

GetLastMobileOk returns a tuple with the LastMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMobile

`func (o *LimitedUserFriend) SetLastMobile(v time.Time)`

SetLastMobile sets LastMobile field to given value.


### SetLastMobileNil

`func (o *LimitedUserFriend) SetLastMobileNil(b bool)`

 SetLastMobileNil sets the value for LastMobile to be an explicit nil

### UnsetLastMobile
`func (o *LimitedUserFriend) UnsetLastMobile()`

UnsetLastMobile ensures that no value is present for LastMobile, not even an explicit nil
### GetLastPlatform

`func (o *LimitedUserFriend) GetLastPlatform() string`

GetLastPlatform returns the LastPlatform field if non-nil, zero value otherwise.

### GetLastPlatformOk

`func (o *LimitedUserFriend) GetLastPlatformOk() (*string, bool)`

GetLastPlatformOk returns a tuple with the LastPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPlatform

`func (o *LimitedUserFriend) SetLastPlatform(v string)`

SetLastPlatform sets LastPlatform field to given value.


### GetLocation

`func (o *LimitedUserFriend) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *LimitedUserFriend) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *LimitedUserFriend) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetPlatform

`func (o *LimitedUserFriend) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *LimitedUserFriend) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *LimitedUserFriend) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetProfilePicOverride

`func (o *LimitedUserFriend) GetProfilePicOverride() string`

GetProfilePicOverride returns the ProfilePicOverride field if non-nil, zero value otherwise.

### GetProfilePicOverrideOk

`func (o *LimitedUserFriend) GetProfilePicOverrideOk() (*string, bool)`

GetProfilePicOverrideOk returns a tuple with the ProfilePicOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePicOverride

`func (o *LimitedUserFriend) SetProfilePicOverride(v string)`

SetProfilePicOverride sets ProfilePicOverride field to given value.

### HasProfilePicOverride

`func (o *LimitedUserFriend) HasProfilePicOverride() bool`

HasProfilePicOverride returns a boolean if a field has been set.

### GetProfilePicOverrideThumbnail

`func (o *LimitedUserFriend) GetProfilePicOverrideThumbnail() string`

GetProfilePicOverrideThumbnail returns the ProfilePicOverrideThumbnail field if non-nil, zero value otherwise.

### GetProfilePicOverrideThumbnailOk

`func (o *LimitedUserFriend) GetProfilePicOverrideThumbnailOk() (*string, bool)`

GetProfilePicOverrideThumbnailOk returns a tuple with the ProfilePicOverrideThumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePicOverrideThumbnail

`func (o *LimitedUserFriend) SetProfilePicOverrideThumbnail(v string)`

SetProfilePicOverrideThumbnail sets ProfilePicOverrideThumbnail field to given value.

### HasProfilePicOverrideThumbnail

`func (o *LimitedUserFriend) HasProfilePicOverrideThumbnail() bool`

HasProfilePicOverrideThumbnail returns a boolean if a field has been set.

### GetStatus

`func (o *LimitedUserFriend) GetStatus() UserStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LimitedUserFriend) GetStatusOk() (*UserStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LimitedUserFriend) SetStatus(v UserStatus)`

SetStatus sets Status field to given value.


### GetStatusDescription

`func (o *LimitedUserFriend) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *LimitedUserFriend) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *LimitedUserFriend) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.


### GetTags

`func (o *LimitedUserFriend) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LimitedUserFriend) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LimitedUserFriend) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetUserIcon

`func (o *LimitedUserFriend) GetUserIcon() string`

GetUserIcon returns the UserIcon field if non-nil, zero value otherwise.

### GetUserIconOk

`func (o *LimitedUserFriend) GetUserIconOk() (*string, bool)`

GetUserIconOk returns a tuple with the UserIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIcon

`func (o *LimitedUserFriend) SetUserIcon(v string)`

SetUserIcon sets UserIcon field to given value.

### HasUserIcon

`func (o *LimitedUserFriend) HasUserIcon() bool`

HasUserIcon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


