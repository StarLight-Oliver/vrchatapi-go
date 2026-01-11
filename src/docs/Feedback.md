# Feedback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommenterId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**CommenterName** | **string** |  | 
**ContentAuthorId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**ContentAuthorName** | **NullableString** |  | 
**ContentId** | **string** |  | 
**ContentName** | Pointer to **string** |  | [optional] 
**ContentType** | **string** |  | 
**ContentVersion** | **NullableInt32** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** |  | 
**Reason** | **string** |  | 
**Tags** | **[]string** |  | 
**Type** | **string** |  | 

## Methods

### NewFeedback

`func NewFeedback(commenterId string, commenterName string, contentAuthorId string, contentAuthorName NullableString, contentId string, contentType string, contentVersion NullableInt32, id string, reason string, tags []string, type_ string, ) *Feedback`

NewFeedback instantiates a new Feedback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackWithDefaults

`func NewFeedbackWithDefaults() *Feedback`

NewFeedbackWithDefaults instantiates a new Feedback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommenterId

`func (o *Feedback) GetCommenterId() string`

GetCommenterId returns the CommenterId field if non-nil, zero value otherwise.

### GetCommenterIdOk

`func (o *Feedback) GetCommenterIdOk() (*string, bool)`

GetCommenterIdOk returns a tuple with the CommenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterId

`func (o *Feedback) SetCommenterId(v string)`

SetCommenterId sets CommenterId field to given value.


### GetCommenterName

`func (o *Feedback) GetCommenterName() string`

GetCommenterName returns the CommenterName field if non-nil, zero value otherwise.

### GetCommenterNameOk

`func (o *Feedback) GetCommenterNameOk() (*string, bool)`

GetCommenterNameOk returns a tuple with the CommenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterName

`func (o *Feedback) SetCommenterName(v string)`

SetCommenterName sets CommenterName field to given value.


### GetContentAuthorId

`func (o *Feedback) GetContentAuthorId() string`

GetContentAuthorId returns the ContentAuthorId field if non-nil, zero value otherwise.

### GetContentAuthorIdOk

`func (o *Feedback) GetContentAuthorIdOk() (*string, bool)`

GetContentAuthorIdOk returns a tuple with the ContentAuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentAuthorId

`func (o *Feedback) SetContentAuthorId(v string)`

SetContentAuthorId sets ContentAuthorId field to given value.


### GetContentAuthorName

`func (o *Feedback) GetContentAuthorName() string`

GetContentAuthorName returns the ContentAuthorName field if non-nil, zero value otherwise.

### GetContentAuthorNameOk

`func (o *Feedback) GetContentAuthorNameOk() (*string, bool)`

GetContentAuthorNameOk returns a tuple with the ContentAuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentAuthorName

`func (o *Feedback) SetContentAuthorName(v string)`

SetContentAuthorName sets ContentAuthorName field to given value.


### SetContentAuthorNameNil

`func (o *Feedback) SetContentAuthorNameNil(b bool)`

 SetContentAuthorNameNil sets the value for ContentAuthorName to be an explicit nil

### UnsetContentAuthorName
`func (o *Feedback) UnsetContentAuthorName()`

UnsetContentAuthorName ensures that no value is present for ContentAuthorName, not even an explicit nil
### GetContentId

`func (o *Feedback) GetContentId() string`

GetContentId returns the ContentId field if non-nil, zero value otherwise.

### GetContentIdOk

`func (o *Feedback) GetContentIdOk() (*string, bool)`

GetContentIdOk returns a tuple with the ContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentId

`func (o *Feedback) SetContentId(v string)`

SetContentId sets ContentId field to given value.


### GetContentName

`func (o *Feedback) GetContentName() string`

GetContentName returns the ContentName field if non-nil, zero value otherwise.

### GetContentNameOk

`func (o *Feedback) GetContentNameOk() (*string, bool)`

GetContentNameOk returns a tuple with the ContentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentName

`func (o *Feedback) SetContentName(v string)`

SetContentName sets ContentName field to given value.

### HasContentName

`func (o *Feedback) HasContentName() bool`

HasContentName returns a boolean if a field has been set.

### GetContentType

`func (o *Feedback) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *Feedback) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *Feedback) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetContentVersion

`func (o *Feedback) GetContentVersion() int32`

GetContentVersion returns the ContentVersion field if non-nil, zero value otherwise.

### GetContentVersionOk

`func (o *Feedback) GetContentVersionOk() (*int32, bool)`

GetContentVersionOk returns a tuple with the ContentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentVersion

`func (o *Feedback) SetContentVersion(v int32)`

SetContentVersion sets ContentVersion field to given value.


### SetContentVersionNil

`func (o *Feedback) SetContentVersionNil(b bool)`

 SetContentVersionNil sets the value for ContentVersion to be an explicit nil

### UnsetContentVersion
`func (o *Feedback) UnsetContentVersion()`

UnsetContentVersion ensures that no value is present for ContentVersion, not even an explicit nil
### GetDescription

`func (o *Feedback) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Feedback) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Feedback) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Feedback) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Feedback) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Feedback) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *Feedback) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Feedback) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Feedback) SetId(v string)`

SetId sets Id field to given value.


### GetReason

`func (o *Feedback) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Feedback) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Feedback) SetReason(v string)`

SetReason sets Reason field to given value.


### GetTags

`func (o *Feedback) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Feedback) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Feedback) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetType

`func (o *Feedback) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Feedback) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Feedback) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


