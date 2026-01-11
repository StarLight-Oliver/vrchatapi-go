# AgreementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgreementCode** | [**AgreementCode**](AgreementCode.md) |  | [default to CONTENT_COPYRIGHT_OWNED]
**AgreementFulltext** | **string** | The full text of the agreement (currently &#x60;By clicking OK, I certify that I have the necessary rights to upload this content and that it will not infringe on any third-party legal or intellectual property rights.&#x60;). | 
**ContentId** | **string** | The id of the content being uploaded, such as a WorldID, AvatarID, or PropID. | 
**Version** | **int32** | The version of the agreement (currently &#x60;1&#x60;). | 

## Methods

### NewAgreementRequest

`func NewAgreementRequest(agreementCode AgreementCode, agreementFulltext string, contentId string, version int32, ) *AgreementRequest`

NewAgreementRequest instantiates a new AgreementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgreementRequestWithDefaults

`func NewAgreementRequestWithDefaults() *AgreementRequest`

NewAgreementRequestWithDefaults instantiates a new AgreementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgreementCode

`func (o *AgreementRequest) GetAgreementCode() AgreementCode`

GetAgreementCode returns the AgreementCode field if non-nil, zero value otherwise.

### GetAgreementCodeOk

`func (o *AgreementRequest) GetAgreementCodeOk() (*AgreementCode, bool)`

GetAgreementCodeOk returns a tuple with the AgreementCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementCode

`func (o *AgreementRequest) SetAgreementCode(v AgreementCode)`

SetAgreementCode sets AgreementCode field to given value.


### GetAgreementFulltext

`func (o *AgreementRequest) GetAgreementFulltext() string`

GetAgreementFulltext returns the AgreementFulltext field if non-nil, zero value otherwise.

### GetAgreementFulltextOk

`func (o *AgreementRequest) GetAgreementFulltextOk() (*string, bool)`

GetAgreementFulltextOk returns a tuple with the AgreementFulltext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementFulltext

`func (o *AgreementRequest) SetAgreementFulltext(v string)`

SetAgreementFulltext sets AgreementFulltext field to given value.


### GetContentId

`func (o *AgreementRequest) GetContentId() string`

GetContentId returns the ContentId field if non-nil, zero value otherwise.

### GetContentIdOk

`func (o *AgreementRequest) GetContentIdOk() (*string, bool)`

GetContentIdOk returns a tuple with the ContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentId

`func (o *AgreementRequest) SetContentId(v string)`

SetContentId sets ContentId field to given value.


### GetVersion

`func (o *AgreementRequest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AgreementRequest) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AgreementRequest) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


