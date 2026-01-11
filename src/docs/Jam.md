# Jam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**Id** | **string** |  | 
**IsVisible** | **bool** |  | 
**MoreInfo** | **string** |  | 
**State** | **string** | One of: - submissions_open - closed | 
**StateChangeDates** | [**JamStateChangeDates**](JamStateChangeDates.md) |  | 
**SubmissionContentGateDate** | **NullableTime** |  | 
**SubmissionContentGated** | **bool** |  | 
**Title** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewJam

`func NewJam(description string, id string, isVisible bool, moreInfo string, state string, stateChangeDates JamStateChangeDates, submissionContentGateDate NullableTime, submissionContentGated bool, title string, updatedAt time.Time, ) *Jam`

NewJam instantiates a new Jam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJamWithDefaults

`func NewJamWithDefaults() *Jam`

NewJamWithDefaults instantiates a new Jam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Jam) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Jam) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Jam) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *Jam) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Jam) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Jam) SetId(v string)`

SetId sets Id field to given value.


### GetIsVisible

`func (o *Jam) GetIsVisible() bool`

GetIsVisible returns the IsVisible field if non-nil, zero value otherwise.

### GetIsVisibleOk

`func (o *Jam) GetIsVisibleOk() (*bool, bool)`

GetIsVisibleOk returns a tuple with the IsVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVisible

`func (o *Jam) SetIsVisible(v bool)`

SetIsVisible sets IsVisible field to given value.


### GetMoreInfo

`func (o *Jam) GetMoreInfo() string`

GetMoreInfo returns the MoreInfo field if non-nil, zero value otherwise.

### GetMoreInfoOk

`func (o *Jam) GetMoreInfoOk() (*string, bool)`

GetMoreInfoOk returns a tuple with the MoreInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoreInfo

`func (o *Jam) SetMoreInfo(v string)`

SetMoreInfo sets MoreInfo field to given value.


### GetState

`func (o *Jam) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Jam) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Jam) SetState(v string)`

SetState sets State field to given value.


### GetStateChangeDates

`func (o *Jam) GetStateChangeDates() JamStateChangeDates`

GetStateChangeDates returns the StateChangeDates field if non-nil, zero value otherwise.

### GetStateChangeDatesOk

`func (o *Jam) GetStateChangeDatesOk() (*JamStateChangeDates, bool)`

GetStateChangeDatesOk returns a tuple with the StateChangeDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateChangeDates

`func (o *Jam) SetStateChangeDates(v JamStateChangeDates)`

SetStateChangeDates sets StateChangeDates field to given value.


### GetSubmissionContentGateDate

`func (o *Jam) GetSubmissionContentGateDate() time.Time`

GetSubmissionContentGateDate returns the SubmissionContentGateDate field if non-nil, zero value otherwise.

### GetSubmissionContentGateDateOk

`func (o *Jam) GetSubmissionContentGateDateOk() (*time.Time, bool)`

GetSubmissionContentGateDateOk returns a tuple with the SubmissionContentGateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionContentGateDate

`func (o *Jam) SetSubmissionContentGateDate(v time.Time)`

SetSubmissionContentGateDate sets SubmissionContentGateDate field to given value.


### SetSubmissionContentGateDateNil

`func (o *Jam) SetSubmissionContentGateDateNil(b bool)`

 SetSubmissionContentGateDateNil sets the value for SubmissionContentGateDate to be an explicit nil

### UnsetSubmissionContentGateDate
`func (o *Jam) UnsetSubmissionContentGateDate()`

UnsetSubmissionContentGateDate ensures that no value is present for SubmissionContentGateDate, not even an explicit nil
### GetSubmissionContentGated

`func (o *Jam) GetSubmissionContentGated() bool`

GetSubmissionContentGated returns the SubmissionContentGated field if non-nil, zero value otherwise.

### GetSubmissionContentGatedOk

`func (o *Jam) GetSubmissionContentGatedOk() (*bool, bool)`

GetSubmissionContentGatedOk returns a tuple with the SubmissionContentGated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionContentGated

`func (o *Jam) SetSubmissionContentGated(v bool)`

SetSubmissionContentGated sets SubmissionContentGated field to given value.


### GetTitle

`func (o *Jam) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Jam) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Jam) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUpdatedAt

`func (o *Jam) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Jam) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Jam) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


