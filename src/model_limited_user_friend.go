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

// checks if the LimitedUserFriend type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LimitedUserFriend{}

// LimitedUserFriend User object received when querying your friends list
type LimitedUserFriend struct {
	Bio *string `json:"bio,omitempty"`
	//
	BioLinks []string `json:"bioLinks,omitempty"`
	// When profilePicOverride is not empty, use it instead.
	CurrentAvatarImageUrl *string  `json:"currentAvatarImageUrl,omitempty"`
	CurrentAvatarTags     []string `json:"currentAvatarTags,omitempty"`
	// When profilePicOverride is not empty, use it instead.
	CurrentAvatarThumbnailImageUrl *string       `json:"currentAvatarThumbnailImageUrl,omitempty"`
	DeveloperType                  DeveloperType `json:"developerType"`
	DisplayName                    string        `json:"displayName"`
	FriendKey                      string        `json:"friendKey"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	Id           string       `json:"id"`
	ImageUrl     string       `json:"imageUrl"`
	IsFriend     bool         `json:"isFriend"`
	LastActivity NullableTime `json:"last_activity"`
	LastLogin    NullableTime `json:"last_login"`
	LastMobile   NullableTime `json:"last_mobile"`
	// This is normally `android`, `ios`, `standalonewindows`, `web`, or the empty value ``, but also supposedly can be any random Unity version such as `2019.2.4-801-Release` or `2019.2.2-772-Release` or even `unknownplatform`.
	LastPlatform                string     `json:"last_platform"`
	Location                    string     `json:"location"`
	Platform                    string     `json:"platform"`
	ProfilePicOverride          *string    `json:"profilePicOverride,omitempty"`
	ProfilePicOverrideThumbnail *string    `json:"profilePicOverrideThumbnail,omitempty"`
	Status                      UserStatus `json:"status"`
	StatusDescription           string     `json:"statusDescription"`
	// <- Always empty.
	Tags     []string `json:"tags"`
	UserIcon *string  `json:"userIcon,omitempty"`
}

type _LimitedUserFriend LimitedUserFriend

// NewLimitedUserFriend instantiates a new LimitedUserFriend object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLimitedUserFriend(developerType DeveloperType, displayName string, friendKey string, id string, imageUrl string, isFriend bool, lastActivity NullableTime, lastLogin NullableTime, lastMobile NullableTime, lastPlatform string, location string, platform string, status UserStatus, statusDescription string, tags []string) *LimitedUserFriend {
	this := LimitedUserFriend{}
	this.DeveloperType = developerType
	this.DisplayName = displayName
	this.FriendKey = friendKey
	this.Id = id
	this.ImageUrl = imageUrl
	this.IsFriend = isFriend
	this.LastActivity = lastActivity
	this.LastLogin = lastLogin
	this.LastMobile = lastMobile
	this.LastPlatform = lastPlatform
	this.Location = location
	this.Platform = platform
	this.Status = status
	this.StatusDescription = statusDescription
	this.Tags = tags
	return &this
}

// NewLimitedUserFriendWithDefaults instantiates a new LimitedUserFriend object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLimitedUserFriendWithDefaults() *LimitedUserFriend {
	this := LimitedUserFriend{}
	var developerType DeveloperType = DeveloperType_NONE
	this.DeveloperType = developerType
	var status UserStatus = UserStatus_OFFLINE
	this.Status = status
	return &this
}

// GetBio returns the Bio field value if set, zero value otherwise.
func (o *LimitedUserFriend) GetBio() string {
	if o == nil || IsNil(o.Bio) {
		var ret string
		return ret
	}
	return *o.Bio
}

// GetBioOk returns a tuple with the Bio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedUserFriend) GetBioOk() (*string, bool) {
	if o == nil || IsNil(o.Bio) {
		return nil, false
	}
	return o.Bio, true
}

// HasBio returns a boolean if a field has been set.
func (o *LimitedUserFriend) HasBio() bool {
	if o != nil && !IsNil(o.Bio) {
		return true
	}

	return false
}

// SetBio gets a reference to the given string and assigns it to the Bio field.
func (o *LimitedUserFriend) SetBio(v string) {
	o.Bio = &v
}

// GetBioLinks returns the BioLinks field value if set, zero value otherwise.
func (o *LimitedUserFriend) GetBioLinks() []string {
	if o == nil || IsNil(o.BioLinks) {
		var ret []string
		return ret
	}
	return o.BioLinks
}

// GetBioLinksOk returns a tuple with the BioLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedUserFriend) GetBioLinksOk() ([]string, bool) {
	if o == nil || IsNil(o.BioLinks) {
		return nil, false
	}
	return o.BioLinks, true
}

// HasBioLinks returns a boolean if a field has been set.
func (o *LimitedUserFriend) HasBioLinks() bool {
	if o != nil && !IsNil(o.BioLinks) {
		return true
	}

	return false
}

// SetBioLinks gets a reference to the given []string and assigns it to the BioLinks field.
func (o *LimitedUserFriend) SetBioLinks(v []string) {
	o.BioLinks = v
}

// GetCurrentAvatarImageUrl returns the CurrentAvatarImageUrl field value if set, zero value otherwise.
func (o *LimitedUserFriend) GetCurrentAvatarImageUrl() string {
	if o == nil || IsNil(o.CurrentAvatarImageUrl) {
		var ret string
		return ret
	}
	return *o.CurrentAvatarImageUrl
}

// GetCurrentAvatarImageUrlOk returns a tuple with the CurrentAvatarImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedUserFriend) GetCurrentAvatarImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentAvatarImageUrl) {
		return nil, false
	}
	return o.CurrentAvatarImageUrl, true
}

// HasCurrentAvatarImageUrl returns a boolean if a field has been set.
func (o *LimitedUserFriend) HasCurrentAvatarImageUrl() bool {
	if o != nil && !IsNil(o.CurrentAvatarImageUrl) {
		return true
	}

	return false
}

// SetCurrentAvatarImageUrl gets a reference to the given string and assigns it to the CurrentAvatarImageUrl field.
func (o *LimitedUserFriend) SetCurrentAvatarImageUrl(v string) {
	o.CurrentAvatarImageUrl = &v
}

// GetCurrentAvatarTags returns the CurrentAvatarTags field value if set, zero value otherwise.
func (o *LimitedUserFriend) GetCurrentAvatarTags() []string {
	if o == nil || IsNil(o.CurrentAvatarTags) {
		var ret []string
		return ret
	}
	return o.CurrentAvatarTags
}

// GetCurrentAvatarTagsOk returns a tuple with the CurrentAvatarTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedUserFriend) GetCurrentAvatarTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.CurrentAvatarTags) {
		return nil, false
	}
	return o.CurrentAvatarTags, true
}

// HasCurrentAvatarTags returns a boolean if a field has been set.
func (o *LimitedUserFriend) HasCurrentAvatarTags() bool {
	if o != nil && !IsNil(o.CurrentAvatarTags) {
		return true
	}

	return false
}

// SetCurrentAvatarTags gets a reference to the given []string and assigns it to the CurrentAvatarTags field.
func (o *LimitedUserFriend) SetCurrentAvatarTags(v []string) {
	o.CurrentAvatarTags = v
}

// GetCurrentAvatarThumbnailImageUrl returns the CurrentAvatarThumbnailImageUrl field value if set, zero value otherwise.
func (o *LimitedUserFriend) GetCurrentAvatarThumbnailImageUrl() string {
	if o == nil || IsNil(o.CurrentAvatarThumbnailImageUrl) {
		var ret string
		return ret
	}
	return *o.CurrentAvatarThumbnailImageUrl
}

// GetCurrentAvatarThumbnailImageUrlOk returns a tuple with the CurrentAvatarThumbnailImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedUserFriend) GetCurrentAvatarThumbnailImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentAvatarThumbnailImageUrl) {
		return nil, false
	}
	return o.CurrentAvatarThumbnailImageUrl, true
}

// HasCurrentAvatarThumbnailImageUrl returns a boolean if a field has been set.
func (o *LimitedUserFriend) HasCurrentAvatarThumbnailImageUrl() bool {
	if o != nil && !IsNil(o.CurrentAvatarThumbnailImageUrl) {
		return true
	}

	return false
}

// SetCurrentAvatarThumbnailImageUrl gets a reference to the given string and assigns it to the CurrentAvatarThumbnailImageUrl field.
func (o *LimitedUserFriend) SetCurrentAvatarThumbnailImageUrl(v string) {
	o.CurrentAvatarThumbnailImageUrl = &v
}

// GetDeveloperType returns the DeveloperType field value
func (o *LimitedUserFriend) GetDeveloperType() DeveloperType {
	if o == nil {
		var ret DeveloperType
		return ret
	}

	return o.DeveloperType
}

// GetDeveloperTypeOk returns a tuple with the DeveloperType field value
// and a boolean to check if the value has been set.
func (o *LimitedUserFriend) GetDeveloperTypeOk() (*DeveloperType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeveloperType, true
}

// SetDeveloperType sets field value
func (o *LimitedUserFriend) SetDeveloperType(v DeveloperType) {
	o.DeveloperType = v
}

// GetDisplayName returns the DisplayName field value
func (o *LimitedUserFriend) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *LimitedUserFriend) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *LimitedUserFriend) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetFriendKey returns the FriendKey field value
func (o *LimitedUserFriend) GetFriendKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FriendKey
}

// GetFriendKeyOk returns a tuple with the FriendKey field value
// and a boolean to check if the value has been set.
func (o *LimitedUserFriend) GetFriendKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FriendKey, true
}

// SetFriendKey sets field value
func (o *LimitedUserFriend) SetFriendKey(v string) {
	o.FriendKey = v
}

// GetId returns the Id field value
func (o *LimitedUserFriend) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LimitedUserFriend) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LimitedUserFriend) SetId(v string) {
	o.Id = v
}

// GetImageUrl returns the ImageUrl field value
func (o *LimitedUserFriend) GetImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
func (o *LimitedUserFriend) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageUrl, true
}

// SetImageUrl sets field value
func (o *LimitedUserFriend) SetImageUrl(v string) {
	o.ImageUrl = v
}

// GetIsFriend returns the IsFriend field value
func (o *LimitedUserFriend) GetIsFriend() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsFriend
}

// GetIsFriendOk returns a tuple with the IsFriend field value
// and a boolean to check if the value has been set.
func (o *LimitedUserFriend) GetIsFriendOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFriend, true
}

// SetIsFriend sets field value
func (o *LimitedUserFriend) SetIsFriend(v bool) {
	o.IsFriend = v
}

// GetLastActivity returns the LastActivity field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *LimitedUserFriend) GetLastActivity() time.Time {
	if o == nil || o.LastActivity.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastActivity.Get()
}

// GetLastActivityOk returns a tuple with the LastActivity field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LimitedUserFriend) GetLastActivityOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastActivity.Get(), o.LastActivity.IsSet()
}

// SetLastActivity sets field value
func (o *LimitedUserFriend) SetLastActivity(v time.Time) {
	o.LastActivity.Set(&v)
}

// GetLastLogin returns the LastLogin field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *LimitedUserFriend) GetLastLogin() time.Time {
	if o == nil || o.LastLogin.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastLogin.Get()
}

// GetLastLoginOk returns a tuple with the LastLogin field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LimitedUserFriend) GetLastLoginOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastLogin.Get(), o.LastLogin.IsSet()
}

// SetLastLogin sets field value
func (o *LimitedUserFriend) SetLastLogin(v time.Time) {
	o.LastLogin.Set(&v)
}

// GetLastMobile returns the LastMobile field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *LimitedUserFriend) GetLastMobile() time.Time {
	if o == nil || o.LastMobile.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastMobile.Get()
}

// GetLastMobileOk returns a tuple with the LastMobile field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LimitedUserFriend) GetLastMobileOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastMobile.Get(), o.LastMobile.IsSet()
}

// SetLastMobile sets field value
func (o *LimitedUserFriend) SetLastMobile(v time.Time) {
	o.LastMobile.Set(&v)
}

// GetLastPlatform returns the LastPlatform field value
func (o *LimitedUserFriend) GetLastPlatform() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastPlatform
}

// GetLastPlatformOk returns a tuple with the LastPlatform field value
// and a boolean to check if the value has been set.
func (o *LimitedUserFriend) GetLastPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastPlatform, true
}

// SetLastPlatform sets field value
func (o *LimitedUserFriend) SetLastPlatform(v string) {
	o.LastPlatform = v
}

// GetLocation returns the Location field value
func (o *LimitedUserFriend) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *LimitedUserFriend) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *LimitedUserFriend) SetLocation(v string) {
	o.Location = v
}

// GetPlatform returns the Platform field value
func (o *LimitedUserFriend) GetPlatform() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value
// and a boolean to check if the value has been set.
func (o *LimitedUserFriend) GetPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Platform, true
}

// SetPlatform sets field value
func (o *LimitedUserFriend) SetPlatform(v string) {
	o.Platform = v
}

// GetProfilePicOverride returns the ProfilePicOverride field value if set, zero value otherwise.
func (o *LimitedUserFriend) GetProfilePicOverride() string {
	if o == nil || IsNil(o.ProfilePicOverride) {
		var ret string
		return ret
	}
	return *o.ProfilePicOverride
}

// GetProfilePicOverrideOk returns a tuple with the ProfilePicOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedUserFriend) GetProfilePicOverrideOk() (*string, bool) {
	if o == nil || IsNil(o.ProfilePicOverride) {
		return nil, false
	}
	return o.ProfilePicOverride, true
}

// HasProfilePicOverride returns a boolean if a field has been set.
func (o *LimitedUserFriend) HasProfilePicOverride() bool {
	if o != nil && !IsNil(o.ProfilePicOverride) {
		return true
	}

	return false
}

// SetProfilePicOverride gets a reference to the given string and assigns it to the ProfilePicOverride field.
func (o *LimitedUserFriend) SetProfilePicOverride(v string) {
	o.ProfilePicOverride = &v
}

// GetProfilePicOverrideThumbnail returns the ProfilePicOverrideThumbnail field value if set, zero value otherwise.
func (o *LimitedUserFriend) GetProfilePicOverrideThumbnail() string {
	if o == nil || IsNil(o.ProfilePicOverrideThumbnail) {
		var ret string
		return ret
	}
	return *o.ProfilePicOverrideThumbnail
}

// GetProfilePicOverrideThumbnailOk returns a tuple with the ProfilePicOverrideThumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedUserFriend) GetProfilePicOverrideThumbnailOk() (*string, bool) {
	if o == nil || IsNil(o.ProfilePicOverrideThumbnail) {
		return nil, false
	}
	return o.ProfilePicOverrideThumbnail, true
}

// HasProfilePicOverrideThumbnail returns a boolean if a field has been set.
func (o *LimitedUserFriend) HasProfilePicOverrideThumbnail() bool {
	if o != nil && !IsNil(o.ProfilePicOverrideThumbnail) {
		return true
	}

	return false
}

// SetProfilePicOverrideThumbnail gets a reference to the given string and assigns it to the ProfilePicOverrideThumbnail field.
func (o *LimitedUserFriend) SetProfilePicOverrideThumbnail(v string) {
	o.ProfilePicOverrideThumbnail = &v
}

// GetStatus returns the Status field value
func (o *LimitedUserFriend) GetStatus() UserStatus {
	if o == nil {
		var ret UserStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *LimitedUserFriend) GetStatusOk() (*UserStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *LimitedUserFriend) SetStatus(v UserStatus) {
	o.Status = v
}

// GetStatusDescription returns the StatusDescription field value
func (o *LimitedUserFriend) GetStatusDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatusDescription
}

// GetStatusDescriptionOk returns a tuple with the StatusDescription field value
// and a boolean to check if the value has been set.
func (o *LimitedUserFriend) GetStatusDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusDescription, true
}

// SetStatusDescription sets field value
func (o *LimitedUserFriend) SetStatusDescription(v string) {
	o.StatusDescription = v
}

// GetTags returns the Tags field value
func (o *LimitedUserFriend) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *LimitedUserFriend) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *LimitedUserFriend) SetTags(v []string) {
	o.Tags = v
}

// GetUserIcon returns the UserIcon field value if set, zero value otherwise.
func (o *LimitedUserFriend) GetUserIcon() string {
	if o == nil || IsNil(o.UserIcon) {
		var ret string
		return ret
	}
	return *o.UserIcon
}

// GetUserIconOk returns a tuple with the UserIcon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedUserFriend) GetUserIconOk() (*string, bool) {
	if o == nil || IsNil(o.UserIcon) {
		return nil, false
	}
	return o.UserIcon, true
}

// HasUserIcon returns a boolean if a field has been set.
func (o *LimitedUserFriend) HasUserIcon() bool {
	if o != nil && !IsNil(o.UserIcon) {
		return true
	}

	return false
}

// SetUserIcon gets a reference to the given string and assigns it to the UserIcon field.
func (o *LimitedUserFriend) SetUserIcon(v string) {
	o.UserIcon = &v
}

func (o LimitedUserFriend) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LimitedUserFriend) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bio) {
		toSerialize["bio"] = o.Bio
	}
	if !IsNil(o.BioLinks) {
		toSerialize["bioLinks"] = o.BioLinks
	}
	if !IsNil(o.CurrentAvatarImageUrl) {
		toSerialize["currentAvatarImageUrl"] = o.CurrentAvatarImageUrl
	}
	if !IsNil(o.CurrentAvatarTags) {
		toSerialize["currentAvatarTags"] = o.CurrentAvatarTags
	}
	if !IsNil(o.CurrentAvatarThumbnailImageUrl) {
		toSerialize["currentAvatarThumbnailImageUrl"] = o.CurrentAvatarThumbnailImageUrl
	}
	toSerialize["developerType"] = o.DeveloperType
	toSerialize["displayName"] = o.DisplayName
	toSerialize["friendKey"] = o.FriendKey
	toSerialize["id"] = o.Id
	toSerialize["imageUrl"] = o.ImageUrl
	toSerialize["isFriend"] = o.IsFriend
	toSerialize["last_activity"] = o.LastActivity.Get()
	toSerialize["last_login"] = o.LastLogin.Get()
	toSerialize["last_mobile"] = o.LastMobile.Get()
	toSerialize["last_platform"] = o.LastPlatform
	toSerialize["location"] = o.Location
	toSerialize["platform"] = o.Platform
	if !IsNil(o.ProfilePicOverride) {
		toSerialize["profilePicOverride"] = o.ProfilePicOverride
	}
	if !IsNil(o.ProfilePicOverrideThumbnail) {
		toSerialize["profilePicOverrideThumbnail"] = o.ProfilePicOverrideThumbnail
	}
	toSerialize["status"] = o.Status
	toSerialize["statusDescription"] = o.StatusDescription
	toSerialize["tags"] = o.Tags
	if !IsNil(o.UserIcon) {
		toSerialize["userIcon"] = o.UserIcon
	}
	return toSerialize, nil
}

func (o *LimitedUserFriend) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"developerType",
		"displayName",
		"friendKey",
		"id",
		"imageUrl",
		"isFriend",
		"last_activity",
		"last_login",
		"last_mobile",
		"last_platform",
		"location",
		"platform",
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

	varLimitedUserFriend := _LimitedUserFriend{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLimitedUserFriend)

	if err != nil {
		return err
	}

	*o = LimitedUserFriend(varLimitedUserFriend)

	return err
}

type NullableLimitedUserFriend struct {
	value *LimitedUserFriend
	isSet bool
}

func (v NullableLimitedUserFriend) Get() *LimitedUserFriend {
	return v.value
}

func (v *NullableLimitedUserFriend) Set(val *LimitedUserFriend) {
	v.value = val
	v.isSet = true
}

func (v NullableLimitedUserFriend) IsSet() bool {
	return v.isSet
}

func (v *NullableLimitedUserFriend) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLimitedUserFriend(val *LimitedUserFriend) *NullableLimitedUserFriend {
	return &NullableLimitedUserFriend{value: val, isSet: true}
}

func (v NullableLimitedUserFriend) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLimitedUserFriend) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
