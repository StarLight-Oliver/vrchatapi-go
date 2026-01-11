# Pending2FAResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QrCodeDataUrl** | **string** |  | 
**Secret** | **string** |  | 

## Methods

### NewPending2FAResult

`func NewPending2FAResult(qrCodeDataUrl string, secret string, ) *Pending2FAResult`

NewPending2FAResult instantiates a new Pending2FAResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPending2FAResultWithDefaults

`func NewPending2FAResultWithDefaults() *Pending2FAResult`

NewPending2FAResultWithDefaults instantiates a new Pending2FAResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQrCodeDataUrl

`func (o *Pending2FAResult) GetQrCodeDataUrl() string`

GetQrCodeDataUrl returns the QrCodeDataUrl field if non-nil, zero value otherwise.

### GetQrCodeDataUrlOk

`func (o *Pending2FAResult) GetQrCodeDataUrlOk() (*string, bool)`

GetQrCodeDataUrlOk returns a tuple with the QrCodeDataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCodeDataUrl

`func (o *Pending2FAResult) SetQrCodeDataUrl(v string)`

SetQrCodeDataUrl sets QrCodeDataUrl field to given value.


### GetSecret

`func (o *Pending2FAResult) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *Pending2FAResult) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *Pending2FAResult) SetSecret(v string)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


