# NotificationDetailVoteToKick

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InitiatorUserId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**UserToKickId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 

## Methods

### NewNotificationDetailVoteToKick

`func NewNotificationDetailVoteToKick(initiatorUserId string, userToKickId string, ) *NotificationDetailVoteToKick`

NewNotificationDetailVoteToKick instantiates a new NotificationDetailVoteToKick object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationDetailVoteToKickWithDefaults

`func NewNotificationDetailVoteToKickWithDefaults() *NotificationDetailVoteToKick`

NewNotificationDetailVoteToKickWithDefaults instantiates a new NotificationDetailVoteToKick object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitiatorUserId

`func (o *NotificationDetailVoteToKick) GetInitiatorUserId() string`

GetInitiatorUserId returns the InitiatorUserId field if non-nil, zero value otherwise.

### GetInitiatorUserIdOk

`func (o *NotificationDetailVoteToKick) GetInitiatorUserIdOk() (*string, bool)`

GetInitiatorUserIdOk returns a tuple with the InitiatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorUserId

`func (o *NotificationDetailVoteToKick) SetInitiatorUserId(v string)`

SetInitiatorUserId sets InitiatorUserId field to given value.


### GetUserToKickId

`func (o *NotificationDetailVoteToKick) GetUserToKickId() string`

GetUserToKickId returns the UserToKickId field if non-nil, zero value otherwise.

### GetUserToKickIdOk

`func (o *NotificationDetailVoteToKick) GetUserToKickIdOk() (*string, bool)`

GetUserToKickIdOk returns a tuple with the UserToKickId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToKickId

`func (o *NotificationDetailVoteToKick) SetUserToKickId(v string)`

SetUserToKickId sets UserToKickId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


