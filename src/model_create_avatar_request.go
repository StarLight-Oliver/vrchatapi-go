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
)

// checks if the CreateAvatarRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAvatarRequest{}

// CreateAvatarRequest struct for CreateAvatarRequest
type CreateAvatarRequest struct {
	AssetUrl     *string `json:"assetUrl,omitempty"`
	AssetVersion *string `json:"assetVersion,omitempty"`
	// A date and time of the pattern `M/d/yyyy h:mm:ss tt` (see C Sharp `System.DateTime`)
	CreatedAt   *string `json:"created_at,omitempty"`
	Description *string `json:"description,omitempty"`
	Id          *string `json:"id,omitempty"`
	ImageUrl    string  `json:"imageUrl"`
	Name        string  `json:"name"`
	// This is normally `android`, `ios`, `standalonewindows`, `web`, or the empty value ``, but also supposedly can be any random Unity version such as `2019.2.4-801-Release` or `2019.2.2-772-Release` or even `unknownplatform`.
	Platform      *string        `json:"platform,omitempty"`
	ReleaseStatus *ReleaseStatus `json:"releaseStatus,omitempty"`
	//
	Tags              []string `json:"tags,omitempty"`
	ThumbnailImageUrl *string  `json:"thumbnailImageUrl,omitempty"`
	UnityPackageUrl   *string  `json:"unityPackageUrl,omitempty"`
	UnityVersion      *string  `json:"unityVersion,omitempty"`
	// A date and time of the pattern `M/d/yyyy h:mm:ss tt` (see C Sharp `System.DateTime`)
	UpdatedAt *string `json:"updated_at,omitempty"`
	Version   *int32  `json:"version,omitempty"`
}

type _CreateAvatarRequest CreateAvatarRequest

// NewCreateAvatarRequest instantiates a new CreateAvatarRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAvatarRequest(imageUrl string, name string) *CreateAvatarRequest {
	this := CreateAvatarRequest{}
	this.ImageUrl = imageUrl
	this.Name = name
	var releaseStatus ReleaseStatus = ReleaseStatus_PUBLIC
	this.ReleaseStatus = &releaseStatus
	var unityVersion string = "5.3.4p1"
	this.UnityVersion = &unityVersion
	var version int32 = 1
	this.Version = &version
	return &this
}

// NewCreateAvatarRequestWithDefaults instantiates a new CreateAvatarRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAvatarRequestWithDefaults() *CreateAvatarRequest {
	this := CreateAvatarRequest{}
	var releaseStatus ReleaseStatus = ReleaseStatus_PUBLIC
	this.ReleaseStatus = &releaseStatus
	var unityVersion string = "5.3.4p1"
	this.UnityVersion = &unityVersion
	var version int32 = 1
	this.Version = &version
	return &this
}

// GetAssetUrl returns the AssetUrl field value if set, zero value otherwise.
func (o *CreateAvatarRequest) GetAssetUrl() string {
	if o == nil || IsNil(o.AssetUrl) {
		var ret string
		return ret
	}
	return *o.AssetUrl
}

// GetAssetUrlOk returns a tuple with the AssetUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAvatarRequest) GetAssetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AssetUrl) {
		return nil, false
	}
	return o.AssetUrl, true
}

// HasAssetUrl returns a boolean if a field has been set.
func (o *CreateAvatarRequest) HasAssetUrl() bool {
	if o != nil && !IsNil(o.AssetUrl) {
		return true
	}

	return false
}

// SetAssetUrl gets a reference to the given string and assigns it to the AssetUrl field.
func (o *CreateAvatarRequest) SetAssetUrl(v string) {
	o.AssetUrl = &v
}

// GetAssetVersion returns the AssetVersion field value if set, zero value otherwise.
func (o *CreateAvatarRequest) GetAssetVersion() string {
	if o == nil || IsNil(o.AssetVersion) {
		var ret string
		return ret
	}
	return *o.AssetVersion
}

// GetAssetVersionOk returns a tuple with the AssetVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAvatarRequest) GetAssetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.AssetVersion) {
		return nil, false
	}
	return o.AssetVersion, true
}

// HasAssetVersion returns a boolean if a field has been set.
func (o *CreateAvatarRequest) HasAssetVersion() bool {
	if o != nil && !IsNil(o.AssetVersion) {
		return true
	}

	return false
}

// SetAssetVersion gets a reference to the given string and assigns it to the AssetVersion field.
func (o *CreateAvatarRequest) SetAssetVersion(v string) {
	o.AssetVersion = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CreateAvatarRequest) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAvatarRequest) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CreateAvatarRequest) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *CreateAvatarRequest) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateAvatarRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAvatarRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateAvatarRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateAvatarRequest) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateAvatarRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAvatarRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateAvatarRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateAvatarRequest) SetId(v string) {
	o.Id = &v
}

// GetImageUrl returns the ImageUrl field value
func (o *CreateAvatarRequest) GetImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
func (o *CreateAvatarRequest) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageUrl, true
}

// SetImageUrl sets field value
func (o *CreateAvatarRequest) SetImageUrl(v string) {
	o.ImageUrl = v
}

// GetName returns the Name field value
func (o *CreateAvatarRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateAvatarRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateAvatarRequest) SetName(v string) {
	o.Name = v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *CreateAvatarRequest) GetPlatform() string {
	if o == nil || IsNil(o.Platform) {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAvatarRequest) GetPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *CreateAvatarRequest) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *CreateAvatarRequest) SetPlatform(v string) {
	o.Platform = &v
}

// GetReleaseStatus returns the ReleaseStatus field value if set, zero value otherwise.
func (o *CreateAvatarRequest) GetReleaseStatus() ReleaseStatus {
	if o == nil || IsNil(o.ReleaseStatus) {
		var ret ReleaseStatus
		return ret
	}
	return *o.ReleaseStatus
}

// GetReleaseStatusOk returns a tuple with the ReleaseStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAvatarRequest) GetReleaseStatusOk() (*ReleaseStatus, bool) {
	if o == nil || IsNil(o.ReleaseStatus) {
		return nil, false
	}
	return o.ReleaseStatus, true
}

// HasReleaseStatus returns a boolean if a field has been set.
func (o *CreateAvatarRequest) HasReleaseStatus() bool {
	if o != nil && !IsNil(o.ReleaseStatus) {
		return true
	}

	return false
}

// SetReleaseStatus gets a reference to the given ReleaseStatus and assigns it to the ReleaseStatus field.
func (o *CreateAvatarRequest) SetReleaseStatus(v ReleaseStatus) {
	o.ReleaseStatus = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateAvatarRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAvatarRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateAvatarRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreateAvatarRequest) SetTags(v []string) {
	o.Tags = v
}

// GetThumbnailImageUrl returns the ThumbnailImageUrl field value if set, zero value otherwise.
func (o *CreateAvatarRequest) GetThumbnailImageUrl() string {
	if o == nil || IsNil(o.ThumbnailImageUrl) {
		var ret string
		return ret
	}
	return *o.ThumbnailImageUrl
}

// GetThumbnailImageUrlOk returns a tuple with the ThumbnailImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAvatarRequest) GetThumbnailImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ThumbnailImageUrl) {
		return nil, false
	}
	return o.ThumbnailImageUrl, true
}

// HasThumbnailImageUrl returns a boolean if a field has been set.
func (o *CreateAvatarRequest) HasThumbnailImageUrl() bool {
	if o != nil && !IsNil(o.ThumbnailImageUrl) {
		return true
	}

	return false
}

// SetThumbnailImageUrl gets a reference to the given string and assigns it to the ThumbnailImageUrl field.
func (o *CreateAvatarRequest) SetThumbnailImageUrl(v string) {
	o.ThumbnailImageUrl = &v
}

// GetUnityPackageUrl returns the UnityPackageUrl field value if set, zero value otherwise.
func (o *CreateAvatarRequest) GetUnityPackageUrl() string {
	if o == nil || IsNil(o.UnityPackageUrl) {
		var ret string
		return ret
	}
	return *o.UnityPackageUrl
}

// GetUnityPackageUrlOk returns a tuple with the UnityPackageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAvatarRequest) GetUnityPackageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.UnityPackageUrl) {
		return nil, false
	}
	return o.UnityPackageUrl, true
}

// HasUnityPackageUrl returns a boolean if a field has been set.
func (o *CreateAvatarRequest) HasUnityPackageUrl() bool {
	if o != nil && !IsNil(o.UnityPackageUrl) {
		return true
	}

	return false
}

// SetUnityPackageUrl gets a reference to the given string and assigns it to the UnityPackageUrl field.
func (o *CreateAvatarRequest) SetUnityPackageUrl(v string) {
	o.UnityPackageUrl = &v
}

// GetUnityVersion returns the UnityVersion field value if set, zero value otherwise.
func (o *CreateAvatarRequest) GetUnityVersion() string {
	if o == nil || IsNil(o.UnityVersion) {
		var ret string
		return ret
	}
	return *o.UnityVersion
}

// GetUnityVersionOk returns a tuple with the UnityVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAvatarRequest) GetUnityVersionOk() (*string, bool) {
	if o == nil || IsNil(o.UnityVersion) {
		return nil, false
	}
	return o.UnityVersion, true
}

// HasUnityVersion returns a boolean if a field has been set.
func (o *CreateAvatarRequest) HasUnityVersion() bool {
	if o != nil && !IsNil(o.UnityVersion) {
		return true
	}

	return false
}

// SetUnityVersion gets a reference to the given string and assigns it to the UnityVersion field.
func (o *CreateAvatarRequest) SetUnityVersion(v string) {
	o.UnityVersion = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CreateAvatarRequest) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAvatarRequest) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CreateAvatarRequest) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *CreateAvatarRequest) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CreateAvatarRequest) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAvatarRequest) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CreateAvatarRequest) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *CreateAvatarRequest) SetVersion(v int32) {
	o.Version = &v
}

func (o CreateAvatarRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAvatarRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssetUrl) {
		toSerialize["assetUrl"] = o.AssetUrl
	}
	if !IsNil(o.AssetVersion) {
		toSerialize["assetVersion"] = o.AssetVersion
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["imageUrl"] = o.ImageUrl
	toSerialize["name"] = o.Name
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.ReleaseStatus) {
		toSerialize["releaseStatus"] = o.ReleaseStatus
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.ThumbnailImageUrl) {
		toSerialize["thumbnailImageUrl"] = o.ThumbnailImageUrl
	}
	if !IsNil(o.UnityPackageUrl) {
		toSerialize["unityPackageUrl"] = o.UnityPackageUrl
	}
	if !IsNil(o.UnityVersion) {
		toSerialize["unityVersion"] = o.UnityVersion
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

func (o *CreateAvatarRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"imageUrl",
		"name",
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

	varCreateAvatarRequest := _CreateAvatarRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateAvatarRequest)

	if err != nil {
		return err
	}

	*o = CreateAvatarRequest(varCreateAvatarRequest)

	return err
}

type NullableCreateAvatarRequest struct {
	value *CreateAvatarRequest
	isSet bool
}

func (v NullableCreateAvatarRequest) Get() *CreateAvatarRequest {
	return v.value
}

func (v *NullableCreateAvatarRequest) Set(val *CreateAvatarRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAvatarRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAvatarRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAvatarRequest(val *CreateAvatarRequest) *NullableCreateAvatarRequest {
	return &NullableCreateAvatarRequest{value: val, isSet: true}
}

func (v NullableCreateAvatarRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAvatarRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
