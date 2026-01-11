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

// checks if the CalendarEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CalendarEvent{}

// CalendarEvent An event scheduled on a group's calendar
type CalendarEvent struct {
	AccessType                   CalendarEventAccess   `json:"accessType"`
	Category                     CalendarEventCategory `json:"category"`
	CloseInstanceAfterEndMinutes *int32                `json:"closeInstanceAfterEndMinutes,omitempty"`
	CreatedAt                    *time.Time            `json:"createdAt,omitempty"`
	DeletedAt                    NullableTime          `json:"deletedAt,omitempty"`
	Description                  string                `json:"description"`
	DurationInMs                 *int32                `json:"durationInMs,omitempty"`
	EndsAt                       time.Time             `json:"endsAt"`
	Featured                     *bool                 `json:"featured,omitempty"`
	GuestEarlyJoinMinutes        *int32                `json:"guestEarlyJoinMinutes,omitempty"`
	HostEarlyJoinMinutes         *int32                `json:"hostEarlyJoinMinutes,omitempty"`
	Id                           string                `json:"id"`
	ImageId                      *string               `json:"imageId,omitempty"`
	ImageUrl                     NullableString        `json:"imageUrl,omitempty"`
	InterestedUserCount          *int32                `json:"interestedUserCount,omitempty"`
	IsDraft                      *bool                 `json:"isDraft,omitempty"`
	// Languages that might be spoken at this event
	Languages []string                `json:"languages,omitempty"`
	OwnerId   *string                 `json:"ownerId,omitempty"`
	Platforms []CalendarEventPlatform `json:"platforms,omitempty"`
	// Group roles that may join this event
	RoleIds  []string  `json:"roleIds,omitempty"`
	StartsAt time.Time `json:"startsAt"`
	// Custom tags for this event
	Tags                 []string                   `json:"tags,omitempty"`
	Title                string                     `json:"title"`
	Type                 *string                    `json:"type,omitempty"`
	UpdatedAt            *time.Time                 `json:"updatedAt,omitempty"`
	UserInterest         *CalendarEventUserInterest `json:"userInterest,omitempty"`
	UsesInstanceOverflow *bool                      `json:"usesInstanceOverflow,omitempty"`
}

type _CalendarEvent CalendarEvent

// NewCalendarEvent instantiates a new CalendarEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCalendarEvent(accessType CalendarEventAccess, category CalendarEventCategory, description string, endsAt time.Time, id string, startsAt time.Time, title string) *CalendarEvent {
	this := CalendarEvent{}
	this.AccessType = accessType
	this.Category = category
	this.Description = description
	this.EndsAt = endsAt
	this.Id = id
	this.StartsAt = startsAt
	this.Title = title
	return &this
}

// NewCalendarEventWithDefaults instantiates a new CalendarEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCalendarEventWithDefaults() *CalendarEvent {
	this := CalendarEvent{}
	var accessType CalendarEventAccess = CalendarEventAccess_PUBLIC
	this.AccessType = accessType
	var category CalendarEventCategory = CalendarEventCategory_OTHER
	this.Category = category
	return &this
}

// GetAccessType returns the AccessType field value
func (o *CalendarEvent) GetAccessType() CalendarEventAccess {
	if o == nil {
		var ret CalendarEventAccess
		return ret
	}

	return o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value
// and a boolean to check if the value has been set.
func (o *CalendarEvent) GetAccessTypeOk() (*CalendarEventAccess, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessType, true
}

// SetAccessType sets field value
func (o *CalendarEvent) SetAccessType(v CalendarEventAccess) {
	o.AccessType = v
}

// GetCategory returns the Category field value
func (o *CalendarEvent) GetCategory() CalendarEventCategory {
	if o == nil {
		var ret CalendarEventCategory
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *CalendarEvent) GetCategoryOk() (*CalendarEventCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *CalendarEvent) SetCategory(v CalendarEventCategory) {
	o.Category = v
}

// GetCloseInstanceAfterEndMinutes returns the CloseInstanceAfterEndMinutes field value if set, zero value otherwise.
func (o *CalendarEvent) GetCloseInstanceAfterEndMinutes() int32 {
	if o == nil || IsNil(o.CloseInstanceAfterEndMinutes) {
		var ret int32
		return ret
	}
	return *o.CloseInstanceAfterEndMinutes
}

// GetCloseInstanceAfterEndMinutesOk returns a tuple with the CloseInstanceAfterEndMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarEvent) GetCloseInstanceAfterEndMinutesOk() (*int32, bool) {
	if o == nil || IsNil(o.CloseInstanceAfterEndMinutes) {
		return nil, false
	}
	return o.CloseInstanceAfterEndMinutes, true
}

// HasCloseInstanceAfterEndMinutes returns a boolean if a field has been set.
func (o *CalendarEvent) HasCloseInstanceAfterEndMinutes() bool {
	if o != nil && !IsNil(o.CloseInstanceAfterEndMinutes) {
		return true
	}

	return false
}

// SetCloseInstanceAfterEndMinutes gets a reference to the given int32 and assigns it to the CloseInstanceAfterEndMinutes field.
func (o *CalendarEvent) SetCloseInstanceAfterEndMinutes(v int32) {
	o.CloseInstanceAfterEndMinutes = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CalendarEvent) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarEvent) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CalendarEvent) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CalendarEvent) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CalendarEvent) GetDeletedAt() time.Time {
	if o == nil || IsNil(o.DeletedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CalendarEvent) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *CalendarEvent) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *CalendarEvent) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *CalendarEvent) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *CalendarEvent) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetDescription returns the Description field value
func (o *CalendarEvent) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CalendarEvent) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CalendarEvent) SetDescription(v string) {
	o.Description = v
}

// GetDurationInMs returns the DurationInMs field value if set, zero value otherwise.
func (o *CalendarEvent) GetDurationInMs() int32 {
	if o == nil || IsNil(o.DurationInMs) {
		var ret int32
		return ret
	}
	return *o.DurationInMs
}

// GetDurationInMsOk returns a tuple with the DurationInMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarEvent) GetDurationInMsOk() (*int32, bool) {
	if o == nil || IsNil(o.DurationInMs) {
		return nil, false
	}
	return o.DurationInMs, true
}

// HasDurationInMs returns a boolean if a field has been set.
func (o *CalendarEvent) HasDurationInMs() bool {
	if o != nil && !IsNil(o.DurationInMs) {
		return true
	}

	return false
}

// SetDurationInMs gets a reference to the given int32 and assigns it to the DurationInMs field.
func (o *CalendarEvent) SetDurationInMs(v int32) {
	o.DurationInMs = &v
}

// GetEndsAt returns the EndsAt field value
func (o *CalendarEvent) GetEndsAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndsAt
}

// GetEndsAtOk returns a tuple with the EndsAt field value
// and a boolean to check if the value has been set.
func (o *CalendarEvent) GetEndsAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndsAt, true
}

// SetEndsAt sets field value
func (o *CalendarEvent) SetEndsAt(v time.Time) {
	o.EndsAt = v
}

// GetFeatured returns the Featured field value if set, zero value otherwise.
func (o *CalendarEvent) GetFeatured() bool {
	if o == nil || IsNil(o.Featured) {
		var ret bool
		return ret
	}
	return *o.Featured
}

// GetFeaturedOk returns a tuple with the Featured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarEvent) GetFeaturedOk() (*bool, bool) {
	if o == nil || IsNil(o.Featured) {
		return nil, false
	}
	return o.Featured, true
}

// HasFeatured returns a boolean if a field has been set.
func (o *CalendarEvent) HasFeatured() bool {
	if o != nil && !IsNil(o.Featured) {
		return true
	}

	return false
}

// SetFeatured gets a reference to the given bool and assigns it to the Featured field.
func (o *CalendarEvent) SetFeatured(v bool) {
	o.Featured = &v
}

// GetGuestEarlyJoinMinutes returns the GuestEarlyJoinMinutes field value if set, zero value otherwise.
func (o *CalendarEvent) GetGuestEarlyJoinMinutes() int32 {
	if o == nil || IsNil(o.GuestEarlyJoinMinutes) {
		var ret int32
		return ret
	}
	return *o.GuestEarlyJoinMinutes
}

// GetGuestEarlyJoinMinutesOk returns a tuple with the GuestEarlyJoinMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarEvent) GetGuestEarlyJoinMinutesOk() (*int32, bool) {
	if o == nil || IsNil(o.GuestEarlyJoinMinutes) {
		return nil, false
	}
	return o.GuestEarlyJoinMinutes, true
}

// HasGuestEarlyJoinMinutes returns a boolean if a field has been set.
func (o *CalendarEvent) HasGuestEarlyJoinMinutes() bool {
	if o != nil && !IsNil(o.GuestEarlyJoinMinutes) {
		return true
	}

	return false
}

// SetGuestEarlyJoinMinutes gets a reference to the given int32 and assigns it to the GuestEarlyJoinMinutes field.
func (o *CalendarEvent) SetGuestEarlyJoinMinutes(v int32) {
	o.GuestEarlyJoinMinutes = &v
}

// GetHostEarlyJoinMinutes returns the HostEarlyJoinMinutes field value if set, zero value otherwise.
func (o *CalendarEvent) GetHostEarlyJoinMinutes() int32 {
	if o == nil || IsNil(o.HostEarlyJoinMinutes) {
		var ret int32
		return ret
	}
	return *o.HostEarlyJoinMinutes
}

// GetHostEarlyJoinMinutesOk returns a tuple with the HostEarlyJoinMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarEvent) GetHostEarlyJoinMinutesOk() (*int32, bool) {
	if o == nil || IsNil(o.HostEarlyJoinMinutes) {
		return nil, false
	}
	return o.HostEarlyJoinMinutes, true
}

// HasHostEarlyJoinMinutes returns a boolean if a field has been set.
func (o *CalendarEvent) HasHostEarlyJoinMinutes() bool {
	if o != nil && !IsNil(o.HostEarlyJoinMinutes) {
		return true
	}

	return false
}

// SetHostEarlyJoinMinutes gets a reference to the given int32 and assigns it to the HostEarlyJoinMinutes field.
func (o *CalendarEvent) SetHostEarlyJoinMinutes(v int32) {
	o.HostEarlyJoinMinutes = &v
}

// GetId returns the Id field value
func (o *CalendarEvent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CalendarEvent) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CalendarEvent) SetId(v string) {
	o.Id = v
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *CalendarEvent) GetImageId() string {
	if o == nil || IsNil(o.ImageId) {
		var ret string
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarEvent) GetImageIdOk() (*string, bool) {
	if o == nil || IsNil(o.ImageId) {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *CalendarEvent) HasImageId() bool {
	if o != nil && !IsNil(o.ImageId) {
		return true
	}

	return false
}

// SetImageId gets a reference to the given string and assigns it to the ImageId field.
func (o *CalendarEvent) SetImageId(v string) {
	o.ImageId = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CalendarEvent) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ImageUrl.Get()
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CalendarEvent) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageUrl.Get(), o.ImageUrl.IsSet()
}

// HasImageUrl returns a boolean if a field has been set.
func (o *CalendarEvent) HasImageUrl() bool {
	if o != nil && o.ImageUrl.IsSet() {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given NullableString and assigns it to the ImageUrl field.
func (o *CalendarEvent) SetImageUrl(v string) {
	o.ImageUrl.Set(&v)
}

// SetImageUrlNil sets the value for ImageUrl to be an explicit nil
func (o *CalendarEvent) SetImageUrlNil() {
	o.ImageUrl.Set(nil)
}

// UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
func (o *CalendarEvent) UnsetImageUrl() {
	o.ImageUrl.Unset()
}

// GetInterestedUserCount returns the InterestedUserCount field value if set, zero value otherwise.
func (o *CalendarEvent) GetInterestedUserCount() int32 {
	if o == nil || IsNil(o.InterestedUserCount) {
		var ret int32
		return ret
	}
	return *o.InterestedUserCount
}

// GetInterestedUserCountOk returns a tuple with the InterestedUserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarEvent) GetInterestedUserCountOk() (*int32, bool) {
	if o == nil || IsNil(o.InterestedUserCount) {
		return nil, false
	}
	return o.InterestedUserCount, true
}

// HasInterestedUserCount returns a boolean if a field has been set.
func (o *CalendarEvent) HasInterestedUserCount() bool {
	if o != nil && !IsNil(o.InterestedUserCount) {
		return true
	}

	return false
}

// SetInterestedUserCount gets a reference to the given int32 and assigns it to the InterestedUserCount field.
func (o *CalendarEvent) SetInterestedUserCount(v int32) {
	o.InterestedUserCount = &v
}

// GetIsDraft returns the IsDraft field value if set, zero value otherwise.
func (o *CalendarEvent) GetIsDraft() bool {
	if o == nil || IsNil(o.IsDraft) {
		var ret bool
		return ret
	}
	return *o.IsDraft
}

// GetIsDraftOk returns a tuple with the IsDraft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarEvent) GetIsDraftOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDraft) {
		return nil, false
	}
	return o.IsDraft, true
}

// HasIsDraft returns a boolean if a field has been set.
func (o *CalendarEvent) HasIsDraft() bool {
	if o != nil && !IsNil(o.IsDraft) {
		return true
	}

	return false
}

// SetIsDraft gets a reference to the given bool and assigns it to the IsDraft field.
func (o *CalendarEvent) SetIsDraft(v bool) {
	o.IsDraft = &v
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *CalendarEvent) GetLanguages() []string {
	if o == nil || IsNil(o.Languages) {
		var ret []string
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarEvent) GetLanguagesOk() ([]string, bool) {
	if o == nil || IsNil(o.Languages) {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *CalendarEvent) HasLanguages() bool {
	if o != nil && !IsNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []string and assigns it to the Languages field.
func (o *CalendarEvent) SetLanguages(v []string) {
	o.Languages = v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *CalendarEvent) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarEvent) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *CalendarEvent) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *CalendarEvent) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetPlatforms returns the Platforms field value if set, zero value otherwise.
func (o *CalendarEvent) GetPlatforms() []CalendarEventPlatform {
	if o == nil || IsNil(o.Platforms) {
		var ret []CalendarEventPlatform
		return ret
	}
	return o.Platforms
}

// GetPlatformsOk returns a tuple with the Platforms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarEvent) GetPlatformsOk() ([]CalendarEventPlatform, bool) {
	if o == nil || IsNil(o.Platforms) {
		return nil, false
	}
	return o.Platforms, true
}

// HasPlatforms returns a boolean if a field has been set.
func (o *CalendarEvent) HasPlatforms() bool {
	if o != nil && !IsNil(o.Platforms) {
		return true
	}

	return false
}

// SetPlatforms gets a reference to the given []CalendarEventPlatform and assigns it to the Platforms field.
func (o *CalendarEvent) SetPlatforms(v []CalendarEventPlatform) {
	o.Platforms = v
}

// GetRoleIds returns the RoleIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CalendarEvent) GetRoleIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RoleIds
}

// GetRoleIdsOk returns a tuple with the RoleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CalendarEvent) GetRoleIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RoleIds) {
		return nil, false
	}
	return o.RoleIds, true
}

// HasRoleIds returns a boolean if a field has been set.
func (o *CalendarEvent) HasRoleIds() bool {
	if o != nil && !IsNil(o.RoleIds) {
		return true
	}

	return false
}

// SetRoleIds gets a reference to the given []string and assigns it to the RoleIds field.
func (o *CalendarEvent) SetRoleIds(v []string) {
	o.RoleIds = v
}

// GetStartsAt returns the StartsAt field value
func (o *CalendarEvent) GetStartsAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value
// and a boolean to check if the value has been set.
func (o *CalendarEvent) GetStartsAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartsAt, true
}

// SetStartsAt sets field value
func (o *CalendarEvent) SetStartsAt(v time.Time) {
	o.StartsAt = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CalendarEvent) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarEvent) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CalendarEvent) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CalendarEvent) SetTags(v []string) {
	o.Tags = v
}

// GetTitle returns the Title field value
func (o *CalendarEvent) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CalendarEvent) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *CalendarEvent) SetTitle(v string) {
	o.Title = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CalendarEvent) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarEvent) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CalendarEvent) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CalendarEvent) SetType(v string) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CalendarEvent) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarEvent) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CalendarEvent) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CalendarEvent) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUserInterest returns the UserInterest field value if set, zero value otherwise.
func (o *CalendarEvent) GetUserInterest() CalendarEventUserInterest {
	if o == nil || IsNil(o.UserInterest) {
		var ret CalendarEventUserInterest
		return ret
	}
	return *o.UserInterest
}

// GetUserInterestOk returns a tuple with the UserInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarEvent) GetUserInterestOk() (*CalendarEventUserInterest, bool) {
	if o == nil || IsNil(o.UserInterest) {
		return nil, false
	}
	return o.UserInterest, true
}

// HasUserInterest returns a boolean if a field has been set.
func (o *CalendarEvent) HasUserInterest() bool {
	if o != nil && !IsNil(o.UserInterest) {
		return true
	}

	return false
}

// SetUserInterest gets a reference to the given CalendarEventUserInterest and assigns it to the UserInterest field.
func (o *CalendarEvent) SetUserInterest(v CalendarEventUserInterest) {
	o.UserInterest = &v
}

// GetUsesInstanceOverflow returns the UsesInstanceOverflow field value if set, zero value otherwise.
func (o *CalendarEvent) GetUsesInstanceOverflow() bool {
	if o == nil || IsNil(o.UsesInstanceOverflow) {
		var ret bool
		return ret
	}
	return *o.UsesInstanceOverflow
}

// GetUsesInstanceOverflowOk returns a tuple with the UsesInstanceOverflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarEvent) GetUsesInstanceOverflowOk() (*bool, bool) {
	if o == nil || IsNil(o.UsesInstanceOverflow) {
		return nil, false
	}
	return o.UsesInstanceOverflow, true
}

// HasUsesInstanceOverflow returns a boolean if a field has been set.
func (o *CalendarEvent) HasUsesInstanceOverflow() bool {
	if o != nil && !IsNil(o.UsesInstanceOverflow) {
		return true
	}

	return false
}

// SetUsesInstanceOverflow gets a reference to the given bool and assigns it to the UsesInstanceOverflow field.
func (o *CalendarEvent) SetUsesInstanceOverflow(v bool) {
	o.UsesInstanceOverflow = &v
}

func (o CalendarEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CalendarEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accessType"] = o.AccessType
	toSerialize["category"] = o.Category
	if !IsNil(o.CloseInstanceAfterEndMinutes) {
		toSerialize["closeInstanceAfterEndMinutes"] = o.CloseInstanceAfterEndMinutes
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.DeletedAt.IsSet() {
		toSerialize["deletedAt"] = o.DeletedAt.Get()
	}
	toSerialize["description"] = o.Description
	if !IsNil(o.DurationInMs) {
		toSerialize["durationInMs"] = o.DurationInMs
	}
	toSerialize["endsAt"] = o.EndsAt
	if !IsNil(o.Featured) {
		toSerialize["featured"] = o.Featured
	}
	if !IsNil(o.GuestEarlyJoinMinutes) {
		toSerialize["guestEarlyJoinMinutes"] = o.GuestEarlyJoinMinutes
	}
	if !IsNil(o.HostEarlyJoinMinutes) {
		toSerialize["hostEarlyJoinMinutes"] = o.HostEarlyJoinMinutes
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.ImageId) {
		toSerialize["imageId"] = o.ImageId
	}
	if o.ImageUrl.IsSet() {
		toSerialize["imageUrl"] = o.ImageUrl.Get()
	}
	if !IsNil(o.InterestedUserCount) {
		toSerialize["interestedUserCount"] = o.InterestedUserCount
	}
	if !IsNil(o.IsDraft) {
		toSerialize["isDraft"] = o.IsDraft
	}
	if !IsNil(o.Languages) {
		toSerialize["languages"] = o.Languages
	}
	if !IsNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}
	if !IsNil(o.Platforms) {
		toSerialize["platforms"] = o.Platforms
	}
	if o.RoleIds != nil {
		toSerialize["roleIds"] = o.RoleIds
	}
	toSerialize["startsAt"] = o.StartsAt
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["title"] = o.Title
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.UserInterest) {
		toSerialize["userInterest"] = o.UserInterest
	}
	if !IsNil(o.UsesInstanceOverflow) {
		toSerialize["usesInstanceOverflow"] = o.UsesInstanceOverflow
	}
	return toSerialize, nil
}

func (o *CalendarEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accessType",
		"category",
		"description",
		"endsAt",
		"id",
		"startsAt",
		"title",
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

	varCalendarEvent := _CalendarEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCalendarEvent)

	if err != nil {
		return err
	}

	*o = CalendarEvent(varCalendarEvent)

	return err
}

type NullableCalendarEvent struct {
	value *CalendarEvent
	isSet bool
}

func (v NullableCalendarEvent) Get() *CalendarEvent {
	return v.value
}

func (v *NullableCalendarEvent) Set(val *CalendarEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableCalendarEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableCalendarEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCalendarEvent(val *CalendarEvent) *NullableCalendarEvent {
	return &NullableCalendarEvent{value: val, isSet: true}
}

func (v NullableCalendarEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCalendarEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
