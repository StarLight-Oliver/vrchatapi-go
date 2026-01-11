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

// checks if the FileAnalysis type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileAnalysis{}

// FileAnalysis struct for FileAnalysis
type FileAnalysis struct {
	AvatarStats FileAnalysisAvatarStats `json:"avatarStats"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	EncryptionKey *string `json:"encryptionKey,omitempty"`
	FileSize int32 `json:"fileSize"`
	PerformanceRating *string `json:"performanceRating,omitempty"`
	Success bool `json:"success"`
	UncompressedSize int32 `json:"uncompressedSize"`
}

type _FileAnalysis FileAnalysis

// NewFileAnalysis instantiates a new FileAnalysis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileAnalysis(avatarStats FileAnalysisAvatarStats, fileSize int32, success bool, uncompressedSize int32) *FileAnalysis {
	this := FileAnalysis{}
	this.AvatarStats = avatarStats
	this.FileSize = fileSize
	this.Success = success
	this.UncompressedSize = uncompressedSize
	return &this
}

// NewFileAnalysisWithDefaults instantiates a new FileAnalysis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileAnalysisWithDefaults() *FileAnalysis {
	this := FileAnalysis{}
	return &this
}

// GetAvatarStats returns the AvatarStats field value
func (o *FileAnalysis) GetAvatarStats() FileAnalysisAvatarStats {
	if o == nil {
		var ret FileAnalysisAvatarStats
		return ret
	}

	return o.AvatarStats
}

// GetAvatarStatsOk returns a tuple with the AvatarStats field value
// and a boolean to check if the value has been set.
func (o *FileAnalysis) GetAvatarStatsOk() (*FileAnalysisAvatarStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvatarStats, true
}

// SetAvatarStats sets field value
func (o *FileAnalysis) SetAvatarStats(v FileAnalysisAvatarStats) {
	o.AvatarStats = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FileAnalysis) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileAnalysis) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FileAnalysis) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *FileAnalysis) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetEncryptionKey returns the EncryptionKey field value if set, zero value otherwise.
func (o *FileAnalysis) GetEncryptionKey() string {
	if o == nil || IsNil(o.EncryptionKey) {
		var ret string
		return ret
	}
	return *o.EncryptionKey
}

// GetEncryptionKeyOk returns a tuple with the EncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileAnalysis) GetEncryptionKeyOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionKey) {
		return nil, false
	}
	return o.EncryptionKey, true
}

// HasEncryptionKey returns a boolean if a field has been set.
func (o *FileAnalysis) HasEncryptionKey() bool {
	if o != nil && !IsNil(o.EncryptionKey) {
		return true
	}

	return false
}

// SetEncryptionKey gets a reference to the given string and assigns it to the EncryptionKey field.
func (o *FileAnalysis) SetEncryptionKey(v string) {
	o.EncryptionKey = &v
}

// GetFileSize returns the FileSize field value
func (o *FileAnalysis) GetFileSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value
// and a boolean to check if the value has been set.
func (o *FileAnalysis) GetFileSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileSize, true
}

// SetFileSize sets field value
func (o *FileAnalysis) SetFileSize(v int32) {
	o.FileSize = v
}

// GetPerformanceRating returns the PerformanceRating field value if set, zero value otherwise.
func (o *FileAnalysis) GetPerformanceRating() string {
	if o == nil || IsNil(o.PerformanceRating) {
		var ret string
		return ret
	}
	return *o.PerformanceRating
}

// GetPerformanceRatingOk returns a tuple with the PerformanceRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileAnalysis) GetPerformanceRatingOk() (*string, bool) {
	if o == nil || IsNil(o.PerformanceRating) {
		return nil, false
	}
	return o.PerformanceRating, true
}

// HasPerformanceRating returns a boolean if a field has been set.
func (o *FileAnalysis) HasPerformanceRating() bool {
	if o != nil && !IsNil(o.PerformanceRating) {
		return true
	}

	return false
}

// SetPerformanceRating gets a reference to the given string and assigns it to the PerformanceRating field.
func (o *FileAnalysis) SetPerformanceRating(v string) {
	o.PerformanceRating = &v
}

// GetSuccess returns the Success field value
func (o *FileAnalysis) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *FileAnalysis) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *FileAnalysis) SetSuccess(v bool) {
	o.Success = v
}

// GetUncompressedSize returns the UncompressedSize field value
func (o *FileAnalysis) GetUncompressedSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UncompressedSize
}

// GetUncompressedSizeOk returns a tuple with the UncompressedSize field value
// and a boolean to check if the value has been set.
func (o *FileAnalysis) GetUncompressedSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UncompressedSize, true
}

// SetUncompressedSize sets field value
func (o *FileAnalysis) SetUncompressedSize(v int32) {
	o.UncompressedSize = v
}

func (o FileAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileAnalysis) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["avatarStats"] = o.AvatarStats
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.EncryptionKey) {
		toSerialize["encryptionKey"] = o.EncryptionKey
	}
	toSerialize["fileSize"] = o.FileSize
	if !IsNil(o.PerformanceRating) {
		toSerialize["performanceRating"] = o.PerformanceRating
	}
	toSerialize["success"] = o.Success
	toSerialize["uncompressedSize"] = o.UncompressedSize
	return toSerialize, nil
}

func (o *FileAnalysis) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"avatarStats",
		"fileSize",
		"success",
		"uncompressedSize",
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

	varFileAnalysis := _FileAnalysis{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFileAnalysis)

	if err != nil {
		return err
	}

	*o = FileAnalysis(varFileAnalysis)

	return err
}

type NullableFileAnalysis struct {
	value *FileAnalysis
	isSet bool
}

func (v NullableFileAnalysis) Get() *FileAnalysis {
	return v.value
}

func (v *NullableFileAnalysis) Set(val *FileAnalysis) {
	v.value = val
	v.isSet = true
}

func (v NullableFileAnalysis) IsSet() bool {
	return v.isSet
}

func (v *NullableFileAnalysis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileAnalysis(val *FileAnalysis) *NullableFileAnalysis {
	return &NullableFileAnalysis{value: val, isSet: true}
}

func (v NullableFileAnalysis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileAnalysis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


