# UpdateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedTOSVersion** | Pointer to **int32** |  | [optional] 
**Bio** | Pointer to **string** |  | [optional] 
**BioLinks** | Pointer to **[]string** |  | [optional] 
**Birthday** | Pointer to **string** |  | [optional] 
**ContentFilters** | Pointer to [**[]ContentFilter**](ContentFilter.md) | These tags begin with &#x60;content_&#x60; and control content gating | [optional] 
**CurrentPassword** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** | MUST specify currentPassword as well to change display name | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**HasSharedConnectionsOptOut** | Pointer to **bool** | Opt out of the Mutuals feature | [optional] 
**IsBoopingEnabled** | Pointer to **bool** |  | [optional] 
**Password** | Pointer to **string** | MUST specify currentPassword as well to change password | [optional] 
**Pronouns** | Pointer to **string** |  | [optional] 
**RevertDisplayName** | Pointer to **bool** | MUST specify currentPassword as well to revert display name | [optional] 
**Status** | Pointer to [**UserStatus**](UserStatus.md) |  | [optional] [default to OFFLINE]
**StatusDescription** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |   | [optional] 
**Unsubscribe** | Pointer to **bool** |  | [optional] 
**UserIcon** | Pointer to **string** | MUST be a valid VRChat /file/ url. | [optional] 

## Methods

### NewUpdateUserRequest

`func NewUpdateUserRequest() *UpdateUserRequest`

NewUpdateUserRequest instantiates a new UpdateUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserRequestWithDefaults

`func NewUpdateUserRequestWithDefaults() *UpdateUserRequest`

NewUpdateUserRequestWithDefaults instantiates a new UpdateUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedTOSVersion

`func (o *UpdateUserRequest) GetAcceptedTOSVersion() int32`

GetAcceptedTOSVersion returns the AcceptedTOSVersion field if non-nil, zero value otherwise.

### GetAcceptedTOSVersionOk

`func (o *UpdateUserRequest) GetAcceptedTOSVersionOk() (*int32, bool)`

GetAcceptedTOSVersionOk returns a tuple with the AcceptedTOSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedTOSVersion

`func (o *UpdateUserRequest) SetAcceptedTOSVersion(v int32)`

SetAcceptedTOSVersion sets AcceptedTOSVersion field to given value.

### HasAcceptedTOSVersion

`func (o *UpdateUserRequest) HasAcceptedTOSVersion() bool`

HasAcceptedTOSVersion returns a boolean if a field has been set.

### GetBio

`func (o *UpdateUserRequest) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *UpdateUserRequest) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *UpdateUserRequest) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *UpdateUserRequest) HasBio() bool`

HasBio returns a boolean if a field has been set.

### GetBioLinks

`func (o *UpdateUserRequest) GetBioLinks() []string`

GetBioLinks returns the BioLinks field if non-nil, zero value otherwise.

### GetBioLinksOk

`func (o *UpdateUserRequest) GetBioLinksOk() (*[]string, bool)`

GetBioLinksOk returns a tuple with the BioLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBioLinks

`func (o *UpdateUserRequest) SetBioLinks(v []string)`

SetBioLinks sets BioLinks field to given value.

### HasBioLinks

`func (o *UpdateUserRequest) HasBioLinks() bool`

HasBioLinks returns a boolean if a field has been set.

### GetBirthday

`func (o *UpdateUserRequest) GetBirthday() string`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *UpdateUserRequest) GetBirthdayOk() (*string, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *UpdateUserRequest) SetBirthday(v string)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *UpdateUserRequest) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### GetContentFilters

`func (o *UpdateUserRequest) GetContentFilters() []ContentFilter`

GetContentFilters returns the ContentFilters field if non-nil, zero value otherwise.

### GetContentFiltersOk

`func (o *UpdateUserRequest) GetContentFiltersOk() (*[]ContentFilter, bool)`

GetContentFiltersOk returns a tuple with the ContentFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFilters

`func (o *UpdateUserRequest) SetContentFilters(v []ContentFilter)`

SetContentFilters sets ContentFilters field to given value.

### HasContentFilters

`func (o *UpdateUserRequest) HasContentFilters() bool`

HasContentFilters returns a boolean if a field has been set.

### GetCurrentPassword

`func (o *UpdateUserRequest) GetCurrentPassword() string`

GetCurrentPassword returns the CurrentPassword field if non-nil, zero value otherwise.

### GetCurrentPasswordOk

`func (o *UpdateUserRequest) GetCurrentPasswordOk() (*string, bool)`

GetCurrentPasswordOk returns a tuple with the CurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPassword

`func (o *UpdateUserRequest) SetCurrentPassword(v string)`

SetCurrentPassword sets CurrentPassword field to given value.

### HasCurrentPassword

`func (o *UpdateUserRequest) HasCurrentPassword() bool`

HasCurrentPassword returns a boolean if a field has been set.

### GetDisplayName

`func (o *UpdateUserRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateUserRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateUserRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateUserRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateUserRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetHasSharedConnectionsOptOut

`func (o *UpdateUserRequest) GetHasSharedConnectionsOptOut() bool`

GetHasSharedConnectionsOptOut returns the HasSharedConnectionsOptOut field if non-nil, zero value otherwise.

### GetHasSharedConnectionsOptOutOk

`func (o *UpdateUserRequest) GetHasSharedConnectionsOptOutOk() (*bool, bool)`

GetHasSharedConnectionsOptOutOk returns a tuple with the HasSharedConnectionsOptOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSharedConnectionsOptOut

`func (o *UpdateUserRequest) SetHasSharedConnectionsOptOut(v bool)`

SetHasSharedConnectionsOptOut sets HasSharedConnectionsOptOut field to given value.

### HasHasSharedConnectionsOptOut

`func (o *UpdateUserRequest) HasHasSharedConnectionsOptOut() bool`

HasHasSharedConnectionsOptOut returns a boolean if a field has been set.

### GetIsBoopingEnabled

`func (o *UpdateUserRequest) GetIsBoopingEnabled() bool`

GetIsBoopingEnabled returns the IsBoopingEnabled field if non-nil, zero value otherwise.

### GetIsBoopingEnabledOk

`func (o *UpdateUserRequest) GetIsBoopingEnabledOk() (*bool, bool)`

GetIsBoopingEnabledOk returns a tuple with the IsBoopingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBoopingEnabled

`func (o *UpdateUserRequest) SetIsBoopingEnabled(v bool)`

SetIsBoopingEnabled sets IsBoopingEnabled field to given value.

### HasIsBoopingEnabled

`func (o *UpdateUserRequest) HasIsBoopingEnabled() bool`

HasIsBoopingEnabled returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateUserRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateUserRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateUserRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateUserRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPronouns

`func (o *UpdateUserRequest) GetPronouns() string`

GetPronouns returns the Pronouns field if non-nil, zero value otherwise.

### GetPronounsOk

`func (o *UpdateUserRequest) GetPronounsOk() (*string, bool)`

GetPronounsOk returns a tuple with the Pronouns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPronouns

`func (o *UpdateUserRequest) SetPronouns(v string)`

SetPronouns sets Pronouns field to given value.

### HasPronouns

`func (o *UpdateUserRequest) HasPronouns() bool`

HasPronouns returns a boolean if a field has been set.

### GetRevertDisplayName

`func (o *UpdateUserRequest) GetRevertDisplayName() bool`

GetRevertDisplayName returns the RevertDisplayName field if non-nil, zero value otherwise.

### GetRevertDisplayNameOk

`func (o *UpdateUserRequest) GetRevertDisplayNameOk() (*bool, bool)`

GetRevertDisplayNameOk returns a tuple with the RevertDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevertDisplayName

`func (o *UpdateUserRequest) SetRevertDisplayName(v bool)`

SetRevertDisplayName sets RevertDisplayName field to given value.

### HasRevertDisplayName

`func (o *UpdateUserRequest) HasRevertDisplayName() bool`

HasRevertDisplayName returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateUserRequest) GetStatus() UserStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateUserRequest) GetStatusOk() (*UserStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateUserRequest) SetStatus(v UserStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateUserRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDescription

`func (o *UpdateUserRequest) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *UpdateUserRequest) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *UpdateUserRequest) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.

### HasStatusDescription

`func (o *UpdateUserRequest) HasStatusDescription() bool`

HasStatusDescription returns a boolean if a field has been set.

### GetTags

`func (o *UpdateUserRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateUserRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateUserRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateUserRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUnsubscribe

`func (o *UpdateUserRequest) GetUnsubscribe() bool`

GetUnsubscribe returns the Unsubscribe field if non-nil, zero value otherwise.

### GetUnsubscribeOk

`func (o *UpdateUserRequest) GetUnsubscribeOk() (*bool, bool)`

GetUnsubscribeOk returns a tuple with the Unsubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribe

`func (o *UpdateUserRequest) SetUnsubscribe(v bool)`

SetUnsubscribe sets Unsubscribe field to given value.

### HasUnsubscribe

`func (o *UpdateUserRequest) HasUnsubscribe() bool`

HasUnsubscribe returns a boolean if a field has been set.

### GetUserIcon

`func (o *UpdateUserRequest) GetUserIcon() string`

GetUserIcon returns the UserIcon field if non-nil, zero value otherwise.

### GetUserIconOk

`func (o *UpdateUserRequest) GetUserIconOk() (*string, bool)`

GetUserIconOk returns a tuple with the UserIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIcon

`func (o *UpdateUserRequest) SetUserIcon(v string)`

SetUserIcon sets UserIcon field to given value.

### HasUserIcon

`func (o *UpdateUserRequest) HasUserIcon() bool`

HasUserIcon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


