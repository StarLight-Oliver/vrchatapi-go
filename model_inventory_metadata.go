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

// checks if the InventoryMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryMetadata{}

// InventoryMetadata struct for InventoryMetadata
type InventoryMetadata struct {
	Animated *bool `json:"animated,omitempty"`
	AnimationStyle *string `json:"animationStyle,omitempty"`
	AssetBundleId *string `json:"assetBundleId,omitempty"`
	FileId *string `json:"fileId,omitempty"`
	ImageUrl *string `json:"imageUrl,omitempty"`
	// Only in bundles
	InventoryItemsToInstantiate []string `json:"inventoryItemsToInstantiate,omitempty"`
	MaskTag *string `json:"maskTag,omitempty"`
	PropId *string `json:"propId,omitempty"`
}

// NewInventoryMetadata instantiates a new InventoryMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryMetadata() *InventoryMetadata {
	this := InventoryMetadata{}
	return &this
}

// NewInventoryMetadataWithDefaults instantiates a new InventoryMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryMetadataWithDefaults() *InventoryMetadata {
	this := InventoryMetadata{}
	return &this
}

// GetAnimated returns the Animated field value if set, zero value otherwise.
func (o *InventoryMetadata) GetAnimated() bool {
	if o == nil || IsNil(o.Animated) {
		var ret bool
		return ret
	}
	return *o.Animated
}

// GetAnimatedOk returns a tuple with the Animated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryMetadata) GetAnimatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Animated) {
		return nil, false
	}
	return o.Animated, true
}

// HasAnimated returns a boolean if a field has been set.
func (o *InventoryMetadata) HasAnimated() bool {
	if o != nil && !IsNil(o.Animated) {
		return true
	}

	return false
}

// SetAnimated gets a reference to the given bool and assigns it to the Animated field.
func (o *InventoryMetadata) SetAnimated(v bool) {
	o.Animated = &v
}

// GetAnimationStyle returns the AnimationStyle field value if set, zero value otherwise.
func (o *InventoryMetadata) GetAnimationStyle() string {
	if o == nil || IsNil(o.AnimationStyle) {
		var ret string
		return ret
	}
	return *o.AnimationStyle
}

// GetAnimationStyleOk returns a tuple with the AnimationStyle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryMetadata) GetAnimationStyleOk() (*string, bool) {
	if o == nil || IsNil(o.AnimationStyle) {
		return nil, false
	}
	return o.AnimationStyle, true
}

// HasAnimationStyle returns a boolean if a field has been set.
func (o *InventoryMetadata) HasAnimationStyle() bool {
	if o != nil && !IsNil(o.AnimationStyle) {
		return true
	}

	return false
}

// SetAnimationStyle gets a reference to the given string and assigns it to the AnimationStyle field.
func (o *InventoryMetadata) SetAnimationStyle(v string) {
	o.AnimationStyle = &v
}

// GetAssetBundleId returns the AssetBundleId field value if set, zero value otherwise.
func (o *InventoryMetadata) GetAssetBundleId() string {
	if o == nil || IsNil(o.AssetBundleId) {
		var ret string
		return ret
	}
	return *o.AssetBundleId
}

// GetAssetBundleIdOk returns a tuple with the AssetBundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryMetadata) GetAssetBundleIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssetBundleId) {
		return nil, false
	}
	return o.AssetBundleId, true
}

// HasAssetBundleId returns a boolean if a field has been set.
func (o *InventoryMetadata) HasAssetBundleId() bool {
	if o != nil && !IsNil(o.AssetBundleId) {
		return true
	}

	return false
}

// SetAssetBundleId gets a reference to the given string and assigns it to the AssetBundleId field.
func (o *InventoryMetadata) SetAssetBundleId(v string) {
	o.AssetBundleId = &v
}

// GetFileId returns the FileId field value if set, zero value otherwise.
func (o *InventoryMetadata) GetFileId() string {
	if o == nil || IsNil(o.FileId) {
		var ret string
		return ret
	}
	return *o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryMetadata) GetFileIdOk() (*string, bool) {
	if o == nil || IsNil(o.FileId) {
		return nil, false
	}
	return o.FileId, true
}

// HasFileId returns a boolean if a field has been set.
func (o *InventoryMetadata) HasFileId() bool {
	if o != nil && !IsNil(o.FileId) {
		return true
	}

	return false
}

// SetFileId gets a reference to the given string and assigns it to the FileId field.
func (o *InventoryMetadata) SetFileId(v string) {
	o.FileId = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *InventoryMetadata) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryMetadata) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *InventoryMetadata) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *InventoryMetadata) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetInventoryItemsToInstantiate returns the InventoryItemsToInstantiate field value if set, zero value otherwise.
func (o *InventoryMetadata) GetInventoryItemsToInstantiate() []string {
	if o == nil || IsNil(o.InventoryItemsToInstantiate) {
		var ret []string
		return ret
	}
	return o.InventoryItemsToInstantiate
}

// GetInventoryItemsToInstantiateOk returns a tuple with the InventoryItemsToInstantiate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryMetadata) GetInventoryItemsToInstantiateOk() ([]string, bool) {
	if o == nil || IsNil(o.InventoryItemsToInstantiate) {
		return nil, false
	}
	return o.InventoryItemsToInstantiate, true
}

// HasInventoryItemsToInstantiate returns a boolean if a field has been set.
func (o *InventoryMetadata) HasInventoryItemsToInstantiate() bool {
	if o != nil && !IsNil(o.InventoryItemsToInstantiate) {
		return true
	}

	return false
}

// SetInventoryItemsToInstantiate gets a reference to the given []string and assigns it to the InventoryItemsToInstantiate field.
func (o *InventoryMetadata) SetInventoryItemsToInstantiate(v []string) {
	o.InventoryItemsToInstantiate = v
}

// GetMaskTag returns the MaskTag field value if set, zero value otherwise.
func (o *InventoryMetadata) GetMaskTag() string {
	if o == nil || IsNil(o.MaskTag) {
		var ret string
		return ret
	}
	return *o.MaskTag
}

// GetMaskTagOk returns a tuple with the MaskTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryMetadata) GetMaskTagOk() (*string, bool) {
	if o == nil || IsNil(o.MaskTag) {
		return nil, false
	}
	return o.MaskTag, true
}

// HasMaskTag returns a boolean if a field has been set.
func (o *InventoryMetadata) HasMaskTag() bool {
	if o != nil && !IsNil(o.MaskTag) {
		return true
	}

	return false
}

// SetMaskTag gets a reference to the given string and assigns it to the MaskTag field.
func (o *InventoryMetadata) SetMaskTag(v string) {
	o.MaskTag = &v
}

// GetPropId returns the PropId field value if set, zero value otherwise.
func (o *InventoryMetadata) GetPropId() string {
	if o == nil || IsNil(o.PropId) {
		var ret string
		return ret
	}
	return *o.PropId
}

// GetPropIdOk returns a tuple with the PropId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryMetadata) GetPropIdOk() (*string, bool) {
	if o == nil || IsNil(o.PropId) {
		return nil, false
	}
	return o.PropId, true
}

// HasPropId returns a boolean if a field has been set.
func (o *InventoryMetadata) HasPropId() bool {
	if o != nil && !IsNil(o.PropId) {
		return true
	}

	return false
}

// SetPropId gets a reference to the given string and assigns it to the PropId field.
func (o *InventoryMetadata) SetPropId(v string) {
	o.PropId = &v
}

func (o InventoryMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InventoryMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Animated) {
		toSerialize["animated"] = o.Animated
	}
	if !IsNil(o.AnimationStyle) {
		toSerialize["animationStyle"] = o.AnimationStyle
	}
	if !IsNil(o.AssetBundleId) {
		toSerialize["assetBundleId"] = o.AssetBundleId
	}
	if !IsNil(o.FileId) {
		toSerialize["fileId"] = o.FileId
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	if !IsNil(o.InventoryItemsToInstantiate) {
		toSerialize["inventoryItemsToInstantiate"] = o.InventoryItemsToInstantiate
	}
	if !IsNil(o.MaskTag) {
		toSerialize["maskTag"] = o.MaskTag
	}
	if !IsNil(o.PropId) {
		toSerialize["propId"] = o.PropId
	}
	return toSerialize, nil
}

type NullableInventoryMetadata struct {
	value *InventoryMetadata
	isSet bool
}

func (v NullableInventoryMetadata) Get() *InventoryMetadata {
	return v.value
}

func (v *NullableInventoryMetadata) Set(val *InventoryMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryMetadata(val *InventoryMetadata) *NullableInventoryMetadata {
	return &NullableInventoryMetadata{value: val, isSet: true}
}

func (v NullableInventoryMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


