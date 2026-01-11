# RespondNotificationV2Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseData** | Pointer to **string** |  | [optional] [default to ""]
**ResponseType** | **string** |  | 

## Methods

### NewRespondNotificationV2Request

`func NewRespondNotificationV2Request(responseType string, ) *RespondNotificationV2Request`

NewRespondNotificationV2Request instantiates a new RespondNotificationV2Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRespondNotificationV2RequestWithDefaults

`func NewRespondNotificationV2RequestWithDefaults() *RespondNotificationV2Request`

NewRespondNotificationV2RequestWithDefaults instantiates a new RespondNotificationV2Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseData

`func (o *RespondNotificationV2Request) GetResponseData() string`

GetResponseData returns the ResponseData field if non-nil, zero value otherwise.

### GetResponseDataOk

`func (o *RespondNotificationV2Request) GetResponseDataOk() (*string, bool)`

GetResponseDataOk returns a tuple with the ResponseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseData

`func (o *RespondNotificationV2Request) SetResponseData(v string)`

SetResponseData sets ResponseData field to given value.

### HasResponseData

`func (o *RespondNotificationV2Request) HasResponseData() bool`

HasResponseData returns a boolean if a field has been set.

### GetResponseType

`func (o *RespondNotificationV2Request) GetResponseType() string`

GetResponseType returns the ResponseType field if non-nil, zero value otherwise.

### GetResponseTypeOk

`func (o *RespondNotificationV2Request) GetResponseTypeOk() (*string, bool)`

GetResponseTypeOk returns a tuple with the ResponseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseType

`func (o *RespondNotificationV2Request) SetResponseType(v string)`

SetResponseType sets ResponseType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


