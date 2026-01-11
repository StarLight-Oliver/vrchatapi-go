# BoopRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmojiId** | Pointer to **string** | Either a FileID or a string constant for default emojis | [optional] 
**EmojiVersion** | Pointer to **int32** |  | [optional] 
**InventoryItemId** | Pointer to **string** |  | [optional] 

## Methods

### NewBoopRequest

`func NewBoopRequest() *BoopRequest`

NewBoopRequest instantiates a new BoopRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoopRequestWithDefaults

`func NewBoopRequestWithDefaults() *BoopRequest`

NewBoopRequestWithDefaults instantiates a new BoopRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmojiId

`func (o *BoopRequest) GetEmojiId() string`

GetEmojiId returns the EmojiId field if non-nil, zero value otherwise.

### GetEmojiIdOk

`func (o *BoopRequest) GetEmojiIdOk() (*string, bool)`

GetEmojiIdOk returns a tuple with the EmojiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiId

`func (o *BoopRequest) SetEmojiId(v string)`

SetEmojiId sets EmojiId field to given value.

### HasEmojiId

`func (o *BoopRequest) HasEmojiId() bool`

HasEmojiId returns a boolean if a field has been set.

### GetEmojiVersion

`func (o *BoopRequest) GetEmojiVersion() int32`

GetEmojiVersion returns the EmojiVersion field if non-nil, zero value otherwise.

### GetEmojiVersionOk

`func (o *BoopRequest) GetEmojiVersionOk() (*int32, bool)`

GetEmojiVersionOk returns a tuple with the EmojiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiVersion

`func (o *BoopRequest) SetEmojiVersion(v int32)`

SetEmojiVersion sets EmojiVersion field to given value.

### HasEmojiVersion

`func (o *BoopRequest) HasEmojiVersion() bool`

HasEmojiVersion returns a boolean if a field has been set.

### GetInventoryItemId

`func (o *BoopRequest) GetInventoryItemId() string`

GetInventoryItemId returns the InventoryItemId field if non-nil, zero value otherwise.

### GetInventoryItemIdOk

`func (o *BoopRequest) GetInventoryItemIdOk() (*string, bool)`

GetInventoryItemIdOk returns a tuple with the InventoryItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryItemId

`func (o *BoopRequest) SetInventoryItemId(v string)`

SetInventoryItemId sets InventoryItemId field to given value.

### HasInventoryItemId

`func (o *BoopRequest) HasInventoryItemId() bool`

HasInventoryItemId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


