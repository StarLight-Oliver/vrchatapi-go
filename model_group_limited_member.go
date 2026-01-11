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
)

// checks if the GroupLimitedMember type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupLimitedMember{}

// GroupLimitedMember May be null when attempting to retrieve group membership for a user who is not part of the group
type GroupLimitedMember struct {
	// Only visible via the /groups/:groupId/members endpoint, **not** when fetching a specific user.
	BannedAt NullableTime `json:"bannedAt,omitempty"`
	// Only visible via the /groups/:groupId/members endpoint, **not** when fetching a specific user.
	CreatedAt             NullableTime `json:"createdAt,omitempty"`
	GroupId               *string      `json:"groupId,omitempty"`
	HasJoinedFromPurchase *bool        `json:"hasJoinedFromPurchase,omitempty"`
	Id                    *string      `json:"id,omitempty"`
	// Whether the user is representing the group. This makes the group show up above the name tag in-game.
	IsRepresenting                   *bool        `json:"isRepresenting,omitempty"`
	IsSubscribedToAnnouncements      *bool        `json:"isSubscribedToAnnouncements,omitempty"`
	IsSubscribedToEventAnnouncements *bool        `json:"isSubscribedToEventAnnouncements,omitempty"`
	JoinedAt                         NullableTime `json:"joinedAt,omitempty"`
	LastPostReadAt                   NullableTime `json:"lastPostReadAt,omitempty"`
	MRoleIds                         []string     `json:"mRoleIds,omitempty"`
	// Only visible via the /groups/:groupId/members endpoint, **not** when fetching a specific user.
	ManagerNotes     NullableString     `json:"managerNotes,omitempty"`
	MembershipStatus *GroupMemberStatus `json:"membershipStatus,omitempty"`
	RoleIds          []string           `json:"roleIds,omitempty"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	UserId     *string `json:"userId,omitempty"`
	Visibility *string `json:"visibility,omitempty"`
}

// NewGroupLimitedMember instantiates a new GroupLimitedMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupLimitedMember() *GroupLimitedMember {
	this := GroupLimitedMember{}
	var isRepresenting bool = false
	this.IsRepresenting = &isRepresenting
	var isSubscribedToAnnouncements bool = false
	this.IsSubscribedToAnnouncements = &isSubscribedToAnnouncements
	var membershipStatus GroupMemberStatus = GroupMemberStatus_INACTIVE
	this.MembershipStatus = &membershipStatus
	return &this
}

// NewGroupLimitedMemberWithDefaults instantiates a new GroupLimitedMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupLimitedMemberWithDefaults() *GroupLimitedMember {
	this := GroupLimitedMember{}
	var isRepresenting bool = false
	this.IsRepresenting = &isRepresenting
	var isSubscribedToAnnouncements bool = false
	this.IsSubscribedToAnnouncements = &isSubscribedToAnnouncements
	var membershipStatus GroupMemberStatus = GroupMemberStatus_INACTIVE
	this.MembershipStatus = &membershipStatus
	return &this
}

// GetBannedAt returns the BannedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupLimitedMember) GetBannedAt() time.Time {
	if o == nil || IsNil(o.BannedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.BannedAt.Get()
}

// GetBannedAtOk returns a tuple with the BannedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupLimitedMember) GetBannedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.BannedAt.Get(), o.BannedAt.IsSet()
}

// HasBannedAt returns a boolean if a field has been set.
func (o *GroupLimitedMember) HasBannedAt() bool {
	if o != nil && o.BannedAt.IsSet() {
		return true
	}

	return false
}

// SetBannedAt gets a reference to the given NullableTime and assigns it to the BannedAt field.
func (o *GroupLimitedMember) SetBannedAt(v time.Time) {
	o.BannedAt.Set(&v)
}

// SetBannedAtNil sets the value for BannedAt to be an explicit nil
func (o *GroupLimitedMember) SetBannedAtNil() {
	o.BannedAt.Set(nil)
}

// UnsetBannedAt ensures that no value is present for BannedAt, not even an explicit nil
func (o *GroupLimitedMember) UnsetBannedAt() {
	o.BannedAt.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupLimitedMember) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupLimitedMember) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GroupLimitedMember) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *GroupLimitedMember) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *GroupLimitedMember) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *GroupLimitedMember) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *GroupLimitedMember) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupLimitedMember) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *GroupLimitedMember) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *GroupLimitedMember) SetGroupId(v string) {
	o.GroupId = &v
}

// GetHasJoinedFromPurchase returns the HasJoinedFromPurchase field value if set, zero value otherwise.
func (o *GroupLimitedMember) GetHasJoinedFromPurchase() bool {
	if o == nil || IsNil(o.HasJoinedFromPurchase) {
		var ret bool
		return ret
	}
	return *o.HasJoinedFromPurchase
}

// GetHasJoinedFromPurchaseOk returns a tuple with the HasJoinedFromPurchase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupLimitedMember) GetHasJoinedFromPurchaseOk() (*bool, bool) {
	if o == nil || IsNil(o.HasJoinedFromPurchase) {
		return nil, false
	}
	return o.HasJoinedFromPurchase, true
}

// HasHasJoinedFromPurchase returns a boolean if a field has been set.
func (o *GroupLimitedMember) HasHasJoinedFromPurchase() bool {
	if o != nil && !IsNil(o.HasJoinedFromPurchase) {
		return true
	}

	return false
}

// SetHasJoinedFromPurchase gets a reference to the given bool and assigns it to the HasJoinedFromPurchase field.
func (o *GroupLimitedMember) SetHasJoinedFromPurchase(v bool) {
	o.HasJoinedFromPurchase = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GroupLimitedMember) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupLimitedMember) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GroupLimitedMember) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GroupLimitedMember) SetId(v string) {
	o.Id = &v
}

// GetIsRepresenting returns the IsRepresenting field value if set, zero value otherwise.
func (o *GroupLimitedMember) GetIsRepresenting() bool {
	if o == nil || IsNil(o.IsRepresenting) {
		var ret bool
		return ret
	}
	return *o.IsRepresenting
}

// GetIsRepresentingOk returns a tuple with the IsRepresenting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupLimitedMember) GetIsRepresentingOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRepresenting) {
		return nil, false
	}
	return o.IsRepresenting, true
}

// HasIsRepresenting returns a boolean if a field has been set.
func (o *GroupLimitedMember) HasIsRepresenting() bool {
	if o != nil && !IsNil(o.IsRepresenting) {
		return true
	}

	return false
}

// SetIsRepresenting gets a reference to the given bool and assigns it to the IsRepresenting field.
func (o *GroupLimitedMember) SetIsRepresenting(v bool) {
	o.IsRepresenting = &v
}

// GetIsSubscribedToAnnouncements returns the IsSubscribedToAnnouncements field value if set, zero value otherwise.
func (o *GroupLimitedMember) GetIsSubscribedToAnnouncements() bool {
	if o == nil || IsNil(o.IsSubscribedToAnnouncements) {
		var ret bool
		return ret
	}
	return *o.IsSubscribedToAnnouncements
}

// GetIsSubscribedToAnnouncementsOk returns a tuple with the IsSubscribedToAnnouncements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupLimitedMember) GetIsSubscribedToAnnouncementsOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSubscribedToAnnouncements) {
		return nil, false
	}
	return o.IsSubscribedToAnnouncements, true
}

// HasIsSubscribedToAnnouncements returns a boolean if a field has been set.
func (o *GroupLimitedMember) HasIsSubscribedToAnnouncements() bool {
	if o != nil && !IsNil(o.IsSubscribedToAnnouncements) {
		return true
	}

	return false
}

// SetIsSubscribedToAnnouncements gets a reference to the given bool and assigns it to the IsSubscribedToAnnouncements field.
func (o *GroupLimitedMember) SetIsSubscribedToAnnouncements(v bool) {
	o.IsSubscribedToAnnouncements = &v
}

// GetIsSubscribedToEventAnnouncements returns the IsSubscribedToEventAnnouncements field value if set, zero value otherwise.
func (o *GroupLimitedMember) GetIsSubscribedToEventAnnouncements() bool {
	if o == nil || IsNil(o.IsSubscribedToEventAnnouncements) {
		var ret bool
		return ret
	}
	return *o.IsSubscribedToEventAnnouncements
}

// GetIsSubscribedToEventAnnouncementsOk returns a tuple with the IsSubscribedToEventAnnouncements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupLimitedMember) GetIsSubscribedToEventAnnouncementsOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSubscribedToEventAnnouncements) {
		return nil, false
	}
	return o.IsSubscribedToEventAnnouncements, true
}

// HasIsSubscribedToEventAnnouncements returns a boolean if a field has been set.
func (o *GroupLimitedMember) HasIsSubscribedToEventAnnouncements() bool {
	if o != nil && !IsNil(o.IsSubscribedToEventAnnouncements) {
		return true
	}

	return false
}

// SetIsSubscribedToEventAnnouncements gets a reference to the given bool and assigns it to the IsSubscribedToEventAnnouncements field.
func (o *GroupLimitedMember) SetIsSubscribedToEventAnnouncements(v bool) {
	o.IsSubscribedToEventAnnouncements = &v
}

// GetJoinedAt returns the JoinedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupLimitedMember) GetJoinedAt() time.Time {
	if o == nil || IsNil(o.JoinedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.JoinedAt.Get()
}

// GetJoinedAtOk returns a tuple with the JoinedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupLimitedMember) GetJoinedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.JoinedAt.Get(), o.JoinedAt.IsSet()
}

// HasJoinedAt returns a boolean if a field has been set.
func (o *GroupLimitedMember) HasJoinedAt() bool {
	if o != nil && o.JoinedAt.IsSet() {
		return true
	}

	return false
}

// SetJoinedAt gets a reference to the given NullableTime and assigns it to the JoinedAt field.
func (o *GroupLimitedMember) SetJoinedAt(v time.Time) {
	o.JoinedAt.Set(&v)
}

// SetJoinedAtNil sets the value for JoinedAt to be an explicit nil
func (o *GroupLimitedMember) SetJoinedAtNil() {
	o.JoinedAt.Set(nil)
}

// UnsetJoinedAt ensures that no value is present for JoinedAt, not even an explicit nil
func (o *GroupLimitedMember) UnsetJoinedAt() {
	o.JoinedAt.Unset()
}

// GetLastPostReadAt returns the LastPostReadAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupLimitedMember) GetLastPostReadAt() time.Time {
	if o == nil || IsNil(o.LastPostReadAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastPostReadAt.Get()
}

// GetLastPostReadAtOk returns a tuple with the LastPostReadAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupLimitedMember) GetLastPostReadAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastPostReadAt.Get(), o.LastPostReadAt.IsSet()
}

// HasLastPostReadAt returns a boolean if a field has been set.
func (o *GroupLimitedMember) HasLastPostReadAt() bool {
	if o != nil && o.LastPostReadAt.IsSet() {
		return true
	}

	return false
}

// SetLastPostReadAt gets a reference to the given NullableTime and assigns it to the LastPostReadAt field.
func (o *GroupLimitedMember) SetLastPostReadAt(v time.Time) {
	o.LastPostReadAt.Set(&v)
}

// SetLastPostReadAtNil sets the value for LastPostReadAt to be an explicit nil
func (o *GroupLimitedMember) SetLastPostReadAtNil() {
	o.LastPostReadAt.Set(nil)
}

// UnsetLastPostReadAt ensures that no value is present for LastPostReadAt, not even an explicit nil
func (o *GroupLimitedMember) UnsetLastPostReadAt() {
	o.LastPostReadAt.Unset()
}

// GetMRoleIds returns the MRoleIds field value if set, zero value otherwise.
func (o *GroupLimitedMember) GetMRoleIds() []string {
	if o == nil || IsNil(o.MRoleIds) {
		var ret []string
		return ret
	}
	return o.MRoleIds
}

// GetMRoleIdsOk returns a tuple with the MRoleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupLimitedMember) GetMRoleIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.MRoleIds) {
		return nil, false
	}
	return o.MRoleIds, true
}

// HasMRoleIds returns a boolean if a field has been set.
func (o *GroupLimitedMember) HasMRoleIds() bool {
	if o != nil && !IsNil(o.MRoleIds) {
		return true
	}

	return false
}

// SetMRoleIds gets a reference to the given []string and assigns it to the MRoleIds field.
func (o *GroupLimitedMember) SetMRoleIds(v []string) {
	o.MRoleIds = v
}

// GetManagerNotes returns the ManagerNotes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupLimitedMember) GetManagerNotes() string {
	if o == nil || IsNil(o.ManagerNotes.Get()) {
		var ret string
		return ret
	}
	return *o.ManagerNotes.Get()
}

// GetManagerNotesOk returns a tuple with the ManagerNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupLimitedMember) GetManagerNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ManagerNotes.Get(), o.ManagerNotes.IsSet()
}

// HasManagerNotes returns a boolean if a field has been set.
func (o *GroupLimitedMember) HasManagerNotes() bool {
	if o != nil && o.ManagerNotes.IsSet() {
		return true
	}

	return false
}

// SetManagerNotes gets a reference to the given NullableString and assigns it to the ManagerNotes field.
func (o *GroupLimitedMember) SetManagerNotes(v string) {
	o.ManagerNotes.Set(&v)
}

// SetManagerNotesNil sets the value for ManagerNotes to be an explicit nil
func (o *GroupLimitedMember) SetManagerNotesNil() {
	o.ManagerNotes.Set(nil)
}

// UnsetManagerNotes ensures that no value is present for ManagerNotes, not even an explicit nil
func (o *GroupLimitedMember) UnsetManagerNotes() {
	o.ManagerNotes.Unset()
}

// GetMembershipStatus returns the MembershipStatus field value if set, zero value otherwise.
func (o *GroupLimitedMember) GetMembershipStatus() GroupMemberStatus {
	if o == nil || IsNil(o.MembershipStatus) {
		var ret GroupMemberStatus
		return ret
	}
	return *o.MembershipStatus
}

// GetMembershipStatusOk returns a tuple with the MembershipStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupLimitedMember) GetMembershipStatusOk() (*GroupMemberStatus, bool) {
	if o == nil || IsNil(o.MembershipStatus) {
		return nil, false
	}
	return o.MembershipStatus, true
}

// HasMembershipStatus returns a boolean if a field has been set.
func (o *GroupLimitedMember) HasMembershipStatus() bool {
	if o != nil && !IsNil(o.MembershipStatus) {
		return true
	}

	return false
}

// SetMembershipStatus gets a reference to the given GroupMemberStatus and assigns it to the MembershipStatus field.
func (o *GroupLimitedMember) SetMembershipStatus(v GroupMemberStatus) {
	o.MembershipStatus = &v
}

// GetRoleIds returns the RoleIds field value if set, zero value otherwise.
func (o *GroupLimitedMember) GetRoleIds() []string {
	if o == nil || IsNil(o.RoleIds) {
		var ret []string
		return ret
	}
	return o.RoleIds
}

// GetRoleIdsOk returns a tuple with the RoleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupLimitedMember) GetRoleIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RoleIds) {
		return nil, false
	}
	return o.RoleIds, true
}

// HasRoleIds returns a boolean if a field has been set.
func (o *GroupLimitedMember) HasRoleIds() bool {
	if o != nil && !IsNil(o.RoleIds) {
		return true
	}

	return false
}

// SetRoleIds gets a reference to the given []string and assigns it to the RoleIds field.
func (o *GroupLimitedMember) SetRoleIds(v []string) {
	o.RoleIds = v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *GroupLimitedMember) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupLimitedMember) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *GroupLimitedMember) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *GroupLimitedMember) SetUserId(v string) {
	o.UserId = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *GroupLimitedMember) GetVisibility() string {
	if o == nil || IsNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupLimitedMember) GetVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *GroupLimitedMember) HasVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *GroupLimitedMember) SetVisibility(v string) {
	o.Visibility = &v
}

func (o GroupLimitedMember) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupLimitedMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BannedAt.IsSet() {
		toSerialize["bannedAt"] = o.BannedAt.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["createdAt"] = o.CreatedAt.Get()
	}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !IsNil(o.HasJoinedFromPurchase) {
		toSerialize["hasJoinedFromPurchase"] = o.HasJoinedFromPurchase
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsRepresenting) {
		toSerialize["isRepresenting"] = o.IsRepresenting
	}
	if !IsNil(o.IsSubscribedToAnnouncements) {
		toSerialize["isSubscribedToAnnouncements"] = o.IsSubscribedToAnnouncements
	}
	if !IsNil(o.IsSubscribedToEventAnnouncements) {
		toSerialize["isSubscribedToEventAnnouncements"] = o.IsSubscribedToEventAnnouncements
	}
	if o.JoinedAt.IsSet() {
		toSerialize["joinedAt"] = o.JoinedAt.Get()
	}
	if o.LastPostReadAt.IsSet() {
		toSerialize["lastPostReadAt"] = o.LastPostReadAt.Get()
	}
	if !IsNil(o.MRoleIds) {
		toSerialize["mRoleIds"] = o.MRoleIds
	}
	if o.ManagerNotes.IsSet() {
		toSerialize["managerNotes"] = o.ManagerNotes.Get()
	}
	if !IsNil(o.MembershipStatus) {
		toSerialize["membershipStatus"] = o.MembershipStatus
	}
	if !IsNil(o.RoleIds) {
		toSerialize["roleIds"] = o.RoleIds
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	return toSerialize, nil
}

type NullableGroupLimitedMember struct {
	value *GroupLimitedMember
	isSet bool
}

func (v NullableGroupLimitedMember) Get() *GroupLimitedMember {
	return v.value
}

func (v *NullableGroupLimitedMember) Set(val *GroupLimitedMember) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupLimitedMember) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupLimitedMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupLimitedMember(val *GroupLimitedMember) *NullableGroupLimitedMember {
	return &NullableGroupLimitedMember{value: val, isSet: true}
}

func (v NullableGroupLimitedMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupLimitedMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
