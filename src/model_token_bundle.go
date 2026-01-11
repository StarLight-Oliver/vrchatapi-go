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

// checks if the TokenBundle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenBundle{}

// TokenBundle struct for TokenBundle
type TokenBundle struct {
	// price of the bundle
	Amount int32 `json:"amount"`
	AppleProductId string `json:"appleProductId"`
	Description string `json:"description"`
	GoogleProductId *string `json:"googleProductId,omitempty"`
	Id string `json:"id"`
	// direct url to image
	ImageUrl string `json:"imageUrl"`
	OculusSku string `json:"oculusSku"`
	SteamItemId string `json:"steamItemId"`
	// number of tokens received
	Tokens int32 `json:"tokens"`
}

type _TokenBundle TokenBundle

// NewTokenBundle instantiates a new TokenBundle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenBundle(amount int32, appleProductId string, description string, id string, imageUrl string, oculusSku string, steamItemId string, tokens int32) *TokenBundle {
	this := TokenBundle{}
	this.Amount = amount
	this.AppleProductId = appleProductId
	this.Description = description
	this.Id = id
	this.ImageUrl = imageUrl
	this.OculusSku = oculusSku
	this.SteamItemId = steamItemId
	this.Tokens = tokens
	return &this
}

// NewTokenBundleWithDefaults instantiates a new TokenBundle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenBundleWithDefaults() *TokenBundle {
	this := TokenBundle{}
	return &this
}

// GetAmount returns the Amount field value
func (o *TokenBundle) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TokenBundle) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TokenBundle) SetAmount(v int32) {
	o.Amount = v
}

// GetAppleProductId returns the AppleProductId field value
func (o *TokenBundle) GetAppleProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppleProductId
}

// GetAppleProductIdOk returns a tuple with the AppleProductId field value
// and a boolean to check if the value has been set.
func (o *TokenBundle) GetAppleProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppleProductId, true
}

// SetAppleProductId sets field value
func (o *TokenBundle) SetAppleProductId(v string) {
	o.AppleProductId = v
}

// GetDescription returns the Description field value
func (o *TokenBundle) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TokenBundle) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TokenBundle) SetDescription(v string) {
	o.Description = v
}

// GetGoogleProductId returns the GoogleProductId field value if set, zero value otherwise.
func (o *TokenBundle) GetGoogleProductId() string {
	if o == nil || IsNil(o.GoogleProductId) {
		var ret string
		return ret
	}
	return *o.GoogleProductId
}

// GetGoogleProductIdOk returns a tuple with the GoogleProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenBundle) GetGoogleProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.GoogleProductId) {
		return nil, false
	}
	return o.GoogleProductId, true
}

// HasGoogleProductId returns a boolean if a field has been set.
func (o *TokenBundle) HasGoogleProductId() bool {
	if o != nil && !IsNil(o.GoogleProductId) {
		return true
	}

	return false
}

// SetGoogleProductId gets a reference to the given string and assigns it to the GoogleProductId field.
func (o *TokenBundle) SetGoogleProductId(v string) {
	o.GoogleProductId = &v
}

// GetId returns the Id field value
func (o *TokenBundle) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TokenBundle) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TokenBundle) SetId(v string) {
	o.Id = v
}

// GetImageUrl returns the ImageUrl field value
func (o *TokenBundle) GetImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
func (o *TokenBundle) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageUrl, true
}

// SetImageUrl sets field value
func (o *TokenBundle) SetImageUrl(v string) {
	o.ImageUrl = v
}

// GetOculusSku returns the OculusSku field value
func (o *TokenBundle) GetOculusSku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OculusSku
}

// GetOculusSkuOk returns a tuple with the OculusSku field value
// and a boolean to check if the value has been set.
func (o *TokenBundle) GetOculusSkuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OculusSku, true
}

// SetOculusSku sets field value
func (o *TokenBundle) SetOculusSku(v string) {
	o.OculusSku = v
}

// GetSteamItemId returns the SteamItemId field value
func (o *TokenBundle) GetSteamItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SteamItemId
}

// GetSteamItemIdOk returns a tuple with the SteamItemId field value
// and a boolean to check if the value has been set.
func (o *TokenBundle) GetSteamItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SteamItemId, true
}

// SetSteamItemId sets field value
func (o *TokenBundle) SetSteamItemId(v string) {
	o.SteamItemId = v
}

// GetTokens returns the Tokens field value
func (o *TokenBundle) GetTokens() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value
// and a boolean to check if the value has been set.
func (o *TokenBundle) GetTokensOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tokens, true
}

// SetTokens sets field value
func (o *TokenBundle) SetTokens(v int32) {
	o.Tokens = v
}

func (o TokenBundle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenBundle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["appleProductId"] = o.AppleProductId
	toSerialize["description"] = o.Description
	if !IsNil(o.GoogleProductId) {
		toSerialize["googleProductId"] = o.GoogleProductId
	}
	toSerialize["id"] = o.Id
	toSerialize["imageUrl"] = o.ImageUrl
	toSerialize["oculusSku"] = o.OculusSku
	toSerialize["steamItemId"] = o.SteamItemId
	toSerialize["tokens"] = o.Tokens
	return toSerialize, nil
}

func (o *TokenBundle) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"appleProductId",
		"description",
		"id",
		"imageUrl",
		"oculusSku",
		"steamItemId",
		"tokens",
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

	varTokenBundle := _TokenBundle{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTokenBundle)

	if err != nil {
		return err
	}

	*o = TokenBundle(varTokenBundle)

	return err
}

type NullableTokenBundle struct {
	value *TokenBundle
	isSet bool
}

func (v NullableTokenBundle) Get() *TokenBundle {
	return v.value
}

func (v *NullableTokenBundle) Set(val *TokenBundle) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenBundle) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenBundle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenBundle(val *TokenBundle) *NullableTokenBundle {
	return &NullableTokenBundle{value: val, isSet: true}
}

func (v NullableTokenBundle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenBundle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


