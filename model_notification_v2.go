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
	"time"
)

// checks if the NotificationV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationV2{}

// NotificationV2
type NotificationV2 struct {
	CanDelete       bool                       `json:"canDelete"`
	Category        string                     `json:"category"`
	CreatedAt       time.Time                  `json:"createdAt"`
	Data            NotificationV2Data         `json:"data"`
	Details         *NotificationV2DetailsBoop `json:"details,omitempty"`
	ExpiresAt       time.Time                  `json:"expiresAt"`
	ExpiryAfterSeen NullableInt32              `json:"expiryAfterSeen"`
	Id              string                     `json:"id"`
	IgnoreDND       bool                       `json:"ignoreDND"`
	ImageUrl        NullableString             `json:"imageUrl"`
	IsSystem        bool                       `json:"isSystem"`
	Link            string                     `json:"link"`
	LinkText        string                     `json:"linkText"`
	LinkTextKey     NullableString             `json:"linkTextKey"`
	Message         string                     `json:"message"`
	MessageKey      NullableString             `json:"messageKey,omitempty"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	ReceiverUserId         string                   `json:"receiverUserId"`
	RelatedNotificationsId NullableString           `json:"relatedNotificationsId"`
	RequireSeen            bool                     `json:"requireSeen"`
	Responses              []NotificationV2Response `json:"responses"`
	Seen                   bool                     `json:"seen"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	SenderUserId   string             `json:"senderUserId"`
	SenderUsername NullableString     `json:"senderUsername"`
	Title          string             `json:"title"`
	TitleKey       NullableString     `json:"titleKey"`
	Type           NotificationV2Type `json:"type"`
	UpdatedAt      time.Time          `json:"updatedAt"`
	Version        int32              `json:"version"`
}

type _NotificationV2 NotificationV2

// NewNotificationV2 instantiates a new NotificationV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationV2(canDelete bool, category string, createdAt time.Time, data NotificationV2Data, expiresAt time.Time, expiryAfterSeen NullableInt32, id string, ignoreDND bool, imageUrl NullableString, isSystem bool, link string, linkText string, linkTextKey NullableString, message string, receiverUserId string, relatedNotificationsId NullableString, requireSeen bool, responses []NotificationV2Response, seen bool, senderUserId string, senderUsername NullableString, title string, titleKey NullableString, type_ NotificationV2Type, updatedAt time.Time, version int32) *NotificationV2 {
	this := NotificationV2{}
	this.CanDelete = canDelete
	this.Category = category
	this.CreatedAt = createdAt
	this.Data = data
	this.ExpiresAt = expiresAt
	this.ExpiryAfterSeen = expiryAfterSeen
	this.Id = id
	this.IgnoreDND = ignoreDND
	this.ImageUrl = imageUrl
	this.IsSystem = isSystem
	this.Link = link
	this.LinkText = linkText
	this.LinkTextKey = linkTextKey
	this.Message = message
	this.ReceiverUserId = receiverUserId
	this.RelatedNotificationsId = relatedNotificationsId
	this.RequireSeen = requireSeen
	this.Responses = responses
	this.Seen = seen
	this.SenderUserId = senderUserId
	this.SenderUsername = senderUsername
	this.Title = title
	this.TitleKey = titleKey
	this.Type = type_
	this.UpdatedAt = updatedAt
	this.Version = version
	return &this
}

// NewNotificationV2WithDefaults instantiates a new NotificationV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationV2WithDefaults() *NotificationV2 {
	this := NotificationV2{}
	var type_ NotificationV2Type = NotificationV2Type_GROUP_ANNOUNCEMENT
	this.Type = type_
	var version int32 = 2
	this.Version = version
	return &this
}

// GetCanDelete returns the CanDelete field value
func (o *NotificationV2) GetCanDelete() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanDelete
}

// GetCanDeleteOk returns a tuple with the CanDelete field value
// and a boolean to check if the value has been set.
func (o *NotificationV2) GetCanDeleteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanDelete, true
}

// SetCanDelete sets field value
func (o *NotificationV2) SetCanDelete(v bool) {
	o.CanDelete = v
}

// GetCategory returns the Category field value
func (o *NotificationV2) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *NotificationV2) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *NotificationV2) SetCategory(v string) {
	o.Category = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *NotificationV2) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *NotificationV2) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *NotificationV2) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetData returns the Data field value
func (o *NotificationV2) GetData() NotificationV2Data {
	if o == nil {
		var ret NotificationV2Data
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *NotificationV2) GetDataOk() (*NotificationV2Data, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *NotificationV2) SetData(v NotificationV2Data) {
	o.Data = v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *NotificationV2) GetDetails() NotificationV2DetailsBoop {
	if o == nil || IsNil(o.Details) {
		var ret NotificationV2DetailsBoop
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationV2) GetDetailsOk() (*NotificationV2DetailsBoop, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *NotificationV2) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given NotificationV2DetailsBoop and assigns it to the Details field.
func (o *NotificationV2) SetDetails(v NotificationV2DetailsBoop) {
	o.Details = &v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *NotificationV2) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *NotificationV2) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *NotificationV2) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

// GetExpiryAfterSeen returns the ExpiryAfterSeen field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *NotificationV2) GetExpiryAfterSeen() int32 {
	if o == nil || o.ExpiryAfterSeen.Get() == nil {
		var ret int32
		return ret
	}

	return *o.ExpiryAfterSeen.Get()
}

// GetExpiryAfterSeenOk returns a tuple with the ExpiryAfterSeen field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationV2) GetExpiryAfterSeenOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiryAfterSeen.Get(), o.ExpiryAfterSeen.IsSet()
}

// SetExpiryAfterSeen sets field value
func (o *NotificationV2) SetExpiryAfterSeen(v int32) {
	o.ExpiryAfterSeen.Set(&v)
}

// GetId returns the Id field value
func (o *NotificationV2) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NotificationV2) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NotificationV2) SetId(v string) {
	o.Id = v
}

// GetIgnoreDND returns the IgnoreDND field value
func (o *NotificationV2) GetIgnoreDND() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IgnoreDND
}

// GetIgnoreDNDOk returns a tuple with the IgnoreDND field value
// and a boolean to check if the value has been set.
func (o *NotificationV2) GetIgnoreDNDOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IgnoreDND, true
}

// SetIgnoreDND sets field value
func (o *NotificationV2) SetIgnoreDND(v bool) {
	o.IgnoreDND = v
}

// GetImageUrl returns the ImageUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NotificationV2) GetImageUrl() string {
	if o == nil || o.ImageUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.ImageUrl.Get()
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationV2) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageUrl.Get(), o.ImageUrl.IsSet()
}

// SetImageUrl sets field value
func (o *NotificationV2) SetImageUrl(v string) {
	o.ImageUrl.Set(&v)
}

// GetIsSystem returns the IsSystem field value
func (o *NotificationV2) GetIsSystem() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSystem
}

// GetIsSystemOk returns a tuple with the IsSystem field value
// and a boolean to check if the value has been set.
func (o *NotificationV2) GetIsSystemOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSystem, true
}

// SetIsSystem sets field value
func (o *NotificationV2) SetIsSystem(v bool) {
	o.IsSystem = v
}

// GetLink returns the Link field value
func (o *NotificationV2) GetLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Link
}

// GetLinkOk returns a tuple with the Link field value
// and a boolean to check if the value has been set.
func (o *NotificationV2) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Link, true
}

// SetLink sets field value
func (o *NotificationV2) SetLink(v string) {
	o.Link = v
}

// GetLinkText returns the LinkText field value
func (o *NotificationV2) GetLinkText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkText
}

// GetLinkTextOk returns a tuple with the LinkText field value
// and a boolean to check if the value has been set.
func (o *NotificationV2) GetLinkTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkText, true
}

// SetLinkText sets field value
func (o *NotificationV2) SetLinkText(v string) {
	o.LinkText = v
}

// GetLinkTextKey returns the LinkTextKey field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NotificationV2) GetLinkTextKey() string {
	if o == nil || o.LinkTextKey.Get() == nil {
		var ret string
		return ret
	}

	return *o.LinkTextKey.Get()
}

// GetLinkTextKeyOk returns a tuple with the LinkTextKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationV2) GetLinkTextKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LinkTextKey.Get(), o.LinkTextKey.IsSet()
}

// SetLinkTextKey sets field value
func (o *NotificationV2) SetLinkTextKey(v string) {
	o.LinkTextKey.Set(&v)
}

// GetMessage returns the Message field value
func (o *NotificationV2) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *NotificationV2) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *NotificationV2) SetMessage(v string) {
	o.Message = v
}

// GetMessageKey returns the MessageKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NotificationV2) GetMessageKey() string {
	if o == nil || IsNil(o.MessageKey.Get()) {
		var ret string
		return ret
	}
	return *o.MessageKey.Get()
}

// GetMessageKeyOk returns a tuple with the MessageKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationV2) GetMessageKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MessageKey.Get(), o.MessageKey.IsSet()
}

// HasMessageKey returns a boolean if a field has been set.
func (o *NotificationV2) HasMessageKey() bool {
	if o != nil && o.MessageKey.IsSet() {
		return true
	}

	return false
}

// SetMessageKey gets a reference to the given NullableString and assigns it to the MessageKey field.
func (o *NotificationV2) SetMessageKey(v string) {
	o.MessageKey.Set(&v)
}

// SetMessageKeyNil sets the value for MessageKey to be an explicit nil
func (o *NotificationV2) SetMessageKeyNil() {
	o.MessageKey.Set(nil)
}

// UnsetMessageKey ensures that no value is present for MessageKey, not even an explicit nil
func (o *NotificationV2) UnsetMessageKey() {
	o.MessageKey.Unset()
}

// GetReceiverUserId returns the ReceiverUserId field value
func (o *NotificationV2) GetReceiverUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReceiverUserId
}

// GetReceiverUserIdOk returns a tuple with the ReceiverUserId field value
// and a boolean to check if the value has been set.
func (o *NotificationV2) GetReceiverUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceiverUserId, true
}

// SetReceiverUserId sets field value
func (o *NotificationV2) SetReceiverUserId(v string) {
	o.ReceiverUserId = v
}

// GetRelatedNotificationsId returns the RelatedNotificationsId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NotificationV2) GetRelatedNotificationsId() string {
	if o == nil || o.RelatedNotificationsId.Get() == nil {
		var ret string
		return ret
	}

	return *o.RelatedNotificationsId.Get()
}

// GetRelatedNotificationsIdOk returns a tuple with the RelatedNotificationsId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationV2) GetRelatedNotificationsIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RelatedNotificationsId.Get(), o.RelatedNotificationsId.IsSet()
}

// SetRelatedNotificationsId sets field value
func (o *NotificationV2) SetRelatedNotificationsId(v string) {
	o.RelatedNotificationsId.Set(&v)
}

// GetRequireSeen returns the RequireSeen field value
func (o *NotificationV2) GetRequireSeen() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RequireSeen
}

// GetRequireSeenOk returns a tuple with the RequireSeen field value
// and a boolean to check if the value has been set.
func (o *NotificationV2) GetRequireSeenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequireSeen, true
}

// SetRequireSeen sets field value
func (o *NotificationV2) SetRequireSeen(v bool) {
	o.RequireSeen = v
}

// GetResponses returns the Responses field value
func (o *NotificationV2) GetResponses() []NotificationV2Response {
	if o == nil {
		var ret []NotificationV2Response
		return ret
	}

	return o.Responses
}

// GetResponsesOk returns a tuple with the Responses field value
// and a boolean to check if the value has been set.
func (o *NotificationV2) GetResponsesOk() ([]NotificationV2Response, bool) {
	if o == nil {
		return nil, false
	}
	return o.Responses, true
}

// SetResponses sets field value
func (o *NotificationV2) SetResponses(v []NotificationV2Response) {
	o.Responses = v
}

// GetSeen returns the Seen field value
func (o *NotificationV2) GetSeen() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Seen
}

// GetSeenOk returns a tuple with the Seen field value
// and a boolean to check if the value has been set.
func (o *NotificationV2) GetSeenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Seen, true
}

// SetSeen sets field value
func (o *NotificationV2) SetSeen(v bool) {
	o.Seen = v
}

// GetSenderUserId returns the SenderUserId field value
func (o *NotificationV2) GetSenderUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SenderUserId
}

// GetSenderUserIdOk returns a tuple with the SenderUserId field value
// and a boolean to check if the value has been set.
func (o *NotificationV2) GetSenderUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderUserId, true
}

// SetSenderUserId sets field value
func (o *NotificationV2) SetSenderUserId(v string) {
	o.SenderUserId = v
}

// GetSenderUsername returns the SenderUsername field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NotificationV2) GetSenderUsername() string {
	if o == nil || o.SenderUsername.Get() == nil {
		var ret string
		return ret
	}

	return *o.SenderUsername.Get()
}

// GetSenderUsernameOk returns a tuple with the SenderUsername field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationV2) GetSenderUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SenderUsername.Get(), o.SenderUsername.IsSet()
}

// SetSenderUsername sets field value
func (o *NotificationV2) SetSenderUsername(v string) {
	o.SenderUsername.Set(&v)
}

// GetTitle returns the Title field value
func (o *NotificationV2) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *NotificationV2) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *NotificationV2) SetTitle(v string) {
	o.Title = v
}

// GetTitleKey returns the TitleKey field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NotificationV2) GetTitleKey() string {
	if o == nil || o.TitleKey.Get() == nil {
		var ret string
		return ret
	}

	return *o.TitleKey.Get()
}

// GetTitleKeyOk returns a tuple with the TitleKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NotificationV2) GetTitleKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TitleKey.Get(), o.TitleKey.IsSet()
}

// SetTitleKey sets field value
func (o *NotificationV2) SetTitleKey(v string) {
	o.TitleKey.Set(&v)
}

// GetType returns the Type field value
func (o *NotificationV2) GetType() NotificationV2Type {
	if o == nil {
		var ret NotificationV2Type
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NotificationV2) GetTypeOk() (*NotificationV2Type, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NotificationV2) SetType(v NotificationV2Type) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *NotificationV2) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *NotificationV2) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *NotificationV2) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetVersion returns the Version field value
func (o *NotificationV2) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *NotificationV2) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *NotificationV2) SetVersion(v int32) {
	o.Version = v
}

func (o NotificationV2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["canDelete"] = o.CanDelete
	toSerialize["category"] = o.Category
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["data"] = o.Data
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	toSerialize["expiresAt"] = o.ExpiresAt
	toSerialize["expiryAfterSeen"] = o.ExpiryAfterSeen.Get()
	toSerialize["id"] = o.Id
	toSerialize["ignoreDND"] = o.IgnoreDND
	toSerialize["imageUrl"] = o.ImageUrl.Get()
	toSerialize["isSystem"] = o.IsSystem
	toSerialize["link"] = o.Link
	toSerialize["linkText"] = o.LinkText
	toSerialize["linkTextKey"] = o.LinkTextKey.Get()
	toSerialize["message"] = o.Message
	if o.MessageKey.IsSet() {
		toSerialize["messageKey"] = o.MessageKey.Get()
	}
	toSerialize["receiverUserId"] = o.ReceiverUserId
	toSerialize["relatedNotificationsId"] = o.RelatedNotificationsId.Get()
	toSerialize["requireSeen"] = o.RequireSeen
	toSerialize["responses"] = o.Responses
	toSerialize["seen"] = o.Seen
	toSerialize["senderUserId"] = o.SenderUserId
	toSerialize["senderUsername"] = o.SenderUsername.Get()
	toSerialize["title"] = o.Title
	toSerialize["titleKey"] = o.TitleKey.Get()
	toSerialize["type"] = o.Type
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

func (o *NotificationV2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"canDelete",
		"category",
		"createdAt",
		"data",
		"expiresAt",
		"expiryAfterSeen",
		"id",
		"ignoreDND",
		"imageUrl",
		"isSystem",
		"link",
		"linkText",
		"linkTextKey",
		"message",
		"receiverUserId",
		"relatedNotificationsId",
		"requireSeen",
		"responses",
		"seen",
		"senderUserId",
		"senderUsername",
		"title",
		"titleKey",
		"type",
		"updatedAt",
		"version",
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

	varNotificationV2 := _NotificationV2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNotificationV2)

	if err != nil {
		return err
	}

	*o = NotificationV2(varNotificationV2)

	return err
}

type NullableNotificationV2 struct {
	value *NotificationV2
	isSet bool
}

func (v NullableNotificationV2) Get() *NotificationV2 {
	return v.value
}

func (v *NullableNotificationV2) Set(val *NotificationV2) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationV2) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationV2(val *NotificationV2) *NullableNotificationV2 {
	return &NullableNotificationV2{value: val, isSet: true}
}

func (v NullableNotificationV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
