# MutualFriend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarThumbnail** | Pointer to **string** | When profilePicOverride is not empty, use it instead. | [optional] 
**CurrentAvatarImageUrl** | **string** | When profilePicOverride is not empty, use it instead. | 
**CurrentAvatarTags** | Pointer to **[]string** |  | [optional] 
**CurrentAvatarThumbnailImageUrl** | Pointer to **string** | When profilePicOverride is not empty, use it instead. | [optional] 
**DisplayName** | **string** |  | 
**Id** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**ImageUrl** | **string** |  | 
**ProfilePicOverride** | Pointer to **string** |  | [optional] 
**Status** | [**UserStatus**](UserStatus.md) |  | [default to OFFLINE]
**StatusDescription** | **string** |  | 

## Methods

### NewMutualFriend

`func NewMutualFriend(currentAvatarImageUrl string, displayName string, id string, imageUrl string, status UserStatus, statusDescription string, ) *MutualFriend`

NewMutualFriend instantiates a new MutualFriend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutualFriendWithDefaults

`func NewMutualFriendWithDefaults() *MutualFriend`

NewMutualFriendWithDefaults instantiates a new MutualFriend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarThumbnail

`func (o *MutualFriend) GetAvatarThumbnail() string`

GetAvatarThumbnail returns the AvatarThumbnail field if non-nil, zero value otherwise.

### GetAvatarThumbnailOk

`func (o *MutualFriend) GetAvatarThumbnailOk() (*string, bool)`

GetAvatarThumbnailOk returns a tuple with the AvatarThumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarThumbnail

`func (o *MutualFriend) SetAvatarThumbnail(v string)`

SetAvatarThumbnail sets AvatarThumbnail field to given value.

### HasAvatarThumbnail

`func (o *MutualFriend) HasAvatarThumbnail() bool`

HasAvatarThumbnail returns a boolean if a field has been set.

### GetCurrentAvatarImageUrl

`func (o *MutualFriend) GetCurrentAvatarImageUrl() string`

GetCurrentAvatarImageUrl returns the CurrentAvatarImageUrl field if non-nil, zero value otherwise.

### GetCurrentAvatarImageUrlOk

`func (o *MutualFriend) GetCurrentAvatarImageUrlOk() (*string, bool)`

GetCurrentAvatarImageUrlOk returns a tuple with the CurrentAvatarImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAvatarImageUrl

`func (o *MutualFriend) SetCurrentAvatarImageUrl(v string)`

SetCurrentAvatarImageUrl sets CurrentAvatarImageUrl field to given value.


### GetCurrentAvatarTags

`func (o *MutualFriend) GetCurrentAvatarTags() []string`

GetCurrentAvatarTags returns the CurrentAvatarTags field if non-nil, zero value otherwise.

### GetCurrentAvatarTagsOk

`func (o *MutualFriend) GetCurrentAvatarTagsOk() (*[]string, bool)`

GetCurrentAvatarTagsOk returns a tuple with the CurrentAvatarTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAvatarTags

`func (o *MutualFriend) SetCurrentAvatarTags(v []string)`

SetCurrentAvatarTags sets CurrentAvatarTags field to given value.

### HasCurrentAvatarTags

`func (o *MutualFriend) HasCurrentAvatarTags() bool`

HasCurrentAvatarTags returns a boolean if a field has been set.

### GetCurrentAvatarThumbnailImageUrl

`func (o *MutualFriend) GetCurrentAvatarThumbnailImageUrl() string`

GetCurrentAvatarThumbnailImageUrl returns the CurrentAvatarThumbnailImageUrl field if non-nil, zero value otherwise.

### GetCurrentAvatarThumbnailImageUrlOk

`func (o *MutualFriend) GetCurrentAvatarThumbnailImageUrlOk() (*string, bool)`

GetCurrentAvatarThumbnailImageUrlOk returns a tuple with the CurrentAvatarThumbnailImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAvatarThumbnailImageUrl

`func (o *MutualFriend) SetCurrentAvatarThumbnailImageUrl(v string)`

SetCurrentAvatarThumbnailImageUrl sets CurrentAvatarThumbnailImageUrl field to given value.

### HasCurrentAvatarThumbnailImageUrl

`func (o *MutualFriend) HasCurrentAvatarThumbnailImageUrl() bool`

HasCurrentAvatarThumbnailImageUrl returns a boolean if a field has been set.

### GetDisplayName

`func (o *MutualFriend) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MutualFriend) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MutualFriend) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetId

`func (o *MutualFriend) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MutualFriend) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MutualFriend) SetId(v string)`

SetId sets Id field to given value.


### GetImageUrl

`func (o *MutualFriend) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *MutualFriend) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *MutualFriend) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.


### GetProfilePicOverride

`func (o *MutualFriend) GetProfilePicOverride() string`

GetProfilePicOverride returns the ProfilePicOverride field if non-nil, zero value otherwise.

### GetProfilePicOverrideOk

`func (o *MutualFriend) GetProfilePicOverrideOk() (*string, bool)`

GetProfilePicOverrideOk returns a tuple with the ProfilePicOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePicOverride

`func (o *MutualFriend) SetProfilePicOverride(v string)`

SetProfilePicOverride sets ProfilePicOverride field to given value.

### HasProfilePicOverride

`func (o *MutualFriend) HasProfilePicOverride() bool`

HasProfilePicOverride returns a boolean if a field has been set.

### GetStatus

`func (o *MutualFriend) GetStatus() UserStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MutualFriend) GetStatusOk() (*UserStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MutualFriend) SetStatus(v UserStatus)`

SetStatus sets Status field to given value.


### GetStatusDescription

`func (o *MutualFriend) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *MutualFriend) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *MutualFriend) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


