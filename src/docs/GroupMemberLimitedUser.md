# GroupMemberLimitedUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentAvatarTags** | Pointer to **[]string** |  | [optional] 
**CurrentAvatarThumbnailImageUrl** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**IconUrl** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 
**ProfilePicOverride** | Pointer to **string** |  | [optional] 
**ThumbnailUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGroupMemberLimitedUser

`func NewGroupMemberLimitedUser() *GroupMemberLimitedUser`

NewGroupMemberLimitedUser instantiates a new GroupMemberLimitedUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupMemberLimitedUserWithDefaults

`func NewGroupMemberLimitedUserWithDefaults() *GroupMemberLimitedUser`

NewGroupMemberLimitedUserWithDefaults instantiates a new GroupMemberLimitedUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentAvatarTags

`func (o *GroupMemberLimitedUser) GetCurrentAvatarTags() []string`

GetCurrentAvatarTags returns the CurrentAvatarTags field if non-nil, zero value otherwise.

### GetCurrentAvatarTagsOk

`func (o *GroupMemberLimitedUser) GetCurrentAvatarTagsOk() (*[]string, bool)`

GetCurrentAvatarTagsOk returns a tuple with the CurrentAvatarTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAvatarTags

`func (o *GroupMemberLimitedUser) SetCurrentAvatarTags(v []string)`

SetCurrentAvatarTags sets CurrentAvatarTags field to given value.

### HasCurrentAvatarTags

`func (o *GroupMemberLimitedUser) HasCurrentAvatarTags() bool`

HasCurrentAvatarTags returns a boolean if a field has been set.

### GetCurrentAvatarThumbnailImageUrl

`func (o *GroupMemberLimitedUser) GetCurrentAvatarThumbnailImageUrl() string`

GetCurrentAvatarThumbnailImageUrl returns the CurrentAvatarThumbnailImageUrl field if non-nil, zero value otherwise.

### GetCurrentAvatarThumbnailImageUrlOk

`func (o *GroupMemberLimitedUser) GetCurrentAvatarThumbnailImageUrlOk() (*string, bool)`

GetCurrentAvatarThumbnailImageUrlOk returns a tuple with the CurrentAvatarThumbnailImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAvatarThumbnailImageUrl

`func (o *GroupMemberLimitedUser) SetCurrentAvatarThumbnailImageUrl(v string)`

SetCurrentAvatarThumbnailImageUrl sets CurrentAvatarThumbnailImageUrl field to given value.

### HasCurrentAvatarThumbnailImageUrl

`func (o *GroupMemberLimitedUser) HasCurrentAvatarThumbnailImageUrl() bool`

HasCurrentAvatarThumbnailImageUrl returns a boolean if a field has been set.

### SetCurrentAvatarThumbnailImageUrlNil

`func (o *GroupMemberLimitedUser) SetCurrentAvatarThumbnailImageUrlNil(b bool)`

 SetCurrentAvatarThumbnailImageUrlNil sets the value for CurrentAvatarThumbnailImageUrl to be an explicit nil

### UnsetCurrentAvatarThumbnailImageUrl
`func (o *GroupMemberLimitedUser) UnsetCurrentAvatarThumbnailImageUrl()`

UnsetCurrentAvatarThumbnailImageUrl ensures that no value is present for CurrentAvatarThumbnailImageUrl, not even an explicit nil
### GetDisplayName

`func (o *GroupMemberLimitedUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GroupMemberLimitedUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GroupMemberLimitedUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GroupMemberLimitedUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIconUrl

`func (o *GroupMemberLimitedUser) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *GroupMemberLimitedUser) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *GroupMemberLimitedUser) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *GroupMemberLimitedUser) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetId

`func (o *GroupMemberLimitedUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupMemberLimitedUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupMemberLimitedUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupMemberLimitedUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProfilePicOverride

`func (o *GroupMemberLimitedUser) GetProfilePicOverride() string`

GetProfilePicOverride returns the ProfilePicOverride field if non-nil, zero value otherwise.

### GetProfilePicOverrideOk

`func (o *GroupMemberLimitedUser) GetProfilePicOverrideOk() (*string, bool)`

GetProfilePicOverrideOk returns a tuple with the ProfilePicOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePicOverride

`func (o *GroupMemberLimitedUser) SetProfilePicOverride(v string)`

SetProfilePicOverride sets ProfilePicOverride field to given value.

### HasProfilePicOverride

`func (o *GroupMemberLimitedUser) HasProfilePicOverride() bool`

HasProfilePicOverride returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *GroupMemberLimitedUser) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *GroupMemberLimitedUser) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *GroupMemberLimitedUser) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *GroupMemberLimitedUser) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### SetThumbnailUrlNil

`func (o *GroupMemberLimitedUser) SetThumbnailUrlNil(b bool)`

 SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil

### UnsetThumbnailUrl
`func (o *GroupMemberLimitedUser) UnsetThumbnailUrl()`

UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


