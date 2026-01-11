# LimitedUserInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgeVerificationStatus** | [**AgeVerificationStatus**](AgeVerificationStatus.md) |  | 
**AgeVerified** | **bool** | &#x60;true&#x60; if, user is age verified (not 18+). | 
**AllowAvatarCopying** | **bool** |  | 
**Bio** | Pointer to **string** |  | [optional] 
**BioLinks** | Pointer to **[]string** |   | [optional] 
**CurrentAvatarImageUrl** | **string** | When profilePicOverride is not empty, use it instead. | 
**CurrentAvatarTags** | **[]string** |  | 
**CurrentAvatarThumbnailImageUrl** | **string** | When profilePicOverride is not empty, use it instead. | 
**DateJoined** | **NullableTime** |  | 
**DeveloperType** | [**DeveloperType**](DeveloperType.md) |  | [default to NONE]
**DisplayName** | **string** |  | 
**FriendKey** | **string** |  | 
**Id** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**ImageUrl** | Pointer to **string** |  | [optional] 
**IsFriend** | **bool** |  | 
**LastActivity** | **NullableTime** |  | 
**LastMobile** | Pointer to **NullableTime** |  | [optional] 
**LastPlatform** | **string** | This is normally &#x60;android&#x60;, &#x60;ios&#x60;, &#x60;standalonewindows&#x60;, &#x60;web&#x60;, or the empty value &#x60;&#x60;, but also supposedly can be any random Unity version such as &#x60;2019.2.4-801-Release&#x60; or &#x60;2019.2.2-772-Release&#x60; or even &#x60;unknownplatform&#x60;. | 
**Platform** | Pointer to **string** | This is normally &#x60;android&#x60;, &#x60;ios&#x60;, &#x60;standalonewindows&#x60;, &#x60;web&#x60;, or the empty value &#x60;&#x60;, but also supposedly can be any random Unity version such as &#x60;2019.2.4-801-Release&#x60; or &#x60;2019.2.2-772-Release&#x60; or even &#x60;unknownplatform&#x60;. | [optional] 
**ProfilePicOverride** | Pointer to **string** |  | [optional] 
**ProfilePicOverrideThumbnail** | Pointer to **string** |  | [optional] 
**Pronouns** | **string** |  | 
**State** | Pointer to [**UserState**](UserState.md) |  | [optional] [default to OFFLINE]
**Status** | [**UserStatus**](UserStatus.md) |  | [default to OFFLINE]
**StatusDescription** | **string** |  | 
**Tags** | **[]string** |  | 
**UserIcon** | Pointer to **string** |  | [optional] 

## Methods

### NewLimitedUserInstance

`func NewLimitedUserInstance(ageVerificationStatus AgeVerificationStatus, ageVerified bool, allowAvatarCopying bool, currentAvatarImageUrl string, currentAvatarTags []string, currentAvatarThumbnailImageUrl string, dateJoined NullableTime, developerType DeveloperType, displayName string, friendKey string, id string, isFriend bool, lastActivity NullableTime, lastPlatform string, pronouns string, status UserStatus, statusDescription string, tags []string, ) *LimitedUserInstance`

NewLimitedUserInstance instantiates a new LimitedUserInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLimitedUserInstanceWithDefaults

`func NewLimitedUserInstanceWithDefaults() *LimitedUserInstance`

NewLimitedUserInstanceWithDefaults instantiates a new LimitedUserInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgeVerificationStatus

`func (o *LimitedUserInstance) GetAgeVerificationStatus() AgeVerificationStatus`

GetAgeVerificationStatus returns the AgeVerificationStatus field if non-nil, zero value otherwise.

### GetAgeVerificationStatusOk

`func (o *LimitedUserInstance) GetAgeVerificationStatusOk() (*AgeVerificationStatus, bool)`

GetAgeVerificationStatusOk returns a tuple with the AgeVerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeVerificationStatus

`func (o *LimitedUserInstance) SetAgeVerificationStatus(v AgeVerificationStatus)`

SetAgeVerificationStatus sets AgeVerificationStatus field to given value.


### GetAgeVerified

`func (o *LimitedUserInstance) GetAgeVerified() bool`

GetAgeVerified returns the AgeVerified field if non-nil, zero value otherwise.

### GetAgeVerifiedOk

`func (o *LimitedUserInstance) GetAgeVerifiedOk() (*bool, bool)`

GetAgeVerifiedOk returns a tuple with the AgeVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeVerified

`func (o *LimitedUserInstance) SetAgeVerified(v bool)`

SetAgeVerified sets AgeVerified field to given value.


### GetAllowAvatarCopying

`func (o *LimitedUserInstance) GetAllowAvatarCopying() bool`

GetAllowAvatarCopying returns the AllowAvatarCopying field if non-nil, zero value otherwise.

### GetAllowAvatarCopyingOk

`func (o *LimitedUserInstance) GetAllowAvatarCopyingOk() (*bool, bool)`

GetAllowAvatarCopyingOk returns a tuple with the AllowAvatarCopying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAvatarCopying

`func (o *LimitedUserInstance) SetAllowAvatarCopying(v bool)`

SetAllowAvatarCopying sets AllowAvatarCopying field to given value.


### GetBio

`func (o *LimitedUserInstance) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *LimitedUserInstance) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *LimitedUserInstance) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *LimitedUserInstance) HasBio() bool`

HasBio returns a boolean if a field has been set.

### GetBioLinks

`func (o *LimitedUserInstance) GetBioLinks() []string`

GetBioLinks returns the BioLinks field if non-nil, zero value otherwise.

### GetBioLinksOk

`func (o *LimitedUserInstance) GetBioLinksOk() (*[]string, bool)`

GetBioLinksOk returns a tuple with the BioLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBioLinks

`func (o *LimitedUserInstance) SetBioLinks(v []string)`

SetBioLinks sets BioLinks field to given value.

### HasBioLinks

`func (o *LimitedUserInstance) HasBioLinks() bool`

HasBioLinks returns a boolean if a field has been set.

### GetCurrentAvatarImageUrl

`func (o *LimitedUserInstance) GetCurrentAvatarImageUrl() string`

GetCurrentAvatarImageUrl returns the CurrentAvatarImageUrl field if non-nil, zero value otherwise.

### GetCurrentAvatarImageUrlOk

`func (o *LimitedUserInstance) GetCurrentAvatarImageUrlOk() (*string, bool)`

GetCurrentAvatarImageUrlOk returns a tuple with the CurrentAvatarImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAvatarImageUrl

`func (o *LimitedUserInstance) SetCurrentAvatarImageUrl(v string)`

SetCurrentAvatarImageUrl sets CurrentAvatarImageUrl field to given value.


### GetCurrentAvatarTags

`func (o *LimitedUserInstance) GetCurrentAvatarTags() []string`

GetCurrentAvatarTags returns the CurrentAvatarTags field if non-nil, zero value otherwise.

### GetCurrentAvatarTagsOk

`func (o *LimitedUserInstance) GetCurrentAvatarTagsOk() (*[]string, bool)`

GetCurrentAvatarTagsOk returns a tuple with the CurrentAvatarTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAvatarTags

`func (o *LimitedUserInstance) SetCurrentAvatarTags(v []string)`

SetCurrentAvatarTags sets CurrentAvatarTags field to given value.


### GetCurrentAvatarThumbnailImageUrl

`func (o *LimitedUserInstance) GetCurrentAvatarThumbnailImageUrl() string`

GetCurrentAvatarThumbnailImageUrl returns the CurrentAvatarThumbnailImageUrl field if non-nil, zero value otherwise.

### GetCurrentAvatarThumbnailImageUrlOk

`func (o *LimitedUserInstance) GetCurrentAvatarThumbnailImageUrlOk() (*string, bool)`

GetCurrentAvatarThumbnailImageUrlOk returns a tuple with the CurrentAvatarThumbnailImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAvatarThumbnailImageUrl

`func (o *LimitedUserInstance) SetCurrentAvatarThumbnailImageUrl(v string)`

SetCurrentAvatarThumbnailImageUrl sets CurrentAvatarThumbnailImageUrl field to given value.


### GetDateJoined

`func (o *LimitedUserInstance) GetDateJoined() time.Time`

GetDateJoined returns the DateJoined field if non-nil, zero value otherwise.

### GetDateJoinedOk

`func (o *LimitedUserInstance) GetDateJoinedOk() (*time.Time, bool)`

GetDateJoinedOk returns a tuple with the DateJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateJoined

`func (o *LimitedUserInstance) SetDateJoined(v time.Time)`

SetDateJoined sets DateJoined field to given value.


### SetDateJoinedNil

`func (o *LimitedUserInstance) SetDateJoinedNil(b bool)`

 SetDateJoinedNil sets the value for DateJoined to be an explicit nil

### UnsetDateJoined
`func (o *LimitedUserInstance) UnsetDateJoined()`

UnsetDateJoined ensures that no value is present for DateJoined, not even an explicit nil
### GetDeveloperType

`func (o *LimitedUserInstance) GetDeveloperType() DeveloperType`

GetDeveloperType returns the DeveloperType field if non-nil, zero value otherwise.

### GetDeveloperTypeOk

`func (o *LimitedUserInstance) GetDeveloperTypeOk() (*DeveloperType, bool)`

GetDeveloperTypeOk returns a tuple with the DeveloperType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperType

`func (o *LimitedUserInstance) SetDeveloperType(v DeveloperType)`

SetDeveloperType sets DeveloperType field to given value.


### GetDisplayName

`func (o *LimitedUserInstance) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *LimitedUserInstance) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *LimitedUserInstance) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetFriendKey

`func (o *LimitedUserInstance) GetFriendKey() string`

GetFriendKey returns the FriendKey field if non-nil, zero value otherwise.

### GetFriendKeyOk

`func (o *LimitedUserInstance) GetFriendKeyOk() (*string, bool)`

GetFriendKeyOk returns a tuple with the FriendKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendKey

`func (o *LimitedUserInstance) SetFriendKey(v string)`

SetFriendKey sets FriendKey field to given value.


### GetId

`func (o *LimitedUserInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LimitedUserInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LimitedUserInstance) SetId(v string)`

SetId sets Id field to given value.


### GetImageUrl

`func (o *LimitedUserInstance) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *LimitedUserInstance) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *LimitedUserInstance) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *LimitedUserInstance) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetIsFriend

`func (o *LimitedUserInstance) GetIsFriend() bool`

GetIsFriend returns the IsFriend field if non-nil, zero value otherwise.

### GetIsFriendOk

`func (o *LimitedUserInstance) GetIsFriendOk() (*bool, bool)`

GetIsFriendOk returns a tuple with the IsFriend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFriend

`func (o *LimitedUserInstance) SetIsFriend(v bool)`

SetIsFriend sets IsFriend field to given value.


### GetLastActivity

`func (o *LimitedUserInstance) GetLastActivity() time.Time`

GetLastActivity returns the LastActivity field if non-nil, zero value otherwise.

### GetLastActivityOk

`func (o *LimitedUserInstance) GetLastActivityOk() (*time.Time, bool)`

GetLastActivityOk returns a tuple with the LastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivity

`func (o *LimitedUserInstance) SetLastActivity(v time.Time)`

SetLastActivity sets LastActivity field to given value.


### SetLastActivityNil

`func (o *LimitedUserInstance) SetLastActivityNil(b bool)`

 SetLastActivityNil sets the value for LastActivity to be an explicit nil

### UnsetLastActivity
`func (o *LimitedUserInstance) UnsetLastActivity()`

UnsetLastActivity ensures that no value is present for LastActivity, not even an explicit nil
### GetLastMobile

`func (o *LimitedUserInstance) GetLastMobile() time.Time`

GetLastMobile returns the LastMobile field if non-nil, zero value otherwise.

### GetLastMobileOk

`func (o *LimitedUserInstance) GetLastMobileOk() (*time.Time, bool)`

GetLastMobileOk returns a tuple with the LastMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMobile

`func (o *LimitedUserInstance) SetLastMobile(v time.Time)`

SetLastMobile sets LastMobile field to given value.

### HasLastMobile

`func (o *LimitedUserInstance) HasLastMobile() bool`

HasLastMobile returns a boolean if a field has been set.

### SetLastMobileNil

`func (o *LimitedUserInstance) SetLastMobileNil(b bool)`

 SetLastMobileNil sets the value for LastMobile to be an explicit nil

### UnsetLastMobile
`func (o *LimitedUserInstance) UnsetLastMobile()`

UnsetLastMobile ensures that no value is present for LastMobile, not even an explicit nil
### GetLastPlatform

`func (o *LimitedUserInstance) GetLastPlatform() string`

GetLastPlatform returns the LastPlatform field if non-nil, zero value otherwise.

### GetLastPlatformOk

`func (o *LimitedUserInstance) GetLastPlatformOk() (*string, bool)`

GetLastPlatformOk returns a tuple with the LastPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPlatform

`func (o *LimitedUserInstance) SetLastPlatform(v string)`

SetLastPlatform sets LastPlatform field to given value.


### GetPlatform

`func (o *LimitedUserInstance) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *LimitedUserInstance) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *LimitedUserInstance) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *LimitedUserInstance) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetProfilePicOverride

`func (o *LimitedUserInstance) GetProfilePicOverride() string`

GetProfilePicOverride returns the ProfilePicOverride field if non-nil, zero value otherwise.

### GetProfilePicOverrideOk

`func (o *LimitedUserInstance) GetProfilePicOverrideOk() (*string, bool)`

GetProfilePicOverrideOk returns a tuple with the ProfilePicOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePicOverride

`func (o *LimitedUserInstance) SetProfilePicOverride(v string)`

SetProfilePicOverride sets ProfilePicOverride field to given value.

### HasProfilePicOverride

`func (o *LimitedUserInstance) HasProfilePicOverride() bool`

HasProfilePicOverride returns a boolean if a field has been set.

### GetProfilePicOverrideThumbnail

`func (o *LimitedUserInstance) GetProfilePicOverrideThumbnail() string`

GetProfilePicOverrideThumbnail returns the ProfilePicOverrideThumbnail field if non-nil, zero value otherwise.

### GetProfilePicOverrideThumbnailOk

`func (o *LimitedUserInstance) GetProfilePicOverrideThumbnailOk() (*string, bool)`

GetProfilePicOverrideThumbnailOk returns a tuple with the ProfilePicOverrideThumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePicOverrideThumbnail

`func (o *LimitedUserInstance) SetProfilePicOverrideThumbnail(v string)`

SetProfilePicOverrideThumbnail sets ProfilePicOverrideThumbnail field to given value.

### HasProfilePicOverrideThumbnail

`func (o *LimitedUserInstance) HasProfilePicOverrideThumbnail() bool`

HasProfilePicOverrideThumbnail returns a boolean if a field has been set.

### GetPronouns

`func (o *LimitedUserInstance) GetPronouns() string`

GetPronouns returns the Pronouns field if non-nil, zero value otherwise.

### GetPronounsOk

`func (o *LimitedUserInstance) GetPronounsOk() (*string, bool)`

GetPronounsOk returns a tuple with the Pronouns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPronouns

`func (o *LimitedUserInstance) SetPronouns(v string)`

SetPronouns sets Pronouns field to given value.


### GetState

`func (o *LimitedUserInstance) GetState() UserState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LimitedUserInstance) GetStateOk() (*UserState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LimitedUserInstance) SetState(v UserState)`

SetState sets State field to given value.

### HasState

`func (o *LimitedUserInstance) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatus

`func (o *LimitedUserInstance) GetStatus() UserStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LimitedUserInstance) GetStatusOk() (*UserStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LimitedUserInstance) SetStatus(v UserStatus)`

SetStatus sets Status field to given value.


### GetStatusDescription

`func (o *LimitedUserInstance) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *LimitedUserInstance) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *LimitedUserInstance) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.


### GetTags

`func (o *LimitedUserInstance) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LimitedUserInstance) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LimitedUserInstance) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetUserIcon

`func (o *LimitedUserInstance) GetUserIcon() string`

GetUserIcon returns the UserIcon field if non-nil, zero value otherwise.

### GetUserIconOk

`func (o *LimitedUserInstance) GetUserIconOk() (*string, bool)`

GetUserIconOk returns a tuple with the UserIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIcon

`func (o *LimitedUserInstance) SetUserIcon(v string)`

SetUserIcon sets UserIcon field to given value.

### HasUserIcon

`func (o *LimitedUserInstance) HasUserIcon() bool`

HasUserIcon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


