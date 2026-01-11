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

// checks if the FavoritedWorld type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FavoritedWorld{}

// FavoritedWorld
type FavoritedWorld struct {
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	AuthorId               *string                  `json:"authorId,omitempty"`
	AuthorName             string                   `json:"authorName"`
	Capacity               int32                    `json:"capacity"`
	CreatedAt              *time.Time               `json:"created_at,omitempty"`
	DefaultContentSettings *InstanceContentSettings `json:"defaultContentSettings,omitempty"`
	Description            *string                  `json:"description,omitempty"`
	FavoriteGroup          string                   `json:"favoriteGroup"`
	FavoriteId             string                   `json:"favoriteId"`
	Favorites              *int32                   `json:"favorites,omitempty"`
	Featured               *bool                    `json:"featured,omitempty"`
	Heat                   *int32                   `json:"heat,omitempty"`
	// WorldID be \"offline\" on User profiles if you are not friends with that user.
	Id                  string         `json:"id"`
	ImageUrl            string         `json:"imageUrl"`
	LabsPublicationDate *string        `json:"labsPublicationDate,omitempty"`
	Name                string         `json:"name"`
	Occupants           int32          `json:"occupants"`
	Organization        *string        `json:"organization,omitempty"`
	Popularity          *int32         `json:"popularity,omitempty"`
	PreviewYoutubeId    NullableString `json:"previewYoutubeId,omitempty"`
	PublicationDate     *string        `json:"publicationDate,omitempty"`
	RecommendedCapacity *int32         `json:"recommendedCapacity,omitempty"`
	ReleaseStatus       ReleaseStatus  `json:"releaseStatus"`
	//
	Tags              []string `json:"tags,omitempty"`
	ThumbnailImageUrl string   `json:"thumbnailImageUrl"`
	UdonProducts      []string `json:"udonProducts,omitempty"`
	//
	UnityPackages []UnityPackage `json:"unityPackages,omitempty"`
	UpdatedAt     *time.Time     `json:"updated_at,omitempty"`
	UrlList       []string       `json:"urlList,omitempty"`
	Version       *int32         `json:"version,omitempty"`
	Visits        *int32         `json:"visits,omitempty"`
}

type _FavoritedWorld FavoritedWorld

// NewFavoritedWorld instantiates a new FavoritedWorld object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFavoritedWorld(authorName string, capacity int32, favoriteGroup string, favoriteId string, id string, imageUrl string, name string, occupants int32, releaseStatus ReleaseStatus, thumbnailImageUrl string) *FavoritedWorld {
	this := FavoritedWorld{}
	this.AuthorName = authorName
	this.Capacity = capacity
	this.FavoriteGroup = favoriteGroup
	this.FavoriteId = favoriteId
	var favorites int32 = 0
	this.Favorites = &favorites
	var featured bool = false
	this.Featured = &featured
	var heat int32 = 0
	this.Heat = &heat
	this.Id = id
	this.ImageUrl = imageUrl
	this.Name = name
	this.Occupants = occupants
	var organization string = "vrchat"
	this.Organization = &organization
	var popularity int32 = 0
	this.Popularity = &popularity
	this.ReleaseStatus = releaseStatus
	this.ThumbnailImageUrl = thumbnailImageUrl
	var visits int32 = 0
	this.Visits = &visits
	return &this
}

// NewFavoritedWorldWithDefaults instantiates a new FavoritedWorld object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFavoritedWorldWithDefaults() *FavoritedWorld {
	this := FavoritedWorld{}
	var favorites int32 = 0
	this.Favorites = &favorites
	var featured bool = false
	this.Featured = &featured
	var heat int32 = 0
	this.Heat = &heat
	var occupants int32 = 0
	this.Occupants = occupants
	var organization string = "vrchat"
	this.Organization = &organization
	var popularity int32 = 0
	this.Popularity = &popularity
	var releaseStatus ReleaseStatus = ReleaseStatus_PUBLIC
	this.ReleaseStatus = releaseStatus
	var visits int32 = 0
	this.Visits = &visits
	return &this
}

// GetAuthorId returns the AuthorId field value if set, zero value otherwise.
func (o *FavoritedWorld) GetAuthorId() string {
	if o == nil || IsNil(o.AuthorId) {
		var ret string
		return ret
	}
	return *o.AuthorId
}

// GetAuthorIdOk returns a tuple with the AuthorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FavoritedWorld) GetAuthorIdOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorId) {
		return nil, false
	}
	return o.AuthorId, true
}

// HasAuthorId returns a boolean if a field has been set.
func (o *FavoritedWorld) HasAuthorId() bool {
	if o != nil && !IsNil(o.AuthorId) {
		return true
	}

	return false
}

// SetAuthorId gets a reference to the given string and assigns it to the AuthorId field.
func (o *FavoritedWorld) SetAuthorId(v string) {
	o.AuthorId = &v
}

// GetAuthorName returns the AuthorName field value
func (o *FavoritedWorld) GetAuthorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorName
}

// GetAuthorNameOk returns a tuple with the AuthorName field value
// and a boolean to check if the value has been set.
func (o *FavoritedWorld) GetAuthorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorName, true
}

// SetAuthorName sets field value
func (o *FavoritedWorld) SetAuthorName(v string) {
	o.AuthorName = v
}

// GetCapacity returns the Capacity field value
func (o *FavoritedWorld) GetCapacity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value
// and a boolean to check if the value has been set.
func (o *FavoritedWorld) GetCapacityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capacity, true
}

// SetCapacity sets field value
func (o *FavoritedWorld) SetCapacity(v int32) {
	o.Capacity = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FavoritedWorld) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FavoritedWorld) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FavoritedWorld) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *FavoritedWorld) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDefaultContentSettings returns the DefaultContentSettings field value if set, zero value otherwise.
func (o *FavoritedWorld) GetDefaultContentSettings() InstanceContentSettings {
	if o == nil || IsNil(o.DefaultContentSettings) {
		var ret InstanceContentSettings
		return ret
	}
	return *o.DefaultContentSettings
}

// GetDefaultContentSettingsOk returns a tuple with the DefaultContentSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FavoritedWorld) GetDefaultContentSettingsOk() (*InstanceContentSettings, bool) {
	if o == nil || IsNil(o.DefaultContentSettings) {
		return nil, false
	}
	return o.DefaultContentSettings, true
}

// HasDefaultContentSettings returns a boolean if a field has been set.
func (o *FavoritedWorld) HasDefaultContentSettings() bool {
	if o != nil && !IsNil(o.DefaultContentSettings) {
		return true
	}

	return false
}

// SetDefaultContentSettings gets a reference to the given InstanceContentSettings and assigns it to the DefaultContentSettings field.
func (o *FavoritedWorld) SetDefaultContentSettings(v InstanceContentSettings) {
	o.DefaultContentSettings = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FavoritedWorld) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FavoritedWorld) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FavoritedWorld) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FavoritedWorld) SetDescription(v string) {
	o.Description = &v
}

// GetFavoriteGroup returns the FavoriteGroup field value
func (o *FavoritedWorld) GetFavoriteGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FavoriteGroup
}

// GetFavoriteGroupOk returns a tuple with the FavoriteGroup field value
// and a boolean to check if the value has been set.
func (o *FavoritedWorld) GetFavoriteGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FavoriteGroup, true
}

// SetFavoriteGroup sets field value
func (o *FavoritedWorld) SetFavoriteGroup(v string) {
	o.FavoriteGroup = v
}

// GetFavoriteId returns the FavoriteId field value
func (o *FavoritedWorld) GetFavoriteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FavoriteId
}

// GetFavoriteIdOk returns a tuple with the FavoriteId field value
// and a boolean to check if the value has been set.
func (o *FavoritedWorld) GetFavoriteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FavoriteId, true
}

// SetFavoriteId sets field value
func (o *FavoritedWorld) SetFavoriteId(v string) {
	o.FavoriteId = v
}

// GetFavorites returns the Favorites field value if set, zero value otherwise.
func (o *FavoritedWorld) GetFavorites() int32 {
	if o == nil || IsNil(o.Favorites) {
		var ret int32
		return ret
	}
	return *o.Favorites
}

// GetFavoritesOk returns a tuple with the Favorites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FavoritedWorld) GetFavoritesOk() (*int32, bool) {
	if o == nil || IsNil(o.Favorites) {
		return nil, false
	}
	return o.Favorites, true
}

// HasFavorites returns a boolean if a field has been set.
func (o *FavoritedWorld) HasFavorites() bool {
	if o != nil && !IsNil(o.Favorites) {
		return true
	}

	return false
}

// SetFavorites gets a reference to the given int32 and assigns it to the Favorites field.
func (o *FavoritedWorld) SetFavorites(v int32) {
	o.Favorites = &v
}

// GetFeatured returns the Featured field value if set, zero value otherwise.
func (o *FavoritedWorld) GetFeatured() bool {
	if o == nil || IsNil(o.Featured) {
		var ret bool
		return ret
	}
	return *o.Featured
}

// GetFeaturedOk returns a tuple with the Featured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FavoritedWorld) GetFeaturedOk() (*bool, bool) {
	if o == nil || IsNil(o.Featured) {
		return nil, false
	}
	return o.Featured, true
}

// HasFeatured returns a boolean if a field has been set.
func (o *FavoritedWorld) HasFeatured() bool {
	if o != nil && !IsNil(o.Featured) {
		return true
	}

	return false
}

// SetFeatured gets a reference to the given bool and assigns it to the Featured field.
func (o *FavoritedWorld) SetFeatured(v bool) {
	o.Featured = &v
}

// GetHeat returns the Heat field value if set, zero value otherwise.
func (o *FavoritedWorld) GetHeat() int32 {
	if o == nil || IsNil(o.Heat) {
		var ret int32
		return ret
	}
	return *o.Heat
}

// GetHeatOk returns a tuple with the Heat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FavoritedWorld) GetHeatOk() (*int32, bool) {
	if o == nil || IsNil(o.Heat) {
		return nil, false
	}
	return o.Heat, true
}

// HasHeat returns a boolean if a field has been set.
func (o *FavoritedWorld) HasHeat() bool {
	if o != nil && !IsNil(o.Heat) {
		return true
	}

	return false
}

// SetHeat gets a reference to the given int32 and assigns it to the Heat field.
func (o *FavoritedWorld) SetHeat(v int32) {
	o.Heat = &v
}

// GetId returns the Id field value
func (o *FavoritedWorld) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FavoritedWorld) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FavoritedWorld) SetId(v string) {
	o.Id = v
}

// GetImageUrl returns the ImageUrl field value
func (o *FavoritedWorld) GetImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
func (o *FavoritedWorld) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageUrl, true
}

// SetImageUrl sets field value
func (o *FavoritedWorld) SetImageUrl(v string) {
	o.ImageUrl = v
}

// GetLabsPublicationDate returns the LabsPublicationDate field value if set, zero value otherwise.
func (o *FavoritedWorld) GetLabsPublicationDate() string {
	if o == nil || IsNil(o.LabsPublicationDate) {
		var ret string
		return ret
	}
	return *o.LabsPublicationDate
}

// GetLabsPublicationDateOk returns a tuple with the LabsPublicationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FavoritedWorld) GetLabsPublicationDateOk() (*string, bool) {
	if o == nil || IsNil(o.LabsPublicationDate) {
		return nil, false
	}
	return o.LabsPublicationDate, true
}

// HasLabsPublicationDate returns a boolean if a field has been set.
func (o *FavoritedWorld) HasLabsPublicationDate() bool {
	if o != nil && !IsNil(o.LabsPublicationDate) {
		return true
	}

	return false
}

// SetLabsPublicationDate gets a reference to the given string and assigns it to the LabsPublicationDate field.
func (o *FavoritedWorld) SetLabsPublicationDate(v string) {
	o.LabsPublicationDate = &v
}

// GetName returns the Name field value
func (o *FavoritedWorld) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FavoritedWorld) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FavoritedWorld) SetName(v string) {
	o.Name = v
}

// GetOccupants returns the Occupants field value
func (o *FavoritedWorld) GetOccupants() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Occupants
}

// GetOccupantsOk returns a tuple with the Occupants field value
// and a boolean to check if the value has been set.
func (o *FavoritedWorld) GetOccupantsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Occupants, true
}

// SetOccupants sets field value
func (o *FavoritedWorld) SetOccupants(v int32) {
	o.Occupants = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *FavoritedWorld) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FavoritedWorld) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *FavoritedWorld) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *FavoritedWorld) SetOrganization(v string) {
	o.Organization = &v
}

// GetPopularity returns the Popularity field value if set, zero value otherwise.
func (o *FavoritedWorld) GetPopularity() int32 {
	if o == nil || IsNil(o.Popularity) {
		var ret int32
		return ret
	}
	return *o.Popularity
}

// GetPopularityOk returns a tuple with the Popularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FavoritedWorld) GetPopularityOk() (*int32, bool) {
	if o == nil || IsNil(o.Popularity) {
		return nil, false
	}
	return o.Popularity, true
}

// HasPopularity returns a boolean if a field has been set.
func (o *FavoritedWorld) HasPopularity() bool {
	if o != nil && !IsNil(o.Popularity) {
		return true
	}

	return false
}

// SetPopularity gets a reference to the given int32 and assigns it to the Popularity field.
func (o *FavoritedWorld) SetPopularity(v int32) {
	o.Popularity = &v
}

// GetPreviewYoutubeId returns the PreviewYoutubeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FavoritedWorld) GetPreviewYoutubeId() string {
	if o == nil || IsNil(o.PreviewYoutubeId.Get()) {
		var ret string
		return ret
	}
	return *o.PreviewYoutubeId.Get()
}

// GetPreviewYoutubeIdOk returns a tuple with the PreviewYoutubeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FavoritedWorld) GetPreviewYoutubeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviewYoutubeId.Get(), o.PreviewYoutubeId.IsSet()
}

// HasPreviewYoutubeId returns a boolean if a field has been set.
func (o *FavoritedWorld) HasPreviewYoutubeId() bool {
	if o != nil && o.PreviewYoutubeId.IsSet() {
		return true
	}

	return false
}

// SetPreviewYoutubeId gets a reference to the given NullableString and assigns it to the PreviewYoutubeId field.
func (o *FavoritedWorld) SetPreviewYoutubeId(v string) {
	o.PreviewYoutubeId.Set(&v)
}

// SetPreviewYoutubeIdNil sets the value for PreviewYoutubeId to be an explicit nil
func (o *FavoritedWorld) SetPreviewYoutubeIdNil() {
	o.PreviewYoutubeId.Set(nil)
}

// UnsetPreviewYoutubeId ensures that no value is present for PreviewYoutubeId, not even an explicit nil
func (o *FavoritedWorld) UnsetPreviewYoutubeId() {
	o.PreviewYoutubeId.Unset()
}

// GetPublicationDate returns the PublicationDate field value if set, zero value otherwise.
func (o *FavoritedWorld) GetPublicationDate() string {
	if o == nil || IsNil(o.PublicationDate) {
		var ret string
		return ret
	}
	return *o.PublicationDate
}

// GetPublicationDateOk returns a tuple with the PublicationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FavoritedWorld) GetPublicationDateOk() (*string, bool) {
	if o == nil || IsNil(o.PublicationDate) {
		return nil, false
	}
	return o.PublicationDate, true
}

// HasPublicationDate returns a boolean if a field has been set.
func (o *FavoritedWorld) HasPublicationDate() bool {
	if o != nil && !IsNil(o.PublicationDate) {
		return true
	}

	return false
}

// SetPublicationDate gets a reference to the given string and assigns it to the PublicationDate field.
func (o *FavoritedWorld) SetPublicationDate(v string) {
	o.PublicationDate = &v
}

// GetRecommendedCapacity returns the RecommendedCapacity field value if set, zero value otherwise.
func (o *FavoritedWorld) GetRecommendedCapacity() int32 {
	if o == nil || IsNil(o.RecommendedCapacity) {
		var ret int32
		return ret
	}
	return *o.RecommendedCapacity
}

// GetRecommendedCapacityOk returns a tuple with the RecommendedCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FavoritedWorld) GetRecommendedCapacityOk() (*int32, bool) {
	if o == nil || IsNil(o.RecommendedCapacity) {
		return nil, false
	}
	return o.RecommendedCapacity, true
}

// HasRecommendedCapacity returns a boolean if a field has been set.
func (o *FavoritedWorld) HasRecommendedCapacity() bool {
	if o != nil && !IsNil(o.RecommendedCapacity) {
		return true
	}

	return false
}

// SetRecommendedCapacity gets a reference to the given int32 and assigns it to the RecommendedCapacity field.
func (o *FavoritedWorld) SetRecommendedCapacity(v int32) {
	o.RecommendedCapacity = &v
}

// GetReleaseStatus returns the ReleaseStatus field value
func (o *FavoritedWorld) GetReleaseStatus() ReleaseStatus {
	if o == nil {
		var ret ReleaseStatus
		return ret
	}

	return o.ReleaseStatus
}

// GetReleaseStatusOk returns a tuple with the ReleaseStatus field value
// and a boolean to check if the value has been set.
func (o *FavoritedWorld) GetReleaseStatusOk() (*ReleaseStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleaseStatus, true
}

// SetReleaseStatus sets field value
func (o *FavoritedWorld) SetReleaseStatus(v ReleaseStatus) {
	o.ReleaseStatus = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FavoritedWorld) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FavoritedWorld) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FavoritedWorld) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FavoritedWorld) SetTags(v []string) {
	o.Tags = v
}

// GetThumbnailImageUrl returns the ThumbnailImageUrl field value
func (o *FavoritedWorld) GetThumbnailImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThumbnailImageUrl
}

// GetThumbnailImageUrlOk returns a tuple with the ThumbnailImageUrl field value
// and a boolean to check if the value has been set.
func (o *FavoritedWorld) GetThumbnailImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThumbnailImageUrl, true
}

// SetThumbnailImageUrl sets field value
func (o *FavoritedWorld) SetThumbnailImageUrl(v string) {
	o.ThumbnailImageUrl = v
}

// GetUdonProducts returns the UdonProducts field value if set, zero value otherwise.
func (o *FavoritedWorld) GetUdonProducts() []string {
	if o == nil || IsNil(o.UdonProducts) {
		var ret []string
		return ret
	}
	return o.UdonProducts
}

// GetUdonProductsOk returns a tuple with the UdonProducts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FavoritedWorld) GetUdonProductsOk() ([]string, bool) {
	if o == nil || IsNil(o.UdonProducts) {
		return nil, false
	}
	return o.UdonProducts, true
}

// HasUdonProducts returns a boolean if a field has been set.
func (o *FavoritedWorld) HasUdonProducts() bool {
	if o != nil && !IsNil(o.UdonProducts) {
		return true
	}

	return false
}

// SetUdonProducts gets a reference to the given []string and assigns it to the UdonProducts field.
func (o *FavoritedWorld) SetUdonProducts(v []string) {
	o.UdonProducts = v
}

// GetUnityPackages returns the UnityPackages field value if set, zero value otherwise.
func (o *FavoritedWorld) GetUnityPackages() []UnityPackage {
	if o == nil || IsNil(o.UnityPackages) {
		var ret []UnityPackage
		return ret
	}
	return o.UnityPackages
}

// GetUnityPackagesOk returns a tuple with the UnityPackages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FavoritedWorld) GetUnityPackagesOk() ([]UnityPackage, bool) {
	if o == nil || IsNil(o.UnityPackages) {
		return nil, false
	}
	return o.UnityPackages, true
}

// HasUnityPackages returns a boolean if a field has been set.
func (o *FavoritedWorld) HasUnityPackages() bool {
	if o != nil && !IsNil(o.UnityPackages) {
		return true
	}

	return false
}

// SetUnityPackages gets a reference to the given []UnityPackage and assigns it to the UnityPackages field.
func (o *FavoritedWorld) SetUnityPackages(v []UnityPackage) {
	o.UnityPackages = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *FavoritedWorld) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FavoritedWorld) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *FavoritedWorld) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *FavoritedWorld) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUrlList returns the UrlList field value if set, zero value otherwise.
func (o *FavoritedWorld) GetUrlList() []string {
	if o == nil || IsNil(o.UrlList) {
		var ret []string
		return ret
	}
	return o.UrlList
}

// GetUrlListOk returns a tuple with the UrlList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FavoritedWorld) GetUrlListOk() ([]string, bool) {
	if o == nil || IsNil(o.UrlList) {
		return nil, false
	}
	return o.UrlList, true
}

// HasUrlList returns a boolean if a field has been set.
func (o *FavoritedWorld) HasUrlList() bool {
	if o != nil && !IsNil(o.UrlList) {
		return true
	}

	return false
}

// SetUrlList gets a reference to the given []string and assigns it to the UrlList field.
func (o *FavoritedWorld) SetUrlList(v []string) {
	o.UrlList = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *FavoritedWorld) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FavoritedWorld) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *FavoritedWorld) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *FavoritedWorld) SetVersion(v int32) {
	o.Version = &v
}

// GetVisits returns the Visits field value if set, zero value otherwise.
func (o *FavoritedWorld) GetVisits() int32 {
	if o == nil || IsNil(o.Visits) {
		var ret int32
		return ret
	}
	return *o.Visits
}

// GetVisitsOk returns a tuple with the Visits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FavoritedWorld) GetVisitsOk() (*int32, bool) {
	if o == nil || IsNil(o.Visits) {
		return nil, false
	}
	return o.Visits, true
}

// HasVisits returns a boolean if a field has been set.
func (o *FavoritedWorld) HasVisits() bool {
	if o != nil && !IsNil(o.Visits) {
		return true
	}

	return false
}

// SetVisits gets a reference to the given int32 and assigns it to the Visits field.
func (o *FavoritedWorld) SetVisits(v int32) {
	o.Visits = &v
}

func (o FavoritedWorld) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FavoritedWorld) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthorId) {
		toSerialize["authorId"] = o.AuthorId
	}
	toSerialize["authorName"] = o.AuthorName
	toSerialize["capacity"] = o.Capacity
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.DefaultContentSettings) {
		toSerialize["defaultContentSettings"] = o.DefaultContentSettings
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["favoriteGroup"] = o.FavoriteGroup
	toSerialize["favoriteId"] = o.FavoriteId
	if !IsNil(o.Favorites) {
		toSerialize["favorites"] = o.Favorites
	}
	if !IsNil(o.Featured) {
		toSerialize["featured"] = o.Featured
	}
	if !IsNil(o.Heat) {
		toSerialize["heat"] = o.Heat
	}
	toSerialize["id"] = o.Id
	toSerialize["imageUrl"] = o.ImageUrl
	if !IsNil(o.LabsPublicationDate) {
		toSerialize["labsPublicationDate"] = o.LabsPublicationDate
	}
	toSerialize["name"] = o.Name
	toSerialize["occupants"] = o.Occupants
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.Popularity) {
		toSerialize["popularity"] = o.Popularity
	}
	if o.PreviewYoutubeId.IsSet() {
		toSerialize["previewYoutubeId"] = o.PreviewYoutubeId.Get()
	}
	if !IsNil(o.PublicationDate) {
		toSerialize["publicationDate"] = o.PublicationDate
	}
	if !IsNil(o.RecommendedCapacity) {
		toSerialize["recommendedCapacity"] = o.RecommendedCapacity
	}
	toSerialize["releaseStatus"] = o.ReleaseStatus
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["thumbnailImageUrl"] = o.ThumbnailImageUrl
	if !IsNil(o.UdonProducts) {
		toSerialize["udonProducts"] = o.UdonProducts
	}
	if !IsNil(o.UnityPackages) {
		toSerialize["unityPackages"] = o.UnityPackages
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.UrlList) {
		toSerialize["urlList"] = o.UrlList
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Visits) {
		toSerialize["visits"] = o.Visits
	}
	return toSerialize, nil
}

func (o *FavoritedWorld) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authorName",
		"capacity",
		"favoriteGroup",
		"favoriteId",
		"id",
		"imageUrl",
		"name",
		"occupants",
		"releaseStatus",
		"thumbnailImageUrl",
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

	varFavoritedWorld := _FavoritedWorld{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFavoritedWorld)

	if err != nil {
		return err
	}

	*o = FavoritedWorld(varFavoritedWorld)

	return err
}

type NullableFavoritedWorld struct {
	value *FavoritedWorld
	isSet bool
}

func (v NullableFavoritedWorld) Get() *FavoritedWorld {
	return v.value
}

func (v *NullableFavoritedWorld) Set(val *FavoritedWorld) {
	v.value = val
	v.isSet = true
}

func (v NullableFavoritedWorld) IsSet() bool {
	return v.isSet
}

func (v *NullableFavoritedWorld) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFavoritedWorld(val *FavoritedWorld) *NullableFavoritedWorld {
	return &NullableFavoritedWorld{value: val, isSet: true}
}

func (v NullableFavoritedWorld) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFavoritedWorld) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
