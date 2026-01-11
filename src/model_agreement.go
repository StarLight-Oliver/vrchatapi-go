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

// checks if the Agreement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Agreement{}

// Agreement struct for Agreement
type Agreement struct {
	AgreementCode AgreementCode `json:"agreementCode"`
	// The full text of the agreement.
	AgreementFulltext *string `json:"agreementFulltext,omitempty"`
	// The id of the content being uploaded, such as a WorldID, AvatarID, or PropID.
	ContentId string `json:"contentId"`
	// When the agreement was created
	Created string `json:"created"`
	// The id of the agreement.
	Id   string   `json:"id"`
	Tags []string `json:"tags"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	UserId string `json:"userId"`
	// The version of the agreement.
	Version int32 `json:"version"`
}

type _Agreement Agreement

// NewAgreement instantiates a new Agreement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgreement(agreementCode AgreementCode, contentId string, created string, id string, tags []string, userId string, version int32) *Agreement {
	this := Agreement{}
	this.AgreementCode = agreementCode
	this.ContentId = contentId
	this.Created = created
	this.Id = id
	this.Tags = tags
	this.UserId = userId
	this.Version = version
	return &this
}

// NewAgreementWithDefaults instantiates a new Agreement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgreementWithDefaults() *Agreement {
	this := Agreement{}
	var agreementCode AgreementCode = CONTENT_COPYRIGHT_OWNED
	this.AgreementCode = agreementCode
	return &this
}

// GetAgreementCode returns the AgreementCode field value
func (o *Agreement) GetAgreementCode() AgreementCode {
	if o == nil {
		var ret AgreementCode
		return ret
	}

	return o.AgreementCode
}

// GetAgreementCodeOk returns a tuple with the AgreementCode field value
// and a boolean to check if the value has been set.
func (o *Agreement) GetAgreementCodeOk() (*AgreementCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgreementCode, true
}

// SetAgreementCode sets field value
func (o *Agreement) SetAgreementCode(v AgreementCode) {
	o.AgreementCode = v
}

// GetAgreementFulltext returns the AgreementFulltext field value if set, zero value otherwise.
func (o *Agreement) GetAgreementFulltext() string {
	if o == nil || IsNil(o.AgreementFulltext) {
		var ret string
		return ret
	}
	return *o.AgreementFulltext
}

// GetAgreementFulltextOk returns a tuple with the AgreementFulltext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agreement) GetAgreementFulltextOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementFulltext) {
		return nil, false
	}
	return o.AgreementFulltext, true
}

// HasAgreementFulltext returns a boolean if a field has been set.
func (o *Agreement) HasAgreementFulltext() bool {
	if o != nil && !IsNil(o.AgreementFulltext) {
		return true
	}

	return false
}

// SetAgreementFulltext gets a reference to the given string and assigns it to the AgreementFulltext field.
func (o *Agreement) SetAgreementFulltext(v string) {
	o.AgreementFulltext = &v
}

// GetContentId returns the ContentId field value
func (o *Agreement) GetContentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentId
}

// GetContentIdOk returns a tuple with the ContentId field value
// and a boolean to check if the value has been set.
func (o *Agreement) GetContentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentId, true
}

// SetContentId sets field value
func (o *Agreement) SetContentId(v string) {
	o.ContentId = v
}

// GetCreated returns the Created field value
func (o *Agreement) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Agreement) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Agreement) SetCreated(v string) {
	o.Created = v
}

// GetId returns the Id field value
func (o *Agreement) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Agreement) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Agreement) SetId(v string) {
	o.Id = v
}

// GetTags returns the Tags field value
func (o *Agreement) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *Agreement) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *Agreement) SetTags(v []string) {
	o.Tags = v
}

// GetUserId returns the UserId field value
func (o *Agreement) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *Agreement) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *Agreement) SetUserId(v string) {
	o.UserId = v
}

// GetVersion returns the Version field value
func (o *Agreement) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *Agreement) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *Agreement) SetVersion(v int32) {
	o.Version = v
}

func (o Agreement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Agreement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["agreementCode"] = o.AgreementCode
	if !IsNil(o.AgreementFulltext) {
		toSerialize["agreementFulltext"] = o.AgreementFulltext
	}
	toSerialize["contentId"] = o.ContentId
	toSerialize["created"] = o.Created
	toSerialize["id"] = o.Id
	toSerialize["tags"] = o.Tags
	toSerialize["userId"] = o.UserId
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

func (o *Agreement) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"agreementCode",
		"contentId",
		"created",
		"id",
		"tags",
		"userId",
		"version",
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

	varAgreement := _Agreement{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAgreement)

	if err != nil {
		return err
	}

	*o = Agreement(varAgreement)

	return err
}

type NullableAgreement struct {
	value *Agreement
	isSet bool
}

func (v NullableAgreement) Get() *Agreement {
	return v.value
}

func (v *NullableAgreement) Set(val *Agreement) {
	v.value = val
	v.isSet = true
}

func (v NullableAgreement) IsSet() bool {
	return v.isSet
}

func (v *NullableAgreement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgreement(val *Agreement) *NullableAgreement {
	return &NullableAgreement{value: val, isSet: true}
}

func (v NullableAgreement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgreement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
