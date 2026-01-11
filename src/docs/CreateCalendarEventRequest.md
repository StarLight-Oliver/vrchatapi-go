# CreateCalendarEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | [**CalendarEventAccess**](CalendarEventAccess.md) |  | [default to PUBLIC]
**Category** | [**CalendarEventCategory**](CalendarEventCategory.md) |  | [default to OTHER]
**CloseInstanceAfterEndMinutes** | Pointer to **int32** |  | [optional] 
**Description** | **string** |  | 
**EndsAt** | **time.Time** | Time the event ends at | 
**Featured** | Pointer to **bool** |  | [optional] 
**GuestEarlyJoinMinutes** | Pointer to **int32** |  | [optional] 
**HostEarlyJoinMinutes** | Pointer to **int32** |  | [optional] 
**ImageId** | Pointer to **string** |  | [optional] 
**IsDraft** | Pointer to **bool** |  | [optional] 
**Languages** | Pointer to **[]string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**Platforms** | Pointer to [**[]CalendarEventPlatform**](CalendarEventPlatform.md) |  | [optional] 
**RoleIds** | Pointer to **[]string** |  | [optional] 
**SendCreationNotification** | **bool** | Send notification to group members. | 
**StartsAt** | **time.Time** | Time the event starts at | 
**Tags** | Pointer to **[]string** |  | [optional] 
**Title** | **string** | Event title | 
**UsesInstanceOverflow** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateCalendarEventRequest

`func NewCreateCalendarEventRequest(accessType CalendarEventAccess, category CalendarEventCategory, description string, endsAt time.Time, sendCreationNotification bool, startsAt time.Time, title string, ) *CreateCalendarEventRequest`

NewCreateCalendarEventRequest instantiates a new CreateCalendarEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCalendarEventRequestWithDefaults

`func NewCreateCalendarEventRequestWithDefaults() *CreateCalendarEventRequest`

NewCreateCalendarEventRequestWithDefaults instantiates a new CreateCalendarEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *CreateCalendarEventRequest) GetAccessType() CalendarEventAccess`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *CreateCalendarEventRequest) GetAccessTypeOk() (*CalendarEventAccess, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *CreateCalendarEventRequest) SetAccessType(v CalendarEventAccess)`

SetAccessType sets AccessType field to given value.


### GetCategory

`func (o *CreateCalendarEventRequest) GetCategory() CalendarEventCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CreateCalendarEventRequest) GetCategoryOk() (*CalendarEventCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CreateCalendarEventRequest) SetCategory(v CalendarEventCategory)`

SetCategory sets Category field to given value.


### GetCloseInstanceAfterEndMinutes

`func (o *CreateCalendarEventRequest) GetCloseInstanceAfterEndMinutes() int32`

GetCloseInstanceAfterEndMinutes returns the CloseInstanceAfterEndMinutes field if non-nil, zero value otherwise.

### GetCloseInstanceAfterEndMinutesOk

`func (o *CreateCalendarEventRequest) GetCloseInstanceAfterEndMinutesOk() (*int32, bool)`

GetCloseInstanceAfterEndMinutesOk returns a tuple with the CloseInstanceAfterEndMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseInstanceAfterEndMinutes

`func (o *CreateCalendarEventRequest) SetCloseInstanceAfterEndMinutes(v int32)`

SetCloseInstanceAfterEndMinutes sets CloseInstanceAfterEndMinutes field to given value.

### HasCloseInstanceAfterEndMinutes

`func (o *CreateCalendarEventRequest) HasCloseInstanceAfterEndMinutes() bool`

HasCloseInstanceAfterEndMinutes returns a boolean if a field has been set.

### GetDescription

`func (o *CreateCalendarEventRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCalendarEventRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCalendarEventRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEndsAt

`func (o *CreateCalendarEventRequest) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *CreateCalendarEventRequest) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *CreateCalendarEventRequest) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.


### GetFeatured

`func (o *CreateCalendarEventRequest) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *CreateCalendarEventRequest) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *CreateCalendarEventRequest) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *CreateCalendarEventRequest) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetGuestEarlyJoinMinutes

`func (o *CreateCalendarEventRequest) GetGuestEarlyJoinMinutes() int32`

GetGuestEarlyJoinMinutes returns the GuestEarlyJoinMinutes field if non-nil, zero value otherwise.

### GetGuestEarlyJoinMinutesOk

`func (o *CreateCalendarEventRequest) GetGuestEarlyJoinMinutesOk() (*int32, bool)`

GetGuestEarlyJoinMinutesOk returns a tuple with the GuestEarlyJoinMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestEarlyJoinMinutes

`func (o *CreateCalendarEventRequest) SetGuestEarlyJoinMinutes(v int32)`

SetGuestEarlyJoinMinutes sets GuestEarlyJoinMinutes field to given value.

### HasGuestEarlyJoinMinutes

`func (o *CreateCalendarEventRequest) HasGuestEarlyJoinMinutes() bool`

HasGuestEarlyJoinMinutes returns a boolean if a field has been set.

### GetHostEarlyJoinMinutes

`func (o *CreateCalendarEventRequest) GetHostEarlyJoinMinutes() int32`

GetHostEarlyJoinMinutes returns the HostEarlyJoinMinutes field if non-nil, zero value otherwise.

### GetHostEarlyJoinMinutesOk

`func (o *CreateCalendarEventRequest) GetHostEarlyJoinMinutesOk() (*int32, bool)`

GetHostEarlyJoinMinutesOk returns a tuple with the HostEarlyJoinMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostEarlyJoinMinutes

`func (o *CreateCalendarEventRequest) SetHostEarlyJoinMinutes(v int32)`

SetHostEarlyJoinMinutes sets HostEarlyJoinMinutes field to given value.

### HasHostEarlyJoinMinutes

`func (o *CreateCalendarEventRequest) HasHostEarlyJoinMinutes() bool`

HasHostEarlyJoinMinutes returns a boolean if a field has been set.

### GetImageId

`func (o *CreateCalendarEventRequest) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *CreateCalendarEventRequest) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *CreateCalendarEventRequest) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *CreateCalendarEventRequest) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetIsDraft

`func (o *CreateCalendarEventRequest) GetIsDraft() bool`

GetIsDraft returns the IsDraft field if non-nil, zero value otherwise.

### GetIsDraftOk

`func (o *CreateCalendarEventRequest) GetIsDraftOk() (*bool, bool)`

GetIsDraftOk returns a tuple with the IsDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDraft

`func (o *CreateCalendarEventRequest) SetIsDraft(v bool)`

SetIsDraft sets IsDraft field to given value.

### HasIsDraft

`func (o *CreateCalendarEventRequest) HasIsDraft() bool`

HasIsDraft returns a boolean if a field has been set.

### GetLanguages

`func (o *CreateCalendarEventRequest) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *CreateCalendarEventRequest) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *CreateCalendarEventRequest) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *CreateCalendarEventRequest) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetParentId

`func (o *CreateCalendarEventRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreateCalendarEventRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreateCalendarEventRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CreateCalendarEventRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetPlatforms

`func (o *CreateCalendarEventRequest) GetPlatforms() []CalendarEventPlatform`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *CreateCalendarEventRequest) GetPlatformsOk() (*[]CalendarEventPlatform, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *CreateCalendarEventRequest) SetPlatforms(v []CalendarEventPlatform)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *CreateCalendarEventRequest) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetRoleIds

`func (o *CreateCalendarEventRequest) GetRoleIds() []string`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *CreateCalendarEventRequest) GetRoleIdsOk() (*[]string, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *CreateCalendarEventRequest) SetRoleIds(v []string)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *CreateCalendarEventRequest) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.

### GetSendCreationNotification

`func (o *CreateCalendarEventRequest) GetSendCreationNotification() bool`

GetSendCreationNotification returns the SendCreationNotification field if non-nil, zero value otherwise.

### GetSendCreationNotificationOk

`func (o *CreateCalendarEventRequest) GetSendCreationNotificationOk() (*bool, bool)`

GetSendCreationNotificationOk returns a tuple with the SendCreationNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendCreationNotification

`func (o *CreateCalendarEventRequest) SetSendCreationNotification(v bool)`

SetSendCreationNotification sets SendCreationNotification field to given value.


### GetStartsAt

`func (o *CreateCalendarEventRequest) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *CreateCalendarEventRequest) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *CreateCalendarEventRequest) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.


### GetTags

`func (o *CreateCalendarEventRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateCalendarEventRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateCalendarEventRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateCalendarEventRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTitle

`func (o *CreateCalendarEventRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateCalendarEventRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateCalendarEventRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUsesInstanceOverflow

`func (o *CreateCalendarEventRequest) GetUsesInstanceOverflow() bool`

GetUsesInstanceOverflow returns the UsesInstanceOverflow field if non-nil, zero value otherwise.

### GetUsesInstanceOverflowOk

`func (o *CreateCalendarEventRequest) GetUsesInstanceOverflowOk() (*bool, bool)`

GetUsesInstanceOverflowOk returns a tuple with the UsesInstanceOverflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesInstanceOverflow

`func (o *CreateCalendarEventRequest) SetUsesInstanceOverflow(v bool)`

SetUsesInstanceOverflow sets UsesInstanceOverflow field to given value.

### HasUsesInstanceOverflow

`func (o *CreateCalendarEventRequest) HasUsesInstanceOverflow() bool`

HasUsesInstanceOverflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


