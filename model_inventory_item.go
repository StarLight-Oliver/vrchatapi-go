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

// checks if the InventoryItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryItem{}

// InventoryItem struct for InventoryItem
type InventoryItem struct {
	Collections       []string                                   `json:"collections"`
	CreatedAt         time.Time                                  `json:"created_at"`
	DefaultAttributes map[string]InventoryDefaultAttributesValue `json:"defaultAttributes"`
	Description       string                                     `json:"description"`
	EquipSlot         *InventoryEquipSlot                        `json:"equipSlot,omitempty"`
	EquipSlots        []InventoryEquipSlot                       `json:"equipSlots,omitempty"`
	ExpiryDate        NullableTime                               `json:"expiryDate,omitempty"`
	Flags             []string                                   `json:"flags"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	HolderId               string                  `json:"holderId"`
	Id                     string                  `json:"id"`
	ImageUrl               string                  `json:"imageUrl"`
	IsArchived             bool                    `json:"isArchived"`
	IsSeen                 bool                    `json:"isSeen"`
	ItemType               InventoryItemType       `json:"itemType"`
	ItemTypeLabel          string                  `json:"itemTypeLabel"`
	Metadata               InventoryMetadata       `json:"metadata"`
	Name                   string                  `json:"name"`
	Quantifiable           bool                    `json:"quantifiable"`
	Tags                   []string                `json:"tags"`
	TemplateId             string                  `json:"templateId"`
	TemplateCreatedAt      time.Time               `json:"template_created_at"`
	TemplateUpdatedAt      time.Time               `json:"template_updated_at"`
	UpdatedAt              time.Time               `json:"updated_at"`
	UserAttributes         InventoryUserAttributes `json:"userAttributes"`
	ValidateUserAttributes bool                    `json:"validateUserAttributes"`
}

type _InventoryItem InventoryItem

// NewInventoryItem instantiates a new InventoryItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryItem(collections []string, createdAt time.Time, defaultAttributes map[string]InventoryDefaultAttributesValue, description string, flags []string, holderId string, id string, imageUrl string, isArchived bool, isSeen bool, itemType InventoryItemType, itemTypeLabel string, metadata InventoryMetadata, name string, quantifiable bool, tags []string, templateId string, templateCreatedAt time.Time, templateUpdatedAt time.Time, updatedAt time.Time, userAttributes InventoryUserAttributes, validateUserAttributes bool) *InventoryItem {
	this := InventoryItem{}
	this.Collections = collections
	this.CreatedAt = createdAt
	this.DefaultAttributes = defaultAttributes
	this.Description = description
	var equipSlot InventoryEquipSlot = InventoryEquipSlot_EMPTY
	this.EquipSlot = &equipSlot
	this.Flags = flags
	this.HolderId = holderId
	this.Id = id
	this.ImageUrl = imageUrl
	this.IsArchived = isArchived
	this.IsSeen = isSeen
	this.ItemType = itemType
	this.ItemTypeLabel = itemTypeLabel
	this.Metadata = metadata
	this.Name = name
	this.Quantifiable = quantifiable
	this.Tags = tags
	this.TemplateId = templateId
	this.TemplateCreatedAt = templateCreatedAt
	this.TemplateUpdatedAt = templateUpdatedAt
	this.UpdatedAt = updatedAt
	this.UserAttributes = userAttributes
	this.ValidateUserAttributes = validateUserAttributes
	return &this
}

// NewInventoryItemWithDefaults instantiates a new InventoryItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryItemWithDefaults() *InventoryItem {
	this := InventoryItem{}
	var equipSlot InventoryEquipSlot = InventoryEquipSlot_EMPTY
	this.EquipSlot = &equipSlot
	var itemType InventoryItemType = InventoryItemType_BUNDLE
	this.ItemType = itemType
	return &this
}

// GetCollections returns the Collections field value
func (o *InventoryItem) GetCollections() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Collections
}

// GetCollectionsOk returns a tuple with the Collections field value
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetCollectionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Collections, true
}

// SetCollections sets field value
func (o *InventoryItem) SetCollections(v []string) {
	o.Collections = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *InventoryItem) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *InventoryItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDefaultAttributes returns the DefaultAttributes field value
func (o *InventoryItem) GetDefaultAttributes() map[string]InventoryDefaultAttributesValue {
	if o == nil {
		var ret map[string]InventoryDefaultAttributesValue
		return ret
	}

	return o.DefaultAttributes
}

// GetDefaultAttributesOk returns a tuple with the DefaultAttributes field value
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetDefaultAttributesOk() (*map[string]InventoryDefaultAttributesValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultAttributes, true
}

// SetDefaultAttributes sets field value
func (o *InventoryItem) SetDefaultAttributes(v map[string]InventoryDefaultAttributesValue) {
	o.DefaultAttributes = v
}

// GetDescription returns the Description field value
func (o *InventoryItem) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *InventoryItem) SetDescription(v string) {
	o.Description = v
}

// GetEquipSlot returns the EquipSlot field value if set, zero value otherwise.
func (o *InventoryItem) GetEquipSlot() InventoryEquipSlot {
	if o == nil || IsNil(o.EquipSlot) {
		var ret InventoryEquipSlot
		return ret
	}
	return *o.EquipSlot
}

// GetEquipSlotOk returns a tuple with the EquipSlot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetEquipSlotOk() (*InventoryEquipSlot, bool) {
	if o == nil || IsNil(o.EquipSlot) {
		return nil, false
	}
	return o.EquipSlot, true
}

// HasEquipSlot returns a boolean if a field has been set.
func (o *InventoryItem) HasEquipSlot() bool {
	if o != nil && !IsNil(o.EquipSlot) {
		return true
	}

	return false
}

// SetEquipSlot gets a reference to the given InventoryEquipSlot and assigns it to the EquipSlot field.
func (o *InventoryItem) SetEquipSlot(v InventoryEquipSlot) {
	o.EquipSlot = &v
}

// GetEquipSlots returns the EquipSlots field value if set, zero value otherwise.
func (o *InventoryItem) GetEquipSlots() []InventoryEquipSlot {
	if o == nil || IsNil(o.EquipSlots) {
		var ret []InventoryEquipSlot
		return ret
	}
	return o.EquipSlots
}

// GetEquipSlotsOk returns a tuple with the EquipSlots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetEquipSlotsOk() ([]InventoryEquipSlot, bool) {
	if o == nil || IsNil(o.EquipSlots) {
		return nil, false
	}
	return o.EquipSlots, true
}

// HasEquipSlots returns a boolean if a field has been set.
func (o *InventoryItem) HasEquipSlots() bool {
	if o != nil && !IsNil(o.EquipSlots) {
		return true
	}

	return false
}

// SetEquipSlots gets a reference to the given []InventoryEquipSlot and assigns it to the EquipSlots field.
func (o *InventoryItem) SetEquipSlots(v []InventoryEquipSlot) {
	o.EquipSlots = v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InventoryItem) GetExpiryDate() time.Time {
	if o == nil || IsNil(o.ExpiryDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpiryDate.Get()
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InventoryItem) GetExpiryDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiryDate.Get(), o.ExpiryDate.IsSet()
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *InventoryItem) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate.IsSet() {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given NullableTime and assigns it to the ExpiryDate field.
func (o *InventoryItem) SetExpiryDate(v time.Time) {
	o.ExpiryDate.Set(&v)
}

// SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil
func (o *InventoryItem) SetExpiryDateNil() {
	o.ExpiryDate.Set(nil)
}

// UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
func (o *InventoryItem) UnsetExpiryDate() {
	o.ExpiryDate.Unset()
}

// GetFlags returns the Flags field value
func (o *InventoryItem) GetFlags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetFlagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Flags, true
}

// SetFlags sets field value
func (o *InventoryItem) SetFlags(v []string) {
	o.Flags = v
}

// GetHolderId returns the HolderId field value
func (o *InventoryItem) GetHolderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HolderId
}

// GetHolderIdOk returns a tuple with the HolderId field value
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetHolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HolderId, true
}

// SetHolderId sets field value
func (o *InventoryItem) SetHolderId(v string) {
	o.HolderId = v
}

// GetId returns the Id field value
func (o *InventoryItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InventoryItem) SetId(v string) {
	o.Id = v
}

// GetImageUrl returns the ImageUrl field value
func (o *InventoryItem) GetImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageUrl, true
}

// SetImageUrl sets field value
func (o *InventoryItem) SetImageUrl(v string) {
	o.ImageUrl = v
}

// GetIsArchived returns the IsArchived field value
func (o *InventoryItem) GetIsArchived() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsArchived
}

// GetIsArchivedOk returns a tuple with the IsArchived field value
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetIsArchivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsArchived, true
}

// SetIsArchived sets field value
func (o *InventoryItem) SetIsArchived(v bool) {
	o.IsArchived = v
}

// GetIsSeen returns the IsSeen field value
func (o *InventoryItem) GetIsSeen() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSeen
}

// GetIsSeenOk returns a tuple with the IsSeen field value
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetIsSeenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSeen, true
}

// SetIsSeen sets field value
func (o *InventoryItem) SetIsSeen(v bool) {
	o.IsSeen = v
}

// GetItemType returns the ItemType field value
func (o *InventoryItem) GetItemType() InventoryItemType {
	if o == nil {
		var ret InventoryItemType
		return ret
	}

	return o.ItemType
}

// GetItemTypeOk returns a tuple with the ItemType field value
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetItemTypeOk() (*InventoryItemType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemType, true
}

// SetItemType sets field value
func (o *InventoryItem) SetItemType(v InventoryItemType) {
	o.ItemType = v
}

// GetItemTypeLabel returns the ItemTypeLabel field value
func (o *InventoryItem) GetItemTypeLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemTypeLabel
}

// GetItemTypeLabelOk returns a tuple with the ItemTypeLabel field value
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetItemTypeLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemTypeLabel, true
}

// SetItemTypeLabel sets field value
func (o *InventoryItem) SetItemTypeLabel(v string) {
	o.ItemTypeLabel = v
}

// GetMetadata returns the Metadata field value
func (o *InventoryItem) GetMetadata() InventoryMetadata {
	if o == nil {
		var ret InventoryMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetMetadataOk() (*InventoryMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *InventoryItem) SetMetadata(v InventoryMetadata) {
	o.Metadata = v
}

// GetName returns the Name field value
func (o *InventoryItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InventoryItem) SetName(v string) {
	o.Name = v
}

// GetQuantifiable returns the Quantifiable field value
func (o *InventoryItem) GetQuantifiable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Quantifiable
}

// GetQuantifiableOk returns a tuple with the Quantifiable field value
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetQuantifiableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantifiable, true
}

// SetQuantifiable sets field value
func (o *InventoryItem) SetQuantifiable(v bool) {
	o.Quantifiable = v
}

// GetTags returns the Tags field value
func (o *InventoryItem) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *InventoryItem) SetTags(v []string) {
	o.Tags = v
}

// GetTemplateId returns the TemplateId field value
func (o *InventoryItem) GetTemplateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetTemplateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateId, true
}

// SetTemplateId sets field value
func (o *InventoryItem) SetTemplateId(v string) {
	o.TemplateId = v
}

// GetTemplateCreatedAt returns the TemplateCreatedAt field value
func (o *InventoryItem) GetTemplateCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TemplateCreatedAt
}

// GetTemplateCreatedAtOk returns a tuple with the TemplateCreatedAt field value
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetTemplateCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateCreatedAt, true
}

// SetTemplateCreatedAt sets field value
func (o *InventoryItem) SetTemplateCreatedAt(v time.Time) {
	o.TemplateCreatedAt = v
}

// GetTemplateUpdatedAt returns the TemplateUpdatedAt field value
func (o *InventoryItem) GetTemplateUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TemplateUpdatedAt
}

// GetTemplateUpdatedAtOk returns a tuple with the TemplateUpdatedAt field value
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetTemplateUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateUpdatedAt, true
}

// SetTemplateUpdatedAt sets field value
func (o *InventoryItem) SetTemplateUpdatedAt(v time.Time) {
	o.TemplateUpdatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *InventoryItem) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *InventoryItem) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetUserAttributes returns the UserAttributes field value
func (o *InventoryItem) GetUserAttributes() InventoryUserAttributes {
	if o == nil {
		var ret InventoryUserAttributes
		return ret
	}

	return o.UserAttributes
}

// GetUserAttributesOk returns a tuple with the UserAttributes field value
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetUserAttributesOk() (*InventoryUserAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserAttributes, true
}

// SetUserAttributes sets field value
func (o *InventoryItem) SetUserAttributes(v InventoryUserAttributes) {
	o.UserAttributes = v
}

// GetValidateUserAttributes returns the ValidateUserAttributes field value
func (o *InventoryItem) GetValidateUserAttributes() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ValidateUserAttributes
}

// GetValidateUserAttributesOk returns a tuple with the ValidateUserAttributes field value
// and a boolean to check if the value has been set.
func (o *InventoryItem) GetValidateUserAttributesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidateUserAttributes, true
}

// SetValidateUserAttributes sets field value
func (o *InventoryItem) SetValidateUserAttributes(v bool) {
	o.ValidateUserAttributes = v
}

func (o InventoryItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collections"] = o.Collections
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["defaultAttributes"] = o.DefaultAttributes
	toSerialize["description"] = o.Description
	if !IsNil(o.EquipSlot) {
		toSerialize["equipSlot"] = o.EquipSlot
	}
	if !IsNil(o.EquipSlots) {
		toSerialize["equipSlots"] = o.EquipSlots
	}
	if o.ExpiryDate.IsSet() {
		toSerialize["expiryDate"] = o.ExpiryDate.Get()
	}
	toSerialize["flags"] = o.Flags
	toSerialize["holderId"] = o.HolderId
	toSerialize["id"] = o.Id
	toSerialize["imageUrl"] = o.ImageUrl
	toSerialize["isArchived"] = o.IsArchived
	toSerialize["isSeen"] = o.IsSeen
	toSerialize["itemType"] = o.ItemType
	toSerialize["itemTypeLabel"] = o.ItemTypeLabel
	toSerialize["metadata"] = o.Metadata
	toSerialize["name"] = o.Name
	toSerialize["quantifiable"] = o.Quantifiable
	toSerialize["tags"] = o.Tags
	toSerialize["templateId"] = o.TemplateId
	toSerialize["template_created_at"] = o.TemplateCreatedAt
	toSerialize["template_updated_at"] = o.TemplateUpdatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["userAttributes"] = o.UserAttributes
	toSerialize["validateUserAttributes"] = o.ValidateUserAttributes
	return toSerialize, nil
}

func (o *InventoryItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"collections",
		"created_at",
		"defaultAttributes",
		"description",
		"flags",
		"holderId",
		"id",
		"imageUrl",
		"isArchived",
		"isSeen",
		"itemType",
		"itemTypeLabel",
		"metadata",
		"name",
		"quantifiable",
		"tags",
		"templateId",
		"template_created_at",
		"template_updated_at",
		"updated_at",
		"userAttributes",
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

	varInventoryItem := _InventoryItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInventoryItem)

	if err != nil {
		return err
	}

	*o = InventoryItem(varInventoryItem)

	return err
}

type NullableInventoryItem struct {
	value *InventoryItem
	isSet bool
}

func (v NullableInventoryItem) Get() *InventoryItem {
	return v.value
}

func (v *NullableInventoryItem) Set(val *InventoryItem) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryItem) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryItem(val *InventoryItem) *NullableInventoryItem {
	return &NullableInventoryItem{value: val, isSet: true}
}

func (v NullableInventoryItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
