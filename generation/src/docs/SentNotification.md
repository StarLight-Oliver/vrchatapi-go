# SentNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**Details** | [**SentNotificationDetails**](SentNotificationDetails.md) |  | 
**Id** | **string** |  | 
**Message** | **string** |  | 
**ReceiverUserId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**SenderUserId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**SenderUsername** | Pointer to **string** | -| **DEPRECATED:** VRChat API no longer return usernames of other users. [See issue by Tupper for more information](https://github.com/pypy-vrc/VRCX/issues/429). | [optional] 
**Type** | [**NotificationType**](NotificationType.md) |  | [default to FRIEND_REQUEST]

## Methods

### NewSentNotification

`func NewSentNotification(createdAt time.Time, details SentNotificationDetails, id string, message string, receiverUserId string, senderUserId string, type_ NotificationType, ) *SentNotification`

NewSentNotification instantiates a new SentNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSentNotificationWithDefaults

`func NewSentNotificationWithDefaults() *SentNotification`

NewSentNotificationWithDefaults instantiates a new SentNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *SentNotification) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SentNotification) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SentNotification) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDetails

`func (o *SentNotification) GetDetails() SentNotificationDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *SentNotification) GetDetailsOk() (*SentNotificationDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *SentNotification) SetDetails(v SentNotificationDetails)`

SetDetails sets Details field to given value.


### GetId

`func (o *SentNotification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SentNotification) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SentNotification) SetId(v string)`

SetId sets Id field to given value.


### GetMessage

`func (o *SentNotification) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SentNotification) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SentNotification) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetReceiverUserId

`func (o *SentNotification) GetReceiverUserId() string`

GetReceiverUserId returns the ReceiverUserId field if non-nil, zero value otherwise.

### GetReceiverUserIdOk

`func (o *SentNotification) GetReceiverUserIdOk() (*string, bool)`

GetReceiverUserIdOk returns a tuple with the ReceiverUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverUserId

`func (o *SentNotification) SetReceiverUserId(v string)`

SetReceiverUserId sets ReceiverUserId field to given value.


### GetSenderUserId

`func (o *SentNotification) GetSenderUserId() string`

GetSenderUserId returns the SenderUserId field if non-nil, zero value otherwise.

### GetSenderUserIdOk

`func (o *SentNotification) GetSenderUserIdOk() (*string, bool)`

GetSenderUserIdOk returns a tuple with the SenderUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderUserId

`func (o *SentNotification) SetSenderUserId(v string)`

SetSenderUserId sets SenderUserId field to given value.


### GetSenderUsername

`func (o *SentNotification) GetSenderUsername() string`

GetSenderUsername returns the SenderUsername field if non-nil, zero value otherwise.

### GetSenderUsernameOk

`func (o *SentNotification) GetSenderUsernameOk() (*string, bool)`

GetSenderUsernameOk returns a tuple with the SenderUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderUsername

`func (o *SentNotification) SetSenderUsername(v string)`

SetSenderUsername sets SenderUsername field to given value.

### HasSenderUsername

`func (o *SentNotification) HasSenderUsername() bool`

HasSenderUsername returns a boolean if a field has been set.

### GetType

`func (o *SentNotification) GetType() NotificationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SentNotification) GetTypeOk() (*NotificationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SentNotification) SetType(v NotificationType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


