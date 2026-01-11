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
	"time"
)

// checks if the InventoryTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryTemplate{}

// InventoryTemplate struct for InventoryTemplate
type InventoryTemplate struct {
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	AuthorId               string                        `json:"authorId"`
	Collections            []string                      `json:"collections"`
	CreatedAt              time.Time                     `json:"created_at"`
	DefaultAttributes      map[string]interface{}        `json:"defaultAttributes"`
	Description            string                        `json:"description"`
	EquipSlots             []string                      `json:"equipSlots"`
	Flags                  []string                      `json:"flags"`
	Id                     string                        `json:"id"`
	ImageUrl               string                        `json:"imageUrl"`
	ItemType               InventoryItemType             `json:"itemType"`
	ItemTypeLabel          string                        `json:"itemTypeLabel"`
	Metadata               *InventoryMetadata            `json:"metadata,omitempty"`
	Name                   string                        `json:"name"`
	NotificationDetails    *InventoryNotificationDetails `json:"notificationDetails,omitempty"`
	Status                 string                        `json:"status"`
	Tags                   []string                      `json:"tags"`
	UpdatedAt              time.Time                     `json:"updated_at"`
	ValidateUserAttributes bool                          `json:"validateUserAttributes"`
}

type _InventoryTemplate InventoryTemplate

// NewInventoryTemplate instantiates a new InventoryTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryTemplate(authorId string, collections []string, createdAt time.Time, defaultAttributes map[string]interface{}, description string, equipSlots []string, flags []string, id string, imageUrl string, itemType InventoryItemType, itemTypeLabel string, name string, status string, tags []string, updatedAt time.Time, validateUserAttributes bool) *InventoryTemplate {
	this := InventoryTemplate{}
	this.AuthorId = authorId
	this.Collections = collections
	this.CreatedAt = createdAt
	this.DefaultAttributes = defaultAttributes
	this.Description = description
	this.EquipSlots = equipSlots
	this.Flags = flags
	this.Id = id
	this.ImageUrl = imageUrl
	this.ItemType = itemType
	this.ItemTypeLabel = itemTypeLabel
	this.Name = name
	this.Status = status
	this.Tags = tags
	this.UpdatedAt = updatedAt
	this.ValidateUserAttributes = validateUserAttributes
	return &this
}

// NewInventoryTemplateWithDefaults instantiates a new InventoryTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryTemplateWithDefaults() *InventoryTemplate {
	this := InventoryTemplate{}
	var itemType InventoryItemType = InventoryItemType_BUNDLE
	this.ItemType = itemType
	return &this
}

// GetAuthorId returns the AuthorId field value
func (o *InventoryTemplate) GetAuthorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorId
}

// GetAuthorIdOk returns a tuple with the AuthorId field value
// and a boolean to check if the value has been set.
func (o *InventoryTemplate) GetAuthorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorId, true
}

// SetAuthorId sets field value
func (o *InventoryTemplate) SetAuthorId(v string) {
	o.AuthorId = v
}

// GetCollections returns the Collections field value
func (o *InventoryTemplate) GetCollections() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Collections
}

// GetCollectionsOk returns a tuple with the Collections field value
// and a boolean to check if the value has been set.
func (o *InventoryTemplate) GetCollectionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Collections, true
}

// SetCollections sets field value
func (o *InventoryTemplate) SetCollections(v []string) {
	o.Collections = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *InventoryTemplate) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *InventoryTemplate) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *InventoryTemplate) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDefaultAttributes returns the DefaultAttributes field value
func (o *InventoryTemplate) GetDefaultAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.DefaultAttributes
}

// GetDefaultAttributesOk returns a tuple with the DefaultAttributes field value
// and a boolean to check if the value has been set.
func (o *InventoryTemplate) GetDefaultAttributesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.DefaultAttributes, true
}

// SetDefaultAttributes sets field value
func (o *InventoryTemplate) SetDefaultAttributes(v map[string]interface{}) {
	o.DefaultAttributes = v
}

// GetDescription returns the Description field value
func (o *InventoryTemplate) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *InventoryTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *InventoryTemplate) SetDescription(v string) {
	o.Description = v
}

// GetEquipSlots returns the EquipSlots field value
func (o *InventoryTemplate) GetEquipSlots() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EquipSlots
}

// GetEquipSlotsOk returns a tuple with the EquipSlots field value
// and a boolean to check if the value has been set.
func (o *InventoryTemplate) GetEquipSlotsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EquipSlots, true
}

// SetEquipSlots sets field value
func (o *InventoryTemplate) SetEquipSlots(v []string) {
	o.EquipSlots = v
}

// GetFlags returns the Flags field value
func (o *InventoryTemplate) GetFlags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value
// and a boolean to check if the value has been set.
func (o *InventoryTemplate) GetFlagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Flags, true
}

// SetFlags sets field value
func (o *InventoryTemplate) SetFlags(v []string) {
	o.Flags = v
}

// GetId returns the Id field value
func (o *InventoryTemplate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InventoryTemplate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InventoryTemplate) SetId(v string) {
	o.Id = v
}

// GetImageUrl returns the ImageUrl field value
func (o *InventoryTemplate) GetImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
func (o *InventoryTemplate) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageUrl, true
}

// SetImageUrl sets field value
func (o *InventoryTemplate) SetImageUrl(v string) {
	o.ImageUrl = v
}

// GetItemType returns the ItemType field value
func (o *InventoryTemplate) GetItemType() InventoryItemType {
	if o == nil {
		var ret InventoryItemType
		return ret
	}

	return o.ItemType
}

// GetItemTypeOk returns a tuple with the ItemType field value
// and a boolean to check if the value has been set.
func (o *InventoryTemplate) GetItemTypeOk() (*InventoryItemType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemType, true
}

// SetItemType sets field value
func (o *InventoryTemplate) SetItemType(v InventoryItemType) {
	o.ItemType = v
}

// GetItemTypeLabel returns the ItemTypeLabel field value
func (o *InventoryTemplate) GetItemTypeLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemTypeLabel
}

// GetItemTypeLabelOk returns a tuple with the ItemTypeLabel field value
// and a boolean to check if the value has been set.
func (o *InventoryTemplate) GetItemTypeLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemTypeLabel, true
}

// SetItemTypeLabel sets field value
func (o *InventoryTemplate) SetItemTypeLabel(v string) {
	o.ItemTypeLabel = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *InventoryTemplate) GetMetadata() InventoryMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret InventoryMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryTemplate) GetMetadataOk() (*InventoryMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *InventoryTemplate) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given InventoryMetadata and assigns it to the Metadata field.
func (o *InventoryTemplate) SetMetadata(v InventoryMetadata) {
	o.Metadata = &v
}

// GetName returns the Name field value
func (o *InventoryTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InventoryTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InventoryTemplate) SetName(v string) {
	o.Name = v
}

// GetNotificationDetails returns the NotificationDetails field value if set, zero value otherwise.
func (o *InventoryTemplate) GetNotificationDetails() InventoryNotificationDetails {
	if o == nil || IsNil(o.NotificationDetails) {
		var ret InventoryNotificationDetails
		return ret
	}
	return *o.NotificationDetails
}

// GetNotificationDetailsOk returns a tuple with the NotificationDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryTemplate) GetNotificationDetailsOk() (*InventoryNotificationDetails, bool) {
	if o == nil || IsNil(o.NotificationDetails) {
		return nil, false
	}
	return o.NotificationDetails, true
}

// HasNotificationDetails returns a boolean if a field has been set.
func (o *InventoryTemplate) HasNotificationDetails() bool {
	if o != nil && !IsNil(o.NotificationDetails) {
		return true
	}

	return false
}

// SetNotificationDetails gets a reference to the given InventoryNotificationDetails and assigns it to the NotificationDetails field.
func (o *InventoryTemplate) SetNotificationDetails(v InventoryNotificationDetails) {
	o.NotificationDetails = &v
}

// GetStatus returns the Status field value
func (o *InventoryTemplate) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InventoryTemplate) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InventoryTemplate) SetStatus(v string) {
	o.Status = v
}

// GetTags returns the Tags field value
func (o *InventoryTemplate) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *InventoryTemplate) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *InventoryTemplate) SetTags(v []string) {
	o.Tags = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *InventoryTemplate) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *InventoryTemplate) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *InventoryTemplate) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetValidateUserAttributes returns the ValidateUserAttributes field value
func (o *InventoryTemplate) GetValidateUserAttributes() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ValidateUserAttributes
}

// GetValidateUserAttributesOk returns a tuple with the ValidateUserAttributes field value
// and a boolean to check if the value has been set.
func (o *InventoryTemplate) GetValidateUserAttributesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidateUserAttributes, true
}

// SetValidateUserAttributes sets field value
func (o *InventoryTemplate) SetValidateUserAttributes(v bool) {
	o.ValidateUserAttributes = v
}

func (o InventoryTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authorId"] = o.AuthorId
	toSerialize["collections"] = o.Collections
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["defaultAttributes"] = o.DefaultAttributes
	toSerialize["description"] = o.Description
	toSerialize["equipSlots"] = o.EquipSlots
	toSerialize["flags"] = o.Flags
	toSerialize["id"] = o.Id
	toSerialize["imageUrl"] = o.ImageUrl
	toSerialize["itemType"] = o.ItemType
	toSerialize["itemTypeLabel"] = o.ItemTypeLabel
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.NotificationDetails) {
		toSerialize["notificationDetails"] = o.NotificationDetails
	}
	toSerialize["status"] = o.Status
	toSerialize["tags"] = o.Tags
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["validateUserAttributes"] = o.ValidateUserAttributes
	return toSerialize, nil
}

func (o *InventoryTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authorId",
		"collections",
		"created_at",
		"defaultAttributes",
		"description",
		"equipSlots",
		"flags",
		"id",
		"imageUrl",
		"itemType",
		"itemTypeLabel",
		"name",
		"status",
		"tags",
		"updated_at",
		"validateUserAttributes",
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

	varInventoryTemplate := _InventoryTemplate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInventoryTemplate)

	if err != nil {
		return err
	}

	*o = InventoryTemplate(varInventoryTemplate)

	return err
}

type NullableInventoryTemplate struct {
	value *InventoryTemplate
	isSet bool
}

func (v NullableInventoryTemplate) Get() *InventoryTemplate {
	return v.value
}

func (v *NullableInventoryTemplate) Set(val *InventoryTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryTemplate(val *InventoryTemplate) *NullableInventoryTemplate {
	return &NullableInventoryTemplate{value: val, isSet: true}
}

func (v NullableInventoryTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
