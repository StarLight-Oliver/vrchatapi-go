# TokenBundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | price of the bundle | 
**AppleProductId** | **string** |  | 
**Description** | **string** |  | 
**GoogleProductId** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**ImageUrl** | **string** | direct url to image | 
**OculusSku** | **string** |  | 
**SteamItemId** | **string** |  | 
**Tokens** | **int32** | number of tokens received | 

## Methods

### NewTokenBundle

`func NewTokenBundle(amount int32, appleProductId string, description string, id string, imageUrl string, oculusSku string, steamItemId string, tokens int32, ) *TokenBundle`

NewTokenBundle instantiates a new TokenBundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenBundleWithDefaults

`func NewTokenBundleWithDefaults() *TokenBundle`

NewTokenBundleWithDefaults instantiates a new TokenBundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *TokenBundle) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TokenBundle) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TokenBundle) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetAppleProductId

`func (o *TokenBundle) GetAppleProductId() string`

GetAppleProductId returns the AppleProductId field if non-nil, zero value otherwise.

### GetAppleProductIdOk

`func (o *TokenBundle) GetAppleProductIdOk() (*string, bool)`

GetAppleProductIdOk returns a tuple with the AppleProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleProductId

`func (o *TokenBundle) SetAppleProductId(v string)`

SetAppleProductId sets AppleProductId field to given value.


### GetDescription

`func (o *TokenBundle) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TokenBundle) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TokenBundle) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetGoogleProductId

`func (o *TokenBundle) GetGoogleProductId() string`

GetGoogleProductId returns the GoogleProductId field if non-nil, zero value otherwise.

### GetGoogleProductIdOk

`func (o *TokenBundle) GetGoogleProductIdOk() (*string, bool)`

GetGoogleProductIdOk returns a tuple with the GoogleProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleProductId

`func (o *TokenBundle) SetGoogleProductId(v string)`

SetGoogleProductId sets GoogleProductId field to given value.

### HasGoogleProductId

`func (o *TokenBundle) HasGoogleProductId() bool`

HasGoogleProductId returns a boolean if a field has been set.

### GetId

`func (o *TokenBundle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenBundle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenBundle) SetId(v string)`

SetId sets Id field to given value.


### GetImageUrl

`func (o *TokenBundle) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *TokenBundle) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *TokenBundle) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.


### GetOculusSku

`func (o *TokenBundle) GetOculusSku() string`

GetOculusSku returns the OculusSku field if non-nil, zero value otherwise.

### GetOculusSkuOk

`func (o *TokenBundle) GetOculusSkuOk() (*string, bool)`

GetOculusSkuOk returns a tuple with the OculusSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOculusSku

`func (o *TokenBundle) SetOculusSku(v string)`

SetOculusSku sets OculusSku field to given value.


### GetSteamItemId

`func (o *TokenBundle) GetSteamItemId() string`

GetSteamItemId returns the SteamItemId field if non-nil, zero value otherwise.

### GetSteamItemIdOk

`func (o *TokenBundle) GetSteamItemIdOk() (*string, bool)`

GetSteamItemIdOk returns a tuple with the SteamItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamItemId

`func (o *TokenBundle) SetSteamItemId(v string)`

SetSteamItemId sets SteamItemId field to given value.


### GetTokens

`func (o *TokenBundle) GetTokens() int32`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *TokenBundle) GetTokensOk() (*int32, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *TokenBundle) SetTokens(v int32)`

SetTokens sets Tokens field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


