# UpdateCalendarEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** |  | [optional] 
**CloseInstanceAfterEndMinutes** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EndsAt** | Pointer to **time.Time** | Time the vent starts at | [optional] 
**Featured** | Pointer to **bool** |  | [optional] 
**GuestEarlyJoinMinutes** | Pointer to **int32** |  | [optional] 
**HostEarlyJoinMinutes** | Pointer to **int32** |  | [optional] 
**ImageId** | Pointer to **string** |  | [optional] 
**IsDraft** | Pointer to **bool** |  | [optional] 
**Languages** | Pointer to **[]string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**Platforms** | Pointer to **[]string** |  | [optional] 
**RoleIds** | Pointer to **[]string** |  | [optional] 
**SendCreationNotification** | Pointer to **bool** | Send notification to group members. | [optional] [default to false]
**StartsAt** | Pointer to **time.Time** | Time the vent starts at | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Title** | Pointer to **string** | Event title | [optional] 
**UsesInstanceOverflow** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateCalendarEventRequest

`func NewUpdateCalendarEventRequest() *UpdateCalendarEventRequest`

NewUpdateCalendarEventRequest instantiates a new UpdateCalendarEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCalendarEventRequestWithDefaults

`func NewUpdateCalendarEventRequestWithDefaults() *UpdateCalendarEventRequest`

NewUpdateCalendarEventRequestWithDefaults instantiates a new UpdateCalendarEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *UpdateCalendarEventRequest) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *UpdateCalendarEventRequest) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *UpdateCalendarEventRequest) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *UpdateCalendarEventRequest) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCloseInstanceAfterEndMinutes

`func (o *UpdateCalendarEventRequest) GetCloseInstanceAfterEndMinutes() int32`

GetCloseInstanceAfterEndMinutes returns the CloseInstanceAfterEndMinutes field if non-nil, zero value otherwise.

### GetCloseInstanceAfterEndMinutesOk

`func (o *UpdateCalendarEventRequest) GetCloseInstanceAfterEndMinutesOk() (*int32, bool)`

GetCloseInstanceAfterEndMinutesOk returns a tuple with the CloseInstanceAfterEndMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseInstanceAfterEndMinutes

`func (o *UpdateCalendarEventRequest) SetCloseInstanceAfterEndMinutes(v int32)`

SetCloseInstanceAfterEndMinutes sets CloseInstanceAfterEndMinutes field to given value.

### HasCloseInstanceAfterEndMinutes

`func (o *UpdateCalendarEventRequest) HasCloseInstanceAfterEndMinutes() bool`

HasCloseInstanceAfterEndMinutes returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateCalendarEventRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCalendarEventRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCalendarEventRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCalendarEventRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndsAt

`func (o *UpdateCalendarEventRequest) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *UpdateCalendarEventRequest) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *UpdateCalendarEventRequest) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.

### HasEndsAt

`func (o *UpdateCalendarEventRequest) HasEndsAt() bool`

HasEndsAt returns a boolean if a field has been set.

### GetFeatured

`func (o *UpdateCalendarEventRequest) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *UpdateCalendarEventRequest) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *UpdateCalendarEventRequest) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *UpdateCalendarEventRequest) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetGuestEarlyJoinMinutes

`func (o *UpdateCalendarEventRequest) GetGuestEarlyJoinMinutes() int32`

GetGuestEarlyJoinMinutes returns the GuestEarlyJoinMinutes field if non-nil, zero value otherwise.

### GetGuestEarlyJoinMinutesOk

`func (o *UpdateCalendarEventRequest) GetGuestEarlyJoinMinutesOk() (*int32, bool)`

GetGuestEarlyJoinMinutesOk returns a tuple with the GuestEarlyJoinMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestEarlyJoinMinutes

`func (o *UpdateCalendarEventRequest) SetGuestEarlyJoinMinutes(v int32)`

SetGuestEarlyJoinMinutes sets GuestEarlyJoinMinutes field to given value.

### HasGuestEarlyJoinMinutes

`func (o *UpdateCalendarEventRequest) HasGuestEarlyJoinMinutes() bool`

HasGuestEarlyJoinMinutes returns a boolean if a field has been set.

### GetHostEarlyJoinMinutes

`func (o *UpdateCalendarEventRequest) GetHostEarlyJoinMinutes() int32`

GetHostEarlyJoinMinutes returns the HostEarlyJoinMinutes field if non-nil, zero value otherwise.

### GetHostEarlyJoinMinutesOk

`func (o *UpdateCalendarEventRequest) GetHostEarlyJoinMinutesOk() (*int32, bool)`

GetHostEarlyJoinMinutesOk returns a tuple with the HostEarlyJoinMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostEarlyJoinMinutes

`func (o *UpdateCalendarEventRequest) SetHostEarlyJoinMinutes(v int32)`

SetHostEarlyJoinMinutes sets HostEarlyJoinMinutes field to given value.

### HasHostEarlyJoinMinutes

`func (o *UpdateCalendarEventRequest) HasHostEarlyJoinMinutes() bool`

HasHostEarlyJoinMinutes returns a boolean if a field has been set.

### GetImageId

`func (o *UpdateCalendarEventRequest) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *UpdateCalendarEventRequest) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *UpdateCalendarEventRequest) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *UpdateCalendarEventRequest) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetIsDraft

`func (o *UpdateCalendarEventRequest) GetIsDraft() bool`

GetIsDraft returns the IsDraft field if non-nil, zero value otherwise.

### GetIsDraftOk

`func (o *UpdateCalendarEventRequest) GetIsDraftOk() (*bool, bool)`

GetIsDraftOk returns a tuple with the IsDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDraft

`func (o *UpdateCalendarEventRequest) SetIsDraft(v bool)`

SetIsDraft sets IsDraft field to given value.

### HasIsDraft

`func (o *UpdateCalendarEventRequest) HasIsDraft() bool`

HasIsDraft returns a boolean if a field has been set.

### GetLanguages

`func (o *UpdateCalendarEventRequest) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *UpdateCalendarEventRequest) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *UpdateCalendarEventRequest) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *UpdateCalendarEventRequest) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetParentId

`func (o *UpdateCalendarEventRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *UpdateCalendarEventRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *UpdateCalendarEventRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *UpdateCalendarEventRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetPlatforms

`func (o *UpdateCalendarEventRequest) GetPlatforms() []string`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *UpdateCalendarEventRequest) GetPlatformsOk() (*[]string, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *UpdateCalendarEventRequest) SetPlatforms(v []string)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *UpdateCalendarEventRequest) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetRoleIds

`func (o *UpdateCalendarEventRequest) GetRoleIds() []string`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *UpdateCalendarEventRequest) GetRoleIdsOk() (*[]string, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *UpdateCalendarEventRequest) SetRoleIds(v []string)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *UpdateCalendarEventRequest) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.

### GetSendCreationNotification

`func (o *UpdateCalendarEventRequest) GetSendCreationNotification() bool`

GetSendCreationNotification returns the SendCreationNotification field if non-nil, zero value otherwise.

### GetSendCreationNotificationOk

`func (o *UpdateCalendarEventRequest) GetSendCreationNotificationOk() (*bool, bool)`

GetSendCreationNotificationOk returns a tuple with the SendCreationNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendCreationNotification

`func (o *UpdateCalendarEventRequest) SetSendCreationNotification(v bool)`

SetSendCreationNotification sets SendCreationNotification field to given value.

### HasSendCreationNotification

`func (o *UpdateCalendarEventRequest) HasSendCreationNotification() bool`

HasSendCreationNotification returns a boolean if a field has been set.

### GetStartsAt

`func (o *UpdateCalendarEventRequest) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *UpdateCalendarEventRequest) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *UpdateCalendarEventRequest) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *UpdateCalendarEventRequest) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetTags

`func (o *UpdateCalendarEventRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateCalendarEventRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateCalendarEventRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateCalendarEventRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTitle

`func (o *UpdateCalendarEventRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateCalendarEventRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateCalendarEventRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateCalendarEventRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUsesInstanceOverflow

`func (o *UpdateCalendarEventRequest) GetUsesInstanceOverflow() bool`

GetUsesInstanceOverflow returns the UsesInstanceOverflow field if non-nil, zero value otherwise.

### GetUsesInstanceOverflowOk

`func (o *UpdateCalendarEventRequest) GetUsesInstanceOverflowOk() (*bool, bool)`

GetUsesInstanceOverflowOk returns a tuple with the UsesInstanceOverflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesInstanceOverflow

`func (o *UpdateCalendarEventRequest) SetUsesInstanceOverflow(v bool)`

SetUsesInstanceOverflow sets UsesInstanceOverflow field to given value.

### HasUsesInstanceOverflow

`func (o *UpdateCalendarEventRequest) HasUsesInstanceOverflow() bool`

HasUsesInstanceOverflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


