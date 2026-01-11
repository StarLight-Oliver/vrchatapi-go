/*
VRChat API Documentation

![VRChat API Banner](https://vrchatapi.github.io/assets/img/api_banner_1500x400.png)  # Welcome to the VRChat API  Before we begin, we would like to state this is a **COMMUNITY DRIVEN PROJECT**. This means that everything you read on here was written by the community itself and is **not** officially supported by VRChat. The documentation is provided \"AS IS\", and any action you take towards VRChat is completely your own responsibility.  The documentation and additional libraries SHALL ONLY be used for applications interacting with VRChat's API in accordance with their [Terms of Service](https://hello.vrchat.com/legal) and [Community Guidelines](https://hello.vrchat.com/community-guidelines), and MUST NOT be used for modifying the client, \"avatar ripping\", or other illegal activities. Malicious usage or spamming the API may result in account termination. Certain parts of the API are also more sensitive than others, for example moderation, so please tread extra carefully and read the warnings when present.  ![Tupper Policy on API](https://i.imgur.com/yLlW7Ok.png)  Finally, use of the API using applications other than the approved methods (website, VRChat application, Unity SDK) is not officially supported. VRChat provides no guarantee or support for external applications using the API. Access to API endpoints may break **at any time, without notice**. Therefore, please **do not ping** VRChat Staff in the VRChat Discord if you are having API problems, as they do not provide API support. We will make a best effort in keeping this documentation and associated language libraries up to date, but things might be outdated or missing. If you find that something is no longer valid, please contact us on Discord or [create an issue](https://github.com/vrchatapi/specification/issues) and tell us so we can fix it.  # Getting Started  The VRChat API can be used to programmatically retrieve or update information regarding your profile, friends, avatars, worlds and more. The API consists of two parts, \"Photon\" which is only used in-game, and the \"Web API\" which is used by both the game and the website. This documentation focuses only on the Web API.  The API is designed around the REST ideology, providing semi-simple and usually predictable URIs to access and modify objects. Requests support standard HTTP methods like GET, PUT, POST, and DELETE and standard status codes. Response bodies are always UTF-8 encoded JSON objects, unless explicitly documented otherwise.  <div class=\"callout callout-error\">   <strong>🛑 Warning! Do not touch Photon!</strong><br>   Photon is only used by the in-game client and should <b>not</b> be touched. Doing so may result in permanent account termination. </div>  <div class=\"callout callout-info\">   <strong>ℹ️ Authentication</strong><br>   Read <a href=\"#tag--authentication\">Authentication</a> for how to log in. </div>  # Using the API  For simply exploring what the API can do it is strongly recommended to download [Insomnia](https://insomnia.rest/download), a free and open-source API client that's great for sending requests to the API in an orderly fashion. Insomnia allows you to send data in the format that's required for VRChat's API. It is also possible to try out the API in your browser, by first logging in at [vrchat.com/home](https://vrchat.com/home/) and then going to [vrchat.com/api/1/auth/user](https://vrchat.com/api/1/auth/user), but the information will be much harder to work with.  For more permanent operation such as software development it is instead recommended to use one of the existing language SDKs. This community project maintains API libraries in several languages, which allows you to interact with the API with simple function calls rather than having to implement the HTTP protocol yourself. Most of these libraries are automatically generated from the API specification, sometimes with additional helpful wrapper code to make usage easier. This allows them to be almost automatically updated and expanded upon as soon as a new feature is introduced in the specification itself. The libraries can be found on [GitHub](https://github.com/vrchatapi) or following:  * [NodeJS (JavaScript)](https://www.npmjs.com/package/vrchat) * [Dart](https://pub.dev/packages/vrchat_dart) * [Rust](https://crates.io/crates/vrchatapi) * [C#](https://github.com/vrchatapi/vrchatapi-csharp) * [Python](https://github.com/vrchatapi/vrchatapi-python)  # Pagination  Most endpoints enforce pagination, meaning they will only return 10 entries by default, and never more than 100.<br> Using both the limit and offset parameters allows you to easily paginate through a large number of objects.  | Query Parameter | Type | Description | | ----------|--|------- | | `n` | integer  | The number of objects to return. This value often defaults to 10. Highest limit is always 100.| | `offset` | integer  | A zero-based offset from the default object sorting.|  If a request returns fewer objects than the `limit` parameter, there are no more items available to return.  # Contribution  Do you want to get involved in the documentation effort? Do you want to help improve one of the language API libraries? This project is an [OPEN Open Source Project](https://openopensource.org)! This means that individuals making significant and valuable contributions are given commit-access to the project. It also means we are very open and welcoming of new people making contributions, unlike some more guarded open-source projects.  [![Discord](https://img.shields.io/static/v1?label=vrchatapi&message=discord&color=blueviolet&style=for-the-badge)](https://discord.gg/qjZE9C9fkB)

API version: 1.20.7
Contact: vrchatapi.lpv0t@aries.fyi
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vrchatapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the CurrentUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CurrentUser{}

// CurrentUser struct for CurrentUser
type CurrentUser struct {
	AcceptedPrivacyVersion *int32         `json:"acceptedPrivacyVersion,omitempty"`
	AcceptedTOSVersion     int32          `json:"acceptedTOSVersion"`
	AccountDeletionDate    NullableString `json:"accountDeletionDate,omitempty"`
	//
	AccountDeletionLog []AccountDeletionLog `json:"accountDeletionLog,omitempty"`
	//
	ActiveFriends         []string              `json:"activeFriends,omitempty"`
	AgeVerificationStatus AgeVerificationStatus `json:"ageVerificationStatus"`
	// `true` if, user is age verified (not 18+).
	AgeVerified        bool `json:"ageVerified"`
	AllowAvatarCopying bool `json:"allowAvatarCopying"`
	// The auth token for NEWLY REGISTERED ACCOUNTS ONLY (/auth/register)
	AuthToken *string `json:"authToken,omitempty"`
	//
	Badges []Badge `json:"badges,omitempty"`
	Bio    string  `json:"bio"`
	//
	BioLinks []string `json:"bioLinks"`
	// These tags begin with `content_` and control content gating
	ContentFilters []string `json:"contentFilters,omitempty"`
	CurrentAvatar  string   `json:"currentAvatar"`
	// When profilePicOverride is not empty, use it instead.
	CurrentAvatarImageUrl string   `json:"currentAvatarImageUrl"`
	CurrentAvatarTags     []string `json:"currentAvatarTags"`
	// When profilePicOverride is not empty, use it instead.
	CurrentAvatarThumbnailImageUrl string          `json:"currentAvatarThumbnailImageUrl"`
	DateJoined                     string          `json:"date_joined"`
	DeveloperType                  DeveloperType   `json:"developerType"`
	DiscordDetails                 *DiscordDetails `json:"discordDetails,omitempty"`
	// https://discord.com/developers/docs/reference#snowflakes
	DiscordId      *string `json:"discordId,omitempty"`
	DisplayName    string  `json:"displayName"`
	EmailVerified  bool    `json:"emailVerified"`
	FallbackAvatar *string `json:"fallbackAvatar,omitempty"`
	// Always empty array.
	// Deprecated
	FriendGroupNames          []string               `json:"friendGroupNames"`
	FriendKey                 string                 `json:"friendKey"`
	Friends                   []string               `json:"friends"`
	GoogleDetails             map[string]interface{} `json:"googleDetails,omitempty"`
	GoogleId                  *string                `json:"googleId,omitempty"`
	HasBirthday               bool                   `json:"hasBirthday"`
	HasEmail                  bool                   `json:"hasEmail"`
	HasLoggedInFromClient     bool                   `json:"hasLoggedInFromClient"`
	HasPendingEmail           bool                   `json:"hasPendingEmail"`
	HideContentFilterSettings *bool                  `json:"hideContentFilterSettings,omitempty"`
	// WorldID be \"offline\" on User profiles if you are not friends with that user.
	HomeLocation string `json:"homeLocation"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	Id               string       `json:"id"`
	IsAdult          bool         `json:"isAdult"`
	IsBoopingEnabled *bool        `json:"isBoopingEnabled,omitempty"`
	IsFriend         bool         `json:"isFriend"`
	LastActivity     *time.Time   `json:"last_activity,omitempty"`
	LastLogin        time.Time    `json:"last_login"`
	LastMobile       NullableTime `json:"last_mobile"`
	// This is normally `android`, `ios`, `standalonewindows`, `web`, or the empty value ``, but also supposedly can be any random Unity version such as `2019.2.4-801-Release` or `2019.2.2-772-Release` or even `unknownplatform`.
	LastPlatform           string   `json:"last_platform"`
	ObfuscatedEmail        string   `json:"obfuscatedEmail"`
	ObfuscatedPendingEmail string   `json:"obfuscatedPendingEmail"`
	OculusId               string   `json:"oculusId"`
	OfflineFriends         []string `json:"offlineFriends,omitempty"`
	OnlineFriends          []string `json:"onlineFriends,omitempty"`
	//
	PastDisplayNames            []PastDisplayName                 `json:"pastDisplayNames"`
	PicoId                      *string                           `json:"picoId,omitempty"`
	PlatformHistory             []CurrentUserPlatformHistoryInner `json:"platform_history,omitempty"`
	Presence                    *CurrentUserPresence              `json:"presence,omitempty"`
	ProfilePicOverride          string                            `json:"profilePicOverride"`
	ProfilePicOverrideThumbnail string                            `json:"profilePicOverrideThumbnail"`
	Pronouns                    string                            `json:"pronouns"`
	PronounsHistory             []string                          `json:"pronounsHistory"`
	QueuedInstance              NullableString                    `json:"queuedInstance,omitempty"`
	ReceiveMobileInvitations    *bool                             `json:"receiveMobileInvitations,omitempty"`
	State                       UserState                         `json:"state"`
	Status                      UserStatus                        `json:"status"`
	StatusDescription           string                            `json:"statusDescription"`
	StatusFirstTime             bool                              `json:"statusFirstTime"`
	StatusHistory               []string                          `json:"statusHistory"`
	SteamDetails                map[string]interface{}            `json:"steamDetails"`
	SteamId                     string                            `json:"steamId"`
	Tags                        []string                          `json:"tags"`
	TwoFactorAuthEnabled        bool                              `json:"twoFactorAuthEnabled"`
	TwoFactorAuthEnabledDate    NullableTime                      `json:"twoFactorAuthEnabledDate,omitempty"`
	Unsubscribe                 bool                              `json:"unsubscribe"`
	UpdatedAt                   *time.Time                        `json:"updated_at,omitempty"`
	UserIcon                    string                            `json:"userIcon"`
	UserLanguage                NullableString                    `json:"userLanguage,omitempty"`
	UserLanguageCode            NullableString                    `json:"userLanguageCode,omitempty"`
	// -| **DEPRECATED:** VRChat API no longer return usernames of other users. [See issue by Tupper for more information](https://github.com/pypy-vrc/VRCX/issues/429).
	// Deprecated
	Username              *string `json:"username,omitempty"`
	UsesGeneratedPassword bool    `json:"usesGeneratedPassword"`
	ViveId                *string `json:"viveId,omitempty"`
}

type _CurrentUser CurrentUser

// NewCurrentUser instantiates a new CurrentUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrentUser(acceptedTOSVersion int32, ageVerificationStatus AgeVerificationStatus, ageVerified bool, allowAvatarCopying bool, bio string, bioLinks []string, currentAvatar string, currentAvatarImageUrl string, currentAvatarTags []string, currentAvatarThumbnailImageUrl string, dateJoined string, developerType DeveloperType, displayName string, emailVerified bool, friendGroupNames []string, friendKey string, friends []string, hasBirthday bool, hasEmail bool, hasLoggedInFromClient bool, hasPendingEmail bool, homeLocation string, id string, isAdult bool, isFriend bool, lastLogin time.Time, lastMobile NullableTime, lastPlatform string, obfuscatedEmail string, obfuscatedPendingEmail string, oculusId string, pastDisplayNames []PastDisplayName, profilePicOverride string, profilePicOverrideThumbnail string, pronouns string, pronounsHistory []string, state UserState, status UserStatus, statusDescription string, statusFirstTime bool, statusHistory []string, steamDetails map[string]interface{}, steamId string, tags []string, twoFactorAuthEnabled bool, unsubscribe bool, userIcon string, usesGeneratedPassword bool) *CurrentUser {
	this := CurrentUser{}
	this.AcceptedTOSVersion = acceptedTOSVersion
	this.AgeVerificationStatus = ageVerificationStatus
	this.AgeVerified = ageVerified
	this.AllowAvatarCopying = allowAvatarCopying
	this.Bio = bio
	this.BioLinks = bioLinks
	this.CurrentAvatar = currentAvatar
	this.CurrentAvatarImageUrl = currentAvatarImageUrl
	this.CurrentAvatarTags = currentAvatarTags
	this.CurrentAvatarThumbnailImageUrl = currentAvatarThumbnailImageUrl
	this.DateJoined = dateJoined
	this.DeveloperType = developerType
	this.DisplayName = displayName
	this.EmailVerified = emailVerified
	this.FriendGroupNames = friendGroupNames
	this.FriendKey = friendKey
	this.Friends = friends
	this.HasBirthday = hasBirthday
	this.HasEmail = hasEmail
	this.HasLoggedInFromClient = hasLoggedInFromClient
	this.HasPendingEmail = hasPendingEmail
	this.HomeLocation = homeLocation
	this.Id = id
	this.IsAdult = isAdult
	var isBoopingEnabled bool = true
	this.IsBoopingEnabled = &isBoopingEnabled
	this.IsFriend = isFriend
	this.LastLogin = lastLogin
	this.LastMobile = lastMobile
	this.LastPlatform = lastPlatform
	this.ObfuscatedEmail = obfuscatedEmail
	this.ObfuscatedPendingEmail = obfuscatedPendingEmail
	this.OculusId = oculusId
	this.PastDisplayNames = pastDisplayNames
	this.ProfilePicOverride = profilePicOverride
	this.ProfilePicOverrideThumbnail = profilePicOverrideThumbnail
	this.Pronouns = pronouns
	this.PronounsHistory = pronounsHistory
	this.State = state
	this.Status = status
	this.StatusDescription = statusDescription
	this.StatusFirstTime = statusFirstTime
	this.StatusHistory = statusHistory
	this.SteamDetails = steamDetails
	this.SteamId = steamId
	this.Tags = tags
	this.TwoFactorAuthEnabled = twoFactorAuthEnabled
	this.Unsubscribe = unsubscribe
	this.UserIcon = userIcon
	this.UsesGeneratedPassword = usesGeneratedPassword
	return &this
}

// NewCurrentUserWithDefaults instantiates a new CurrentUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrentUserWithDefaults() *CurrentUser {
	this := CurrentUser{}
	var developerType DeveloperType = DeveloperType_NONE
	this.DeveloperType = developerType
	var isBoopingEnabled bool = true
	this.IsBoopingEnabled = &isBoopingEnabled
	var isFriend bool = false
	this.IsFriend = isFriend
	var state UserState = UserState_OFFLINE
	this.State = state
	var status UserStatus = UserStatus_OFFLINE
	this.Status = status
	return &this
}

// GetAcceptedPrivacyVersion returns the AcceptedPrivacyVersion field value if set, zero value otherwise.
func (o *CurrentUser) GetAcceptedPrivacyVersion() int32 {
	if o == nil || IsNil(o.AcceptedPrivacyVersion) {
		var ret int32
		return ret
	}
	return *o.AcceptedPrivacyVersion
}

// GetAcceptedPrivacyVersionOk returns a tuple with the AcceptedPrivacyVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetAcceptedPrivacyVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.AcceptedPrivacyVersion) {
		return nil, false
	}
	return o.AcceptedPrivacyVersion, true
}

// HasAcceptedPrivacyVersion returns a boolean if a field has been set.
func (o *CurrentUser) HasAcceptedPrivacyVersion() bool {
	if o != nil && !IsNil(o.AcceptedPrivacyVersion) {
		return true
	}

	return false
}

// SetAcceptedPrivacyVersion gets a reference to the given int32 and assigns it to the AcceptedPrivacyVersion field.
func (o *CurrentUser) SetAcceptedPrivacyVersion(v int32) {
	o.AcceptedPrivacyVersion = &v
}

// GetAcceptedTOSVersion returns the AcceptedTOSVersion field value
func (o *CurrentUser) GetAcceptedTOSVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AcceptedTOSVersion
}

// GetAcceptedTOSVersionOk returns a tuple with the AcceptedTOSVersion field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetAcceptedTOSVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcceptedTOSVersion, true
}

// SetAcceptedTOSVersion sets field value
func (o *CurrentUser) SetAcceptedTOSVersion(v int32) {
	o.AcceptedTOSVersion = v
}

// GetAccountDeletionDate returns the AccountDeletionDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CurrentUser) GetAccountDeletionDate() string {
	if o == nil || IsNil(o.AccountDeletionDate.Get()) {
		var ret string
		return ret
	}
	return *o.AccountDeletionDate.Get()
}

// GetAccountDeletionDateOk returns a tuple with the AccountDeletionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CurrentUser) GetAccountDeletionDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountDeletionDate.Get(), o.AccountDeletionDate.IsSet()
}

// HasAccountDeletionDate returns a boolean if a field has been set.
func (o *CurrentUser) HasAccountDeletionDate() bool {
	if o != nil && o.AccountDeletionDate.IsSet() {
		return true
	}

	return false
}

// SetAccountDeletionDate gets a reference to the given NullableString and assigns it to the AccountDeletionDate field.
func (o *CurrentUser) SetAccountDeletionDate(v string) {
	o.AccountDeletionDate.Set(&v)
}

// SetAccountDeletionDateNil sets the value for AccountDeletionDate to be an explicit nil
func (o *CurrentUser) SetAccountDeletionDateNil() {
	o.AccountDeletionDate.Set(nil)
}

// UnsetAccountDeletionDate ensures that no value is present for AccountDeletionDate, not even an explicit nil
func (o *CurrentUser) UnsetAccountDeletionDate() {
	o.AccountDeletionDate.Unset()
}

// GetAccountDeletionLog returns the AccountDeletionLog field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CurrentUser) GetAccountDeletionLog() []AccountDeletionLog {
	if o == nil {
		var ret []AccountDeletionLog
		return ret
	}
	return o.AccountDeletionLog
}

// GetAccountDeletionLogOk returns a tuple with the AccountDeletionLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CurrentUser) GetAccountDeletionLogOk() ([]AccountDeletionLog, bool) {
	if o == nil || IsNil(o.AccountDeletionLog) {
		return nil, false
	}
	return o.AccountDeletionLog, true
}

// HasAccountDeletionLog returns a boolean if a field has been set.
func (o *CurrentUser) HasAccountDeletionLog() bool {
	if o != nil && !IsNil(o.AccountDeletionLog) {
		return true
	}

	return false
}

// SetAccountDeletionLog gets a reference to the given []AccountDeletionLog and assigns it to the AccountDeletionLog field.
func (o *CurrentUser) SetAccountDeletionLog(v []AccountDeletionLog) {
	o.AccountDeletionLog = v
}

// GetActiveFriends returns the ActiveFriends field value if set, zero value otherwise.
func (o *CurrentUser) GetActiveFriends() []string {
	if o == nil || IsNil(o.ActiveFriends) {
		var ret []string
		return ret
	}
	return o.ActiveFriends
}

// GetActiveFriendsOk returns a tuple with the ActiveFriends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetActiveFriendsOk() ([]string, bool) {
	if o == nil || IsNil(o.ActiveFriends) {
		return nil, false
	}
	return o.ActiveFriends, true
}

// HasActiveFriends returns a boolean if a field has been set.
func (o *CurrentUser) HasActiveFriends() bool {
	if o != nil && !IsNil(o.ActiveFriends) {
		return true
	}

	return false
}

// SetActiveFriends gets a reference to the given []string and assigns it to the ActiveFriends field.
func (o *CurrentUser) SetActiveFriends(v []string) {
	o.ActiveFriends = v
}

// GetAgeVerificationStatus returns the AgeVerificationStatus field value
func (o *CurrentUser) GetAgeVerificationStatus() AgeVerificationStatus {
	if o == nil {
		var ret AgeVerificationStatus
		return ret
	}

	return o.AgeVerificationStatus
}

// GetAgeVerificationStatusOk returns a tuple with the AgeVerificationStatus field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetAgeVerificationStatusOk() (*AgeVerificationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgeVerificationStatus, true
}

// SetAgeVerificationStatus sets field value
func (o *CurrentUser) SetAgeVerificationStatus(v AgeVerificationStatus) {
	o.AgeVerificationStatus = v
}

// GetAgeVerified returns the AgeVerified field value
func (o *CurrentUser) GetAgeVerified() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AgeVerified
}

// GetAgeVerifiedOk returns a tuple with the AgeVerified field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetAgeVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgeVerified, true
}

// SetAgeVerified sets field value
func (o *CurrentUser) SetAgeVerified(v bool) {
	o.AgeVerified = v
}

// GetAllowAvatarCopying returns the AllowAvatarCopying field value
func (o *CurrentUser) GetAllowAvatarCopying() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowAvatarCopying
}

// GetAllowAvatarCopyingOk returns a tuple with the AllowAvatarCopying field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetAllowAvatarCopyingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowAvatarCopying, true
}

// SetAllowAvatarCopying sets field value
func (o *CurrentUser) SetAllowAvatarCopying(v bool) {
	o.AllowAvatarCopying = v
}

// GetAuthToken returns the AuthToken field value if set, zero value otherwise.
func (o *CurrentUser) GetAuthToken() string {
	if o == nil || IsNil(o.AuthToken) {
		var ret string
		return ret
	}
	return *o.AuthToken
}

// GetAuthTokenOk returns a tuple with the AuthToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetAuthTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AuthToken) {
		return nil, false
	}
	return o.AuthToken, true
}

// HasAuthToken returns a boolean if a field has been set.
func (o *CurrentUser) HasAuthToken() bool {
	if o != nil && !IsNil(o.AuthToken) {
		return true
	}

	return false
}

// SetAuthToken gets a reference to the given string and assigns it to the AuthToken field.
func (o *CurrentUser) SetAuthToken(v string) {
	o.AuthToken = &v
}

// GetBadges returns the Badges field value if set, zero value otherwise.
func (o *CurrentUser) GetBadges() []Badge {
	if o == nil || IsNil(o.Badges) {
		var ret []Badge
		return ret
	}
	return o.Badges
}

// GetBadgesOk returns a tuple with the Badges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetBadgesOk() ([]Badge, bool) {
	if o == nil || IsNil(o.Badges) {
		return nil, false
	}
	return o.Badges, true
}

// HasBadges returns a boolean if a field has been set.
func (o *CurrentUser) HasBadges() bool {
	if o != nil && !IsNil(o.Badges) {
		return true
	}

	return false
}

// SetBadges gets a reference to the given []Badge and assigns it to the Badges field.
func (o *CurrentUser) SetBadges(v []Badge) {
	o.Badges = v
}

// GetBio returns the Bio field value
func (o *CurrentUser) GetBio() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bio
}

// GetBioOk returns a tuple with the Bio field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetBioOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bio, true
}

// SetBio sets field value
func (o *CurrentUser) SetBio(v string) {
	o.Bio = v
}

// GetBioLinks returns the BioLinks field value
func (o *CurrentUser) GetBioLinks() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BioLinks
}

// GetBioLinksOk returns a tuple with the BioLinks field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetBioLinksOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BioLinks, true
}

// SetBioLinks sets field value
func (o *CurrentUser) SetBioLinks(v []string) {
	o.BioLinks = v
}

// GetContentFilters returns the ContentFilters field value if set, zero value otherwise.
func (o *CurrentUser) GetContentFilters() []string {
	if o == nil || IsNil(o.ContentFilters) {
		var ret []string
		return ret
	}
	return o.ContentFilters
}

// GetContentFiltersOk returns a tuple with the ContentFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetContentFiltersOk() ([]string, bool) {
	if o == nil || IsNil(o.ContentFilters) {
		return nil, false
	}
	return o.ContentFilters, true
}

// HasContentFilters returns a boolean if a field has been set.
func (o *CurrentUser) HasContentFilters() bool {
	if o != nil && !IsNil(o.ContentFilters) {
		return true
	}

	return false
}

// SetContentFilters gets a reference to the given []string and assigns it to the ContentFilters field.
func (o *CurrentUser) SetContentFilters(v []string) {
	o.ContentFilters = v
}

// GetCurrentAvatar returns the CurrentAvatar field value
func (o *CurrentUser) GetCurrentAvatar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentAvatar
}

// GetCurrentAvatarOk returns a tuple with the CurrentAvatar field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetCurrentAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentAvatar, true
}

// SetCurrentAvatar sets field value
func (o *CurrentUser) SetCurrentAvatar(v string) {
	o.CurrentAvatar = v
}

// GetCurrentAvatarImageUrl returns the CurrentAvatarImageUrl field value
func (o *CurrentUser) GetCurrentAvatarImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentAvatarImageUrl
}

// GetCurrentAvatarImageUrlOk returns a tuple with the CurrentAvatarImageUrl field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetCurrentAvatarImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentAvatarImageUrl, true
}

// SetCurrentAvatarImageUrl sets field value
func (o *CurrentUser) SetCurrentAvatarImageUrl(v string) {
	o.CurrentAvatarImageUrl = v
}

// GetCurrentAvatarTags returns the CurrentAvatarTags field value
func (o *CurrentUser) GetCurrentAvatarTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CurrentAvatarTags
}

// GetCurrentAvatarTagsOk returns a tuple with the CurrentAvatarTags field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetCurrentAvatarTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentAvatarTags, true
}

// SetCurrentAvatarTags sets field value
func (o *CurrentUser) SetCurrentAvatarTags(v []string) {
	o.CurrentAvatarTags = v
}

// GetCurrentAvatarThumbnailImageUrl returns the CurrentAvatarThumbnailImageUrl field value
func (o *CurrentUser) GetCurrentAvatarThumbnailImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentAvatarThumbnailImageUrl
}

// GetCurrentAvatarThumbnailImageUrlOk returns a tuple with the CurrentAvatarThumbnailImageUrl field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetCurrentAvatarThumbnailImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentAvatarThumbnailImageUrl, true
}

// SetCurrentAvatarThumbnailImageUrl sets field value
func (o *CurrentUser) SetCurrentAvatarThumbnailImageUrl(v string) {
	o.CurrentAvatarThumbnailImageUrl = v
}

// GetDateJoined returns the DateJoined field value
func (o *CurrentUser) GetDateJoined() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DateJoined
}

// GetDateJoinedOk returns a tuple with the DateJoined field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetDateJoinedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateJoined, true
}

// SetDateJoined sets field value
func (o *CurrentUser) SetDateJoined(v string) {
	o.DateJoined = v
}

// GetDeveloperType returns the DeveloperType field value
func (o *CurrentUser) GetDeveloperType() DeveloperType {
	if o == nil {
		var ret DeveloperType
		return ret
	}

	return o.DeveloperType
}

// GetDeveloperTypeOk returns a tuple with the DeveloperType field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetDeveloperTypeOk() (*DeveloperType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeveloperType, true
}

// SetDeveloperType sets field value
func (o *CurrentUser) SetDeveloperType(v DeveloperType) {
	o.DeveloperType = v
}

// GetDiscordDetails returns the DiscordDetails field value if set, zero value otherwise.
func (o *CurrentUser) GetDiscordDetails() DiscordDetails {
	if o == nil || IsNil(o.DiscordDetails) {
		var ret DiscordDetails
		return ret
	}
	return *o.DiscordDetails
}

// GetDiscordDetailsOk returns a tuple with the DiscordDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetDiscordDetailsOk() (*DiscordDetails, bool) {
	if o == nil || IsNil(o.DiscordDetails) {
		return nil, false
	}
	return o.DiscordDetails, true
}

// HasDiscordDetails returns a boolean if a field has been set.
func (o *CurrentUser) HasDiscordDetails() bool {
	if o != nil && !IsNil(o.DiscordDetails) {
		return true
	}

	return false
}

// SetDiscordDetails gets a reference to the given DiscordDetails and assigns it to the DiscordDetails field.
func (o *CurrentUser) SetDiscordDetails(v DiscordDetails) {
	o.DiscordDetails = &v
}

// GetDiscordId returns the DiscordId field value if set, zero value otherwise.
func (o *CurrentUser) GetDiscordId() string {
	if o == nil || IsNil(o.DiscordId) {
		var ret string
		return ret
	}
	return *o.DiscordId
}

// GetDiscordIdOk returns a tuple with the DiscordId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetDiscordIdOk() (*string, bool) {
	if o == nil || IsNil(o.DiscordId) {
		return nil, false
	}
	return o.DiscordId, true
}

// HasDiscordId returns a boolean if a field has been set.
func (o *CurrentUser) HasDiscordId() bool {
	if o != nil && !IsNil(o.DiscordId) {
		return true
	}

	return false
}

// SetDiscordId gets a reference to the given string and assigns it to the DiscordId field.
func (o *CurrentUser) SetDiscordId(v string) {
	o.DiscordId = &v
}

// GetDisplayName returns the DisplayName field value
func (o *CurrentUser) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *CurrentUser) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetEmailVerified returns the EmailVerified field value
func (o *CurrentUser) GetEmailVerified() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EmailVerified
}

// GetEmailVerifiedOk returns a tuple with the EmailVerified field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetEmailVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailVerified, true
}

// SetEmailVerified sets field value
func (o *CurrentUser) SetEmailVerified(v bool) {
	o.EmailVerified = v
}

// GetFallbackAvatar returns the FallbackAvatar field value if set, zero value otherwise.
func (o *CurrentUser) GetFallbackAvatar() string {
	if o == nil || IsNil(o.FallbackAvatar) {
		var ret string
		return ret
	}
	return *o.FallbackAvatar
}

// GetFallbackAvatarOk returns a tuple with the FallbackAvatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetFallbackAvatarOk() (*string, bool) {
	if o == nil || IsNil(o.FallbackAvatar) {
		return nil, false
	}
	return o.FallbackAvatar, true
}

// HasFallbackAvatar returns a boolean if a field has been set.
func (o *CurrentUser) HasFallbackAvatar() bool {
	if o != nil && !IsNil(o.FallbackAvatar) {
		return true
	}

	return false
}

// SetFallbackAvatar gets a reference to the given string and assigns it to the FallbackAvatar field.
func (o *CurrentUser) SetFallbackAvatar(v string) {
	o.FallbackAvatar = &v
}

// GetFriendGroupNames returns the FriendGroupNames field value
// Deprecated
func (o *CurrentUser) GetFriendGroupNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.FriendGroupNames
}

// GetFriendGroupNamesOk returns a tuple with the FriendGroupNames field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *CurrentUser) GetFriendGroupNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FriendGroupNames, true
}

// SetFriendGroupNames sets field value
// Deprecated
func (o *CurrentUser) SetFriendGroupNames(v []string) {
	o.FriendGroupNames = v
}

// GetFriendKey returns the FriendKey field value
func (o *CurrentUser) GetFriendKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FriendKey
}

// GetFriendKeyOk returns a tuple with the FriendKey field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetFriendKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FriendKey, true
}

// SetFriendKey sets field value
func (o *CurrentUser) SetFriendKey(v string) {
	o.FriendKey = v
}

// GetFriends returns the Friends field value
func (o *CurrentUser) GetFriends() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Friends
}

// GetFriendsOk returns a tuple with the Friends field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetFriendsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Friends, true
}

// SetFriends sets field value
func (o *CurrentUser) SetFriends(v []string) {
	o.Friends = v
}

// GetGoogleDetails returns the GoogleDetails field value if set, zero value otherwise.
func (o *CurrentUser) GetGoogleDetails() map[string]interface{} {
	if o == nil || IsNil(o.GoogleDetails) {
		var ret map[string]interface{}
		return ret
	}
	return o.GoogleDetails
}

// GetGoogleDetailsOk returns a tuple with the GoogleDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetGoogleDetailsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.GoogleDetails) {
		return map[string]interface{}{}, false
	}
	return o.GoogleDetails, true
}

// HasGoogleDetails returns a boolean if a field has been set.
func (o *CurrentUser) HasGoogleDetails() bool {
	if o != nil && !IsNil(o.GoogleDetails) {
		return true
	}

	return false
}

// SetGoogleDetails gets a reference to the given map[string]interface{} and assigns it to the GoogleDetails field.
func (o *CurrentUser) SetGoogleDetails(v map[string]interface{}) {
	o.GoogleDetails = v
}

// GetGoogleId returns the GoogleId field value if set, zero value otherwise.
func (o *CurrentUser) GetGoogleId() string {
	if o == nil || IsNil(o.GoogleId) {
		var ret string
		return ret
	}
	return *o.GoogleId
}

// GetGoogleIdOk returns a tuple with the GoogleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetGoogleIdOk() (*string, bool) {
	if o == nil || IsNil(o.GoogleId) {
		return nil, false
	}
	return o.GoogleId, true
}

// HasGoogleId returns a boolean if a field has been set.
func (o *CurrentUser) HasGoogleId() bool {
	if o != nil && !IsNil(o.GoogleId) {
		return true
	}

	return false
}

// SetGoogleId gets a reference to the given string and assigns it to the GoogleId field.
func (o *CurrentUser) SetGoogleId(v string) {
	o.GoogleId = &v
}

// GetHasBirthday returns the HasBirthday field value
func (o *CurrentUser) GetHasBirthday() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasBirthday
}

// GetHasBirthdayOk returns a tuple with the HasBirthday field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetHasBirthdayOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasBirthday, true
}

// SetHasBirthday sets field value
func (o *CurrentUser) SetHasBirthday(v bool) {
	o.HasBirthday = v
}

// GetHasEmail returns the HasEmail field value
func (o *CurrentUser) GetHasEmail() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasEmail
}

// GetHasEmailOk returns a tuple with the HasEmail field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetHasEmailOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasEmail, true
}

// SetHasEmail sets field value
func (o *CurrentUser) SetHasEmail(v bool) {
	o.HasEmail = v
}

// GetHasLoggedInFromClient returns the HasLoggedInFromClient field value
func (o *CurrentUser) GetHasLoggedInFromClient() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasLoggedInFromClient
}

// GetHasLoggedInFromClientOk returns a tuple with the HasLoggedInFromClient field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetHasLoggedInFromClientOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasLoggedInFromClient, true
}

// SetHasLoggedInFromClient sets field value
func (o *CurrentUser) SetHasLoggedInFromClient(v bool) {
	o.HasLoggedInFromClient = v
}

// GetHasPendingEmail returns the HasPendingEmail field value
func (o *CurrentUser) GetHasPendingEmail() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasPendingEmail
}

// GetHasPendingEmailOk returns a tuple with the HasPendingEmail field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetHasPendingEmailOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasPendingEmail, true
}

// SetHasPendingEmail sets field value
func (o *CurrentUser) SetHasPendingEmail(v bool) {
	o.HasPendingEmail = v
}

// GetHideContentFilterSettings returns the HideContentFilterSettings field value if set, zero value otherwise.
func (o *CurrentUser) GetHideContentFilterSettings() bool {
	if o == nil || IsNil(o.HideContentFilterSettings) {
		var ret bool
		return ret
	}
	return *o.HideContentFilterSettings
}

// GetHideContentFilterSettingsOk returns a tuple with the HideContentFilterSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetHideContentFilterSettingsOk() (*bool, bool) {
	if o == nil || IsNil(o.HideContentFilterSettings) {
		return nil, false
	}
	return o.HideContentFilterSettings, true
}

// HasHideContentFilterSettings returns a boolean if a field has been set.
func (o *CurrentUser) HasHideContentFilterSettings() bool {
	if o != nil && !IsNil(o.HideContentFilterSettings) {
		return true
	}

	return false
}

// SetHideContentFilterSettings gets a reference to the given bool and assigns it to the HideContentFilterSettings field.
func (o *CurrentUser) SetHideContentFilterSettings(v bool) {
	o.HideContentFilterSettings = &v
}

// GetHomeLocation returns the HomeLocation field value
func (o *CurrentUser) GetHomeLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HomeLocation
}

// GetHomeLocationOk returns a tuple with the HomeLocation field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetHomeLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HomeLocation, true
}

// SetHomeLocation sets field value
func (o *CurrentUser) SetHomeLocation(v string) {
	o.HomeLocation = v
}

// GetId returns the Id field value
func (o *CurrentUser) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CurrentUser) SetId(v string) {
	o.Id = v
}

// GetIsAdult returns the IsAdult field value
func (o *CurrentUser) GetIsAdult() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAdult
}

// GetIsAdultOk returns a tuple with the IsAdult field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetIsAdultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsAdult, true
}

// SetIsAdult sets field value
func (o *CurrentUser) SetIsAdult(v bool) {
	o.IsAdult = v
}

// GetIsBoopingEnabled returns the IsBoopingEnabled field value if set, zero value otherwise.
func (o *CurrentUser) GetIsBoopingEnabled() bool {
	if o == nil || IsNil(o.IsBoopingEnabled) {
		var ret bool
		return ret
	}
	return *o.IsBoopingEnabled
}

// GetIsBoopingEnabledOk returns a tuple with the IsBoopingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetIsBoopingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBoopingEnabled) {
		return nil, false
	}
	return o.IsBoopingEnabled, true
}

// HasIsBoopingEnabled returns a boolean if a field has been set.
func (o *CurrentUser) HasIsBoopingEnabled() bool {
	if o != nil && !IsNil(o.IsBoopingEnabled) {
		return true
	}

	return false
}

// SetIsBoopingEnabled gets a reference to the given bool and assigns it to the IsBoopingEnabled field.
func (o *CurrentUser) SetIsBoopingEnabled(v bool) {
	o.IsBoopingEnabled = &v
}

// GetIsFriend returns the IsFriend field value
func (o *CurrentUser) GetIsFriend() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsFriend
}

// GetIsFriendOk returns a tuple with the IsFriend field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetIsFriendOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFriend, true
}

// SetIsFriend sets field value
func (o *CurrentUser) SetIsFriend(v bool) {
	o.IsFriend = v
}

// GetLastActivity returns the LastActivity field value if set, zero value otherwise.
func (o *CurrentUser) GetLastActivity() time.Time {
	if o == nil || IsNil(o.LastActivity) {
		var ret time.Time
		return ret
	}
	return *o.LastActivity
}

// GetLastActivityOk returns a tuple with the LastActivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetLastActivityOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastActivity) {
		return nil, false
	}
	return o.LastActivity, true
}

// HasLastActivity returns a boolean if a field has been set.
func (o *CurrentUser) HasLastActivity() bool {
	if o != nil && !IsNil(o.LastActivity) {
		return true
	}

	return false
}

// SetLastActivity gets a reference to the given time.Time and assigns it to the LastActivity field.
func (o *CurrentUser) SetLastActivity(v time.Time) {
	o.LastActivity = &v
}

// GetLastLogin returns the LastLogin field value
func (o *CurrentUser) GetLastLogin() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastLogin
}

// GetLastLoginOk returns a tuple with the LastLogin field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetLastLoginOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastLogin, true
}

// SetLastLogin sets field value
func (o *CurrentUser) SetLastLogin(v time.Time) {
	o.LastLogin = v
}

// GetLastMobile returns the LastMobile field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *CurrentUser) GetLastMobile() time.Time {
	if o == nil || o.LastMobile.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastMobile.Get()
}

// GetLastMobileOk returns a tuple with the LastMobile field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CurrentUser) GetLastMobileOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastMobile.Get(), o.LastMobile.IsSet()
}

// SetLastMobile sets field value
func (o *CurrentUser) SetLastMobile(v time.Time) {
	o.LastMobile.Set(&v)
}

// GetLastPlatform returns the LastPlatform field value
func (o *CurrentUser) GetLastPlatform() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastPlatform
}

// GetLastPlatformOk returns a tuple with the LastPlatform field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetLastPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastPlatform, true
}

// SetLastPlatform sets field value
func (o *CurrentUser) SetLastPlatform(v string) {
	o.LastPlatform = v
}

// GetObfuscatedEmail returns the ObfuscatedEmail field value
func (o *CurrentUser) GetObfuscatedEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObfuscatedEmail
}

// GetObfuscatedEmailOk returns a tuple with the ObfuscatedEmail field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetObfuscatedEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObfuscatedEmail, true
}

// SetObfuscatedEmail sets field value
func (o *CurrentUser) SetObfuscatedEmail(v string) {
	o.ObfuscatedEmail = v
}

// GetObfuscatedPendingEmail returns the ObfuscatedPendingEmail field value
func (o *CurrentUser) GetObfuscatedPendingEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObfuscatedPendingEmail
}

// GetObfuscatedPendingEmailOk returns a tuple with the ObfuscatedPendingEmail field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetObfuscatedPendingEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObfuscatedPendingEmail, true
}

// SetObfuscatedPendingEmail sets field value
func (o *CurrentUser) SetObfuscatedPendingEmail(v string) {
	o.ObfuscatedPendingEmail = v
}

// GetOculusId returns the OculusId field value
func (o *CurrentUser) GetOculusId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OculusId
}

// GetOculusIdOk returns a tuple with the OculusId field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetOculusIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OculusId, true
}

// SetOculusId sets field value
func (o *CurrentUser) SetOculusId(v string) {
	o.OculusId = v
}

// GetOfflineFriends returns the OfflineFriends field value if set, zero value otherwise.
func (o *CurrentUser) GetOfflineFriends() []string {
	if o == nil || IsNil(o.OfflineFriends) {
		var ret []string
		return ret
	}
	return o.OfflineFriends
}

// GetOfflineFriendsOk returns a tuple with the OfflineFriends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetOfflineFriendsOk() ([]string, bool) {
	if o == nil || IsNil(o.OfflineFriends) {
		return nil, false
	}
	return o.OfflineFriends, true
}

// HasOfflineFriends returns a boolean if a field has been set.
func (o *CurrentUser) HasOfflineFriends() bool {
	if o != nil && !IsNil(o.OfflineFriends) {
		return true
	}

	return false
}

// SetOfflineFriends gets a reference to the given []string and assigns it to the OfflineFriends field.
func (o *CurrentUser) SetOfflineFriends(v []string) {
	o.OfflineFriends = v
}

// GetOnlineFriends returns the OnlineFriends field value if set, zero value otherwise.
func (o *CurrentUser) GetOnlineFriends() []string {
	if o == nil || IsNil(o.OnlineFriends) {
		var ret []string
		return ret
	}
	return o.OnlineFriends
}

// GetOnlineFriendsOk returns a tuple with the OnlineFriends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetOnlineFriendsOk() ([]string, bool) {
	if o == nil || IsNil(o.OnlineFriends) {
		return nil, false
	}
	return o.OnlineFriends, true
}

// HasOnlineFriends returns a boolean if a field has been set.
func (o *CurrentUser) HasOnlineFriends() bool {
	if o != nil && !IsNil(o.OnlineFriends) {
		return true
	}

	return false
}

// SetOnlineFriends gets a reference to the given []string and assigns it to the OnlineFriends field.
func (o *CurrentUser) SetOnlineFriends(v []string) {
	o.OnlineFriends = v
}

// GetPastDisplayNames returns the PastDisplayNames field value
func (o *CurrentUser) GetPastDisplayNames() []PastDisplayName {
	if o == nil {
		var ret []PastDisplayName
		return ret
	}

	return o.PastDisplayNames
}

// GetPastDisplayNamesOk returns a tuple with the PastDisplayNames field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetPastDisplayNamesOk() ([]PastDisplayName, bool) {
	if o == nil {
		return nil, false
	}
	return o.PastDisplayNames, true
}

// SetPastDisplayNames sets field value
func (o *CurrentUser) SetPastDisplayNames(v []PastDisplayName) {
	o.PastDisplayNames = v
}

// GetPicoId returns the PicoId field value if set, zero value otherwise.
func (o *CurrentUser) GetPicoId() string {
	if o == nil || IsNil(o.PicoId) {
		var ret string
		return ret
	}
	return *o.PicoId
}

// GetPicoIdOk returns a tuple with the PicoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetPicoIdOk() (*string, bool) {
	if o == nil || IsNil(o.PicoId) {
		return nil, false
	}
	return o.PicoId, true
}

// HasPicoId returns a boolean if a field has been set.
func (o *CurrentUser) HasPicoId() bool {
	if o != nil && !IsNil(o.PicoId) {
		return true
	}

	return false
}

// SetPicoId gets a reference to the given string and assigns it to the PicoId field.
func (o *CurrentUser) SetPicoId(v string) {
	o.PicoId = &v
}

// GetPlatformHistory returns the PlatformHistory field value if set, zero value otherwise.
func (o *CurrentUser) GetPlatformHistory() []CurrentUserPlatformHistoryInner {
	if o == nil || IsNil(o.PlatformHistory) {
		var ret []CurrentUserPlatformHistoryInner
		return ret
	}
	return o.PlatformHistory
}

// GetPlatformHistoryOk returns a tuple with the PlatformHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetPlatformHistoryOk() ([]CurrentUserPlatformHistoryInner, bool) {
	if o == nil || IsNil(o.PlatformHistory) {
		return nil, false
	}
	return o.PlatformHistory, true
}

// HasPlatformHistory returns a boolean if a field has been set.
func (o *CurrentUser) HasPlatformHistory() bool {
	if o != nil && !IsNil(o.PlatformHistory) {
		return true
	}

	return false
}

// SetPlatformHistory gets a reference to the given []CurrentUserPlatformHistoryInner and assigns it to the PlatformHistory field.
func (o *CurrentUser) SetPlatformHistory(v []CurrentUserPlatformHistoryInner) {
	o.PlatformHistory = v
}

// GetPresence returns the Presence field value if set, zero value otherwise.
func (o *CurrentUser) GetPresence() CurrentUserPresence {
	if o == nil || IsNil(o.Presence) {
		var ret CurrentUserPresence
		return ret
	}
	return *o.Presence
}

// GetPresenceOk returns a tuple with the Presence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetPresenceOk() (*CurrentUserPresence, bool) {
	if o == nil || IsNil(o.Presence) {
		return nil, false
	}
	return o.Presence, true
}

// HasPresence returns a boolean if a field has been set.
func (o *CurrentUser) HasPresence() bool {
	if o != nil && !IsNil(o.Presence) {
		return true
	}

	return false
}

// SetPresence gets a reference to the given CurrentUserPresence and assigns it to the Presence field.
func (o *CurrentUser) SetPresence(v CurrentUserPresence) {
	o.Presence = &v
}

// GetProfilePicOverride returns the ProfilePicOverride field value
func (o *CurrentUser) GetProfilePicOverride() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfilePicOverride
}

// GetProfilePicOverrideOk returns a tuple with the ProfilePicOverride field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetProfilePicOverrideOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfilePicOverride, true
}

// SetProfilePicOverride sets field value
func (o *CurrentUser) SetProfilePicOverride(v string) {
	o.ProfilePicOverride = v
}

// GetProfilePicOverrideThumbnail returns the ProfilePicOverrideThumbnail field value
func (o *CurrentUser) GetProfilePicOverrideThumbnail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfilePicOverrideThumbnail
}

// GetProfilePicOverrideThumbnailOk returns a tuple with the ProfilePicOverrideThumbnail field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetProfilePicOverrideThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfilePicOverrideThumbnail, true
}

// SetProfilePicOverrideThumbnail sets field value
func (o *CurrentUser) SetProfilePicOverrideThumbnail(v string) {
	o.ProfilePicOverrideThumbnail = v
}

// GetPronouns returns the Pronouns field value
func (o *CurrentUser) GetPronouns() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pronouns
}

// GetPronounsOk returns a tuple with the Pronouns field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetPronounsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pronouns, true
}

// SetPronouns sets field value
func (o *CurrentUser) SetPronouns(v string) {
	o.Pronouns = v
}

// GetPronounsHistory returns the PronounsHistory field value
func (o *CurrentUser) GetPronounsHistory() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PronounsHistory
}

// GetPronounsHistoryOk returns a tuple with the PronounsHistory field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetPronounsHistoryOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PronounsHistory, true
}

// SetPronounsHistory sets field value
func (o *CurrentUser) SetPronounsHistory(v []string) {
	o.PronounsHistory = v
}

// GetQueuedInstance returns the QueuedInstance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CurrentUser) GetQueuedInstance() string {
	if o == nil || IsNil(o.QueuedInstance.Get()) {
		var ret string
		return ret
	}
	return *o.QueuedInstance.Get()
}

// GetQueuedInstanceOk returns a tuple with the QueuedInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CurrentUser) GetQueuedInstanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.QueuedInstance.Get(), o.QueuedInstance.IsSet()
}

// HasQueuedInstance returns a boolean if a field has been set.
func (o *CurrentUser) HasQueuedInstance() bool {
	if o != nil && o.QueuedInstance.IsSet() {
		return true
	}

	return false
}

// SetQueuedInstance gets a reference to the given NullableString and assigns it to the QueuedInstance field.
func (o *CurrentUser) SetQueuedInstance(v string) {
	o.QueuedInstance.Set(&v)
}

// SetQueuedInstanceNil sets the value for QueuedInstance to be an explicit nil
func (o *CurrentUser) SetQueuedInstanceNil() {
	o.QueuedInstance.Set(nil)
}

// UnsetQueuedInstance ensures that no value is present for QueuedInstance, not even an explicit nil
func (o *CurrentUser) UnsetQueuedInstance() {
	o.QueuedInstance.Unset()
}

// GetReceiveMobileInvitations returns the ReceiveMobileInvitations field value if set, zero value otherwise.
func (o *CurrentUser) GetReceiveMobileInvitations() bool {
	if o == nil || IsNil(o.ReceiveMobileInvitations) {
		var ret bool
		return ret
	}
	return *o.ReceiveMobileInvitations
}

// GetReceiveMobileInvitationsOk returns a tuple with the ReceiveMobileInvitations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetReceiveMobileInvitationsOk() (*bool, bool) {
	if o == nil || IsNil(o.ReceiveMobileInvitations) {
		return nil, false
	}
	return o.ReceiveMobileInvitations, true
}

// HasReceiveMobileInvitations returns a boolean if a field has been set.
func (o *CurrentUser) HasReceiveMobileInvitations() bool {
	if o != nil && !IsNil(o.ReceiveMobileInvitations) {
		return true
	}

	return false
}

// SetReceiveMobileInvitations gets a reference to the given bool and assigns it to the ReceiveMobileInvitations field.
func (o *CurrentUser) SetReceiveMobileInvitations(v bool) {
	o.ReceiveMobileInvitations = &v
}

// GetState returns the State field value
func (o *CurrentUser) GetState() UserState {
	if o == nil {
		var ret UserState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetStateOk() (*UserState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CurrentUser) SetState(v UserState) {
	o.State = v
}

// GetStatus returns the Status field value
func (o *CurrentUser) GetStatus() UserStatus {
	if o == nil {
		var ret UserStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetStatusOk() (*UserStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CurrentUser) SetStatus(v UserStatus) {
	o.Status = v
}

// GetStatusDescription returns the StatusDescription field value
func (o *CurrentUser) GetStatusDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatusDescription
}

// GetStatusDescriptionOk returns a tuple with the StatusDescription field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetStatusDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusDescription, true
}

// SetStatusDescription sets field value
func (o *CurrentUser) SetStatusDescription(v string) {
	o.StatusDescription = v
}

// GetStatusFirstTime returns the StatusFirstTime field value
func (o *CurrentUser) GetStatusFirstTime() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.StatusFirstTime
}

// GetStatusFirstTimeOk returns a tuple with the StatusFirstTime field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetStatusFirstTimeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusFirstTime, true
}

// SetStatusFirstTime sets field value
func (o *CurrentUser) SetStatusFirstTime(v bool) {
	o.StatusFirstTime = v
}

// GetStatusHistory returns the StatusHistory field value
func (o *CurrentUser) GetStatusHistory() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.StatusHistory
}

// GetStatusHistoryOk returns a tuple with the StatusHistory field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetStatusHistoryOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StatusHistory, true
}

// SetStatusHistory sets field value
func (o *CurrentUser) SetStatusHistory(v []string) {
	o.StatusHistory = v
}

// GetSteamDetails returns the SteamDetails field value
func (o *CurrentUser) GetSteamDetails() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.SteamDetails
}

// GetSteamDetailsOk returns a tuple with the SteamDetails field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetSteamDetailsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.SteamDetails, true
}

// SetSteamDetails sets field value
func (o *CurrentUser) SetSteamDetails(v map[string]interface{}) {
	o.SteamDetails = v
}

// GetSteamId returns the SteamId field value
func (o *CurrentUser) GetSteamId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SteamId
}

// GetSteamIdOk returns a tuple with the SteamId field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetSteamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SteamId, true
}

// SetSteamId sets field value
func (o *CurrentUser) SetSteamId(v string) {
	o.SteamId = v
}

// GetTags returns the Tags field value
func (o *CurrentUser) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *CurrentUser) SetTags(v []string) {
	o.Tags = v
}

// GetTwoFactorAuthEnabled returns the TwoFactorAuthEnabled field value
func (o *CurrentUser) GetTwoFactorAuthEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.TwoFactorAuthEnabled
}

// GetTwoFactorAuthEnabledOk returns a tuple with the TwoFactorAuthEnabled field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetTwoFactorAuthEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TwoFactorAuthEnabled, true
}

// SetTwoFactorAuthEnabled sets field value
func (o *CurrentUser) SetTwoFactorAuthEnabled(v bool) {
	o.TwoFactorAuthEnabled = v
}

// GetTwoFactorAuthEnabledDate returns the TwoFactorAuthEnabledDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CurrentUser) GetTwoFactorAuthEnabledDate() time.Time {
	if o == nil || IsNil(o.TwoFactorAuthEnabledDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.TwoFactorAuthEnabledDate.Get()
}

// GetTwoFactorAuthEnabledDateOk returns a tuple with the TwoFactorAuthEnabledDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CurrentUser) GetTwoFactorAuthEnabledDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.TwoFactorAuthEnabledDate.Get(), o.TwoFactorAuthEnabledDate.IsSet()
}

// HasTwoFactorAuthEnabledDate returns a boolean if a field has been set.
func (o *CurrentUser) HasTwoFactorAuthEnabledDate() bool {
	if o != nil && o.TwoFactorAuthEnabledDate.IsSet() {
		return true
	}

	return false
}

// SetTwoFactorAuthEnabledDate gets a reference to the given NullableTime and assigns it to the TwoFactorAuthEnabledDate field.
func (o *CurrentUser) SetTwoFactorAuthEnabledDate(v time.Time) {
	o.TwoFactorAuthEnabledDate.Set(&v)
}

// SetTwoFactorAuthEnabledDateNil sets the value for TwoFactorAuthEnabledDate to be an explicit nil
func (o *CurrentUser) SetTwoFactorAuthEnabledDateNil() {
	o.TwoFactorAuthEnabledDate.Set(nil)
}

// UnsetTwoFactorAuthEnabledDate ensures that no value is present for TwoFactorAuthEnabledDate, not even an explicit nil
func (o *CurrentUser) UnsetTwoFactorAuthEnabledDate() {
	o.TwoFactorAuthEnabledDate.Unset()
}

// GetUnsubscribe returns the Unsubscribe field value
func (o *CurrentUser) GetUnsubscribe() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Unsubscribe
}

// GetUnsubscribeOk returns a tuple with the Unsubscribe field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetUnsubscribeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unsubscribe, true
}

// SetUnsubscribe sets field value
func (o *CurrentUser) SetUnsubscribe(v bool) {
	o.Unsubscribe = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CurrentUser) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CurrentUser) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CurrentUser) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUserIcon returns the UserIcon field value
func (o *CurrentUser) GetUserIcon() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserIcon
}

// GetUserIconOk returns a tuple with the UserIcon field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetUserIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserIcon, true
}

// SetUserIcon sets field value
func (o *CurrentUser) SetUserIcon(v string) {
	o.UserIcon = v
}

// GetUserLanguage returns the UserLanguage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CurrentUser) GetUserLanguage() string {
	if o == nil || IsNil(o.UserLanguage.Get()) {
		var ret string
		return ret
	}
	return *o.UserLanguage.Get()
}

// GetUserLanguageOk returns a tuple with the UserLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CurrentUser) GetUserLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserLanguage.Get(), o.UserLanguage.IsSet()
}

// HasUserLanguage returns a boolean if a field has been set.
func (o *CurrentUser) HasUserLanguage() bool {
	if o != nil && o.UserLanguage.IsSet() {
		return true
	}

	return false
}

// SetUserLanguage gets a reference to the given NullableString and assigns it to the UserLanguage field.
func (o *CurrentUser) SetUserLanguage(v string) {
	o.UserLanguage.Set(&v)
}

// SetUserLanguageNil sets the value for UserLanguage to be an explicit nil
func (o *CurrentUser) SetUserLanguageNil() {
	o.UserLanguage.Set(nil)
}

// UnsetUserLanguage ensures that no value is present for UserLanguage, not even an explicit nil
func (o *CurrentUser) UnsetUserLanguage() {
	o.UserLanguage.Unset()
}

// GetUserLanguageCode returns the UserLanguageCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CurrentUser) GetUserLanguageCode() string {
	if o == nil || IsNil(o.UserLanguageCode.Get()) {
		var ret string
		return ret
	}
	return *o.UserLanguageCode.Get()
}

// GetUserLanguageCodeOk returns a tuple with the UserLanguageCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CurrentUser) GetUserLanguageCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserLanguageCode.Get(), o.UserLanguageCode.IsSet()
}

// HasUserLanguageCode returns a boolean if a field has been set.
func (o *CurrentUser) HasUserLanguageCode() bool {
	if o != nil && o.UserLanguageCode.IsSet() {
		return true
	}

	return false
}

// SetUserLanguageCode gets a reference to the given NullableString and assigns it to the UserLanguageCode field.
func (o *CurrentUser) SetUserLanguageCode(v string) {
	o.UserLanguageCode.Set(&v)
}

// SetUserLanguageCodeNil sets the value for UserLanguageCode to be an explicit nil
func (o *CurrentUser) SetUserLanguageCodeNil() {
	o.UserLanguageCode.Set(nil)
}

// UnsetUserLanguageCode ensures that no value is present for UserLanguageCode, not even an explicit nil
func (o *CurrentUser) UnsetUserLanguageCode() {
	o.UserLanguageCode.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise.
// Deprecated
func (o *CurrentUser) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CurrentUser) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *CurrentUser) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
// Deprecated
func (o *CurrentUser) SetUsername(v string) {
	o.Username = &v
}

// GetUsesGeneratedPassword returns the UsesGeneratedPassword field value
func (o *CurrentUser) GetUsesGeneratedPassword() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UsesGeneratedPassword
}

// GetUsesGeneratedPasswordOk returns a tuple with the UsesGeneratedPassword field value
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetUsesGeneratedPasswordOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsesGeneratedPassword, true
}

// SetUsesGeneratedPassword sets field value
func (o *CurrentUser) SetUsesGeneratedPassword(v bool) {
	o.UsesGeneratedPassword = v
}

// GetViveId returns the ViveId field value if set, zero value otherwise.
func (o *CurrentUser) GetViveId() string {
	if o == nil || IsNil(o.ViveId) {
		var ret string
		return ret
	}
	return *o.ViveId
}

// GetViveIdOk returns a tuple with the ViveId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentUser) GetViveIdOk() (*string, bool) {
	if o == nil || IsNil(o.ViveId) {
		return nil, false
	}
	return o.ViveId, true
}

// HasViveId returns a boolean if a field has been set.
func (o *CurrentUser) HasViveId() bool {
	if o != nil && !IsNil(o.ViveId) {
		return true
	}

	return false
}

// SetViveId gets a reference to the given string and assigns it to the ViveId field.
func (o *CurrentUser) SetViveId(v string) {
	o.ViveId = &v
}

func (o CurrentUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CurrentUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AcceptedPrivacyVersion) {
		toSerialize["acceptedPrivacyVersion"] = o.AcceptedPrivacyVersion
	}
	toSerialize["acceptedTOSVersion"] = o.AcceptedTOSVersion
	if o.AccountDeletionDate.IsSet() {
		toSerialize["accountDeletionDate"] = o.AccountDeletionDate.Get()
	}
	if o.AccountDeletionLog != nil {
		toSerialize["accountDeletionLog"] = o.AccountDeletionLog
	}
	if !IsNil(o.ActiveFriends) {
		toSerialize["activeFriends"] = o.ActiveFriends
	}
	toSerialize["ageVerificationStatus"] = o.AgeVerificationStatus
	toSerialize["ageVerified"] = o.AgeVerified
	toSerialize["allowAvatarCopying"] = o.AllowAvatarCopying
	if !IsNil(o.AuthToken) {
		toSerialize["authToken"] = o.AuthToken
	}
	if !IsNil(o.Badges) {
		toSerialize["badges"] = o.Badges
	}
	toSerialize["bio"] = o.Bio
	toSerialize["bioLinks"] = o.BioLinks
	if !IsNil(o.ContentFilters) {
		toSerialize["contentFilters"] = o.ContentFilters
	}
	toSerialize["currentAvatar"] = o.CurrentAvatar
	toSerialize["currentAvatarImageUrl"] = o.CurrentAvatarImageUrl
	toSerialize["currentAvatarTags"] = o.CurrentAvatarTags
	toSerialize["currentAvatarThumbnailImageUrl"] = o.CurrentAvatarThumbnailImageUrl
	toSerialize["date_joined"] = o.DateJoined
	toSerialize["developerType"] = o.DeveloperType
	if !IsNil(o.DiscordDetails) {
		toSerialize["discordDetails"] = o.DiscordDetails
	}
	if !IsNil(o.DiscordId) {
		toSerialize["discordId"] = o.DiscordId
	}
	toSerialize["displayName"] = o.DisplayName
	toSerialize["emailVerified"] = o.EmailVerified
	if !IsNil(o.FallbackAvatar) {
		toSerialize["fallbackAvatar"] = o.FallbackAvatar
	}
	toSerialize["friendGroupNames"] = o.FriendGroupNames
	toSerialize["friendKey"] = o.FriendKey
	toSerialize["friends"] = o.Friends
	if !IsNil(o.GoogleDetails) {
		toSerialize["googleDetails"] = o.GoogleDetails
	}
	if !IsNil(o.GoogleId) {
		toSerialize["googleId"] = o.GoogleId
	}
	toSerialize["hasBirthday"] = o.HasBirthday
	toSerialize["hasEmail"] = o.HasEmail
	toSerialize["hasLoggedInFromClient"] = o.HasLoggedInFromClient
	toSerialize["hasPendingEmail"] = o.HasPendingEmail
	if !IsNil(o.HideContentFilterSettings) {
		toSerialize["hideContentFilterSettings"] = o.HideContentFilterSettings
	}
	toSerialize["homeLocation"] = o.HomeLocation
	toSerialize["id"] = o.Id
	toSerialize["isAdult"] = o.IsAdult
	if !IsNil(o.IsBoopingEnabled) {
		toSerialize["isBoopingEnabled"] = o.IsBoopingEnabled
	}
	toSerialize["isFriend"] = o.IsFriend
	if !IsNil(o.LastActivity) {
		toSerialize["last_activity"] = o.LastActivity
	}
	toSerialize["last_login"] = o.LastLogin
	toSerialize["last_mobile"] = o.LastMobile.Get()
	toSerialize["last_platform"] = o.LastPlatform
	toSerialize["obfuscatedEmail"] = o.ObfuscatedEmail
	toSerialize["obfuscatedPendingEmail"] = o.ObfuscatedPendingEmail
	toSerialize["oculusId"] = o.OculusId
	if !IsNil(o.OfflineFriends) {
		toSerialize["offlineFriends"] = o.OfflineFriends
	}
	if !IsNil(o.OnlineFriends) {
		toSerialize["onlineFriends"] = o.OnlineFriends
	}
	toSerialize["pastDisplayNames"] = o.PastDisplayNames
	if !IsNil(o.PicoId) {
		toSerialize["picoId"] = o.PicoId
	}
	if !IsNil(o.PlatformHistory) {
		toSerialize["platform_history"] = o.PlatformHistory
	}
	if !IsNil(o.Presence) {
		toSerialize["presence"] = o.Presence
	}
	toSerialize["profilePicOverride"] = o.ProfilePicOverride
	toSerialize["profilePicOverrideThumbnail"] = o.ProfilePicOverrideThumbnail
	toSerialize["pronouns"] = o.Pronouns
	toSerialize["pronounsHistory"] = o.PronounsHistory
	if o.QueuedInstance.IsSet() {
		toSerialize["queuedInstance"] = o.QueuedInstance.Get()
	}
	if !IsNil(o.ReceiveMobileInvitations) {
		toSerialize["receiveMobileInvitations"] = o.ReceiveMobileInvitations
	}
	toSerialize["state"] = o.State
	toSerialize["status"] = o.Status
	toSerialize["statusDescription"] = o.StatusDescription
	toSerialize["statusFirstTime"] = o.StatusFirstTime
	toSerialize["statusHistory"] = o.StatusHistory
	toSerialize["steamDetails"] = o.SteamDetails
	toSerialize["steamId"] = o.SteamId
	toSerialize["tags"] = o.Tags
	toSerialize["twoFactorAuthEnabled"] = o.TwoFactorAuthEnabled
	if o.TwoFactorAuthEnabledDate.IsSet() {
		toSerialize["twoFactorAuthEnabledDate"] = o.TwoFactorAuthEnabledDate.Get()
	}
	toSerialize["unsubscribe"] = o.Unsubscribe
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	toSerialize["userIcon"] = o.UserIcon
	if o.UserLanguage.IsSet() {
		toSerialize["userLanguage"] = o.UserLanguage.Get()
	}
	if o.UserLanguageCode.IsSet() {
		toSerialize["userLanguageCode"] = o.UserLanguageCode.Get()
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	toSerialize["usesGeneratedPassword"] = o.UsesGeneratedPassword
	if !IsNil(o.ViveId) {
		toSerialize["viveId"] = o.ViveId
	}
	return toSerialize, nil
}

func (o *CurrentUser) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"acceptedTOSVersion",
		"ageVerificationStatus",
		"ageVerified",
		"allowAvatarCopying",
		"bio",
		"bioLinks",
		"currentAvatar",
		"currentAvatarImageUrl",
		"currentAvatarTags",
		"currentAvatarThumbnailImageUrl",
		"date_joined",
		"developerType",
		"displayName",
		"emailVerified",
		"friendGroupNames",
		"friendKey",
		"friends",
		"hasBirthday",
		"hasEmail",
		"hasLoggedInFromClient",
		"hasPendingEmail",
		"homeLocation",
		"id",
		"isAdult",
		"isFriend",
		"last_login",
		"last_mobile",
		"last_platform",
		"obfuscatedEmail",
		"obfuscatedPendingEmail",
		"oculusId",
		"pastDisplayNames",
		"profilePicOverride",
		"profilePicOverrideThumbnail",
		"pronouns",
		"pronounsHistory",
		"state",
		"status",
		"statusDescription",
		"statusFirstTime",
		"statusHistory",
		"steamDetails",
		"steamId",
		"tags",
		"twoFactorAuthEnabled",
		"unsubscribe",
		"userIcon",
		"usesGeneratedPassword",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCurrentUser := _CurrentUser{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCurrentUser)

	if err != nil {
		return err
	}

	*o = CurrentUser(varCurrentUser)

	return err
}

type NullableCurrentUser struct {
	value *CurrentUser
	isSet bool
}

func (v NullableCurrentUser) Get() *CurrentUser {
	return v.value
}

func (v *NullableCurrentUser) Set(val *CurrentUser) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrentUser) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrentUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrentUser(val *CurrentUser) *NullableCurrentUser {
	return &NullableCurrentUser{value: val, isSet: true}
}

func (v NullableCurrentUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrentUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
