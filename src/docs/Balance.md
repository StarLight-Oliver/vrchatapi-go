# Balance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | **int32** |  | [default to 0]
**NoTransactions** | Pointer to **bool** |  | [optional] 
**TiliaResponse** | Pointer to **bool** |  | [optional] 

## Methods

### NewBalance

`func NewBalance(balance int32, ) *Balance`

NewBalance instantiates a new Balance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceWithDefaults

`func NewBalanceWithDefaults() *Balance`

NewBalanceWithDefaults instantiates a new Balance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *Balance) GetBalance() int32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *Balance) GetBalanceOk() (*int32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *Balance) SetBalance(v int32)`

SetBalance sets Balance field to given value.


### GetNoTransactions

`func (o *Balance) GetNoTransactions() bool`

GetNoTransactions returns the NoTransactions field if non-nil, zero value otherwise.

### GetNoTransactionsOk

`func (o *Balance) GetNoTransactionsOk() (*bool, bool)`

GetNoTransactionsOk returns a tuple with the NoTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoTransactions

`func (o *Balance) SetNoTransactions(v bool)`

SetNoTransactions sets NoTransactions field to given value.

### HasNoTransactions

`func (o *Balance) HasNoTransactions() bool`

HasNoTransactions returns a boolean if a field has been set.

### GetTiliaResponse

`func (o *Balance) GetTiliaResponse() bool`

GetTiliaResponse returns the TiliaResponse field if non-nil, zero value otherwise.

### GetTiliaResponseOk

`func (o *Balance) GetTiliaResponseOk() (*bool, bool)`

GetTiliaResponseOk returns a tuple with the TiliaResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiliaResponse

`func (o *Balance) SetTiliaResponse(v bool)`

SetTiliaResponse sets TiliaResponse field to given value.

### HasTiliaResponse

`func (o *Balance) HasTiliaResponse() bool`

HasTiliaResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


