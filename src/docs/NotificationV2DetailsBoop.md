# NotificationV2DetailsBoop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmojiId** | **string** | Either a FileID or a string constant for default emojis | 
**EmojiVersion** | **NullableInt32** |  | 
**InventoryItemId** | **string** |  | 

## Methods

### NewNotificationV2DetailsBoop

`func NewNotificationV2DetailsBoop(emojiId string, emojiVersion NullableInt32, inventoryItemId string, ) *NotificationV2DetailsBoop`

NewNotificationV2DetailsBoop instantiates a new NotificationV2DetailsBoop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationV2DetailsBoopWithDefaults

`func NewNotificationV2DetailsBoopWithDefaults() *NotificationV2DetailsBoop`

NewNotificationV2DetailsBoopWithDefaults instantiates a new NotificationV2DetailsBoop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmojiId

`func (o *NotificationV2DetailsBoop) GetEmojiId() string`

GetEmojiId returns the EmojiId field if non-nil, zero value otherwise.

### GetEmojiIdOk

`func (o *NotificationV2DetailsBoop) GetEmojiIdOk() (*string, bool)`

GetEmojiIdOk returns a tuple with the EmojiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiId

`func (o *NotificationV2DetailsBoop) SetEmojiId(v string)`

SetEmojiId sets EmojiId field to given value.


### GetEmojiVersion

`func (o *NotificationV2DetailsBoop) GetEmojiVersion() int32`

GetEmojiVersion returns the EmojiVersion field if non-nil, zero value otherwise.

### GetEmojiVersionOk

`func (o *NotificationV2DetailsBoop) GetEmojiVersionOk() (*int32, bool)`

GetEmojiVersionOk returns a tuple with the EmojiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiVersion

`func (o *NotificationV2DetailsBoop) SetEmojiVersion(v int32)`

SetEmojiVersion sets EmojiVersion field to given value.


### SetEmojiVersionNil

`func (o *NotificationV2DetailsBoop) SetEmojiVersionNil(b bool)`

 SetEmojiVersionNil sets the value for EmojiVersion to be an explicit nil

### UnsetEmojiVersion
`func (o *NotificationV2DetailsBoop) UnsetEmojiVersion()`

UnsetEmojiVersion ensures that no value is present for EmojiVersion, not even an explicit nil
### GetInventoryItemId

`func (o *NotificationV2DetailsBoop) GetInventoryItemId() string`

GetInventoryItemId returns the InventoryItemId field if non-nil, zero value otherwise.

### GetInventoryItemIdOk

`func (o *NotificationV2DetailsBoop) GetInventoryItemIdOk() (*string, bool)`

GetInventoryItemIdOk returns a tuple with the InventoryItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryItemId

`func (o *NotificationV2DetailsBoop) SetInventoryItemId(v string)`

SetInventoryItemId sets InventoryItemId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


