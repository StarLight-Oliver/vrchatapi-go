# NotificationDetailRequestInviteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InResponseTo** | **string** |  | 
**RequestMessage** | Pointer to **string** | Used when using InviteMessage Slot. | [optional] 

## Methods

### NewNotificationDetailRequestInviteResponse

`func NewNotificationDetailRequestInviteResponse(inResponseTo string, ) *NotificationDetailRequestInviteResponse`

NewNotificationDetailRequestInviteResponse instantiates a new NotificationDetailRequestInviteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationDetailRequestInviteResponseWithDefaults

`func NewNotificationDetailRequestInviteResponseWithDefaults() *NotificationDetailRequestInviteResponse`

NewNotificationDetailRequestInviteResponseWithDefaults instantiates a new NotificationDetailRequestInviteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInResponseTo

`func (o *NotificationDetailRequestInviteResponse) GetInResponseTo() string`

GetInResponseTo returns the InResponseTo field if non-nil, zero value otherwise.

### GetInResponseToOk

`func (o *NotificationDetailRequestInviteResponse) GetInResponseToOk() (*string, bool)`

GetInResponseToOk returns a tuple with the InResponseTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInResponseTo

`func (o *NotificationDetailRequestInviteResponse) SetInResponseTo(v string)`

SetInResponseTo sets InResponseTo field to given value.


### GetRequestMessage

`func (o *NotificationDetailRequestInviteResponse) GetRequestMessage() string`

GetRequestMessage returns the RequestMessage field if non-nil, zero value otherwise.

### GetRequestMessageOk

`func (o *NotificationDetailRequestInviteResponse) GetRequestMessageOk() (*string, bool)`

GetRequestMessageOk returns a tuple with the RequestMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMessage

`func (o *NotificationDetailRequestInviteResponse) SetRequestMessage(v string)`

SetRequestMessage sets RequestMessage field to given value.

### HasRequestMessage

`func (o *NotificationDetailRequestInviteResponse) HasRequestMessage() bool`

HasRequestMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


