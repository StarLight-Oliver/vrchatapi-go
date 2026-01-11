# TwoFactorRecoveryCodes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Otp** | Pointer to [**[]TwoFactorRecoveryCodesOtpInner**](TwoFactorRecoveryCodesOtpInner.md) |  | [optional] 
**RequiresTwoFactorAuth** | Pointer to **[]string** |  | [optional] 

## Methods

### NewTwoFactorRecoveryCodes

`func NewTwoFactorRecoveryCodes() *TwoFactorRecoveryCodes`

NewTwoFactorRecoveryCodes instantiates a new TwoFactorRecoveryCodes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTwoFactorRecoveryCodesWithDefaults

`func NewTwoFactorRecoveryCodesWithDefaults() *TwoFactorRecoveryCodes`

NewTwoFactorRecoveryCodesWithDefaults instantiates a new TwoFactorRecoveryCodes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOtp

`func (o *TwoFactorRecoveryCodes) GetOtp() []TwoFactorRecoveryCodesOtpInner`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *TwoFactorRecoveryCodes) GetOtpOk() (*[]TwoFactorRecoveryCodesOtpInner, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *TwoFactorRecoveryCodes) SetOtp(v []TwoFactorRecoveryCodesOtpInner)`

SetOtp sets Otp field to given value.

### HasOtp

`func (o *TwoFactorRecoveryCodes) HasOtp() bool`

HasOtp returns a boolean if a field has been set.

### GetRequiresTwoFactorAuth

`func (o *TwoFactorRecoveryCodes) GetRequiresTwoFactorAuth() []string`

GetRequiresTwoFactorAuth returns the RequiresTwoFactorAuth field if non-nil, zero value otherwise.

### GetRequiresTwoFactorAuthOk

`func (o *TwoFactorRecoveryCodes) GetRequiresTwoFactorAuthOk() (*[]string, bool)`

GetRequiresTwoFactorAuthOk returns a tuple with the RequiresTwoFactorAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresTwoFactorAuth

`func (o *TwoFactorRecoveryCodes) SetRequiresTwoFactorAuth(v []string)`

SetRequiresTwoFactorAuth sets RequiresTwoFactorAuth field to given value.

### HasRequiresTwoFactorAuth

`func (o *TwoFactorRecoveryCodes) HasRequiresTwoFactorAuth() bool`

HasRequiresTwoFactorAuth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


