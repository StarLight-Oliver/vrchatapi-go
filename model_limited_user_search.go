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
)

// checks if the LimitedUserSearch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LimitedUserSearch{}

// LimitedUserSearch User object received when searching
type LimitedUserSearch struct {
	Bio *string `json:"bio,omitempty"`
	//
	BioLinks []string `json:"bioLinks,omitempty"`
	// When profilePicOverride is not empty, use it instead.
	CurrentAvatarImageUrl string   `json:"currentAvatarImageUrl"`
	CurrentAvatarTags     []string `json:"currentAvatarTags"`
	// When profilePicOverride is not empty, use it instead.
	CurrentAvatarThumbnailImageUrl string        `json:"currentAvatarThumbnailImageUrl"`
	DeveloperType                  DeveloperType `json:"developerType"`
	DisplayName                    string        `json:"displayName"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	Id       string `json:"id"`
	IsFriend bool   `json:"isFriend"`
	// This is normally `android`, `ios`, `standalonewindows`, `web`, or the empty value ``, but also supposedly can be any random Unity version such as `2019.2.4-801-Release` or `2019.2.2-772-Release` or even `unknownplatform`.
	LastPlatform       string     `json:"last_platform"`
	ProfilePicOverride *string    `json:"profilePicOverride,omitempty"`
	Pronouns           *string    `json:"pronouns,omitempty"`
	Status             UserStatus `json:"status"`
	StatusDescription  string     `json:"statusDescription"`
	// <- Always empty.
	Tags     []string `json:"tags"`
	UserIcon *string  `json:"userIcon,omitempty"`
}

type _LimitedUserSearch LimitedUserSearch

// NewLimitedUserSearch instantiates a new LimitedUserSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLimitedUserSearch(currentAvatarImageUrl string, currentAvatarTags []string, currentAvatarThumbnailImageUrl string, developerType DeveloperType, displayName string, id string, isFriend bool, lastPlatform string, status UserStatus, statusDescription string, tags []string) *LimitedUserSearch {
	this := LimitedUserSearch{}
	this.CurrentAvatarImageUrl = currentAvatarImageUrl
	this.CurrentAvatarTags = currentAvatarTags
	this.CurrentAvatarThumbnailImageUrl = currentAvatarThumbnailImageUrl
	this.DeveloperType = developerType
	this.DisplayName = displayName
	this.Id = id
	this.IsFriend = isFriend
	this.LastPlatform = lastPlatform
	this.Status = status
	this.StatusDescription = statusDescription
	this.Tags = tags
	return &this
}

// NewLimitedUserSearchWithDefaults instantiates a new LimitedUserSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLimitedUserSearchWithDefaults() *LimitedUserSearch {
	this := LimitedUserSearch{}
	var developerType DeveloperType = DeveloperType_NONE
	this.DeveloperType = developerType
	var status UserStatus = UserStatus_OFFLINE
	this.Status = status
	return &this
}

// GetBio returns the Bio field value if set, zero value otherwise.
func (o *LimitedUserSearch) GetBio() string {
	if o == nil || IsNil(o.Bio) {
		var ret string
		return ret
	}
	return *o.Bio
}

// GetBioOk returns a tuple with the Bio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedUserSearch) GetBioOk() (*string, bool) {
	if o == nil || IsNil(o.Bio) {
		return nil, false
	}
	return o.Bio, true
}

// HasBio returns a boolean if a field has been set.
func (o *LimitedUserSearch) HasBio() bool {
	if o != nil && !IsNil(o.Bio) {
		return true
	}

	return false
}

// SetBio gets a reference to the given string and assigns it to the Bio field.
func (o *LimitedUserSearch) SetBio(v string) {
	o.Bio = &v
}

// GetBioLinks returns the BioLinks field value if set, zero value otherwise.
func (o *LimitedUserSearch) GetBioLinks() []string {
	if o == nil || IsNil(o.BioLinks) {
		var ret []string
		return ret
	}
	return o.BioLinks
}

// GetBioLinksOk returns a tuple with the BioLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedUserSearch) GetBioLinksOk() ([]string, bool) {
	if o == nil || IsNil(o.BioLinks) {
		return nil, false
	}
	return o.BioLinks, true
}

// HasBioLinks returns a boolean if a field has been set.
func (o *LimitedUserSearch) HasBioLinks() bool {
	if o != nil && !IsNil(o.BioLinks) {
		return true
	}

	return false
}

// SetBioLinks gets a reference to the given []string and assigns it to the BioLinks field.
func (o *LimitedUserSearch) SetBioLinks(v []string) {
	o.BioLinks = v
}

// GetCurrentAvatarImageUrl returns the CurrentAvatarImageUrl field value
func (o *LimitedUserSearch) GetCurrentAvatarImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentAvatarImageUrl
}

// GetCurrentAvatarImageUrlOk returns a tuple with the CurrentAvatarImageUrl field value
// and a boolean to check if the value has been set.
func (o *LimitedUserSearch) GetCurrentAvatarImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentAvatarImageUrl, true
}

// SetCurrentAvatarImageUrl sets field value
func (o *LimitedUserSearch) SetCurrentAvatarImageUrl(v string) {
	o.CurrentAvatarImageUrl = v
}

// GetCurrentAvatarTags returns the CurrentAvatarTags field value
func (o *LimitedUserSearch) GetCurrentAvatarTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CurrentAvatarTags
}

// GetCurrentAvatarTagsOk returns a tuple with the CurrentAvatarTags field value
// and a boolean to check if the value has been set.
func (o *LimitedUserSearch) GetCurrentAvatarTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentAvatarTags, true
}

// SetCurrentAvatarTags sets field value
func (o *LimitedUserSearch) SetCurrentAvatarTags(v []string) {
	o.CurrentAvatarTags = v
}

// GetCurrentAvatarThumbnailImageUrl returns the CurrentAvatarThumbnailImageUrl field value
func (o *LimitedUserSearch) GetCurrentAvatarThumbnailImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentAvatarThumbnailImageUrl
}

// GetCurrentAvatarThumbnailImageUrlOk returns a tuple with the CurrentAvatarThumbnailImageUrl field value
// and a boolean to check if the value has been set.
func (o *LimitedUserSearch) GetCurrentAvatarThumbnailImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentAvatarThumbnailImageUrl, true
}

// SetCurrentAvatarThumbnailImageUrl sets field value
func (o *LimitedUserSearch) SetCurrentAvatarThumbnailImageUrl(v string) {
	o.CurrentAvatarThumbnailImageUrl = v
}

// GetDeveloperType returns the DeveloperType field value
func (o *LimitedUserSearch) GetDeveloperType() DeveloperType {
	if o == nil {
		var ret DeveloperType
		return ret
	}

	return o.DeveloperType
}

// GetDeveloperTypeOk returns a tuple with the DeveloperType field value
// and a boolean to check if the value has been set.
func (o *LimitedUserSearch) GetDeveloperTypeOk() (*DeveloperType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeveloperType, true
}

// SetDeveloperType sets field value
func (o *LimitedUserSearch) SetDeveloperType(v DeveloperType) {
	o.DeveloperType = v
}

// GetDisplayName returns the DisplayName field value
func (o *LimitedUserSearch) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *LimitedUserSearch) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *LimitedUserSearch) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetId returns the Id field value
func (o *LimitedUserSearch) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LimitedUserSearch) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LimitedUserSearch) SetId(v string) {
	o.Id = v
}

// GetIsFriend returns the IsFriend field value
func (o *LimitedUserSearch) GetIsFriend() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsFriend
}

// GetIsFriendOk returns a tuple with the IsFriend field value
// and a boolean to check if the value has been set.
func (o *LimitedUserSearch) GetIsFriendOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFriend, true
}

// SetIsFriend sets field value
func (o *LimitedUserSearch) SetIsFriend(v bool) {
	o.IsFriend = v
}

// GetLastPlatform returns the LastPlatform field value
func (o *LimitedUserSearch) GetLastPlatform() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastPlatform
}

// GetLastPlatformOk returns a tuple with the LastPlatform field value
// and a boolean to check if the value has been set.
func (o *LimitedUserSearch) GetLastPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastPlatform, true
}

// SetLastPlatform sets field value
func (o *LimitedUserSearch) SetLastPlatform(v string) {
	o.LastPlatform = v
}

// GetProfilePicOverride returns the ProfilePicOverride field value if set, zero value otherwise.
func (o *LimitedUserSearch) GetProfilePicOverride() string {
	if o == nil || IsNil(o.ProfilePicOverride) {
		var ret string
		return ret
	}
	return *o.ProfilePicOverride
}

// GetProfilePicOverrideOk returns a tuple with the ProfilePicOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedUserSearch) GetProfilePicOverrideOk() (*string, bool) {
	if o == nil || IsNil(o.ProfilePicOverride) {
		return nil, false
	}
	return o.ProfilePicOverride, true
}

// HasProfilePicOverride returns a boolean if a field has been set.
func (o *LimitedUserSearch) HasProfilePicOverride() bool {
	if o != nil && !IsNil(o.ProfilePicOverride) {
		return true
	}

	return false
}

// SetProfilePicOverride gets a reference to the given string and assigns it to the ProfilePicOverride field.
func (o *LimitedUserSearch) SetProfilePicOverride(v string) {
	o.ProfilePicOverride = &v
}

// GetPronouns returns the Pronouns field value if set, zero value otherwise.
func (o *LimitedUserSearch) GetPronouns() string {
	if o == nil || IsNil(o.Pronouns) {
		var ret string
		return ret
	}
	return *o.Pronouns
}

// GetPronounsOk returns a tuple with the Pronouns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedUserSearch) GetPronounsOk() (*string, bool) {
	if o == nil || IsNil(o.Pronouns) {
		return nil, false
	}
	return o.Pronouns, true
}

// HasPronouns returns a boolean if a field has been set.
func (o *LimitedUserSearch) HasPronouns() bool {
	if o != nil && !IsNil(o.Pronouns) {
		return true
	}

	return false
}

// SetPronouns gets a reference to the given string and assigns it to the Pronouns field.
func (o *LimitedUserSearch) SetPronouns(v string) {
	o.Pronouns = &v
}

// GetStatus returns the Status field value
func (o *LimitedUserSearch) GetStatus() UserStatus {
	if o == nil {
		var ret UserStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *LimitedUserSearch) GetStatusOk() (*UserStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *LimitedUserSearch) SetStatus(v UserStatus) {
	o.Status = v
}

// GetStatusDescription returns the StatusDescription field value
func (o *LimitedUserSearch) GetStatusDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatusDescription
}

// GetStatusDescriptionOk returns a tuple with the StatusDescription field value
// and a boolean to check if the value has been set.
func (o *LimitedUserSearch) GetStatusDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusDescription, true
}

// SetStatusDescription sets field value
func (o *LimitedUserSearch) SetStatusDescription(v string) {
	o.StatusDescription = v
}

// GetTags returns the Tags field value
func (o *LimitedUserSearch) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *LimitedUserSearch) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *LimitedUserSearch) SetTags(v []string) {
	o.Tags = v
}

// GetUserIcon returns the UserIcon field value if set, zero value otherwise.
func (o *LimitedUserSearch) GetUserIcon() string {
	if o == nil || IsNil(o.UserIcon) {
		var ret string
		return ret
	}
	return *o.UserIcon
}

// GetUserIconOk returns a tuple with the UserIcon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedUserSearch) GetUserIconOk() (*string, bool) {
	if o == nil || IsNil(o.UserIcon) {
		return nil, false
	}
	return o.UserIcon, true
}

// HasUserIcon returns a boolean if a field has been set.
func (o *LimitedUserSearch) HasUserIcon() bool {
	if o != nil && !IsNil(o.UserIcon) {
		return true
	}

	return false
}

// SetUserIcon gets a reference to the given string and assigns it to the UserIcon field.
func (o *LimitedUserSearch) SetUserIcon(v string) {
	o.UserIcon = &v
}

func (o LimitedUserSearch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LimitedUserSearch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bio) {
		toSerialize["bio"] = o.Bio
	}
	if !IsNil(o.BioLinks) {
		toSerialize["bioLinks"] = o.BioLinks
	}
	toSerialize["currentAvatarImageUrl"] = o.CurrentAvatarImageUrl
	toSerialize["currentAvatarTags"] = o.CurrentAvatarTags
	toSerialize["currentAvatarThumbnailImageUrl"] = o.CurrentAvatarThumbnailImageUrl
	toSerialize["developerType"] = o.DeveloperType
	toSerialize["displayName"] = o.DisplayName
	toSerialize["id"] = o.Id
	toSerialize["isFriend"] = o.IsFriend
	toSerialize["last_platform"] = o.LastPlatform
	if !IsNil(o.ProfilePicOverride) {
		toSerialize["profilePicOverride"] = o.ProfilePicOverride
	}
	if !IsNil(o.Pronouns) {
		toSerialize["pronouns"] = o.Pronouns
	}
	toSerialize["status"] = o.Status
	toSerialize["statusDescription"] = o.StatusDescription
	toSerialize["tags"] = o.Tags
	if !IsNil(o.UserIcon) {
		toSerialize["userIcon"] = o.UserIcon
	}
	return toSerialize, nil
}

func (o *LimitedUserSearch) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"currentAvatarImageUrl",
		"currentAvatarTags",
		"currentAvatarThumbnailImageUrl",
		"developerType",
		"displayName",
		"id",
		"isFriend",
		"last_platform",
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

	varLimitedUserSearch := _LimitedUserSearch{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLimitedUserSearch)

	if err != nil {
		return err
	}

	*o = LimitedUserSearch(varLimitedUserSearch)

	return err
}

type NullableLimitedUserSearch struct {
	value *LimitedUserSearch
	isSet bool
}

func (v NullableLimitedUserSearch) Get() *LimitedUserSearch {
	return v.value
}

func (v *NullableLimitedUserSearch) Set(val *LimitedUserSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableLimitedUserSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableLimitedUserSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLimitedUserSearch(val *LimitedUserSearch) *NullableLimitedUserSearch {
	return &NullableLimitedUserSearch{value: val, isSet: true}
}

func (v NullableLimitedUserSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLimitedUserSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
