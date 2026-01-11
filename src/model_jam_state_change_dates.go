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
)

// checks if the JamStateChangeDates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JamStateChangeDates{}

// JamStateChangeDates struct for JamStateChangeDates
type JamStateChangeDates struct {
	Closed NullableTime `json:"closed,omitempty"`
	SubmissionsClosed NullableTime `json:"submissionsClosed,omitempty"`
	SubmissionsOpened NullableTime `json:"submissionsOpened,omitempty"`
	WinnersSelected NullableTime `json:"winnersSelected,omitempty"`
}

// NewJamStateChangeDates instantiates a new JamStateChangeDates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJamStateChangeDates() *JamStateChangeDates {
	this := JamStateChangeDates{}
	return &this
}

// NewJamStateChangeDatesWithDefaults instantiates a new JamStateChangeDates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJamStateChangeDatesWithDefaults() *JamStateChangeDates {
	this := JamStateChangeDates{}
	return &this
}

// GetClosed returns the Closed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JamStateChangeDates) GetClosed() time.Time {
	if o == nil || IsNil(o.Closed.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Closed.Get()
}

// GetClosedOk returns a tuple with the Closed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JamStateChangeDates) GetClosedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Closed.Get(), o.Closed.IsSet()
}

// HasClosed returns a boolean if a field has been set.
func (o *JamStateChangeDates) HasClosed() bool {
	if o != nil && o.Closed.IsSet() {
		return true
	}

	return false
}

// SetClosed gets a reference to the given NullableTime and assigns it to the Closed field.
func (o *JamStateChangeDates) SetClosed(v time.Time) {
	o.Closed.Set(&v)
}
// SetClosedNil sets the value for Closed to be an explicit nil
func (o *JamStateChangeDates) SetClosedNil() {
	o.Closed.Set(nil)
}

// UnsetClosed ensures that no value is present for Closed, not even an explicit nil
func (o *JamStateChangeDates) UnsetClosed() {
	o.Closed.Unset()
}

// GetSubmissionsClosed returns the SubmissionsClosed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JamStateChangeDates) GetSubmissionsClosed() time.Time {
	if o == nil || IsNil(o.SubmissionsClosed.Get()) {
		var ret time.Time
		return ret
	}
	return *o.SubmissionsClosed.Get()
}

// GetSubmissionsClosedOk returns a tuple with the SubmissionsClosed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JamStateChangeDates) GetSubmissionsClosedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubmissionsClosed.Get(), o.SubmissionsClosed.IsSet()
}

// HasSubmissionsClosed returns a boolean if a field has been set.
func (o *JamStateChangeDates) HasSubmissionsClosed() bool {
	if o != nil && o.SubmissionsClosed.IsSet() {
		return true
	}

	return false
}

// SetSubmissionsClosed gets a reference to the given NullableTime and assigns it to the SubmissionsClosed field.
func (o *JamStateChangeDates) SetSubmissionsClosed(v time.Time) {
	o.SubmissionsClosed.Set(&v)
}
// SetSubmissionsClosedNil sets the value for SubmissionsClosed to be an explicit nil
func (o *JamStateChangeDates) SetSubmissionsClosedNil() {
	o.SubmissionsClosed.Set(nil)
}

// UnsetSubmissionsClosed ensures that no value is present for SubmissionsClosed, not even an explicit nil
func (o *JamStateChangeDates) UnsetSubmissionsClosed() {
	o.SubmissionsClosed.Unset()
}

// GetSubmissionsOpened returns the SubmissionsOpened field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JamStateChangeDates) GetSubmissionsOpened() time.Time {
	if o == nil || IsNil(o.SubmissionsOpened.Get()) {
		var ret time.Time
		return ret
	}
	return *o.SubmissionsOpened.Get()
}

// GetSubmissionsOpenedOk returns a tuple with the SubmissionsOpened field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JamStateChangeDates) GetSubmissionsOpenedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubmissionsOpened.Get(), o.SubmissionsOpened.IsSet()
}

// HasSubmissionsOpened returns a boolean if a field has been set.
func (o *JamStateChangeDates) HasSubmissionsOpened() bool {
	if o != nil && o.SubmissionsOpened.IsSet() {
		return true
	}

	return false
}

// SetSubmissionsOpened gets a reference to the given NullableTime and assigns it to the SubmissionsOpened field.
func (o *JamStateChangeDates) SetSubmissionsOpened(v time.Time) {
	o.SubmissionsOpened.Set(&v)
}
// SetSubmissionsOpenedNil sets the value for SubmissionsOpened to be an explicit nil
func (o *JamStateChangeDates) SetSubmissionsOpenedNil() {
	o.SubmissionsOpened.Set(nil)
}

// UnsetSubmissionsOpened ensures that no value is present for SubmissionsOpened, not even an explicit nil
func (o *JamStateChangeDates) UnsetSubmissionsOpened() {
	o.SubmissionsOpened.Unset()
}

// GetWinnersSelected returns the WinnersSelected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JamStateChangeDates) GetWinnersSelected() time.Time {
	if o == nil || IsNil(o.WinnersSelected.Get()) {
		var ret time.Time
		return ret
	}
	return *o.WinnersSelected.Get()
}

// GetWinnersSelectedOk returns a tuple with the WinnersSelected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JamStateChangeDates) GetWinnersSelectedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.WinnersSelected.Get(), o.WinnersSelected.IsSet()
}

// HasWinnersSelected returns a boolean if a field has been set.
func (o *JamStateChangeDates) HasWinnersSelected() bool {
	if o != nil && o.WinnersSelected.IsSet() {
		return true
	}

	return false
}

// SetWinnersSelected gets a reference to the given NullableTime and assigns it to the WinnersSelected field.
func (o *JamStateChangeDates) SetWinnersSelected(v time.Time) {
	o.WinnersSelected.Set(&v)
}
// SetWinnersSelectedNil sets the value for WinnersSelected to be an explicit nil
func (o *JamStateChangeDates) SetWinnersSelectedNil() {
	o.WinnersSelected.Set(nil)
}

// UnsetWinnersSelected ensures that no value is present for WinnersSelected, not even an explicit nil
func (o *JamStateChangeDates) UnsetWinnersSelected() {
	o.WinnersSelected.Unset()
}

func (o JamStateChangeDates) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JamStateChangeDates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Closed.IsSet() {
		toSerialize["closed"] = o.Closed.Get()
	}
	if o.SubmissionsClosed.IsSet() {
		toSerialize["submissionsClosed"] = o.SubmissionsClosed.Get()
	}
	if o.SubmissionsOpened.IsSet() {
		toSerialize["submissionsOpened"] = o.SubmissionsOpened.Get()
	}
	if o.WinnersSelected.IsSet() {
		toSerialize["winnersSelected"] = o.WinnersSelected.Get()
	}
	return toSerialize, nil
}

type NullableJamStateChangeDates struct {
	value *JamStateChangeDates
	isSet bool
}

func (v NullableJamStateChangeDates) Get() *JamStateChangeDates {
	return v.value
}

func (v *NullableJamStateChangeDates) Set(val *JamStateChangeDates) {
	v.value = val
	v.isSet = true
}

func (v NullableJamStateChangeDates) IsSet() bool {
	return v.isSet
}

func (v *NullableJamStateChangeDates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJamStateChangeDates(val *JamStateChangeDates) *NullableJamStateChangeDates {
	return &NullableJamStateChangeDates{value: val, isSet: true}
}

func (v NullableJamStateChangeDates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJamStateChangeDates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


