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

// checks if the ModerationReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModerationReport{}

// ModerationReport struct for ModerationReport
type ModerationReport struct {
	// Valid values are the keys of the object `$.reportOptions[type]` from `GET /config`. Descriptions of these are found at `$.reportCategories[type]`.
	Category string `json:"category"`
	ContentId string `json:"contentId"`
	ContentName string `json:"contentName"`
	ContentThumbnailImageUrl NullableString `json:"contentThumbnailImageUrl"`
	// The subjective reason for the report
	Description string `json:"description"`
	EvidenceRequired bool `json:"evidenceRequired"`
	Id string `json:"id"`
	// Valid values are the strings in the array `$.reportOptions[type][category]` from `GET /config`. Descriptions of these are found at `$.reportReasons[type]`.
	Reason string `json:"reason"`
	SupportRequired bool `json:"supportRequired"`
	// Valid values are the keys of the object `$.reportOptions` from `GET /config`.
	Type string `json:"type"`
}

type _ModerationReport ModerationReport

// NewModerationReport instantiates a new ModerationReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModerationReport(category string, contentId string, contentName string, contentThumbnailImageUrl NullableString, description string, evidenceRequired bool, id string, reason string, supportRequired bool, type_ string) *ModerationReport {
	this := ModerationReport{}
	this.Category = category
	this.ContentId = contentId
	this.ContentName = contentName
	this.ContentThumbnailImageUrl = contentThumbnailImageUrl
	this.Description = description
	this.EvidenceRequired = evidenceRequired
	this.Id = id
	this.Reason = reason
	this.SupportRequired = supportRequired
	this.Type = type_
	return &this
}

// NewModerationReportWithDefaults instantiates a new ModerationReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModerationReportWithDefaults() *ModerationReport {
	this := ModerationReport{}
	return &this
}

// GetCategory returns the Category field value
func (o *ModerationReport) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *ModerationReport) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *ModerationReport) SetCategory(v string) {
	o.Category = v
}

// GetContentId returns the ContentId field value
func (o *ModerationReport) GetContentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentId
}

// GetContentIdOk returns a tuple with the ContentId field value
// and a boolean to check if the value has been set.
func (o *ModerationReport) GetContentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentId, true
}

// SetContentId sets field value
func (o *ModerationReport) SetContentId(v string) {
	o.ContentId = v
}

// GetContentName returns the ContentName field value
func (o *ModerationReport) GetContentName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentName
}

// GetContentNameOk returns a tuple with the ContentName field value
// and a boolean to check if the value has been set.
func (o *ModerationReport) GetContentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentName, true
}

// SetContentName sets field value
func (o *ModerationReport) SetContentName(v string) {
	o.ContentName = v
}

// GetContentThumbnailImageUrl returns the ContentThumbnailImageUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ModerationReport) GetContentThumbnailImageUrl() string {
	if o == nil || o.ContentThumbnailImageUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.ContentThumbnailImageUrl.Get()
}

// GetContentThumbnailImageUrlOk returns a tuple with the ContentThumbnailImageUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModerationReport) GetContentThumbnailImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentThumbnailImageUrl.Get(), o.ContentThumbnailImageUrl.IsSet()
}

// SetContentThumbnailImageUrl sets field value
func (o *ModerationReport) SetContentThumbnailImageUrl(v string) {
	o.ContentThumbnailImageUrl.Set(&v)
}

// GetDescription returns the Description field value
func (o *ModerationReport) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ModerationReport) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ModerationReport) SetDescription(v string) {
	o.Description = v
}

// GetEvidenceRequired returns the EvidenceRequired field value
func (o *ModerationReport) GetEvidenceRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EvidenceRequired
}

// GetEvidenceRequiredOk returns a tuple with the EvidenceRequired field value
// and a boolean to check if the value has been set.
func (o *ModerationReport) GetEvidenceRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvidenceRequired, true
}

// SetEvidenceRequired sets field value
func (o *ModerationReport) SetEvidenceRequired(v bool) {
	o.EvidenceRequired = v
}

// GetId returns the Id field value
func (o *ModerationReport) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ModerationReport) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ModerationReport) SetId(v string) {
	o.Id = v
}

// GetReason returns the Reason field value
func (o *ModerationReport) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *ModerationReport) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *ModerationReport) SetReason(v string) {
	o.Reason = v
}

// GetSupportRequired returns the SupportRequired field value
func (o *ModerationReport) GetSupportRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SupportRequired
}

// GetSupportRequiredOk returns a tuple with the SupportRequired field value
// and a boolean to check if the value has been set.
func (o *ModerationReport) GetSupportRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportRequired, true
}

// SetSupportRequired sets field value
func (o *ModerationReport) SetSupportRequired(v bool) {
	o.SupportRequired = v
}

// GetType returns the Type field value
func (o *ModerationReport) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ModerationReport) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ModerationReport) SetType(v string) {
	o.Type = v
}

func (o ModerationReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModerationReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["category"] = o.Category
	toSerialize["contentId"] = o.ContentId
	toSerialize["contentName"] = o.ContentName
	toSerialize["contentThumbnailImageUrl"] = o.ContentThumbnailImageUrl.Get()
	toSerialize["description"] = o.Description
	toSerialize["evidenceRequired"] = o.EvidenceRequired
	toSerialize["id"] = o.Id
	toSerialize["reason"] = o.Reason
	toSerialize["supportRequired"] = o.SupportRequired
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *ModerationReport) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"category",
		"contentId",
		"contentName",
		"contentThumbnailImageUrl",
		"description",
		"evidenceRequired",
		"id",
		"reason",
		"supportRequired",
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

	varModerationReport := _ModerationReport{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModerationReport)

	if err != nil {
		return err
	}

	*o = ModerationReport(varModerationReport)

	return err
}

type NullableModerationReport struct {
	value *ModerationReport
	isSet bool
}

func (v NullableModerationReport) Get() *ModerationReport {
	return v.value
}

func (v *NullableModerationReport) Set(val *ModerationReport) {
	v.value = val
	v.isSet = true
}

func (v NullableModerationReport) IsSet() bool {
	return v.isSet
}

func (v *NullableModerationReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModerationReport(val *ModerationReport) *NullableModerationReport {
	return &NullableModerationReport{value: val, isSet: true}
}

func (v NullableModerationReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModerationReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


