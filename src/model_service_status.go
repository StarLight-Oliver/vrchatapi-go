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

// checks if the ServiceStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceStatus{}

// ServiceStatus Status information for a service request
type ServiceStatus struct {
	CreatedAt time.Time `json:"created_at"`
	// The id of this service, NOT the id of the thing this service was requested for.
	Id string `json:"id"`
	Progress []map[string]interface{} `json:"progress"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	RequesterUserId string `json:"requesterUserId"`
	State string `json:"state"`
	// The id of the thing this service was requested for.
	SubjectId string `json:"subjectId"`
	// The kind of the thing this service was requested for.
	SubjectType string `json:"subjectType"`
	// The kind of service that was requested.
	Type string `json:"type"`
	UpdatedAt time.Time `json:"updated_at"`
}

type _ServiceStatus ServiceStatus

// NewServiceStatus instantiates a new ServiceStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceStatus(createdAt time.Time, id string, progress []map[string]interface{}, requesterUserId string, state string, subjectId string, subjectType string, type_ string, updatedAt time.Time) *ServiceStatus {
	this := ServiceStatus{}
	this.CreatedAt = createdAt
	this.Id = id
	this.Progress = progress
	this.RequesterUserId = requesterUserId
	this.State = state
	this.SubjectId = subjectId
	this.SubjectType = subjectType
	this.Type = type_
	this.UpdatedAt = updatedAt
	return &this
}

// NewServiceStatusWithDefaults instantiates a new ServiceStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceStatusWithDefaults() *ServiceStatus {
	this := ServiceStatus{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *ServiceStatus) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ServiceStatus) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ServiceStatus) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetId returns the Id field value
func (o *ServiceStatus) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServiceStatus) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServiceStatus) SetId(v string) {
	o.Id = v
}

// GetProgress returns the Progress field value
func (o *ServiceStatus) GetProgress() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Progress
}

// GetProgressOk returns a tuple with the Progress field value
// and a boolean to check if the value has been set.
func (o *ServiceStatus) GetProgressOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Progress, true
}

// SetProgress sets field value
func (o *ServiceStatus) SetProgress(v []map[string]interface{}) {
	o.Progress = v
}

// GetRequesterUserId returns the RequesterUserId field value
func (o *ServiceStatus) GetRequesterUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequesterUserId
}

// GetRequesterUserIdOk returns a tuple with the RequesterUserId field value
// and a boolean to check if the value has been set.
func (o *ServiceStatus) GetRequesterUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequesterUserId, true
}

// SetRequesterUserId sets field value
func (o *ServiceStatus) SetRequesterUserId(v string) {
	o.RequesterUserId = v
}

// GetState returns the State field value
func (o *ServiceStatus) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ServiceStatus) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ServiceStatus) SetState(v string) {
	o.State = v
}

// GetSubjectId returns the SubjectId field value
func (o *ServiceStatus) GetSubjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubjectId
}

// GetSubjectIdOk returns a tuple with the SubjectId field value
// and a boolean to check if the value has been set.
func (o *ServiceStatus) GetSubjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectId, true
}

// SetSubjectId sets field value
func (o *ServiceStatus) SetSubjectId(v string) {
	o.SubjectId = v
}

// GetSubjectType returns the SubjectType field value
func (o *ServiceStatus) GetSubjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubjectType
}

// GetSubjectTypeOk returns a tuple with the SubjectType field value
// and a boolean to check if the value has been set.
func (o *ServiceStatus) GetSubjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectType, true
}

// SetSubjectType sets field value
func (o *ServiceStatus) SetSubjectType(v string) {
	o.SubjectType = v
}

// GetType returns the Type field value
func (o *ServiceStatus) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServiceStatus) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ServiceStatus) SetType(v string) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ServiceStatus) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ServiceStatus) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ServiceStatus) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o ServiceStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["id"] = o.Id
	toSerialize["progress"] = o.Progress
	toSerialize["requesterUserId"] = o.RequesterUserId
	toSerialize["state"] = o.State
	toSerialize["subjectId"] = o.SubjectId
	toSerialize["subjectType"] = o.SubjectType
	toSerialize["type"] = o.Type
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *ServiceStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created_at",
		"id",
		"progress",
		"requesterUserId",
		"state",
		"subjectId",
		"subjectType",
		"type",
		"updated_at",
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

	varServiceStatus := _ServiceStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceStatus)

	if err != nil {
		return err
	}

	*o = ServiceStatus(varServiceStatus)

	return err
}

type NullableServiceStatus struct {
	value *ServiceStatus
	isSet bool
}

func (v NullableServiceStatus) Get() *ServiceStatus {
	return v.value
}

func (v *NullableServiceStatus) Set(val *ServiceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceStatus(val *ServiceStatus) *NullableServiceStatus {
	return &NullableServiceStatus{value: val, isSet: true}
}

func (v NullableServiceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


