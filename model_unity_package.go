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
	"bytes"
	"fmt"
)

// checks if the UnityPackage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnityPackage{}

// UnityPackage 
type UnityPackage struct {
	AssetUrl NullableString `json:"assetUrl,omitempty"`
	AssetUrlObject map[string]interface{} `json:"assetUrlObject,omitempty"`
	AssetVersion int32 `json:"assetVersion"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Id string `json:"id"`
	ImpostorUrl NullableString `json:"impostorUrl,omitempty"`
	ImpostorizerVersion *string `json:"impostorizerVersion,omitempty"`
	PerformanceRating *PerformanceRatings `json:"performanceRating,omitempty"`
	// This is normally `android`, `ios`, `standalonewindows`, `web`, or the empty value ``, but also supposedly can be any random Unity version such as `2019.2.4-801-Release` or `2019.2.2-772-Release` or even `unknownplatform`.
	Platform string `json:"platform"`
	PluginUrl *string `json:"pluginUrl,omitempty"`
	PluginUrlObject map[string]interface{} `json:"pluginUrlObject,omitempty"`
	ScanStatus *string `json:"scanStatus,omitempty"`
	UnitySortNumber *int64 `json:"unitySortNumber,omitempty"`
	UnityVersion string `json:"unityVersion"`
	Variant *string `json:"variant,omitempty"`
	WorldSignature NullableString `json:"worldSignature,omitempty"`
}

type _UnityPackage UnityPackage

// NewUnityPackage instantiates a new UnityPackage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnityPackage(assetVersion int32, id string, platform string, unityVersion string) *UnityPackage {
	this := UnityPackage{}
	this.AssetVersion = assetVersion
	this.Id = id
	this.Platform = platform
	this.UnityVersion = unityVersion
	return &this
}

// NewUnityPackageWithDefaults instantiates a new UnityPackage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnityPackageWithDefaults() *UnityPackage {
	this := UnityPackage{}
	var unityVersion string = "5.3.4p1"
	this.UnityVersion = unityVersion
	return &this
}

// GetAssetUrl returns the AssetUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnityPackage) GetAssetUrl() string {
	if o == nil || IsNil(o.AssetUrl.Get()) {
		var ret string
		return ret
	}
	return *o.AssetUrl.Get()
}

// GetAssetUrlOk returns a tuple with the AssetUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnityPackage) GetAssetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetUrl.Get(), o.AssetUrl.IsSet()
}

// HasAssetUrl returns a boolean if a field has been set.
func (o *UnityPackage) HasAssetUrl() bool {
	if o != nil && o.AssetUrl.IsSet() {
		return true
	}

	return false
}

// SetAssetUrl gets a reference to the given NullableString and assigns it to the AssetUrl field.
func (o *UnityPackage) SetAssetUrl(v string) {
	o.AssetUrl.Set(&v)
}
// SetAssetUrlNil sets the value for AssetUrl to be an explicit nil
func (o *UnityPackage) SetAssetUrlNil() {
	o.AssetUrl.Set(nil)
}

// UnsetAssetUrl ensures that no value is present for AssetUrl, not even an explicit nil
func (o *UnityPackage) UnsetAssetUrl() {
	o.AssetUrl.Unset()
}

// GetAssetUrlObject returns the AssetUrlObject field value if set, zero value otherwise.
func (o *UnityPackage) GetAssetUrlObject() map[string]interface{} {
	if o == nil || IsNil(o.AssetUrlObject) {
		var ret map[string]interface{}
		return ret
	}
	return o.AssetUrlObject
}

// GetAssetUrlObjectOk returns a tuple with the AssetUrlObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnityPackage) GetAssetUrlObjectOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AssetUrlObject) {
		return map[string]interface{}{}, false
	}
	return o.AssetUrlObject, true
}

// HasAssetUrlObject returns a boolean if a field has been set.
func (o *UnityPackage) HasAssetUrlObject() bool {
	if o != nil && !IsNil(o.AssetUrlObject) {
		return true
	}

	return false
}

// SetAssetUrlObject gets a reference to the given map[string]interface{} and assigns it to the AssetUrlObject field.
func (o *UnityPackage) SetAssetUrlObject(v map[string]interface{}) {
	o.AssetUrlObject = v
}

// GetAssetVersion returns the AssetVersion field value
func (o *UnityPackage) GetAssetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AssetVersion
}

// GetAssetVersionOk returns a tuple with the AssetVersion field value
// and a boolean to check if the value has been set.
func (o *UnityPackage) GetAssetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetVersion, true
}

// SetAssetVersion sets field value
func (o *UnityPackage) SetAssetVersion(v int32) {
	o.AssetVersion = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *UnityPackage) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnityPackage) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *UnityPackage) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *UnityPackage) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetId returns the Id field value
func (o *UnityPackage) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UnityPackage) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UnityPackage) SetId(v string) {
	o.Id = v
}

// GetImpostorUrl returns the ImpostorUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnityPackage) GetImpostorUrl() string {
	if o == nil || IsNil(o.ImpostorUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ImpostorUrl.Get()
}

// GetImpostorUrlOk returns a tuple with the ImpostorUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnityPackage) GetImpostorUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImpostorUrl.Get(), o.ImpostorUrl.IsSet()
}

// HasImpostorUrl returns a boolean if a field has been set.
func (o *UnityPackage) HasImpostorUrl() bool {
	if o != nil && o.ImpostorUrl.IsSet() {
		return true
	}

	return false
}

// SetImpostorUrl gets a reference to the given NullableString and assigns it to the ImpostorUrl field.
func (o *UnityPackage) SetImpostorUrl(v string) {
	o.ImpostorUrl.Set(&v)
}
// SetImpostorUrlNil sets the value for ImpostorUrl to be an explicit nil
func (o *UnityPackage) SetImpostorUrlNil() {
	o.ImpostorUrl.Set(nil)
}

// UnsetImpostorUrl ensures that no value is present for ImpostorUrl, not even an explicit nil
func (o *UnityPackage) UnsetImpostorUrl() {
	o.ImpostorUrl.Unset()
}

// GetImpostorizerVersion returns the ImpostorizerVersion field value if set, zero value otherwise.
func (o *UnityPackage) GetImpostorizerVersion() string {
	if o == nil || IsNil(o.ImpostorizerVersion) {
		var ret string
		return ret
	}
	return *o.ImpostorizerVersion
}

// GetImpostorizerVersionOk returns a tuple with the ImpostorizerVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnityPackage) GetImpostorizerVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ImpostorizerVersion) {
		return nil, false
	}
	return o.ImpostorizerVersion, true
}

// HasImpostorizerVersion returns a boolean if a field has been set.
func (o *UnityPackage) HasImpostorizerVersion() bool {
	if o != nil && !IsNil(o.ImpostorizerVersion) {
		return true
	}

	return false
}

// SetImpostorizerVersion gets a reference to the given string and assigns it to the ImpostorizerVersion field.
func (o *UnityPackage) SetImpostorizerVersion(v string) {
	o.ImpostorizerVersion = &v
}

// GetPerformanceRating returns the PerformanceRating field value if set, zero value otherwise.
func (o *UnityPackage) GetPerformanceRating() PerformanceRatings {
	if o == nil || IsNil(o.PerformanceRating) {
		var ret PerformanceRatings
		return ret
	}
	return *o.PerformanceRating
}

// GetPerformanceRatingOk returns a tuple with the PerformanceRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnityPackage) GetPerformanceRatingOk() (*PerformanceRatings, bool) {
	if o == nil || IsNil(o.PerformanceRating) {
		return nil, false
	}
	return o.PerformanceRating, true
}

// HasPerformanceRating returns a boolean if a field has been set.
func (o *UnityPackage) HasPerformanceRating() bool {
	if o != nil && !IsNil(o.PerformanceRating) {
		return true
	}

	return false
}

// SetPerformanceRating gets a reference to the given PerformanceRatings and assigns it to the PerformanceRating field.
func (o *UnityPackage) SetPerformanceRating(v PerformanceRatings) {
	o.PerformanceRating = &v
}

// GetPlatform returns the Platform field value
func (o *UnityPackage) GetPlatform() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value
// and a boolean to check if the value has been set.
func (o *UnityPackage) GetPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Platform, true
}

// SetPlatform sets field value
func (o *UnityPackage) SetPlatform(v string) {
	o.Platform = v
}

// GetPluginUrl returns the PluginUrl field value if set, zero value otherwise.
func (o *UnityPackage) GetPluginUrl() string {
	if o == nil || IsNil(o.PluginUrl) {
		var ret string
		return ret
	}
	return *o.PluginUrl
}

// GetPluginUrlOk returns a tuple with the PluginUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnityPackage) GetPluginUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PluginUrl) {
		return nil, false
	}
	return o.PluginUrl, true
}

// HasPluginUrl returns a boolean if a field has been set.
func (o *UnityPackage) HasPluginUrl() bool {
	if o != nil && !IsNil(o.PluginUrl) {
		return true
	}

	return false
}

// SetPluginUrl gets a reference to the given string and assigns it to the PluginUrl field.
func (o *UnityPackage) SetPluginUrl(v string) {
	o.PluginUrl = &v
}

// GetPluginUrlObject returns the PluginUrlObject field value if set, zero value otherwise.
func (o *UnityPackage) GetPluginUrlObject() map[string]interface{} {
	if o == nil || IsNil(o.PluginUrlObject) {
		var ret map[string]interface{}
		return ret
	}
	return o.PluginUrlObject
}

// GetPluginUrlObjectOk returns a tuple with the PluginUrlObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnityPackage) GetPluginUrlObjectOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.PluginUrlObject) {
		return map[string]interface{}{}, false
	}
	return o.PluginUrlObject, true
}

// HasPluginUrlObject returns a boolean if a field has been set.
func (o *UnityPackage) HasPluginUrlObject() bool {
	if o != nil && !IsNil(o.PluginUrlObject) {
		return true
	}

	return false
}

// SetPluginUrlObject gets a reference to the given map[string]interface{} and assigns it to the PluginUrlObject field.
func (o *UnityPackage) SetPluginUrlObject(v map[string]interface{}) {
	o.PluginUrlObject = v
}

// GetScanStatus returns the ScanStatus field value if set, zero value otherwise.
func (o *UnityPackage) GetScanStatus() string {
	if o == nil || IsNil(o.ScanStatus) {
		var ret string
		return ret
	}
	return *o.ScanStatus
}

// GetScanStatusOk returns a tuple with the ScanStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnityPackage) GetScanStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ScanStatus) {
		return nil, false
	}
	return o.ScanStatus, true
}

// HasScanStatus returns a boolean if a field has been set.
func (o *UnityPackage) HasScanStatus() bool {
	if o != nil && !IsNil(o.ScanStatus) {
		return true
	}

	return false
}

// SetScanStatus gets a reference to the given string and assigns it to the ScanStatus field.
func (o *UnityPackage) SetScanStatus(v string) {
	o.ScanStatus = &v
}

// GetUnitySortNumber returns the UnitySortNumber field value if set, zero value otherwise.
func (o *UnityPackage) GetUnitySortNumber() int64 {
	if o == nil || IsNil(o.UnitySortNumber) {
		var ret int64
		return ret
	}
	return *o.UnitySortNumber
}

// GetUnitySortNumberOk returns a tuple with the UnitySortNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnityPackage) GetUnitySortNumberOk() (*int64, bool) {
	if o == nil || IsNil(o.UnitySortNumber) {
		return nil, false
	}
	return o.UnitySortNumber, true
}

// HasUnitySortNumber returns a boolean if a field has been set.
func (o *UnityPackage) HasUnitySortNumber() bool {
	if o != nil && !IsNil(o.UnitySortNumber) {
		return true
	}

	return false
}

// SetUnitySortNumber gets a reference to the given int64 and assigns it to the UnitySortNumber field.
func (o *UnityPackage) SetUnitySortNumber(v int64) {
	o.UnitySortNumber = &v
}

// GetUnityVersion returns the UnityVersion field value
func (o *UnityPackage) GetUnityVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnityVersion
}

// GetUnityVersionOk returns a tuple with the UnityVersion field value
// and a boolean to check if the value has been set.
func (o *UnityPackage) GetUnityVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnityVersion, true
}

// SetUnityVersion sets field value
func (o *UnityPackage) SetUnityVersion(v string) {
	o.UnityVersion = v
}

// GetVariant returns the Variant field value if set, zero value otherwise.
func (o *UnityPackage) GetVariant() string {
	if o == nil || IsNil(o.Variant) {
		var ret string
		return ret
	}
	return *o.Variant
}

// GetVariantOk returns a tuple with the Variant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnityPackage) GetVariantOk() (*string, bool) {
	if o == nil || IsNil(o.Variant) {
		return nil, false
	}
	return o.Variant, true
}

// HasVariant returns a boolean if a field has been set.
func (o *UnityPackage) HasVariant() bool {
	if o != nil && !IsNil(o.Variant) {
		return true
	}

	return false
}

// SetVariant gets a reference to the given string and assigns it to the Variant field.
func (o *UnityPackage) SetVariant(v string) {
	o.Variant = &v
}

// GetWorldSignature returns the WorldSignature field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnityPackage) GetWorldSignature() string {
	if o == nil || IsNil(o.WorldSignature.Get()) {
		var ret string
		return ret
	}
	return *o.WorldSignature.Get()
}

// GetWorldSignatureOk returns a tuple with the WorldSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnityPackage) GetWorldSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorldSignature.Get(), o.WorldSignature.IsSet()
}

// HasWorldSignature returns a boolean if a field has been set.
func (o *UnityPackage) HasWorldSignature() bool {
	if o != nil && o.WorldSignature.IsSet() {
		return true
	}

	return false
}

// SetWorldSignature gets a reference to the given NullableString and assigns it to the WorldSignature field.
func (o *UnityPackage) SetWorldSignature(v string) {
	o.WorldSignature.Set(&v)
}
// SetWorldSignatureNil sets the value for WorldSignature to be an explicit nil
func (o *UnityPackage) SetWorldSignatureNil() {
	o.WorldSignature.Set(nil)
}

// UnsetWorldSignature ensures that no value is present for WorldSignature, not even an explicit nil
func (o *UnityPackage) UnsetWorldSignature() {
	o.WorldSignature.Unset()
}

func (o UnityPackage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnityPackage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AssetUrl.IsSet() {
		toSerialize["assetUrl"] = o.AssetUrl.Get()
	}
	if !IsNil(o.AssetUrlObject) {
		toSerialize["assetUrlObject"] = o.AssetUrlObject
	}
	toSerialize["assetVersion"] = o.AssetVersion
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	toSerialize["id"] = o.Id
	if o.ImpostorUrl.IsSet() {
		toSerialize["impostorUrl"] = o.ImpostorUrl.Get()
	}
	if !IsNil(o.ImpostorizerVersion) {
		toSerialize["impostorizerVersion"] = o.ImpostorizerVersion
	}
	if !IsNil(o.PerformanceRating) {
		toSerialize["performanceRating"] = o.PerformanceRating
	}
	toSerialize["platform"] = o.Platform
	if !IsNil(o.PluginUrl) {
		toSerialize["pluginUrl"] = o.PluginUrl
	}
	if !IsNil(o.PluginUrlObject) {
		toSerialize["pluginUrlObject"] = o.PluginUrlObject
	}
	if !IsNil(o.ScanStatus) {
		toSerialize["scanStatus"] = o.ScanStatus
	}
	if !IsNil(o.UnitySortNumber) {
		toSerialize["unitySortNumber"] = o.UnitySortNumber
	}
	toSerialize["unityVersion"] = o.UnityVersion
	if !IsNil(o.Variant) {
		toSerialize["variant"] = o.Variant
	}
	if o.WorldSignature.IsSet() {
		toSerialize["worldSignature"] = o.WorldSignature.Get()
	}
	return toSerialize, nil
}

func (o *UnityPackage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"assetVersion",
		"id",
		"platform",
		"unityVersion",
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

	varUnityPackage := _UnityPackage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnityPackage)

	if err != nil {
		return err
	}

	*o = UnityPackage(varUnityPackage)

	return err
}

type NullableUnityPackage struct {
	value *UnityPackage
	isSet bool
}

func (v NullableUnityPackage) Get() *UnityPackage {
	return v.value
}

func (v *NullableUnityPackage) Set(val *UnityPackage) {
	v.value = val
	v.isSet = true
}

func (v NullableUnityPackage) IsSet() bool {
	return v.isSet
}

func (v *NullableUnityPackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnityPackage(val *UnityPackage) *NullableUnityPackage {
	return &NullableUnityPackage{value: val, isSet: true}
}

func (v NullableUnityPackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnityPackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


