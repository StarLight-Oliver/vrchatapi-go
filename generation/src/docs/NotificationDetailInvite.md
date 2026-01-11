# NotificationDetailInvite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InviteMessage** | Pointer to **string** |  | [optional] 
**WorldId** | **string** | Represents a unique location, consisting of a world identifier and an instance identifier, or \&quot;offline\&quot; if the user is not on your friends list. | 
**WorldName** | **string** |  | 

## Methods

### NewNotificationDetailInvite

`func NewNotificationDetailInvite(worldId string, worldName string, ) *NotificationDetailInvite`

NewNotificationDetailInvite instantiates a new NotificationDetailInvite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationDetailInviteWithDefaults

`func NewNotificationDetailInviteWithDefaults() *NotificationDetailInvite`

NewNotificationDetailInviteWithDefaults instantiates a new NotificationDetailInvite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInviteMessage

`func (o *NotificationDetailInvite) GetInviteMessage() string`

GetInviteMessage returns the InviteMessage field if non-nil, zero value otherwise.

### GetInviteMessageOk

`func (o *NotificationDetailInvite) GetInviteMessageOk() (*string, bool)`

GetInviteMessageOk returns a tuple with the InviteMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteMessage

`func (o *NotificationDetailInvite) SetInviteMessage(v string)`

SetInviteMessage sets InviteMessage field to given value.

### HasInviteMessage

`func (o *NotificationDetailInvite) HasInviteMessage() bool`

HasInviteMessage returns a boolean if a field has been set.

### GetWorldId

`func (o *NotificationDetailInvite) GetWorldId() string`

GetWorldId returns the WorldId field if non-nil, zero value otherwise.

### GetWorldIdOk

`func (o *NotificationDetailInvite) GetWorldIdOk() (*string, bool)`

GetWorldIdOk returns a tuple with the WorldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorldId

`func (o *NotificationDetailInvite) SetWorldId(v string)`

SetWorldId sets WorldId field to given value.


### GetWorldName

`func (o *NotificationDetailInvite) GetWorldName() string`

GetWorldName returns the WorldName field if non-nil, zero value otherwise.

### GetWorldNameOk

`func (o *NotificationDetailInvite) GetWorldNameOk() (*string, bool)`

GetWorldNameOk returns a tuple with the WorldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorldName

`func (o *NotificationDetailInvite) SetWorldName(v string)`

SetWorldName sets WorldName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


