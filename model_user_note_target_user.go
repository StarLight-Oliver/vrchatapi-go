/*
VRChat API Documentation

![VRChat API Banner](https://vrchatapi.github.io/assets/img/api_banner_1500x400.png)  # Welcome to the VRChat API  Before we begin, we would like to state this is a **COMMUNITY DRIVEN PROJECT**. This means that everything you read on here was written by the community itself and is **not** officially supported by VRChat. The documentation is provided \"AS IS\", and any action you take towards VRChat is completely your own responsibility.  The documentation and additional libraries SHALL ONLY be used for applications interacting with VRChat's API in accordance with their [Terms of Service](https://hello.vrchat.com/legal) and [Community Guidelines](https://hello.vrchat.com/community-guidelines), and MUST NOT be used for modifying the client, \"avatar ripping\", or other illegal activities. Malicious usage or spamming the API may result in account termination. Certain parts of the API are also more sensitive than others, for example moderation, so please tread extra carefully and read the warnings when present.  ![Tupper Policy on API](https://i.imgur.com/yLlW7Ok.png)  Finally, use of the API using applications other than the approved methods (website, VRChat application, Unity SDK) is not officially supported. VRChat provides no guarantee or support for external applications using the API. Access to API endpoints may break **at any time, without notice**. Therefore, please **do not ping** VRChat Staff in the VRChat Discord if you are having API problems, as they do not provide API support. We will make a best effort in keeping this documentation and associated language libraries up to date, but things might be outdated or missing. If you find that something is no longer valid, please contact us on Discord or [create an issue](https://github.com/vrchatapi/specification/issues) and tell us so we can fix it.  # Getting Started  The VRChat API can be used to programmatically retrieve or update information regarding your profile, friends, avatars, worlds and more. The API consists of two parts, \"Photon\" which is only used in-game, and the \"Web API\" which is used by both the game and the website. This documentation focuses only on the Web API.  The API is designed around the REST ideology, providing semi-simple and usually predictable URIs to access and modify objects. Requests support standard HTTP methods like GET, PUT, POST, and DELETE and standard status codes. Response bodies are always UTF-8 encoded JSON objects, unless explicitly documented otherwise.  <div class=\"callout callout-error\">   <strong>🛑 Warning! Do not touch Photon!</strong><br>   Photon is only used by the in-game client and should <b>not</b> be touched. Doing so may result in permanent account termination. </div>  <div class=\"callout callout-info\">   <strong>ℹ️ Authentication</strong><br>   Read <a href=\"#tag--authentication\">Authentication</a> for how to log in. </div>  # Using the API  For simply exploring what the API can do it is strongly recommended to download [Insomnia](https://insomnia.rest/download), a free and open-source API client that's great for sending requests to the API in an orderly fashion. Insomnia allows you to send data in the format that's required for VRChat's API. It is also possible to try out the API in your browser, by first logging in at [vrchat.com/home](https://vrchat.com/home/) and then going to [vrchat.com/api/1/auth/user](https://vrchat.com/api/1/auth/user), but the information will be much harder to work with.  For more permanent operation such as software development it is instead recommended to use one of the existing language SDKs. This community project maintains API libraries in several languages, which allows you to interact with the API with simple function calls rather than having to implement the HTTP protocol yourself. Most of these libraries are automatically generated from the API specification, sometimes with additional helpful wrapper code to make usage easier. This allows them to be almost automatically updated and expanded upon as soon as a new feature is introduced in the specification itself. The libraries can be found on [GitHub](https://github.com/vrchatapi) or following:  * [NodeJS (JavaScript)](https://www.npmjs.com/package/vrchat) * [Dart](https://pub.dev/packages/vrchat_dart) * [Rust](https://crates.io/crates/vrchatapi) * [C#](https://github.com/vrchatapi/vrchatapi-csharp) * [Python](https://github.com/vrchatapi/vrchatapi-python)  # Pagination  Most endpoints enforce pagination, meaning they will only return 10 entries by default, and never more than 100.<br> Using both the limit and offset parameters allows you to easily paginate through a large number of objects.  | Query Parameter | Type | Description | | ----------|--|------- | | `n` | integer  | The number of objects to return. This value often defaults to 10. Highest limit is always 100.| | `offset` | integer  | A zero-based offset from the default object sorting.|  If a request returns fewer objects than the `limit` parameter, there are no more items available to return.  # Contribution  Do you want to get involved in the documentation effort? Do you want to help improve one of the language API libraries? This project is an [OPEN Open Source Project](https://openopensource.org)! This means that individuals making significant and valuable contributions are given commit-access to the project. It also means we are very open and welcoming of new people making contributions, unlike some more guarded open-source projects.  [![Discord](https://img.shields.io/static/v1?label=vrchatapi&message=discord&color=blueviolet&style=for-the-badge)](https://discord.gg/qjZE9C9fkB)

API version: 1.20.7
Contact: vrchatapi.lpv0t@aries.fyi
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vrchatapi

import (
	"encoding/json"
)

// checks if the UserNoteTargetUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserNoteTargetUser{}

// UserNoteTargetUser struct for UserNoteTargetUser
type UserNoteTargetUser struct {
	Id *string `json:"id,omitempty"`
	CurrentAvatarTags []string `json:"currentAvatarTags,omitempty"`
	// When profilePicOverride is not empty, use it instead.
	CurrentAvatarThumbnailImageUrl *string `json:"currentAvatarThumbnailImageUrl,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	ProfilePicOverride NullableString `json:"profilePicOverride,omitempty"`
	UserIcon *string `json:"userIcon,omitempty"`
}

// NewUserNoteTargetUser instantiates a new UserNoteTargetUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserNoteTargetUser() *UserNoteTargetUser {
	this := UserNoteTargetUser{}
	return &this
}

// NewUserNoteTargetUserWithDefaults instantiates a new UserNoteTargetUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserNoteTargetUserWithDefaults() *UserNoteTargetUser {
	this := UserNoteTargetUser{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserNoteTargetUser) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserNoteTargetUser) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserNoteTargetUser) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserNoteTargetUser) SetId(v string) {
	o.Id = &v
}

// GetCurrentAvatarTags returns the CurrentAvatarTags field value if set, zero value otherwise.
func (o *UserNoteTargetUser) GetCurrentAvatarTags() []string {
	if o == nil || IsNil(o.CurrentAvatarTags) {
		var ret []string
		return ret
	}
	return o.CurrentAvatarTags
}

// GetCurrentAvatarTagsOk returns a tuple with the CurrentAvatarTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserNoteTargetUser) GetCurrentAvatarTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.CurrentAvatarTags) {
		return nil, false
	}
	return o.CurrentAvatarTags, true
}

// HasCurrentAvatarTags returns a boolean if a field has been set.
func (o *UserNoteTargetUser) HasCurrentAvatarTags() bool {
	if o != nil && !IsNil(o.CurrentAvatarTags) {
		return true
	}

	return false
}

// SetCurrentAvatarTags gets a reference to the given []string and assigns it to the CurrentAvatarTags field.
func (o *UserNoteTargetUser) SetCurrentAvatarTags(v []string) {
	o.CurrentAvatarTags = v
}

// GetCurrentAvatarThumbnailImageUrl returns the CurrentAvatarThumbnailImageUrl field value if set, zero value otherwise.
func (o *UserNoteTargetUser) GetCurrentAvatarThumbnailImageUrl() string {
	if o == nil || IsNil(o.CurrentAvatarThumbnailImageUrl) {
		var ret string
		return ret
	}
	return *o.CurrentAvatarThumbnailImageUrl
}

// GetCurrentAvatarThumbnailImageUrlOk returns a tuple with the CurrentAvatarThumbnailImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserNoteTargetUser) GetCurrentAvatarThumbnailImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentAvatarThumbnailImageUrl) {
		return nil, false
	}
	return o.CurrentAvatarThumbnailImageUrl, true
}

// HasCurrentAvatarThumbnailImageUrl returns a boolean if a field has been set.
func (o *UserNoteTargetUser) HasCurrentAvatarThumbnailImageUrl() bool {
	if o != nil && !IsNil(o.CurrentAvatarThumbnailImageUrl) {
		return true
	}

	return false
}

// SetCurrentAvatarThumbnailImageUrl gets a reference to the given string and assigns it to the CurrentAvatarThumbnailImageUrl field.
func (o *UserNoteTargetUser) SetCurrentAvatarThumbnailImageUrl(v string) {
	o.CurrentAvatarThumbnailImageUrl = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *UserNoteTargetUser) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserNoteTargetUser) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UserNoteTargetUser) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *UserNoteTargetUser) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetProfilePicOverride returns the ProfilePicOverride field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserNoteTargetUser) GetProfilePicOverride() string {
	if o == nil || IsNil(o.ProfilePicOverride.Get()) {
		var ret string
		return ret
	}
	return *o.ProfilePicOverride.Get()
}

// GetProfilePicOverrideOk returns a tuple with the ProfilePicOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserNoteTargetUser) GetProfilePicOverrideOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfilePicOverride.Get(), o.ProfilePicOverride.IsSet()
}

// HasProfilePicOverride returns a boolean if a field has been set.
func (o *UserNoteTargetUser) HasProfilePicOverride() bool {
	if o != nil && o.ProfilePicOverride.IsSet() {
		return true
	}

	return false
}

// SetProfilePicOverride gets a reference to the given NullableString and assigns it to the ProfilePicOverride field.
func (o *UserNoteTargetUser) SetProfilePicOverride(v string) {
	o.ProfilePicOverride.Set(&v)
}
// SetProfilePicOverrideNil sets the value for ProfilePicOverride to be an explicit nil
func (o *UserNoteTargetUser) SetProfilePicOverrideNil() {
	o.ProfilePicOverride.Set(nil)
}

// UnsetProfilePicOverride ensures that no value is present for ProfilePicOverride, not even an explicit nil
func (o *UserNoteTargetUser) UnsetProfilePicOverride() {
	o.ProfilePicOverride.Unset()
}

// GetUserIcon returns the UserIcon field value if set, zero value otherwise.
func (o *UserNoteTargetUser) GetUserIcon() string {
	if o == nil || IsNil(o.UserIcon) {
		var ret string
		return ret
	}
	return *o.UserIcon
}

// GetUserIconOk returns a tuple with the UserIcon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserNoteTargetUser) GetUserIconOk() (*string, bool) {
	if o == nil || IsNil(o.UserIcon) {
		return nil, false
	}
	return o.UserIcon, true
}

// HasUserIcon returns a boolean if a field has been set.
func (o *UserNoteTargetUser) HasUserIcon() bool {
	if o != nil && !IsNil(o.UserIcon) {
		return true
	}

	return false
}

// SetUserIcon gets a reference to the given string and assigns it to the UserIcon field.
func (o *UserNoteTargetUser) SetUserIcon(v string) {
	o.UserIcon = &v
}

func (o UserNoteTargetUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserNoteTargetUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CurrentAvatarTags) {
		toSerialize["currentAvatarTags"] = o.CurrentAvatarTags
	}
	if !IsNil(o.CurrentAvatarThumbnailImageUrl) {
		toSerialize["currentAvatarThumbnailImageUrl"] = o.CurrentAvatarThumbnailImageUrl
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.ProfilePicOverride.IsSet() {
		toSerialize["profilePicOverride"] = o.ProfilePicOverride.Get()
	}
	if !IsNil(o.UserIcon) {
		toSerialize["userIcon"] = o.UserIcon
	}
	return toSerialize, nil
}

type NullableUserNoteTargetUser struct {
	value *UserNoteTargetUser
	isSet bool
}

func (v NullableUserNoteTargetUser) Get() *UserNoteTargetUser {
	return v.value
}

func (v *NullableUserNoteTargetUser) Set(val *UserNoteTargetUser) {
	v.value = val
	v.isSet = true
}

func (v NullableUserNoteTargetUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUserNoteTargetUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserNoteTargetUser(val *UserNoteTargetUser) *NullableUserNoteTargetUser {
	return &NullableUserNoteTargetUser{value: val, isSet: true}
}

func (v NullableUserNoteTargetUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserNoteTargetUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


