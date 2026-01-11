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

// checks if the File type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &File{}

// File
type File struct {
	AnimationStyle *ImageAnimationStyle `json:"animationStyle,omitempty"`
	Extension      string               `json:"extension"`
	// The number of frames for animated spritesheet images.
	Frames *int32 `json:"frames,omitempty"`
	// The frames per second for animated spritesheet images.
	FramesOverTime            *int32          `json:"framesOverTime,omitempty"`
	Id                        string          `json:"id"`
	LoopStyle                 *ImageLoopStyle `json:"loopStyle,omitempty"`
	MaskTag                   *ImageMask      `json:"maskTag,omitempty"`
	MimeType                  MIMEType        `json:"mimeType"`
	ModifiedThumbnailFileName *string         `json:"modifiedThumbnailFileName,omitempty"`
	//
	Name string `json:"name"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	OwnerId string `json:"ownerId"`
	//
	Tags []string `json:"tags"`
	//
	Versions []FileVersion `json:"versions"`
}

type _File File

// NewFile instantiates a new File object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFile(extension string, id string, mimeType MIMEType, name string, ownerId string, tags []string, versions []FileVersion) *File {
	this := File{}
	this.Extension = extension
	this.Id = id
	var loopStyle ImageLoopStyle = ImageLoopStyle_LINEAR
	this.LoopStyle = &loopStyle
	var maskTag ImageMask = ImageMask_SQUARE
	this.MaskTag = &maskTag
	this.MimeType = mimeType
	this.Name = name
	this.OwnerId = ownerId
	this.Tags = tags
	this.Versions = versions
	return &this
}

// NewFileWithDefaults instantiates a new File object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileWithDefaults() *File {
	this := File{}
	var loopStyle ImageLoopStyle = ImageLoopStyle_LINEAR
	this.LoopStyle = &loopStyle
	var maskTag ImageMask = ImageMask_SQUARE
	this.MaskTag = &maskTag
	var mimeType MIMEType = MIMEType_IMAGE_JPEG
	this.MimeType = mimeType
	return &this
}

// GetAnimationStyle returns the AnimationStyle field value if set, zero value otherwise.
func (o *File) GetAnimationStyle() ImageAnimationStyle {
	if o == nil || IsNil(o.AnimationStyle) {
		var ret ImageAnimationStyle
		return ret
	}
	return *o.AnimationStyle
}

// GetAnimationStyleOk returns a tuple with the AnimationStyle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *File) GetAnimationStyleOk() (*ImageAnimationStyle, bool) {
	if o == nil || IsNil(o.AnimationStyle) {
		return nil, false
	}
	return o.AnimationStyle, true
}

// HasAnimationStyle returns a boolean if a field has been set.
func (o *File) HasAnimationStyle() bool {
	if o != nil && !IsNil(o.AnimationStyle) {
		return true
	}

	return false
}

// SetAnimationStyle gets a reference to the given ImageAnimationStyle and assigns it to the AnimationStyle field.
func (o *File) SetAnimationStyle(v ImageAnimationStyle) {
	o.AnimationStyle = &v
}

// GetExtension returns the Extension field value
func (o *File) GetExtension() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value
// and a boolean to check if the value has been set.
func (o *File) GetExtensionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Extension, true
}

// SetExtension sets field value
func (o *File) SetExtension(v string) {
	o.Extension = v
}

// GetFrames returns the Frames field value if set, zero value otherwise.
func (o *File) GetFrames() int32 {
	if o == nil || IsNil(o.Frames) {
		var ret int32
		return ret
	}
	return *o.Frames
}

// GetFramesOk returns a tuple with the Frames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *File) GetFramesOk() (*int32, bool) {
	if o == nil || IsNil(o.Frames) {
		return nil, false
	}
	return o.Frames, true
}

// HasFrames returns a boolean if a field has been set.
func (o *File) HasFrames() bool {
	if o != nil && !IsNil(o.Frames) {
		return true
	}

	return false
}

// SetFrames gets a reference to the given int32 and assigns it to the Frames field.
func (o *File) SetFrames(v int32) {
	o.Frames = &v
}

// GetFramesOverTime returns the FramesOverTime field value if set, zero value otherwise.
func (o *File) GetFramesOverTime() int32 {
	if o == nil || IsNil(o.FramesOverTime) {
		var ret int32
		return ret
	}
	return *o.FramesOverTime
}

// GetFramesOverTimeOk returns a tuple with the FramesOverTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *File) GetFramesOverTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.FramesOverTime) {
		return nil, false
	}
	return o.FramesOverTime, true
}

// HasFramesOverTime returns a boolean if a field has been set.
func (o *File) HasFramesOverTime() bool {
	if o != nil && !IsNil(o.FramesOverTime) {
		return true
	}

	return false
}

// SetFramesOverTime gets a reference to the given int32 and assigns it to the FramesOverTime field.
func (o *File) SetFramesOverTime(v int32) {
	o.FramesOverTime = &v
}

// GetId returns the Id field value
func (o *File) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *File) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *File) SetId(v string) {
	o.Id = v
}

// GetLoopStyle returns the LoopStyle field value if set, zero value otherwise.
func (o *File) GetLoopStyle() ImageLoopStyle {
	if o == nil || IsNil(o.LoopStyle) {
		var ret ImageLoopStyle
		return ret
	}
	return *o.LoopStyle
}

// GetLoopStyleOk returns a tuple with the LoopStyle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *File) GetLoopStyleOk() (*ImageLoopStyle, bool) {
	if o == nil || IsNil(o.LoopStyle) {
		return nil, false
	}
	return o.LoopStyle, true
}

// HasLoopStyle returns a boolean if a field has been set.
func (o *File) HasLoopStyle() bool {
	if o != nil && !IsNil(o.LoopStyle) {
		return true
	}

	return false
}

// SetLoopStyle gets a reference to the given ImageLoopStyle and assigns it to the LoopStyle field.
func (o *File) SetLoopStyle(v ImageLoopStyle) {
	o.LoopStyle = &v
}

// GetMaskTag returns the MaskTag field value if set, zero value otherwise.
func (o *File) GetMaskTag() ImageMask {
	if o == nil || IsNil(o.MaskTag) {
		var ret ImageMask
		return ret
	}
	return *o.MaskTag
}

// GetMaskTagOk returns a tuple with the MaskTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *File) GetMaskTagOk() (*ImageMask, bool) {
	if o == nil || IsNil(o.MaskTag) {
		return nil, false
	}
	return o.MaskTag, true
}

// HasMaskTag returns a boolean if a field has been set.
func (o *File) HasMaskTag() bool {
	if o != nil && !IsNil(o.MaskTag) {
		return true
	}

	return false
}

// SetMaskTag gets a reference to the given ImageMask and assigns it to the MaskTag field.
func (o *File) SetMaskTag(v ImageMask) {
	o.MaskTag = &v
}

// GetMimeType returns the MimeType field value
func (o *File) GetMimeType() MIMEType {
	if o == nil {
		var ret MIMEType
		return ret
	}

	return o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value
// and a boolean to check if the value has been set.
func (o *File) GetMimeTypeOk() (*MIMEType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MimeType, true
}

// SetMimeType sets field value
func (o *File) SetMimeType(v MIMEType) {
	o.MimeType = v
}

// GetModifiedThumbnailFileName returns the ModifiedThumbnailFileName field value if set, zero value otherwise.
func (o *File) GetModifiedThumbnailFileName() string {
	if o == nil || IsNil(o.ModifiedThumbnailFileName) {
		var ret string
		return ret
	}
	return *o.ModifiedThumbnailFileName
}

// GetModifiedThumbnailFileNameOk returns a tuple with the ModifiedThumbnailFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *File) GetModifiedThumbnailFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.ModifiedThumbnailFileName) {
		return nil, false
	}
	return o.ModifiedThumbnailFileName, true
}

// HasModifiedThumbnailFileName returns a boolean if a field has been set.
func (o *File) HasModifiedThumbnailFileName() bool {
	if o != nil && !IsNil(o.ModifiedThumbnailFileName) {
		return true
	}

	return false
}

// SetModifiedThumbnailFileName gets a reference to the given string and assigns it to the ModifiedThumbnailFileName field.
func (o *File) SetModifiedThumbnailFileName(v string) {
	o.ModifiedThumbnailFileName = &v
}

// GetName returns the Name field value
func (o *File) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *File) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *File) SetName(v string) {
	o.Name = v
}

// GetOwnerId returns the OwnerId field value
func (o *File) GetOwnerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value
// and a boolean to check if the value has been set.
func (o *File) GetOwnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerId, true
}

// SetOwnerId sets field value
func (o *File) SetOwnerId(v string) {
	o.OwnerId = v
}

// GetTags returns the Tags field value
func (o *File) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *File) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *File) SetTags(v []string) {
	o.Tags = v
}

// GetVersions returns the Versions field value
func (o *File) GetVersions() []FileVersion {
	if o == nil {
		var ret []FileVersion
		return ret
	}

	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value
// and a boolean to check if the value has been set.
func (o *File) GetVersionsOk() ([]FileVersion, bool) {
	if o == nil {
		return nil, false
	}
	return o.Versions, true
}

// SetVersions sets field value
func (o *File) SetVersions(v []FileVersion) {
	o.Versions = v
}

func (o File) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o File) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AnimationStyle) {
		toSerialize["animationStyle"] = o.AnimationStyle
	}
	toSerialize["extension"] = o.Extension
	if !IsNil(o.Frames) {
		toSerialize["frames"] = o.Frames
	}
	if !IsNil(o.FramesOverTime) {
		toSerialize["framesOverTime"] = o.FramesOverTime
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.LoopStyle) {
		toSerialize["loopStyle"] = o.LoopStyle
	}
	if !IsNil(o.MaskTag) {
		toSerialize["maskTag"] = o.MaskTag
	}
	toSerialize["mimeType"] = o.MimeType
	if !IsNil(o.ModifiedThumbnailFileName) {
		toSerialize["modifiedThumbnailFileName"] = o.ModifiedThumbnailFileName
	}
	toSerialize["name"] = o.Name
	toSerialize["ownerId"] = o.OwnerId
	toSerialize["tags"] = o.Tags
	toSerialize["versions"] = o.Versions
	return toSerialize, nil
}

func (o *File) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"extension",
		"id",
		"mimeType",
		"name",
		"ownerId",
		"tags",
		"versions",
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

	varFile := _File{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFile)

	if err != nil {
		return err
	}

	*o = File(varFile)

	return err
}

type NullableFile struct {
	value *File
	isSet bool
}

func (v NullableFile) Get() *File {
	return v.value
}

func (v *NullableFile) Set(val *File) {
	v.value = val
	v.isSet = true
}

func (v NullableFile) IsSet() bool {
	return v.isSet
}

func (v *NullableFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFile(val *File) *NullableFile {
	return &NullableFile{value: val, isSet: true}
}

func (v NullableFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
