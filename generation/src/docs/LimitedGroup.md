# LimitedGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BannerId** | Pointer to **NullableString** |  | [optional] 
**BannerUrl** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Discriminator** | Pointer to **string** |  | [optional] 
**Galleries** | Pointer to [**[]GroupGallery**](GroupGallery.md) |   | [optional] 
**IconId** | Pointer to **NullableString** |  | [optional] 
**IconUrl** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsSearchable** | Pointer to **bool** |  | [optional] 
**MemberCount** | Pointer to **int32** |  | [optional] 
**MembershipStatus** | Pointer to [**GroupMemberStatus**](GroupMemberStatus.md) |  | [optional] [default to INACTIVE]
**Name** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 
**Rules** | Pointer to **NullableString** |  | [optional] 
**ShortCode** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |   | [optional] 

## Methods

### NewLimitedGroup

`func NewLimitedGroup() *LimitedGroup`

NewLimitedGroup instantiates a new LimitedGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLimitedGroupWithDefaults

`func NewLimitedGroupWithDefaults() *LimitedGroup`

NewLimitedGroupWithDefaults instantiates a new LimitedGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBannerId

`func (o *LimitedGroup) GetBannerId() string`

GetBannerId returns the BannerId field if non-nil, zero value otherwise.

### GetBannerIdOk

`func (o *LimitedGroup) GetBannerIdOk() (*string, bool)`

GetBannerIdOk returns a tuple with the BannerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerId

`func (o *LimitedGroup) SetBannerId(v string)`

SetBannerId sets BannerId field to given value.

### HasBannerId

`func (o *LimitedGroup) HasBannerId() bool`

HasBannerId returns a boolean if a field has been set.

### SetBannerIdNil

`func (o *LimitedGroup) SetBannerIdNil(b bool)`

 SetBannerIdNil sets the value for BannerId to be an explicit nil

### UnsetBannerId
`func (o *LimitedGroup) UnsetBannerId()`

UnsetBannerId ensures that no value is present for BannerId, not even an explicit nil
### GetBannerUrl

`func (o *LimitedGroup) GetBannerUrl() string`

GetBannerUrl returns the BannerUrl field if non-nil, zero value otherwise.

### GetBannerUrlOk

`func (o *LimitedGroup) GetBannerUrlOk() (*string, bool)`

GetBannerUrlOk returns a tuple with the BannerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerUrl

`func (o *LimitedGroup) SetBannerUrl(v string)`

SetBannerUrl sets BannerUrl field to given value.

### HasBannerUrl

`func (o *LimitedGroup) HasBannerUrl() bool`

HasBannerUrl returns a boolean if a field has been set.

### SetBannerUrlNil

`func (o *LimitedGroup) SetBannerUrlNil(b bool)`

 SetBannerUrlNil sets the value for BannerUrl to be an explicit nil

### UnsetBannerUrl
`func (o *LimitedGroup) UnsetBannerUrl()`

UnsetBannerUrl ensures that no value is present for BannerUrl, not even an explicit nil
### GetCreatedAt

`func (o *LimitedGroup) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LimitedGroup) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LimitedGroup) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LimitedGroup) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *LimitedGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LimitedGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LimitedGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LimitedGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscriminator

`func (o *LimitedGroup) GetDiscriminator() string`

GetDiscriminator returns the Discriminator field if non-nil, zero value otherwise.

### GetDiscriminatorOk

`func (o *LimitedGroup) GetDiscriminatorOk() (*string, bool)`

GetDiscriminatorOk returns a tuple with the Discriminator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscriminator

`func (o *LimitedGroup) SetDiscriminator(v string)`

SetDiscriminator sets Discriminator field to given value.

### HasDiscriminator

`func (o *LimitedGroup) HasDiscriminator() bool`

HasDiscriminator returns a boolean if a field has been set.

### GetGalleries

`func (o *LimitedGroup) GetGalleries() []GroupGallery`

GetGalleries returns the Galleries field if non-nil, zero value otherwise.

### GetGalleriesOk

`func (o *LimitedGroup) GetGalleriesOk() (*[]GroupGallery, bool)`

GetGalleriesOk returns a tuple with the Galleries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGalleries

`func (o *LimitedGroup) SetGalleries(v []GroupGallery)`

SetGalleries sets Galleries field to given value.

### HasGalleries

`func (o *LimitedGroup) HasGalleries() bool`

HasGalleries returns a boolean if a field has been set.

### GetIconId

`func (o *LimitedGroup) GetIconId() string`

GetIconId returns the IconId field if non-nil, zero value otherwise.

### GetIconIdOk

`func (o *LimitedGroup) GetIconIdOk() (*string, bool)`

GetIconIdOk returns a tuple with the IconId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconId

`func (o *LimitedGroup) SetIconId(v string)`

SetIconId sets IconId field to given value.

### HasIconId

`func (o *LimitedGroup) HasIconId() bool`

HasIconId returns a boolean if a field has been set.

### SetIconIdNil

`func (o *LimitedGroup) SetIconIdNil(b bool)`

 SetIconIdNil sets the value for IconId to be an explicit nil

### UnsetIconId
`func (o *LimitedGroup) UnsetIconId()`

UnsetIconId ensures that no value is present for IconId, not even an explicit nil
### GetIconUrl

`func (o *LimitedGroup) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *LimitedGroup) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *LimitedGroup) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *LimitedGroup) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *LimitedGroup) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *LimitedGroup) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetId

`func (o *LimitedGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LimitedGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LimitedGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LimitedGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsSearchable

`func (o *LimitedGroup) GetIsSearchable() bool`

GetIsSearchable returns the IsSearchable field if non-nil, zero value otherwise.

### GetIsSearchableOk

`func (o *LimitedGroup) GetIsSearchableOk() (*bool, bool)`

GetIsSearchableOk returns a tuple with the IsSearchable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSearchable

`func (o *LimitedGroup) SetIsSearchable(v bool)`

SetIsSearchable sets IsSearchable field to given value.

### HasIsSearchable

`func (o *LimitedGroup) HasIsSearchable() bool`

HasIsSearchable returns a boolean if a field has been set.

### GetMemberCount

`func (o *LimitedGroup) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *LimitedGroup) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *LimitedGroup) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.

### HasMemberCount

`func (o *LimitedGroup) HasMemberCount() bool`

HasMemberCount returns a boolean if a field has been set.

### GetMembershipStatus

`func (o *LimitedGroup) GetMembershipStatus() GroupMemberStatus`

GetMembershipStatus returns the MembershipStatus field if non-nil, zero value otherwise.

### GetMembershipStatusOk

`func (o *LimitedGroup) GetMembershipStatusOk() (*GroupMemberStatus, bool)`

GetMembershipStatusOk returns a tuple with the MembershipStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipStatus

`func (o *LimitedGroup) SetMembershipStatus(v GroupMemberStatus)`

SetMembershipStatus sets MembershipStatus field to given value.

### HasMembershipStatus

`func (o *LimitedGroup) HasMembershipStatus() bool`

HasMembershipStatus returns a boolean if a field has been set.

### GetName

`func (o *LimitedGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LimitedGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LimitedGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LimitedGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnerId

`func (o *LimitedGroup) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *LimitedGroup) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *LimitedGroup) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *LimitedGroup) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetRules

`func (o *LimitedGroup) GetRules() string`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *LimitedGroup) GetRulesOk() (*string, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *LimitedGroup) SetRules(v string)`

SetRules sets Rules field to given value.

### HasRules

`func (o *LimitedGroup) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRulesNil

`func (o *LimitedGroup) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *LimitedGroup) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil
### GetShortCode

`func (o *LimitedGroup) GetShortCode() string`

GetShortCode returns the ShortCode field if non-nil, zero value otherwise.

### GetShortCodeOk

`func (o *LimitedGroup) GetShortCodeOk() (*string, bool)`

GetShortCodeOk returns a tuple with the ShortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortCode

`func (o *LimitedGroup) SetShortCode(v string)`

SetShortCode sets ShortCode field to given value.

### HasShortCode

`func (o *LimitedGroup) HasShortCode() bool`

HasShortCode returns a boolean if a field has been set.

### GetTags

`func (o *LimitedGroup) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LimitedGroup) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LimitedGroup) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LimitedGroup) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


