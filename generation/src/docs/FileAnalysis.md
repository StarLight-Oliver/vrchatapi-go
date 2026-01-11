# FileAnalysis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarStats** | [**FileAnalysisAvatarStats**](FileAnalysisAvatarStats.md) |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**EncryptionKey** | Pointer to **string** |  | [optional] 
**FileSize** | **int32** |  | 
**PerformanceRating** | Pointer to **string** |  | [optional] 
**Success** | **bool** |  | 
**UncompressedSize** | **int32** |  | 

## Methods

### NewFileAnalysis

`func NewFileAnalysis(avatarStats FileAnalysisAvatarStats, fileSize int32, success bool, uncompressedSize int32, ) *FileAnalysis`

NewFileAnalysis instantiates a new FileAnalysis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileAnalysisWithDefaults

`func NewFileAnalysisWithDefaults() *FileAnalysis`

NewFileAnalysisWithDefaults instantiates a new FileAnalysis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarStats

`func (o *FileAnalysis) GetAvatarStats() FileAnalysisAvatarStats`

GetAvatarStats returns the AvatarStats field if non-nil, zero value otherwise.

### GetAvatarStatsOk

`func (o *FileAnalysis) GetAvatarStatsOk() (*FileAnalysisAvatarStats, bool)`

GetAvatarStatsOk returns a tuple with the AvatarStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarStats

`func (o *FileAnalysis) SetAvatarStats(v FileAnalysisAvatarStats)`

SetAvatarStats sets AvatarStats field to given value.


### GetCreatedAt

`func (o *FileAnalysis) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FileAnalysis) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FileAnalysis) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FileAnalysis) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEncryptionKey

`func (o *FileAnalysis) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *FileAnalysis) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *FileAnalysis) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *FileAnalysis) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### GetFileSize

`func (o *FileAnalysis) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *FileAnalysis) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *FileAnalysis) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.


### GetPerformanceRating

`func (o *FileAnalysis) GetPerformanceRating() string`

GetPerformanceRating returns the PerformanceRating field if non-nil, zero value otherwise.

### GetPerformanceRatingOk

`func (o *FileAnalysis) GetPerformanceRatingOk() (*string, bool)`

GetPerformanceRatingOk returns a tuple with the PerformanceRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceRating

`func (o *FileAnalysis) SetPerformanceRating(v string)`

SetPerformanceRating sets PerformanceRating field to given value.

### HasPerformanceRating

`func (o *FileAnalysis) HasPerformanceRating() bool`

HasPerformanceRating returns a boolean if a field has been set.

### GetSuccess

`func (o *FileAnalysis) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *FileAnalysis) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *FileAnalysis) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetUncompressedSize

`func (o *FileAnalysis) GetUncompressedSize() int32`

GetUncompressedSize returns the UncompressedSize field if non-nil, zero value otherwise.

### GetUncompressedSizeOk

`func (o *FileAnalysis) GetUncompressedSizeOk() (*int32, bool)`

GetUncompressedSizeOk returns a tuple with the UncompressedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncompressedSize

`func (o *FileAnalysis) SetUncompressedSize(v int32)`

SetUncompressedSize sets UncompressedSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


