# ChangeUserTagsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | **[]string** | The tags being added or removed. | 

## Methods

### NewChangeUserTagsRequest

`func NewChangeUserTagsRequest(tags []string, ) *ChangeUserTagsRequest`

NewChangeUserTagsRequest instantiates a new ChangeUserTagsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeUserTagsRequestWithDefaults

`func NewChangeUserTagsRequestWithDefaults() *ChangeUserTagsRequest`

NewChangeUserTagsRequestWithDefaults instantiates a new ChangeUserTagsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *ChangeUserTagsRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ChangeUserTagsRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ChangeUserTagsRequest) SetTags(v []string)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


