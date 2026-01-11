# LimitedUserGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BannerId** | Pointer to **NullableString** |  | [optional] 
**BannerUrl** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Discriminator** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**IconId** | Pointer to **NullableString** |  | [optional] 
**IconUrl** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsRepresenting** | Pointer to **bool** |  | [optional] 
**LastPostCreatedAt** | Pointer to **NullableTime** |  | [optional] 
**LastPostReadAt** | Pointer to **NullableTime** |  | [optional] 
**MemberCount** | Pointer to **int32** |  | [optional] 
**MemberVisibility** | Pointer to **string** |  | [optional] 
**MutualGroup** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 
**Privacy** | Pointer to **string** |  | [optional] 
**ShortCode** | Pointer to **string** |  | [optional] 

## Methods

### NewLimitedUserGroups

`func NewLimitedUserGroups() *LimitedUserGroups`

NewLimitedUserGroups instantiates a new LimitedUserGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLimitedUserGroupsWithDefaults

`func NewLimitedUserGroupsWithDefaults() *LimitedUserGroups`

NewLimitedUserGroupsWithDefaults instantiates a new LimitedUserGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBannerId

`func (o *LimitedUserGroups) GetBannerId() string`

GetBannerId returns the BannerId field if non-nil, zero value otherwise.

### GetBannerIdOk

`func (o *LimitedUserGroups) GetBannerIdOk() (*string, bool)`

GetBannerIdOk returns a tuple with the BannerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerId

`func (o *LimitedUserGroups) SetBannerId(v string)`

SetBannerId sets BannerId field to given value.

### HasBannerId

`func (o *LimitedUserGroups) HasBannerId() bool`

HasBannerId returns a boolean if a field has been set.

### SetBannerIdNil

`func (o *LimitedUserGroups) SetBannerIdNil(b bool)`

 SetBannerIdNil sets the value for BannerId to be an explicit nil

### UnsetBannerId
`func (o *LimitedUserGroups) UnsetBannerId()`

UnsetBannerId ensures that no value is present for BannerId, not even an explicit nil
### GetBannerUrl

`func (o *LimitedUserGroups) GetBannerUrl() string`

GetBannerUrl returns the BannerUrl field if non-nil, zero value otherwise.

### GetBannerUrlOk

`func (o *LimitedUserGroups) GetBannerUrlOk() (*string, bool)`

GetBannerUrlOk returns a tuple with the BannerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerUrl

`func (o *LimitedUserGroups) SetBannerUrl(v string)`

SetBannerUrl sets BannerUrl field to given value.

### HasBannerUrl

`func (o *LimitedUserGroups) HasBannerUrl() bool`

HasBannerUrl returns a boolean if a field has been set.

### SetBannerUrlNil

`func (o *LimitedUserGroups) SetBannerUrlNil(b bool)`

 SetBannerUrlNil sets the value for BannerUrl to be an explicit nil

### UnsetBannerUrl
`func (o *LimitedUserGroups) UnsetBannerUrl()`

UnsetBannerUrl ensures that no value is present for BannerUrl, not even an explicit nil
### GetDescription

`func (o *LimitedUserGroups) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LimitedUserGroups) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LimitedUserGroups) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LimitedUserGroups) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscriminator

`func (o *LimitedUserGroups) GetDiscriminator() string`

GetDiscriminator returns the Discriminator field if non-nil, zero value otherwise.

### GetDiscriminatorOk

`func (o *LimitedUserGroups) GetDiscriminatorOk() (*string, bool)`

GetDiscriminatorOk returns a tuple with the Discriminator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscriminator

`func (o *LimitedUserGroups) SetDiscriminator(v string)`

SetDiscriminator sets Discriminator field to given value.

### HasDiscriminator

`func (o *LimitedUserGroups) HasDiscriminator() bool`

HasDiscriminator returns a boolean if a field has been set.

### GetGroupId

`func (o *LimitedUserGroups) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *LimitedUserGroups) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *LimitedUserGroups) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *LimitedUserGroups) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetIconId

`func (o *LimitedUserGroups) GetIconId() string`

GetIconId returns the IconId field if non-nil, zero value otherwise.

### GetIconIdOk

`func (o *LimitedUserGroups) GetIconIdOk() (*string, bool)`

GetIconIdOk returns a tuple with the IconId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconId

`func (o *LimitedUserGroups) SetIconId(v string)`

SetIconId sets IconId field to given value.

### HasIconId

`func (o *LimitedUserGroups) HasIconId() bool`

HasIconId returns a boolean if a field has been set.

### SetIconIdNil

`func (o *LimitedUserGroups) SetIconIdNil(b bool)`

 SetIconIdNil sets the value for IconId to be an explicit nil

### UnsetIconId
`func (o *LimitedUserGroups) UnsetIconId()`

UnsetIconId ensures that no value is present for IconId, not even an explicit nil
### GetIconUrl

`func (o *LimitedUserGroups) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *LimitedUserGroups) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *LimitedUserGroups) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *LimitedUserGroups) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *LimitedUserGroups) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *LimitedUserGroups) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetId

`func (o *LimitedUserGroups) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LimitedUserGroups) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LimitedUserGroups) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LimitedUserGroups) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsRepresenting

`func (o *LimitedUserGroups) GetIsRepresenting() bool`

GetIsRepresenting returns the IsRepresenting field if non-nil, zero value otherwise.

### GetIsRepresentingOk

`func (o *LimitedUserGroups) GetIsRepresentingOk() (*bool, bool)`

GetIsRepresentingOk returns a tuple with the IsRepresenting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRepresenting

`func (o *LimitedUserGroups) SetIsRepresenting(v bool)`

SetIsRepresenting sets IsRepresenting field to given value.

### HasIsRepresenting

`func (o *LimitedUserGroups) HasIsRepresenting() bool`

HasIsRepresenting returns a boolean if a field has been set.

### GetLastPostCreatedAt

`func (o *LimitedUserGroups) GetLastPostCreatedAt() time.Time`

GetLastPostCreatedAt returns the LastPostCreatedAt field if non-nil, zero value otherwise.

### GetLastPostCreatedAtOk

`func (o *LimitedUserGroups) GetLastPostCreatedAtOk() (*time.Time, bool)`

GetLastPostCreatedAtOk returns a tuple with the LastPostCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPostCreatedAt

`func (o *LimitedUserGroups) SetLastPostCreatedAt(v time.Time)`

SetLastPostCreatedAt sets LastPostCreatedAt field to given value.

### HasLastPostCreatedAt

`func (o *LimitedUserGroups) HasLastPostCreatedAt() bool`

HasLastPostCreatedAt returns a boolean if a field has been set.

### SetLastPostCreatedAtNil

`func (o *LimitedUserGroups) SetLastPostCreatedAtNil(b bool)`

 SetLastPostCreatedAtNil sets the value for LastPostCreatedAt to be an explicit nil

### UnsetLastPostCreatedAt
`func (o *LimitedUserGroups) UnsetLastPostCreatedAt()`

UnsetLastPostCreatedAt ensures that no value is present for LastPostCreatedAt, not even an explicit nil
### GetLastPostReadAt

`func (o *LimitedUserGroups) GetLastPostReadAt() time.Time`

GetLastPostReadAt returns the LastPostReadAt field if non-nil, zero value otherwise.

### GetLastPostReadAtOk

`func (o *LimitedUserGroups) GetLastPostReadAtOk() (*time.Time, bool)`

GetLastPostReadAtOk returns a tuple with the LastPostReadAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPostReadAt

`func (o *LimitedUserGroups) SetLastPostReadAt(v time.Time)`

SetLastPostReadAt sets LastPostReadAt field to given value.

### HasLastPostReadAt

`func (o *LimitedUserGroups) HasLastPostReadAt() bool`

HasLastPostReadAt returns a boolean if a field has been set.

### SetLastPostReadAtNil

`func (o *LimitedUserGroups) SetLastPostReadAtNil(b bool)`

 SetLastPostReadAtNil sets the value for LastPostReadAt to be an explicit nil

### UnsetLastPostReadAt
`func (o *LimitedUserGroups) UnsetLastPostReadAt()`

UnsetLastPostReadAt ensures that no value is present for LastPostReadAt, not even an explicit nil
### GetMemberCount

`func (o *LimitedUserGroups) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *LimitedUserGroups) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *LimitedUserGroups) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.

### HasMemberCount

`func (o *LimitedUserGroups) HasMemberCount() bool`

HasMemberCount returns a boolean if a field has been set.

### GetMemberVisibility

`func (o *LimitedUserGroups) GetMemberVisibility() string`

GetMemberVisibility returns the MemberVisibility field if non-nil, zero value otherwise.

### GetMemberVisibilityOk

`func (o *LimitedUserGroups) GetMemberVisibilityOk() (*string, bool)`

GetMemberVisibilityOk returns a tuple with the MemberVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberVisibility

`func (o *LimitedUserGroups) SetMemberVisibility(v string)`

SetMemberVisibility sets MemberVisibility field to given value.

### HasMemberVisibility

`func (o *LimitedUserGroups) HasMemberVisibility() bool`

HasMemberVisibility returns a boolean if a field has been set.

### GetMutualGroup

`func (o *LimitedUserGroups) GetMutualGroup() bool`

GetMutualGroup returns the MutualGroup field if non-nil, zero value otherwise.

### GetMutualGroupOk

`func (o *LimitedUserGroups) GetMutualGroupOk() (*bool, bool)`

GetMutualGroupOk returns a tuple with the MutualGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutualGroup

`func (o *LimitedUserGroups) SetMutualGroup(v bool)`

SetMutualGroup sets MutualGroup field to given value.

### HasMutualGroup

`func (o *LimitedUserGroups) HasMutualGroup() bool`

HasMutualGroup returns a boolean if a field has been set.

### GetName

`func (o *LimitedUserGroups) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LimitedUserGroups) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LimitedUserGroups) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LimitedUserGroups) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnerId

`func (o *LimitedUserGroups) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *LimitedUserGroups) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *LimitedUserGroups) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *LimitedUserGroups) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetPrivacy

`func (o *LimitedUserGroups) GetPrivacy() string`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *LimitedUserGroups) GetPrivacyOk() (*string, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *LimitedUserGroups) SetPrivacy(v string)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *LimitedUserGroups) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetShortCode

`func (o *LimitedUserGroups) GetShortCode() string`

GetShortCode returns the ShortCode field if non-nil, zero value otherwise.

### GetShortCodeOk

`func (o *LimitedUserGroups) GetShortCodeOk() (*string, bool)`

GetShortCodeOk returns a tuple with the ShortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortCode

`func (o *LimitedUserGroups) SetShortCode(v string)`

SetShortCode sets ShortCode field to given value.

### HasShortCode

`func (o *LimitedUserGroups) HasShortCode() bool`

HasShortCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


