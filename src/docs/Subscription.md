# Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** |  | 
**AppleProductId** | Pointer to **string** |  | [optional] 
**Description** | **string** |  | 
**GooglePlanId** | Pointer to **string** |  | [optional] 
**GoogleProductId** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**OculusSku** | Pointer to **string** |  | [optional] 
**Period** | [**SubscriptionPeriod**](SubscriptionPeriod.md) |  | [default to MONTH]
**PicoSku** | Pointer to **string** |  | [optional] 
**SteamItemId** | **string** |  | 
**Tier** | **int32** |  | 

## Methods

### NewSubscription

`func NewSubscription(amount float32, description string, id string, period SubscriptionPeriod, steamItemId string, tier int32, ) *Subscription`

NewSubscription instantiates a new Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionWithDefaults

`func NewSubscriptionWithDefaults() *Subscription`

NewSubscriptionWithDefaults instantiates a new Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Subscription) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Subscription) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Subscription) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetAppleProductId

`func (o *Subscription) GetAppleProductId() string`

GetAppleProductId returns the AppleProductId field if non-nil, zero value otherwise.

### GetAppleProductIdOk

`func (o *Subscription) GetAppleProductIdOk() (*string, bool)`

GetAppleProductIdOk returns a tuple with the AppleProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppleProductId

`func (o *Subscription) SetAppleProductId(v string)`

SetAppleProductId sets AppleProductId field to given value.

### HasAppleProductId

`func (o *Subscription) HasAppleProductId() bool`

HasAppleProductId returns a boolean if a field has been set.

### GetDescription

`func (o *Subscription) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Subscription) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Subscription) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetGooglePlanId

`func (o *Subscription) GetGooglePlanId() string`

GetGooglePlanId returns the GooglePlanId field if non-nil, zero value otherwise.

### GetGooglePlanIdOk

`func (o *Subscription) GetGooglePlanIdOk() (*string, bool)`

GetGooglePlanIdOk returns a tuple with the GooglePlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGooglePlanId

`func (o *Subscription) SetGooglePlanId(v string)`

SetGooglePlanId sets GooglePlanId field to given value.

### HasGooglePlanId

`func (o *Subscription) HasGooglePlanId() bool`

HasGooglePlanId returns a boolean if a field has been set.

### GetGoogleProductId

`func (o *Subscription) GetGoogleProductId() string`

GetGoogleProductId returns the GoogleProductId field if non-nil, zero value otherwise.

### GetGoogleProductIdOk

`func (o *Subscription) GetGoogleProductIdOk() (*string, bool)`

GetGoogleProductIdOk returns a tuple with the GoogleProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleProductId

`func (o *Subscription) SetGoogleProductId(v string)`

SetGoogleProductId sets GoogleProductId field to given value.

### HasGoogleProductId

`func (o *Subscription) HasGoogleProductId() bool`

HasGoogleProductId returns a boolean if a field has been set.

### GetId

`func (o *Subscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subscription) SetId(v string)`

SetId sets Id field to given value.


### GetOculusSku

`func (o *Subscription) GetOculusSku() string`

GetOculusSku returns the OculusSku field if non-nil, zero value otherwise.

### GetOculusSkuOk

`func (o *Subscription) GetOculusSkuOk() (*string, bool)`

GetOculusSkuOk returns a tuple with the OculusSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOculusSku

`func (o *Subscription) SetOculusSku(v string)`

SetOculusSku sets OculusSku field to given value.

### HasOculusSku

`func (o *Subscription) HasOculusSku() bool`

HasOculusSku returns a boolean if a field has been set.

### GetPeriod

`func (o *Subscription) GetPeriod() SubscriptionPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *Subscription) GetPeriodOk() (*SubscriptionPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *Subscription) SetPeriod(v SubscriptionPeriod)`

SetPeriod sets Period field to given value.


### GetPicoSku

`func (o *Subscription) GetPicoSku() string`

GetPicoSku returns the PicoSku field if non-nil, zero value otherwise.

### GetPicoSkuOk

`func (o *Subscription) GetPicoSkuOk() (*string, bool)`

GetPicoSkuOk returns a tuple with the PicoSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicoSku

`func (o *Subscription) SetPicoSku(v string)`

SetPicoSku sets PicoSku field to given value.

### HasPicoSku

`func (o *Subscription) HasPicoSku() bool`

HasPicoSku returns a boolean if a field has been set.

### GetSteamItemId

`func (o *Subscription) GetSteamItemId() string`

GetSteamItemId returns the SteamItemId field if non-nil, zero value otherwise.

### GetSteamItemIdOk

`func (o *Subscription) GetSteamItemIdOk() (*string, bool)`

GetSteamItemIdOk returns a tuple with the SteamItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamItemId

`func (o *Subscription) SetSteamItemId(v string)`

SetSteamItemId sets SteamItemId field to given value.


### GetTier

`func (o *Subscription) GetTier() int32`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *Subscription) GetTierOk() (*int32, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *Subscription) SetTier(v int32)`

SetTier sets Tier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


