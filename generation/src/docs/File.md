# File

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnimationStyle** | Pointer to [**ImageAnimationStyle**](ImageAnimationStyle.md) |  | [optional] 
**Extension** | **string** |  | 
**Frames** | Pointer to **int32** | The number of frames for animated spritesheet images. | [optional] 
**FramesOverTime** | Pointer to **int32** | The frames per second for animated spritesheet images. | [optional] 
**Id** | **string** |  | 
**LoopStyle** | Pointer to [**ImageLoopStyle**](ImageLoopStyle.md) |  | [optional] [default to LINEAR]
**MaskTag** | Pointer to [**ImageMask**](ImageMask.md) |  | [optional] [default to SQUARE]
**MimeType** | [**MIMEType**](MIMEType.md) |  | [default to IMAGE_JPEG]
**ModifiedThumbnailFileName** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**OwnerId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**Tags** | **[]string** |   | 
**Versions** | [**[]FileVersion**](FileVersion.md) |   | 

## Methods

### NewFile

`func NewFile(extension string, id string, mimeType MIMEType, name string, ownerId string, tags []string, versions []FileVersion, ) *File`

NewFile instantiates a new File object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileWithDefaults

`func NewFileWithDefaults() *File`

NewFileWithDefaults instantiates a new File object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnimationStyle

`func (o *File) GetAnimationStyle() ImageAnimationStyle`

GetAnimationStyle returns the AnimationStyle field if non-nil, zero value otherwise.

### GetAnimationStyleOk

`func (o *File) GetAnimationStyleOk() (*ImageAnimationStyle, bool)`

GetAnimationStyleOk returns a tuple with the AnimationStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnimationStyle

`func (o *File) SetAnimationStyle(v ImageAnimationStyle)`

SetAnimationStyle sets AnimationStyle field to given value.

### HasAnimationStyle

`func (o *File) HasAnimationStyle() bool`

HasAnimationStyle returns a boolean if a field has been set.

### GetExtension

`func (o *File) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *File) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *File) SetExtension(v string)`

SetExtension sets Extension field to given value.


### GetFrames

`func (o *File) GetFrames() int32`

GetFrames returns the Frames field if non-nil, zero value otherwise.

### GetFramesOk

`func (o *File) GetFramesOk() (*int32, bool)`

GetFramesOk returns a tuple with the Frames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrames

`func (o *File) SetFrames(v int32)`

SetFrames sets Frames field to given value.

### HasFrames

`func (o *File) HasFrames() bool`

HasFrames returns a boolean if a field has been set.

### GetFramesOverTime

`func (o *File) GetFramesOverTime() int32`

GetFramesOverTime returns the FramesOverTime field if non-nil, zero value otherwise.

### GetFramesOverTimeOk

`func (o *File) GetFramesOverTimeOk() (*int32, bool)`

GetFramesOverTimeOk returns a tuple with the FramesOverTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramesOverTime

`func (o *File) SetFramesOverTime(v int32)`

SetFramesOverTime sets FramesOverTime field to given value.

### HasFramesOverTime

`func (o *File) HasFramesOverTime() bool`

HasFramesOverTime returns a boolean if a field has been set.

### GetId

`func (o *File) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *File) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *File) SetId(v string)`

SetId sets Id field to given value.


### GetLoopStyle

`func (o *File) GetLoopStyle() ImageLoopStyle`

GetLoopStyle returns the LoopStyle field if non-nil, zero value otherwise.

### GetLoopStyleOk

`func (o *File) GetLoopStyleOk() (*ImageLoopStyle, bool)`

GetLoopStyleOk returns a tuple with the LoopStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopStyle

`func (o *File) SetLoopStyle(v ImageLoopStyle)`

SetLoopStyle sets LoopStyle field to given value.

### HasLoopStyle

`func (o *File) HasLoopStyle() bool`

HasLoopStyle returns a boolean if a field has been set.

### GetMaskTag

`func (o *File) GetMaskTag() ImageMask`

GetMaskTag returns the MaskTag field if non-nil, zero value otherwise.

### GetMaskTagOk

`func (o *File) GetMaskTagOk() (*ImageMask, bool)`

GetMaskTagOk returns a tuple with the MaskTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskTag

`func (o *File) SetMaskTag(v ImageMask)`

SetMaskTag sets MaskTag field to given value.

### HasMaskTag

`func (o *File) HasMaskTag() bool`

HasMaskTag returns a boolean if a field has been set.

### GetMimeType

`func (o *File) GetMimeType() MIMEType`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *File) GetMimeTypeOk() (*MIMEType, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *File) SetMimeType(v MIMEType)`

SetMimeType sets MimeType field to given value.


### GetModifiedThumbnailFileName

`func (o *File) GetModifiedThumbnailFileName() string`

GetModifiedThumbnailFileName returns the ModifiedThumbnailFileName field if non-nil, zero value otherwise.

### GetModifiedThumbnailFileNameOk

`func (o *File) GetModifiedThumbnailFileNameOk() (*string, bool)`

GetModifiedThumbnailFileNameOk returns a tuple with the ModifiedThumbnailFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedThumbnailFileName

`func (o *File) SetModifiedThumbnailFileName(v string)`

SetModifiedThumbnailFileName sets ModifiedThumbnailFileName field to given value.

### HasModifiedThumbnailFileName

`func (o *File) HasModifiedThumbnailFileName() bool`

HasModifiedThumbnailFileName returns a boolean if a field has been set.

### GetName

`func (o *File) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *File) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *File) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerId

`func (o *File) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *File) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *File) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetTags

`func (o *File) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *File) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *File) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetVersions

`func (o *File) GetVersions() []FileVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *File) GetVersionsOk() (*[]FileVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *File) SetVersions(v []FileVersion)`

SetVersions sets Versions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


