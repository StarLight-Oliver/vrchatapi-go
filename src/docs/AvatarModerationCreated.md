# AvatarModerationCreated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarModerationType** | [**AvatarModerationType**](AvatarModerationType.md) |  | 
**Created** | **int64** | Timestamp in milliseconds since Unix epoch | 
**TargetAvatarId** | **string** |  | 

## Methods

### NewAvatarModerationCreated

`func NewAvatarModerationCreated(avatarModerationType AvatarModerationType, created int64, targetAvatarId string, ) *AvatarModerationCreated`

NewAvatarModerationCreated instantiates a new AvatarModerationCreated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvatarModerationCreatedWithDefaults

`func NewAvatarModerationCreatedWithDefaults() *AvatarModerationCreated`

NewAvatarModerationCreatedWithDefaults instantiates a new AvatarModerationCreated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarModerationType

`func (o *AvatarModerationCreated) GetAvatarModerationType() AvatarModerationType`

GetAvatarModerationType returns the AvatarModerationType field if non-nil, zero value otherwise.

### GetAvatarModerationTypeOk

`func (o *AvatarModerationCreated) GetAvatarModerationTypeOk() (*AvatarModerationType, bool)`

GetAvatarModerationTypeOk returns a tuple with the AvatarModerationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarModerationType

`func (o *AvatarModerationCreated) SetAvatarModerationType(v AvatarModerationType)`

SetAvatarModerationType sets AvatarModerationType field to given value.


### GetCreated

`func (o *AvatarModerationCreated) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AvatarModerationCreated) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AvatarModerationCreated) SetCreated(v int64)`

SetCreated sets Created field to given value.


### GetTargetAvatarId

`func (o *AvatarModerationCreated) GetTargetAvatarId() string`

GetTargetAvatarId returns the TargetAvatarId field if non-nil, zero value otherwise.

### GetTargetAvatarIdOk

`func (o *AvatarModerationCreated) GetTargetAvatarIdOk() (*string, bool)`

GetTargetAvatarIdOk returns a tuple with the TargetAvatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAvatarId

`func (o *AvatarModerationCreated) SetTargetAvatarId(v string)`

SetTargetAvatarId sets TargetAvatarId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


