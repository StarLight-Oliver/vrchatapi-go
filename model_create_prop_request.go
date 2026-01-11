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
	"bytes"
	"fmt"
)

// checks if the CreatePropRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePropRequest{}

// CreatePropRequest struct for CreatePropRequest
type CreatePropRequest struct {
	AssetUrl string `json:"assetUrl"`
	AssetVersion int32 `json:"assetVersion"`
	Description string `json:"description"`
	Id string `json:"id"`
	ImageUrl string `json:"imageUrl"`
	Name string `json:"name"`
	// This is normally `android`, `ios`, `standalonewindows`, `web`, or the empty value ``, but also supposedly can be any random Unity version such as `2019.2.4-801-Release` or `2019.2.2-772-Release` or even `unknownplatform`.
	Platform string `json:"platform"`
	PropSignature *string `json:"propSignature,omitempty"`
	// How a prop is summoned and interacted with. 0: the prop fixed to some surface in the world 1: the prop is a pickup and may be held by users 2: ???
	SpawnType int32 `json:"spawnType"`
	Tags []string `json:"tags"`
	UnityVersion string `json:"unityVersion"`
	// Bitmask for restrictions on what world surfaces a prop may be summoned. 0: no restrictions 1: floors 2: walls 4: ceilings
	WorldPlacementMask int32 `json:"worldPlacementMask"`
}

type _CreatePropRequest CreatePropRequest

// NewCreatePropRequest instantiates a new CreatePropRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePropRequest(assetUrl string, assetVersion int32, description string, id string, imageUrl string, name string, platform string, spawnType int32, tags []string, unityVersion string, worldPlacementMask int32) *CreatePropRequest {
	this := CreatePropRequest{}
	this.AssetUrl = assetUrl
	this.AssetVersion = assetVersion
	this.Description = description
	this.Id = id
	this.ImageUrl = imageUrl
	this.Name = name
	this.Platform = platform
	this.SpawnType = spawnType
	this.Tags = tags
	this.UnityVersion = unityVersion
	this.WorldPlacementMask = worldPlacementMask
	return &this
}

// NewCreatePropRequestWithDefaults instantiates a new CreatePropRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePropRequestWithDefaults() *CreatePropRequest {
	this := CreatePropRequest{}
	var spawnType int32 = 1
	this.SpawnType = spawnType
	var worldPlacementMask int32 = 1
	this.WorldPlacementMask = worldPlacementMask
	return &this
}

// GetAssetUrl returns the AssetUrl field value
func (o *CreatePropRequest) GetAssetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetUrl
}

// GetAssetUrlOk returns a tuple with the AssetUrl field value
// and a boolean to check if the value has been set.
func (o *CreatePropRequest) GetAssetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetUrl, true
}

// SetAssetUrl sets field value
func (o *CreatePropRequest) SetAssetUrl(v string) {
	o.AssetUrl = v
}

// GetAssetVersion returns the AssetVersion field value
func (o *CreatePropRequest) GetAssetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AssetVersion
}

// GetAssetVersionOk returns a tuple with the AssetVersion field value
// and a boolean to check if the value has been set.
func (o *CreatePropRequest) GetAssetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetVersion, true
}

// SetAssetVersion sets field value
func (o *CreatePropRequest) SetAssetVersion(v int32) {
	o.AssetVersion = v
}

// GetDescription returns the Description field value
func (o *CreatePropRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreatePropRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreatePropRequest) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value
func (o *CreatePropRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreatePropRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreatePropRequest) SetId(v string) {
	o.Id = v
}

// GetImageUrl returns the ImageUrl field value
func (o *CreatePropRequest) GetImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
func (o *CreatePropRequest) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageUrl, true
}

// SetImageUrl sets field value
func (o *CreatePropRequest) SetImageUrl(v string) {
	o.ImageUrl = v
}

// GetName returns the Name field value
func (o *CreatePropRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreatePropRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreatePropRequest) SetName(v string) {
	o.Name = v
}

// GetPlatform returns the Platform field value
func (o *CreatePropRequest) GetPlatform() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value
// and a boolean to check if the value has been set.
func (o *CreatePropRequest) GetPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Platform, true
}

// SetPlatform sets field value
func (o *CreatePropRequest) SetPlatform(v string) {
	o.Platform = v
}

// GetPropSignature returns the PropSignature field value if set, zero value otherwise.
func (o *CreatePropRequest) GetPropSignature() string {
	if o == nil || IsNil(o.PropSignature) {
		var ret string
		return ret
	}
	return *o.PropSignature
}

// GetPropSignatureOk returns a tuple with the PropSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePropRequest) GetPropSignatureOk() (*string, bool) {
	if o == nil || IsNil(o.PropSignature) {
		return nil, false
	}
	return o.PropSignature, true
}

// HasPropSignature returns a boolean if a field has been set.
func (o *CreatePropRequest) HasPropSignature() bool {
	if o != nil && !IsNil(o.PropSignature) {
		return true
	}

	return false
}

// SetPropSignature gets a reference to the given string and assigns it to the PropSignature field.
func (o *CreatePropRequest) SetPropSignature(v string) {
	o.PropSignature = &v
}

// GetSpawnType returns the SpawnType field value
func (o *CreatePropRequest) GetSpawnType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SpawnType
}

// GetSpawnTypeOk returns a tuple with the SpawnType field value
// and a boolean to check if the value has been set.
func (o *CreatePropRequest) GetSpawnTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpawnType, true
}

// SetSpawnType sets field value
func (o *CreatePropRequest) SetSpawnType(v int32) {
	o.SpawnType = v
}

// GetTags returns the Tags field value
func (o *CreatePropRequest) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *CreatePropRequest) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *CreatePropRequest) SetTags(v []string) {
	o.Tags = v
}

// GetUnityVersion returns the UnityVersion field value
func (o *CreatePropRequest) GetUnityVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnityVersion
}

// GetUnityVersionOk returns a tuple with the UnityVersion field value
// and a boolean to check if the value has been set.
func (o *CreatePropRequest) GetUnityVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnityVersion, true
}

// SetUnityVersion sets field value
func (o *CreatePropRequest) SetUnityVersion(v string) {
	o.UnityVersion = v
}

// GetWorldPlacementMask returns the WorldPlacementMask field value
func (o *CreatePropRequest) GetWorldPlacementMask() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WorldPlacementMask
}

// GetWorldPlacementMaskOk returns a tuple with the WorldPlacementMask field value
// and a boolean to check if the value has been set.
func (o *CreatePropRequest) GetWorldPlacementMaskOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorldPlacementMask, true
}

// SetWorldPlacementMask sets field value
func (o *CreatePropRequest) SetWorldPlacementMask(v int32) {
	o.WorldPlacementMask = v
}

func (o CreatePropRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePropRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["assetUrl"] = o.AssetUrl
	toSerialize["assetVersion"] = o.AssetVersion
	toSerialize["description"] = o.Description
	toSerialize["id"] = o.Id
	toSerialize["imageUrl"] = o.ImageUrl
	toSerialize["name"] = o.Name
	toSerialize["platform"] = o.Platform
	if !IsNil(o.PropSignature) {
		toSerialize["propSignature"] = o.PropSignature
	}
	toSerialize["spawnType"] = o.SpawnType
	toSerialize["tags"] = o.Tags
	toSerialize["unityVersion"] = o.UnityVersion
	toSerialize["worldPlacementMask"] = o.WorldPlacementMask
	return toSerialize, nil
}

func (o *CreatePropRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"assetUrl",
		"assetVersion",
		"description",
		"id",
		"imageUrl",
		"name",
		"platform",
		"spawnType",
		"tags",
		"unityVersion",
		"worldPlacementMask",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreatePropRequest := _CreatePropRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreatePropRequest)

	if err != nil {
		return err
	}

	*o = CreatePropRequest(varCreatePropRequest)

	return err
}

type NullableCreatePropRequest struct {
	value *CreatePropRequest
	isSet bool
}

func (v NullableCreatePropRequest) Get() *CreatePropRequest {
	return v.value
}

func (v *NullableCreatePropRequest) Set(val *CreatePropRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePropRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePropRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePropRequest(val *CreatePropRequest) *NullableCreatePropRequest {
	return &NullableCreatePropRequest{value: val, isSet: true}
}

func (v NullableCreatePropRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePropRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


