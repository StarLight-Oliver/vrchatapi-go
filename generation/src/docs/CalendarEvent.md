# CalendarEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | [**CalendarEventAccess**](CalendarEventAccess.md) |  | [default to PUBLIC]
**Category** | [**CalendarEventCategory**](CalendarEventCategory.md) |  | [default to OTHER]
**CloseInstanceAfterEndMinutes** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **NullableTime** |  | [optional] 
**Description** | **string** |  | 
**DurationInMs** | Pointer to **int32** |  | [optional] 
**EndsAt** | **time.Time** |  | 
**Featured** | Pointer to **bool** |  | [optional] 
**GuestEarlyJoinMinutes** | Pointer to **int32** |  | [optional] 
**HostEarlyJoinMinutes** | Pointer to **int32** |  | [optional] 
**Id** | **string** |  | 
**ImageId** | Pointer to **string** |  | [optional] 
**ImageUrl** | Pointer to **NullableString** |  | [optional] 
**InterestedUserCount** | Pointer to **int32** |  | [optional] 
**IsDraft** | Pointer to **bool** |  | [optional] 
**Languages** | Pointer to **[]string** | Languages that might be spoken at this event | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**Platforms** | Pointer to [**[]CalendarEventPlatform**](CalendarEventPlatform.md) |  | [optional] 
**RoleIds** | Pointer to **[]string** | Group roles that may join this event | [optional] 
**StartsAt** | **time.Time** |  | 
**Tags** | Pointer to **[]string** | Custom tags for this event | [optional] 
**Title** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UserInterest** | Pointer to [**CalendarEventUserInterest**](CalendarEventUserInterest.md) |  | [optional] 
**UsesInstanceOverflow** | Pointer to **bool** |  | [optional] 

## Methods

### NewCalendarEvent

`func NewCalendarEvent(accessType CalendarEventAccess, category CalendarEventCategory, description string, endsAt time.Time, id string, startsAt time.Time, title string, ) *CalendarEvent`

NewCalendarEvent instantiates a new CalendarEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalendarEventWithDefaults

`func NewCalendarEventWithDefaults() *CalendarEvent`

NewCalendarEventWithDefaults instantiates a new CalendarEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *CalendarEvent) GetAccessType() CalendarEventAccess`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *CalendarEvent) GetAccessTypeOk() (*CalendarEventAccess, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *CalendarEvent) SetAccessType(v CalendarEventAccess)`

SetAccessType sets AccessType field to given value.


### GetCategory

`func (o *CalendarEvent) GetCategory() CalendarEventCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CalendarEvent) GetCategoryOk() (*CalendarEventCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CalendarEvent) SetCategory(v CalendarEventCategory)`

SetCategory sets Category field to given value.


### GetCloseInstanceAfterEndMinutes

`func (o *CalendarEvent) GetCloseInstanceAfterEndMinutes() int32`

GetCloseInstanceAfterEndMinutes returns the CloseInstanceAfterEndMinutes field if non-nil, zero value otherwise.

### GetCloseInstanceAfterEndMinutesOk

`func (o *CalendarEvent) GetCloseInstanceAfterEndMinutesOk() (*int32, bool)`

GetCloseInstanceAfterEndMinutesOk returns a tuple with the CloseInstanceAfterEndMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseInstanceAfterEndMinutes

`func (o *CalendarEvent) SetCloseInstanceAfterEndMinutes(v int32)`

SetCloseInstanceAfterEndMinutes sets CloseInstanceAfterEndMinutes field to given value.

### HasCloseInstanceAfterEndMinutes

`func (o *CalendarEvent) HasCloseInstanceAfterEndMinutes() bool`

HasCloseInstanceAfterEndMinutes returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CalendarEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CalendarEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CalendarEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CalendarEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *CalendarEvent) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *CalendarEvent) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *CalendarEvent) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *CalendarEvent) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *CalendarEvent) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *CalendarEvent) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetDescription

`func (o *CalendarEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CalendarEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CalendarEvent) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDurationInMs

`func (o *CalendarEvent) GetDurationInMs() int32`

GetDurationInMs returns the DurationInMs field if non-nil, zero value otherwise.

### GetDurationInMsOk

`func (o *CalendarEvent) GetDurationInMsOk() (*int32, bool)`

GetDurationInMsOk returns a tuple with the DurationInMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInMs

`func (o *CalendarEvent) SetDurationInMs(v int32)`

SetDurationInMs sets DurationInMs field to given value.

### HasDurationInMs

`func (o *CalendarEvent) HasDurationInMs() bool`

HasDurationInMs returns a boolean if a field has been set.

### GetEndsAt

`func (o *CalendarEvent) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *CalendarEvent) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *CalendarEvent) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.


### GetFeatured

`func (o *CalendarEvent) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *CalendarEvent) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *CalendarEvent) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *CalendarEvent) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetGuestEarlyJoinMinutes

`func (o *CalendarEvent) GetGuestEarlyJoinMinutes() int32`

GetGuestEarlyJoinMinutes returns the GuestEarlyJoinMinutes field if non-nil, zero value otherwise.

### GetGuestEarlyJoinMinutesOk

`func (o *CalendarEvent) GetGuestEarlyJoinMinutesOk() (*int32, bool)`

GetGuestEarlyJoinMinutesOk returns a tuple with the GuestEarlyJoinMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestEarlyJoinMinutes

`func (o *CalendarEvent) SetGuestEarlyJoinMinutes(v int32)`

SetGuestEarlyJoinMinutes sets GuestEarlyJoinMinutes field to given value.

### HasGuestEarlyJoinMinutes

`func (o *CalendarEvent) HasGuestEarlyJoinMinutes() bool`

HasGuestEarlyJoinMinutes returns a boolean if a field has been set.

### GetHostEarlyJoinMinutes

`func (o *CalendarEvent) GetHostEarlyJoinMinutes() int32`

GetHostEarlyJoinMinutes returns the HostEarlyJoinMinutes field if non-nil, zero value otherwise.

### GetHostEarlyJoinMinutesOk

`func (o *CalendarEvent) GetHostEarlyJoinMinutesOk() (*int32, bool)`

GetHostEarlyJoinMinutesOk returns a tuple with the HostEarlyJoinMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostEarlyJoinMinutes

`func (o *CalendarEvent) SetHostEarlyJoinMinutes(v int32)`

SetHostEarlyJoinMinutes sets HostEarlyJoinMinutes field to given value.

### HasHostEarlyJoinMinutes

`func (o *CalendarEvent) HasHostEarlyJoinMinutes() bool`

HasHostEarlyJoinMinutes returns a boolean if a field has been set.

### GetId

`func (o *CalendarEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CalendarEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CalendarEvent) SetId(v string)`

SetId sets Id field to given value.


### GetImageId

`func (o *CalendarEvent) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *CalendarEvent) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *CalendarEvent) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *CalendarEvent) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetImageUrl

`func (o *CalendarEvent) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *CalendarEvent) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *CalendarEvent) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *CalendarEvent) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *CalendarEvent) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *CalendarEvent) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetInterestedUserCount

`func (o *CalendarEvent) GetInterestedUserCount() int32`

GetInterestedUserCount returns the InterestedUserCount field if non-nil, zero value otherwise.

### GetInterestedUserCountOk

`func (o *CalendarEvent) GetInterestedUserCountOk() (*int32, bool)`

GetInterestedUserCountOk returns a tuple with the InterestedUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestedUserCount

`func (o *CalendarEvent) SetInterestedUserCount(v int32)`

SetInterestedUserCount sets InterestedUserCount field to given value.

### HasInterestedUserCount

`func (o *CalendarEvent) HasInterestedUserCount() bool`

HasInterestedUserCount returns a boolean if a field has been set.

### GetIsDraft

`func (o *CalendarEvent) GetIsDraft() bool`

GetIsDraft returns the IsDraft field if non-nil, zero value otherwise.

### GetIsDraftOk

`func (o *CalendarEvent) GetIsDraftOk() (*bool, bool)`

GetIsDraftOk returns a tuple with the IsDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDraft

`func (o *CalendarEvent) SetIsDraft(v bool)`

SetIsDraft sets IsDraft field to given value.

### HasIsDraft

`func (o *CalendarEvent) HasIsDraft() bool`

HasIsDraft returns a boolean if a field has been set.

### GetLanguages

`func (o *CalendarEvent) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *CalendarEvent) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *CalendarEvent) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *CalendarEvent) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetOwnerId

`func (o *CalendarEvent) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *CalendarEvent) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *CalendarEvent) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *CalendarEvent) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetPlatforms

`func (o *CalendarEvent) GetPlatforms() []CalendarEventPlatform`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *CalendarEvent) GetPlatformsOk() (*[]CalendarEventPlatform, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *CalendarEvent) SetPlatforms(v []CalendarEventPlatform)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *CalendarEvent) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetRoleIds

`func (o *CalendarEvent) GetRoleIds() []string`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *CalendarEvent) GetRoleIdsOk() (*[]string, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *CalendarEvent) SetRoleIds(v []string)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *CalendarEvent) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.

### SetRoleIdsNil

`func (o *CalendarEvent) SetRoleIdsNil(b bool)`

 SetRoleIdsNil sets the value for RoleIds to be an explicit nil

### UnsetRoleIds
`func (o *CalendarEvent) UnsetRoleIds()`

UnsetRoleIds ensures that no value is present for RoleIds, not even an explicit nil
### GetStartsAt

`func (o *CalendarEvent) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *CalendarEvent) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *CalendarEvent) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.


### GetTags

`func (o *CalendarEvent) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CalendarEvent) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CalendarEvent) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CalendarEvent) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTitle

`func (o *CalendarEvent) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CalendarEvent) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CalendarEvent) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *CalendarEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CalendarEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CalendarEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CalendarEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CalendarEvent) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CalendarEvent) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CalendarEvent) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CalendarEvent) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserInterest

`func (o *CalendarEvent) GetUserInterest() CalendarEventUserInterest`

GetUserInterest returns the UserInterest field if non-nil, zero value otherwise.

### GetUserInterestOk

`func (o *CalendarEvent) GetUserInterestOk() (*CalendarEventUserInterest, bool)`

GetUserInterestOk returns a tuple with the UserInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInterest

`func (o *CalendarEvent) SetUserInterest(v CalendarEventUserInterest)`

SetUserInterest sets UserInterest field to given value.

### HasUserInterest

`func (o *CalendarEvent) HasUserInterest() bool`

HasUserInterest returns a boolean if a field has been set.

### GetUsesInstanceOverflow

`func (o *CalendarEvent) GetUsesInstanceOverflow() bool`

GetUsesInstanceOverflow returns the UsesInstanceOverflow field if non-nil, zero value otherwise.

### GetUsesInstanceOverflowOk

`func (o *CalendarEvent) GetUsesInstanceOverflowOk() (*bool, bool)`

GetUsesInstanceOverflowOk returns a tuple with the UsesInstanceOverflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesInstanceOverflow

`func (o *CalendarEvent) SetUsesInstanceOverflow(v bool)`

SetUsesInstanceOverflow sets UsesInstanceOverflow field to given value.

### HasUsesInstanceOverflow

`func (o *CalendarEvent) HasUsesInstanceOverflow() bool`

HasUsesInstanceOverflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


