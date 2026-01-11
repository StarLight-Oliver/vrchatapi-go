# PermissionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Badges** | Pointer to **[]string** | Badges afforded the user by this permission | [optional] 
**Max** | Pointer to **int32** | Maximum value afforded the user by this permission | [optional] 
**MaxFavoriteGroups** | Pointer to **map[string]int32** | Maximum favorite groups afforded the user by this permission | [optional] 
**MaxFavoritesPerGroup** | Pointer to **map[string]int32** | Maximum favorites per group afforded the user by this permission | [optional] 
**Tags** | Pointer to **[]string** | Tags afforded the user by this permission | [optional] 

## Methods

### NewPermissionData

`func NewPermissionData() *PermissionData`

NewPermissionData instantiates a new PermissionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionDataWithDefaults

`func NewPermissionDataWithDefaults() *PermissionData`

NewPermissionDataWithDefaults instantiates a new PermissionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBadges

`func (o *PermissionData) GetBadges() []string`

GetBadges returns the Badges field if non-nil, zero value otherwise.

### GetBadgesOk

`func (o *PermissionData) GetBadgesOk() (*[]string, bool)`

GetBadgesOk returns a tuple with the Badges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadges

`func (o *PermissionData) SetBadges(v []string)`

SetBadges sets Badges field to given value.

### HasBadges

`func (o *PermissionData) HasBadges() bool`

HasBadges returns a boolean if a field has been set.

### GetMax

`func (o *PermissionData) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *PermissionData) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *PermissionData) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *PermissionData) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMaxFavoriteGroups

`func (o *PermissionData) GetMaxFavoriteGroups() map[string]int32`

GetMaxFavoriteGroups returns the MaxFavoriteGroups field if non-nil, zero value otherwise.

### GetMaxFavoriteGroupsOk

`func (o *PermissionData) GetMaxFavoriteGroupsOk() (*map[string]int32, bool)`

GetMaxFavoriteGroupsOk returns a tuple with the MaxFavoriteGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFavoriteGroups

`func (o *PermissionData) SetMaxFavoriteGroups(v map[string]int32)`

SetMaxFavoriteGroups sets MaxFavoriteGroups field to given value.

### HasMaxFavoriteGroups

`func (o *PermissionData) HasMaxFavoriteGroups() bool`

HasMaxFavoriteGroups returns a boolean if a field has been set.

### GetMaxFavoritesPerGroup

`func (o *PermissionData) GetMaxFavoritesPerGroup() map[string]int32`

GetMaxFavoritesPerGroup returns the MaxFavoritesPerGroup field if non-nil, zero value otherwise.

### GetMaxFavoritesPerGroupOk

`func (o *PermissionData) GetMaxFavoritesPerGroupOk() (*map[string]int32, bool)`

GetMaxFavoritesPerGroupOk returns a tuple with the MaxFavoritesPerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFavoritesPerGroup

`func (o *PermissionData) SetMaxFavoritesPerGroup(v map[string]int32)`

SetMaxFavoritesPerGroup sets MaxFavoritesPerGroup field to given value.

### HasMaxFavoritesPerGroup

`func (o *PermissionData) HasMaxFavoritesPerGroup() bool`

HasMaxFavoritesPerGroup returns a boolean if a field has been set.

### GetTags

`func (o *PermissionData) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PermissionData) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PermissionData) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PermissionData) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


