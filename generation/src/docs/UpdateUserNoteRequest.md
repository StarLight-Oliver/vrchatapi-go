# UpdateUserNoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Note** | **string** |  | 
**TargetUserId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 

## Methods

### NewUpdateUserNoteRequest

`func NewUpdateUserNoteRequest(note string, targetUserId string, ) *UpdateUserNoteRequest`

NewUpdateUserNoteRequest instantiates a new UpdateUserNoteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserNoteRequestWithDefaults

`func NewUpdateUserNoteRequestWithDefaults() *UpdateUserNoteRequest`

NewUpdateUserNoteRequestWithDefaults instantiates a new UpdateUserNoteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNote

`func (o *UpdateUserNoteRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *UpdateUserNoteRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *UpdateUserNoteRequest) SetNote(v string)`

SetNote sets Note field to given value.


### GetTargetUserId

`func (o *UpdateUserNoteRequest) GetTargetUserId() string`

GetTargetUserId returns the TargetUserId field if non-nil, zero value otherwise.

### GetTargetUserIdOk

`func (o *UpdateUserNoteRequest) GetTargetUserIdOk() (*string, bool)`

GetTargetUserIdOk returns a tuple with the TargetUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUserId

`func (o *UpdateUserNoteRequest) SetTargetUserId(v string)`

SetTargetUserId sets TargetUserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


