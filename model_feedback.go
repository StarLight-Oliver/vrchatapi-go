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
	"bytes"
	"fmt"
)

// checks if the Feedback type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Feedback{}

// Feedback struct for Feedback
type Feedback struct {
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	CommenterId string `json:"commenterId"`
	CommenterName string `json:"commenterName"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	ContentAuthorId string `json:"contentAuthorId"`
	ContentAuthorName NullableString `json:"contentAuthorName"`
	ContentId string `json:"contentId"`
	ContentName *string `json:"contentName,omitempty"`
	ContentType string `json:"contentType"`
	ContentVersion NullableInt32 `json:"contentVersion"`
	Description NullableString `json:"description,omitempty"`
	Id string `json:"id"`
	Reason string `json:"reason"`
	Tags []string `json:"tags"`
	Type string `json:"type"`
}

type _Feedback Feedback

// NewFeedback instantiates a new Feedback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedback(commenterId string, commenterName string, contentAuthorId string, contentAuthorName NullableString, contentId string, contentType string, contentVersion NullableInt32, id string, reason string, tags []string, type_ string) *Feedback {
	this := Feedback{}
	this.CommenterId = commenterId
	this.CommenterName = commenterName
	this.ContentAuthorId = contentAuthorId
	this.ContentAuthorName = contentAuthorName
	this.ContentId = contentId
	this.ContentType = contentType
	this.ContentVersion = contentVersion
	this.Id = id
	this.Reason = reason
	this.Tags = tags
	this.Type = type_
	return &this
}

// NewFeedbackWithDefaults instantiates a new Feedback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedbackWithDefaults() *Feedback {
	this := Feedback{}
	return &this
}

// GetCommenterId returns the CommenterId field value
func (o *Feedback) GetCommenterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommenterId
}

// GetCommenterIdOk returns a tuple with the CommenterId field value
// and a boolean to check if the value has been set.
func (o *Feedback) GetCommenterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommenterId, true
}

// SetCommenterId sets field value
func (o *Feedback) SetCommenterId(v string) {
	o.CommenterId = v
}

// GetCommenterName returns the CommenterName field value
func (o *Feedback) GetCommenterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommenterName
}

// GetCommenterNameOk returns a tuple with the CommenterName field value
// and a boolean to check if the value has been set.
func (o *Feedback) GetCommenterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommenterName, true
}

// SetCommenterName sets field value
func (o *Feedback) SetCommenterName(v string) {
	o.CommenterName = v
}

// GetContentAuthorId returns the ContentAuthorId field value
func (o *Feedback) GetContentAuthorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentAuthorId
}

// GetContentAuthorIdOk returns a tuple with the ContentAuthorId field value
// and a boolean to check if the value has been set.
func (o *Feedback) GetContentAuthorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentAuthorId, true
}

// SetContentAuthorId sets field value
func (o *Feedback) SetContentAuthorId(v string) {
	o.ContentAuthorId = v
}

// GetContentAuthorName returns the ContentAuthorName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Feedback) GetContentAuthorName() string {
	if o == nil || o.ContentAuthorName.Get() == nil {
		var ret string
		return ret
	}

	return *o.ContentAuthorName.Get()
}

// GetContentAuthorNameOk returns a tuple with the ContentAuthorName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Feedback) GetContentAuthorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentAuthorName.Get(), o.ContentAuthorName.IsSet()
}

// SetContentAuthorName sets field value
func (o *Feedback) SetContentAuthorName(v string) {
	o.ContentAuthorName.Set(&v)
}

// GetContentId returns the ContentId field value
func (o *Feedback) GetContentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentId
}

// GetContentIdOk returns a tuple with the ContentId field value
// and a boolean to check if the value has been set.
func (o *Feedback) GetContentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentId, true
}

// SetContentId sets field value
func (o *Feedback) SetContentId(v string) {
	o.ContentId = v
}

// GetContentName returns the ContentName field value if set, zero value otherwise.
func (o *Feedback) GetContentName() string {
	if o == nil || IsNil(o.ContentName) {
		var ret string
		return ret
	}
	return *o.ContentName
}

// GetContentNameOk returns a tuple with the ContentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Feedback) GetContentNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContentName) {
		return nil, false
	}
	return o.ContentName, true
}

// HasContentName returns a boolean if a field has been set.
func (o *Feedback) HasContentName() bool {
	if o != nil && !IsNil(o.ContentName) {
		return true
	}

	return false
}

// SetContentName gets a reference to the given string and assigns it to the ContentName field.
func (o *Feedback) SetContentName(v string) {
	o.ContentName = &v
}

// GetContentType returns the ContentType field value
func (o *Feedback) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *Feedback) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *Feedback) SetContentType(v string) {
	o.ContentType = v
}

// GetContentVersion returns the ContentVersion field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Feedback) GetContentVersion() int32 {
	if o == nil || o.ContentVersion.Get() == nil {
		var ret int32
		return ret
	}

	return *o.ContentVersion.Get()
}

// GetContentVersionOk returns a tuple with the ContentVersion field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Feedback) GetContentVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentVersion.Get(), o.ContentVersion.IsSet()
}

// SetContentVersion sets field value
func (o *Feedback) SetContentVersion(v int32) {
	o.ContentVersion.Set(&v)
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Feedback) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Feedback) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Feedback) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Feedback) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Feedback) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Feedback) UnsetDescription() {
	o.Description.Unset()
}

// GetId returns the Id field value
func (o *Feedback) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Feedback) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Feedback) SetId(v string) {
	o.Id = v
}

// GetReason returns the Reason field value
func (o *Feedback) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *Feedback) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *Feedback) SetReason(v string) {
	o.Reason = v
}

// GetTags returns the Tags field value
func (o *Feedback) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *Feedback) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *Feedback) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value
func (o *Feedback) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Feedback) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Feedback) SetType(v string) {
	o.Type = v
}

func (o Feedback) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Feedback) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["commenterId"] = o.CommenterId
	toSerialize["commenterName"] = o.CommenterName
	toSerialize["contentAuthorId"] = o.ContentAuthorId
	toSerialize["contentAuthorName"] = o.ContentAuthorName.Get()
	toSerialize["contentId"] = o.ContentId
	if !IsNil(o.ContentName) {
		toSerialize["contentName"] = o.ContentName
	}
	toSerialize["contentType"] = o.ContentType
	toSerialize["contentVersion"] = o.ContentVersion.Get()
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["id"] = o.Id
	toSerialize["reason"] = o.Reason
	toSerialize["tags"] = o.Tags
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *Feedback) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"commenterId",
		"commenterName",
		"contentAuthorId",
		"contentAuthorName",
		"contentId",
		"contentType",
		"contentVersion",
		"id",
		"reason",
		"tags",
		"type",
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

	varFeedback := _Feedback{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFeedback)

	if err != nil {
		return err
	}

	*o = Feedback(varFeedback)

	return err
}

type NullableFeedback struct {
	value *Feedback
	isSet bool
}

func (v NullableFeedback) Get() *Feedback {
	return v.value
}

func (v *NullableFeedback) Set(val *Feedback) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedback) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedback(val *Feedback) *NullableFeedback {
	return &NullableFeedback{value: val, isSet: true}
}

func (v NullableFeedback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


