# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agreement** | Pointer to [**TransactionAgreement**](TransactionAgreement.md) |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**Error** | **NullableString** |  | 
**Id** | **string** |  | 
**IsGift** | Pointer to **bool** |  | [optional] [default to false]
**IsTokens** | Pointer to **bool** |  | [optional] [default to false]
**Sandbox** | **bool** |  | [default to false]
**Status** | [**TransactionStatus**](TransactionStatus.md) |  | [default to ACTIVE]
**Steam** | Pointer to [**TransactionSteamInfo**](TransactionSteamInfo.md) |  | [optional] 
**Subscription** | [**Subscription**](Subscription.md) |  | 
**UpdatedAt** | **time.Time** |  | 
**UserDisplayName** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 

## Methods

### NewTransaction

`func NewTransaction(createdAt time.Time, error_ NullableString, id string, sandbox bool, status TransactionStatus, subscription Subscription, updatedAt time.Time, ) *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgreement

`func (o *Transaction) GetAgreement() TransactionAgreement`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *Transaction) GetAgreementOk() (*TransactionAgreement, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *Transaction) SetAgreement(v TransactionAgreement)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *Transaction) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Transaction) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Transaction) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Transaction) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetError

`func (o *Transaction) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Transaction) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Transaction) SetError(v string)`

SetError sets Error field to given value.


### SetErrorNil

`func (o *Transaction) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *Transaction) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetId

`func (o *Transaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Transaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Transaction) SetId(v string)`

SetId sets Id field to given value.


### GetIsGift

`func (o *Transaction) GetIsGift() bool`

GetIsGift returns the IsGift field if non-nil, zero value otherwise.

### GetIsGiftOk

`func (o *Transaction) GetIsGiftOk() (*bool, bool)`

GetIsGiftOk returns a tuple with the IsGift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGift

`func (o *Transaction) SetIsGift(v bool)`

SetIsGift sets IsGift field to given value.

### HasIsGift

`func (o *Transaction) HasIsGift() bool`

HasIsGift returns a boolean if a field has been set.

### GetIsTokens

`func (o *Transaction) GetIsTokens() bool`

GetIsTokens returns the IsTokens field if non-nil, zero value otherwise.

### GetIsTokensOk

`func (o *Transaction) GetIsTokensOk() (*bool, bool)`

GetIsTokensOk returns a tuple with the IsTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTokens

`func (o *Transaction) SetIsTokens(v bool)`

SetIsTokens sets IsTokens field to given value.

### HasIsTokens

`func (o *Transaction) HasIsTokens() bool`

HasIsTokens returns a boolean if a field has been set.

### GetSandbox

`func (o *Transaction) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *Transaction) GetSandboxOk() (*bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *Transaction) SetSandbox(v bool)`

SetSandbox sets Sandbox field to given value.


### GetStatus

`func (o *Transaction) GetStatus() TransactionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Transaction) GetStatusOk() (*TransactionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Transaction) SetStatus(v TransactionStatus)`

SetStatus sets Status field to given value.


### GetSteam

`func (o *Transaction) GetSteam() TransactionSteamInfo`

GetSteam returns the Steam field if non-nil, zero value otherwise.

### GetSteamOk

`func (o *Transaction) GetSteamOk() (*TransactionSteamInfo, bool)`

GetSteamOk returns a tuple with the Steam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteam

`func (o *Transaction) SetSteam(v TransactionSteamInfo)`

SetSteam sets Steam field to given value.

### HasSteam

`func (o *Transaction) HasSteam() bool`

HasSteam returns a boolean if a field has been set.

### GetSubscription

`func (o *Transaction) GetSubscription() Subscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *Transaction) GetSubscriptionOk() (*Subscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *Transaction) SetSubscription(v Subscription)`

SetSubscription sets Subscription field to given value.


### GetUpdatedAt

`func (o *Transaction) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Transaction) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Transaction) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserDisplayName

`func (o *Transaction) GetUserDisplayName() string`

GetUserDisplayName returns the UserDisplayName field if non-nil, zero value otherwise.

### GetUserDisplayNameOk

`func (o *Transaction) GetUserDisplayNameOk() (*string, bool)`

GetUserDisplayNameOk returns a tuple with the UserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDisplayName

`func (o *Transaction) SetUserDisplayName(v string)`

SetUserDisplayName sets UserDisplayName field to given value.

### HasUserDisplayName

`func (o *Transaction) HasUserDisplayName() bool`

HasUserDisplayName returns a boolean if a field has been set.

### GetUserId

`func (o *Transaction) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Transaction) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Transaction) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Transaction) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


