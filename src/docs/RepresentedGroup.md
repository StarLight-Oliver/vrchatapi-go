# RepresentedGroup

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
**IsRepresenting** | Pointer to **bool** |  | [optional] 
**MemberCount** | Pointer to **int32** |  | [optional] 
**MemberVisibility** | Pointer to [**GroupUserVisibility**](GroupUserVisibility.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 
**Privacy** | Pointer to [**GroupPrivacy**](GroupPrivacy.md) |  | [optional] [default to DEFAULT]
**ShortCode** | Pointer to **string** |  | [optional] 

## Methods

### NewRepresentedGroup

`func NewRepresentedGroup() *RepresentedGroup`

NewRepresentedGroup instantiates a new RepresentedGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepresentedGroupWithDefaults

`func NewRepresentedGroupWithDefaults() *RepresentedGroup`

NewRepresentedGroupWithDefaults instantiates a new RepresentedGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBannerId

`func (o *RepresentedGroup) GetBannerId() string`

GetBannerId returns the BannerId field if non-nil, zero value otherwise.

### GetBannerIdOk

`func (o *RepresentedGroup) GetBannerIdOk() (*string, bool)`

GetBannerIdOk returns a tuple with the BannerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerId

`func (o *RepresentedGroup) SetBannerId(v string)`

SetBannerId sets BannerId field to given value.

### HasBannerId

`func (o *RepresentedGroup) HasBannerId() bool`

HasBannerId returns a boolean if a field has been set.

### SetBannerIdNil

`func (o *RepresentedGroup) SetBannerIdNil(b bool)`

 SetBannerIdNil sets the value for BannerId to be an explicit nil

### UnsetBannerId
`func (o *RepresentedGroup) UnsetBannerId()`

UnsetBannerId ensures that no value is present for BannerId, not even an explicit nil
### GetBannerUrl

`func (o *RepresentedGroup) GetBannerUrl() string`

GetBannerUrl returns the BannerUrl field if non-nil, zero value otherwise.

### GetBannerUrlOk

`func (o *RepresentedGroup) GetBannerUrlOk() (*string, bool)`

GetBannerUrlOk returns a tuple with the BannerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerUrl

`func (o *RepresentedGroup) SetBannerUrl(v string)`

SetBannerUrl sets BannerUrl field to given value.

### HasBannerUrl

`func (o *RepresentedGroup) HasBannerUrl() bool`

HasBannerUrl returns a boolean if a field has been set.

### SetBannerUrlNil

`func (o *RepresentedGroup) SetBannerUrlNil(b bool)`

 SetBannerUrlNil sets the value for BannerUrl to be an explicit nil

### UnsetBannerUrl
`func (o *RepresentedGroup) UnsetBannerUrl()`

UnsetBannerUrl ensures that no value is present for BannerUrl, not even an explicit nil
### GetDescription

`func (o *RepresentedGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RepresentedGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RepresentedGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RepresentedGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscriminator

`func (o *RepresentedGroup) GetDiscriminator() string`

GetDiscriminator returns the Discriminator field if non-nil, zero value otherwise.

### GetDiscriminatorOk

`func (o *RepresentedGroup) GetDiscriminatorOk() (*string, bool)`

GetDiscriminatorOk returns a tuple with the Discriminator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscriminator

`func (o *RepresentedGroup) SetDiscriminator(v string)`

SetDiscriminator sets Discriminator field to given value.

### HasDiscriminator

`func (o *RepresentedGroup) HasDiscriminator() bool`

HasDiscriminator returns a boolean if a field has been set.

### GetGroupId

`func (o *RepresentedGroup) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *RepresentedGroup) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *RepresentedGroup) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *RepresentedGroup) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetIconId

`func (o *RepresentedGroup) GetIconId() string`

GetIconId returns the IconId field if non-nil, zero value otherwise.

### GetIconIdOk

`func (o *RepresentedGroup) GetIconIdOk() (*string, bool)`

GetIconIdOk returns a tuple with the IconId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconId

`func (o *RepresentedGroup) SetIconId(v string)`

SetIconId sets IconId field to given value.

### HasIconId

`func (o *RepresentedGroup) HasIconId() bool`

HasIconId returns a boolean if a field has been set.

### SetIconIdNil

`func (o *RepresentedGroup) SetIconIdNil(b bool)`

 SetIconIdNil sets the value for IconId to be an explicit nil

### UnsetIconId
`func (o *RepresentedGroup) UnsetIconId()`

UnsetIconId ensures that no value is present for IconId, not even an explicit nil
### GetIconUrl

`func (o *RepresentedGroup) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *RepresentedGroup) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *RepresentedGroup) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *RepresentedGroup) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *RepresentedGroup) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *RepresentedGroup) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetIsRepresenting

`func (o *RepresentedGroup) GetIsRepresenting() bool`

GetIsRepresenting returns the IsRepresenting field if non-nil, zero value otherwise.

### GetIsRepresentingOk

`func (o *RepresentedGroup) GetIsRepresentingOk() (*bool, bool)`

GetIsRepresentingOk returns a tuple with the IsRepresenting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRepresenting

`func (o *RepresentedGroup) SetIsRepresenting(v bool)`

SetIsRepresenting sets IsRepresenting field to given value.

### HasIsRepresenting

`func (o *RepresentedGroup) HasIsRepresenting() bool`

HasIsRepresenting returns a boolean if a field has been set.

### GetMemberCount

`func (o *RepresentedGroup) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *RepresentedGroup) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *RepresentedGroup) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.

### HasMemberCount

`func (o *RepresentedGroup) HasMemberCount() bool`

HasMemberCount returns a boolean if a field has been set.

### GetMemberVisibility

`func (o *RepresentedGroup) GetMemberVisibility() GroupUserVisibility`

GetMemberVisibility returns the MemberVisibility field if non-nil, zero value otherwise.

### GetMemberVisibilityOk

`func (o *RepresentedGroup) GetMemberVisibilityOk() (*GroupUserVisibility, bool)`

GetMemberVisibilityOk returns a tuple with the MemberVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberVisibility

`func (o *RepresentedGroup) SetMemberVisibility(v GroupUserVisibility)`

SetMemberVisibility sets MemberVisibility field to given value.

### HasMemberVisibility

`func (o *RepresentedGroup) HasMemberVisibility() bool`

HasMemberVisibility returns a boolean if a field has been set.

### GetName

`func (o *RepresentedGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RepresentedGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RepresentedGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RepresentedGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnerId

`func (o *RepresentedGroup) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *RepresentedGroup) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *RepresentedGroup) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *RepresentedGroup) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetPrivacy

`func (o *RepresentedGroup) GetPrivacy() GroupPrivacy`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *RepresentedGroup) GetPrivacyOk() (*GroupPrivacy, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *RepresentedGroup) SetPrivacy(v GroupPrivacy)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *RepresentedGroup) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetShortCode

`func (o *RepresentedGroup) GetShortCode() string`

GetShortCode returns the ShortCode field if non-nil, zero value otherwise.

### GetShortCodeOk

`func (o *RepresentedGroup) GetShortCodeOk() (*string, bool)`

GetShortCodeOk returns a tuple with the ShortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortCode

`func (o *RepresentedGroup) SetShortCode(v string)`

SetShortCode sets ShortCode field to given value.

### HasShortCode

`func (o *RepresentedGroup) HasShortCode() bool`

HasShortCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


