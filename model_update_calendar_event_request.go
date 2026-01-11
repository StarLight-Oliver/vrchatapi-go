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

// checks if the UpdateCalendarEventRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCalendarEventRequest{}

// UpdateCalendarEventRequest struct for UpdateCalendarEventRequest
type UpdateCalendarEventRequest struct {
	Category *string `json:"category,omitempty"`
	CloseInstanceAfterEndMinutes *int32 `json:"closeInstanceAfterEndMinutes,omitempty"`
	Description *string `json:"description,omitempty"`
	// Time the vent starts at
	EndsAt *time.Time `json:"endsAt,omitempty"`
	Featured *bool `json:"featured,omitempty"`
	GuestEarlyJoinMinutes *int32 `json:"guestEarlyJoinMinutes,omitempty"`
	HostEarlyJoinMinutes *int32 `json:"hostEarlyJoinMinutes,omitempty"`
	ImageId *string `json:"imageId,omitempty"`
	IsDraft *bool `json:"isDraft,omitempty"`
	Languages []string `json:"languages,omitempty"`
	ParentId *string `json:"parentId,omitempty"`
	Platforms []string `json:"platforms,omitempty"`
	RoleIds []string `json:"roleIds,omitempty"`
	// Send notification to group members.
	SendCreationNotification *bool `json:"sendCreationNotification,omitempty"`
	// Time the vent starts at
	StartsAt *time.Time `json:"startsAt,omitempty"`
	Tags []string `json:"tags,omitempty"`
	// Event title
	Title *string `json:"title,omitempty"`
	UsesInstanceOverflow *bool `json:"usesInstanceOverflow,omitempty"`
}

// NewUpdateCalendarEventRequest instantiates a new UpdateCalendarEventRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCalendarEventRequest() *UpdateCalendarEventRequest {
	this := UpdateCalendarEventRequest{}
	var sendCreationNotification bool = false
	this.SendCreationNotification = &sendCreationNotification
	return &this
}

// NewUpdateCalendarEventRequestWithDefaults instantiates a new UpdateCalendarEventRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCalendarEventRequestWithDefaults() *UpdateCalendarEventRequest {
	this := UpdateCalendarEventRequest{}
	var sendCreationNotification bool = false
	this.SendCreationNotification = &sendCreationNotification
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *UpdateCalendarEventRequest) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCalendarEventRequest) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *UpdateCalendarEventRequest) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *UpdateCalendarEventRequest) SetCategory(v string) {
	o.Category = &v
}

// GetCloseInstanceAfterEndMinutes returns the CloseInstanceAfterEndMinutes field value if set, zero value otherwise.
func (o *UpdateCalendarEventRequest) GetCloseInstanceAfterEndMinutes() int32 {
	if o == nil || IsNil(o.CloseInstanceAfterEndMinutes) {
		var ret int32
		return ret
	}
	return *o.CloseInstanceAfterEndMinutes
}

// GetCloseInstanceAfterEndMinutesOk returns a tuple with the CloseInstanceAfterEndMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCalendarEventRequest) GetCloseInstanceAfterEndMinutesOk() (*int32, bool) {
	if o == nil || IsNil(o.CloseInstanceAfterEndMinutes) {
		return nil, false
	}
	return o.CloseInstanceAfterEndMinutes, true
}

// HasCloseInstanceAfterEndMinutes returns a boolean if a field has been set.
func (o *UpdateCalendarEventRequest) HasCloseInstanceAfterEndMinutes() bool {
	if o != nil && !IsNil(o.CloseInstanceAfterEndMinutes) {
		return true
	}

	return false
}

// SetCloseInstanceAfterEndMinutes gets a reference to the given int32 and assigns it to the CloseInstanceAfterEndMinutes field.
func (o *UpdateCalendarEventRequest) SetCloseInstanceAfterEndMinutes(v int32) {
	o.CloseInstanceAfterEndMinutes = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateCalendarEventRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCalendarEventRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateCalendarEventRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateCalendarEventRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEndsAt returns the EndsAt field value if set, zero value otherwise.
func (o *UpdateCalendarEventRequest) GetEndsAt() time.Time {
	if o == nil || IsNil(o.EndsAt) {
		var ret time.Time
		return ret
	}
	return *o.EndsAt
}

// GetEndsAtOk returns a tuple with the EndsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCalendarEventRequest) GetEndsAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndsAt) {
		return nil, false
	}
	return o.EndsAt, true
}

// HasEndsAt returns a boolean if a field has been set.
func (o *UpdateCalendarEventRequest) HasEndsAt() bool {
	if o != nil && !IsNil(o.EndsAt) {
		return true
	}

	return false
}

// SetEndsAt gets a reference to the given time.Time and assigns it to the EndsAt field.
func (o *UpdateCalendarEventRequest) SetEndsAt(v time.Time) {
	o.EndsAt = &v
}

// GetFeatured returns the Featured field value if set, zero value otherwise.
func (o *UpdateCalendarEventRequest) GetFeatured() bool {
	if o == nil || IsNil(o.Featured) {
		var ret bool
		return ret
	}
	return *o.Featured
}

// GetFeaturedOk returns a tuple with the Featured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCalendarEventRequest) GetFeaturedOk() (*bool, bool) {
	if o == nil || IsNil(o.Featured) {
		return nil, false
	}
	return o.Featured, true
}

// HasFeatured returns a boolean if a field has been set.
func (o *UpdateCalendarEventRequest) HasFeatured() bool {
	if o != nil && !IsNil(o.Featured) {
		return true
	}

	return false
}

// SetFeatured gets a reference to the given bool and assigns it to the Featured field.
func (o *UpdateCalendarEventRequest) SetFeatured(v bool) {
	o.Featured = &v
}

// GetGuestEarlyJoinMinutes returns the GuestEarlyJoinMinutes field value if set, zero value otherwise.
func (o *UpdateCalendarEventRequest) GetGuestEarlyJoinMinutes() int32 {
	if o == nil || IsNil(o.GuestEarlyJoinMinutes) {
		var ret int32
		return ret
	}
	return *o.GuestEarlyJoinMinutes
}

// GetGuestEarlyJoinMinutesOk returns a tuple with the GuestEarlyJoinMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCalendarEventRequest) GetGuestEarlyJoinMinutesOk() (*int32, bool) {
	if o == nil || IsNil(o.GuestEarlyJoinMinutes) {
		return nil, false
	}
	return o.GuestEarlyJoinMinutes, true
}

// HasGuestEarlyJoinMinutes returns a boolean if a field has been set.
func (o *UpdateCalendarEventRequest) HasGuestEarlyJoinMinutes() bool {
	if o != nil && !IsNil(o.GuestEarlyJoinMinutes) {
		return true
	}

	return false
}

// SetGuestEarlyJoinMinutes gets a reference to the given int32 and assigns it to the GuestEarlyJoinMinutes field.
func (o *UpdateCalendarEventRequest) SetGuestEarlyJoinMinutes(v int32) {
	o.GuestEarlyJoinMinutes = &v
}

// GetHostEarlyJoinMinutes returns the HostEarlyJoinMinutes field value if set, zero value otherwise.
func (o *UpdateCalendarEventRequest) GetHostEarlyJoinMinutes() int32 {
	if o == nil || IsNil(o.HostEarlyJoinMinutes) {
		var ret int32
		return ret
	}
	return *o.HostEarlyJoinMinutes
}

// GetHostEarlyJoinMinutesOk returns a tuple with the HostEarlyJoinMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCalendarEventRequest) GetHostEarlyJoinMinutesOk() (*int32, bool) {
	if o == nil || IsNil(o.HostEarlyJoinMinutes) {
		return nil, false
	}
	return o.HostEarlyJoinMinutes, true
}

// HasHostEarlyJoinMinutes returns a boolean if a field has been set.
func (o *UpdateCalendarEventRequest) HasHostEarlyJoinMinutes() bool {
	if o != nil && !IsNil(o.HostEarlyJoinMinutes) {
		return true
	}

	return false
}

// SetHostEarlyJoinMinutes gets a reference to the given int32 and assigns it to the HostEarlyJoinMinutes field.
func (o *UpdateCalendarEventRequest) SetHostEarlyJoinMinutes(v int32) {
	o.HostEarlyJoinMinutes = &v
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *UpdateCalendarEventRequest) GetImageId() string {
	if o == nil || IsNil(o.ImageId) {
		var ret string
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCalendarEventRequest) GetImageIdOk() (*string, bool) {
	if o == nil || IsNil(o.ImageId) {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *UpdateCalendarEventRequest) HasImageId() bool {
	if o != nil && !IsNil(o.ImageId) {
		return true
	}

	return false
}

// SetImageId gets a reference to the given string and assigns it to the ImageId field.
func (o *UpdateCalendarEventRequest) SetImageId(v string) {
	o.ImageId = &v
}

// GetIsDraft returns the IsDraft field value if set, zero value otherwise.
func (o *UpdateCalendarEventRequest) GetIsDraft() bool {
	if o == nil || IsNil(o.IsDraft) {
		var ret bool
		return ret
	}
	return *o.IsDraft
}

// GetIsDraftOk returns a tuple with the IsDraft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCalendarEventRequest) GetIsDraftOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDraft) {
		return nil, false
	}
	return o.IsDraft, true
}

// HasIsDraft returns a boolean if a field has been set.
func (o *UpdateCalendarEventRequest) HasIsDraft() bool {
	if o != nil && !IsNil(o.IsDraft) {
		return true
	}

	return false
}

// SetIsDraft gets a reference to the given bool and assigns it to the IsDraft field.
func (o *UpdateCalendarEventRequest) SetIsDraft(v bool) {
	o.IsDraft = &v
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *UpdateCalendarEventRequest) GetLanguages() []string {
	if o == nil || IsNil(o.Languages) {
		var ret []string
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCalendarEventRequest) GetLanguagesOk() ([]string, bool) {
	if o == nil || IsNil(o.Languages) {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *UpdateCalendarEventRequest) HasLanguages() bool {
	if o != nil && !IsNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []string and assigns it to the Languages field.
func (o *UpdateCalendarEventRequest) SetLanguages(v []string) {
	o.Languages = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *UpdateCalendarEventRequest) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCalendarEventRequest) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *UpdateCalendarEventRequest) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *UpdateCalendarEventRequest) SetParentId(v string) {
	o.ParentId = &v
}

// GetPlatforms returns the Platforms field value if set, zero value otherwise.
func (o *UpdateCalendarEventRequest) GetPlatforms() []string {
	if o == nil || IsNil(o.Platforms) {
		var ret []string
		return ret
	}
	return o.Platforms
}

// GetPlatformsOk returns a tuple with the Platforms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCalendarEventRequest) GetPlatformsOk() ([]string, bool) {
	if o == nil || IsNil(o.Platforms) {
		return nil, false
	}
	return o.Platforms, true
}

// HasPlatforms returns a boolean if a field has been set.
func (o *UpdateCalendarEventRequest) HasPlatforms() bool {
	if o != nil && !IsNil(o.Platforms) {
		return true
	}

	return false
}

// SetPlatforms gets a reference to the given []string and assigns it to the Platforms field.
func (o *UpdateCalendarEventRequest) SetPlatforms(v []string) {
	o.Platforms = v
}

// GetRoleIds returns the RoleIds field value if set, zero value otherwise.
func (o *UpdateCalendarEventRequest) GetRoleIds() []string {
	if o == nil || IsNil(o.RoleIds) {
		var ret []string
		return ret
	}
	return o.RoleIds
}

// GetRoleIdsOk returns a tuple with the RoleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCalendarEventRequest) GetRoleIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RoleIds) {
		return nil, false
	}
	return o.RoleIds, true
}

// HasRoleIds returns a boolean if a field has been set.
func (o *UpdateCalendarEventRequest) HasRoleIds() bool {
	if o != nil && !IsNil(o.RoleIds) {
		return true
	}

	return false
}

// SetRoleIds gets a reference to the given []string and assigns it to the RoleIds field.
func (o *UpdateCalendarEventRequest) SetRoleIds(v []string) {
	o.RoleIds = v
}

// GetSendCreationNotification returns the SendCreationNotification field value if set, zero value otherwise.
func (o *UpdateCalendarEventRequest) GetSendCreationNotification() bool {
	if o == nil || IsNil(o.SendCreationNotification) {
		var ret bool
		return ret
	}
	return *o.SendCreationNotification
}

// GetSendCreationNotificationOk returns a tuple with the SendCreationNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCalendarEventRequest) GetSendCreationNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.SendCreationNotification) {
		return nil, false
	}
	return o.SendCreationNotification, true
}

// HasSendCreationNotification returns a boolean if a field has been set.
func (o *UpdateCalendarEventRequest) HasSendCreationNotification() bool {
	if o != nil && !IsNil(o.SendCreationNotification) {
		return true
	}

	return false
}

// SetSendCreationNotification gets a reference to the given bool and assigns it to the SendCreationNotification field.
func (o *UpdateCalendarEventRequest) SetSendCreationNotification(v bool) {
	o.SendCreationNotification = &v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise.
func (o *UpdateCalendarEventRequest) GetStartsAt() time.Time {
	if o == nil || IsNil(o.StartsAt) {
		var ret time.Time
		return ret
	}
	return *o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCalendarEventRequest) GetStartsAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartsAt) {
		return nil, false
	}
	return o.StartsAt, true
}

// HasStartsAt returns a boolean if a field has been set.
func (o *UpdateCalendarEventRequest) HasStartsAt() bool {
	if o != nil && !IsNil(o.StartsAt) {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given time.Time and assigns it to the StartsAt field.
func (o *UpdateCalendarEventRequest) SetStartsAt(v time.Time) {
	o.StartsAt = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateCalendarEventRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCalendarEventRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateCalendarEventRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UpdateCalendarEventRequest) SetTags(v []string) {
	o.Tags = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *UpdateCalendarEventRequest) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCalendarEventRequest) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UpdateCalendarEventRequest) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *UpdateCalendarEventRequest) SetTitle(v string) {
	o.Title = &v
}

// GetUsesInstanceOverflow returns the UsesInstanceOverflow field value if set, zero value otherwise.
func (o *UpdateCalendarEventRequest) GetUsesInstanceOverflow() bool {
	if o == nil || IsNil(o.UsesInstanceOverflow) {
		var ret bool
		return ret
	}
	return *o.UsesInstanceOverflow
}

// GetUsesInstanceOverflowOk returns a tuple with the UsesInstanceOverflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCalendarEventRequest) GetUsesInstanceOverflowOk() (*bool, bool) {
	if o == nil || IsNil(o.UsesInstanceOverflow) {
		return nil, false
	}
	return o.UsesInstanceOverflow, true
}

// HasUsesInstanceOverflow returns a boolean if a field has been set.
func (o *UpdateCalendarEventRequest) HasUsesInstanceOverflow() bool {
	if o != nil && !IsNil(o.UsesInstanceOverflow) {
		return true
	}

	return false
}

// SetUsesInstanceOverflow gets a reference to the given bool and assigns it to the UsesInstanceOverflow field.
func (o *UpdateCalendarEventRequest) SetUsesInstanceOverflow(v bool) {
	o.UsesInstanceOverflow = &v
}

func (o UpdateCalendarEventRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCalendarEventRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.CloseInstanceAfterEndMinutes) {
		toSerialize["closeInstanceAfterEndMinutes"] = o.CloseInstanceAfterEndMinutes
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.EndsAt) {
		toSerialize["endsAt"] = o.EndsAt
	}
	if !IsNil(o.Featured) {
		toSerialize["featured"] = o.Featured
	}
	if !IsNil(o.GuestEarlyJoinMinutes) {
		toSerialize["guestEarlyJoinMinutes"] = o.GuestEarlyJoinMinutes
	}
	if !IsNil(o.HostEarlyJoinMinutes) {
		toSerialize["hostEarlyJoinMinutes"] = o.HostEarlyJoinMinutes
	}
	if !IsNil(o.ImageId) {
		toSerialize["imageId"] = o.ImageId
	}
	if !IsNil(o.IsDraft) {
		toSerialize["isDraft"] = o.IsDraft
	}
	if !IsNil(o.Languages) {
		toSerialize["languages"] = o.Languages
	}
	if !IsNil(o.ParentId) {
		toSerialize["parentId"] = o.ParentId
	}
	if !IsNil(o.Platforms) {
		toSerialize["platforms"] = o.Platforms
	}
	if !IsNil(o.RoleIds) {
		toSerialize["roleIds"] = o.RoleIds
	}
	if !IsNil(o.SendCreationNotification) {
		toSerialize["sendCreationNotification"] = o.SendCreationNotification
	}
	if !IsNil(o.StartsAt) {
		toSerialize["startsAt"] = o.StartsAt
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.UsesInstanceOverflow) {
		toSerialize["usesInstanceOverflow"] = o.UsesInstanceOverflow
	}
	return toSerialize, nil
}

type NullableUpdateCalendarEventRequest struct {
	value *UpdateCalendarEventRequest
	isSet bool
}

func (v NullableUpdateCalendarEventRequest) Get() *UpdateCalendarEventRequest {
	return v.value
}

func (v *NullableUpdateCalendarEventRequest) Set(val *UpdateCalendarEventRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCalendarEventRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCalendarEventRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCalendarEventRequest(val *UpdateCalendarEventRequest) *NullableUpdateCalendarEventRequest {
	return &NullableUpdateCalendarEventRequest{value: val, isSet: true}
}

func (v NullableUpdateCalendarEventRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCalendarEventRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


