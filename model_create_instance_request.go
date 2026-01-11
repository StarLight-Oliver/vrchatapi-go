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

// checks if the CreateInstanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateInstanceRequest{}

// CreateInstanceRequest struct for CreateInstanceRequest
type CreateInstanceRequest struct {
	AgeGate         *bool   `json:"ageGate,omitempty"`
	CalendarEntryId *string `json:"calendarEntryId,omitempty"`
	// Only applies to invite type instances to make them invite+
	CanRequestInvite *bool `json:"canRequestInvite,omitempty"`
	// The time after which users won't be allowed to join the instance. This doesn't work for public instances.
	ClosedAt        *time.Time               `json:"closedAt,omitempty"`
	ContentSettings *InstanceContentSettings `json:"contentSettings,omitempty"`
	DisplayName     NullableString           `json:"displayName,omitempty"`
	GroupAccessType *GroupAccessType         `json:"groupAccessType,omitempty"`
	// Currently unused, but will eventually be a flag to set if the closing of the instance should kick people.
	HardClose                  *bool        `json:"hardClose,omitempty"`
	InstancePersistenceEnabled NullableBool `json:"instancePersistenceEnabled,omitempty"`
	InviteOnly                 *bool        `json:"inviteOnly,omitempty"`
	// A groupId if the instance type is \"group\", null if instance type is public, or a userId otherwise
	OwnerId                  NullableString `json:"ownerId,omitempty"`
	PlayerPersistenceEnabled NullableBool   `json:"playerPersistenceEnabled,omitempty"`
	QueueEnabled             *bool          `json:"queueEnabled,omitempty"`
	Region                   InstanceRegion `json:"region"`
	// Group roleIds that are allowed to join if the type is \"group\" and groupAccessType is \"member\"
	RoleIds []string     `json:"roleIds,omitempty"`
	Type    InstanceType `json:"type"`
	// WorldID be \"offline\" on User profiles if you are not friends with that user.
	WorldId string `json:"worldId"`
}

type _CreateInstanceRequest CreateInstanceRequest

// NewCreateInstanceRequest instantiates a new CreateInstanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInstanceRequest(region InstanceRegion, type_ InstanceType, worldId string) *CreateInstanceRequest {
	this := CreateInstanceRequest{}
	var ageGate bool = false
	this.AgeGate = &ageGate
	var canRequestInvite bool = false
	this.CanRequestInvite = &canRequestInvite
	var groupAccessType GroupAccessType = GroupAccessType_MEMBERS
	this.GroupAccessType = &groupAccessType
	var hardClose bool = false
	this.HardClose = &hardClose
	var inviteOnly bool = false
	this.InviteOnly = &inviteOnly
	var queueEnabled bool = false
	this.QueueEnabled = &queueEnabled
	this.Region = region
	this.Type = type_
	this.WorldId = worldId
	return &this
}

// NewCreateInstanceRequestWithDefaults instantiates a new CreateInstanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInstanceRequestWithDefaults() *CreateInstanceRequest {
	this := CreateInstanceRequest{}
	var ageGate bool = false
	this.AgeGate = &ageGate
	var canRequestInvite bool = false
	this.CanRequestInvite = &canRequestInvite
	var groupAccessType GroupAccessType = GroupAccessType_MEMBERS
	this.GroupAccessType = &groupAccessType
	var hardClose bool = false
	this.HardClose = &hardClose
	var inviteOnly bool = false
	this.InviteOnly = &inviteOnly
	var queueEnabled bool = false
	this.QueueEnabled = &queueEnabled
	var region InstanceRegion = InstanceRegion_US
	this.Region = region
	return &this
}

// GetAgeGate returns the AgeGate field value if set, zero value otherwise.
func (o *CreateInstanceRequest) GetAgeGate() bool {
	if o == nil || IsNil(o.AgeGate) {
		var ret bool
		return ret
	}
	return *o.AgeGate
}

// GetAgeGateOk returns a tuple with the AgeGate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceRequest) GetAgeGateOk() (*bool, bool) {
	if o == nil || IsNil(o.AgeGate) {
		return nil, false
	}
	return o.AgeGate, true
}

// HasAgeGate returns a boolean if a field has been set.
func (o *CreateInstanceRequest) HasAgeGate() bool {
	if o != nil && !IsNil(o.AgeGate) {
		return true
	}

	return false
}

// SetAgeGate gets a reference to the given bool and assigns it to the AgeGate field.
func (o *CreateInstanceRequest) SetAgeGate(v bool) {
	o.AgeGate = &v
}

// GetCalendarEntryId returns the CalendarEntryId field value if set, zero value otherwise.
func (o *CreateInstanceRequest) GetCalendarEntryId() string {
	if o == nil || IsNil(o.CalendarEntryId) {
		var ret string
		return ret
	}
	return *o.CalendarEntryId
}

// GetCalendarEntryIdOk returns a tuple with the CalendarEntryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceRequest) GetCalendarEntryIdOk() (*string, bool) {
	if o == nil || IsNil(o.CalendarEntryId) {
		return nil, false
	}
	return o.CalendarEntryId, true
}

// HasCalendarEntryId returns a boolean if a field has been set.
func (o *CreateInstanceRequest) HasCalendarEntryId() bool {
	if o != nil && !IsNil(o.CalendarEntryId) {
		return true
	}

	return false
}

// SetCalendarEntryId gets a reference to the given string and assigns it to the CalendarEntryId field.
func (o *CreateInstanceRequest) SetCalendarEntryId(v string) {
	o.CalendarEntryId = &v
}

// GetCanRequestInvite returns the CanRequestInvite field value if set, zero value otherwise.
func (o *CreateInstanceRequest) GetCanRequestInvite() bool {
	if o == nil || IsNil(o.CanRequestInvite) {
		var ret bool
		return ret
	}
	return *o.CanRequestInvite
}

// GetCanRequestInviteOk returns a tuple with the CanRequestInvite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceRequest) GetCanRequestInviteOk() (*bool, bool) {
	if o == nil || IsNil(o.CanRequestInvite) {
		return nil, false
	}
	return o.CanRequestInvite, true
}

// HasCanRequestInvite returns a boolean if a field has been set.
func (o *CreateInstanceRequest) HasCanRequestInvite() bool {
	if o != nil && !IsNil(o.CanRequestInvite) {
		return true
	}

	return false
}

// SetCanRequestInvite gets a reference to the given bool and assigns it to the CanRequestInvite field.
func (o *CreateInstanceRequest) SetCanRequestInvite(v bool) {
	o.CanRequestInvite = &v
}

// GetClosedAt returns the ClosedAt field value if set, zero value otherwise.
func (o *CreateInstanceRequest) GetClosedAt() time.Time {
	if o == nil || IsNil(o.ClosedAt) {
		var ret time.Time
		return ret
	}
	return *o.ClosedAt
}

// GetClosedAtOk returns a tuple with the ClosedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceRequest) GetClosedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ClosedAt) {
		return nil, false
	}
	return o.ClosedAt, true
}

// HasClosedAt returns a boolean if a field has been set.
func (o *CreateInstanceRequest) HasClosedAt() bool {
	if o != nil && !IsNil(o.ClosedAt) {
		return true
	}

	return false
}

// SetClosedAt gets a reference to the given time.Time and assigns it to the ClosedAt field.
func (o *CreateInstanceRequest) SetClosedAt(v time.Time) {
	o.ClosedAt = &v
}

// GetContentSettings returns the ContentSettings field value if set, zero value otherwise.
func (o *CreateInstanceRequest) GetContentSettings() InstanceContentSettings {
	if o == nil || IsNil(o.ContentSettings) {
		var ret InstanceContentSettings
		return ret
	}
	return *o.ContentSettings
}

// GetContentSettingsOk returns a tuple with the ContentSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceRequest) GetContentSettingsOk() (*InstanceContentSettings, bool) {
	if o == nil || IsNil(o.ContentSettings) {
		return nil, false
	}
	return o.ContentSettings, true
}

// HasContentSettings returns a boolean if a field has been set.
func (o *CreateInstanceRequest) HasContentSettings() bool {
	if o != nil && !IsNil(o.ContentSettings) {
		return true
	}

	return false
}

// SetContentSettings gets a reference to the given InstanceContentSettings and assigns it to the ContentSettings field.
func (o *CreateInstanceRequest) SetContentSettings(v InstanceContentSettings) {
	o.ContentSettings = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateInstanceRequest) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateInstanceRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CreateInstanceRequest) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *CreateInstanceRequest) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}

// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *CreateInstanceRequest) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *CreateInstanceRequest) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetGroupAccessType returns the GroupAccessType field value if set, zero value otherwise.
func (o *CreateInstanceRequest) GetGroupAccessType() GroupAccessType {
	if o == nil || IsNil(o.GroupAccessType) {
		var ret GroupAccessType
		return ret
	}
	return *o.GroupAccessType
}

// GetGroupAccessTypeOk returns a tuple with the GroupAccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceRequest) GetGroupAccessTypeOk() (*GroupAccessType, bool) {
	if o == nil || IsNil(o.GroupAccessType) {
		return nil, false
	}
	return o.GroupAccessType, true
}

// HasGroupAccessType returns a boolean if a field has been set.
func (o *CreateInstanceRequest) HasGroupAccessType() bool {
	if o != nil && !IsNil(o.GroupAccessType) {
		return true
	}

	return false
}

// SetGroupAccessType gets a reference to the given GroupAccessType and assigns it to the GroupAccessType field.
func (o *CreateInstanceRequest) SetGroupAccessType(v GroupAccessType) {
	o.GroupAccessType = &v
}

// GetHardClose returns the HardClose field value if set, zero value otherwise.
func (o *CreateInstanceRequest) GetHardClose() bool {
	if o == nil || IsNil(o.HardClose) {
		var ret bool
		return ret
	}
	return *o.HardClose
}

// GetHardCloseOk returns a tuple with the HardClose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceRequest) GetHardCloseOk() (*bool, bool) {
	if o == nil || IsNil(o.HardClose) {
		return nil, false
	}
	return o.HardClose, true
}

// HasHardClose returns a boolean if a field has been set.
func (o *CreateInstanceRequest) HasHardClose() bool {
	if o != nil && !IsNil(o.HardClose) {
		return true
	}

	return false
}

// SetHardClose gets a reference to the given bool and assigns it to the HardClose field.
func (o *CreateInstanceRequest) SetHardClose(v bool) {
	o.HardClose = &v
}

// GetInstancePersistenceEnabled returns the InstancePersistenceEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateInstanceRequest) GetInstancePersistenceEnabled() bool {
	if o == nil || IsNil(o.InstancePersistenceEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.InstancePersistenceEnabled.Get()
}

// GetInstancePersistenceEnabledOk returns a tuple with the InstancePersistenceEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateInstanceRequest) GetInstancePersistenceEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstancePersistenceEnabled.Get(), o.InstancePersistenceEnabled.IsSet()
}

// HasInstancePersistenceEnabled returns a boolean if a field has been set.
func (o *CreateInstanceRequest) HasInstancePersistenceEnabled() bool {
	if o != nil && o.InstancePersistenceEnabled.IsSet() {
		return true
	}

	return false
}

// SetInstancePersistenceEnabled gets a reference to the given NullableBool and assigns it to the InstancePersistenceEnabled field.
func (o *CreateInstanceRequest) SetInstancePersistenceEnabled(v bool) {
	o.InstancePersistenceEnabled.Set(&v)
}

// SetInstancePersistenceEnabledNil sets the value for InstancePersistenceEnabled to be an explicit nil
func (o *CreateInstanceRequest) SetInstancePersistenceEnabledNil() {
	o.InstancePersistenceEnabled.Set(nil)
}

// UnsetInstancePersistenceEnabled ensures that no value is present for InstancePersistenceEnabled, not even an explicit nil
func (o *CreateInstanceRequest) UnsetInstancePersistenceEnabled() {
	o.InstancePersistenceEnabled.Unset()
}

// GetInviteOnly returns the InviteOnly field value if set, zero value otherwise.
func (o *CreateInstanceRequest) GetInviteOnly() bool {
	if o == nil || IsNil(o.InviteOnly) {
		var ret bool
		return ret
	}
	return *o.InviteOnly
}

// GetInviteOnlyOk returns a tuple with the InviteOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceRequest) GetInviteOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.InviteOnly) {
		return nil, false
	}
	return o.InviteOnly, true
}

// HasInviteOnly returns a boolean if a field has been set.
func (o *CreateInstanceRequest) HasInviteOnly() bool {
	if o != nil && !IsNil(o.InviteOnly) {
		return true
	}

	return false
}

// SetInviteOnly gets a reference to the given bool and assigns it to the InviteOnly field.
func (o *CreateInstanceRequest) SetInviteOnly(v bool) {
	o.InviteOnly = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateInstanceRequest) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId.Get()) {
		var ret string
		return ret
	}
	return *o.OwnerId.Get()
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateInstanceRequest) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerId.Get(), o.OwnerId.IsSet()
}

// HasOwnerId returns a boolean if a field has been set.
func (o *CreateInstanceRequest) HasOwnerId() bool {
	if o != nil && o.OwnerId.IsSet() {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given NullableString and assigns it to the OwnerId field.
func (o *CreateInstanceRequest) SetOwnerId(v string) {
	o.OwnerId.Set(&v)
}

// SetOwnerIdNil sets the value for OwnerId to be an explicit nil
func (o *CreateInstanceRequest) SetOwnerIdNil() {
	o.OwnerId.Set(nil)
}

// UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
func (o *CreateInstanceRequest) UnsetOwnerId() {
	o.OwnerId.Unset()
}

// GetPlayerPersistenceEnabled returns the PlayerPersistenceEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateInstanceRequest) GetPlayerPersistenceEnabled() bool {
	if o == nil || IsNil(o.PlayerPersistenceEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.PlayerPersistenceEnabled.Get()
}

// GetPlayerPersistenceEnabledOk returns a tuple with the PlayerPersistenceEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateInstanceRequest) GetPlayerPersistenceEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlayerPersistenceEnabled.Get(), o.PlayerPersistenceEnabled.IsSet()
}

// HasPlayerPersistenceEnabled returns a boolean if a field has been set.
func (o *CreateInstanceRequest) HasPlayerPersistenceEnabled() bool {
	if o != nil && o.PlayerPersistenceEnabled.IsSet() {
		return true
	}

	return false
}

// SetPlayerPersistenceEnabled gets a reference to the given NullableBool and assigns it to the PlayerPersistenceEnabled field.
func (o *CreateInstanceRequest) SetPlayerPersistenceEnabled(v bool) {
	o.PlayerPersistenceEnabled.Set(&v)
}

// SetPlayerPersistenceEnabledNil sets the value for PlayerPersistenceEnabled to be an explicit nil
func (o *CreateInstanceRequest) SetPlayerPersistenceEnabledNil() {
	o.PlayerPersistenceEnabled.Set(nil)
}

// UnsetPlayerPersistenceEnabled ensures that no value is present for PlayerPersistenceEnabled, not even an explicit nil
func (o *CreateInstanceRequest) UnsetPlayerPersistenceEnabled() {
	o.PlayerPersistenceEnabled.Unset()
}

// GetQueueEnabled returns the QueueEnabled field value if set, zero value otherwise.
func (o *CreateInstanceRequest) GetQueueEnabled() bool {
	if o == nil || IsNil(o.QueueEnabled) {
		var ret bool
		return ret
	}
	return *o.QueueEnabled
}

// GetQueueEnabledOk returns a tuple with the QueueEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceRequest) GetQueueEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.QueueEnabled) {
		return nil, false
	}
	return o.QueueEnabled, true
}

// HasQueueEnabled returns a boolean if a field has been set.
func (o *CreateInstanceRequest) HasQueueEnabled() bool {
	if o != nil && !IsNil(o.QueueEnabled) {
		return true
	}

	return false
}

// SetQueueEnabled gets a reference to the given bool and assigns it to the QueueEnabled field.
func (o *CreateInstanceRequest) SetQueueEnabled(v bool) {
	o.QueueEnabled = &v
}

// GetRegion returns the Region field value
func (o *CreateInstanceRequest) GetRegion() InstanceRegion {
	if o == nil {
		var ret InstanceRegion
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *CreateInstanceRequest) GetRegionOk() (*InstanceRegion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *CreateInstanceRequest) SetRegion(v InstanceRegion) {
	o.Region = v
}

// GetRoleIds returns the RoleIds field value if set, zero value otherwise.
func (o *CreateInstanceRequest) GetRoleIds() []string {
	if o == nil || IsNil(o.RoleIds) {
		var ret []string
		return ret
	}
	return o.RoleIds
}

// GetRoleIdsOk returns a tuple with the RoleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstanceRequest) GetRoleIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RoleIds) {
		return nil, false
	}
	return o.RoleIds, true
}

// HasRoleIds returns a boolean if a field has been set.
func (o *CreateInstanceRequest) HasRoleIds() bool {
	if o != nil && !IsNil(o.RoleIds) {
		return true
	}

	return false
}

// SetRoleIds gets a reference to the given []string and assigns it to the RoleIds field.
func (o *CreateInstanceRequest) SetRoleIds(v []string) {
	o.RoleIds = v
}

// GetType returns the Type field value
func (o *CreateInstanceRequest) GetType() InstanceType {
	if o == nil {
		var ret InstanceType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateInstanceRequest) GetTypeOk() (*InstanceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateInstanceRequest) SetType(v InstanceType) {
	o.Type = v
}

// GetWorldId returns the WorldId field value
func (o *CreateInstanceRequest) GetWorldId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorldId
}

// GetWorldIdOk returns a tuple with the WorldId field value
// and a boolean to check if the value has been set.
func (o *CreateInstanceRequest) GetWorldIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorldId, true
}

// SetWorldId sets field value
func (o *CreateInstanceRequest) SetWorldId(v string) {
	o.WorldId = v
}

func (o CreateInstanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateInstanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgeGate) {
		toSerialize["ageGate"] = o.AgeGate
	}
	if !IsNil(o.CalendarEntryId) {
		toSerialize["calendarEntryId"] = o.CalendarEntryId
	}
	if !IsNil(o.CanRequestInvite) {
		toSerialize["canRequestInvite"] = o.CanRequestInvite
	}
	if !IsNil(o.ClosedAt) {
		toSerialize["closedAt"] = o.ClosedAt
	}
	if !IsNil(o.ContentSettings) {
		toSerialize["contentSettings"] = o.ContentSettings
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if !IsNil(o.GroupAccessType) {
		toSerialize["groupAccessType"] = o.GroupAccessType
	}
	if !IsNil(o.HardClose) {
		toSerialize["hardClose"] = o.HardClose
	}
	if o.InstancePersistenceEnabled.IsSet() {
		toSerialize["instancePersistenceEnabled"] = o.InstancePersistenceEnabled.Get()
	}
	if !IsNil(o.InviteOnly) {
		toSerialize["inviteOnly"] = o.InviteOnly
	}
	if o.OwnerId.IsSet() {
		toSerialize["ownerId"] = o.OwnerId.Get()
	}
	if o.PlayerPersistenceEnabled.IsSet() {
		toSerialize["playerPersistenceEnabled"] = o.PlayerPersistenceEnabled.Get()
	}
	if !IsNil(o.QueueEnabled) {
		toSerialize["queueEnabled"] = o.QueueEnabled
	}
	toSerialize["region"] = o.Region
	if !IsNil(o.RoleIds) {
		toSerialize["roleIds"] = o.RoleIds
	}
	toSerialize["type"] = o.Type
	toSerialize["worldId"] = o.WorldId
	return toSerialize, nil
}

func (o *CreateInstanceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"region",
		"type",
		"worldId",
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

	varCreateInstanceRequest := _CreateInstanceRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateInstanceRequest)

	if err != nil {
		return err
	}

	*o = CreateInstanceRequest(varCreateInstanceRequest)

	return err
}

type NullableCreateInstanceRequest struct {
	value *CreateInstanceRequest
	isSet bool
}

func (v NullableCreateInstanceRequest) Get() *CreateInstanceRequest {
	return v.value
}

func (v *NullableCreateInstanceRequest) Set(val *CreateInstanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInstanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInstanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInstanceRequest(val *CreateInstanceRequest) *NullableCreateInstanceRequest {
	return &NullableCreateInstanceRequest{value: val, isSet: true}
}

func (v NullableCreateInstanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInstanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
