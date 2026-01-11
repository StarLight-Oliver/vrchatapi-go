# UserNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**Id** | **string** |  | 
**Note** | **string** |  | 
**TargetUser** | Pointer to [**UserNoteTargetUser**](UserNoteTargetUser.md) |  | [optional] 
**TargetUserId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**UserId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 

## Methods

### NewUserNote

`func NewUserNote(createdAt time.Time, id string, note string, targetUserId string, userId string, ) *UserNote`

NewUserNote instantiates a new UserNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserNoteWithDefaults

`func NewUserNoteWithDefaults() *UserNote`

NewUserNoteWithDefaults instantiates a new UserNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *UserNote) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserNote) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserNote) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *UserNote) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserNote) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserNote) SetId(v string)`

SetId sets Id field to given value.


### GetNote

`func (o *UserNote) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *UserNote) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *UserNote) SetNote(v string)`

SetNote sets Note field to given value.


### GetTargetUser

`func (o *UserNote) GetTargetUser() UserNoteTargetUser`

GetTargetUser returns the TargetUser field if non-nil, zero value otherwise.

### GetTargetUserOk

`func (o *UserNote) GetTargetUserOk() (*UserNoteTargetUser, bool)`

GetTargetUserOk returns a tuple with the TargetUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUser

`func (o *UserNote) SetTargetUser(v UserNoteTargetUser)`

SetTargetUser sets TargetUser field to given value.

### HasTargetUser

`func (o *UserNote) HasTargetUser() bool`

HasTargetUser returns a boolean if a field has been set.

### GetTargetUserId

`func (o *UserNote) GetTargetUserId() string`

GetTargetUserId returns the TargetUserId field if non-nil, zero value otherwise.

### GetTargetUserIdOk

`func (o *UserNote) GetTargetUserIdOk() (*string, bool)`

GetTargetUserIdOk returns a tuple with the TargetUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUserId

`func (o *UserNote) SetTargetUserId(v string)`

SetTargetUserId sets TargetUserId field to given value.


### GetUserId

`func (o *UserNote) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserNote) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserNote) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


