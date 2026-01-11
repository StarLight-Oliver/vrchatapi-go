# RegisterUserAccount200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedPrivacyVersion** | Pointer to **int32** |  | [optional] 
**AcceptedTOSVersion** | **int32** |  | 
**AccountDeletionDate** | Pointer to **NullableString** |  | [optional] 
**AccountDeletionLog** | Pointer to [**[]AccountDeletionLog**](AccountDeletionLog.md) |   | [optional] 
**ActiveFriends** | Pointer to **[]string** |   | [optional] 
**AgeVerificationStatus** | [**AgeVerificationStatus**](AgeVerificationStatus.md) |  | 
**AgeVerified** | **bool** | &#x60;true&#x60; if, user is age verified (not 18+). | 
**AllowAvatarCopying** | **bool** |  | 
**AuthToken** | Pointer to **string** | The auth token for NEWLY REGISTERED ACCOUNTS ONLY (/auth/register) | [optional] 
**Badges** | Pointer to [**[]Badge**](Badge.md) |   | [optional] 
**Bio** | **string** |  | 
**BioLinks** | **[]string** |   | 
**ContentFilters** | Pointer to **[]string** | These tags begin with &#x60;content_&#x60; and control content gating | [optional] 
**CurrentAvatar** | **string** |  | 
**CurrentAvatarImageUrl** | **string** | When profilePicOverride is not empty, use it instead. | 
**CurrentAvatarTags** | **[]string** |  | 
**CurrentAvatarThumbnailImageUrl** | **string** | When profilePicOverride is not empty, use it instead. | 
**DateJoined** | **string** |  | 
**DeveloperType** | [**DeveloperType**](DeveloperType.md) |  | [default to NONE]
**DiscordDetails** | Pointer to [**DiscordDetails**](DiscordDetails.md) |  | [optional] 
**DiscordId** | Pointer to **string** | https://discord.com/developers/docs/reference#snowflakes | [optional] 
**DisplayName** | **string** |  | 
**EmailVerified** | **bool** |  | 
**FallbackAvatar** | Pointer to **string** |  | [optional] 
**FriendGroupNames** | **[]string** | Always empty array. | 
**FriendKey** | **string** |  | 
**Friends** | **[]string** |  | 
**GoogleDetails** | Pointer to **map[string]interface{}** |  | [optional] 
**GoogleId** | Pointer to **string** |  | [optional] 
**HasBirthday** | **bool** |  | 
**HasEmail** | **bool** |  | 
**HasLoggedInFromClient** | **bool** |  | 
**HasPendingEmail** | **bool** |  | 
**HideContentFilterSettings** | Pointer to **bool** |  | [optional] 
**HomeLocation** | **string** | WorldID be \&quot;offline\&quot; on User profiles if you are not friends with that user. | 
**Id** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**IsAdult** | **bool** |  | 
**IsBoopingEnabled** | Pointer to **bool** |  | [optional] [default to true]
**IsFriend** | **bool** |  | [default to false]
**LastActivity** | Pointer to **time.Time** |  | [optional] 
**LastLogin** | **time.Time** |  | 
**LastMobile** | **NullableTime** |  | 
**LastPlatform** | **string** | This is normally &#x60;android&#x60;, &#x60;ios&#x60;, &#x60;standalonewindows&#x60;, &#x60;web&#x60;, or the empty value &#x60;&#x60;, but also supposedly can be any random Unity version such as &#x60;2019.2.4-801-Release&#x60; or &#x60;2019.2.2-772-Release&#x60; or even &#x60;unknownplatform&#x60;. | 
**ObfuscatedEmail** | **string** |  | 
**ObfuscatedPendingEmail** | **string** |  | 
**OculusId** | **string** |  | 
**OfflineFriends** | Pointer to **[]string** |  | [optional] 
**OnlineFriends** | Pointer to **[]string** |  | [optional] 
**PastDisplayNames** | [**[]PastDisplayName**](PastDisplayName.md) |   | 
**PicoId** | Pointer to **string** |  | [optional] 
**PlatformHistory** | Pointer to [**[]CurrentUserPlatformHistoryInner**](CurrentUserPlatformHistoryInner.md) |  | [optional] 
**Presence** | Pointer to [**CurrentUserPresence**](CurrentUserPresence.md) |  | [optional] 
**ProfilePicOverride** | **string** |  | 
**ProfilePicOverrideThumbnail** | **string** |  | 
**Pronouns** | **string** |  | 
**PronounsHistory** | **[]string** |  | 
**QueuedInstance** | Pointer to **NullableString** |  | [optional] 
**ReceiveMobileInvitations** | Pointer to **bool** |  | [optional] 
**State** | [**UserState**](UserState.md) |  | [default to OFFLINE]
**Status** | [**UserStatus**](UserStatus.md) |  | [default to OFFLINE]
**StatusDescription** | **string** |  | 
**StatusFirstTime** | **bool** |  | 
**StatusHistory** | **[]string** |  | 
**SteamDetails** | **map[string]interface{}** |  | 
**SteamId** | **string** |  | 
**Tags** | **[]string** |  | 
**TwoFactorAuthEnabled** | **bool** |  | 
**TwoFactorAuthEnabledDate** | Pointer to **NullableTime** |  | [optional] 
**Unsubscribe** | **bool** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UserIcon** | **string** |  | 
**UserLanguage** | Pointer to **NullableString** |  | [optional] 
**UserLanguageCode** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **string** | -| **DEPRECATED:** VRChat API no longer return usernames of other users. [See issue by Tupper for more information](https://github.com/pypy-vrc/VRCX/issues/429). | [optional] 
**UsesGeneratedPassword** | **bool** |  | 
**ViveId** | Pointer to **string** |  | [optional] 
**RequiresTwoFactorAuth** | [**[]TwoFactorAuthType**](TwoFactorAuthType.md) |  | 

## Methods

### NewRegisterUserAccount200Response

`func NewRegisterUserAccount200Response(acceptedTOSVersion int32, ageVerificationStatus AgeVerificationStatus, ageVerified bool, allowAvatarCopying bool, bio string, bioLinks []string, currentAvatar string, currentAvatarImageUrl string, currentAvatarTags []string, currentAvatarThumbnailImageUrl string, dateJoined string, developerType DeveloperType, displayName string, emailVerified bool, friendGroupNames []string, friendKey string, friends []string, hasBirthday bool, hasEmail bool, hasLoggedInFromClient bool, hasPendingEmail bool, homeLocation string, id string, isAdult bool, isFriend bool, lastLogin time.Time, lastMobile NullableTime, lastPlatform string, obfuscatedEmail string, obfuscatedPendingEmail string, oculusId string, pastDisplayNames []PastDisplayName, profilePicOverride string, profilePicOverrideThumbnail string, pronouns string, pronounsHistory []string, state UserState, status UserStatus, statusDescription string, statusFirstTime bool, statusHistory []string, steamDetails map[string]interface{}, steamId string, tags []string, twoFactorAuthEnabled bool, unsubscribe bool, userIcon string, usesGeneratedPassword bool, requiresTwoFactorAuth []TwoFactorAuthType, ) *RegisterUserAccount200Response`

NewRegisterUserAccount200Response instantiates a new RegisterUserAccount200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterUserAccount200ResponseWithDefaults

`func NewRegisterUserAccount200ResponseWithDefaults() *RegisterUserAccount200Response`

NewRegisterUserAccount200ResponseWithDefaults instantiates a new RegisterUserAccount200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedPrivacyVersion

`func (o *RegisterUserAccount200Response) GetAcceptedPrivacyVersion() int32`

GetAcceptedPrivacyVersion returns the AcceptedPrivacyVersion field if non-nil, zero value otherwise.

### GetAcceptedPrivacyVersionOk

`func (o *RegisterUserAccount200Response) GetAcceptedPrivacyVersionOk() (*int32, bool)`

GetAcceptedPrivacyVersionOk returns a tuple with the AcceptedPrivacyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedPrivacyVersion

`func (o *RegisterUserAccount200Response) SetAcceptedPrivacyVersion(v int32)`

SetAcceptedPrivacyVersion sets AcceptedPrivacyVersion field to given value.

### HasAcceptedPrivacyVersion

`func (o *RegisterUserAccount200Response) HasAcceptedPrivacyVersion() bool`

HasAcceptedPrivacyVersion returns a boolean if a field has been set.

### GetAcceptedTOSVersion

`func (o *RegisterUserAccount200Response) GetAcceptedTOSVersion() int32`

GetAcceptedTOSVersion returns the AcceptedTOSVersion field if non-nil, zero value otherwise.

### GetAcceptedTOSVersionOk

`func (o *RegisterUserAccount200Response) GetAcceptedTOSVersionOk() (*int32, bool)`

GetAcceptedTOSVersionOk returns a tuple with the AcceptedTOSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedTOSVersion

`func (o *RegisterUserAccount200Response) SetAcceptedTOSVersion(v int32)`

SetAcceptedTOSVersion sets AcceptedTOSVersion field to given value.


### GetAccountDeletionDate

`func (o *RegisterUserAccount200Response) GetAccountDeletionDate() string`

GetAccountDeletionDate returns the AccountDeletionDate field if non-nil, zero value otherwise.

### GetAccountDeletionDateOk

`func (o *RegisterUserAccount200Response) GetAccountDeletionDateOk() (*string, bool)`

GetAccountDeletionDateOk returns a tuple with the AccountDeletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDeletionDate

`func (o *RegisterUserAccount200Response) SetAccountDeletionDate(v string)`

SetAccountDeletionDate sets AccountDeletionDate field to given value.

### HasAccountDeletionDate

`func (o *RegisterUserAccount200Response) HasAccountDeletionDate() bool`

HasAccountDeletionDate returns a boolean if a field has been set.

### SetAccountDeletionDateNil

`func (o *RegisterUserAccount200Response) SetAccountDeletionDateNil(b bool)`

 SetAccountDeletionDateNil sets the value for AccountDeletionDate to be an explicit nil

### UnsetAccountDeletionDate
`func (o *RegisterUserAccount200Response) UnsetAccountDeletionDate()`

UnsetAccountDeletionDate ensures that no value is present for AccountDeletionDate, not even an explicit nil
### GetAccountDeletionLog

`func (o *RegisterUserAccount200Response) GetAccountDeletionLog() []AccountDeletionLog`

GetAccountDeletionLog returns the AccountDeletionLog field if non-nil, zero value otherwise.

### GetAccountDeletionLogOk

`func (o *RegisterUserAccount200Response) GetAccountDeletionLogOk() (*[]AccountDeletionLog, bool)`

GetAccountDeletionLogOk returns a tuple with the AccountDeletionLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDeletionLog

`func (o *RegisterUserAccount200Response) SetAccountDeletionLog(v []AccountDeletionLog)`

SetAccountDeletionLog sets AccountDeletionLog field to given value.

### HasAccountDeletionLog

`func (o *RegisterUserAccount200Response) HasAccountDeletionLog() bool`

HasAccountDeletionLog returns a boolean if a field has been set.

### SetAccountDeletionLogNil

`func (o *RegisterUserAccount200Response) SetAccountDeletionLogNil(b bool)`

 SetAccountDeletionLogNil sets the value for AccountDeletionLog to be an explicit nil

### UnsetAccountDeletionLog
`func (o *RegisterUserAccount200Response) UnsetAccountDeletionLog()`

UnsetAccountDeletionLog ensures that no value is present for AccountDeletionLog, not even an explicit nil
### GetActiveFriends

`func (o *RegisterUserAccount200Response) GetActiveFriends() []string`

GetActiveFriends returns the ActiveFriends field if non-nil, zero value otherwise.

### GetActiveFriendsOk

`func (o *RegisterUserAccount200Response) GetActiveFriendsOk() (*[]string, bool)`

GetActiveFriendsOk returns a tuple with the ActiveFriends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveFriends

`func (o *RegisterUserAccount200Response) SetActiveFriends(v []string)`

SetActiveFriends sets ActiveFriends field to given value.

### HasActiveFriends

`func (o *RegisterUserAccount200Response) HasActiveFriends() bool`

HasActiveFriends returns a boolean if a field has been set.

### GetAgeVerificationStatus

`func (o *RegisterUserAccount200Response) GetAgeVerificationStatus() AgeVerificationStatus`

GetAgeVerificationStatus returns the AgeVerificationStatus field if non-nil, zero value otherwise.

### GetAgeVerificationStatusOk

`func (o *RegisterUserAccount200Response) GetAgeVerificationStatusOk() (*AgeVerificationStatus, bool)`

GetAgeVerificationStatusOk returns a tuple with the AgeVerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeVerificationStatus

`func (o *RegisterUserAccount200Response) SetAgeVerificationStatus(v AgeVerificationStatus)`

SetAgeVerificationStatus sets AgeVerificationStatus field to given value.


### GetAgeVerified

`func (o *RegisterUserAccount200Response) GetAgeVerified() bool`

GetAgeVerified returns the AgeVerified field if non-nil, zero value otherwise.

### GetAgeVerifiedOk

`func (o *RegisterUserAccount200Response) GetAgeVerifiedOk() (*bool, bool)`

GetAgeVerifiedOk returns a tuple with the AgeVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeVerified

`func (o *RegisterUserAccount200Response) SetAgeVerified(v bool)`

SetAgeVerified sets AgeVerified field to given value.


### GetAllowAvatarCopying

`func (o *RegisterUserAccount200Response) GetAllowAvatarCopying() bool`

GetAllowAvatarCopying returns the AllowAvatarCopying field if non-nil, zero value otherwise.

### GetAllowAvatarCopyingOk

`func (o *RegisterUserAccount200Response) GetAllowAvatarCopyingOk() (*bool, bool)`

GetAllowAvatarCopyingOk returns a tuple with the AllowAvatarCopying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAvatarCopying

`func (o *RegisterUserAccount200Response) SetAllowAvatarCopying(v bool)`

SetAllowAvatarCopying sets AllowAvatarCopying field to given value.


### GetAuthToken

`func (o *RegisterUserAccount200Response) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *RegisterUserAccount200Response) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *RegisterUserAccount200Response) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.

### HasAuthToken

`func (o *RegisterUserAccount200Response) HasAuthToken() bool`

HasAuthToken returns a boolean if a field has been set.

### GetBadges

`func (o *RegisterUserAccount200Response) GetBadges() []Badge`

GetBadges returns the Badges field if non-nil, zero value otherwise.

### GetBadgesOk

`func (o *RegisterUserAccount200Response) GetBadgesOk() (*[]Badge, bool)`

GetBadgesOk returns a tuple with the Badges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadges

`func (o *RegisterUserAccount200Response) SetBadges(v []Badge)`

SetBadges sets Badges field to given value.

### HasBadges

`func (o *RegisterUserAccount200Response) HasBadges() bool`

HasBadges returns a boolean if a field has been set.

### GetBio

`func (o *RegisterUserAccount200Response) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *RegisterUserAccount200Response) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *RegisterUserAccount200Response) SetBio(v string)`

SetBio sets Bio field to given value.


### GetBioLinks

`func (o *RegisterUserAccount200Response) GetBioLinks() []string`

GetBioLinks returns the BioLinks field if non-nil, zero value otherwise.

### GetBioLinksOk

`func (o *RegisterUserAccount200Response) GetBioLinksOk() (*[]string, bool)`

GetBioLinksOk returns a tuple with the BioLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBioLinks

`func (o *RegisterUserAccount200Response) SetBioLinks(v []string)`

SetBioLinks sets BioLinks field to given value.


### GetContentFilters

`func (o *RegisterUserAccount200Response) GetContentFilters() []string`

GetContentFilters returns the ContentFilters field if non-nil, zero value otherwise.

### GetContentFiltersOk

`func (o *RegisterUserAccount200Response) GetContentFiltersOk() (*[]string, bool)`

GetContentFiltersOk returns a tuple with the ContentFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFilters

`func (o *RegisterUserAccount200Response) SetContentFilters(v []string)`

SetContentFilters sets ContentFilters field to given value.

### HasContentFilters

`func (o *RegisterUserAccount200Response) HasContentFilters() bool`

HasContentFilters returns a boolean if a field has been set.

### GetCurrentAvatar

`func (o *RegisterUserAccount200Response) GetCurrentAvatar() string`

GetCurrentAvatar returns the CurrentAvatar field if non-nil, zero value otherwise.

### GetCurrentAvatarOk

`func (o *RegisterUserAccount200Response) GetCurrentAvatarOk() (*string, bool)`

GetCurrentAvatarOk returns a tuple with the CurrentAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAvatar

`func (o *RegisterUserAccount200Response) SetCurrentAvatar(v string)`

SetCurrentAvatar sets CurrentAvatar field to given value.


### GetCurrentAvatarImageUrl

`func (o *RegisterUserAccount200Response) GetCurrentAvatarImageUrl() string`

GetCurrentAvatarImageUrl returns the CurrentAvatarImageUrl field if non-nil, zero value otherwise.

### GetCurrentAvatarImageUrlOk

`func (o *RegisterUserAccount200Response) GetCurrentAvatarImageUrlOk() (*string, bool)`

GetCurrentAvatarImageUrlOk returns a tuple with the CurrentAvatarImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAvatarImageUrl

`func (o *RegisterUserAccount200Response) SetCurrentAvatarImageUrl(v string)`

SetCurrentAvatarImageUrl sets CurrentAvatarImageUrl field to given value.


### GetCurrentAvatarTags

`func (o *RegisterUserAccount200Response) GetCurrentAvatarTags() []string`

GetCurrentAvatarTags returns the CurrentAvatarTags field if non-nil, zero value otherwise.

### GetCurrentAvatarTagsOk

`func (o *RegisterUserAccount200Response) GetCurrentAvatarTagsOk() (*[]string, bool)`

GetCurrentAvatarTagsOk returns a tuple with the CurrentAvatarTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAvatarTags

`func (o *RegisterUserAccount200Response) SetCurrentAvatarTags(v []string)`

SetCurrentAvatarTags sets CurrentAvatarTags field to given value.


### GetCurrentAvatarThumbnailImageUrl

`func (o *RegisterUserAccount200Response) GetCurrentAvatarThumbnailImageUrl() string`

GetCurrentAvatarThumbnailImageUrl returns the CurrentAvatarThumbnailImageUrl field if non-nil, zero value otherwise.

### GetCurrentAvatarThumbnailImageUrlOk

`func (o *RegisterUserAccount200Response) GetCurrentAvatarThumbnailImageUrlOk() (*string, bool)`

GetCurrentAvatarThumbnailImageUrlOk returns a tuple with the CurrentAvatarThumbnailImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAvatarThumbnailImageUrl

`func (o *RegisterUserAccount200Response) SetCurrentAvatarThumbnailImageUrl(v string)`

SetCurrentAvatarThumbnailImageUrl sets CurrentAvatarThumbnailImageUrl field to given value.


### GetDateJoined

`func (o *RegisterUserAccount200Response) GetDateJoined() string`

GetDateJoined returns the DateJoined field if non-nil, zero value otherwise.

### GetDateJoinedOk

`func (o *RegisterUserAccount200Response) GetDateJoinedOk() (*string, bool)`

GetDateJoinedOk returns a tuple with the DateJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateJoined

`func (o *RegisterUserAccount200Response) SetDateJoined(v string)`

SetDateJoined sets DateJoined field to given value.


### GetDeveloperType

`func (o *RegisterUserAccount200Response) GetDeveloperType() DeveloperType`

GetDeveloperType returns the DeveloperType field if non-nil, zero value otherwise.

### GetDeveloperTypeOk

`func (o *RegisterUserAccount200Response) GetDeveloperTypeOk() (*DeveloperType, bool)`

GetDeveloperTypeOk returns a tuple with the DeveloperType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperType

`func (o *RegisterUserAccount200Response) SetDeveloperType(v DeveloperType)`

SetDeveloperType sets DeveloperType field to given value.


### GetDiscordDetails

`func (o *RegisterUserAccount200Response) GetDiscordDetails() DiscordDetails`

GetDiscordDetails returns the DiscordDetails field if non-nil, zero value otherwise.

### GetDiscordDetailsOk

`func (o *RegisterUserAccount200Response) GetDiscordDetailsOk() (*DiscordDetails, bool)`

GetDiscordDetailsOk returns a tuple with the DiscordDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordDetails

`func (o *RegisterUserAccount200Response) SetDiscordDetails(v DiscordDetails)`

SetDiscordDetails sets DiscordDetails field to given value.

### HasDiscordDetails

`func (o *RegisterUserAccount200Response) HasDiscordDetails() bool`

HasDiscordDetails returns a boolean if a field has been set.

### GetDiscordId

`func (o *RegisterUserAccount200Response) GetDiscordId() string`

GetDiscordId returns the DiscordId field if non-nil, zero value otherwise.

### GetDiscordIdOk

`func (o *RegisterUserAccount200Response) GetDiscordIdOk() (*string, bool)`

GetDiscordIdOk returns a tuple with the DiscordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscordId

`func (o *RegisterUserAccount200Response) SetDiscordId(v string)`

SetDiscordId sets DiscordId field to given value.

### HasDiscordId

`func (o *RegisterUserAccount200Response) HasDiscordId() bool`

HasDiscordId returns a boolean if a field has been set.

### GetDisplayName

`func (o *RegisterUserAccount200Response) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RegisterUserAccount200Response) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *RegisterUserAccount200Response) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetEmailVerified

`func (o *RegisterUserAccount200Response) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *RegisterUserAccount200Response) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *RegisterUserAccount200Response) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.


### GetFallbackAvatar

`func (o *RegisterUserAccount200Response) GetFallbackAvatar() string`

GetFallbackAvatar returns the FallbackAvatar field if non-nil, zero value otherwise.

### GetFallbackAvatarOk

`func (o *RegisterUserAccount200Response) GetFallbackAvatarOk() (*string, bool)`

GetFallbackAvatarOk returns a tuple with the FallbackAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackAvatar

`func (o *RegisterUserAccount200Response) SetFallbackAvatar(v string)`

SetFallbackAvatar sets FallbackAvatar field to given value.

### HasFallbackAvatar

`func (o *RegisterUserAccount200Response) HasFallbackAvatar() bool`

HasFallbackAvatar returns a boolean if a field has been set.

### GetFriendGroupNames

`func (o *RegisterUserAccount200Response) GetFriendGroupNames() []string`

GetFriendGroupNames returns the FriendGroupNames field if non-nil, zero value otherwise.

### GetFriendGroupNamesOk

`func (o *RegisterUserAccount200Response) GetFriendGroupNamesOk() (*[]string, bool)`

GetFriendGroupNamesOk returns a tuple with the FriendGroupNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendGroupNames

`func (o *RegisterUserAccount200Response) SetFriendGroupNames(v []string)`

SetFriendGroupNames sets FriendGroupNames field to given value.


### GetFriendKey

`func (o *RegisterUserAccount200Response) GetFriendKey() string`

GetFriendKey returns the FriendKey field if non-nil, zero value otherwise.

### GetFriendKeyOk

`func (o *RegisterUserAccount200Response) GetFriendKeyOk() (*string, bool)`

GetFriendKeyOk returns a tuple with the FriendKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendKey

`func (o *RegisterUserAccount200Response) SetFriendKey(v string)`

SetFriendKey sets FriendKey field to given value.


### GetFriends

`func (o *RegisterUserAccount200Response) GetFriends() []string`

GetFriends returns the Friends field if non-nil, zero value otherwise.

### GetFriendsOk

`func (o *RegisterUserAccount200Response) GetFriendsOk() (*[]string, bool)`

GetFriendsOk returns a tuple with the Friends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriends

`func (o *RegisterUserAccount200Response) SetFriends(v []string)`

SetFriends sets Friends field to given value.


### GetGoogleDetails

`func (o *RegisterUserAccount200Response) GetGoogleDetails() map[string]interface{}`

GetGoogleDetails returns the GoogleDetails field if non-nil, zero value otherwise.

### GetGoogleDetailsOk

`func (o *RegisterUserAccount200Response) GetGoogleDetailsOk() (*map[string]interface{}, bool)`

GetGoogleDetailsOk returns a tuple with the GoogleDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleDetails

`func (o *RegisterUserAccount200Response) SetGoogleDetails(v map[string]interface{})`

SetGoogleDetails sets GoogleDetails field to given value.

### HasGoogleDetails

`func (o *RegisterUserAccount200Response) HasGoogleDetails() bool`

HasGoogleDetails returns a boolean if a field has been set.

### GetGoogleId

`func (o *RegisterUserAccount200Response) GetGoogleId() string`

GetGoogleId returns the GoogleId field if non-nil, zero value otherwise.

### GetGoogleIdOk

`func (o *RegisterUserAccount200Response) GetGoogleIdOk() (*string, bool)`

GetGoogleIdOk returns a tuple with the GoogleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleId

`func (o *RegisterUserAccount200Response) SetGoogleId(v string)`

SetGoogleId sets GoogleId field to given value.

### HasGoogleId

`func (o *RegisterUserAccount200Response) HasGoogleId() bool`

HasGoogleId returns a boolean if a field has been set.

### GetHasBirthday

`func (o *RegisterUserAccount200Response) GetHasBirthday() bool`

GetHasBirthday returns the HasBirthday field if non-nil, zero value otherwise.

### GetHasBirthdayOk

`func (o *RegisterUserAccount200Response) GetHasBirthdayOk() (*bool, bool)`

GetHasBirthdayOk returns a tuple with the HasBirthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasBirthday

`func (o *RegisterUserAccount200Response) SetHasBirthday(v bool)`

SetHasBirthday sets HasBirthday field to given value.


### GetHasEmail

`func (o *RegisterUserAccount200Response) GetHasEmail() bool`

GetHasEmail returns the HasEmail field if non-nil, zero value otherwise.

### GetHasEmailOk

`func (o *RegisterUserAccount200Response) GetHasEmailOk() (*bool, bool)`

GetHasEmailOk returns a tuple with the HasEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEmail

`func (o *RegisterUserAccount200Response) SetHasEmail(v bool)`

SetHasEmail sets HasEmail field to given value.


### GetHasLoggedInFromClient

`func (o *RegisterUserAccount200Response) GetHasLoggedInFromClient() bool`

GetHasLoggedInFromClient returns the HasLoggedInFromClient field if non-nil, zero value otherwise.

### GetHasLoggedInFromClientOk

`func (o *RegisterUserAccount200Response) GetHasLoggedInFromClientOk() (*bool, bool)`

GetHasLoggedInFromClientOk returns a tuple with the HasLoggedInFromClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLoggedInFromClient

`func (o *RegisterUserAccount200Response) SetHasLoggedInFromClient(v bool)`

SetHasLoggedInFromClient sets HasLoggedInFromClient field to given value.


### GetHasPendingEmail

`func (o *RegisterUserAccount200Response) GetHasPendingEmail() bool`

GetHasPendingEmail returns the HasPendingEmail field if non-nil, zero value otherwise.

### GetHasPendingEmailOk

`func (o *RegisterUserAccount200Response) GetHasPendingEmailOk() (*bool, bool)`

GetHasPendingEmailOk returns a tuple with the HasPendingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPendingEmail

`func (o *RegisterUserAccount200Response) SetHasPendingEmail(v bool)`

SetHasPendingEmail sets HasPendingEmail field to given value.


### GetHideContentFilterSettings

`func (o *RegisterUserAccount200Response) GetHideContentFilterSettings() bool`

GetHideContentFilterSettings returns the HideContentFilterSettings field if non-nil, zero value otherwise.

### GetHideContentFilterSettingsOk

`func (o *RegisterUserAccount200Response) GetHideContentFilterSettingsOk() (*bool, bool)`

GetHideContentFilterSettingsOk returns a tuple with the HideContentFilterSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideContentFilterSettings

`func (o *RegisterUserAccount200Response) SetHideContentFilterSettings(v bool)`

SetHideContentFilterSettings sets HideContentFilterSettings field to given value.

### HasHideContentFilterSettings

`func (o *RegisterUserAccount200Response) HasHideContentFilterSettings() bool`

HasHideContentFilterSettings returns a boolean if a field has been set.

### GetHomeLocation

`func (o *RegisterUserAccount200Response) GetHomeLocation() string`

GetHomeLocation returns the HomeLocation field if non-nil, zero value otherwise.

### GetHomeLocationOk

`func (o *RegisterUserAccount200Response) GetHomeLocationOk() (*string, bool)`

GetHomeLocationOk returns a tuple with the HomeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeLocation

`func (o *RegisterUserAccount200Response) SetHomeLocation(v string)`

SetHomeLocation sets HomeLocation field to given value.


### GetId

`func (o *RegisterUserAccount200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegisterUserAccount200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegisterUserAccount200Response) SetId(v string)`

SetId sets Id field to given value.


### GetIsAdult

`func (o *RegisterUserAccount200Response) GetIsAdult() bool`

GetIsAdult returns the IsAdult field if non-nil, zero value otherwise.

### GetIsAdultOk

`func (o *RegisterUserAccount200Response) GetIsAdultOk() (*bool, bool)`

GetIsAdultOk returns a tuple with the IsAdult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdult

`func (o *RegisterUserAccount200Response) SetIsAdult(v bool)`

SetIsAdult sets IsAdult field to given value.


### GetIsBoopingEnabled

`func (o *RegisterUserAccount200Response) GetIsBoopingEnabled() bool`

GetIsBoopingEnabled returns the IsBoopingEnabled field if non-nil, zero value otherwise.

### GetIsBoopingEnabledOk

`func (o *RegisterUserAccount200Response) GetIsBoopingEnabledOk() (*bool, bool)`

GetIsBoopingEnabledOk returns a tuple with the IsBoopingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBoopingEnabled

`func (o *RegisterUserAccount200Response) SetIsBoopingEnabled(v bool)`

SetIsBoopingEnabled sets IsBoopingEnabled field to given value.

### HasIsBoopingEnabled

`func (o *RegisterUserAccount200Response) HasIsBoopingEnabled() bool`

HasIsBoopingEnabled returns a boolean if a field has been set.

### GetIsFriend

`func (o *RegisterUserAccount200Response) GetIsFriend() bool`

GetIsFriend returns the IsFriend field if non-nil, zero value otherwise.

### GetIsFriendOk

`func (o *RegisterUserAccount200Response) GetIsFriendOk() (*bool, bool)`

GetIsFriendOk returns a tuple with the IsFriend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFriend

`func (o *RegisterUserAccount200Response) SetIsFriend(v bool)`

SetIsFriend sets IsFriend field to given value.


### GetLastActivity

`func (o *RegisterUserAccount200Response) GetLastActivity() time.Time`

GetLastActivity returns the LastActivity field if non-nil, zero value otherwise.

### GetLastActivityOk

`func (o *RegisterUserAccount200Response) GetLastActivityOk() (*time.Time, bool)`

GetLastActivityOk returns a tuple with the LastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivity

`func (o *RegisterUserAccount200Response) SetLastActivity(v time.Time)`

SetLastActivity sets LastActivity field to given value.

### HasLastActivity

`func (o *RegisterUserAccount200Response) HasLastActivity() bool`

HasLastActivity returns a boolean if a field has been set.

### GetLastLogin

`func (o *RegisterUserAccount200Response) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *RegisterUserAccount200Response) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *RegisterUserAccount200Response) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.


### GetLastMobile

`func (o *RegisterUserAccount200Response) GetLastMobile() time.Time`

GetLastMobile returns the LastMobile field if non-nil, zero value otherwise.

### GetLastMobileOk

`func (o *RegisterUserAccount200Response) GetLastMobileOk() (*time.Time, bool)`

GetLastMobileOk returns a tuple with the LastMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMobile

`func (o *RegisterUserAccount200Response) SetLastMobile(v time.Time)`

SetLastMobile sets LastMobile field to given value.


### SetLastMobileNil

`func (o *RegisterUserAccount200Response) SetLastMobileNil(b bool)`

 SetLastMobileNil sets the value for LastMobile to be an explicit nil

### UnsetLastMobile
`func (o *RegisterUserAccount200Response) UnsetLastMobile()`

UnsetLastMobile ensures that no value is present for LastMobile, not even an explicit nil
### GetLastPlatform

`func (o *RegisterUserAccount200Response) GetLastPlatform() string`

GetLastPlatform returns the LastPlatform field if non-nil, zero value otherwise.

### GetLastPlatformOk

`func (o *RegisterUserAccount200Response) GetLastPlatformOk() (*string, bool)`

GetLastPlatformOk returns a tuple with the LastPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPlatform

`func (o *RegisterUserAccount200Response) SetLastPlatform(v string)`

SetLastPlatform sets LastPlatform field to given value.


### GetObfuscatedEmail

`func (o *RegisterUserAccount200Response) GetObfuscatedEmail() string`

GetObfuscatedEmail returns the ObfuscatedEmail field if non-nil, zero value otherwise.

### GetObfuscatedEmailOk

`func (o *RegisterUserAccount200Response) GetObfuscatedEmailOk() (*string, bool)`

GetObfuscatedEmailOk returns a tuple with the ObfuscatedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObfuscatedEmail

`func (o *RegisterUserAccount200Response) SetObfuscatedEmail(v string)`

SetObfuscatedEmail sets ObfuscatedEmail field to given value.


### GetObfuscatedPendingEmail

`func (o *RegisterUserAccount200Response) GetObfuscatedPendingEmail() string`

GetObfuscatedPendingEmail returns the ObfuscatedPendingEmail field if non-nil, zero value otherwise.

### GetObfuscatedPendingEmailOk

`func (o *RegisterUserAccount200Response) GetObfuscatedPendingEmailOk() (*string, bool)`

GetObfuscatedPendingEmailOk returns a tuple with the ObfuscatedPendingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObfuscatedPendingEmail

`func (o *RegisterUserAccount200Response) SetObfuscatedPendingEmail(v string)`

SetObfuscatedPendingEmail sets ObfuscatedPendingEmail field to given value.


### GetOculusId

`func (o *RegisterUserAccount200Response) GetOculusId() string`

GetOculusId returns the OculusId field if non-nil, zero value otherwise.

### GetOculusIdOk

`func (o *RegisterUserAccount200Response) GetOculusIdOk() (*string, bool)`

GetOculusIdOk returns a tuple with the OculusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOculusId

`func (o *RegisterUserAccount200Response) SetOculusId(v string)`

SetOculusId sets OculusId field to given value.


### GetOfflineFriends

`func (o *RegisterUserAccount200Response) GetOfflineFriends() []string`

GetOfflineFriends returns the OfflineFriends field if non-nil, zero value otherwise.

### GetOfflineFriendsOk

`func (o *RegisterUserAccount200Response) GetOfflineFriendsOk() (*[]string, bool)`

GetOfflineFriendsOk returns a tuple with the OfflineFriends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflineFriends

`func (o *RegisterUserAccount200Response) SetOfflineFriends(v []string)`

SetOfflineFriends sets OfflineFriends field to given value.

### HasOfflineFriends

`func (o *RegisterUserAccount200Response) HasOfflineFriends() bool`

HasOfflineFriends returns a boolean if a field has been set.

### GetOnlineFriends

`func (o *RegisterUserAccount200Response) GetOnlineFriends() []string`

GetOnlineFriends returns the OnlineFriends field if non-nil, zero value otherwise.

### GetOnlineFriendsOk

`func (o *RegisterUserAccount200Response) GetOnlineFriendsOk() (*[]string, bool)`

GetOnlineFriendsOk returns a tuple with the OnlineFriends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineFriends

`func (o *RegisterUserAccount200Response) SetOnlineFriends(v []string)`

SetOnlineFriends sets OnlineFriends field to given value.

### HasOnlineFriends

`func (o *RegisterUserAccount200Response) HasOnlineFriends() bool`

HasOnlineFriends returns a boolean if a field has been set.

### GetPastDisplayNames

`func (o *RegisterUserAccount200Response) GetPastDisplayNames() []PastDisplayName`

GetPastDisplayNames returns the PastDisplayNames field if non-nil, zero value otherwise.

### GetPastDisplayNamesOk

`func (o *RegisterUserAccount200Response) GetPastDisplayNamesOk() (*[]PastDisplayName, bool)`

GetPastDisplayNamesOk returns a tuple with the PastDisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastDisplayNames

`func (o *RegisterUserAccount200Response) SetPastDisplayNames(v []PastDisplayName)`

SetPastDisplayNames sets PastDisplayNames field to given value.


### GetPicoId

`func (o *RegisterUserAccount200Response) GetPicoId() string`

GetPicoId returns the PicoId field if non-nil, zero value otherwise.

### GetPicoIdOk

`func (o *RegisterUserAccount200Response) GetPicoIdOk() (*string, bool)`

GetPicoIdOk returns a tuple with the PicoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicoId

`func (o *RegisterUserAccount200Response) SetPicoId(v string)`

SetPicoId sets PicoId field to given value.

### HasPicoId

`func (o *RegisterUserAccount200Response) HasPicoId() bool`

HasPicoId returns a boolean if a field has been set.

### GetPlatformHistory

`func (o *RegisterUserAccount200Response) GetPlatformHistory() []CurrentUserPlatformHistoryInner`

GetPlatformHistory returns the PlatformHistory field if non-nil, zero value otherwise.

### GetPlatformHistoryOk

`func (o *RegisterUserAccount200Response) GetPlatformHistoryOk() (*[]CurrentUserPlatformHistoryInner, bool)`

GetPlatformHistoryOk returns a tuple with the PlatformHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformHistory

`func (o *RegisterUserAccount200Response) SetPlatformHistory(v []CurrentUserPlatformHistoryInner)`

SetPlatformHistory sets PlatformHistory field to given value.

### HasPlatformHistory

`func (o *RegisterUserAccount200Response) HasPlatformHistory() bool`

HasPlatformHistory returns a boolean if a field has been set.

### GetPresence

`func (o *RegisterUserAccount200Response) GetPresence() CurrentUserPresence`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *RegisterUserAccount200Response) GetPresenceOk() (*CurrentUserPresence, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *RegisterUserAccount200Response) SetPresence(v CurrentUserPresence)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *RegisterUserAccount200Response) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetProfilePicOverride

`func (o *RegisterUserAccount200Response) GetProfilePicOverride() string`

GetProfilePicOverride returns the ProfilePicOverride field if non-nil, zero value otherwise.

### GetProfilePicOverrideOk

`func (o *RegisterUserAccount200Response) GetProfilePicOverrideOk() (*string, bool)`

GetProfilePicOverrideOk returns a tuple with the ProfilePicOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePicOverride

`func (o *RegisterUserAccount200Response) SetProfilePicOverride(v string)`

SetProfilePicOverride sets ProfilePicOverride field to given value.


### GetProfilePicOverrideThumbnail

`func (o *RegisterUserAccount200Response) GetProfilePicOverrideThumbnail() string`

GetProfilePicOverrideThumbnail returns the ProfilePicOverrideThumbnail field if non-nil, zero value otherwise.

### GetProfilePicOverrideThumbnailOk

`func (o *RegisterUserAccount200Response) GetProfilePicOverrideThumbnailOk() (*string, bool)`

GetProfilePicOverrideThumbnailOk returns a tuple with the ProfilePicOverrideThumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePicOverrideThumbnail

`func (o *RegisterUserAccount200Response) SetProfilePicOverrideThumbnail(v string)`

SetProfilePicOverrideThumbnail sets ProfilePicOverrideThumbnail field to given value.


### GetPronouns

`func (o *RegisterUserAccount200Response) GetPronouns() string`

GetPronouns returns the Pronouns field if non-nil, zero value otherwise.

### GetPronounsOk

`func (o *RegisterUserAccount200Response) GetPronounsOk() (*string, bool)`

GetPronounsOk returns a tuple with the Pronouns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPronouns

`func (o *RegisterUserAccount200Response) SetPronouns(v string)`

SetPronouns sets Pronouns field to given value.


### GetPronounsHistory

`func (o *RegisterUserAccount200Response) GetPronounsHistory() []string`

GetPronounsHistory returns the PronounsHistory field if non-nil, zero value otherwise.

### GetPronounsHistoryOk

`func (o *RegisterUserAccount200Response) GetPronounsHistoryOk() (*[]string, bool)`

GetPronounsHistoryOk returns a tuple with the PronounsHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPronounsHistory

`func (o *RegisterUserAccount200Response) SetPronounsHistory(v []string)`

SetPronounsHistory sets PronounsHistory field to given value.


### GetQueuedInstance

`func (o *RegisterUserAccount200Response) GetQueuedInstance() string`

GetQueuedInstance returns the QueuedInstance field if non-nil, zero value otherwise.

### GetQueuedInstanceOk

`func (o *RegisterUserAccount200Response) GetQueuedInstanceOk() (*string, bool)`

GetQueuedInstanceOk returns a tuple with the QueuedInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedInstance

`func (o *RegisterUserAccount200Response) SetQueuedInstance(v string)`

SetQueuedInstance sets QueuedInstance field to given value.

### HasQueuedInstance

`func (o *RegisterUserAccount200Response) HasQueuedInstance() bool`

HasQueuedInstance returns a boolean if a field has been set.

### SetQueuedInstanceNil

`func (o *RegisterUserAccount200Response) SetQueuedInstanceNil(b bool)`

 SetQueuedInstanceNil sets the value for QueuedInstance to be an explicit nil

### UnsetQueuedInstance
`func (o *RegisterUserAccount200Response) UnsetQueuedInstance()`

UnsetQueuedInstance ensures that no value is present for QueuedInstance, not even an explicit nil
### GetReceiveMobileInvitations

`func (o *RegisterUserAccount200Response) GetReceiveMobileInvitations() bool`

GetReceiveMobileInvitations returns the ReceiveMobileInvitations field if non-nil, zero value otherwise.

### GetReceiveMobileInvitationsOk

`func (o *RegisterUserAccount200Response) GetReceiveMobileInvitationsOk() (*bool, bool)`

GetReceiveMobileInvitationsOk returns a tuple with the ReceiveMobileInvitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveMobileInvitations

`func (o *RegisterUserAccount200Response) SetReceiveMobileInvitations(v bool)`

SetReceiveMobileInvitations sets ReceiveMobileInvitations field to given value.

### HasReceiveMobileInvitations

`func (o *RegisterUserAccount200Response) HasReceiveMobileInvitations() bool`

HasReceiveMobileInvitations returns a boolean if a field has been set.

### GetState

`func (o *RegisterUserAccount200Response) GetState() UserState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RegisterUserAccount200Response) GetStateOk() (*UserState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RegisterUserAccount200Response) SetState(v UserState)`

SetState sets State field to given value.


### GetStatus

`func (o *RegisterUserAccount200Response) GetStatus() UserStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RegisterUserAccount200Response) GetStatusOk() (*UserStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RegisterUserAccount200Response) SetStatus(v UserStatus)`

SetStatus sets Status field to given value.


### GetStatusDescription

`func (o *RegisterUserAccount200Response) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *RegisterUserAccount200Response) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *RegisterUserAccount200Response) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.


### GetStatusFirstTime

`func (o *RegisterUserAccount200Response) GetStatusFirstTime() bool`

GetStatusFirstTime returns the StatusFirstTime field if non-nil, zero value otherwise.

### GetStatusFirstTimeOk

`func (o *RegisterUserAccount200Response) GetStatusFirstTimeOk() (*bool, bool)`

GetStatusFirstTimeOk returns a tuple with the StatusFirstTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusFirstTime

`func (o *RegisterUserAccount200Response) SetStatusFirstTime(v bool)`

SetStatusFirstTime sets StatusFirstTime field to given value.


### GetStatusHistory

`func (o *RegisterUserAccount200Response) GetStatusHistory() []string`

GetStatusHistory returns the StatusHistory field if non-nil, zero value otherwise.

### GetStatusHistoryOk

`func (o *RegisterUserAccount200Response) GetStatusHistoryOk() (*[]string, bool)`

GetStatusHistoryOk returns a tuple with the StatusHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusHistory

`func (o *RegisterUserAccount200Response) SetStatusHistory(v []string)`

SetStatusHistory sets StatusHistory field to given value.


### GetSteamDetails

`func (o *RegisterUserAccount200Response) GetSteamDetails() map[string]interface{}`

GetSteamDetails returns the SteamDetails field if non-nil, zero value otherwise.

### GetSteamDetailsOk

`func (o *RegisterUserAccount200Response) GetSteamDetailsOk() (*map[string]interface{}, bool)`

GetSteamDetailsOk returns a tuple with the SteamDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamDetails

`func (o *RegisterUserAccount200Response) SetSteamDetails(v map[string]interface{})`

SetSteamDetails sets SteamDetails field to given value.


### GetSteamId

`func (o *RegisterUserAccount200Response) GetSteamId() string`

GetSteamId returns the SteamId field if non-nil, zero value otherwise.

### GetSteamIdOk

`func (o *RegisterUserAccount200Response) GetSteamIdOk() (*string, bool)`

GetSteamIdOk returns a tuple with the SteamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamId

`func (o *RegisterUserAccount200Response) SetSteamId(v string)`

SetSteamId sets SteamId field to given value.


### GetTags

`func (o *RegisterUserAccount200Response) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RegisterUserAccount200Response) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RegisterUserAccount200Response) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetTwoFactorAuthEnabled

`func (o *RegisterUserAccount200Response) GetTwoFactorAuthEnabled() bool`

GetTwoFactorAuthEnabled returns the TwoFactorAuthEnabled field if non-nil, zero value otherwise.

### GetTwoFactorAuthEnabledOk

`func (o *RegisterUserAccount200Response) GetTwoFactorAuthEnabledOk() (*bool, bool)`

GetTwoFactorAuthEnabledOk returns a tuple with the TwoFactorAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuthEnabled

`func (o *RegisterUserAccount200Response) SetTwoFactorAuthEnabled(v bool)`

SetTwoFactorAuthEnabled sets TwoFactorAuthEnabled field to given value.


### GetTwoFactorAuthEnabledDate

`func (o *RegisterUserAccount200Response) GetTwoFactorAuthEnabledDate() time.Time`

GetTwoFactorAuthEnabledDate returns the TwoFactorAuthEnabledDate field if non-nil, zero value otherwise.

### GetTwoFactorAuthEnabledDateOk

`func (o *RegisterUserAccount200Response) GetTwoFactorAuthEnabledDateOk() (*time.Time, bool)`

GetTwoFactorAuthEnabledDateOk returns a tuple with the TwoFactorAuthEnabledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuthEnabledDate

`func (o *RegisterUserAccount200Response) SetTwoFactorAuthEnabledDate(v time.Time)`

SetTwoFactorAuthEnabledDate sets TwoFactorAuthEnabledDate field to given value.

### HasTwoFactorAuthEnabledDate

`func (o *RegisterUserAccount200Response) HasTwoFactorAuthEnabledDate() bool`

HasTwoFactorAuthEnabledDate returns a boolean if a field has been set.

### SetTwoFactorAuthEnabledDateNil

`func (o *RegisterUserAccount200Response) SetTwoFactorAuthEnabledDateNil(b bool)`

 SetTwoFactorAuthEnabledDateNil sets the value for TwoFactorAuthEnabledDate to be an explicit nil

### UnsetTwoFactorAuthEnabledDate
`func (o *RegisterUserAccount200Response) UnsetTwoFactorAuthEnabledDate()`

UnsetTwoFactorAuthEnabledDate ensures that no value is present for TwoFactorAuthEnabledDate, not even an explicit nil
### GetUnsubscribe

`func (o *RegisterUserAccount200Response) GetUnsubscribe() bool`

GetUnsubscribe returns the Unsubscribe field if non-nil, zero value otherwise.

### GetUnsubscribeOk

`func (o *RegisterUserAccount200Response) GetUnsubscribeOk() (*bool, bool)`

GetUnsubscribeOk returns a tuple with the Unsubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribe

`func (o *RegisterUserAccount200Response) SetUnsubscribe(v bool)`

SetUnsubscribe sets Unsubscribe field to given value.


### GetUpdatedAt

`func (o *RegisterUserAccount200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RegisterUserAccount200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RegisterUserAccount200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RegisterUserAccount200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserIcon

`func (o *RegisterUserAccount200Response) GetUserIcon() string`

GetUserIcon returns the UserIcon field if non-nil, zero value otherwise.

### GetUserIconOk

`func (o *RegisterUserAccount200Response) GetUserIconOk() (*string, bool)`

GetUserIconOk returns a tuple with the UserIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIcon

`func (o *RegisterUserAccount200Response) SetUserIcon(v string)`

SetUserIcon sets UserIcon field to given value.


### GetUserLanguage

`func (o *RegisterUserAccount200Response) GetUserLanguage() string`

GetUserLanguage returns the UserLanguage field if non-nil, zero value otherwise.

### GetUserLanguageOk

`func (o *RegisterUserAccount200Response) GetUserLanguageOk() (*string, bool)`

GetUserLanguageOk returns a tuple with the UserLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLanguage

`func (o *RegisterUserAccount200Response) SetUserLanguage(v string)`

SetUserLanguage sets UserLanguage field to given value.

### HasUserLanguage

`func (o *RegisterUserAccount200Response) HasUserLanguage() bool`

HasUserLanguage returns a boolean if a field has been set.

### SetUserLanguageNil

`func (o *RegisterUserAccount200Response) SetUserLanguageNil(b bool)`

 SetUserLanguageNil sets the value for UserLanguage to be an explicit nil

### UnsetUserLanguage
`func (o *RegisterUserAccount200Response) UnsetUserLanguage()`

UnsetUserLanguage ensures that no value is present for UserLanguage, not even an explicit nil
### GetUserLanguageCode

`func (o *RegisterUserAccount200Response) GetUserLanguageCode() string`

GetUserLanguageCode returns the UserLanguageCode field if non-nil, zero value otherwise.

### GetUserLanguageCodeOk

`func (o *RegisterUserAccount200Response) GetUserLanguageCodeOk() (*string, bool)`

GetUserLanguageCodeOk returns a tuple with the UserLanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLanguageCode

`func (o *RegisterUserAccount200Response) SetUserLanguageCode(v string)`

SetUserLanguageCode sets UserLanguageCode field to given value.

### HasUserLanguageCode

`func (o *RegisterUserAccount200Response) HasUserLanguageCode() bool`

HasUserLanguageCode returns a boolean if a field has been set.

### SetUserLanguageCodeNil

`func (o *RegisterUserAccount200Response) SetUserLanguageCodeNil(b bool)`

 SetUserLanguageCodeNil sets the value for UserLanguageCode to be an explicit nil

### UnsetUserLanguageCode
`func (o *RegisterUserAccount200Response) UnsetUserLanguageCode()`

UnsetUserLanguageCode ensures that no value is present for UserLanguageCode, not even an explicit nil
### GetUsername

`func (o *RegisterUserAccount200Response) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RegisterUserAccount200Response) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RegisterUserAccount200Response) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *RegisterUserAccount200Response) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetUsesGeneratedPassword

`func (o *RegisterUserAccount200Response) GetUsesGeneratedPassword() bool`

GetUsesGeneratedPassword returns the UsesGeneratedPassword field if non-nil, zero value otherwise.

### GetUsesGeneratedPasswordOk

`func (o *RegisterUserAccount200Response) GetUsesGeneratedPasswordOk() (*bool, bool)`

GetUsesGeneratedPasswordOk returns a tuple with the UsesGeneratedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesGeneratedPassword

`func (o *RegisterUserAccount200Response) SetUsesGeneratedPassword(v bool)`

SetUsesGeneratedPassword sets UsesGeneratedPassword field to given value.


### GetViveId

`func (o *RegisterUserAccount200Response) GetViveId() string`

GetViveId returns the ViveId field if non-nil, zero value otherwise.

### GetViveIdOk

`func (o *RegisterUserAccount200Response) GetViveIdOk() (*string, bool)`

GetViveIdOk returns a tuple with the ViveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViveId

`func (o *RegisterUserAccount200Response) SetViveId(v string)`

SetViveId sets ViveId field to given value.

### HasViveId

`func (o *RegisterUserAccount200Response) HasViveId() bool`

HasViveId returns a boolean if a field has been set.

### GetRequiresTwoFactorAuth

`func (o *RegisterUserAccount200Response) GetRequiresTwoFactorAuth() []TwoFactorAuthType`

GetRequiresTwoFactorAuth returns the RequiresTwoFactorAuth field if non-nil, zero value otherwise.

### GetRequiresTwoFactorAuthOk

`func (o *RegisterUserAccount200Response) GetRequiresTwoFactorAuthOk() (*[]TwoFactorAuthType, bool)`

GetRequiresTwoFactorAuthOk returns a tuple with the RequiresTwoFactorAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresTwoFactorAuth

`func (o *RegisterUserAccount200Response) SetRequiresTwoFactorAuth(v []TwoFactorAuthType)`

SetRequiresTwoFactorAuth sets RequiresTwoFactorAuth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


