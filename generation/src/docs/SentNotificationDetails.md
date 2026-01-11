# SentNotificationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmojiId** | Pointer to **string** |  | [optional] 
**EmojiVersion** | Pointer to **int32** |  | [optional] 
**InventoryItemId** | Pointer to **string** |  | [optional] 
**InviteMessage** | Pointer to **string** |  | [optional] 
**WorldId** | **string** | Represents a unique location, consisting of a world identifier and an instance identifier, or \&quot;offline\&quot; if the user is not on your friends list. | 
**WorldName** | **string** |  | 
**InResponseTo** | **string** |  | 
**ResponseMessage** | **string** |  | 
**Platform** | Pointer to **string** | This is normally &#x60;android&#x60;, &#x60;ios&#x60;, &#x60;standalonewindows&#x60;, &#x60;web&#x60;, or the empty value &#x60;&#x60;, but also supposedly can be any random Unity version such as &#x60;2019.2.4-801-Release&#x60; or &#x60;2019.2.2-772-Release&#x60; or even &#x60;unknownplatform&#x60;. | [optional] 
**RequestMessage** | Pointer to **string** | Used when using InviteMessage Slot. | [optional] 
**InitiatorUserId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**UserToKickId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 

## Methods

### NewSentNotificationDetails

`func NewSentNotificationDetails(worldId string, worldName string, inResponseTo string, responseMessage string, initiatorUserId string, userToKickId string, ) *SentNotificationDetails`

NewSentNotificationDetails instantiates a new SentNotificationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSentNotificationDetailsWithDefaults

`func NewSentNotificationDetailsWithDefaults() *SentNotificationDetails`

NewSentNotificationDetailsWithDefaults instantiates a new SentNotificationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmojiId

`func (o *SentNotificationDetails) GetEmojiId() string`

GetEmojiId returns the EmojiId field if non-nil, zero value otherwise.

### GetEmojiIdOk

`func (o *SentNotificationDetails) GetEmojiIdOk() (*string, bool)`

GetEmojiIdOk returns a tuple with the EmojiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiId

`func (o *SentNotificationDetails) SetEmojiId(v string)`

SetEmojiId sets EmojiId field to given value.

### HasEmojiId

`func (o *SentNotificationDetails) HasEmojiId() bool`

HasEmojiId returns a boolean if a field has been set.

### GetEmojiVersion

`func (o *SentNotificationDetails) GetEmojiVersion() int32`

GetEmojiVersion returns the EmojiVersion field if non-nil, zero value otherwise.

### GetEmojiVersionOk

`func (o *SentNotificationDetails) GetEmojiVersionOk() (*int32, bool)`

GetEmojiVersionOk returns a tuple with the EmojiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiVersion

`func (o *SentNotificationDetails) SetEmojiVersion(v int32)`

SetEmojiVersion sets EmojiVersion field to given value.

### HasEmojiVersion

`func (o *SentNotificationDetails) HasEmojiVersion() bool`

HasEmojiVersion returns a boolean if a field has been set.

### GetInventoryItemId

`func (o *SentNotificationDetails) GetInventoryItemId() string`

GetInventoryItemId returns the InventoryItemId field if non-nil, zero value otherwise.

### GetInventoryItemIdOk

`func (o *SentNotificationDetails) GetInventoryItemIdOk() (*string, bool)`

GetInventoryItemIdOk returns a tuple with the InventoryItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryItemId

`func (o *SentNotificationDetails) SetInventoryItemId(v string)`

SetInventoryItemId sets InventoryItemId field to given value.

### HasInventoryItemId

`func (o *SentNotificationDetails) HasInventoryItemId() bool`

HasInventoryItemId returns a boolean if a field has been set.

### GetInviteMessage

`func (o *SentNotificationDetails) GetInviteMessage() string`

GetInviteMessage returns the InviteMessage field if non-nil, zero value otherwise.

### GetInviteMessageOk

`func (o *SentNotificationDetails) GetInviteMessageOk() (*string, bool)`

GetInviteMessageOk returns a tuple with the InviteMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteMessage

`func (o *SentNotificationDetails) SetInviteMessage(v string)`

SetInviteMessage sets InviteMessage field to given value.

### HasInviteMessage

`func (o *SentNotificationDetails) HasInviteMessage() bool`

HasInviteMessage returns a boolean if a field has been set.

### GetWorldId

`func (o *SentNotificationDetails) GetWorldId() string`

GetWorldId returns the WorldId field if non-nil, zero value otherwise.

### GetWorldIdOk

`func (o *SentNotificationDetails) GetWorldIdOk() (*string, bool)`

GetWorldIdOk returns a tuple with the WorldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorldId

`func (o *SentNotificationDetails) SetWorldId(v string)`

SetWorldId sets WorldId field to given value.


### GetWorldName

`func (o *SentNotificationDetails) GetWorldName() string`

GetWorldName returns the WorldName field if non-nil, zero value otherwise.

### GetWorldNameOk

`func (o *SentNotificationDetails) GetWorldNameOk() (*string, bool)`

GetWorldNameOk returns a tuple with the WorldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorldName

`func (o *SentNotificationDetails) SetWorldName(v string)`

SetWorldName sets WorldName field to given value.


### GetInResponseTo

`func (o *SentNotificationDetails) GetInResponseTo() string`

GetInResponseTo returns the InResponseTo field if non-nil, zero value otherwise.

### GetInResponseToOk

`func (o *SentNotificationDetails) GetInResponseToOk() (*string, bool)`

GetInResponseToOk returns a tuple with the InResponseTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInResponseTo

`func (o *SentNotificationDetails) SetInResponseTo(v string)`

SetInResponseTo sets InResponseTo field to given value.


### GetResponseMessage

`func (o *SentNotificationDetails) GetResponseMessage() string`

GetResponseMessage returns the ResponseMessage field if non-nil, zero value otherwise.

### GetResponseMessageOk

`func (o *SentNotificationDetails) GetResponseMessageOk() (*string, bool)`

GetResponseMessageOk returns a tuple with the ResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessage

`func (o *SentNotificationDetails) SetResponseMessage(v string)`

SetResponseMessage sets ResponseMessage field to given value.


### GetPlatform

`func (o *SentNotificationDetails) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *SentNotificationDetails) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *SentNotificationDetails) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *SentNotificationDetails) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetRequestMessage

`func (o *SentNotificationDetails) GetRequestMessage() string`

GetRequestMessage returns the RequestMessage field if non-nil, zero value otherwise.

### GetRequestMessageOk

`func (o *SentNotificationDetails) GetRequestMessageOk() (*string, bool)`

GetRequestMessageOk returns a tuple with the RequestMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMessage

`func (o *SentNotificationDetails) SetRequestMessage(v string)`

SetRequestMessage sets RequestMessage field to given value.

### HasRequestMessage

`func (o *SentNotificationDetails) HasRequestMessage() bool`

HasRequestMessage returns a boolean if a field has been set.

### GetInitiatorUserId

`func (o *SentNotificationDetails) GetInitiatorUserId() string`

GetInitiatorUserId returns the InitiatorUserId field if non-nil, zero value otherwise.

### GetInitiatorUserIdOk

`func (o *SentNotificationDetails) GetInitiatorUserIdOk() (*string, bool)`

GetInitiatorUserIdOk returns a tuple with the InitiatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorUserId

`func (o *SentNotificationDetails) SetInitiatorUserId(v string)`

SetInitiatorUserId sets InitiatorUserId field to given value.


### GetUserToKickId

`func (o *SentNotificationDetails) GetUserToKickId() string`

GetUserToKickId returns the UserToKickId field if non-nil, zero value otherwise.

### GetUserToKickIdOk

`func (o *SentNotificationDetails) GetUserToKickIdOk() (*string, bool)`

GetUserToKickIdOk returns a tuple with the UserToKickId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToKickId

`func (o *SentNotificationDetails) SetUserToKickId(v string)`

SetUserToKickId sets UserToKickId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


