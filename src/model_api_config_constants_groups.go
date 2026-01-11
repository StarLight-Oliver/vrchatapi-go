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

// checks if the APIConfigConstantsGROUPS type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIConfigConstantsGROUPS{}

// APIConfigConstantsGROUPS Group-related constants
type APIConfigConstantsGROUPS struct {
	// Maximum group capacity
	CAPACITY *int32 `json:"CAPACITY,omitempty"`
	// Requirements for transferring group ownership
	GROUP_TRANSFER_REQUIREMENTS []string `json:"GROUP_TRANSFER_REQUIREMENTS,omitempty"`
	// Maximum number of invite requests
	MAX_INVITES_REQUESTS *int32 `json:"MAX_INVITES_REQUESTS,omitempty"`
	// Maximum number of joined groups
	MAX_JOINED *int32 `json:"MAX_JOINED,omitempty"`
	// Maximum number of joined groups for VRChat Plus members
	MAX_JOINED_PLUS *int32 `json:"MAX_JOINED_PLUS,omitempty"`
	// Maximum number of supported languages
	MAX_LANGUAGES *int32 `json:"MAX_LANGUAGES,omitempty"`
	// Maximum number of group links
	MAX_LINKS *int32 `json:"MAX_LINKS,omitempty"`
	// Maximum number of management roles in a group
	MAX_MANAGEMENT_ROLES *int32 `json:"MAX_MANAGEMENT_ROLES,omitempty"`
	// Maximum number of groups a user can own
	MAX_OWNED *int32 `json:"MAX_OWNED,omitempty"`
	// Maximum number of roles in a group
	MAX_ROLES *int32 `json:"MAX_ROLES,omitempty"`
}

// NewAPIConfigConstantsGROUPS instantiates a new APIConfigConstantsGROUPS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIConfigConstantsGROUPS() *APIConfigConstantsGROUPS {
	this := APIConfigConstantsGROUPS{}
	var cAPACITY int32 = 100000
	this.CAPACITY = &cAPACITY
	var mAXINVITESREQUESTS int32 = 50
	this.MAX_INVITES_REQUESTS = &mAXINVITESREQUESTS
	var mAXJOINED int32 = 100
	this.MAX_JOINED = &mAXJOINED
	var mAXJOINEDPLUS int32 = 200
	this.MAX_JOINED_PLUS = &mAXJOINEDPLUS
	var mAXLANGUAGES int32 = 10
	this.MAX_LANGUAGES = &mAXLANGUAGES
	var mAXLINKS int32 = 3
	this.MAX_LINKS = &mAXLINKS
	var mAXMANAGEMENTROLES int32 = 5
	this.MAX_MANAGEMENT_ROLES = &mAXMANAGEMENTROLES
	var mAXOWNED int32 = 5
	this.MAX_OWNED = &mAXOWNED
	var mAXROLES int32 = 50
	this.MAX_ROLES = &mAXROLES
	return &this
}

// NewAPIConfigConstantsGROUPSWithDefaults instantiates a new APIConfigConstantsGROUPS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIConfigConstantsGROUPSWithDefaults() *APIConfigConstantsGROUPS {
	this := APIConfigConstantsGROUPS{}
	var cAPACITY int32 = 100000
	this.CAPACITY = &cAPACITY
	var mAXINVITESREQUESTS int32 = 50
	this.MAX_INVITES_REQUESTS = &mAXINVITESREQUESTS
	var mAXJOINED int32 = 100
	this.MAX_JOINED = &mAXJOINED
	var mAXJOINEDPLUS int32 = 200
	this.MAX_JOINED_PLUS = &mAXJOINEDPLUS
	var mAXLANGUAGES int32 = 10
	this.MAX_LANGUAGES = &mAXLANGUAGES
	var mAXLINKS int32 = 3
	this.MAX_LINKS = &mAXLINKS
	var mAXMANAGEMENTROLES int32 = 5
	this.MAX_MANAGEMENT_ROLES = &mAXMANAGEMENTROLES
	var mAXOWNED int32 = 5
	this.MAX_OWNED = &mAXOWNED
	var mAXROLES int32 = 50
	this.MAX_ROLES = &mAXROLES
	return &this
}

// GetCAPACITY returns the CAPACITY field value if set, zero value otherwise.
func (o *APIConfigConstantsGROUPS) GetCAPACITY() int32 {
	if o == nil || IsNil(o.CAPACITY) {
		var ret int32
		return ret
	}
	return *o.CAPACITY
}

// GetCAPACITYOk returns a tuple with the CAPACITY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIConfigConstantsGROUPS) GetCAPACITYOk() (*int32, bool) {
	if o == nil || IsNil(o.CAPACITY) {
		return nil, false
	}
	return o.CAPACITY, true
}

// HasCAPACITY returns a boolean if a field has been set.
func (o *APIConfigConstantsGROUPS) HasCAPACITY() bool {
	if o != nil && !IsNil(o.CAPACITY) {
		return true
	}

	return false
}

// SetCAPACITY gets a reference to the given int32 and assigns it to the CAPACITY field.
func (o *APIConfigConstantsGROUPS) SetCAPACITY(v int32) {
	o.CAPACITY = &v
}

// GetGROUP_TRANSFER_REQUIREMENTS returns the GROUP_TRANSFER_REQUIREMENTS field value if set, zero value otherwise.
func (o *APIConfigConstantsGROUPS) GetGROUP_TRANSFER_REQUIREMENTS() []string {
	if o == nil || IsNil(o.GROUP_TRANSFER_REQUIREMENTS) {
		var ret []string
		return ret
	}
	return o.GROUP_TRANSFER_REQUIREMENTS
}

// GetGROUP_TRANSFER_REQUIREMENTSOk returns a tuple with the GROUP_TRANSFER_REQUIREMENTS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIConfigConstantsGROUPS) GetGROUP_TRANSFER_REQUIREMENTSOk() ([]string, bool) {
	if o == nil || IsNil(o.GROUP_TRANSFER_REQUIREMENTS) {
		return nil, false
	}
	return o.GROUP_TRANSFER_REQUIREMENTS, true
}

// HasGROUP_TRANSFER_REQUIREMENTS returns a boolean if a field has been set.
func (o *APIConfigConstantsGROUPS) HasGROUP_TRANSFER_REQUIREMENTS() bool {
	if o != nil && !IsNil(o.GROUP_TRANSFER_REQUIREMENTS) {
		return true
	}

	return false
}

// SetGROUP_TRANSFER_REQUIREMENTS gets a reference to the given []string and assigns it to the GROUP_TRANSFER_REQUIREMENTS field.
func (o *APIConfigConstantsGROUPS) SetGROUP_TRANSFER_REQUIREMENTS(v []string) {
	o.GROUP_TRANSFER_REQUIREMENTS = v
}

// GetMAX_INVITES_REQUESTS returns the MAX_INVITES_REQUESTS field value if set, zero value otherwise.
func (o *APIConfigConstantsGROUPS) GetMAX_INVITES_REQUESTS() int32 {
	if o == nil || IsNil(o.MAX_INVITES_REQUESTS) {
		var ret int32
		return ret
	}
	return *o.MAX_INVITES_REQUESTS
}

// GetMAX_INVITES_REQUESTSOk returns a tuple with the MAX_INVITES_REQUESTS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIConfigConstantsGROUPS) GetMAX_INVITES_REQUESTSOk() (*int32, bool) {
	if o == nil || IsNil(o.MAX_INVITES_REQUESTS) {
		return nil, false
	}
	return o.MAX_INVITES_REQUESTS, true
}

// HasMAX_INVITES_REQUESTS returns a boolean if a field has been set.
func (o *APIConfigConstantsGROUPS) HasMAX_INVITES_REQUESTS() bool {
	if o != nil && !IsNil(o.MAX_INVITES_REQUESTS) {
		return true
	}

	return false
}

// SetMAX_INVITES_REQUESTS gets a reference to the given int32 and assigns it to the MAX_INVITES_REQUESTS field.
func (o *APIConfigConstantsGROUPS) SetMAX_INVITES_REQUESTS(v int32) {
	o.MAX_INVITES_REQUESTS = &v
}

// GetMAX_JOINED returns the MAX_JOINED field value if set, zero value otherwise.
func (o *APIConfigConstantsGROUPS) GetMAX_JOINED() int32 {
	if o == nil || IsNil(o.MAX_JOINED) {
		var ret int32
		return ret
	}
	return *o.MAX_JOINED
}

// GetMAX_JOINEDOk returns a tuple with the MAX_JOINED field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIConfigConstantsGROUPS) GetMAX_JOINEDOk() (*int32, bool) {
	if o == nil || IsNil(o.MAX_JOINED) {
		return nil, false
	}
	return o.MAX_JOINED, true
}

// HasMAX_JOINED returns a boolean if a field has been set.
func (o *APIConfigConstantsGROUPS) HasMAX_JOINED() bool {
	if o != nil && !IsNil(o.MAX_JOINED) {
		return true
	}

	return false
}

// SetMAX_JOINED gets a reference to the given int32 and assigns it to the MAX_JOINED field.
func (o *APIConfigConstantsGROUPS) SetMAX_JOINED(v int32) {
	o.MAX_JOINED = &v
}

// GetMAX_JOINED_PLUS returns the MAX_JOINED_PLUS field value if set, zero value otherwise.
func (o *APIConfigConstantsGROUPS) GetMAX_JOINED_PLUS() int32 {
	if o == nil || IsNil(o.MAX_JOINED_PLUS) {
		var ret int32
		return ret
	}
	return *o.MAX_JOINED_PLUS
}

// GetMAX_JOINED_PLUSOk returns a tuple with the MAX_JOINED_PLUS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIConfigConstantsGROUPS) GetMAX_JOINED_PLUSOk() (*int32, bool) {
	if o == nil || IsNil(o.MAX_JOINED_PLUS) {
		return nil, false
	}
	return o.MAX_JOINED_PLUS, true
}

// HasMAX_JOINED_PLUS returns a boolean if a field has been set.
func (o *APIConfigConstantsGROUPS) HasMAX_JOINED_PLUS() bool {
	if o != nil && !IsNil(o.MAX_JOINED_PLUS) {
		return true
	}

	return false
}

// SetMAX_JOINED_PLUS gets a reference to the given int32 and assigns it to the MAX_JOINED_PLUS field.
func (o *APIConfigConstantsGROUPS) SetMAX_JOINED_PLUS(v int32) {
	o.MAX_JOINED_PLUS = &v
}

// GetMAX_LANGUAGES returns the MAX_LANGUAGES field value if set, zero value otherwise.
func (o *APIConfigConstantsGROUPS) GetMAX_LANGUAGES() int32 {
	if o == nil || IsNil(o.MAX_LANGUAGES) {
		var ret int32
		return ret
	}
	return *o.MAX_LANGUAGES
}

// GetMAX_LANGUAGESOk returns a tuple with the MAX_LANGUAGES field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIConfigConstantsGROUPS) GetMAX_LANGUAGESOk() (*int32, bool) {
	if o == nil || IsNil(o.MAX_LANGUAGES) {
		return nil, false
	}
	return o.MAX_LANGUAGES, true
}

// HasMAX_LANGUAGES returns a boolean if a field has been set.
func (o *APIConfigConstantsGROUPS) HasMAX_LANGUAGES() bool {
	if o != nil && !IsNil(o.MAX_LANGUAGES) {
		return true
	}

	return false
}

// SetMAX_LANGUAGES gets a reference to the given int32 and assigns it to the MAX_LANGUAGES field.
func (o *APIConfigConstantsGROUPS) SetMAX_LANGUAGES(v int32) {
	o.MAX_LANGUAGES = &v
}

// GetMAX_LINKS returns the MAX_LINKS field value if set, zero value otherwise.
func (o *APIConfigConstantsGROUPS) GetMAX_LINKS() int32 {
	if o == nil || IsNil(o.MAX_LINKS) {
		var ret int32
		return ret
	}
	return *o.MAX_LINKS
}

// GetMAX_LINKSOk returns a tuple with the MAX_LINKS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIConfigConstantsGROUPS) GetMAX_LINKSOk() (*int32, bool) {
	if o == nil || IsNil(o.MAX_LINKS) {
		return nil, false
	}
	return o.MAX_LINKS, true
}

// HasMAX_LINKS returns a boolean if a field has been set.
func (o *APIConfigConstantsGROUPS) HasMAX_LINKS() bool {
	if o != nil && !IsNil(o.MAX_LINKS) {
		return true
	}

	return false
}

// SetMAX_LINKS gets a reference to the given int32 and assigns it to the MAX_LINKS field.
func (o *APIConfigConstantsGROUPS) SetMAX_LINKS(v int32) {
	o.MAX_LINKS = &v
}

// GetMAX_MANAGEMENT_ROLES returns the MAX_MANAGEMENT_ROLES field value if set, zero value otherwise.
func (o *APIConfigConstantsGROUPS) GetMAX_MANAGEMENT_ROLES() int32 {
	if o == nil || IsNil(o.MAX_MANAGEMENT_ROLES) {
		var ret int32
		return ret
	}
	return *o.MAX_MANAGEMENT_ROLES
}

// GetMAX_MANAGEMENT_ROLESOk returns a tuple with the MAX_MANAGEMENT_ROLES field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIConfigConstantsGROUPS) GetMAX_MANAGEMENT_ROLESOk() (*int32, bool) {
	if o == nil || IsNil(o.MAX_MANAGEMENT_ROLES) {
		return nil, false
	}
	return o.MAX_MANAGEMENT_ROLES, true
}

// HasMAX_MANAGEMENT_ROLES returns a boolean if a field has been set.
func (o *APIConfigConstantsGROUPS) HasMAX_MANAGEMENT_ROLES() bool {
	if o != nil && !IsNil(o.MAX_MANAGEMENT_ROLES) {
		return true
	}

	return false
}

// SetMAX_MANAGEMENT_ROLES gets a reference to the given int32 and assigns it to the MAX_MANAGEMENT_ROLES field.
func (o *APIConfigConstantsGROUPS) SetMAX_MANAGEMENT_ROLES(v int32) {
	o.MAX_MANAGEMENT_ROLES = &v
}

// GetMAX_OWNED returns the MAX_OWNED field value if set, zero value otherwise.
func (o *APIConfigConstantsGROUPS) GetMAX_OWNED() int32 {
	if o == nil || IsNil(o.MAX_OWNED) {
		var ret int32
		return ret
	}
	return *o.MAX_OWNED
}

// GetMAX_OWNEDOk returns a tuple with the MAX_OWNED field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIConfigConstantsGROUPS) GetMAX_OWNEDOk() (*int32, bool) {
	if o == nil || IsNil(o.MAX_OWNED) {
		return nil, false
	}
	return o.MAX_OWNED, true
}

// HasMAX_OWNED returns a boolean if a field has been set.
func (o *APIConfigConstantsGROUPS) HasMAX_OWNED() bool {
	if o != nil && !IsNil(o.MAX_OWNED) {
		return true
	}

	return false
}

// SetMAX_OWNED gets a reference to the given int32 and assigns it to the MAX_OWNED field.
func (o *APIConfigConstantsGROUPS) SetMAX_OWNED(v int32) {
	o.MAX_OWNED = &v
}

// GetMAX_ROLES returns the MAX_ROLES field value if set, zero value otherwise.
func (o *APIConfigConstantsGROUPS) GetMAX_ROLES() int32 {
	if o == nil || IsNil(o.MAX_ROLES) {
		var ret int32
		return ret
	}
	return *o.MAX_ROLES
}

// GetMAX_ROLESOk returns a tuple with the MAX_ROLES field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIConfigConstantsGROUPS) GetMAX_ROLESOk() (*int32, bool) {
	if o == nil || IsNil(o.MAX_ROLES) {
		return nil, false
	}
	return o.MAX_ROLES, true
}

// HasMAX_ROLES returns a boolean if a field has been set.
func (o *APIConfigConstantsGROUPS) HasMAX_ROLES() bool {
	if o != nil && !IsNil(o.MAX_ROLES) {
		return true
	}

	return false
}

// SetMAX_ROLES gets a reference to the given int32 and assigns it to the MAX_ROLES field.
func (o *APIConfigConstantsGROUPS) SetMAX_ROLES(v int32) {
	o.MAX_ROLES = &v
}

func (o APIConfigConstantsGROUPS) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIConfigConstantsGROUPS) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CAPACITY) {
		toSerialize["CAPACITY"] = o.CAPACITY
	}
	if !IsNil(o.GROUP_TRANSFER_REQUIREMENTS) {
		toSerialize["GROUP_TRANSFER_REQUIREMENTS"] = o.GROUP_TRANSFER_REQUIREMENTS
	}
	if !IsNil(o.MAX_INVITES_REQUESTS) {
		toSerialize["MAX_INVITES_REQUESTS"] = o.MAX_INVITES_REQUESTS
	}
	if !IsNil(o.MAX_JOINED) {
		toSerialize["MAX_JOINED"] = o.MAX_JOINED
	}
	if !IsNil(o.MAX_JOINED_PLUS) {
		toSerialize["MAX_JOINED_PLUS"] = o.MAX_JOINED_PLUS
	}
	if !IsNil(o.MAX_LANGUAGES) {
		toSerialize["MAX_LANGUAGES"] = o.MAX_LANGUAGES
	}
	if !IsNil(o.MAX_LINKS) {
		toSerialize["MAX_LINKS"] = o.MAX_LINKS
	}
	if !IsNil(o.MAX_MANAGEMENT_ROLES) {
		toSerialize["MAX_MANAGEMENT_ROLES"] = o.MAX_MANAGEMENT_ROLES
	}
	if !IsNil(o.MAX_OWNED) {
		toSerialize["MAX_OWNED"] = o.MAX_OWNED
	}
	if !IsNil(o.MAX_ROLES) {
		toSerialize["MAX_ROLES"] = o.MAX_ROLES
	}
	return toSerialize, nil
}

type NullableAPIConfigConstantsGROUPS struct {
	value *APIConfigConstantsGROUPS
	isSet bool
}

func (v NullableAPIConfigConstantsGROUPS) Get() *APIConfigConstantsGROUPS {
	return v.value
}

func (v *NullableAPIConfigConstantsGROUPS) Set(val *APIConfigConstantsGROUPS) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIConfigConstantsGROUPS) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIConfigConstantsGROUPS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIConfigConstantsGROUPS(val *APIConfigConstantsGROUPS) *NullableAPIConfigConstantsGROUPS {
	return &NullableAPIConfigConstantsGROUPS{value: val, isSet: true}
}

func (v NullableAPIConfigConstantsGROUPS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIConfigConstantsGROUPS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


