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

// NotificationV2Data - struct for NotificationV2Data
type NotificationV2Data struct {
	NotificationV2DataBadgeEarned *NotificationV2DataBadgeEarned
	NotificationV2DataBoop *NotificationV2DataBoop
	NotificationV2DataEventAnnouncement *NotificationV2DataEventAnnouncement
	NotificationV2DataGroupAnnouncement *NotificationV2DataGroupAnnouncement
	NotificationV2DataGroupInformative *NotificationV2DataGroupInformative
	NotificationV2DataGroupTransfer *NotificationV2DataGroupTransfer
	MapmapOfStringAny *map[string]interface{}
}

// NotificationV2DataBadgeEarnedAsNotificationV2Data is a convenience function that returns NotificationV2DataBadgeEarned wrapped in NotificationV2Data
func NotificationV2DataBadgeEarnedAsNotificationV2Data(v *NotificationV2DataBadgeEarned) NotificationV2Data {
	return NotificationV2Data{
		NotificationV2DataBadgeEarned: v,
	}
}

// NotificationV2DataBoopAsNotificationV2Data is a convenience function that returns NotificationV2DataBoop wrapped in NotificationV2Data
func NotificationV2DataBoopAsNotificationV2Data(v *NotificationV2DataBoop) NotificationV2Data {
	return NotificationV2Data{
		NotificationV2DataBoop: v,
	}
}

// NotificationV2DataEventAnnouncementAsNotificationV2Data is a convenience function that returns NotificationV2DataEventAnnouncement wrapped in NotificationV2Data
func NotificationV2DataEventAnnouncementAsNotificationV2Data(v *NotificationV2DataEventAnnouncement) NotificationV2Data {
	return NotificationV2Data{
		NotificationV2DataEventAnnouncement: v,
	}
}

// NotificationV2DataGroupAnnouncementAsNotificationV2Data is a convenience function that returns NotificationV2DataGroupAnnouncement wrapped in NotificationV2Data
func NotificationV2DataGroupAnnouncementAsNotificationV2Data(v *NotificationV2DataGroupAnnouncement) NotificationV2Data {
	return NotificationV2Data{
		NotificationV2DataGroupAnnouncement: v,
	}
}

// NotificationV2DataGroupInformativeAsNotificationV2Data is a convenience function that returns NotificationV2DataGroupInformative wrapped in NotificationV2Data
func NotificationV2DataGroupInformativeAsNotificationV2Data(v *NotificationV2DataGroupInformative) NotificationV2Data {
	return NotificationV2Data{
		NotificationV2DataGroupInformative: v,
	}
}

// NotificationV2DataGroupTransferAsNotificationV2Data is a convenience function that returns NotificationV2DataGroupTransfer wrapped in NotificationV2Data
func NotificationV2DataGroupTransferAsNotificationV2Data(v *NotificationV2DataGroupTransfer) NotificationV2Data {
	return NotificationV2Data{
		NotificationV2DataGroupTransfer: v,
	}
}

// map[string]interface{}AsNotificationV2Data is a convenience function that returns map[string]interface{} wrapped in NotificationV2Data
func MapmapOfStringAnyAsNotificationV2Data(v *map[string]interface{}) NotificationV2Data {
	return NotificationV2Data{
		MapmapOfStringAny: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *NotificationV2Data) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into NotificationV2DataBadgeEarned
	err = newStrictDecoder(data).Decode(&dst.NotificationV2DataBadgeEarned)
	if err == nil {
		jsonNotificationV2DataBadgeEarned, _ := json.Marshal(dst.NotificationV2DataBadgeEarned)
		if string(jsonNotificationV2DataBadgeEarned) == "{}" { // empty struct
			dst.NotificationV2DataBadgeEarned = nil
		} else {
			if err = validator.Validate(dst.NotificationV2DataBadgeEarned); err != nil {
				dst.NotificationV2DataBadgeEarned = nil
			} else {
				match++
			}
		}
	} else {
		dst.NotificationV2DataBadgeEarned = nil
	}

	// try to unmarshal data into NotificationV2DataBoop
	err = newStrictDecoder(data).Decode(&dst.NotificationV2DataBoop)
	if err == nil {
		jsonNotificationV2DataBoop, _ := json.Marshal(dst.NotificationV2DataBoop)
		if string(jsonNotificationV2DataBoop) == "{}" { // empty struct
			dst.NotificationV2DataBoop = nil
		} else {
			if err = validator.Validate(dst.NotificationV2DataBoop); err != nil {
				dst.NotificationV2DataBoop = nil
			} else {
				match++
			}
		}
	} else {
		dst.NotificationV2DataBoop = nil
	}

	// try to unmarshal data into NotificationV2DataEventAnnouncement
	err = newStrictDecoder(data).Decode(&dst.NotificationV2DataEventAnnouncement)
	if err == nil {
		jsonNotificationV2DataEventAnnouncement, _ := json.Marshal(dst.NotificationV2DataEventAnnouncement)
		if string(jsonNotificationV2DataEventAnnouncement) == "{}" { // empty struct
			dst.NotificationV2DataEventAnnouncement = nil
		} else {
			if err = validator.Validate(dst.NotificationV2DataEventAnnouncement); err != nil {
				dst.NotificationV2DataEventAnnouncement = nil
			} else {
				match++
			}
		}
	} else {
		dst.NotificationV2DataEventAnnouncement = nil
	}

	// try to unmarshal data into NotificationV2DataGroupAnnouncement
	err = newStrictDecoder(data).Decode(&dst.NotificationV2DataGroupAnnouncement)
	if err == nil {
		jsonNotificationV2DataGroupAnnouncement, _ := json.Marshal(dst.NotificationV2DataGroupAnnouncement)
		if string(jsonNotificationV2DataGroupAnnouncement) == "{}" { // empty struct
			dst.NotificationV2DataGroupAnnouncement = nil
		} else {
			if err = validator.Validate(dst.NotificationV2DataGroupAnnouncement); err != nil {
				dst.NotificationV2DataGroupAnnouncement = nil
			} else {
				match++
			}
		}
	} else {
		dst.NotificationV2DataGroupAnnouncement = nil
	}

	// try to unmarshal data into NotificationV2DataGroupInformative
	err = newStrictDecoder(data).Decode(&dst.NotificationV2DataGroupInformative)
	if err == nil {
		jsonNotificationV2DataGroupInformative, _ := json.Marshal(dst.NotificationV2DataGroupInformative)
		if string(jsonNotificationV2DataGroupInformative) == "{}" { // empty struct
			dst.NotificationV2DataGroupInformative = nil
		} else {
			if err = validator.Validate(dst.NotificationV2DataGroupInformative); err != nil {
				dst.NotificationV2DataGroupInformative = nil
			} else {
				match++
			}
		}
	} else {
		dst.NotificationV2DataGroupInformative = nil
	}

	// try to unmarshal data into NotificationV2DataGroupTransfer
	err = newStrictDecoder(data).Decode(&dst.NotificationV2DataGroupTransfer)
	if err == nil {
		jsonNotificationV2DataGroupTransfer, _ := json.Marshal(dst.NotificationV2DataGroupTransfer)
		if string(jsonNotificationV2DataGroupTransfer) == "{}" { // empty struct
			dst.NotificationV2DataGroupTransfer = nil
		} else {
			if err = validator.Validate(dst.NotificationV2DataGroupTransfer); err != nil {
				dst.NotificationV2DataGroupTransfer = nil
			} else {
				match++
			}
		}
	} else {
		dst.NotificationV2DataGroupTransfer = nil
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
		dst.NotificationV2DataBadgeEarned = nil
		dst.NotificationV2DataBoop = nil
		dst.NotificationV2DataEventAnnouncement = nil
		dst.NotificationV2DataGroupAnnouncement = nil
		dst.NotificationV2DataGroupInformative = nil
		dst.NotificationV2DataGroupTransfer = nil
		dst.MapmapOfStringAny = nil

		return fmt.Errorf("data matches more than one schema in oneOf(NotificationV2Data)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(NotificationV2Data)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NotificationV2Data) MarshalJSON() ([]byte, error) {
	if src.NotificationV2DataBadgeEarned != nil {
		return json.Marshal(&src.NotificationV2DataBadgeEarned)
	}

	if src.NotificationV2DataBoop != nil {
		return json.Marshal(&src.NotificationV2DataBoop)
	}

	if src.NotificationV2DataEventAnnouncement != nil {
		return json.Marshal(&src.NotificationV2DataEventAnnouncement)
	}

	if src.NotificationV2DataGroupAnnouncement != nil {
		return json.Marshal(&src.NotificationV2DataGroupAnnouncement)
	}

	if src.NotificationV2DataGroupInformative != nil {
		return json.Marshal(&src.NotificationV2DataGroupInformative)
	}

	if src.NotificationV2DataGroupTransfer != nil {
		return json.Marshal(&src.NotificationV2DataGroupTransfer)
	}

	if src.MapmapOfStringAny != nil {
		return json.Marshal(&src.MapmapOfStringAny)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NotificationV2Data) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.NotificationV2DataBadgeEarned != nil {
		return obj.NotificationV2DataBadgeEarned
	}

	if obj.NotificationV2DataBoop != nil {
		return obj.NotificationV2DataBoop
	}

	if obj.NotificationV2DataEventAnnouncement != nil {
		return obj.NotificationV2DataEventAnnouncement
	}

	if obj.NotificationV2DataGroupAnnouncement != nil {
		return obj.NotificationV2DataGroupAnnouncement
	}

	if obj.NotificationV2DataGroupInformative != nil {
		return obj.NotificationV2DataGroupInformative
	}

	if obj.NotificationV2DataGroupTransfer != nil {
		return obj.NotificationV2DataGroupTransfer
	}

	if obj.MapmapOfStringAny != nil {
		return obj.MapmapOfStringAny
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj NotificationV2Data) GetActualInstanceValue() (interface{}) {
	if obj.NotificationV2DataBadgeEarned != nil {
		return *obj.NotificationV2DataBadgeEarned
	}

	if obj.NotificationV2DataBoop != nil {
		return *obj.NotificationV2DataBoop
	}

	if obj.NotificationV2DataEventAnnouncement != nil {
		return *obj.NotificationV2DataEventAnnouncement
	}

	if obj.NotificationV2DataGroupAnnouncement != nil {
		return *obj.NotificationV2DataGroupAnnouncement
	}

	if obj.NotificationV2DataGroupInformative != nil {
		return *obj.NotificationV2DataGroupInformative
	}

	if obj.NotificationV2DataGroupTransfer != nil {
		return *obj.NotificationV2DataGroupTransfer
	}

	if obj.MapmapOfStringAny != nil {
		return *obj.MapmapOfStringAny
	}

	// all schemas are nil
	return nil
}

type NullableNotificationV2Data struct {
	value *NotificationV2Data
	isSet bool
}

func (v NullableNotificationV2Data) Get() *NotificationV2Data {
	return v.value
}

func (v *NullableNotificationV2Data) Set(val *NotificationV2Data) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationV2Data) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationV2Data) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationV2Data(val *NotificationV2Data) *NullableNotificationV2Data {
	return &NullableNotificationV2Data{value: val, isSet: true}
}

func (v NullableNotificationV2Data) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationV2Data) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


