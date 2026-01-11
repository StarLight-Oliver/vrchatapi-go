# Agreement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgreementCode** | [**AgreementCode**](AgreementCode.md) |  | [default to CONTENT_COPYRIGHT_OWNED]
**AgreementFulltext** | Pointer to **string** | The full text of the agreement. | [optional] 
**ContentId** | **string** | The id of the content being uploaded, such as a WorldID, AvatarID, or PropID. | 
**Created** | **string** | When the agreement was created | 
**Id** | **string** | The id of the agreement. | 
**Tags** | **[]string** |  | 
**UserId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**Version** | **int32** | The version of the agreement. | 

## Methods

### NewAgreement

`func NewAgreement(agreementCode AgreementCode, contentId string, created string, id string, tags []string, userId string, version int32, ) *Agreement`

NewAgreement instantiates a new Agreement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgreementWithDefaults

`func NewAgreementWithDefaults() *Agreement`

NewAgreementWithDefaults instantiates a new Agreement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgreementCode

`func (o *Agreement) GetAgreementCode() AgreementCode`

GetAgreementCode returns the AgreementCode field if non-nil, zero value otherwise.

### GetAgreementCodeOk

`func (o *Agreement) GetAgreementCodeOk() (*AgreementCode, bool)`

GetAgreementCodeOk returns a tuple with the AgreementCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementCode

`func (o *Agreement) SetAgreementCode(v AgreementCode)`

SetAgreementCode sets AgreementCode field to given value.


### GetAgreementFulltext

`func (o *Agreement) GetAgreementFulltext() string`

GetAgreementFulltext returns the AgreementFulltext field if non-nil, zero value otherwise.

### GetAgreementFulltextOk

`func (o *Agreement) GetAgreementFulltextOk() (*string, bool)`

GetAgreementFulltextOk returns a tuple with the AgreementFulltext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementFulltext

`func (o *Agreement) SetAgreementFulltext(v string)`

SetAgreementFulltext sets AgreementFulltext field to given value.

### HasAgreementFulltext

`func (o *Agreement) HasAgreementFulltext() bool`

HasAgreementFulltext returns a boolean if a field has been set.

### GetContentId

`func (o *Agreement) GetContentId() string`

GetContentId returns the ContentId field if non-nil, zero value otherwise.

### GetContentIdOk

`func (o *Agreement) GetContentIdOk() (*string, bool)`

GetContentIdOk returns a tuple with the ContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentId

`func (o *Agreement) SetContentId(v string)`

SetContentId sets ContentId field to given value.


### GetCreated

`func (o *Agreement) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Agreement) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Agreement) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetId

`func (o *Agreement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Agreement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Agreement) SetId(v string)`

SetId sets Id field to given value.


### GetTags

`func (o *Agreement) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Agreement) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Agreement) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetUserId

`func (o *Agreement) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Agreement) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Agreement) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetVersion

`func (o *Agreement) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Agreement) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Agreement) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


