# ServiceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**Id** | **string** | The id of this service, NOT the id of the thing this service was requested for. | 
**Progress** | **[]map[string]interface{}** |  | 
**RequesterUserId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**State** | **string** |  | 
**SubjectId** | **string** | The id of the thing this service was requested for. | 
**SubjectType** | **string** | The kind of the thing this service was requested for. | 
**Type** | **string** | The kind of service that was requested. | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewServiceStatus

`func NewServiceStatus(createdAt time.Time, id string, progress []map[string]interface{}, requesterUserId string, state string, subjectId string, subjectType string, type_ string, updatedAt time.Time, ) *ServiceStatus`

NewServiceStatus instantiates a new ServiceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceStatusWithDefaults

`func NewServiceStatusWithDefaults() *ServiceStatus`

NewServiceStatusWithDefaults instantiates a new ServiceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ServiceStatus) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceStatus) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceStatus) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *ServiceStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceStatus) SetId(v string)`

SetId sets Id field to given value.


### GetProgress

`func (o *ServiceStatus) GetProgress() []map[string]interface{}`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ServiceStatus) GetProgressOk() (*[]map[string]interface{}, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ServiceStatus) SetProgress(v []map[string]interface{})`

SetProgress sets Progress field to given value.


### GetRequesterUserId

`func (o *ServiceStatus) GetRequesterUserId() string`

GetRequesterUserId returns the RequesterUserId field if non-nil, zero value otherwise.

### GetRequesterUserIdOk

`func (o *ServiceStatus) GetRequesterUserIdOk() (*string, bool)`

GetRequesterUserIdOk returns a tuple with the RequesterUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterUserId

`func (o *ServiceStatus) SetRequesterUserId(v string)`

SetRequesterUserId sets RequesterUserId field to given value.


### GetState

`func (o *ServiceStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ServiceStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ServiceStatus) SetState(v string)`

SetState sets State field to given value.


### GetSubjectId

`func (o *ServiceStatus) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *ServiceStatus) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *ServiceStatus) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.


### GetSubjectType

`func (o *ServiceStatus) GetSubjectType() string`

GetSubjectType returns the SubjectType field if non-nil, zero value otherwise.

### GetSubjectTypeOk

`func (o *ServiceStatus) GetSubjectTypeOk() (*string, bool)`

GetSubjectTypeOk returns a tuple with the SubjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectType

`func (o *ServiceStatus) SetSubjectType(v string)`

SetSubjectType sets SubjectType field to given value.


### GetType

`func (o *ServiceStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceStatus) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *ServiceStatus) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServiceStatus) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServiceStatus) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


