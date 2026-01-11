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

// checks if the SubmitModerationReportRequestDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubmitModerationReportRequestDetails{}

// SubmitModerationReportRequestDetails Relevant details specific to the type of the report. `fileId` is for the image file attached to an inventory item, such as an emoji. `holderId` is for the user who owns an inventory item, such as a emoji.
type SubmitModerationReportRequestDetails struct {
	FileId *string `json:"fileId,omitempty"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	HolderId *string `json:"holderId,omitempty"`
	// Relevant detail for reports about image content, such as emoji.
	ImageType *string `json:"imageType,omitempty"`
	// Relevant detail for reports taking place from within an instance.
	InstanceAgeGated *bool `json:"instanceAgeGated,omitempty"`
	// Relevant detail for reports taking place from within an instance.
	InstanceType *string `json:"instanceType,omitempty"`
	// Relevant detail for reports about content that might not be tagged properly.
	SuggestedWarnings []ContentFilter `json:"suggestedWarnings,omitempty"`
	// Relevant detail for reports involving another user in the same instance world.
	UserInSameInstance *bool `json:"userInSameInstance,omitempty"`
}

// NewSubmitModerationReportRequestDetails instantiates a new SubmitModerationReportRequestDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitModerationReportRequestDetails() *SubmitModerationReportRequestDetails {
	this := SubmitModerationReportRequestDetails{}
	return &this
}

// NewSubmitModerationReportRequestDetailsWithDefaults instantiates a new SubmitModerationReportRequestDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitModerationReportRequestDetailsWithDefaults() *SubmitModerationReportRequestDetails {
	this := SubmitModerationReportRequestDetails{}
	return &this
}

// GetFileId returns the FileId field value if set, zero value otherwise.
func (o *SubmitModerationReportRequestDetails) GetFileId() string {
	if o == nil || IsNil(o.FileId) {
		var ret string
		return ret
	}
	return *o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitModerationReportRequestDetails) GetFileIdOk() (*string, bool) {
	if o == nil || IsNil(o.FileId) {
		return nil, false
	}
	return o.FileId, true
}

// HasFileId returns a boolean if a field has been set.
func (o *SubmitModerationReportRequestDetails) HasFileId() bool {
	if o != nil && !IsNil(o.FileId) {
		return true
	}

	return false
}

// SetFileId gets a reference to the given string and assigns it to the FileId field.
func (o *SubmitModerationReportRequestDetails) SetFileId(v string) {
	o.FileId = &v
}

// GetHolderId returns the HolderId field value if set, zero value otherwise.
func (o *SubmitModerationReportRequestDetails) GetHolderId() string {
	if o == nil || IsNil(o.HolderId) {
		var ret string
		return ret
	}
	return *o.HolderId
}

// GetHolderIdOk returns a tuple with the HolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitModerationReportRequestDetails) GetHolderIdOk() (*string, bool) {
	if o == nil || IsNil(o.HolderId) {
		return nil, false
	}
	return o.HolderId, true
}

// HasHolderId returns a boolean if a field has been set.
func (o *SubmitModerationReportRequestDetails) HasHolderId() bool {
	if o != nil && !IsNil(o.HolderId) {
		return true
	}

	return false
}

// SetHolderId gets a reference to the given string and assigns it to the HolderId field.
func (o *SubmitModerationReportRequestDetails) SetHolderId(v string) {
	o.HolderId = &v
}

// GetImageType returns the ImageType field value if set, zero value otherwise.
func (o *SubmitModerationReportRequestDetails) GetImageType() string {
	if o == nil || IsNil(o.ImageType) {
		var ret string
		return ret
	}
	return *o.ImageType
}

// GetImageTypeOk returns a tuple with the ImageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitModerationReportRequestDetails) GetImageTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ImageType) {
		return nil, false
	}
	return o.ImageType, true
}

// HasImageType returns a boolean if a field has been set.
func (o *SubmitModerationReportRequestDetails) HasImageType() bool {
	if o != nil && !IsNil(o.ImageType) {
		return true
	}

	return false
}

// SetImageType gets a reference to the given string and assigns it to the ImageType field.
func (o *SubmitModerationReportRequestDetails) SetImageType(v string) {
	o.ImageType = &v
}

// GetInstanceAgeGated returns the InstanceAgeGated field value if set, zero value otherwise.
func (o *SubmitModerationReportRequestDetails) GetInstanceAgeGated() bool {
	if o == nil || IsNil(o.InstanceAgeGated) {
		var ret bool
		return ret
	}
	return *o.InstanceAgeGated
}

// GetInstanceAgeGatedOk returns a tuple with the InstanceAgeGated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitModerationReportRequestDetails) GetInstanceAgeGatedOk() (*bool, bool) {
	if o == nil || IsNil(o.InstanceAgeGated) {
		return nil, false
	}
	return o.InstanceAgeGated, true
}

// HasInstanceAgeGated returns a boolean if a field has been set.
func (o *SubmitModerationReportRequestDetails) HasInstanceAgeGated() bool {
	if o != nil && !IsNil(o.InstanceAgeGated) {
		return true
	}

	return false
}

// SetInstanceAgeGated gets a reference to the given bool and assigns it to the InstanceAgeGated field.
func (o *SubmitModerationReportRequestDetails) SetInstanceAgeGated(v bool) {
	o.InstanceAgeGated = &v
}

// GetInstanceType returns the InstanceType field value if set, zero value otherwise.
func (o *SubmitModerationReportRequestDetails) GetInstanceType() string {
	if o == nil || IsNil(o.InstanceType) {
		var ret string
		return ret
	}
	return *o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitModerationReportRequestDetails) GetInstanceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceType) {
		return nil, false
	}
	return o.InstanceType, true
}

// HasInstanceType returns a boolean if a field has been set.
func (o *SubmitModerationReportRequestDetails) HasInstanceType() bool {
	if o != nil && !IsNil(o.InstanceType) {
		return true
	}

	return false
}

// SetInstanceType gets a reference to the given string and assigns it to the InstanceType field.
func (o *SubmitModerationReportRequestDetails) SetInstanceType(v string) {
	o.InstanceType = &v
}

// GetSuggestedWarnings returns the SuggestedWarnings field value if set, zero value otherwise.
func (o *SubmitModerationReportRequestDetails) GetSuggestedWarnings() []ContentFilter {
	if o == nil || IsNil(o.SuggestedWarnings) {
		var ret []ContentFilter
		return ret
	}
	return o.SuggestedWarnings
}

// GetSuggestedWarningsOk returns a tuple with the SuggestedWarnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitModerationReportRequestDetails) GetSuggestedWarningsOk() ([]ContentFilter, bool) {
	if o == nil || IsNil(o.SuggestedWarnings) {
		return nil, false
	}
	return o.SuggestedWarnings, true
}

// HasSuggestedWarnings returns a boolean if a field has been set.
func (o *SubmitModerationReportRequestDetails) HasSuggestedWarnings() bool {
	if o != nil && !IsNil(o.SuggestedWarnings) {
		return true
	}

	return false
}

// SetSuggestedWarnings gets a reference to the given []ContentFilter and assigns it to the SuggestedWarnings field.
func (o *SubmitModerationReportRequestDetails) SetSuggestedWarnings(v []ContentFilter) {
	o.SuggestedWarnings = v
}

// GetUserInSameInstance returns the UserInSameInstance field value if set, zero value otherwise.
func (o *SubmitModerationReportRequestDetails) GetUserInSameInstance() bool {
	if o == nil || IsNil(o.UserInSameInstance) {
		var ret bool
		return ret
	}
	return *o.UserInSameInstance
}

// GetUserInSameInstanceOk returns a tuple with the UserInSameInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitModerationReportRequestDetails) GetUserInSameInstanceOk() (*bool, bool) {
	if o == nil || IsNil(o.UserInSameInstance) {
		return nil, false
	}
	return o.UserInSameInstance, true
}

// HasUserInSameInstance returns a boolean if a field has been set.
func (o *SubmitModerationReportRequestDetails) HasUserInSameInstance() bool {
	if o != nil && !IsNil(o.UserInSameInstance) {
		return true
	}

	return false
}

// SetUserInSameInstance gets a reference to the given bool and assigns it to the UserInSameInstance field.
func (o *SubmitModerationReportRequestDetails) SetUserInSameInstance(v bool) {
	o.UserInSameInstance = &v
}

func (o SubmitModerationReportRequestDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubmitModerationReportRequestDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileId) {
		toSerialize["fileId"] = o.FileId
	}
	if !IsNil(o.HolderId) {
		toSerialize["holderId"] = o.HolderId
	}
	if !IsNil(o.ImageType) {
		toSerialize["imageType"] = o.ImageType
	}
	if !IsNil(o.InstanceAgeGated) {
		toSerialize["instanceAgeGated"] = o.InstanceAgeGated
	}
	if !IsNil(o.InstanceType) {
		toSerialize["instanceType"] = o.InstanceType
	}
	if !IsNil(o.SuggestedWarnings) {
		toSerialize["suggestedWarnings"] = o.SuggestedWarnings
	}
	if !IsNil(o.UserInSameInstance) {
		toSerialize["userInSameInstance"] = o.UserInSameInstance
	}
	return toSerialize, nil
}

type NullableSubmitModerationReportRequestDetails struct {
	value *SubmitModerationReportRequestDetails
	isSet bool
}

func (v NullableSubmitModerationReportRequestDetails) Get() *SubmitModerationReportRequestDetails {
	return v.value
}

func (v *NullableSubmitModerationReportRequestDetails) Set(val *SubmitModerationReportRequestDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitModerationReportRequestDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitModerationReportRequestDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitModerationReportRequestDetails(val *SubmitModerationReportRequestDetails) *NullableSubmitModerationReportRequestDetails {
	return &NullableSubmitModerationReportRequestDetails{value: val, isSet: true}
}

func (v NullableSubmitModerationReportRequestDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitModerationReportRequestDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


