# CreateGroupPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageId** | Pointer to **string** |  | [optional] 
**RoleIds** | Pointer to **[]string** |   | [optional] 
**SendNotification** | **bool** | Send notification to group members. | [default to false]
**Text** | **string** | Post text | 
**Title** | **string** | Post title | 
**Visibility** | [**GroupPostVisibility**](GroupPostVisibility.md) |  | 

## Methods

### NewCreateGroupPostRequest

`func NewCreateGroupPostRequest(sendNotification bool, text string, title string, visibility GroupPostVisibility, ) *CreateGroupPostRequest`

NewCreateGroupPostRequest instantiates a new CreateGroupPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupPostRequestWithDefaults

`func NewCreateGroupPostRequestWithDefaults() *CreateGroupPostRequest`

NewCreateGroupPostRequestWithDefaults instantiates a new CreateGroupPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageId

`func (o *CreateGroupPostRequest) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *CreateGroupPostRequest) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *CreateGroupPostRequest) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *CreateGroupPostRequest) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetRoleIds

`func (o *CreateGroupPostRequest) GetRoleIds() []string`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *CreateGroupPostRequest) GetRoleIdsOk() (*[]string, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *CreateGroupPostRequest) SetRoleIds(v []string)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *CreateGroupPostRequest) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.

### GetSendNotification

`func (o *CreateGroupPostRequest) GetSendNotification() bool`

GetSendNotification returns the SendNotification field if non-nil, zero value otherwise.

### GetSendNotificationOk

`func (o *CreateGroupPostRequest) GetSendNotificationOk() (*bool, bool)`

GetSendNotificationOk returns a tuple with the SendNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendNotification

`func (o *CreateGroupPostRequest) SetSendNotification(v bool)`

SetSendNotification sets SendNotification field to given value.


### GetText

`func (o *CreateGroupPostRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CreateGroupPostRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CreateGroupPostRequest) SetText(v string)`

SetText sets Text field to given value.


### GetTitle

`func (o *CreateGroupPostRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateGroupPostRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateGroupPostRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetVisibility

`func (o *CreateGroupPostRequest) GetVisibility() GroupPostVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *CreateGroupPostRequest) GetVisibilityOk() (*GroupPostVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *CreateGroupPostRequest) SetVisibility(v GroupPostVisibility)`

SetVisibility sets Visibility field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


