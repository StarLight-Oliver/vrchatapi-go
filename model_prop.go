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

// checks if the Prop type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Prop{}

// Prop
type Prop struct {
	CreatedAt time.Time `json:"_created_at"`
	UpdatedAt time.Time `json:"_updated_at"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	AuthorId        string        `json:"authorId"`
	AuthorName      string        `json:"authorName"`
	Description     string        `json:"description"`
	Id              string        `json:"id"`
	ImageUrl        string        `json:"imageUrl"`
	MaxCountPerUser int32         `json:"maxCountPerUser"`
	Name            string        `json:"name"`
	ReleaseStatus   ReleaseStatus `json:"releaseStatus"`
	// How a prop is summoned and interacted with. 0: the prop fixed to some surface in the world 1: the prop is a pickup and may be held by users 2: ???
	SpawnType         int32              `json:"spawnType"`
	Tags              []string           `json:"tags"`
	ThumbnailImageUrl string             `json:"thumbnailImageUrl"`
	UnityPackageUrl   NullableString     `json:"unityPackageUrl"`
	UnityPackages     []PropUnityPackage `json:"unityPackages"`
	// Bitmask for restrictions on what world surfaces a prop may be summoned. 0: no restrictions 1: floors 2: walls 4: ceilings
	WorldPlacementMask int32 `json:"worldPlacementMask"`
}

type _Prop Prop

// NewProp instantiates a new Prop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProp(createdAt time.Time, updatedAt time.Time, authorId string, authorName string, description string, id string, imageUrl string, maxCountPerUser int32, name string, releaseStatus ReleaseStatus, spawnType int32, tags []string, thumbnailImageUrl string, unityPackageUrl NullableString, unityPackages []PropUnityPackage, worldPlacementMask int32) *Prop {
	this := Prop{}
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.AuthorId = authorId
	this.AuthorName = authorName
	this.Description = description
	this.Id = id
	this.ImageUrl = imageUrl
	this.MaxCountPerUser = maxCountPerUser
	this.Name = name
	this.ReleaseStatus = releaseStatus
	this.SpawnType = spawnType
	this.Tags = tags
	this.ThumbnailImageUrl = thumbnailImageUrl
	this.UnityPackageUrl = unityPackageUrl
	this.UnityPackages = unityPackages
	this.WorldPlacementMask = worldPlacementMask
	return &this
}

// NewPropWithDefaults instantiates a new Prop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropWithDefaults() *Prop {
	this := Prop{}
	var maxCountPerUser int32 = 1
	this.MaxCountPerUser = maxCountPerUser
	var releaseStatus ReleaseStatus = ReleaseStatus_PUBLIC
	this.ReleaseStatus = releaseStatus
	var spawnType int32 = 1
	this.SpawnType = spawnType
	var worldPlacementMask int32 = 1
	this.WorldPlacementMask = worldPlacementMask
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *Prop) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Prop) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Prop) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Prop) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Prop) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Prop) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetAuthorId returns the AuthorId field value
func (o *Prop) GetAuthorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorId
}

// GetAuthorIdOk returns a tuple with the AuthorId field value
// and a boolean to check if the value has been set.
func (o *Prop) GetAuthorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorId, true
}

// SetAuthorId sets field value
func (o *Prop) SetAuthorId(v string) {
	o.AuthorId = v
}

// GetAuthorName returns the AuthorName field value
func (o *Prop) GetAuthorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorName
}

// GetAuthorNameOk returns a tuple with the AuthorName field value
// and a boolean to check if the value has been set.
func (o *Prop) GetAuthorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorName, true
}

// SetAuthorName sets field value
func (o *Prop) SetAuthorName(v string) {
	o.AuthorName = v
}

// GetDescription returns the Description field value
func (o *Prop) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Prop) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Prop) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value
func (o *Prop) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Prop) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Prop) SetId(v string) {
	o.Id = v
}

// GetImageUrl returns the ImageUrl field value
func (o *Prop) GetImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
func (o *Prop) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageUrl, true
}

// SetImageUrl sets field value
func (o *Prop) SetImageUrl(v string) {
	o.ImageUrl = v
}

// GetMaxCountPerUser returns the MaxCountPerUser field value
func (o *Prop) GetMaxCountPerUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxCountPerUser
}

// GetMaxCountPerUserOk returns a tuple with the MaxCountPerUser field value
// and a boolean to check if the value has been set.
func (o *Prop) GetMaxCountPerUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxCountPerUser, true
}

// SetMaxCountPerUser sets field value
func (o *Prop) SetMaxCountPerUser(v int32) {
	o.MaxCountPerUser = v
}

// GetName returns the Name field value
func (o *Prop) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Prop) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Prop) SetName(v string) {
	o.Name = v
}

// GetReleaseStatus returns the ReleaseStatus field value
func (o *Prop) GetReleaseStatus() ReleaseStatus {
	if o == nil {
		var ret ReleaseStatus
		return ret
	}

	return o.ReleaseStatus
}

// GetReleaseStatusOk returns a tuple with the ReleaseStatus field value
// and a boolean to check if the value has been set.
func (o *Prop) GetReleaseStatusOk() (*ReleaseStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleaseStatus, true
}

// SetReleaseStatus sets field value
func (o *Prop) SetReleaseStatus(v ReleaseStatus) {
	o.ReleaseStatus = v
}

// GetSpawnType returns the SpawnType field value
func (o *Prop) GetSpawnType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SpawnType
}

// GetSpawnTypeOk returns a tuple with the SpawnType field value
// and a boolean to check if the value has been set.
func (o *Prop) GetSpawnTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpawnType, true
}

// SetSpawnType sets field value
func (o *Prop) SetSpawnType(v int32) {
	o.SpawnType = v
}

// GetTags returns the Tags field value
func (o *Prop) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *Prop) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *Prop) SetTags(v []string) {
	o.Tags = v
}

// GetThumbnailImageUrl returns the ThumbnailImageUrl field value
func (o *Prop) GetThumbnailImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThumbnailImageUrl
}

// GetThumbnailImageUrlOk returns a tuple with the ThumbnailImageUrl field value
// and a boolean to check if the value has been set.
func (o *Prop) GetThumbnailImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThumbnailImageUrl, true
}

// SetThumbnailImageUrl sets field value
func (o *Prop) SetThumbnailImageUrl(v string) {
	o.ThumbnailImageUrl = v
}

// GetUnityPackageUrl returns the UnityPackageUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Prop) GetUnityPackageUrl() string {
	if o == nil || o.UnityPackageUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.UnityPackageUrl.Get()
}

// GetUnityPackageUrlOk returns a tuple with the UnityPackageUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Prop) GetUnityPackageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnityPackageUrl.Get(), o.UnityPackageUrl.IsSet()
}

// SetUnityPackageUrl sets field value
func (o *Prop) SetUnityPackageUrl(v string) {
	o.UnityPackageUrl.Set(&v)
}

// GetUnityPackages returns the UnityPackages field value
func (o *Prop) GetUnityPackages() []PropUnityPackage {
	if o == nil {
		var ret []PropUnityPackage
		return ret
	}

	return o.UnityPackages
}

// GetUnityPackagesOk returns a tuple with the UnityPackages field value
// and a boolean to check if the value has been set.
func (o *Prop) GetUnityPackagesOk() ([]PropUnityPackage, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnityPackages, true
}

// SetUnityPackages sets field value
func (o *Prop) SetUnityPackages(v []PropUnityPackage) {
	o.UnityPackages = v
}

// GetWorldPlacementMask returns the WorldPlacementMask field value
func (o *Prop) GetWorldPlacementMask() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WorldPlacementMask
}

// GetWorldPlacementMaskOk returns a tuple with the WorldPlacementMask field value
// and a boolean to check if the value has been set.
func (o *Prop) GetWorldPlacementMaskOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorldPlacementMask, true
}

// SetWorldPlacementMask sets field value
func (o *Prop) SetWorldPlacementMask(v int32) {
	o.WorldPlacementMask = v
}

func (o Prop) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Prop) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_created_at"] = o.CreatedAt
	toSerialize["_updated_at"] = o.UpdatedAt
	toSerialize["authorId"] = o.AuthorId
	toSerialize["authorName"] = o.AuthorName
	toSerialize["description"] = o.Description
	toSerialize["id"] = o.Id
	toSerialize["imageUrl"] = o.ImageUrl
	toSerialize["maxCountPerUser"] = o.MaxCountPerUser
	toSerialize["name"] = o.Name
	toSerialize["releaseStatus"] = o.ReleaseStatus
	toSerialize["spawnType"] = o.SpawnType
	toSerialize["tags"] = o.Tags
	toSerialize["thumbnailImageUrl"] = o.ThumbnailImageUrl
	toSerialize["unityPackageUrl"] = o.UnityPackageUrl.Get()
	toSerialize["unityPackages"] = o.UnityPackages
	toSerialize["worldPlacementMask"] = o.WorldPlacementMask
	return toSerialize, nil
}

func (o *Prop) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_created_at",
		"_updated_at",
		"authorId",
		"authorName",
		"description",
		"id",
		"imageUrl",
		"maxCountPerUser",
		"name",
		"releaseStatus",
		"spawnType",
		"tags",
		"thumbnailImageUrl",
		"unityPackageUrl",
		"unityPackages",
		"worldPlacementMask",
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

	varProp := _Prop{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProp)

	if err != nil {
		return err
	}

	*o = Prop(varProp)

	return err
}

type NullableProp struct {
	value *Prop
	isSet bool
}

func (v NullableProp) Get() *Prop {
	return v.value
}

func (v *NullableProp) Set(val *Prop) {
	v.value = val
	v.isSet = true
}

func (v NullableProp) IsSet() bool {
	return v.isSet
}

func (v *NullableProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProp(val *Prop) *NullableProp {
	return &NullableProp{value: val, isSet: true}
}

func (v NullableProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
