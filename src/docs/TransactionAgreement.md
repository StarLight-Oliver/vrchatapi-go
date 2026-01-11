# TransactionAgreement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agreement** | **string** |  | 
**AgreementId** | **string** |  | 
**BillingType** | **string** |  | 
**Currency** | **string** |  | 
**EndDate** | **string** |  | 
**FailedAttempts** | **int32** |  | 
**Frequency** | **int32** |  | 
**ItemId** | **int32** |  | 
**LastAmount** | **float32** |  | 
**LastAmountVat** | **float32** |  | 
**LastPayment** | **string** |  | 
**NextPayment** | **string** |  | 
**Outstanding** | **int32** |  | 
**Period** | **string** |  | 
**RecurringAmt** | **float32** |  | 
**StartDate** | **string** |  | 
**Status** | **string** | This is NOT TransactionStatus, but whatever Steam return. | 
**TimeCreated** | **string** |  | 

## Methods

### NewTransactionAgreement

`func NewTransactionAgreement(agreement string, agreementId string, billingType string, currency string, endDate string, failedAttempts int32, frequency int32, itemId int32, lastAmount float32, lastAmountVat float32, lastPayment string, nextPayment string, outstanding int32, period string, recurringAmt float32, startDate string, status string, timeCreated string, ) *TransactionAgreement`

NewTransactionAgreement instantiates a new TransactionAgreement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionAgreementWithDefaults

`func NewTransactionAgreementWithDefaults() *TransactionAgreement`

NewTransactionAgreementWithDefaults instantiates a new TransactionAgreement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgreement

`func (o *TransactionAgreement) GetAgreement() string`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *TransactionAgreement) GetAgreementOk() (*string, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *TransactionAgreement) SetAgreement(v string)`

SetAgreement sets Agreement field to given value.


### GetAgreementId

`func (o *TransactionAgreement) GetAgreementId() string`

GetAgreementId returns the AgreementId field if non-nil, zero value otherwise.

### GetAgreementIdOk

`func (o *TransactionAgreement) GetAgreementIdOk() (*string, bool)`

GetAgreementIdOk returns a tuple with the AgreementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementId

`func (o *TransactionAgreement) SetAgreementId(v string)`

SetAgreementId sets AgreementId field to given value.


### GetBillingType

`func (o *TransactionAgreement) GetBillingType() string`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *TransactionAgreement) GetBillingTypeOk() (*string, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *TransactionAgreement) SetBillingType(v string)`

SetBillingType sets BillingType field to given value.


### GetCurrency

`func (o *TransactionAgreement) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransactionAgreement) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransactionAgreement) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetEndDate

`func (o *TransactionAgreement) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *TransactionAgreement) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *TransactionAgreement) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetFailedAttempts

`func (o *TransactionAgreement) GetFailedAttempts() int32`

GetFailedAttempts returns the FailedAttempts field if non-nil, zero value otherwise.

### GetFailedAttemptsOk

`func (o *TransactionAgreement) GetFailedAttemptsOk() (*int32, bool)`

GetFailedAttemptsOk returns a tuple with the FailedAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAttempts

`func (o *TransactionAgreement) SetFailedAttempts(v int32)`

SetFailedAttempts sets FailedAttempts field to given value.


### GetFrequency

`func (o *TransactionAgreement) GetFrequency() int32`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *TransactionAgreement) GetFrequencyOk() (*int32, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *TransactionAgreement) SetFrequency(v int32)`

SetFrequency sets Frequency field to given value.


### GetItemId

`func (o *TransactionAgreement) GetItemId() int32`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *TransactionAgreement) GetItemIdOk() (*int32, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *TransactionAgreement) SetItemId(v int32)`

SetItemId sets ItemId field to given value.


### GetLastAmount

`func (o *TransactionAgreement) GetLastAmount() float32`

GetLastAmount returns the LastAmount field if non-nil, zero value otherwise.

### GetLastAmountOk

`func (o *TransactionAgreement) GetLastAmountOk() (*float32, bool)`

GetLastAmountOk returns a tuple with the LastAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAmount

`func (o *TransactionAgreement) SetLastAmount(v float32)`

SetLastAmount sets LastAmount field to given value.


### GetLastAmountVat

`func (o *TransactionAgreement) GetLastAmountVat() float32`

GetLastAmountVat returns the LastAmountVat field if non-nil, zero value otherwise.

### GetLastAmountVatOk

`func (o *TransactionAgreement) GetLastAmountVatOk() (*float32, bool)`

GetLastAmountVatOk returns a tuple with the LastAmountVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAmountVat

`func (o *TransactionAgreement) SetLastAmountVat(v float32)`

SetLastAmountVat sets LastAmountVat field to given value.


### GetLastPayment

`func (o *TransactionAgreement) GetLastPayment() string`

GetLastPayment returns the LastPayment field if non-nil, zero value otherwise.

### GetLastPaymentOk

`func (o *TransactionAgreement) GetLastPaymentOk() (*string, bool)`

GetLastPaymentOk returns a tuple with the LastPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPayment

`func (o *TransactionAgreement) SetLastPayment(v string)`

SetLastPayment sets LastPayment field to given value.


### GetNextPayment

`func (o *TransactionAgreement) GetNextPayment() string`

GetNextPayment returns the NextPayment field if non-nil, zero value otherwise.

### GetNextPaymentOk

`func (o *TransactionAgreement) GetNextPaymentOk() (*string, bool)`

GetNextPaymentOk returns a tuple with the NextPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPayment

`func (o *TransactionAgreement) SetNextPayment(v string)`

SetNextPayment sets NextPayment field to given value.


### GetOutstanding

`func (o *TransactionAgreement) GetOutstanding() int32`

GetOutstanding returns the Outstanding field if non-nil, zero value otherwise.

### GetOutstandingOk

`func (o *TransactionAgreement) GetOutstandingOk() (*int32, bool)`

GetOutstandingOk returns a tuple with the Outstanding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutstanding

`func (o *TransactionAgreement) SetOutstanding(v int32)`

SetOutstanding sets Outstanding field to given value.


### GetPeriod

`func (o *TransactionAgreement) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *TransactionAgreement) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *TransactionAgreement) SetPeriod(v string)`

SetPeriod sets Period field to given value.


### GetRecurringAmt

`func (o *TransactionAgreement) GetRecurringAmt() float32`

GetRecurringAmt returns the RecurringAmt field if non-nil, zero value otherwise.

### GetRecurringAmtOk

`func (o *TransactionAgreement) GetRecurringAmtOk() (*float32, bool)`

GetRecurringAmtOk returns a tuple with the RecurringAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringAmt

`func (o *TransactionAgreement) SetRecurringAmt(v float32)`

SetRecurringAmt sets RecurringAmt field to given value.


### GetStartDate

`func (o *TransactionAgreement) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TransactionAgreement) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TransactionAgreement) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetStatus

`func (o *TransactionAgreement) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransactionAgreement) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransactionAgreement) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimeCreated

`func (o *TransactionAgreement) GetTimeCreated() string`

GetTimeCreated returns the TimeCreated field if non-nil, zero value otherwise.

### GetTimeCreatedOk

`func (o *TransactionAgreement) GetTimeCreatedOk() (*string, bool)`

GetTimeCreatedOk returns a tuple with the TimeCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeCreated

`func (o *TransactionAgreement) SetTimeCreated(v string)`

SetTimeCreated sets TimeCreated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


