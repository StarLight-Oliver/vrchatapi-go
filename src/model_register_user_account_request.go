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

// checks if the RegisterUserAccountRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegisterUserAccountRequest{}

// RegisterUserAccountRequest struct for RegisterUserAccountRequest
type RegisterUserAccountRequest struct {
	// The most recent version of the TOS
	AcceptedTOSVersion int32 `json:"acceptedTOSVersion"`
	// Captcha code
	CaptchaCode string `json:"captchaCode"`
	// Birth day of month
	Day string `json:"day"`
	// Email address
	Email string `json:"email"`
	// Birth month of year
	Month string `json:"month"`
	// Password
	Password string `json:"password"`
	// Whether to recieve promotional emails
	Subscribe bool `json:"subscribe"`
	// Display Name / Username (Username is a sanitized version)
	Username string `json:"username"`
	// Birth year
	Year string `json:"year"`
}

type _RegisterUserAccountRequest RegisterUserAccountRequest

// NewRegisterUserAccountRequest instantiates a new RegisterUserAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterUserAccountRequest(acceptedTOSVersion int32, captchaCode string, day string, email string, month string, password string, subscribe bool, username string, year string) *RegisterUserAccountRequest {
	this := RegisterUserAccountRequest{}
	this.AcceptedTOSVersion = acceptedTOSVersion
	this.CaptchaCode = captchaCode
	this.Day = day
	this.Email = email
	this.Month = month
	this.Password = password
	this.Subscribe = subscribe
	this.Username = username
	this.Year = year
	return &this
}

// NewRegisterUserAccountRequestWithDefaults instantiates a new RegisterUserAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterUserAccountRequestWithDefaults() *RegisterUserAccountRequest {
	this := RegisterUserAccountRequest{}
	return &this
}

// GetAcceptedTOSVersion returns the AcceptedTOSVersion field value
func (o *RegisterUserAccountRequest) GetAcceptedTOSVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AcceptedTOSVersion
}

// GetAcceptedTOSVersionOk returns a tuple with the AcceptedTOSVersion field value
// and a boolean to check if the value has been set.
func (o *RegisterUserAccountRequest) GetAcceptedTOSVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcceptedTOSVersion, true
}

// SetAcceptedTOSVersion sets field value
func (o *RegisterUserAccountRequest) SetAcceptedTOSVersion(v int32) {
	o.AcceptedTOSVersion = v
}

// GetCaptchaCode returns the CaptchaCode field value
func (o *RegisterUserAccountRequest) GetCaptchaCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CaptchaCode
}

// GetCaptchaCodeOk returns a tuple with the CaptchaCode field value
// and a boolean to check if the value has been set.
func (o *RegisterUserAccountRequest) GetCaptchaCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CaptchaCode, true
}

// SetCaptchaCode sets field value
func (o *RegisterUserAccountRequest) SetCaptchaCode(v string) {
	o.CaptchaCode = v
}

// GetDay returns the Day field value
func (o *RegisterUserAccountRequest) GetDay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Day
}

// GetDayOk returns a tuple with the Day field value
// and a boolean to check if the value has been set.
func (o *RegisterUserAccountRequest) GetDayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Day, true
}

// SetDay sets field value
func (o *RegisterUserAccountRequest) SetDay(v string) {
	o.Day = v
}

// GetEmail returns the Email field value
func (o *RegisterUserAccountRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *RegisterUserAccountRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *RegisterUserAccountRequest) SetEmail(v string) {
	o.Email = v
}

// GetMonth returns the Month field value
func (o *RegisterUserAccountRequest) GetMonth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Month
}

// GetMonthOk returns a tuple with the Month field value
// and a boolean to check if the value has been set.
func (o *RegisterUserAccountRequest) GetMonthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Month, true
}

// SetMonth sets field value
func (o *RegisterUserAccountRequest) SetMonth(v string) {
	o.Month = v
}

// GetPassword returns the Password field value
func (o *RegisterUserAccountRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *RegisterUserAccountRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *RegisterUserAccountRequest) SetPassword(v string) {
	o.Password = v
}

// GetSubscribe returns the Subscribe field value
func (o *RegisterUserAccountRequest) GetSubscribe() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Subscribe
}

// GetSubscribeOk returns a tuple with the Subscribe field value
// and a boolean to check if the value has been set.
func (o *RegisterUserAccountRequest) GetSubscribeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscribe, true
}

// SetSubscribe sets field value
func (o *RegisterUserAccountRequest) SetSubscribe(v bool) {
	o.Subscribe = v
}

// GetUsername returns the Username field value
func (o *RegisterUserAccountRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *RegisterUserAccountRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *RegisterUserAccountRequest) SetUsername(v string) {
	o.Username = v
}

// GetYear returns the Year field value
func (o *RegisterUserAccountRequest) GetYear() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Year
}

// GetYearOk returns a tuple with the Year field value
// and a boolean to check if the value has been set.
func (o *RegisterUserAccountRequest) GetYearOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Year, true
}

// SetYear sets field value
func (o *RegisterUserAccountRequest) SetYear(v string) {
	o.Year = v
}

func (o RegisterUserAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegisterUserAccountRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["acceptedTOSVersion"] = o.AcceptedTOSVersion
	toSerialize["captchaCode"] = o.CaptchaCode
	toSerialize["day"] = o.Day
	toSerialize["email"] = o.Email
	toSerialize["month"] = o.Month
	toSerialize["password"] = o.Password
	toSerialize["subscribe"] = o.Subscribe
	toSerialize["username"] = o.Username
	toSerialize["year"] = o.Year
	return toSerialize, nil
}

func (o *RegisterUserAccountRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"acceptedTOSVersion",
		"captchaCode",
		"day",
		"email",
		"month",
		"password",
		"subscribe",
		"username",
		"year",
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

	varRegisterUserAccountRequest := _RegisterUserAccountRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegisterUserAccountRequest)

	if err != nil {
		return err
	}

	*o = RegisterUserAccountRequest(varRegisterUserAccountRequest)

	return err
}

type NullableRegisterUserAccountRequest struct {
	value *RegisterUserAccountRequest
	isSet bool
}

func (v NullableRegisterUserAccountRequest) Get() *RegisterUserAccountRequest {
	return v.value
}

func (v *NullableRegisterUserAccountRequest) Set(val *RegisterUserAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterUserAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterUserAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterUserAccountRequest(val *RegisterUserAccountRequest) *NullableRegisterUserAccountRequest {
	return &NullableRegisterUserAccountRequest{value: val, isSet: true}
}

func (v NullableRegisterUserAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterUserAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


