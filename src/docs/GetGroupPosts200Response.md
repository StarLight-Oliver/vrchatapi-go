# GetGroupPosts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Posts** | Pointer to [**[]GroupPost**](GroupPost.md) |  | [optional] 

## Methods

### NewGetGroupPosts200Response

`func NewGetGroupPosts200Response() *GetGroupPosts200Response`

NewGetGroupPosts200Response instantiates a new GetGroupPosts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGroupPosts200ResponseWithDefaults

`func NewGetGroupPosts200ResponseWithDefaults() *GetGroupPosts200Response`

NewGetGroupPosts200ResponseWithDefaults instantiates a new GetGroupPosts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPosts

`func (o *GetGroupPosts200Response) GetPosts() []GroupPost`

GetPosts returns the Posts field if non-nil, zero value otherwise.

### GetPostsOk

`func (o *GetGroupPosts200Response) GetPostsOk() (*[]GroupPost, bool)`

GetPostsOk returns a tuple with the Posts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosts

`func (o *GetGroupPosts200Response) SetPosts(v []GroupPost)`

SetPosts sets Posts field to given value.

### HasPosts

`func (o *GetGroupPosts200Response) HasPosts() bool`

HasPosts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


