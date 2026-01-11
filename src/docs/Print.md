# Print

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**AuthorName** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Files** | [**PrintFiles**](PrintFiles.md) |  | 
**Id** | **string** |  | 
**Note** | **string** |  | 
**OwnerId** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 
**Timestamp** | **time.Time** |  | 
**WorldId** | **string** | WorldID be \&quot;offline\&quot; on User profiles if you are not friends with that user. | 
**WorldName** | **string** |  | 

## Methods

### NewPrint

`func NewPrint(authorId string, authorName string, createdAt time.Time, files PrintFiles, id string, note string, timestamp time.Time, worldId string, worldName string, ) *Print`

NewPrint instantiates a new Print object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrintWithDefaults

`func NewPrintWithDefaults() *Print`

NewPrintWithDefaults instantiates a new Print object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorId

`func (o *Print) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *Print) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *Print) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.


### GetAuthorName

`func (o *Print) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *Print) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *Print) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.


### GetCreatedAt

`func (o *Print) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Print) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Print) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetFiles

`func (o *Print) GetFiles() PrintFiles`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *Print) GetFilesOk() (*PrintFiles, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *Print) SetFiles(v PrintFiles)`

SetFiles sets Files field to given value.


### GetId

`func (o *Print) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Print) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Print) SetId(v string)`

SetId sets Id field to given value.


### GetNote

`func (o *Print) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *Print) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *Print) SetNote(v string)`

SetNote sets Note field to given value.


### GetOwnerId

`func (o *Print) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Print) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Print) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *Print) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetTimestamp

`func (o *Print) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Print) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Print) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetWorldId

`func (o *Print) GetWorldId() string`

GetWorldId returns the WorldId field if non-nil, zero value otherwise.

### GetWorldIdOk

`func (o *Print) GetWorldIdOk() (*string, bool)`

GetWorldIdOk returns a tuple with the WorldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorldId

`func (o *Print) SetWorldId(v string)`

SetWorldId sets WorldId field to given value.


### GetWorldName

`func (o *Print) GetWorldName() string`

GetWorldName returns the WorldName field if non-nil, zero value otherwise.

### GetWorldNameOk

`func (o *Print) GetWorldNameOk() (*string, bool)`

GetWorldNameOk returns a tuple with the WorldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorldName

`func (o *Print) SetWorldName(v string)`

SetWorldName sets WorldName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


