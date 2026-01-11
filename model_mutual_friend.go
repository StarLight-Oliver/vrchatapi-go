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

// checks if the MutualFriend type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MutualFriend{}

// MutualFriend User object received when querying mutual friends
type MutualFriend struct {
	// When profilePicOverride is not empty, use it instead.
	AvatarThumbnail *string `json:"avatarThumbnail,omitempty"`
	// When profilePicOverride is not empty, use it instead.
	CurrentAvatarImageUrl string   `json:"currentAvatarImageUrl"`
	CurrentAvatarTags     []string `json:"currentAvatarTags,omitempty"`
	// When profilePicOverride is not empty, use it instead.
	CurrentAvatarThumbnailImageUrl *string `json:"currentAvatarThumbnailImageUrl,omitempty"`
	DisplayName                    string  `json:"displayName"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	Id                 string     `json:"id"`
	ImageUrl           string     `json:"imageUrl"`
	ProfilePicOverride *string    `json:"profilePicOverride,omitempty"`
	Status             UserStatus `json:"status"`
	StatusDescription  string     `json:"statusDescription"`
}

type _MutualFriend MutualFriend

// NewMutualFriend instantiates a new MutualFriend object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMutualFriend(currentAvatarImageUrl string, displayName string, id string, imageUrl string, status UserStatus, statusDescription string) *MutualFriend {
	this := MutualFriend{}
	this.CurrentAvatarImageUrl = currentAvatarImageUrl
	this.DisplayName = displayName
	this.Id = id
	this.ImageUrl = imageUrl
	this.Status = status
	this.StatusDescription = statusDescription
	return &this
}

// NewMutualFriendWithDefaults instantiates a new MutualFriend object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMutualFriendWithDefaults() *MutualFriend {
	this := MutualFriend{}
	var status UserStatus = UserStatus_OFFLINE
	this.Status = status
	return &this
}

// GetAvatarThumbnail returns the AvatarThumbnail field value if set, zero value otherwise.
func (o *MutualFriend) GetAvatarThumbnail() string {
	if o == nil || IsNil(o.AvatarThumbnail) {
		var ret string
		return ret
	}
	return *o.AvatarThumbnail
}

// GetAvatarThumbnailOk returns a tuple with the AvatarThumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualFriend) GetAvatarThumbnailOk() (*string, bool) {
	if o == nil || IsNil(o.AvatarThumbnail) {
		return nil, false
	}
	return o.AvatarThumbnail, true
}

// HasAvatarThumbnail returns a boolean if a field has been set.
func (o *MutualFriend) HasAvatarThumbnail() bool {
	if o != nil && !IsNil(o.AvatarThumbnail) {
		return true
	}

	return false
}

// SetAvatarThumbnail gets a reference to the given string and assigns it to the AvatarThumbnail field.
func (o *MutualFriend) SetAvatarThumbnail(v string) {
	o.AvatarThumbnail = &v
}

// GetCurrentAvatarImageUrl returns the CurrentAvatarImageUrl field value
func (o *MutualFriend) GetCurrentAvatarImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentAvatarImageUrl
}

// GetCurrentAvatarImageUrlOk returns a tuple with the CurrentAvatarImageUrl field value
// and a boolean to check if the value has been set.
func (o *MutualFriend) GetCurrentAvatarImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentAvatarImageUrl, true
}

// SetCurrentAvatarImageUrl sets field value
func (o *MutualFriend) SetCurrentAvatarImageUrl(v string) {
	o.CurrentAvatarImageUrl = v
}

// GetCurrentAvatarTags returns the CurrentAvatarTags field value if set, zero value otherwise.
func (o *MutualFriend) GetCurrentAvatarTags() []string {
	if o == nil || IsNil(o.CurrentAvatarTags) {
		var ret []string
		return ret
	}
	return o.CurrentAvatarTags
}

// GetCurrentAvatarTagsOk returns a tuple with the CurrentAvatarTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualFriend) GetCurrentAvatarTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.CurrentAvatarTags) {
		return nil, false
	}
	return o.CurrentAvatarTags, true
}

// HasCurrentAvatarTags returns a boolean if a field has been set.
func (o *MutualFriend) HasCurrentAvatarTags() bool {
	if o != nil && !IsNil(o.CurrentAvatarTags) {
		return true
	}

	return false
}

// SetCurrentAvatarTags gets a reference to the given []string and assigns it to the CurrentAvatarTags field.
func (o *MutualFriend) SetCurrentAvatarTags(v []string) {
	o.CurrentAvatarTags = v
}

// GetCurrentAvatarThumbnailImageUrl returns the CurrentAvatarThumbnailImageUrl field value if set, zero value otherwise.
func (o *MutualFriend) GetCurrentAvatarThumbnailImageUrl() string {
	if o == nil || IsNil(o.CurrentAvatarThumbnailImageUrl) {
		var ret string
		return ret
	}
	return *o.CurrentAvatarThumbnailImageUrl
}

// GetCurrentAvatarThumbnailImageUrlOk returns a tuple with the CurrentAvatarThumbnailImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualFriend) GetCurrentAvatarThumbnailImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentAvatarThumbnailImageUrl) {
		return nil, false
	}
	return o.CurrentAvatarThumbnailImageUrl, true
}

// HasCurrentAvatarThumbnailImageUrl returns a boolean if a field has been set.
func (o *MutualFriend) HasCurrentAvatarThumbnailImageUrl() bool {
	if o != nil && !IsNil(o.CurrentAvatarThumbnailImageUrl) {
		return true
	}

	return false
}

// SetCurrentAvatarThumbnailImageUrl gets a reference to the given string and assigns it to the CurrentAvatarThumbnailImageUrl field.
func (o *MutualFriend) SetCurrentAvatarThumbnailImageUrl(v string) {
	o.CurrentAvatarThumbnailImageUrl = &v
}

// GetDisplayName returns the DisplayName field value
func (o *MutualFriend) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *MutualFriend) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *MutualFriend) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetId returns the Id field value
func (o *MutualFriend) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MutualFriend) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MutualFriend) SetId(v string) {
	o.Id = v
}

// GetImageUrl returns the ImageUrl field value
func (o *MutualFriend) GetImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
func (o *MutualFriend) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageUrl, true
}

// SetImageUrl sets field value
func (o *MutualFriend) SetImageUrl(v string) {
	o.ImageUrl = v
}

// GetProfilePicOverride returns the ProfilePicOverride field value if set, zero value otherwise.
func (o *MutualFriend) GetProfilePicOverride() string {
	if o == nil || IsNil(o.ProfilePicOverride) {
		var ret string
		return ret
	}
	return *o.ProfilePicOverride
}

// GetProfilePicOverrideOk returns a tuple with the ProfilePicOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualFriend) GetProfilePicOverrideOk() (*string, bool) {
	if o == nil || IsNil(o.ProfilePicOverride) {
		return nil, false
	}
	return o.ProfilePicOverride, true
}

// HasProfilePicOverride returns a boolean if a field has been set.
func (o *MutualFriend) HasProfilePicOverride() bool {
	if o != nil && !IsNil(o.ProfilePicOverride) {
		return true
	}

	return false
}

// SetProfilePicOverride gets a reference to the given string and assigns it to the ProfilePicOverride field.
func (o *MutualFriend) SetProfilePicOverride(v string) {
	o.ProfilePicOverride = &v
}

// GetStatus returns the Status field value
func (o *MutualFriend) GetStatus() UserStatus {
	if o == nil {
		var ret UserStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *MutualFriend) GetStatusOk() (*UserStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *MutualFriend) SetStatus(v UserStatus) {
	o.Status = v
}

// GetStatusDescription returns the StatusDescription field value
func (o *MutualFriend) GetStatusDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatusDescription
}

// GetStatusDescriptionOk returns a tuple with the StatusDescription field value
// and a boolean to check if the value has been set.
func (o *MutualFriend) GetStatusDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusDescription, true
}

// SetStatusDescription sets field value
func (o *MutualFriend) SetStatusDescription(v string) {
	o.StatusDescription = v
}

func (o MutualFriend) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MutualFriend) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvatarThumbnail) {
		toSerialize["avatarThumbnail"] = o.AvatarThumbnail
	}
	toSerialize["currentAvatarImageUrl"] = o.CurrentAvatarImageUrl
	if !IsNil(o.CurrentAvatarTags) {
		toSerialize["currentAvatarTags"] = o.CurrentAvatarTags
	}
	if !IsNil(o.CurrentAvatarThumbnailImageUrl) {
		toSerialize["currentAvatarThumbnailImageUrl"] = o.CurrentAvatarThumbnailImageUrl
	}
	toSerialize["displayName"] = o.DisplayName
	toSerialize["id"] = o.Id
	toSerialize["imageUrl"] = o.ImageUrl
	if !IsNil(o.ProfilePicOverride) {
		toSerialize["profilePicOverride"] = o.ProfilePicOverride
	}
	toSerialize["status"] = o.Status
	toSerialize["statusDescription"] = o.StatusDescription
	return toSerialize, nil
}

func (o *MutualFriend) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"currentAvatarImageUrl",
		"displayName",
		"id",
		"imageUrl",
		"status",
		"statusDescription",
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

	varMutualFriend := _MutualFriend{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMutualFriend)

	if err != nil {
		return err
	}

	*o = MutualFriend(varMutualFriend)

	return err
}

type NullableMutualFriend struct {
	value *MutualFriend
	isSet bool
}

func (v NullableMutualFriend) Get() *MutualFriend {
	return v.value
}

func (v *NullableMutualFriend) Set(val *MutualFriend) {
	v.value = val
	v.isSet = true
}

func (v NullableMutualFriend) IsSet() bool {
	return v.isSet
}

func (v *NullableMutualFriend) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMutualFriend(val *MutualFriend) *NullableMutualFriend {
	return &NullableMutualFriend{value: val, isSet: true}
}

func (v NullableMutualFriend) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMutualFriend) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
