# CreateAvatarModerationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarModerationType** | [**AvatarModerationType**](AvatarModerationType.md) |  | 
**TargetAvatarId** | **string** |  | 

## Methods

### NewCreateAvatarModerationRequest

`func NewCreateAvatarModerationRequest(avatarModerationType AvatarModerationType, targetAvatarId string, ) *CreateAvatarModerationRequest`

NewCreateAvatarModerationRequest instantiates a new CreateAvatarModerationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAvatarModerationRequestWithDefaults

`func NewCreateAvatarModerationRequestWithDefaults() *CreateAvatarModerationRequest`

NewCreateAvatarModerationRequestWithDefaults instantiates a new CreateAvatarModerationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarModerationType

`func (o *CreateAvatarModerationRequest) GetAvatarModerationType() AvatarModerationType`

GetAvatarModerationType returns the AvatarModerationType field if non-nil, zero value otherwise.

### GetAvatarModerationTypeOk

`func (o *CreateAvatarModerationRequest) GetAvatarModerationTypeOk() (*AvatarModerationType, bool)`

GetAvatarModerationTypeOk returns a tuple with the AvatarModerationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarModerationType

`func (o *CreateAvatarModerationRequest) SetAvatarModerationType(v AvatarModerationType)`

SetAvatarModerationType sets AvatarModerationType field to given value.


### GetTargetAvatarId

`func (o *CreateAvatarModerationRequest) GetTargetAvatarId() string`

GetTargetAvatarId returns the TargetAvatarId field if non-nil, zero value otherwise.

### GetTargetAvatarIdOk

`func (o *CreateAvatarModerationRequest) GetTargetAvatarIdOk() (*string, bool)`

GetTargetAvatarIdOk returns a tuple with the TargetAvatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAvatarId

`func (o *CreateAvatarModerationRequest) SetTargetAvatarId(v string)`

SetTargetAvatarId sets TargetAvatarId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


