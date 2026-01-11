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
	"fmt"
	"gopkg.in/validator.v2"
)

// SentNotificationDetails - struct for SentNotificationDetails
type SentNotificationDetails struct {
	NotificationDetailBoop *NotificationDetailBoop
	NotificationDetailInvite *NotificationDetailInvite
	NotificationDetailInviteResponse *NotificationDetailInviteResponse
	NotificationDetailRequestInvite *NotificationDetailRequestInvite
	NotificationDetailRequestInviteResponse *NotificationDetailRequestInviteResponse
	NotificationDetailVoteToKick *NotificationDetailVoteToKick
	MapmapOfStringAny *map[string]interface{}
}

// NotificationDetailBoopAsSentNotificationDetails is a convenience function that returns NotificationDetailBoop wrapped in SentNotificationDetails
func NotificationDetailBoopAsSentNotificationDetails(v *NotificationDetailBoop) SentNotificationDetails {
	return SentNotificationDetails{
		NotificationDetailBoop: v,
	}
}

// NotificationDetailInviteAsSentNotificationDetails is a convenience function that returns NotificationDetailInvite wrapped in SentNotificationDetails
func NotificationDetailInviteAsSentNotificationDetails(v *NotificationDetailInvite) SentNotificationDetails {
	return SentNotificationDetails{
		NotificationDetailInvite: v,
	}
}

// NotificationDetailInviteResponseAsSentNotificationDetails is a convenience function that returns NotificationDetailInviteResponse wrapped in SentNotificationDetails
func NotificationDetailInviteResponseAsSentNotificationDetails(v *NotificationDetailInviteResponse) SentNotificationDetails {
	return SentNotificationDetails{
		NotificationDetailInviteResponse: v,
	}
}

// NotificationDetailRequestInviteAsSentNotificationDetails is a convenience function that returns NotificationDetailRequestInvite wrapped in SentNotificationDetails
func NotificationDetailRequestInviteAsSentNotificationDetails(v *NotificationDetailRequestInvite) SentNotificationDetails {
	return SentNotificationDetails{
		NotificationDetailRequestInvite: v,
	}
}

// NotificationDetailRequestInviteResponseAsSentNotificationDetails is a convenience function that returns NotificationDetailRequestInviteResponse wrapped in SentNotificationDetails
func NotificationDetailRequestInviteResponseAsSentNotificationDetails(v *NotificationDetailRequestInviteResponse) SentNotificationDetails {
	return SentNotificationDetails{
		NotificationDetailRequestInviteResponse: v,
	}
}

// NotificationDetailVoteToKickAsSentNotificationDetails is a convenience function that returns NotificationDetailVoteToKick wrapped in SentNotificationDetails
func NotificationDetailVoteToKickAsSentNotificationDetails(v *NotificationDetailVoteToKick) SentNotificationDetails {
	return SentNotificationDetails{
		NotificationDetailVoteToKick: v,
	}
}

// map[string]interface{}AsSentNotificationDetails is a convenience function that returns map[string]interface{} wrapped in SentNotificationDetails
func MapmapOfStringAnyAsSentNotificationDetails(v *map[string]interface{}) SentNotificationDetails {
	return SentNotificationDetails{
		MapmapOfStringAny: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SentNotificationDetails) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into NotificationDetailBoop
	err = newStrictDecoder(data).Decode(&dst.NotificationDetailBoop)
	if err == nil {
		jsonNotificationDetailBoop, _ := json.Marshal(dst.NotificationDetailBoop)
		if string(jsonNotificationDetailBoop) == "{}" { // empty struct
			dst.NotificationDetailBoop = nil
		} else {
			if err = validator.Validate(dst.NotificationDetailBoop); err != nil {
				dst.NotificationDetailBoop = nil
			} else {
				match++
			}
		}
	} else {
		dst.NotificationDetailBoop = nil
	}

	// try to unmarshal data into NotificationDetailInvite
	err = newStrictDecoder(data).Decode(&dst.NotificationDetailInvite)
	if err == nil {
		jsonNotificationDetailInvite, _ := json.Marshal(dst.NotificationDetailInvite)
		if string(jsonNotificationDetailInvite) == "{}" { // empty struct
			dst.NotificationDetailInvite = nil
		} else {
			if err = validator.Validate(dst.NotificationDetailInvite); err != nil {
				dst.NotificationDetailInvite = nil
			} else {
				match++
			}
		}
	} else {
		dst.NotificationDetailInvite = nil
	}

	// try to unmarshal data into NotificationDetailInviteResponse
	err = newStrictDecoder(data).Decode(&dst.NotificationDetailInviteResponse)
	if err == nil {
		jsonNotificationDetailInviteResponse, _ := json.Marshal(dst.NotificationDetailInviteResponse)
		if string(jsonNotificationDetailInviteResponse) == "{}" { // empty struct
			dst.NotificationDetailInviteResponse = nil
		} else {
			if err = validator.Validate(dst.NotificationDetailInviteResponse); err != nil {
				dst.NotificationDetailInviteResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.NotificationDetailInviteResponse = nil
	}

	// try to unmarshal data into NotificationDetailRequestInvite
	err = newStrictDecoder(data).Decode(&dst.NotificationDetailRequestInvite)
	if err == nil {
		jsonNotificationDetailRequestInvite, _ := json.Marshal(dst.NotificationDetailRequestInvite)
		if string(jsonNotificationDetailRequestInvite) == "{}" { // empty struct
			dst.NotificationDetailRequestInvite = nil
		} else {
			if err = validator.Validate(dst.NotificationDetailRequestInvite); err != nil {
				dst.NotificationDetailRequestInvite = nil
			} else {
				match++
			}
		}
	} else {
		dst.NotificationDetailRequestInvite = nil
	}

	// try to unmarshal data into NotificationDetailRequestInviteResponse
	err = newStrictDecoder(data).Decode(&dst.NotificationDetailRequestInviteResponse)
	if err == nil {
		jsonNotificationDetailRequestInviteResponse, _ := json.Marshal(dst.NotificationDetailRequestInviteResponse)
		if string(jsonNotificationDetailRequestInviteResponse) == "{}" { // empty struct
			dst.NotificationDetailRequestInviteResponse = nil
		} else {
			if err = validator.Validate(dst.NotificationDetailRequestInviteResponse); err != nil {
				dst.NotificationDetailRequestInviteResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.NotificationDetailRequestInviteResponse = nil
	}

	// try to unmarshal data into NotificationDetailVoteToKick
	err = newStrictDecoder(data).Decode(&dst.NotificationDetailVoteToKick)
	if err == nil {
		jsonNotificationDetailVoteToKick, _ := json.Marshal(dst.NotificationDetailVoteToKick)
		if string(jsonNotificationDetailVoteToKick) == "{}" { // empty struct
			dst.NotificationDetailVoteToKick = nil
		} else {
			if err = validator.Validate(dst.NotificationDetailVoteToKick); err != nil {
				dst.NotificationDetailVoteToKick = nil
			} else {
				match++
			}
		}
	} else {
		dst.NotificationDetailVoteToKick = nil
	}

	// try to unmarshal data into MapmapOfStringAny
	err = newStrictDecoder(data).Decode(&dst.MapmapOfStringAny)
	if err == nil {
		jsonMapmapOfStringAny, _ := json.Marshal(dst.MapmapOfStringAny)
		if string(jsonMapmapOfStringAny) == "{}" { // empty struct
			dst.MapmapOfStringAny = nil
		} else {
			if err = validator.Validate(dst.MapmapOfStringAny); err != nil {
				dst.MapmapOfStringAny = nil
			} else {
				match++
			}
		}
	} else {
		dst.MapmapOfStringAny = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.NotificationDetailBoop = nil
		dst.NotificationDetailInvite = nil
		dst.NotificationDetailInviteResponse = nil
		dst.NotificationDetailRequestInvite = nil
		dst.NotificationDetailRequestInviteResponse = nil
		dst.NotificationDetailVoteToKick = nil
		dst.MapmapOfStringAny = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SentNotificationDetails)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SentNotificationDetails)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SentNotificationDetails) MarshalJSON() ([]byte, error) {
	if src.NotificationDetailBoop != nil {
		return json.Marshal(&src.NotificationDetailBoop)
	}

	if src.NotificationDetailInvite != nil {
		return json.Marshal(&src.NotificationDetailInvite)
	}

	if src.NotificationDetailInviteResponse != nil {
		return json.Marshal(&src.NotificationDetailInviteResponse)
	}

	if src.NotificationDetailRequestInvite != nil {
		return json.Marshal(&src.NotificationDetailRequestInvite)
	}

	if src.NotificationDetailRequestInviteResponse != nil {
		return json.Marshal(&src.NotificationDetailRequestInviteResponse)
	}

	if src.NotificationDetailVoteToKick != nil {
		return json.Marshal(&src.NotificationDetailVoteToKick)
	}

	if src.MapmapOfStringAny != nil {
		return json.Marshal(&src.MapmapOfStringAny)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SentNotificationDetails) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.NotificationDetailBoop != nil {
		return obj.NotificationDetailBoop
	}

	if obj.NotificationDetailInvite != nil {
		return obj.NotificationDetailInvite
	}

	if obj.NotificationDetailInviteResponse != nil {
		return obj.NotificationDetailInviteResponse
	}

	if obj.NotificationDetailRequestInvite != nil {
		return obj.NotificationDetailRequestInvite
	}

	if obj.NotificationDetailRequestInviteResponse != nil {
		return obj.NotificationDetailRequestInviteResponse
	}

	if obj.NotificationDetailVoteToKick != nil {
		return obj.NotificationDetailVoteToKick
	}

	if obj.MapmapOfStringAny != nil {
		return obj.MapmapOfStringAny
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj SentNotificationDetails) GetActualInstanceValue() (interface{}) {
	if obj.NotificationDetailBoop != nil {
		return *obj.NotificationDetailBoop
	}

	if obj.NotificationDetailInvite != nil {
		return *obj.NotificationDetailInvite
	}

	if obj.NotificationDetailInviteResponse != nil {
		return *obj.NotificationDetailInviteResponse
	}

	if obj.NotificationDetailRequestInvite != nil {
		return *obj.NotificationDetailRequestInvite
	}

	if obj.NotificationDetailRequestInviteResponse != nil {
		return *obj.NotificationDetailRequestInviteResponse
	}

	if obj.NotificationDetailVoteToKick != nil {
		return *obj.NotificationDetailVoteToKick
	}

	if obj.MapmapOfStringAny != nil {
		return *obj.MapmapOfStringAny
	}

	// all schemas are nil
	return nil
}

type NullableSentNotificationDetails struct {
	value *SentNotificationDetails
	isSet bool
}

func (v NullableSentNotificationDetails) Get() *SentNotificationDetails {
	return v.value
}

func (v *NullableSentNotificationDetails) Set(val *SentNotificationDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSentNotificationDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSentNotificationDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSentNotificationDetails(val *SentNotificationDetails) *NullableSentNotificationDetails {
	return &NullableSentNotificationDetails{value: val, isSet: true}
}

func (v NullableSentNotificationDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSentNotificationDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


