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

// checks if the LimitedUserInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LimitedUserInstance{}

// LimitedUserInstance User object received when querying your own instance
type LimitedUserInstance struct {
	AgeVerificationStatus AgeVerificationStatus `json:"ageVerificationStatus"`
	// `true` if, user is age verified (not 18+).
	AgeVerified        bool    `json:"ageVerified"`
	AllowAvatarCopying bool    `json:"allowAvatarCopying"`
	Bio                *string `json:"bio,omitempty"`
	//
	BioLinks []string `json:"bioLinks,omitempty"`
	// When profilePicOverride is not empty, use it instead.
	CurrentAvatarImageUrl string   `json:"currentAvatarImageUrl"`
	CurrentAvatarTags     []string `json:"currentAvatarTags"`
	// When profilePicOverride is not empty, use it instead.
	CurrentAvatarThumbnailImageUrl string        `json:"currentAvatarThumbnailImageUrl"`
	DateJoined                     NullableTime  `json:"date_joined"`
	DeveloperType                  DeveloperType `json:"developerType"`
	DisplayName                    string        `json:"displayName"`
	FriendKey                      string        `json:"friendKey"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	Id           string       `json:"id"`
	ImageUrl     *string      `json:"imageUrl,omitempty"`
	IsFriend     bool         `json:"isFriend"`
	LastActivity NullableTime `json:"last_activity"`
	LastMobile   NullableTime `json:"last_mobile,omitempty"`
	// This is normally `android`, `ios`, `standalonewindows`, `web`, or the empty value ``, but also supposedly can be any random Unity version such as `2019.2.4-801-Release` or `2019.2.2-772-Release` or even `unknownplatform`.
	LastPlatform string `json:"last_platform"`
	// This is normally `android`, `ios`, `standalonewindows`, `web`, or the empty value ``, but also supposedly can be any random Unity version such as `2019.2.4-801-Release` or `2019.2.2-772-Release` or even `unknownplatform`.
	Platform                    *string    `json:"platform,omitempty"`
	ProfilePicOverride          *string    `json:"profilePicOverride,omitempty"`
	ProfilePicOverrideThumbnail *string    `json:"profilePicOverrideThumbnail,omitempty"`
	Pronouns                    string     `json:"pronouns"`
	State                       *UserState `json:"state,omitempty"`
	Status                      UserStatus `json:"status"`
	StatusDescription           string     `json:"statusDescription"`
	Tags                        []string   `json:"tags"`
	UserIcon                    *string    `json:"userIcon,omitempty"`
}

type _LimitedUserInstance LimitedUserInstance

// NewLimitedUserInstance instantiates a new LimitedUserInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLimitedUserInstance(ageVerificationStatus AgeVerificationStatus, ageVerified bool, allowAvatarCopying bool, currentAvatarImageUrl string, currentAvatarTags []string, currentAvatarThumbnailImageUrl string, dateJoined NullableTime, developerType DeveloperType, displayName string, friendKey string, id string, isFriend bool, lastActivity NullableTime, lastPlatform string, pronouns string, status UserStatus, statusDescription string, tags []string) *LimitedUserInstance {
	this := LimitedUserInstance{}
	this.AgeVerificationStatus = ageVerificationStatus
	this.AgeVerified = ageVerified
	this.AllowAvatarCopying = allowAvatarCopying
	this.CurrentAvatarImageUrl = currentAvatarImageUrl
	this.CurrentAvatarTags = currentAvatarTags
	this.CurrentAvatarThumbnailImageUrl = currentAvatarThumbnailImageUrl
	this.DateJoined = dateJoined
	this.DeveloperType = developerType
	this.DisplayName = displayName
	this.FriendKey = friendKey
	this.Id = id
	this.IsFriend = isFriend
	this.LastActivity = lastActivity
	this.LastPlatform = lastPlatform
	this.Pronouns = pronouns
	var state UserState = UserState_OFFLINE
	this.State = &state
	this.Status = status
	this.StatusDescription = statusDescription
	this.Tags = tags
	return &this
}

// NewLimitedUserInstanceWithDefaults instantiates a new LimitedUserInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLimitedUserInstanceWithDefaults() *LimitedUserInstance {
	this := LimitedUserInstance{}
	var developerType DeveloperType = DeveloperType_NONE
	this.DeveloperType = developerType
	var state UserState = UserState_OFFLINE
	this.State = &state
	var status UserStatus = UserStatus_OFFLINE
	this.Status = status
	return &this
}

// GetAgeVerificationStatus returns the AgeVerificationStatus field value
func (o *LimitedUserInstance) GetAgeVerificationStatus() AgeVerificationStatus {
	if o == nil {
		var ret AgeVerificationStatus
		return ret
	}

	return o.AgeVerificationStatus
}

// GetAgeVerificationStatusOk returns a tuple with the AgeVerificationStatus field value
// and a boolean to check if the value has been set.
func (o *LimitedUserInstance) GetAgeVerificationStatusOk() (*AgeVerificationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgeVerificationStatus, true
}

// SetAgeVerificationStatus sets field value
func (o *LimitedUserInstance) SetAgeVerificationStatus(v AgeVerificationStatus) {
	o.AgeVerificationStatus = v
}

// GetAgeVerified returns the AgeVerified field value
func (o *LimitedUserInstance) GetAgeVerified() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AgeVerified
}

// GetAgeVerifiedOk returns a tuple with the AgeVerified field value
// and a boolean to check if the value has been set.
func (o *LimitedUserInstance) GetAgeVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgeVerified, true
}

// SetAgeVerified sets field value
func (o *LimitedUserInstance) SetAgeVerified(v bool) {
	o.AgeVerified = v
}

// GetAllowAvatarCopying returns the AllowAvatarCopying field value
func (o *LimitedUserInstance) GetAllowAvatarCopying() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowAvatarCopying
}

// GetAllowAvatarCopyingOk returns a tuple with the AllowAvatarCopying field value
// and a boolean to check if the value has been set.
func (o *LimitedUserInstance) GetAllowAvatarCopyingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowAvatarCopying, true
}

// SetAllowAvatarCopying sets field value
func (o *LimitedUserInstance) SetAllowAvatarCopying(v bool) {
	o.AllowAvatarCopying = v
}

// GetBio returns the Bio field value if set, zero value otherwise.
func (o *LimitedUserInstance) GetBio() string {
	if o == nil || IsNil(o.Bio) {
		var ret string
		return ret
	}
	return *o.Bio
}

// GetBioOk returns a tuple with the Bio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedUserInstance) GetBioOk() (*string, bool) {
	if o == nil || IsNil(o.Bio) {
		return nil, false
	}
	return o.Bio, true
}

// HasBio returns a boolean if a field has been set.
func (o *LimitedUserInstance) HasBio() bool {
	if o != nil && !IsNil(o.Bio) {
		return true
	}

	return false
}

// SetBio gets a reference to the given string and assigns it to the Bio field.
func (o *LimitedUserInstance) SetBio(v string) {
	o.Bio = &v
}

// GetBioLinks returns the BioLinks field value if set, zero value otherwise.
func (o *LimitedUserInstance) GetBioLinks() []string {
	if o == nil || IsNil(o.BioLinks) {
		var ret []string
		return ret
	}
	return o.BioLinks
}

// GetBioLinksOk returns a tuple with the BioLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedUserInstance) GetBioLinksOk() ([]string, bool) {
	if o == nil || IsNil(o.BioLinks) {
		return nil, false
	}
	return o.BioLinks, true
}

// HasBioLinks returns a boolean if a field has been set.
func (o *LimitedUserInstance) HasBioLinks() bool {
	if o != nil && !IsNil(o.BioLinks) {
		return true
	}

	return false
}

// SetBioLinks gets a reference to the given []string and assigns it to the BioLinks field.
func (o *LimitedUserInstance) SetBioLinks(v []string) {
	o.BioLinks = v
}

// GetCurrentAvatarImageUrl returns the CurrentAvatarImageUrl field value
func (o *LimitedUserInstance) GetCurrentAvatarImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentAvatarImageUrl
}

// GetCurrentAvatarImageUrlOk returns a tuple with the CurrentAvatarImageUrl field value
// and a boolean to check if the value has been set.
func (o *LimitedUserInstance) GetCurrentAvatarImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentAvatarImageUrl, true
}

// SetCurrentAvatarImageUrl sets field value
func (o *LimitedUserInstance) SetCurrentAvatarImageUrl(v string) {
	o.CurrentAvatarImageUrl = v
}

// GetCurrentAvatarTags returns the CurrentAvatarTags field value
func (o *LimitedUserInstance) GetCurrentAvatarTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CurrentAvatarTags
}

// GetCurrentAvatarTagsOk returns a tuple with the CurrentAvatarTags field value
// and a boolean to check if the value has been set.
func (o *LimitedUserInstance) GetCurrentAvatarTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentAvatarTags, true
}

// SetCurrentAvatarTags sets field value
func (o *LimitedUserInstance) SetCurrentAvatarTags(v []string) {
	o.CurrentAvatarTags = v
}

// GetCurrentAvatarThumbnailImageUrl returns the CurrentAvatarThumbnailImageUrl field value
func (o *LimitedUserInstance) GetCurrentAvatarThumbnailImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentAvatarThumbnailImageUrl
}

// GetCurrentAvatarThumbnailImageUrlOk returns a tuple with the CurrentAvatarThumbnailImageUrl field value
// and a boolean to check if the value has been set.
func (o *LimitedUserInstance) GetCurrentAvatarThumbnailImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentAvatarThumbnailImageUrl, true
}

// SetCurrentAvatarThumbnailImageUrl sets field value
func (o *LimitedUserInstance) SetCurrentAvatarThumbnailImageUrl(v string) {
	o.CurrentAvatarThumbnailImageUrl = v
}

// GetDateJoined returns the DateJoined field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *LimitedUserInstance) GetDateJoined() time.Time {
	if o == nil || o.DateJoined.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.DateJoined.Get()
}

// GetDateJoinedOk returns a tuple with the DateJoined field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LimitedUserInstance) GetDateJoinedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateJoined.Get(), o.DateJoined.IsSet()
}

// SetDateJoined sets field value
func (o *LimitedUserInstance) SetDateJoined(v time.Time) {
	o.DateJoined.Set(&v)
}

// GetDeveloperType returns the DeveloperType field value
func (o *LimitedUserInstance) GetDeveloperType() DeveloperType {
	if o == nil {
		var ret DeveloperType
		return ret
	}

	return o.DeveloperType
}

// GetDeveloperTypeOk returns a tuple with the DeveloperType field value
// and a boolean to check if the value has been set.
func (o *LimitedUserInstance) GetDeveloperTypeOk() (*DeveloperType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeveloperType, true
}

// SetDeveloperType sets field value
func (o *LimitedUserInstance) SetDeveloperType(v DeveloperType) {
	o.DeveloperType = v
}

// GetDisplayName returns the DisplayName field value
func (o *LimitedUserInstance) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *LimitedUserInstance) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *LimitedUserInstance) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetFriendKey returns the FriendKey field value
func (o *LimitedUserInstance) GetFriendKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FriendKey
}

// GetFriendKeyOk returns a tuple with the FriendKey field value
// and a boolean to check if the value has been set.
func (o *LimitedUserInstance) GetFriendKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FriendKey, true
}

// SetFriendKey sets field value
func (o *LimitedUserInstance) SetFriendKey(v string) {
	o.FriendKey = v
}

// GetId returns the Id field value
func (o *LimitedUserInstance) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LimitedUserInstance) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LimitedUserInstance) SetId(v string) {
	o.Id = v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *LimitedUserInstance) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedUserInstance) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *LimitedUserInstance) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *LimitedUserInstance) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetIsFriend returns the IsFriend field value
func (o *LimitedUserInstance) GetIsFriend() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsFriend
}

// GetIsFriendOk returns a tuple with the IsFriend field value
// and a boolean to check if the value has been set.
func (o *LimitedUserInstance) GetIsFriendOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFriend, true
}

// SetIsFriend sets field value
func (o *LimitedUserInstance) SetIsFriend(v bool) {
	o.IsFriend = v
}

// GetLastActivity returns the LastActivity field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *LimitedUserInstance) GetLastActivity() time.Time {
	if o == nil || o.LastActivity.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastActivity.Get()
}

// GetLastActivityOk returns a tuple with the LastActivity field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LimitedUserInstance) GetLastActivityOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastActivity.Get(), o.LastActivity.IsSet()
}

// SetLastActivity sets field value
func (o *LimitedUserInstance) SetLastActivity(v time.Time) {
	o.LastActivity.Set(&v)
}

// GetLastMobile returns the LastMobile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LimitedUserInstance) GetLastMobile() time.Time {
	if o == nil || IsNil(o.LastMobile.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastMobile.Get()
}

// GetLastMobileOk returns a tuple with the LastMobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LimitedUserInstance) GetLastMobileOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastMobile.Get(), o.LastMobile.IsSet()
}

// HasLastMobile returns a boolean if a field has been set.
func (o *LimitedUserInstance) HasLastMobile() bool {
	if o != nil && o.LastMobile.IsSet() {
		return true
	}

	return false
}

// SetLastMobile gets a reference to the given NullableTime and assigns it to the LastMobile field.
func (o *LimitedUserInstance) SetLastMobile(v time.Time) {
	o.LastMobile.Set(&v)
}

// SetLastMobileNil sets the value for LastMobile to be an explicit nil
func (o *LimitedUserInstance) SetLastMobileNil() {
	o.LastMobile.Set(nil)
}

// UnsetLastMobile ensures that no value is present for LastMobile, not even an explicit nil
func (o *LimitedUserInstance) UnsetLastMobile() {
	o.LastMobile.Unset()
}

// GetLastPlatform returns the LastPlatform field value
func (o *LimitedUserInstance) GetLastPlatform() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastPlatform
}

// GetLastPlatformOk returns a tuple with the LastPlatform field value
// and a boolean to check if the value has been set.
func (o *LimitedUserInstance) GetLastPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastPlatform, true
}

// SetLastPlatform sets field value
func (o *LimitedUserInstance) SetLastPlatform(v string) {
	o.LastPlatform = v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *LimitedUserInstance) GetPlatform() string {
	if o == nil || IsNil(o.Platform) {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedUserInstance) GetPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *LimitedUserInstance) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *LimitedUserInstance) SetPlatform(v string) {
	o.Platform = &v
}

// GetProfilePicOverride returns the ProfilePicOverride field value if set, zero value otherwise.
func (o *LimitedUserInstance) GetProfilePicOverride() string {
	if o == nil || IsNil(o.ProfilePicOverride) {
		var ret string
		return ret
	}
	return *o.ProfilePicOverride
}

// GetProfilePicOverrideOk returns a tuple with the ProfilePicOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedUserInstance) GetProfilePicOverrideOk() (*string, bool) {
	if o == nil || IsNil(o.ProfilePicOverride) {
		return nil, false
	}
	return o.ProfilePicOverride, true
}

// HasProfilePicOverride returns a boolean if a field has been set.
func (o *LimitedUserInstance) HasProfilePicOverride() bool {
	if o != nil && !IsNil(o.ProfilePicOverride) {
		return true
	}

	return false
}

// SetProfilePicOverride gets a reference to the given string and assigns it to the ProfilePicOverride field.
func (o *LimitedUserInstance) SetProfilePicOverride(v string) {
	o.ProfilePicOverride = &v
}

// GetProfilePicOverrideThumbnail returns the ProfilePicOverrideThumbnail field value if set, zero value otherwise.
func (o *LimitedUserInstance) GetProfilePicOverrideThumbnail() string {
	if o == nil || IsNil(o.ProfilePicOverrideThumbnail) {
		var ret string
		return ret
	}
	return *o.ProfilePicOverrideThumbnail
}

// GetProfilePicOverrideThumbnailOk returns a tuple with the ProfilePicOverrideThumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedUserInstance) GetProfilePicOverrideThumbnailOk() (*string, bool) {
	if o == nil || IsNil(o.ProfilePicOverrideThumbnail) {
		return nil, false
	}
	return o.ProfilePicOverrideThumbnail, true
}

// HasProfilePicOverrideThumbnail returns a boolean if a field has been set.
func (o *LimitedUserInstance) HasProfilePicOverrideThumbnail() bool {
	if o != nil && !IsNil(o.ProfilePicOverrideThumbnail) {
		return true
	}

	return false
}

// SetProfilePicOverrideThumbnail gets a reference to the given string and assigns it to the ProfilePicOverrideThumbnail field.
func (o *LimitedUserInstance) SetProfilePicOverrideThumbnail(v string) {
	o.ProfilePicOverrideThumbnail = &v
}

// GetPronouns returns the Pronouns field value
func (o *LimitedUserInstance) GetPronouns() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pronouns
}

// GetPronounsOk returns a tuple with the Pronouns field value
// and a boolean to check if the value has been set.
func (o *LimitedUserInstance) GetPronounsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pronouns, true
}

// SetPronouns sets field value
func (o *LimitedUserInstance) SetPronouns(v string) {
	o.Pronouns = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *LimitedUserInstance) GetState() UserState {
	if o == nil || IsNil(o.State) {
		var ret UserState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedUserInstance) GetStateOk() (*UserState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *LimitedUserInstance) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given UserState and assigns it to the State field.
func (o *LimitedUserInstance) SetState(v UserState) {
	o.State = &v
}

// GetStatus returns the Status field value
func (o *LimitedUserInstance) GetStatus() UserStatus {
	if o == nil {
		var ret UserStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *LimitedUserInstance) GetStatusOk() (*UserStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *LimitedUserInstance) SetStatus(v UserStatus) {
	o.Status = v
}

// GetStatusDescription returns the StatusDescription field value
func (o *LimitedUserInstance) GetStatusDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatusDescription
}

// GetStatusDescriptionOk returns a tuple with the StatusDescription field value
// and a boolean to check if the value has been set.
func (o *LimitedUserInstance) GetStatusDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusDescription, true
}

// SetStatusDescription sets field value
func (o *LimitedUserInstance) SetStatusDescription(v string) {
	o.StatusDescription = v
}

// GetTags returns the Tags field value
func (o *LimitedUserInstance) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *LimitedUserInstance) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *LimitedUserInstance) SetTags(v []string) {
	o.Tags = v
}

// GetUserIcon returns the UserIcon field value if set, zero value otherwise.
func (o *LimitedUserInstance) GetUserIcon() string {
	if o == nil || IsNil(o.UserIcon) {
		var ret string
		return ret
	}
	return *o.UserIcon
}

// GetUserIconOk returns a tuple with the UserIcon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedUserInstance) GetUserIconOk() (*string, bool) {
	if o == nil || IsNil(o.UserIcon) {
		return nil, false
	}
	return o.UserIcon, true
}

// HasUserIcon returns a boolean if a field has been set.
func (o *LimitedUserInstance) HasUserIcon() bool {
	if o != nil && !IsNil(o.UserIcon) {
		return true
	}

	return false
}

// SetUserIcon gets a reference to the given string and assigns it to the UserIcon field.
func (o *LimitedUserInstance) SetUserIcon(v string) {
	o.UserIcon = &v
}

func (o LimitedUserInstance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LimitedUserInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ageVerificationStatus"] = o.AgeVerificationStatus
	toSerialize["ageVerified"] = o.AgeVerified
	toSerialize["allowAvatarCopying"] = o.AllowAvatarCopying
	if !IsNil(o.Bio) {
		toSerialize["bio"] = o.Bio
	}
	if !IsNil(o.BioLinks) {
		toSerialize["bioLinks"] = o.BioLinks
	}
	toSerialize["currentAvatarImageUrl"] = o.CurrentAvatarImageUrl
	toSerialize["currentAvatarTags"] = o.CurrentAvatarTags
	toSerialize["currentAvatarThumbnailImageUrl"] = o.CurrentAvatarThumbnailImageUrl
	toSerialize["date_joined"] = o.DateJoined.Get()
	toSerialize["developerType"] = o.DeveloperType
	toSerialize["displayName"] = o.DisplayName
	toSerialize["friendKey"] = o.FriendKey
	toSerialize["id"] = o.Id
	if !IsNil(o.ImageUrl) {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	toSerialize["isFriend"] = o.IsFriend
	toSerialize["last_activity"] = o.LastActivity.Get()
	if o.LastMobile.IsSet() {
		toSerialize["last_mobile"] = o.LastMobile.Get()
	}
	toSerialize["last_platform"] = o.LastPlatform
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.ProfilePicOverride) {
		toSerialize["profilePicOverride"] = o.ProfilePicOverride
	}
	if !IsNil(o.ProfilePicOverrideThumbnail) {
		toSerialize["profilePicOverrideThumbnail"] = o.ProfilePicOverrideThumbnail
	}
	toSerialize["pronouns"] = o.Pronouns
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	toSerialize["status"] = o.Status
	toSerialize["statusDescription"] = o.StatusDescription
	toSerialize["tags"] = o.Tags
	if !IsNil(o.UserIcon) {
		toSerialize["userIcon"] = o.UserIcon
	}
	return toSerialize, nil
}

func (o *LimitedUserInstance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ageVerificationStatus",
		"ageVerified",
		"allowAvatarCopying",
		"currentAvatarImageUrl",
		"currentAvatarTags",
		"currentAvatarThumbnailImageUrl",
		"date_joined",
		"developerType",
		"displayName",
		"friendKey",
		"id",
		"isFriend",
		"last_activity",
		"last_platform",
		"pronouns",
		"status",
		"statusDescription",
		"tags",
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

	varLimitedUserInstance := _LimitedUserInstance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLimitedUserInstance)

	if err != nil {
		return err
	}

	*o = LimitedUserInstance(varLimitedUserInstance)

	return err
}

type NullableLimitedUserInstance struct {
	value *LimitedUserInstance
	isSet bool
}

func (v NullableLimitedUserInstance) Get() *LimitedUserInstance {
	return v.value
}

func (v *NullableLimitedUserInstance) Set(val *LimitedUserInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableLimitedUserInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableLimitedUserInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLimitedUserInstance(val *LimitedUserInstance) *NullableLimitedUserInstance {
	return &NullableLimitedUserInstance{value: val, isSet: true}
}

func (v NullableLimitedUserInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLimitedUserInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
