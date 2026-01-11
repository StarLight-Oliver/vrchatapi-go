# Verify2FAResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**Verified** | **bool** |  | 

## Methods

### NewVerify2FAResult

`func NewVerify2FAResult(verified bool, ) *Verify2FAResult`

NewVerify2FAResult instantiates a new Verify2FAResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerify2FAResultWithDefaults

`func NewVerify2FAResultWithDefaults() *Verify2FAResult`

NewVerify2FAResultWithDefaults instantiates a new Verify2FAResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *Verify2FAResult) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Verify2FAResult) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Verify2FAResult) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Verify2FAResult) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVerified

`func (o *Verify2FAResult) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *Verify2FAResult) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *Verify2FAResult) SetVerified(v bool)`

SetVerified sets Verified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


