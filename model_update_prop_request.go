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
)

// checks if the UpdatePropRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePropRequest{}

// UpdatePropRequest struct for UpdatePropRequest
type UpdatePropRequest struct {
	AssetUrl *string `json:"assetUrl,omitempty"`
	AssetVersion *int32 `json:"assetVersion,omitempty"`
	Description *string `json:"description,omitempty"`
	ImageUrl *string `json:"imageUrl,omitempty"`
	Name *string `json:"name,omitempty"`
	// This is normally `android`, `ios`, `standalonewindows`, `web`, or the empty value ``, but also supposedly can be any random Unity version such as `2019.2.4-801-Release` or `2019.2.2-772-Release` or even `unknownplatform`.
	Platform *string `json:"platform,omitempty"`
	PropSignature *string `json:"propSignature,omitempty"`
	// How a prop is summoned and interacted with. 0: the prop fixed to some surface in the world 1: the prop is a pickup and may be held by users 2: ???
	SpawnType *int32 `json:"spawnType,omitempty"`
	Tags []string `json:"tags,omitempty"`
	UnityVersion *string `json:"unityVersion,omitempty"`
	// Bitmask for restrictions on what world surfaces a prop may be summoned. 0: no restrictions 1: floors 2: walls 4: ceilings
	WorldPlacementMask *int32 `json:"worldPlacementMask,omitempty"`
}

// NewUpdatePropRequest instantiates a new UpdatePropRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePropRequest() *UpdatePropRequest {
	this := UpdatePropRequest{}
	var spawnType int32 = 1
	this.SpawnType = &spawnType
	var worldPlacementMask int32 = 1
	this.WorldPlacementMask = &worldPlacementMask
	return &this
}

// NewUpdatePropRequestWithDefaults instantiates a new UpdatePropRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePropRequestWithDefaults() *UpdatePropRequest {
	this := UpdatePropRequest{}
	var spawnType int32 = 1
	this.SpawnType = &spawnType
	var worldPlacementMask int32 = 1
	this.WorldPlacementMask = &worldPlacementMask
	return &this
}

// GetAssetUrl returns the AssetUrl field value if set, zero value otherwise.
func (o *UpdatePropRequest) GetAssetUrl() string {
	if o == nil || IsNil(o.AssetUrl) {
		var ret string
		return ret
	}
	return *o.AssetUrl
}

// GetAssetUrlOk returns a tuple with the AssetUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePropRequest) GetAssetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AssetUrl) {
		return nil, false
	}
	return o.AssetUrl, true
}

// HasAssetUrl returns a boolean if a field has been set.
func (o *UpdatePropRequest) HasAssetUrl() bool {
	if o != nil && !IsNil(o.AssetUrl) {
		return true
	}

	return false
}

// SetAssetUrl gets a reference to the given string and assigns it to the AssetUrl field.
func (o *UpdatePropRequest) SetAssetUrl(v string) {
	o.AssetUrl = &v
}

// GetAssetVersion returns the AssetVersion field value if set, zero value otherwise.
func (o *UpdatePropRequest) GetAssetVersion() int32 {
	if o == nil || IsNil(o.AssetVersion) {
		var ret int32
		return ret
	}
	return *o.AssetVersion
}

// GetAssetVersionOk returns a tuple with the AssetVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePropRequest) GetAssetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.AssetVersion) {
		return nil, false
	}
	return o.AssetVersion, true
}

// HasAssetVersion returns a boolean if a field has been set.
func (o *UpdatePropRequest) HasAssetVersion() bool {
	if o != nil && !IsNil(o.AssetVersion) {
		return true
	}

	return false
}

// SetAssetVersion gets a reference to the given int32 and assigns it to the AssetVersion field.
func (o *UpdatePropRequest) SetAssetVersion(v int32) {
	o.AssetVersion = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdatePropRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePropRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdatePropRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdatePropRequest) SetDescription(v string) {
	o.Description = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *UpdatePropRequest) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePropRequest) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *UpdatePropRequest) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *UpdatePropRequest) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdatePropRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePropRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdatePropRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdatePropRequest) SetName(v string) {
	o.Name = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *UpdatePropRequest) GetPlatform() string {
	if o == nil || IsNil(o.Platform) {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePropRequest) GetPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *UpdatePropRequest) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *UpdatePropRequest) SetPlatform(v string) {
	o.Platform = &v
}

// GetPropSignature returns the PropSignature field value if set, zero value otherwise.
func (o *UpdatePropRequest) GetPropSignature() string {
	if o == nil || IsNil(o.PropSignature) {
		var ret string
		return ret
	}
	return *o.PropSignature
}

// GetPropSignatureOk returns a tuple with the PropSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePropRequest) GetPropSignatureOk() (*string, bool) {
	if o == nil || IsNil(o.PropSignature) {
		return nil, false
	}
	return o.PropSignature, true
}

// HasPropSignature returns a boolean if a field has been set.
func (o *UpdatePropRequest) HasPropSignature() bool {
	if o != nil && !IsNil(o.PropSignature) {
		return true
	}

	return false
}

// SetPropSignature gets a reference to the given string and assigns it to the PropSignature field.
func (o *UpdatePropRequest) SetPropSignature(v string) {
	o.PropSignature = &v
}

// GetSpawnType returns the SpawnType field value if set, zero value otherwise.
func (o *UpdatePropRequest) GetSpawnType() int32 {
	if o == nil || IsNil(o.SpawnType) {
		var ret int32
		return ret
	}
	return *o.SpawnType
}

// GetSpawnTypeOk returns a tuple with the SpawnType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePropRequest) GetSpawnTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.SpawnType) {
		return nil, false
	}
	return o.SpawnType, true
}

// HasSpawnType returns a boolean if a field has been set.
func (o *UpdatePropRequest) HasSpawnType() bool {
	if o != nil && !IsNil(o.SpawnType) {
		return true
	}

	return false
}

// SetSpawnType gets a reference to the given int32 and assigns it to the SpawnType field.
func (o *UpdatePropRequest) SetSpawnType(v int32) {
	o.SpawnType = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdatePropRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePropRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdatePropRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UpdatePropRequest) SetTags(v []string) {
	o.Tags = v
}

// GetUnityVersion returns the UnityVersion field value if set, zero value otherwise.
func (o *UpdatePropRequest) GetUnityVersion() string {
	if o == nil || IsNil(o.UnityVersion) {
		var ret string
		return ret
	}
	return *o.UnityVersion
}

// GetUnityVersionOk returns a tuple with the UnityVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePropRequest) GetUnityVersionOk() (*string, bool) {
	if o == nil || IsNil(o.UnityVersion) {
		return nil, false
	}
	return o.UnityVersion, true
}

// HasUnityVersion returns a boolean if a field has been set.
func (o *UpdatePropRequest) HasUnityVersion() bool {
	if o != nil && !IsNil(o.UnityVersion) {
		return true
	}

	return false
}

// SetUnityVersion gets a reference to the given string and assigns it to the UnityVersion field.
func (o *UpdatePropRequest) SetUnityVersion(v string) {
	o.UnityVersion = &v
}

// GetWorldPlacementMask returns the WorldPlacementMask field value if set, zero value otherwise.
func (o *UpdatePropRequest) GetWorldPlacementMask() int32 {
	if o == nil || IsNil(o.WorldPlacementMask) {
		var ret int32
		return ret
	}
	return *o.WorldPlacementMask
}

// GetWorldPlacementMaskOk returns a tuple with the WorldPlacementMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePropRequest) GetWorldPlacementMaskOk() (*int32, bool) {
	if o == nil || IsNil(o.WorldPlacementMask) {
		return nil, false
	}
	return o.WorldPlacementMask, true
}

// HasWorldPlacementMask returns a boolean if a field has been set.
func (o *UpdatePropRequest) HasWorldPlacementMask() bool {
	if o != nil && !IsNil(o.WorldPlacementMask) {
		return true
	}

	return false
}

// SetWorldPlacementMask gets a reference to the given int32 and assigns it to the WorldPlacementMask field.
func (o *UpdatePropRequest) SetWorldPlacementMask(v int32) {
	o.WorldPlacementMask = &v
}

func (o UpdatePropRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePropRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssetUrl) {
		toSerialize["assetUrl"] = o.AssetUrl
	}
	if !IsNil(o.AssetVersion) {
		toSerialize["assetVersion"] = o.AssetVersion
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.PropSignature) {
		toSerialize["propSignature"] = o.PropSignature
	}
	if !IsNil(o.SpawnType) {
		toSerialize["spawnType"] = o.SpawnType
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.UnityVersion) {
		toSerialize["unityVersion"] = o.UnityVersion
	}
	if !IsNil(o.WorldPlacementMask) {
		toSerialize["worldPlacementMask"] = o.WorldPlacementMask
	}
	return toSerialize, nil
}

type NullableUpdatePropRequest struct {
	value *UpdatePropRequest
	isSet bool
}

func (v NullableUpdatePropRequest) Get() *UpdatePropRequest {
	return v.value
}

func (v *NullableUpdatePropRequest) Set(val *UpdatePropRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePropRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePropRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePropRequest(val *UpdatePropRequest) *NullableUpdatePropRequest {
	return &NullableUpdatePropRequest{value: val, isSet: true}
}

func (v NullableUpdatePropRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePropRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


