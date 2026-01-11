# GetUserGroupInstances200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FetchedAt** | Pointer to **time.Time** |  | [optional] 
**Instances** | Pointer to [**[]Instance**](Instance.md) |  | [optional] 

## Methods

### NewGetUserGroupInstances200Response

`func NewGetUserGroupInstances200Response() *GetUserGroupInstances200Response`

NewGetUserGroupInstances200Response instantiates a new GetUserGroupInstances200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserGroupInstances200ResponseWithDefaults

`func NewGetUserGroupInstances200ResponseWithDefaults() *GetUserGroupInstances200Response`

NewGetUserGroupInstances200ResponseWithDefaults instantiates a new GetUserGroupInstances200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFetchedAt

`func (o *GetUserGroupInstances200Response) GetFetchedAt() time.Time`

GetFetchedAt returns the FetchedAt field if non-nil, zero value otherwise.

### GetFetchedAtOk

`func (o *GetUserGroupInstances200Response) GetFetchedAtOk() (*time.Time, bool)`

GetFetchedAtOk returns a tuple with the FetchedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchedAt

`func (o *GetUserGroupInstances200Response) SetFetchedAt(v time.Time)`

SetFetchedAt sets FetchedAt field to given value.

### HasFetchedAt

`func (o *GetUserGroupInstances200Response) HasFetchedAt() bool`

HasFetchedAt returns a boolean if a field has been set.

### GetInstances

`func (o *GetUserGroupInstances200Response) GetInstances() []Instance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *GetUserGroupInstances200Response) GetInstancesOk() (*[]Instance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *GetUserGroupInstances200Response) SetInstances(v []Instance)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *GetUserGroupInstances200Response) HasInstances() bool`

HasInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


