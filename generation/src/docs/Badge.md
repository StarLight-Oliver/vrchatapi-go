# Badge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedAt** | Pointer to **NullableTime** | only present in CurrentUser badges | [optional] 
**BadgeDescription** | **string** |  | 
**BadgeId** | **string** |  | 
**BadgeImageUrl** | **string** | direct url to image | 
**BadgeName** | **string** |  | 
**Hidden** | Pointer to **NullableBool** | only present in CurrentUser badges | [optional] 
**Showcased** | **bool** |  | 
**UpdatedAt** | Pointer to **NullableTime** | only present in CurrentUser badges | [optional] 

## Methods

### NewBadge

`func NewBadge(badgeDescription string, badgeId string, badgeImageUrl string, badgeName string, showcased bool, ) *Badge`

NewBadge instantiates a new Badge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBadgeWithDefaults

`func NewBadgeWithDefaults() *Badge`

NewBadgeWithDefaults instantiates a new Badge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedAt

`func (o *Badge) GetAssignedAt() time.Time`

GetAssignedAt returns the AssignedAt field if non-nil, zero value otherwise.

### GetAssignedAtOk

`func (o *Badge) GetAssignedAtOk() (*time.Time, bool)`

GetAssignedAtOk returns a tuple with the AssignedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedAt

`func (o *Badge) SetAssignedAt(v time.Time)`

SetAssignedAt sets AssignedAt field to given value.

### HasAssignedAt

`func (o *Badge) HasAssignedAt() bool`

HasAssignedAt returns a boolean if a field has been set.

### SetAssignedAtNil

`func (o *Badge) SetAssignedAtNil(b bool)`

 SetAssignedAtNil sets the value for AssignedAt to be an explicit nil

### UnsetAssignedAt
`func (o *Badge) UnsetAssignedAt()`

UnsetAssignedAt ensures that no value is present for AssignedAt, not even an explicit nil
### GetBadgeDescription

`func (o *Badge) GetBadgeDescription() string`

GetBadgeDescription returns the BadgeDescription field if non-nil, zero value otherwise.

### GetBadgeDescriptionOk

`func (o *Badge) GetBadgeDescriptionOk() (*string, bool)`

GetBadgeDescriptionOk returns a tuple with the BadgeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadgeDescription

`func (o *Badge) SetBadgeDescription(v string)`

SetBadgeDescription sets BadgeDescription field to given value.


### GetBadgeId

`func (o *Badge) GetBadgeId() string`

GetBadgeId returns the BadgeId field if non-nil, zero value otherwise.

### GetBadgeIdOk

`func (o *Badge) GetBadgeIdOk() (*string, bool)`

GetBadgeIdOk returns a tuple with the BadgeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadgeId

`func (o *Badge) SetBadgeId(v string)`

SetBadgeId sets BadgeId field to given value.


### GetBadgeImageUrl

`func (o *Badge) GetBadgeImageUrl() string`

GetBadgeImageUrl returns the BadgeImageUrl field if non-nil, zero value otherwise.

### GetBadgeImageUrlOk

`func (o *Badge) GetBadgeImageUrlOk() (*string, bool)`

GetBadgeImageUrlOk returns a tuple with the BadgeImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadgeImageUrl

`func (o *Badge) SetBadgeImageUrl(v string)`

SetBadgeImageUrl sets BadgeImageUrl field to given value.


### GetBadgeName

`func (o *Badge) GetBadgeName() string`

GetBadgeName returns the BadgeName field if non-nil, zero value otherwise.

### GetBadgeNameOk

`func (o *Badge) GetBadgeNameOk() (*string, bool)`

GetBadgeNameOk returns a tuple with the BadgeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadgeName

`func (o *Badge) SetBadgeName(v string)`

SetBadgeName sets BadgeName field to given value.


### GetHidden

`func (o *Badge) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *Badge) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *Badge) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *Badge) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### SetHiddenNil

`func (o *Badge) SetHiddenNil(b bool)`

 SetHiddenNil sets the value for Hidden to be an explicit nil

### UnsetHidden
`func (o *Badge) UnsetHidden()`

UnsetHidden ensures that no value is present for Hidden, not even an explicit nil
### GetShowcased

`func (o *Badge) GetShowcased() bool`

GetShowcased returns the Showcased field if non-nil, zero value otherwise.

### GetShowcasedOk

`func (o *Badge) GetShowcasedOk() (*bool, bool)`

GetShowcasedOk returns a tuple with the Showcased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowcased

`func (o *Badge) SetShowcased(v bool)`

SetShowcased sets Showcased field to given value.


### GetUpdatedAt

`func (o *Badge) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Badge) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Badge) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Badge) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *Badge) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Badge) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


