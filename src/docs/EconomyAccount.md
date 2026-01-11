# EconomyAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountActivatedOn** | **NullableTime** |  | 
**AccountId** | **NullableString** |  | 
**Blocked** | **bool** |  | 
**CanSpend** | **bool** |  | 
**Source** | **string** |  | 
**UserId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 

## Methods

### NewEconomyAccount

`func NewEconomyAccount(accountActivatedOn NullableTime, accountId NullableString, blocked bool, canSpend bool, source string, userId string, ) *EconomyAccount`

NewEconomyAccount instantiates a new EconomyAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEconomyAccountWithDefaults

`func NewEconomyAccountWithDefaults() *EconomyAccount`

NewEconomyAccountWithDefaults instantiates a new EconomyAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountActivatedOn

`func (o *EconomyAccount) GetAccountActivatedOn() time.Time`

GetAccountActivatedOn returns the AccountActivatedOn field if non-nil, zero value otherwise.

### GetAccountActivatedOnOk

`func (o *EconomyAccount) GetAccountActivatedOnOk() (*time.Time, bool)`

GetAccountActivatedOnOk returns a tuple with the AccountActivatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountActivatedOn

`func (o *EconomyAccount) SetAccountActivatedOn(v time.Time)`

SetAccountActivatedOn sets AccountActivatedOn field to given value.


### SetAccountActivatedOnNil

`func (o *EconomyAccount) SetAccountActivatedOnNil(b bool)`

 SetAccountActivatedOnNil sets the value for AccountActivatedOn to be an explicit nil

### UnsetAccountActivatedOn
`func (o *EconomyAccount) UnsetAccountActivatedOn()`

UnsetAccountActivatedOn ensures that no value is present for AccountActivatedOn, not even an explicit nil
### GetAccountId

`func (o *EconomyAccount) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *EconomyAccount) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *EconomyAccount) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### SetAccountIdNil

`func (o *EconomyAccount) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *EconomyAccount) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetBlocked

`func (o *EconomyAccount) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *EconomyAccount) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *EconomyAccount) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.


### GetCanSpend

`func (o *EconomyAccount) GetCanSpend() bool`

GetCanSpend returns the CanSpend field if non-nil, zero value otherwise.

### GetCanSpendOk

`func (o *EconomyAccount) GetCanSpendOk() (*bool, bool)`

GetCanSpendOk returns a tuple with the CanSpend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSpend

`func (o *EconomyAccount) SetCanSpend(v bool)`

SetCanSpend sets CanSpend field to given value.


### GetSource

`func (o *EconomyAccount) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EconomyAccount) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EconomyAccount) SetSource(v string)`

SetSource sets Source field to given value.


### GetUserId

`func (o *EconomyAccount) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EconomyAccount) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EconomyAccount) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


