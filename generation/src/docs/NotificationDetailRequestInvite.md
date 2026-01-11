# NotificationDetailRequestInvite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | Pointer to **string** | This is normally &#x60;android&#x60;, &#x60;ios&#x60;, &#x60;standalonewindows&#x60;, &#x60;web&#x60;, or the empty value &#x60;&#x60;, but also supposedly can be any random Unity version such as &#x60;2019.2.4-801-Release&#x60; or &#x60;2019.2.2-772-Release&#x60; or even &#x60;unknownplatform&#x60;. | [optional] 
**RequestMessage** | Pointer to **string** | Used when using InviteMessage Slot. | [optional] 

## Methods

### NewNotificationDetailRequestInvite

`func NewNotificationDetailRequestInvite() *NotificationDetailRequestInvite`

NewNotificationDetailRequestInvite instantiates a new NotificationDetailRequestInvite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationDetailRequestInviteWithDefaults

`func NewNotificationDetailRequestInviteWithDefaults() *NotificationDetailRequestInvite`

NewNotificationDetailRequestInviteWithDefaults instantiates a new NotificationDetailRequestInvite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *NotificationDetailRequestInvite) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *NotificationDetailRequestInvite) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *NotificationDetailRequestInvite) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *NotificationDetailRequestInvite) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetRequestMessage

`func (o *NotificationDetailRequestInvite) GetRequestMessage() string`

GetRequestMessage returns the RequestMessage field if non-nil, zero value otherwise.

### GetRequestMessageOk

`func (o *NotificationDetailRequestInvite) GetRequestMessageOk() (*string, bool)`

GetRequestMessageOk returns a tuple with the RequestMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMessage

`func (o *NotificationDetailRequestInvite) SetRequestMessage(v string)`

SetRequestMessage sets RequestMessage field to given value.

### HasRequestMessage

`func (o *NotificationDetailRequestInvite) HasRequestMessage() bool`

HasRequestMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


