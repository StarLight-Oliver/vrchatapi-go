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

// checks if the Instance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Instance{}

// Instance * `hidden` field is only present if InstanceType is `hidden` aka \"Friends+\", and is instance creator. * `friends` field is only present if InstanceType is `friends` aka \"Friends\", and is instance creator. * `private` field is only present if InstanceType is `private` aka \"Invite\" or \"Invite+\", and is instance creator.
type Instance struct {
	Active           *bool          `json:"active,omitempty"`
	AgeGate          NullableBool   `json:"ageGate,omitempty"`
	CalendarEntryId  NullableString `json:"calendarEntryId,omitempty"`
	CanRequestInvite *bool          `json:"canRequestInvite,omitempty"`
	Capacity         *int32         `json:"capacity,omitempty"`
	// Always returns \"unknown\".
	// Deprecated
	ClientNumber    string                   `json:"clientNumber"`
	ClosedAt        NullableTime             `json:"closedAt,omitempty"`
	ContentSettings *InstanceContentSettings `json:"contentSettings,omitempty"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	CreatorId   *string        `json:"creatorId,omitempty"`
	DisplayName NullableString `json:"displayName,omitempty"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	Friends           *string          `json:"friends,omitempty"`
	Full              bool             `json:"full"`
	GameServerVersion NullableInt32    `json:"gameServerVersion,omitempty"`
	GroupAccessType   *GroupAccessType `json:"groupAccessType,omitempty"`
	HardClose         NullableBool     `json:"hardClose,omitempty"`
	HasCapacityForYou *bool            `json:"hasCapacityForYou,omitempty"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	Hidden *string `json:"hidden,omitempty"`
	// InstanceID can be \"offline\" on User profiles if you are not friends with that user and \"private\" if you are friends and user is in private instance.
	Id string `json:"id"`
	// InstanceID can be \"offline\" on User profiles if you are not friends with that user and \"private\" if you are friends and user is in private instance.
	InstanceId                 string       `json:"instanceId"`
	InstancePersistenceEnabled NullableBool `json:"instancePersistenceEnabled,omitempty"`
	// Represents a unique location, consisting of a world identifier and an instance identifier, or \"offline\" if the user is not on your friends list.
	Location string  `json:"location"`
	NUsers   int32   `json:"n_users"`
	Name     string  `json:"name"`
	Nonce    *string `json:"nonce,omitempty"`
	// A groupId if the instance type is \"group\", null if instance type is public, or a userId otherwise
	OwnerId                  NullableString    `json:"ownerId,omitempty"`
	Permanent                bool              `json:"permanent"`
	PhotonRegion             Region            `json:"photonRegion"`
	Platforms                InstancePlatforms `json:"platforms"`
	PlayerPersistenceEnabled NullableBool      `json:"playerPersistenceEnabled,omitempty"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	Private             *string        `json:"private,omitempty"`
	QueueEnabled        bool           `json:"queueEnabled"`
	QueueSize           int32          `json:"queueSize"`
	RecommendedCapacity int32          `json:"recommendedCapacity"`
	Region              InstanceRegion `json:"region"`
	RoleRestricted      *bool          `json:"roleRestricted,omitempty"`
	SecureName          string         `json:"secureName"`
	ShortName           NullableString `json:"shortName,omitempty"`
	Strict              bool           `json:"strict"`
	// The tags array on Instances usually contain the language tags of the people in the instance.
	Tags      []string     `json:"tags"`
	Type      InstanceType `json:"type"`
	UserCount int32        `json:"userCount"`
	// The users field is present on instances created by the requesting user.
	Users []LimitedUserInstance `json:"users,omitempty"`
	World World                 `json:"world"`
	// WorldID be \"offline\" on User profiles if you are not friends with that user.
	WorldId string `json:"worldId"`
}

type _Instance Instance

// NewInstance instantiates a new Instance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstance(clientNumber string, full bool, id string, instanceId string, location string, nUsers int32, name string, permanent bool, photonRegion Region, platforms InstancePlatforms, queueEnabled bool, queueSize int32, recommendedCapacity int32, region InstanceRegion, secureName string, strict bool, tags []string, type_ InstanceType, userCount int32, world World, worldId string) *Instance {
	this := Instance{}
	var active bool = true
	this.Active = &active
	var canRequestInvite bool = true
	this.CanRequestInvite = &canRequestInvite
	this.ClientNumber = clientNumber
	this.Full = full
	var groupAccessType GroupAccessType = GroupAccessType_MEMBERS
	this.GroupAccessType = &groupAccessType
	this.Id = id
	this.InstanceId = instanceId
	this.Location = location
	this.NUsers = nUsers
	this.Name = name
	this.Permanent = permanent
	this.PhotonRegion = photonRegion
	this.Platforms = platforms
	this.QueueEnabled = queueEnabled
	this.QueueSize = queueSize
	this.RecommendedCapacity = recommendedCapacity
	this.Region = region
	this.SecureName = secureName
	this.Strict = strict
	this.Tags = tags
	this.Type = type_
	this.UserCount = userCount
	this.World = world
	this.WorldId = worldId
	return &this
}

// NewInstanceWithDefaults instantiates a new Instance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceWithDefaults() *Instance {
	this := Instance{}
	var active bool = true
	this.Active = &active
	var canRequestInvite bool = true
	this.CanRequestInvite = &canRequestInvite
	var full bool = false
	this.Full = full
	var groupAccessType GroupAccessType = GroupAccessType_MEMBERS
	this.GroupAccessType = &groupAccessType
	var permanent bool = false
	this.Permanent = permanent
	var photonRegion Region = Region_US
	this.PhotonRegion = photonRegion
	var region InstanceRegion = InstanceRegion_US
	this.Region = region
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *Instance) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *Instance) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *Instance) SetActive(v bool) {
	o.Active = &v
}

// GetAgeGate returns the AgeGate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetAgeGate() bool {
	if o == nil || IsNil(o.AgeGate.Get()) {
		var ret bool
		return ret
	}
	return *o.AgeGate.Get()
}

// GetAgeGateOk returns a tuple with the AgeGate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetAgeGateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgeGate.Get(), o.AgeGate.IsSet()
}

// HasAgeGate returns a boolean if a field has been set.
func (o *Instance) HasAgeGate() bool {
	if o != nil && o.AgeGate.IsSet() {
		return true
	}

	return false
}

// SetAgeGate gets a reference to the given NullableBool and assigns it to the AgeGate field.
func (o *Instance) SetAgeGate(v bool) {
	o.AgeGate.Set(&v)
}

// SetAgeGateNil sets the value for AgeGate to be an explicit nil
func (o *Instance) SetAgeGateNil() {
	o.AgeGate.Set(nil)
}

// UnsetAgeGate ensures that no value is present for AgeGate, not even an explicit nil
func (o *Instance) UnsetAgeGate() {
	o.AgeGate.Unset()
}

// GetCalendarEntryId returns the CalendarEntryId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetCalendarEntryId() string {
	if o == nil || IsNil(o.CalendarEntryId.Get()) {
		var ret string
		return ret
	}
	return *o.CalendarEntryId.Get()
}

// GetCalendarEntryIdOk returns a tuple with the CalendarEntryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetCalendarEntryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CalendarEntryId.Get(), o.CalendarEntryId.IsSet()
}

// HasCalendarEntryId returns a boolean if a field has been set.
func (o *Instance) HasCalendarEntryId() bool {
	if o != nil && o.CalendarEntryId.IsSet() {
		return true
	}

	return false
}

// SetCalendarEntryId gets a reference to the given NullableString and assigns it to the CalendarEntryId field.
func (o *Instance) SetCalendarEntryId(v string) {
	o.CalendarEntryId.Set(&v)
}

// SetCalendarEntryIdNil sets the value for CalendarEntryId to be an explicit nil
func (o *Instance) SetCalendarEntryIdNil() {
	o.CalendarEntryId.Set(nil)
}

// UnsetCalendarEntryId ensures that no value is present for CalendarEntryId, not even an explicit nil
func (o *Instance) UnsetCalendarEntryId() {
	o.CalendarEntryId.Unset()
}

// GetCanRequestInvite returns the CanRequestInvite field value if set, zero value otherwise.
func (o *Instance) GetCanRequestInvite() bool {
	if o == nil || IsNil(o.CanRequestInvite) {
		var ret bool
		return ret
	}
	return *o.CanRequestInvite
}

// GetCanRequestInviteOk returns a tuple with the CanRequestInvite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetCanRequestInviteOk() (*bool, bool) {
	if o == nil || IsNil(o.CanRequestInvite) {
		return nil, false
	}
	return o.CanRequestInvite, true
}

// HasCanRequestInvite returns a boolean if a field has been set.
func (o *Instance) HasCanRequestInvite() bool {
	if o != nil && !IsNil(o.CanRequestInvite) {
		return true
	}

	return false
}

// SetCanRequestInvite gets a reference to the given bool and assigns it to the CanRequestInvite field.
func (o *Instance) SetCanRequestInvite(v bool) {
	o.CanRequestInvite = &v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *Instance) GetCapacity() int32 {
	if o == nil || IsNil(o.Capacity) {
		var ret int32
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetCapacityOk() (*int32, bool) {
	if o == nil || IsNil(o.Capacity) {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *Instance) HasCapacity() bool {
	if o != nil && !IsNil(o.Capacity) {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given int32 and assigns it to the Capacity field.
func (o *Instance) SetCapacity(v int32) {
	o.Capacity = &v
}

// GetClientNumber returns the ClientNumber field value
// Deprecated
func (o *Instance) GetClientNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientNumber
}

// GetClientNumberOk returns a tuple with the ClientNumber field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *Instance) GetClientNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientNumber, true
}

// SetClientNumber sets field value
// Deprecated
func (o *Instance) SetClientNumber(v string) {
	o.ClientNumber = v
}

// GetClosedAt returns the ClosedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetClosedAt() time.Time {
	if o == nil || IsNil(o.ClosedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ClosedAt.Get()
}

// GetClosedAtOk returns a tuple with the ClosedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetClosedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClosedAt.Get(), o.ClosedAt.IsSet()
}

// HasClosedAt returns a boolean if a field has been set.
func (o *Instance) HasClosedAt() bool {
	if o != nil && o.ClosedAt.IsSet() {
		return true
	}

	return false
}

// SetClosedAt gets a reference to the given NullableTime and assigns it to the ClosedAt field.
func (o *Instance) SetClosedAt(v time.Time) {
	o.ClosedAt.Set(&v)
}

// SetClosedAtNil sets the value for ClosedAt to be an explicit nil
func (o *Instance) SetClosedAtNil() {
	o.ClosedAt.Set(nil)
}

// UnsetClosedAt ensures that no value is present for ClosedAt, not even an explicit nil
func (o *Instance) UnsetClosedAt() {
	o.ClosedAt.Unset()
}

// GetContentSettings returns the ContentSettings field value if set, zero value otherwise.
func (o *Instance) GetContentSettings() InstanceContentSettings {
	if o == nil || IsNil(o.ContentSettings) {
		var ret InstanceContentSettings
		return ret
	}
	return *o.ContentSettings
}

// GetContentSettingsOk returns a tuple with the ContentSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetContentSettingsOk() (*InstanceContentSettings, bool) {
	if o == nil || IsNil(o.ContentSettings) {
		return nil, false
	}
	return o.ContentSettings, true
}

// HasContentSettings returns a boolean if a field has been set.
func (o *Instance) HasContentSettings() bool {
	if o != nil && !IsNil(o.ContentSettings) {
		return true
	}

	return false
}

// SetContentSettings gets a reference to the given InstanceContentSettings and assigns it to the ContentSettings field.
func (o *Instance) SetContentSettings(v InstanceContentSettings) {
	o.ContentSettings = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *Instance) GetCreatorId() string {
	if o == nil || IsNil(o.CreatorId) {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetCreatorIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorId) {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *Instance) HasCreatorId() bool {
	if o != nil && !IsNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *Instance) SetCreatorId(v string) {
	o.CreatorId = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Instance) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *Instance) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}

// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *Instance) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *Instance) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetFriends returns the Friends field value if set, zero value otherwise.
func (o *Instance) GetFriends() string {
	if o == nil || IsNil(o.Friends) {
		var ret string
		return ret
	}
	return *o.Friends
}

// GetFriendsOk returns a tuple with the Friends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetFriendsOk() (*string, bool) {
	if o == nil || IsNil(o.Friends) {
		return nil, false
	}
	return o.Friends, true
}

// HasFriends returns a boolean if a field has been set.
func (o *Instance) HasFriends() bool {
	if o != nil && !IsNil(o.Friends) {
		return true
	}

	return false
}

// SetFriends gets a reference to the given string and assigns it to the Friends field.
func (o *Instance) SetFriends(v string) {
	o.Friends = &v
}

// GetFull returns the Full field value
func (o *Instance) GetFull() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Full
}

// GetFullOk returns a tuple with the Full field value
// and a boolean to check if the value has been set.
func (o *Instance) GetFullOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Full, true
}

// SetFull sets field value
func (o *Instance) SetFull(v bool) {
	o.Full = v
}

// GetGameServerVersion returns the GameServerVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetGameServerVersion() int32 {
	if o == nil || IsNil(o.GameServerVersion.Get()) {
		var ret int32
		return ret
	}
	return *o.GameServerVersion.Get()
}

// GetGameServerVersionOk returns a tuple with the GameServerVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetGameServerVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.GameServerVersion.Get(), o.GameServerVersion.IsSet()
}

// HasGameServerVersion returns a boolean if a field has been set.
func (o *Instance) HasGameServerVersion() bool {
	if o != nil && o.GameServerVersion.IsSet() {
		return true
	}

	return false
}

// SetGameServerVersion gets a reference to the given NullableInt32 and assigns it to the GameServerVersion field.
func (o *Instance) SetGameServerVersion(v int32) {
	o.GameServerVersion.Set(&v)
}

// SetGameServerVersionNil sets the value for GameServerVersion to be an explicit nil
func (o *Instance) SetGameServerVersionNil() {
	o.GameServerVersion.Set(nil)
}

// UnsetGameServerVersion ensures that no value is present for GameServerVersion, not even an explicit nil
func (o *Instance) UnsetGameServerVersion() {
	o.GameServerVersion.Unset()
}

// GetGroupAccessType returns the GroupAccessType field value if set, zero value otherwise.
func (o *Instance) GetGroupAccessType() GroupAccessType {
	if o == nil || IsNil(o.GroupAccessType) {
		var ret GroupAccessType
		return ret
	}
	return *o.GroupAccessType
}

// GetGroupAccessTypeOk returns a tuple with the GroupAccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetGroupAccessTypeOk() (*GroupAccessType, bool) {
	if o == nil || IsNil(o.GroupAccessType) {
		return nil, false
	}
	return o.GroupAccessType, true
}

// HasGroupAccessType returns a boolean if a field has been set.
func (o *Instance) HasGroupAccessType() bool {
	if o != nil && !IsNil(o.GroupAccessType) {
		return true
	}

	return false
}

// SetGroupAccessType gets a reference to the given GroupAccessType and assigns it to the GroupAccessType field.
func (o *Instance) SetGroupAccessType(v GroupAccessType) {
	o.GroupAccessType = &v
}

// GetHardClose returns the HardClose field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetHardClose() bool {
	if o == nil || IsNil(o.HardClose.Get()) {
		var ret bool
		return ret
	}
	return *o.HardClose.Get()
}

// GetHardCloseOk returns a tuple with the HardClose field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetHardCloseOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HardClose.Get(), o.HardClose.IsSet()
}

// HasHardClose returns a boolean if a field has been set.
func (o *Instance) HasHardClose() bool {
	if o != nil && o.HardClose.IsSet() {
		return true
	}

	return false
}

// SetHardClose gets a reference to the given NullableBool and assigns it to the HardClose field.
func (o *Instance) SetHardClose(v bool) {
	o.HardClose.Set(&v)
}

// SetHardCloseNil sets the value for HardClose to be an explicit nil
func (o *Instance) SetHardCloseNil() {
	o.HardClose.Set(nil)
}

// UnsetHardClose ensures that no value is present for HardClose, not even an explicit nil
func (o *Instance) UnsetHardClose() {
	o.HardClose.Unset()
}

// GetHasCapacityForYou returns the HasCapacityForYou field value if set, zero value otherwise.
func (o *Instance) GetHasCapacityForYou() bool {
	if o == nil || IsNil(o.HasCapacityForYou) {
		var ret bool
		return ret
	}
	return *o.HasCapacityForYou
}

// GetHasCapacityForYouOk returns a tuple with the HasCapacityForYou field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetHasCapacityForYouOk() (*bool, bool) {
	if o == nil || IsNil(o.HasCapacityForYou) {
		return nil, false
	}
	return o.HasCapacityForYou, true
}

// HasHasCapacityForYou returns a boolean if a field has been set.
func (o *Instance) HasHasCapacityForYou() bool {
	if o != nil && !IsNil(o.HasCapacityForYou) {
		return true
	}

	return false
}

// SetHasCapacityForYou gets a reference to the given bool and assigns it to the HasCapacityForYou field.
func (o *Instance) SetHasCapacityForYou(v bool) {
	o.HasCapacityForYou = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *Instance) GetHidden() string {
	if o == nil || IsNil(o.Hidden) {
		var ret string
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetHiddenOk() (*string, bool) {
	if o == nil || IsNil(o.Hidden) {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *Instance) HasHidden() bool {
	if o != nil && !IsNil(o.Hidden) {
		return true
	}

	return false
}

// SetHidden gets a reference to the given string and assigns it to the Hidden field.
func (o *Instance) SetHidden(v string) {
	o.Hidden = &v
}

// GetId returns the Id field value
func (o *Instance) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Instance) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Instance) SetId(v string) {
	o.Id = v
}

// GetInstanceId returns the InstanceId field value
func (o *Instance) GetInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *Instance) GetInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value
func (o *Instance) SetInstanceId(v string) {
	o.InstanceId = v
}

// GetInstancePersistenceEnabled returns the InstancePersistenceEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetInstancePersistenceEnabled() bool {
	if o == nil || IsNil(o.InstancePersistenceEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.InstancePersistenceEnabled.Get()
}

// GetInstancePersistenceEnabledOk returns a tuple with the InstancePersistenceEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetInstancePersistenceEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstancePersistenceEnabled.Get(), o.InstancePersistenceEnabled.IsSet()
}

// HasInstancePersistenceEnabled returns a boolean if a field has been set.
func (o *Instance) HasInstancePersistenceEnabled() bool {
	if o != nil && o.InstancePersistenceEnabled.IsSet() {
		return true
	}

	return false
}

// SetInstancePersistenceEnabled gets a reference to the given NullableBool and assigns it to the InstancePersistenceEnabled field.
func (o *Instance) SetInstancePersistenceEnabled(v bool) {
	o.InstancePersistenceEnabled.Set(&v)
}

// SetInstancePersistenceEnabledNil sets the value for InstancePersistenceEnabled to be an explicit nil
func (o *Instance) SetInstancePersistenceEnabledNil() {
	o.InstancePersistenceEnabled.Set(nil)
}

// UnsetInstancePersistenceEnabled ensures that no value is present for InstancePersistenceEnabled, not even an explicit nil
func (o *Instance) UnsetInstancePersistenceEnabled() {
	o.InstancePersistenceEnabled.Unset()
}

// GetLocation returns the Location field value
func (o *Instance) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *Instance) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *Instance) SetLocation(v string) {
	o.Location = v
}

// GetNUsers returns the NUsers field value
func (o *Instance) GetNUsers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NUsers
}

// GetNUsersOk returns a tuple with the NUsers field value
// and a boolean to check if the value has been set.
func (o *Instance) GetNUsersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NUsers, true
}

// SetNUsers sets field value
func (o *Instance) SetNUsers(v int32) {
	o.NUsers = v
}

// GetName returns the Name field value
func (o *Instance) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Instance) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Instance) SetName(v string) {
	o.Name = v
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *Instance) GetNonce() string {
	if o == nil || IsNil(o.Nonce) {
		var ret string
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetNonceOk() (*string, bool) {
	if o == nil || IsNil(o.Nonce) {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *Instance) HasNonce() bool {
	if o != nil && !IsNil(o.Nonce) {
		return true
	}

	return false
}

// SetNonce gets a reference to the given string and assigns it to the Nonce field.
func (o *Instance) SetNonce(v string) {
	o.Nonce = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId.Get()) {
		var ret string
		return ret
	}
	return *o.OwnerId.Get()
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerId.Get(), o.OwnerId.IsSet()
}

// HasOwnerId returns a boolean if a field has been set.
func (o *Instance) HasOwnerId() bool {
	if o != nil && o.OwnerId.IsSet() {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given NullableString and assigns it to the OwnerId field.
func (o *Instance) SetOwnerId(v string) {
	o.OwnerId.Set(&v)
}

// SetOwnerIdNil sets the value for OwnerId to be an explicit nil
func (o *Instance) SetOwnerIdNil() {
	o.OwnerId.Set(nil)
}

// UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
func (o *Instance) UnsetOwnerId() {
	o.OwnerId.Unset()
}

// GetPermanent returns the Permanent field value
func (o *Instance) GetPermanent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Permanent
}

// GetPermanentOk returns a tuple with the Permanent field value
// and a boolean to check if the value has been set.
func (o *Instance) GetPermanentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permanent, true
}

// SetPermanent sets field value
func (o *Instance) SetPermanent(v bool) {
	o.Permanent = v
}

// GetPhotonRegion returns the PhotonRegion field value
func (o *Instance) GetPhotonRegion() Region {
	if o == nil {
		var ret Region
		return ret
	}

	return o.PhotonRegion
}

// GetPhotonRegionOk returns a tuple with the PhotonRegion field value
// and a boolean to check if the value has been set.
func (o *Instance) GetPhotonRegionOk() (*Region, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhotonRegion, true
}

// SetPhotonRegion sets field value
func (o *Instance) SetPhotonRegion(v Region) {
	o.PhotonRegion = v
}

// GetPlatforms returns the Platforms field value
func (o *Instance) GetPlatforms() InstancePlatforms {
	if o == nil {
		var ret InstancePlatforms
		return ret
	}

	return o.Platforms
}

// GetPlatformsOk returns a tuple with the Platforms field value
// and a boolean to check if the value has been set.
func (o *Instance) GetPlatformsOk() (*InstancePlatforms, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Platforms, true
}

// SetPlatforms sets field value
func (o *Instance) SetPlatforms(v InstancePlatforms) {
	o.Platforms = v
}

// GetPlayerPersistenceEnabled returns the PlayerPersistenceEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetPlayerPersistenceEnabled() bool {
	if o == nil || IsNil(o.PlayerPersistenceEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.PlayerPersistenceEnabled.Get()
}

// GetPlayerPersistenceEnabledOk returns a tuple with the PlayerPersistenceEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetPlayerPersistenceEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlayerPersistenceEnabled.Get(), o.PlayerPersistenceEnabled.IsSet()
}

// HasPlayerPersistenceEnabled returns a boolean if a field has been set.
func (o *Instance) HasPlayerPersistenceEnabled() bool {
	if o != nil && o.PlayerPersistenceEnabled.IsSet() {
		return true
	}

	return false
}

// SetPlayerPersistenceEnabled gets a reference to the given NullableBool and assigns it to the PlayerPersistenceEnabled field.
func (o *Instance) SetPlayerPersistenceEnabled(v bool) {
	o.PlayerPersistenceEnabled.Set(&v)
}

// SetPlayerPersistenceEnabledNil sets the value for PlayerPersistenceEnabled to be an explicit nil
func (o *Instance) SetPlayerPersistenceEnabledNil() {
	o.PlayerPersistenceEnabled.Set(nil)
}

// UnsetPlayerPersistenceEnabled ensures that no value is present for PlayerPersistenceEnabled, not even an explicit nil
func (o *Instance) UnsetPlayerPersistenceEnabled() {
	o.PlayerPersistenceEnabled.Unset()
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *Instance) GetPrivate() string {
	if o == nil || IsNil(o.Private) {
		var ret string
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetPrivateOk() (*string, bool) {
	if o == nil || IsNil(o.Private) {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *Instance) HasPrivate() bool {
	if o != nil && !IsNil(o.Private) {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given string and assigns it to the Private field.
func (o *Instance) SetPrivate(v string) {
	o.Private = &v
}

// GetQueueEnabled returns the QueueEnabled field value
func (o *Instance) GetQueueEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.QueueEnabled
}

// GetQueueEnabledOk returns a tuple with the QueueEnabled field value
// and a boolean to check if the value has been set.
func (o *Instance) GetQueueEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueueEnabled, true
}

// SetQueueEnabled sets field value
func (o *Instance) SetQueueEnabled(v bool) {
	o.QueueEnabled = v
}

// GetQueueSize returns the QueueSize field value
func (o *Instance) GetQueueSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.QueueSize
}

// GetQueueSizeOk returns a tuple with the QueueSize field value
// and a boolean to check if the value has been set.
func (o *Instance) GetQueueSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueueSize, true
}

// SetQueueSize sets field value
func (o *Instance) SetQueueSize(v int32) {
	o.QueueSize = v
}

// GetRecommendedCapacity returns the RecommendedCapacity field value
func (o *Instance) GetRecommendedCapacity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RecommendedCapacity
}

// GetRecommendedCapacityOk returns a tuple with the RecommendedCapacity field value
// and a boolean to check if the value has been set.
func (o *Instance) GetRecommendedCapacityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecommendedCapacity, true
}

// SetRecommendedCapacity sets field value
func (o *Instance) SetRecommendedCapacity(v int32) {
	o.RecommendedCapacity = v
}

// GetRegion returns the Region field value
func (o *Instance) GetRegion() InstanceRegion {
	if o == nil {
		var ret InstanceRegion
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *Instance) GetRegionOk() (*InstanceRegion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *Instance) SetRegion(v InstanceRegion) {
	o.Region = v
}

// GetRoleRestricted returns the RoleRestricted field value if set, zero value otherwise.
func (o *Instance) GetRoleRestricted() bool {
	if o == nil || IsNil(o.RoleRestricted) {
		var ret bool
		return ret
	}
	return *o.RoleRestricted
}

// GetRoleRestrictedOk returns a tuple with the RoleRestricted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetRoleRestrictedOk() (*bool, bool) {
	if o == nil || IsNil(o.RoleRestricted) {
		return nil, false
	}
	return o.RoleRestricted, true
}

// HasRoleRestricted returns a boolean if a field has been set.
func (o *Instance) HasRoleRestricted() bool {
	if o != nil && !IsNil(o.RoleRestricted) {
		return true
	}

	return false
}

// SetRoleRestricted gets a reference to the given bool and assigns it to the RoleRestricted field.
func (o *Instance) SetRoleRestricted(v bool) {
	o.RoleRestricted = &v
}

// GetSecureName returns the SecureName field value
func (o *Instance) GetSecureName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecureName
}

// GetSecureNameOk returns a tuple with the SecureName field value
// and a boolean to check if the value has been set.
func (o *Instance) GetSecureNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecureName, true
}

// SetSecureName sets field value
func (o *Instance) SetSecureName(v string) {
	o.SecureName = v
}

// GetShortName returns the ShortName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Instance) GetShortName() string {
	if o == nil || IsNil(o.ShortName.Get()) {
		var ret string
		return ret
	}
	return *o.ShortName.Get()
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Instance) GetShortNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShortName.Get(), o.ShortName.IsSet()
}

// HasShortName returns a boolean if a field has been set.
func (o *Instance) HasShortName() bool {
	if o != nil && o.ShortName.IsSet() {
		return true
	}

	return false
}

// SetShortName gets a reference to the given NullableString and assigns it to the ShortName field.
func (o *Instance) SetShortName(v string) {
	o.ShortName.Set(&v)
}

// SetShortNameNil sets the value for ShortName to be an explicit nil
func (o *Instance) SetShortNameNil() {
	o.ShortName.Set(nil)
}

// UnsetShortName ensures that no value is present for ShortName, not even an explicit nil
func (o *Instance) UnsetShortName() {
	o.ShortName.Unset()
}

// GetStrict returns the Strict field value
func (o *Instance) GetStrict() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Strict
}

// GetStrictOk returns a tuple with the Strict field value
// and a boolean to check if the value has been set.
func (o *Instance) GetStrictOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Strict, true
}

// SetStrict sets field value
func (o *Instance) SetStrict(v bool) {
	o.Strict = v
}

// GetTags returns the Tags field value
func (o *Instance) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *Instance) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *Instance) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value
func (o *Instance) GetType() InstanceType {
	if o == nil {
		var ret InstanceType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Instance) GetTypeOk() (*InstanceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Instance) SetType(v InstanceType) {
	o.Type = v
}

// GetUserCount returns the UserCount field value
func (o *Instance) GetUserCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserCount
}

// GetUserCountOk returns a tuple with the UserCount field value
// and a boolean to check if the value has been set.
func (o *Instance) GetUserCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserCount, true
}

// SetUserCount sets field value
func (o *Instance) SetUserCount(v int32) {
	o.UserCount = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *Instance) GetUsers() []LimitedUserInstance {
	if o == nil || IsNil(o.Users) {
		var ret []LimitedUserInstance
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Instance) GetUsersOk() ([]LimitedUserInstance, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *Instance) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []LimitedUserInstance and assigns it to the Users field.
func (o *Instance) SetUsers(v []LimitedUserInstance) {
	o.Users = v
}

// GetWorld returns the World field value
func (o *Instance) GetWorld() World {
	if o == nil {
		var ret World
		return ret
	}

	return o.World
}

// GetWorldOk returns a tuple with the World field value
// and a boolean to check if the value has been set.
func (o *Instance) GetWorldOk() (*World, bool) {
	if o == nil {
		return nil, false
	}
	return &o.World, true
}

// SetWorld sets field value
func (o *Instance) SetWorld(v World) {
	o.World = v
}

// GetWorldId returns the WorldId field value
func (o *Instance) GetWorldId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorldId
}

// GetWorldIdOk returns a tuple with the WorldId field value
// and a boolean to check if the value has been set.
func (o *Instance) GetWorldIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorldId, true
}

// SetWorldId sets field value
func (o *Instance) SetWorldId(v string) {
	o.WorldId = v
}

func (o Instance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Instance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if o.AgeGate.IsSet() {
		toSerialize["ageGate"] = o.AgeGate.Get()
	}
	if o.CalendarEntryId.IsSet() {
		toSerialize["calendarEntryId"] = o.CalendarEntryId.Get()
	}
	if !IsNil(o.CanRequestInvite) {
		toSerialize["canRequestInvite"] = o.CanRequestInvite
	}
	if !IsNil(o.Capacity) {
		toSerialize["capacity"] = o.Capacity
	}
	toSerialize["clientNumber"] = o.ClientNumber
	if o.ClosedAt.IsSet() {
		toSerialize["closedAt"] = o.ClosedAt.Get()
	}
	if !IsNil(o.ContentSettings) {
		toSerialize["contentSettings"] = o.ContentSettings
	}
	if !IsNil(o.CreatorId) {
		toSerialize["creatorId"] = o.CreatorId
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if !IsNil(o.Friends) {
		toSerialize["friends"] = o.Friends
	}
	toSerialize["full"] = o.Full
	if o.GameServerVersion.IsSet() {
		toSerialize["gameServerVersion"] = o.GameServerVersion.Get()
	}
	if !IsNil(o.GroupAccessType) {
		toSerialize["groupAccessType"] = o.GroupAccessType
	}
	if o.HardClose.IsSet() {
		toSerialize["hardClose"] = o.HardClose.Get()
	}
	if !IsNil(o.HasCapacityForYou) {
		toSerialize["hasCapacityForYou"] = o.HasCapacityForYou
	}
	if !IsNil(o.Hidden) {
		toSerialize["hidden"] = o.Hidden
	}
	toSerialize["id"] = o.Id
	toSerialize["instanceId"] = o.InstanceId
	if o.InstancePersistenceEnabled.IsSet() {
		toSerialize["instancePersistenceEnabled"] = o.InstancePersistenceEnabled.Get()
	}
	toSerialize["location"] = o.Location
	toSerialize["n_users"] = o.NUsers
	toSerialize["name"] = o.Name
	if !IsNil(o.Nonce) {
		toSerialize["nonce"] = o.Nonce
	}
	if o.OwnerId.IsSet() {
		toSerialize["ownerId"] = o.OwnerId.Get()
	}
	toSerialize["permanent"] = o.Permanent
	toSerialize["photonRegion"] = o.PhotonRegion
	toSerialize["platforms"] = o.Platforms
	if o.PlayerPersistenceEnabled.IsSet() {
		toSerialize["playerPersistenceEnabled"] = o.PlayerPersistenceEnabled.Get()
	}
	if !IsNil(o.Private) {
		toSerialize["private"] = o.Private
	}
	toSerialize["queueEnabled"] = o.QueueEnabled
	toSerialize["queueSize"] = o.QueueSize
	toSerialize["recommendedCapacity"] = o.RecommendedCapacity
	toSerialize["region"] = o.Region
	if !IsNil(o.RoleRestricted) {
		toSerialize["roleRestricted"] = o.RoleRestricted
	}
	toSerialize["secureName"] = o.SecureName
	if o.ShortName.IsSet() {
		toSerialize["shortName"] = o.ShortName.Get()
	}
	toSerialize["strict"] = o.Strict
	toSerialize["tags"] = o.Tags
	toSerialize["type"] = o.Type
	toSerialize["userCount"] = o.UserCount
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	toSerialize["world"] = o.World
	toSerialize["worldId"] = o.WorldId
	return toSerialize, nil
}

func (o *Instance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"clientNumber",
		"full",
		"id",
		"instanceId",
		"location",
		"n_users",
		"name",
		"permanent",
		"photonRegion",
		"platforms",
		"queueEnabled",
		"queueSize",
		"recommendedCapacity",
		"region",
		"secureName",
		"strict",
		"tags",
		"type",
		"userCount",
		"world",
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

	varInstance := _Instance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInstance)

	if err != nil {
		return err
	}

	*o = Instance(varInstance)

	return err
}

type NullableInstance struct {
	value *Instance
	isSet bool
}

func (v NullableInstance) Get() *Instance {
	return v.value
}

func (v *NullableInstance) Set(val *Instance) {
	v.value = val
	v.isSet = true
}

func (v NullableInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstance(val *Instance) *NullableInstance {
	return &NullableInstance{value: val, isSet: true}
}

func (v NullableInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
