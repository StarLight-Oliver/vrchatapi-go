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
	"time"
	"bytes"
	"fmt"
)

// checks if the Badge type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Badge{}

// Badge struct for Badge
type Badge struct {
	// only present in CurrentUser badges
	AssignedAt NullableTime `json:"assignedAt,omitempty"`
	BadgeDescription string `json:"badgeDescription"`
	BadgeId string `json:"badgeId"`
	// direct url to image
	BadgeImageUrl string `json:"badgeImageUrl"`
	BadgeName string `json:"badgeName"`
	// only present in CurrentUser badges
	Hidden NullableBool `json:"hidden,omitempty"`
	Showcased bool `json:"showcased"`
	// only present in CurrentUser badges
	UpdatedAt NullableTime `json:"updatedAt,omitempty"`
}

type _Badge Badge

// NewBadge instantiates a new Badge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBadge(badgeDescription string, badgeId string, badgeImageUrl string, badgeName string, showcased bool) *Badge {
	this := Badge{}
	this.BadgeDescription = badgeDescription
	this.BadgeId = badgeId
	this.BadgeImageUrl = badgeImageUrl
	this.BadgeName = badgeName
	this.Showcased = showcased
	return &this
}

// NewBadgeWithDefaults instantiates a new Badge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBadgeWithDefaults() *Badge {
	this := Badge{}
	return &this
}

// GetAssignedAt returns the AssignedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Badge) GetAssignedAt() time.Time {
	if o == nil || IsNil(o.AssignedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.AssignedAt.Get()
}

// GetAssignedAtOk returns a tuple with the AssignedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Badge) GetAssignedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedAt.Get(), o.AssignedAt.IsSet()
}

// HasAssignedAt returns a boolean if a field has been set.
func (o *Badge) HasAssignedAt() bool {
	if o != nil && o.AssignedAt.IsSet() {
		return true
	}

	return false
}

// SetAssignedAt gets a reference to the given NullableTime and assigns it to the AssignedAt field.
func (o *Badge) SetAssignedAt(v time.Time) {
	o.AssignedAt.Set(&v)
}
// SetAssignedAtNil sets the value for AssignedAt to be an explicit nil
func (o *Badge) SetAssignedAtNil() {
	o.AssignedAt.Set(nil)
}

// UnsetAssignedAt ensures that no value is present for AssignedAt, not even an explicit nil
func (o *Badge) UnsetAssignedAt() {
	o.AssignedAt.Unset()
}

// GetBadgeDescription returns the BadgeDescription field value
func (o *Badge) GetBadgeDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BadgeDescription
}

// GetBadgeDescriptionOk returns a tuple with the BadgeDescription field value
// and a boolean to check if the value has been set.
func (o *Badge) GetBadgeDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BadgeDescription, true
}

// SetBadgeDescription sets field value
func (o *Badge) SetBadgeDescription(v string) {
	o.BadgeDescription = v
}

// GetBadgeId returns the BadgeId field value
func (o *Badge) GetBadgeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BadgeId
}

// GetBadgeIdOk returns a tuple with the BadgeId field value
// and a boolean to check if the value has been set.
func (o *Badge) GetBadgeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BadgeId, true
}

// SetBadgeId sets field value
func (o *Badge) SetBadgeId(v string) {
	o.BadgeId = v
}

// GetBadgeImageUrl returns the BadgeImageUrl field value
func (o *Badge) GetBadgeImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BadgeImageUrl
}

// GetBadgeImageUrlOk returns a tuple with the BadgeImageUrl field value
// and a boolean to check if the value has been set.
func (o *Badge) GetBadgeImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BadgeImageUrl, true
}

// SetBadgeImageUrl sets field value
func (o *Badge) SetBadgeImageUrl(v string) {
	o.BadgeImageUrl = v
}

// GetBadgeName returns the BadgeName field value
func (o *Badge) GetBadgeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BadgeName
}

// GetBadgeNameOk returns a tuple with the BadgeName field value
// and a boolean to check if the value has been set.
func (o *Badge) GetBadgeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BadgeName, true
}

// SetBadgeName sets field value
func (o *Badge) SetBadgeName(v string) {
	o.BadgeName = v
}

// GetHidden returns the Hidden field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Badge) GetHidden() bool {
	if o == nil || IsNil(o.Hidden.Get()) {
		var ret bool
		return ret
	}
	return *o.Hidden.Get()
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Badge) GetHiddenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hidden.Get(), o.Hidden.IsSet()
}

// HasHidden returns a boolean if a field has been set.
func (o *Badge) HasHidden() bool {
	if o != nil && o.Hidden.IsSet() {
		return true
	}

	return false
}

// SetHidden gets a reference to the given NullableBool and assigns it to the Hidden field.
func (o *Badge) SetHidden(v bool) {
	o.Hidden.Set(&v)
}
// SetHiddenNil sets the value for Hidden to be an explicit nil
func (o *Badge) SetHiddenNil() {
	o.Hidden.Set(nil)
}

// UnsetHidden ensures that no value is present for Hidden, not even an explicit nil
func (o *Badge) UnsetHidden() {
	o.Hidden.Unset()
}

// GetShowcased returns the Showcased field value
func (o *Badge) GetShowcased() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Showcased
}

// GetShowcasedOk returns a tuple with the Showcased field value
// and a boolean to check if the value has been set.
func (o *Badge) GetShowcasedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Showcased, true
}

// SetShowcased sets field value
func (o *Badge) SetShowcased(v bool) {
	o.Showcased = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Badge) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Badge) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Badge) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *Badge) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *Badge) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *Badge) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

func (o Badge) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Badge) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AssignedAt.IsSet() {
		toSerialize["assignedAt"] = o.AssignedAt.Get()
	}
	toSerialize["badgeDescription"] = o.BadgeDescription
	toSerialize["badgeId"] = o.BadgeId
	toSerialize["badgeImageUrl"] = o.BadgeImageUrl
	toSerialize["badgeName"] = o.BadgeName
	if o.Hidden.IsSet() {
		toSerialize["hidden"] = o.Hidden.Get()
	}
	toSerialize["showcased"] = o.Showcased
	if o.UpdatedAt.IsSet() {
		toSerialize["updatedAt"] = o.UpdatedAt.Get()
	}
	return toSerialize, nil
}

func (o *Badge) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"badgeDescription",
		"badgeId",
		"badgeImageUrl",
		"badgeName",
		"showcased",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBadge := _Badge{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBadge)

	if err != nil {
		return err
	}

	*o = Badge(varBadge)

	return err
}

type NullableBadge struct {
	value *Badge
	isSet bool
}

func (v NullableBadge) Get() *Badge {
	return v.value
}

func (v *NullableBadge) Set(val *Badge) {
	v.value = val
	v.isSet = true
}

func (v NullableBadge) IsSet() bool {
	return v.isSet
}

func (v *NullableBadge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBadge(val *Badge) *NullableBadge {
	return &NullableBadge{value: val, isSet: true}
}

func (v NullableBadge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBadge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


