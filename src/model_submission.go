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

// checks if the Submission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Submission{}

// Submission 
type Submission struct {
	// Either world ID or avatar ID
	ContentId string `json:"contentId"`
	CreatedAt time.Time `json:"created_at"`
	Description string `json:"description"`
	Id string `json:"id"`
	JamId string `json:"jamId"`
	RatingScore *int32 `json:"ratingScore,omitempty"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	SubmitterId string `json:"submitterId"`
}

type _Submission Submission

// NewSubmission instantiates a new Submission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmission(contentId string, createdAt time.Time, description string, id string, jamId string, submitterId string) *Submission {
	this := Submission{}
	this.ContentId = contentId
	this.CreatedAt = createdAt
	this.Description = description
	this.Id = id
	this.JamId = jamId
	this.SubmitterId = submitterId
	return &this
}

// NewSubmissionWithDefaults instantiates a new Submission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmissionWithDefaults() *Submission {
	this := Submission{}
	return &this
}

// GetContentId returns the ContentId field value
func (o *Submission) GetContentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentId
}

// GetContentIdOk returns a tuple with the ContentId field value
// and a boolean to check if the value has been set.
func (o *Submission) GetContentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentId, true
}

// SetContentId sets field value
func (o *Submission) SetContentId(v string) {
	o.ContentId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Submission) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Submission) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Submission) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value
func (o *Submission) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Submission) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Submission) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value
func (o *Submission) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Submission) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Submission) SetId(v string) {
	o.Id = v
}

// GetJamId returns the JamId field value
func (o *Submission) GetJamId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JamId
}

// GetJamIdOk returns a tuple with the JamId field value
// and a boolean to check if the value has been set.
func (o *Submission) GetJamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JamId, true
}

// SetJamId sets field value
func (o *Submission) SetJamId(v string) {
	o.JamId = v
}

// GetRatingScore returns the RatingScore field value if set, zero value otherwise.
func (o *Submission) GetRatingScore() int32 {
	if o == nil || IsNil(o.RatingScore) {
		var ret int32
		return ret
	}
	return *o.RatingScore
}

// GetRatingScoreOk returns a tuple with the RatingScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Submission) GetRatingScoreOk() (*int32, bool) {
	if o == nil || IsNil(o.RatingScore) {
		return nil, false
	}
	return o.RatingScore, true
}

// HasRatingScore returns a boolean if a field has been set.
func (o *Submission) HasRatingScore() bool {
	if o != nil && !IsNil(o.RatingScore) {
		return true
	}

	return false
}

// SetRatingScore gets a reference to the given int32 and assigns it to the RatingScore field.
func (o *Submission) SetRatingScore(v int32) {
	o.RatingScore = &v
}

// GetSubmitterId returns the SubmitterId field value
func (o *Submission) GetSubmitterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubmitterId
}

// GetSubmitterIdOk returns a tuple with the SubmitterId field value
// and a boolean to check if the value has been set.
func (o *Submission) GetSubmitterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubmitterId, true
}

// SetSubmitterId sets field value
func (o *Submission) SetSubmitterId(v string) {
	o.SubmitterId = v
}

func (o Submission) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Submission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contentId"] = o.ContentId
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["description"] = o.Description
	toSerialize["id"] = o.Id
	toSerialize["jamId"] = o.JamId
	if !IsNil(o.RatingScore) {
		toSerialize["ratingScore"] = o.RatingScore
	}
	toSerialize["submitterId"] = o.SubmitterId
	return toSerialize, nil
}

func (o *Submission) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"contentId",
		"created_at",
		"description",
		"id",
		"jamId",
		"submitterId",
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

	varSubmission := _Submission{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubmission)

	if err != nil {
		return err
	}

	*o = Submission(varSubmission)

	return err
}

type NullableSubmission struct {
	value *Submission
	isSet bool
}

func (v NullableSubmission) Get() *Submission {
	return v.value
}

func (v *NullableSubmission) Set(val *Submission) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmission) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmission(val *Submission) *NullableSubmission {
	return &NullableSubmission{value: val, isSet: true}
}

func (v NullableSubmission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


