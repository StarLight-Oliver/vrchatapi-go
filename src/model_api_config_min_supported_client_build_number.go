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

// checks if the APIConfigMinSupportedClientBuildNumber type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIConfigMinSupportedClientBuildNumber{}

// APIConfigMinSupportedClientBuildNumber Minimum supported client build number for various platforms
type APIConfigMinSupportedClientBuildNumber struct {
	AppStore PlatformBuildInfo `json:"AppStore"`
	Default PlatformBuildInfo `json:"Default"`
	Firebase PlatformBuildInfo `json:"Firebase"`
	FirebaseiOS PlatformBuildInfo `json:"FirebaseiOS"`
	GooglePlay PlatformBuildInfo `json:"GooglePlay"`
	PC PlatformBuildInfo `json:"PC"`
	PicoStore PlatformBuildInfo `json:"PicoStore"`
	QuestAppLab PlatformBuildInfo `json:"QuestAppLab"`
	QuestStore PlatformBuildInfo `json:"QuestStore"`
	TestFlight PlatformBuildInfo `json:"TestFlight"`
	XRElite PlatformBuildInfo `json:"XRElite"`
}

type _APIConfigMinSupportedClientBuildNumber APIConfigMinSupportedClientBuildNumber

// NewAPIConfigMinSupportedClientBuildNumber instantiates a new APIConfigMinSupportedClientBuildNumber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIConfigMinSupportedClientBuildNumber(appStore PlatformBuildInfo, default_ PlatformBuildInfo, firebase PlatformBuildInfo, firebaseiOS PlatformBuildInfo, googlePlay PlatformBuildInfo, pC PlatformBuildInfo, picoStore PlatformBuildInfo, questAppLab PlatformBuildInfo, questStore PlatformBuildInfo, testFlight PlatformBuildInfo, xRElite PlatformBuildInfo) *APIConfigMinSupportedClientBuildNumber {
	this := APIConfigMinSupportedClientBuildNumber{}
	this.AppStore = appStore
	this.Default = default_
	this.Firebase = firebase
	this.FirebaseiOS = firebaseiOS
	this.GooglePlay = googlePlay
	this.PC = pC
	this.PicoStore = picoStore
	this.QuestAppLab = questAppLab
	this.QuestStore = questStore
	this.TestFlight = testFlight
	this.XRElite = xRElite
	return &this
}

// NewAPIConfigMinSupportedClientBuildNumberWithDefaults instantiates a new APIConfigMinSupportedClientBuildNumber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIConfigMinSupportedClientBuildNumberWithDefaults() *APIConfigMinSupportedClientBuildNumber {
	this := APIConfigMinSupportedClientBuildNumber{}
	return &this
}

// GetAppStore returns the AppStore field value
func (o *APIConfigMinSupportedClientBuildNumber) GetAppStore() PlatformBuildInfo {
	if o == nil {
		var ret PlatformBuildInfo
		return ret
	}

	return o.AppStore
}

// GetAppStoreOk returns a tuple with the AppStore field value
// and a boolean to check if the value has been set.
func (o *APIConfigMinSupportedClientBuildNumber) GetAppStoreOk() (*PlatformBuildInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppStore, true
}

// SetAppStore sets field value
func (o *APIConfigMinSupportedClientBuildNumber) SetAppStore(v PlatformBuildInfo) {
	o.AppStore = v
}

// GetDefault returns the Default field value
func (o *APIConfigMinSupportedClientBuildNumber) GetDefault() PlatformBuildInfo {
	if o == nil {
		var ret PlatformBuildInfo
		return ret
	}

	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *APIConfigMinSupportedClientBuildNumber) GetDefaultOk() (*PlatformBuildInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value
func (o *APIConfigMinSupportedClientBuildNumber) SetDefault(v PlatformBuildInfo) {
	o.Default = v
}

// GetFirebase returns the Firebase field value
func (o *APIConfigMinSupportedClientBuildNumber) GetFirebase() PlatformBuildInfo {
	if o == nil {
		var ret PlatformBuildInfo
		return ret
	}

	return o.Firebase
}

// GetFirebaseOk returns a tuple with the Firebase field value
// and a boolean to check if the value has been set.
func (o *APIConfigMinSupportedClientBuildNumber) GetFirebaseOk() (*PlatformBuildInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Firebase, true
}

// SetFirebase sets field value
func (o *APIConfigMinSupportedClientBuildNumber) SetFirebase(v PlatformBuildInfo) {
	o.Firebase = v
}

// GetFirebaseiOS returns the FirebaseiOS field value
func (o *APIConfigMinSupportedClientBuildNumber) GetFirebaseiOS() PlatformBuildInfo {
	if o == nil {
		var ret PlatformBuildInfo
		return ret
	}

	return o.FirebaseiOS
}

// GetFirebaseiOSOk returns a tuple with the FirebaseiOS field value
// and a boolean to check if the value has been set.
func (o *APIConfigMinSupportedClientBuildNumber) GetFirebaseiOSOk() (*PlatformBuildInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirebaseiOS, true
}

// SetFirebaseiOS sets field value
func (o *APIConfigMinSupportedClientBuildNumber) SetFirebaseiOS(v PlatformBuildInfo) {
	o.FirebaseiOS = v
}

// GetGooglePlay returns the GooglePlay field value
func (o *APIConfigMinSupportedClientBuildNumber) GetGooglePlay() PlatformBuildInfo {
	if o == nil {
		var ret PlatformBuildInfo
		return ret
	}

	return o.GooglePlay
}

// GetGooglePlayOk returns a tuple with the GooglePlay field value
// and a boolean to check if the value has been set.
func (o *APIConfigMinSupportedClientBuildNumber) GetGooglePlayOk() (*PlatformBuildInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GooglePlay, true
}

// SetGooglePlay sets field value
func (o *APIConfigMinSupportedClientBuildNumber) SetGooglePlay(v PlatformBuildInfo) {
	o.GooglePlay = v
}

// GetPC returns the PC field value
func (o *APIConfigMinSupportedClientBuildNumber) GetPC() PlatformBuildInfo {
	if o == nil {
		var ret PlatformBuildInfo
		return ret
	}

	return o.PC
}

// GetPCOk returns a tuple with the PC field value
// and a boolean to check if the value has been set.
func (o *APIConfigMinSupportedClientBuildNumber) GetPCOk() (*PlatformBuildInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PC, true
}

// SetPC sets field value
func (o *APIConfigMinSupportedClientBuildNumber) SetPC(v PlatformBuildInfo) {
	o.PC = v
}

// GetPicoStore returns the PicoStore field value
func (o *APIConfigMinSupportedClientBuildNumber) GetPicoStore() PlatformBuildInfo {
	if o == nil {
		var ret PlatformBuildInfo
		return ret
	}

	return o.PicoStore
}

// GetPicoStoreOk returns a tuple with the PicoStore field value
// and a boolean to check if the value has been set.
func (o *APIConfigMinSupportedClientBuildNumber) GetPicoStoreOk() (*PlatformBuildInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PicoStore, true
}

// SetPicoStore sets field value
func (o *APIConfigMinSupportedClientBuildNumber) SetPicoStore(v PlatformBuildInfo) {
	o.PicoStore = v
}

// GetQuestAppLab returns the QuestAppLab field value
func (o *APIConfigMinSupportedClientBuildNumber) GetQuestAppLab() PlatformBuildInfo {
	if o == nil {
		var ret PlatformBuildInfo
		return ret
	}

	return o.QuestAppLab
}

// GetQuestAppLabOk returns a tuple with the QuestAppLab field value
// and a boolean to check if the value has been set.
func (o *APIConfigMinSupportedClientBuildNumber) GetQuestAppLabOk() (*PlatformBuildInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuestAppLab, true
}

// SetQuestAppLab sets field value
func (o *APIConfigMinSupportedClientBuildNumber) SetQuestAppLab(v PlatformBuildInfo) {
	o.QuestAppLab = v
}

// GetQuestStore returns the QuestStore field value
func (o *APIConfigMinSupportedClientBuildNumber) GetQuestStore() PlatformBuildInfo {
	if o == nil {
		var ret PlatformBuildInfo
		return ret
	}

	return o.QuestStore
}

// GetQuestStoreOk returns a tuple with the QuestStore field value
// and a boolean to check if the value has been set.
func (o *APIConfigMinSupportedClientBuildNumber) GetQuestStoreOk() (*PlatformBuildInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuestStore, true
}

// SetQuestStore sets field value
func (o *APIConfigMinSupportedClientBuildNumber) SetQuestStore(v PlatformBuildInfo) {
	o.QuestStore = v
}

// GetTestFlight returns the TestFlight field value
func (o *APIConfigMinSupportedClientBuildNumber) GetTestFlight() PlatformBuildInfo {
	if o == nil {
		var ret PlatformBuildInfo
		return ret
	}

	return o.TestFlight
}

// GetTestFlightOk returns a tuple with the TestFlight field value
// and a boolean to check if the value has been set.
func (o *APIConfigMinSupportedClientBuildNumber) GetTestFlightOk() (*PlatformBuildInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestFlight, true
}

// SetTestFlight sets field value
func (o *APIConfigMinSupportedClientBuildNumber) SetTestFlight(v PlatformBuildInfo) {
	o.TestFlight = v
}

// GetXRElite returns the XRElite field value
func (o *APIConfigMinSupportedClientBuildNumber) GetXRElite() PlatformBuildInfo {
	if o == nil {
		var ret PlatformBuildInfo
		return ret
	}

	return o.XRElite
}

// GetXREliteOk returns a tuple with the XRElite field value
// and a boolean to check if the value has been set.
func (o *APIConfigMinSupportedClientBuildNumber) GetXREliteOk() (*PlatformBuildInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.XRElite, true
}

// SetXRElite sets field value
func (o *APIConfigMinSupportedClientBuildNumber) SetXRElite(v PlatformBuildInfo) {
	o.XRElite = v
}

func (o APIConfigMinSupportedClientBuildNumber) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIConfigMinSupportedClientBuildNumber) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["AppStore"] = o.AppStore
	toSerialize["Default"] = o.Default
	toSerialize["Firebase"] = o.Firebase
	toSerialize["FirebaseiOS"] = o.FirebaseiOS
	toSerialize["GooglePlay"] = o.GooglePlay
	toSerialize["PC"] = o.PC
	toSerialize["PicoStore"] = o.PicoStore
	toSerialize["QuestAppLab"] = o.QuestAppLab
	toSerialize["QuestStore"] = o.QuestStore
	toSerialize["TestFlight"] = o.TestFlight
	toSerialize["XRElite"] = o.XRElite
	return toSerialize, nil
}

func (o *APIConfigMinSupportedClientBuildNumber) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"AppStore",
		"Default",
		"Firebase",
		"FirebaseiOS",
		"GooglePlay",
		"PC",
		"PicoStore",
		"QuestAppLab",
		"QuestStore",
		"TestFlight",
		"XRElite",
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

	varAPIConfigMinSupportedClientBuildNumber := _APIConfigMinSupportedClientBuildNumber{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAPIConfigMinSupportedClientBuildNumber)

	if err != nil {
		return err
	}

	*o = APIConfigMinSupportedClientBuildNumber(varAPIConfigMinSupportedClientBuildNumber)

	return err
}

type NullableAPIConfigMinSupportedClientBuildNumber struct {
	value *APIConfigMinSupportedClientBuildNumber
	isSet bool
}

func (v NullableAPIConfigMinSupportedClientBuildNumber) Get() *APIConfigMinSupportedClientBuildNumber {
	return v.value
}

func (v *NullableAPIConfigMinSupportedClientBuildNumber) Set(val *APIConfigMinSupportedClientBuildNumber) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIConfigMinSupportedClientBuildNumber) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIConfigMinSupportedClientBuildNumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIConfigMinSupportedClientBuildNumber(val *APIConfigMinSupportedClientBuildNumber) *NullableAPIConfigMinSupportedClientBuildNumber {
	return &NullableAPIConfigMinSupportedClientBuildNumber{value: val, isSet: true}
}

func (v NullableAPIConfigMinSupportedClientBuildNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIConfigMinSupportedClientBuildNumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


