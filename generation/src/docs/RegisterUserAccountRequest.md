# RegisterUserAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedTOSVersion** | **int32** | The most recent version of the TOS | 
**CaptchaCode** | **string** | Captcha code | 
**Day** | **string** | Birth day of month | 
**Email** | **string** | Email address | 
**Month** | **string** | Birth month of year | 
**Password** | **string** | Password | 
**Subscribe** | **bool** | Whether to recieve promotional emails | 
**Username** | **string** | Display Name / Username (Username is a sanitized version) | 
**Year** | **string** | Birth year | 

## Methods

### NewRegisterUserAccountRequest

`func NewRegisterUserAccountRequest(acceptedTOSVersion int32, captchaCode string, day string, email string, month string, password string, subscribe bool, username string, year string, ) *RegisterUserAccountRequest`

NewRegisterUserAccountRequest instantiates a new RegisterUserAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterUserAccountRequestWithDefaults

`func NewRegisterUserAccountRequestWithDefaults() *RegisterUserAccountRequest`

NewRegisterUserAccountRequestWithDefaults instantiates a new RegisterUserAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedTOSVersion

`func (o *RegisterUserAccountRequest) GetAcceptedTOSVersion() int32`

GetAcceptedTOSVersion returns the AcceptedTOSVersion field if non-nil, zero value otherwise.

### GetAcceptedTOSVersionOk

`func (o *RegisterUserAccountRequest) GetAcceptedTOSVersionOk() (*int32, bool)`

GetAcceptedTOSVersionOk returns a tuple with the AcceptedTOSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedTOSVersion

`func (o *RegisterUserAccountRequest) SetAcceptedTOSVersion(v int32)`

SetAcceptedTOSVersion sets AcceptedTOSVersion field to given value.


### GetCaptchaCode

`func (o *RegisterUserAccountRequest) GetCaptchaCode() string`

GetCaptchaCode returns the CaptchaCode field if non-nil, zero value otherwise.

### GetCaptchaCodeOk

`func (o *RegisterUserAccountRequest) GetCaptchaCodeOk() (*string, bool)`

GetCaptchaCodeOk returns a tuple with the CaptchaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptchaCode

`func (o *RegisterUserAccountRequest) SetCaptchaCode(v string)`

SetCaptchaCode sets CaptchaCode field to given value.


### GetDay

`func (o *RegisterUserAccountRequest) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *RegisterUserAccountRequest) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *RegisterUserAccountRequest) SetDay(v string)`

SetDay sets Day field to given value.


### GetEmail

`func (o *RegisterUserAccountRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RegisterUserAccountRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RegisterUserAccountRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetMonth

`func (o *RegisterUserAccountRequest) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *RegisterUserAccountRequest) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *RegisterUserAccountRequest) SetMonth(v string)`

SetMonth sets Month field to given value.


### GetPassword

`func (o *RegisterUserAccountRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RegisterUserAccountRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RegisterUserAccountRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetSubscribe

`func (o *RegisterUserAccountRequest) GetSubscribe() bool`

GetSubscribe returns the Subscribe field if non-nil, zero value otherwise.

### GetSubscribeOk

`func (o *RegisterUserAccountRequest) GetSubscribeOk() (*bool, bool)`

GetSubscribeOk returns a tuple with the Subscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribe

`func (o *RegisterUserAccountRequest) SetSubscribe(v bool)`

SetSubscribe sets Subscribe field to given value.


### GetUsername

`func (o *RegisterUserAccountRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RegisterUserAccountRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RegisterUserAccountRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetYear

`func (o *RegisterUserAccountRequest) GetYear() string`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *RegisterUserAccountRequest) GetYearOk() (*string, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *RegisterUserAccountRequest) SetYear(v string)`

SetYear sets Year field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


