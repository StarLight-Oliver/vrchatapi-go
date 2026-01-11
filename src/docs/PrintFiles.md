# PrintFiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** | Link to file, e.g. https://api.vrchat.cloud/api/1/file/file_66fe782d-f2bd-4462-9761-1d766d7b2b26/1/file | [optional] 

## Methods

### NewPrintFiles

`func NewPrintFiles() *PrintFiles`

NewPrintFiles instantiates a new PrintFiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrintFilesWithDefaults

`func NewPrintFilesWithDefaults() *PrintFiles`

NewPrintFilesWithDefaults instantiates a new PrintFiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *PrintFiles) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *PrintFiles) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *PrintFiles) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *PrintFiles) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetImage

`func (o *PrintFiles) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PrintFiles) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PrintFiles) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *PrintFiles) HasImage() bool`

HasImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


