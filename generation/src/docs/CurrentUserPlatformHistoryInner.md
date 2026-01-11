# CurrentUserPlatformHistoryInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsMobile** | Pointer to **bool** |  | [optional] 
**Platform** | Pointer to **string** | This is normally &#x60;android&#x60;, &#x60;ios&#x60;, &#x60;standalonewindows&#x60;, &#x60;web&#x60;, or the empty value &#x60;&#x60;, but also supposedly can be any random Unity version such as &#x60;2019.2.4-801-Release&#x60; or &#x60;2019.2.2-772-Release&#x60; or even &#x60;unknownplatform&#x60;. | [optional] 
**Recorded** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCurrentUserPlatformHistoryInner

`func NewCurrentUserPlatformHistoryInner() *CurrentUserPlatformHistoryInner`

NewCurrentUserPlatformHistoryInner instantiates a new CurrentUserPlatformHistoryInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentUserPlatformHistoryInnerWithDefaults

`func NewCurrentUserPlatformHistoryInnerWithDefaults() *CurrentUserPlatformHistoryInner`

NewCurrentUserPlatformHistoryInnerWithDefaults instantiates a new CurrentUserPlatformHistoryInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsMobile

`func (o *CurrentUserPlatformHistoryInner) GetIsMobile() bool`

GetIsMobile returns the IsMobile field if non-nil, zero value otherwise.

### GetIsMobileOk

`func (o *CurrentUserPlatformHistoryInner) GetIsMobileOk() (*bool, bool)`

GetIsMobileOk returns a tuple with the IsMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMobile

`func (o *CurrentUserPlatformHistoryInner) SetIsMobile(v bool)`

SetIsMobile sets IsMobile field to given value.

### HasIsMobile

`func (o *CurrentUserPlatformHistoryInner) HasIsMobile() bool`

HasIsMobile returns a boolean if a field has been set.

### GetPlatform

`func (o *CurrentUserPlatformHistoryInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CurrentUserPlatformHistoryInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CurrentUserPlatformHistoryInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CurrentUserPlatformHistoryInner) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetRecorded

`func (o *CurrentUserPlatformHistoryInner) GetRecorded() time.Time`

GetRecorded returns the Recorded field if non-nil, zero value otherwise.

### GetRecordedOk

`func (o *CurrentUserPlatformHistoryInner) GetRecordedOk() (*time.Time, bool)`

GetRecordedOk returns a tuple with the Recorded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecorded

`func (o *CurrentUserPlatformHistoryInner) SetRecorded(v time.Time)`

SetRecorded sets Recorded field to given value.

### HasRecorded

`func (o *CurrentUserPlatformHistoryInner) HasRecorded() bool`

HasRecorded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


