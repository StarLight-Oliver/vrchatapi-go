# AgreementStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agreed** | **bool** | Whether the user has agreed for this content. | 
**AgreementCode** | [**AgreementCode**](AgreementCode.md) |  | [default to CONTENT_COPYRIGHT_OWNED]
**ContentId** | **string** | The id of the content being uploaded, such as a WorldID, AvatarID, or PropID. | 
**UserId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**Version** | **int32** | The version of the agreement. | 

## Methods

### NewAgreementStatus

`func NewAgreementStatus(agreed bool, agreementCode AgreementCode, contentId string, userId string, version int32, ) *AgreementStatus`

NewAgreementStatus instantiates a new AgreementStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgreementStatusWithDefaults

`func NewAgreementStatusWithDefaults() *AgreementStatus`

NewAgreementStatusWithDefaults instantiates a new AgreementStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgreed

`func (o *AgreementStatus) GetAgreed() bool`

GetAgreed returns the Agreed field if non-nil, zero value otherwise.

### GetAgreedOk

`func (o *AgreementStatus) GetAgreedOk() (*bool, bool)`

GetAgreedOk returns a tuple with the Agreed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreed

`func (o *AgreementStatus) SetAgreed(v bool)`

SetAgreed sets Agreed field to given value.


### GetAgreementCode

`func (o *AgreementStatus) GetAgreementCode() AgreementCode`

GetAgreementCode returns the AgreementCode field if non-nil, zero value otherwise.

### GetAgreementCodeOk

`func (o *AgreementStatus) GetAgreementCodeOk() (*AgreementCode, bool)`

GetAgreementCodeOk returns a tuple with the AgreementCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementCode

`func (o *AgreementStatus) SetAgreementCode(v AgreementCode)`

SetAgreementCode sets AgreementCode field to given value.


### GetContentId

`func (o *AgreementStatus) GetContentId() string`

GetContentId returns the ContentId field if non-nil, zero value otherwise.

### GetContentIdOk

`func (o *AgreementStatus) GetContentIdOk() (*string, bool)`

GetContentIdOk returns a tuple with the ContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentId

`func (o *AgreementStatus) SetContentId(v string)`

SetContentId sets ContentId field to given value.


### GetUserId

`func (o *AgreementStatus) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AgreementStatus) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AgreementStatus) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetVersion

`func (o *AgreementStatus) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AgreementStatus) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AgreementStatus) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


