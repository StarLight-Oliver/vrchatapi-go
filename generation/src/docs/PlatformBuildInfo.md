# PlatformBuildInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinBuildNumber** | **int32** | Minimum build number required for the platform | 
**RedirectionAddress** | Pointer to **string** | Redirection URL for updating the app | [optional] 

## Methods

### NewPlatformBuildInfo

`func NewPlatformBuildInfo(minBuildNumber int32, ) *PlatformBuildInfo`

NewPlatformBuildInfo instantiates a new PlatformBuildInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformBuildInfoWithDefaults

`func NewPlatformBuildInfoWithDefaults() *PlatformBuildInfo`

NewPlatformBuildInfoWithDefaults instantiates a new PlatformBuildInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinBuildNumber

`func (o *PlatformBuildInfo) GetMinBuildNumber() int32`

GetMinBuildNumber returns the MinBuildNumber field if non-nil, zero value otherwise.

### GetMinBuildNumberOk

`func (o *PlatformBuildInfo) GetMinBuildNumberOk() (*int32, bool)`

GetMinBuildNumberOk returns a tuple with the MinBuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBuildNumber

`func (o *PlatformBuildInfo) SetMinBuildNumber(v int32)`

SetMinBuildNumber sets MinBuildNumber field to given value.


### GetRedirectionAddress

`func (o *PlatformBuildInfo) GetRedirectionAddress() string`

GetRedirectionAddress returns the RedirectionAddress field if non-nil, zero value otherwise.

### GetRedirectionAddressOk

`func (o *PlatformBuildInfo) GetRedirectionAddressOk() (*string, bool)`

GetRedirectionAddressOk returns a tuple with the RedirectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectionAddress

`func (o *PlatformBuildInfo) SetRedirectionAddress(v string)`

SetRedirectionAddress sets RedirectionAddress field to given value.

### HasRedirectionAddress

`func (o *PlatformBuildInfo) HasRedirectionAddress() bool`

HasRedirectionAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


