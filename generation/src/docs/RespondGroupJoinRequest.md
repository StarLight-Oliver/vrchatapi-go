# RespondGroupJoinRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**GroupJoinRequestAction**](GroupJoinRequestAction.md) |  | 
**Block** | Pointer to **bool** | Whether to block the user from requesting again | [optional] 

## Methods

### NewRespondGroupJoinRequest

`func NewRespondGroupJoinRequest(action GroupJoinRequestAction, ) *RespondGroupJoinRequest`

NewRespondGroupJoinRequest instantiates a new RespondGroupJoinRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRespondGroupJoinRequestWithDefaults

`func NewRespondGroupJoinRequestWithDefaults() *RespondGroupJoinRequest`

NewRespondGroupJoinRequestWithDefaults instantiates a new RespondGroupJoinRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *RespondGroupJoinRequest) GetAction() GroupJoinRequestAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RespondGroupJoinRequest) GetActionOk() (*GroupJoinRequestAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RespondGroupJoinRequest) SetAction(v GroupJoinRequestAction)`

SetAction sets Action field to given value.


### GetBlock

`func (o *RespondGroupJoinRequest) GetBlock() bool`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *RespondGroupJoinRequest) GetBlockOk() (*bool, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *RespondGroupJoinRequest) SetBlock(v bool)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *RespondGroupJoinRequest) HasBlock() bool`

HasBlock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


