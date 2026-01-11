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

// checks if the InventoryDrop type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryDrop{}

// InventoryDrop struct for InventoryDrop
type InventoryDrop struct {
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	AuthorId string `json:"authorId"`
	CreatedAt time.Time `json:"created_at"`
	DropExpiryDate NullableTime `json:"dropExpiryDate"`
	EndDropDate time.Time `json:"endDropDate"`
	Id string `json:"id"`
	IsDisabled bool `json:"isDisabled"`
	Name string `json:"name"`
	NotificationDetails InventoryNotificationDetails `json:"notificationDetails"`
	StartDropDate time.Time `json:"startDropDate"`
	Status string `json:"status"`
	Tags []string `json:"tags"`
	TargetGroup string `json:"targetGroup"`
	TemplateIds []string `json:"templateIds"`
	UpdatedAt time.Time `json:"updated_at"`
}

type _InventoryDrop InventoryDrop

// NewInventoryDrop instantiates a new InventoryDrop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryDrop(authorId string, createdAt time.Time, dropExpiryDate NullableTime, endDropDate time.Time, id string, isDisabled bool, name string, notificationDetails InventoryNotificationDetails, startDropDate time.Time, status string, tags []string, targetGroup string, templateIds []string, updatedAt time.Time) *InventoryDrop {
	this := InventoryDrop{}
	this.AuthorId = authorId
	this.CreatedAt = createdAt
	this.DropExpiryDate = dropExpiryDate
	this.EndDropDate = endDropDate
	this.Id = id
	this.IsDisabled = isDisabled
	this.Name = name
	this.NotificationDetails = notificationDetails
	this.StartDropDate = startDropDate
	this.Status = status
	this.Tags = tags
	this.TargetGroup = targetGroup
	this.TemplateIds = templateIds
	this.UpdatedAt = updatedAt
	return &this
}

// NewInventoryDropWithDefaults instantiates a new InventoryDrop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryDropWithDefaults() *InventoryDrop {
	this := InventoryDrop{}
	return &this
}

// GetAuthorId returns the AuthorId field value
func (o *InventoryDrop) GetAuthorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorId
}

// GetAuthorIdOk returns a tuple with the AuthorId field value
// and a boolean to check if the value has been set.
func (o *InventoryDrop) GetAuthorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorId, true
}

// SetAuthorId sets field value
func (o *InventoryDrop) SetAuthorId(v string) {
	o.AuthorId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *InventoryDrop) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *InventoryDrop) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *InventoryDrop) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDropExpiryDate returns the DropExpiryDate field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *InventoryDrop) GetDropExpiryDate() time.Time {
	if o == nil || o.DropExpiryDate.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.DropExpiryDate.Get()
}

// GetDropExpiryDateOk returns a tuple with the DropExpiryDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryDrop) GetDropExpiryDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DropExpiryDate.Get(), o.DropExpiryDate.IsSet()
}

// SetDropExpiryDate sets field value
func (o *InventoryDrop) SetDropExpiryDate(v time.Time) {
	o.DropExpiryDate.Set(&v)
}

// GetEndDropDate returns the EndDropDate field value
func (o *InventoryDrop) GetEndDropDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndDropDate
}

// GetEndDropDateOk returns a tuple with the EndDropDate field value
// and a boolean to check if the value has been set.
func (o *InventoryDrop) GetEndDropDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDropDate, true
}

// SetEndDropDate sets field value
func (o *InventoryDrop) SetEndDropDate(v time.Time) {
	o.EndDropDate = v
}

// GetId returns the Id field value
func (o *InventoryDrop) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InventoryDrop) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InventoryDrop) SetId(v string) {
	o.Id = v
}

// GetIsDisabled returns the IsDisabled field value
func (o *InventoryDrop) GetIsDisabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDisabled
}

// GetIsDisabledOk returns a tuple with the IsDisabled field value
// and a boolean to check if the value has been set.
func (o *InventoryDrop) GetIsDisabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDisabled, true
}

// SetIsDisabled sets field value
func (o *InventoryDrop) SetIsDisabled(v bool) {
	o.IsDisabled = v
}

// GetName returns the Name field value
func (o *InventoryDrop) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InventoryDrop) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InventoryDrop) SetName(v string) {
	o.Name = v
}

// GetNotificationDetails returns the NotificationDetails field value
func (o *InventoryDrop) GetNotificationDetails() InventoryNotificationDetails {
	if o == nil {
		var ret InventoryNotificationDetails
		return ret
	}

	return o.NotificationDetails
}

// GetNotificationDetailsOk returns a tuple with the NotificationDetails field value
// and a boolean to check if the value has been set.
func (o *InventoryDrop) GetNotificationDetailsOk() (*InventoryNotificationDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationDetails, true
}

// SetNotificationDetails sets field value
func (o *InventoryDrop) SetNotificationDetails(v InventoryNotificationDetails) {
	o.NotificationDetails = v
}

// GetStartDropDate returns the StartDropDate field value
func (o *InventoryDrop) GetStartDropDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartDropDate
}

// GetStartDropDateOk returns a tuple with the StartDropDate field value
// and a boolean to check if the value has been set.
func (o *InventoryDrop) GetStartDropDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDropDate, true
}

// SetStartDropDate sets field value
func (o *InventoryDrop) SetStartDropDate(v time.Time) {
	o.StartDropDate = v
}

// GetStatus returns the Status field value
func (o *InventoryDrop) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InventoryDrop) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InventoryDrop) SetStatus(v string) {
	o.Status = v
}

// GetTags returns the Tags field value
func (o *InventoryDrop) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *InventoryDrop) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *InventoryDrop) SetTags(v []string) {
	o.Tags = v
}

// GetTargetGroup returns the TargetGroup field value
func (o *InventoryDrop) GetTargetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetGroup
}

// GetTargetGroupOk returns a tuple with the TargetGroup field value
// and a boolean to check if the value has been set.
func (o *InventoryDrop) GetTargetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetGroup, true
}

// SetTargetGroup sets field value
func (o *InventoryDrop) SetTargetGroup(v string) {
	o.TargetGroup = v
}

// GetTemplateIds returns the TemplateIds field value
func (o *InventoryDrop) GetTemplateIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TemplateIds
}

// GetTemplateIdsOk returns a tuple with the TemplateIds field value
// and a boolean to check if the value has been set.
func (o *InventoryDrop) GetTemplateIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TemplateIds, true
}

// SetTemplateIds sets field value
func (o *InventoryDrop) SetTemplateIds(v []string) {
	o.TemplateIds = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *InventoryDrop) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *InventoryDrop) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *InventoryDrop) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o InventoryDrop) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryDrop) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authorId"] = o.AuthorId
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["dropExpiryDate"] = o.DropExpiryDate.Get()
	toSerialize["endDropDate"] = o.EndDropDate
	toSerialize["id"] = o.Id
	toSerialize["isDisabled"] = o.IsDisabled
	toSerialize["name"] = o.Name
	toSerialize["notificationDetails"] = o.NotificationDetails
	toSerialize["startDropDate"] = o.StartDropDate
	toSerialize["status"] = o.Status
	toSerialize["tags"] = o.Tags
	toSerialize["targetGroup"] = o.TargetGroup
	toSerialize["templateIds"] = o.TemplateIds
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *InventoryDrop) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authorId",
		"created_at",
		"dropExpiryDate",
		"endDropDate",
		"id",
		"isDisabled",
		"name",
		"notificationDetails",
		"startDropDate",
		"status",
		"tags",
		"targetGroup",
		"templateIds",
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

	varInventoryDrop := _InventoryDrop{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInventoryDrop)

	if err != nil {
		return err
	}

	*o = InventoryDrop(varInventoryDrop)

	return err
}

type NullableInventoryDrop struct {
	value *InventoryDrop
	isSet bool
}

func (v NullableInventoryDrop) Get() *InventoryDrop {
	return v.value
}

func (v *NullableInventoryDrop) Set(val *InventoryDrop) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryDrop) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryDrop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryDrop(val *InventoryDrop) *NullableInventoryDrop {
	return &NullableInventoryDrop{value: val, isSet: true}
}

func (v NullableInventoryDrop) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryDrop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


