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

// checks if the LimitedGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LimitedGroup{}

// LimitedGroup struct for LimitedGroup
type LimitedGroup struct {
	BannerId      NullableString `json:"bannerId,omitempty"`
	BannerUrl     NullableString `json:"bannerUrl,omitempty"`
	CreatedAt     *time.Time     `json:"createdAt,omitempty"`
	Description   *string        `json:"description,omitempty"`
	Discriminator *string        `json:"discriminator,omitempty"`
	//
	Galleries        []GroupGallery     `json:"galleries,omitempty"`
	IconId           NullableString     `json:"iconId,omitempty"`
	IconUrl          NullableString     `json:"iconUrl,omitempty"`
	Id               *string            `json:"id,omitempty"`
	IsSearchable     *bool              `json:"isSearchable,omitempty"`
	MemberCount      *int32             `json:"memberCount,omitempty"`
	MembershipStatus *GroupMemberStatus `json:"membershipStatus,omitempty"`
	Name             *string            `json:"name,omitempty"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	OwnerId   *string        `json:"ownerId,omitempty"`
	Rules     NullableString `json:"rules,omitempty"`
	ShortCode *string        `json:"shortCode,omitempty"`
	//
	Tags []string `json:"tags,omitempty"`
}

// NewLimitedGroup instantiates a new LimitedGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLimitedGroup() *LimitedGroup {
	this := LimitedGroup{}
	var membershipStatus GroupMemberStatus = GroupMemberStatus_INACTIVE
	this.MembershipStatus = &membershipStatus
	return &this
}

// NewLimitedGroupWithDefaults instantiates a new LimitedGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLimitedGroupWithDefaults() *LimitedGroup {
	this := LimitedGroup{}
	var membershipStatus GroupMemberStatus = GroupMemberStatus_INACTIVE
	this.MembershipStatus = &membershipStatus
	return &this
}

// GetBannerId returns the BannerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LimitedGroup) GetBannerId() string {
	if o == nil || IsNil(o.BannerId.Get()) {
		var ret string
		return ret
	}
	return *o.BannerId.Get()
}

// GetBannerIdOk returns a tuple with the BannerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LimitedGroup) GetBannerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BannerId.Get(), o.BannerId.IsSet()
}

// HasBannerId returns a boolean if a field has been set.
func (o *LimitedGroup) HasBannerId() bool {
	if o != nil && o.BannerId.IsSet() {
		return true
	}

	return false
}

// SetBannerId gets a reference to the given NullableString and assigns it to the BannerId field.
func (o *LimitedGroup) SetBannerId(v string) {
	o.BannerId.Set(&v)
}

// SetBannerIdNil sets the value for BannerId to be an explicit nil
func (o *LimitedGroup) SetBannerIdNil() {
	o.BannerId.Set(nil)
}

// UnsetBannerId ensures that no value is present for BannerId, not even an explicit nil
func (o *LimitedGroup) UnsetBannerId() {
	o.BannerId.Unset()
}

// GetBannerUrl returns the BannerUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LimitedGroup) GetBannerUrl() string {
	if o == nil || IsNil(o.BannerUrl.Get()) {
		var ret string
		return ret
	}
	return *o.BannerUrl.Get()
}

// GetBannerUrlOk returns a tuple with the BannerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LimitedGroup) GetBannerUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BannerUrl.Get(), o.BannerUrl.IsSet()
}

// HasBannerUrl returns a boolean if a field has been set.
func (o *LimitedGroup) HasBannerUrl() bool {
	if o != nil && o.BannerUrl.IsSet() {
		return true
	}

	return false
}

// SetBannerUrl gets a reference to the given NullableString and assigns it to the BannerUrl field.
func (o *LimitedGroup) SetBannerUrl(v string) {
	o.BannerUrl.Set(&v)
}

// SetBannerUrlNil sets the value for BannerUrl to be an explicit nil
func (o *LimitedGroup) SetBannerUrlNil() {
	o.BannerUrl.Set(nil)
}

// UnsetBannerUrl ensures that no value is present for BannerUrl, not even an explicit nil
func (o *LimitedGroup) UnsetBannerUrl() {
	o.BannerUrl.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *LimitedGroup) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedGroup) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LimitedGroup) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *LimitedGroup) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LimitedGroup) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LimitedGroup) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LimitedGroup) SetDescription(v string) {
	o.Description = &v
}

// GetDiscriminator returns the Discriminator field value if set, zero value otherwise.
func (o *LimitedGroup) GetDiscriminator() string {
	if o == nil || IsNil(o.Discriminator) {
		var ret string
		return ret
	}
	return *o.Discriminator
}

// GetDiscriminatorOk returns a tuple with the Discriminator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedGroup) GetDiscriminatorOk() (*string, bool) {
	if o == nil || IsNil(o.Discriminator) {
		return nil, false
	}
	return o.Discriminator, true
}

// HasDiscriminator returns a boolean if a field has been set.
func (o *LimitedGroup) HasDiscriminator() bool {
	if o != nil && !IsNil(o.Discriminator) {
		return true
	}

	return false
}

// SetDiscriminator gets a reference to the given string and assigns it to the Discriminator field.
func (o *LimitedGroup) SetDiscriminator(v string) {
	o.Discriminator = &v
}

// GetGalleries returns the Galleries field value if set, zero value otherwise.
func (o *LimitedGroup) GetGalleries() []GroupGallery {
	if o == nil || IsNil(o.Galleries) {
		var ret []GroupGallery
		return ret
	}
	return o.Galleries
}

// GetGalleriesOk returns a tuple with the Galleries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedGroup) GetGalleriesOk() ([]GroupGallery, bool) {
	if o == nil || IsNil(o.Galleries) {
		return nil, false
	}
	return o.Galleries, true
}

// HasGalleries returns a boolean if a field has been set.
func (o *LimitedGroup) HasGalleries() bool {
	if o != nil && !IsNil(o.Galleries) {
		return true
	}

	return false
}

// SetGalleries gets a reference to the given []GroupGallery and assigns it to the Galleries field.
func (o *LimitedGroup) SetGalleries(v []GroupGallery) {
	o.Galleries = v
}

// GetIconId returns the IconId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LimitedGroup) GetIconId() string {
	if o == nil || IsNil(o.IconId.Get()) {
		var ret string
		return ret
	}
	return *o.IconId.Get()
}

// GetIconIdOk returns a tuple with the IconId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LimitedGroup) GetIconIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IconId.Get(), o.IconId.IsSet()
}

// HasIconId returns a boolean if a field has been set.
func (o *LimitedGroup) HasIconId() bool {
	if o != nil && o.IconId.IsSet() {
		return true
	}

	return false
}

// SetIconId gets a reference to the given NullableString and assigns it to the IconId field.
func (o *LimitedGroup) SetIconId(v string) {
	o.IconId.Set(&v)
}

// SetIconIdNil sets the value for IconId to be an explicit nil
func (o *LimitedGroup) SetIconIdNil() {
	o.IconId.Set(nil)
}

// UnsetIconId ensures that no value is present for IconId, not even an explicit nil
func (o *LimitedGroup) UnsetIconId() {
	o.IconId.Unset()
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LimitedGroup) GetIconUrl() string {
	if o == nil || IsNil(o.IconUrl.Get()) {
		var ret string
		return ret
	}
	return *o.IconUrl.Get()
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LimitedGroup) GetIconUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IconUrl.Get(), o.IconUrl.IsSet()
}

// HasIconUrl returns a boolean if a field has been set.
func (o *LimitedGroup) HasIconUrl() bool {
	if o != nil && o.IconUrl.IsSet() {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given NullableString and assigns it to the IconUrl field.
func (o *LimitedGroup) SetIconUrl(v string) {
	o.IconUrl.Set(&v)
}

// SetIconUrlNil sets the value for IconUrl to be an explicit nil
func (o *LimitedGroup) SetIconUrlNil() {
	o.IconUrl.Set(nil)
}

// UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
func (o *LimitedGroup) UnsetIconUrl() {
	o.IconUrl.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LimitedGroup) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedGroup) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LimitedGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LimitedGroup) SetId(v string) {
	o.Id = &v
}

// GetIsSearchable returns the IsSearchable field value if set, zero value otherwise.
func (o *LimitedGroup) GetIsSearchable() bool {
	if o == nil || IsNil(o.IsSearchable) {
		var ret bool
		return ret
	}
	return *o.IsSearchable
}

// GetIsSearchableOk returns a tuple with the IsSearchable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedGroup) GetIsSearchableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSearchable) {
		return nil, false
	}
	return o.IsSearchable, true
}

// HasIsSearchable returns a boolean if a field has been set.
func (o *LimitedGroup) HasIsSearchable() bool {
	if o != nil && !IsNil(o.IsSearchable) {
		return true
	}

	return false
}

// SetIsSearchable gets a reference to the given bool and assigns it to the IsSearchable field.
func (o *LimitedGroup) SetIsSearchable(v bool) {
	o.IsSearchable = &v
}

// GetMemberCount returns the MemberCount field value if set, zero value otherwise.
func (o *LimitedGroup) GetMemberCount() int32 {
	if o == nil || IsNil(o.MemberCount) {
		var ret int32
		return ret
	}
	return *o.MemberCount
}

// GetMemberCountOk returns a tuple with the MemberCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedGroup) GetMemberCountOk() (*int32, bool) {
	if o == nil || IsNil(o.MemberCount) {
		return nil, false
	}
	return o.MemberCount, true
}

// HasMemberCount returns a boolean if a field has been set.
func (o *LimitedGroup) HasMemberCount() bool {
	if o != nil && !IsNil(o.MemberCount) {
		return true
	}

	return false
}

// SetMemberCount gets a reference to the given int32 and assigns it to the MemberCount field.
func (o *LimitedGroup) SetMemberCount(v int32) {
	o.MemberCount = &v
}

// GetMembershipStatus returns the MembershipStatus field value if set, zero value otherwise.
func (o *LimitedGroup) GetMembershipStatus() GroupMemberStatus {
	if o == nil || IsNil(o.MembershipStatus) {
		var ret GroupMemberStatus
		return ret
	}
	return *o.MembershipStatus
}

// GetMembershipStatusOk returns a tuple with the MembershipStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedGroup) GetMembershipStatusOk() (*GroupMemberStatus, bool) {
	if o == nil || IsNil(o.MembershipStatus) {
		return nil, false
	}
	return o.MembershipStatus, true
}

// HasMembershipStatus returns a boolean if a field has been set.
func (o *LimitedGroup) HasMembershipStatus() bool {
	if o != nil && !IsNil(o.MembershipStatus) {
		return true
	}

	return false
}

// SetMembershipStatus gets a reference to the given GroupMemberStatus and assigns it to the MembershipStatus field.
func (o *LimitedGroup) SetMembershipStatus(v GroupMemberStatus) {
	o.MembershipStatus = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LimitedGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LimitedGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LimitedGroup) SetName(v string) {
	o.Name = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *LimitedGroup) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedGroup) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *LimitedGroup) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *LimitedGroup) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetRules returns the Rules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LimitedGroup) GetRules() string {
	if o == nil || IsNil(o.Rules.Get()) {
		var ret string
		return ret
	}
	return *o.Rules.Get()
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LimitedGroup) GetRulesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rules.Get(), o.Rules.IsSet()
}

// HasRules returns a boolean if a field has been set.
func (o *LimitedGroup) HasRules() bool {
	if o != nil && o.Rules.IsSet() {
		return true
	}

	return false
}

// SetRules gets a reference to the given NullableString and assigns it to the Rules field.
func (o *LimitedGroup) SetRules(v string) {
	o.Rules.Set(&v)
}

// SetRulesNil sets the value for Rules to be an explicit nil
func (o *LimitedGroup) SetRulesNil() {
	o.Rules.Set(nil)
}

// UnsetRules ensures that no value is present for Rules, not even an explicit nil
func (o *LimitedGroup) UnsetRules() {
	o.Rules.Unset()
}

// GetShortCode returns the ShortCode field value if set, zero value otherwise.
func (o *LimitedGroup) GetShortCode() string {
	if o == nil || IsNil(o.ShortCode) {
		var ret string
		return ret
	}
	return *o.ShortCode
}

// GetShortCodeOk returns a tuple with the ShortCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedGroup) GetShortCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ShortCode) {
		return nil, false
	}
	return o.ShortCode, true
}

// HasShortCode returns a boolean if a field has been set.
func (o *LimitedGroup) HasShortCode() bool {
	if o != nil && !IsNil(o.ShortCode) {
		return true
	}

	return false
}

// SetShortCode gets a reference to the given string and assigns it to the ShortCode field.
func (o *LimitedGroup) SetShortCode(v string) {
	o.ShortCode = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *LimitedGroup) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LimitedGroup) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *LimitedGroup) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *LimitedGroup) SetTags(v []string) {
	o.Tags = v
}

func (o LimitedGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LimitedGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BannerId.IsSet() {
		toSerialize["bannerId"] = o.BannerId.Get()
	}
	if o.BannerUrl.IsSet() {
		toSerialize["bannerUrl"] = o.BannerUrl.Get()
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Discriminator) {
		toSerialize["discriminator"] = o.Discriminator
	}
	if !IsNil(o.Galleries) {
		toSerialize["galleries"] = o.Galleries
	}
	if o.IconId.IsSet() {
		toSerialize["iconId"] = o.IconId.Get()
	}
	if o.IconUrl.IsSet() {
		toSerialize["iconUrl"] = o.IconUrl.Get()
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsSearchable) {
		toSerialize["isSearchable"] = o.IsSearchable
	}
	if !IsNil(o.MemberCount) {
		toSerialize["memberCount"] = o.MemberCount
	}
	if !IsNil(o.MembershipStatus) {
		toSerialize["membershipStatus"] = o.MembershipStatus
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}
	if o.Rules.IsSet() {
		toSerialize["rules"] = o.Rules.Get()
	}
	if !IsNil(o.ShortCode) {
		toSerialize["shortCode"] = o.ShortCode
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableLimitedGroup struct {
	value *LimitedGroup
	isSet bool
}

func (v NullableLimitedGroup) Get() *LimitedGroup {
	return v.value
}

func (v *NullableLimitedGroup) Set(val *LimitedGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableLimitedGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableLimitedGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLimitedGroup(val *LimitedGroup) *NullableLimitedGroup {
	return &NullableLimitedGroup{value: val, isSet: true}
}

func (v NullableLimitedGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLimitedGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
