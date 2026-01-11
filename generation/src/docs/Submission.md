# Submission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentId** | **string** | Either world ID or avatar ID | 
**CreatedAt** | **time.Time** |  | 
**Description** | **string** |  | 
**Id** | **string** |  | 
**JamId** | **string** |  | 
**RatingScore** | Pointer to **int32** |  | [optional] 
**SubmitterId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 

## Methods

### NewSubmission

`func NewSubmission(contentId string, createdAt time.Time, description string, id string, jamId string, submitterId string, ) *Submission`

NewSubmission instantiates a new Submission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionWithDefaults

`func NewSubmissionWithDefaults() *Submission`

NewSubmissionWithDefaults instantiates a new Submission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentId

`func (o *Submission) GetContentId() string`

GetContentId returns the ContentId field if non-nil, zero value otherwise.

### GetContentIdOk

`func (o *Submission) GetContentIdOk() (*string, bool)`

GetContentIdOk returns a tuple with the ContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentId

`func (o *Submission) SetContentId(v string)`

SetContentId sets ContentId field to given value.


### GetCreatedAt

`func (o *Submission) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Submission) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Submission) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *Submission) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Submission) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Submission) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *Submission) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Submission) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Submission) SetId(v string)`

SetId sets Id field to given value.


### GetJamId

`func (o *Submission) GetJamId() string`

GetJamId returns the JamId field if non-nil, zero value otherwise.

### GetJamIdOk

`func (o *Submission) GetJamIdOk() (*string, bool)`

GetJamIdOk returns a tuple with the JamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJamId

`func (o *Submission) SetJamId(v string)`

SetJamId sets JamId field to given value.


### GetRatingScore

`func (o *Submission) GetRatingScore() int32`

GetRatingScore returns the RatingScore field if non-nil, zero value otherwise.

### GetRatingScoreOk

`func (o *Submission) GetRatingScoreOk() (*int32, bool)`

GetRatingScoreOk returns a tuple with the RatingScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingScore

`func (o *Submission) SetRatingScore(v int32)`

SetRatingScore sets RatingScore field to given value.

### HasRatingScore

`func (o *Submission) HasRatingScore() bool`

HasRatingScore returns a boolean if a field has been set.

### GetSubmitterId

`func (o *Submission) GetSubmitterId() string`

GetSubmitterId returns the SubmitterId field if non-nil, zero value otherwise.

### GetSubmitterIdOk

`func (o *Submission) GetSubmitterIdOk() (*string, bool)`

GetSubmitterIdOk returns a tuple with the SubmitterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitterId

`func (o *Submission) SetSubmitterId(v string)`

SetSubmitterId sets SubmitterId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


