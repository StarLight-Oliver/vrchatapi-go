# TransferGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransferTargetId** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 

## Methods

### NewTransferGroupRequest

`func NewTransferGroupRequest() *TransferGroupRequest`

NewTransferGroupRequest instantiates a new TransferGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferGroupRequestWithDefaults

`func NewTransferGroupRequestWithDefaults() *TransferGroupRequest`

NewTransferGroupRequestWithDefaults instantiates a new TransferGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransferTargetId

`func (o *TransferGroupRequest) GetTransferTargetId() string`

GetTransferTargetId returns the TransferTargetId field if non-nil, zero value otherwise.

### GetTransferTargetIdOk

`func (o *TransferGroupRequest) GetTransferTargetIdOk() (*string, bool)`

GetTransferTargetIdOk returns a tuple with the TransferTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferTargetId

`func (o *TransferGroupRequest) SetTransferTargetId(v string)`

SetTransferTargetId sets TransferTargetId field to given value.

### HasTransferTargetId

`func (o *TransferGroupRequest) HasTransferTargetId() bool`

HasTransferTargetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


