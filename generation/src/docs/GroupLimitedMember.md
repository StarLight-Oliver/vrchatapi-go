# GroupLimitedMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BannedAt** | Pointer to **NullableTime** | Only visible via the /groups/:groupId/members endpoint, **not** when fetching a specific user. | [optional] 
**CreatedAt** | Pointer to **NullableTime** | Only visible via the /groups/:groupId/members endpoint, **not** when fetching a specific user. | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**HasJoinedFromPurchase** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsRepresenting** | Pointer to **bool** | Whether the user is representing the group. This makes the group show up above the name tag in-game. | [optional] [default to false]
**IsSubscribedToAnnouncements** | Pointer to **bool** |  | [optional] [default to false]
**IsSubscribedToEventAnnouncements** | Pointer to **bool** |  | [optional] 
**JoinedAt** | Pointer to **NullableTime** |  | [optional] 
**LastPostReadAt** | Pointer to **NullableTime** |  | [optional] 
**MRoleIds** | Pointer to **[]string** |  | [optional] 
**ManagerNotes** | Pointer to **NullableString** | Only visible via the /groups/:groupId/members endpoint, **not** when fetching a specific user. | [optional] 
**MembershipStatus** | Pointer to [**GroupMemberStatus**](GroupMemberStatus.md) |  | [optional] [default to INACTIVE]
**RoleIds** | Pointer to **[]string** |  | [optional] 
**UserId** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 

## Methods

### NewGroupLimitedMember

`func NewGroupLimitedMember() *GroupLimitedMember`

NewGroupLimitedMember instantiates a new GroupLimitedMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupLimitedMemberWithDefaults

`func NewGroupLimitedMemberWithDefaults() *GroupLimitedMember`

NewGroupLimitedMemberWithDefaults instantiates a new GroupLimitedMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBannedAt

`func (o *GroupLimitedMember) GetBannedAt() time.Time`

GetBannedAt returns the BannedAt field if non-nil, zero value otherwise.

### GetBannedAtOk

`func (o *GroupLimitedMember) GetBannedAtOk() (*time.Time, bool)`

GetBannedAtOk returns a tuple with the BannedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedAt

`func (o *GroupLimitedMember) SetBannedAt(v time.Time)`

SetBannedAt sets BannedAt field to given value.

### HasBannedAt

`func (o *GroupLimitedMember) HasBannedAt() bool`

HasBannedAt returns a boolean if a field has been set.

### SetBannedAtNil

`func (o *GroupLimitedMember) SetBannedAtNil(b bool)`

 SetBannedAtNil sets the value for BannedAt to be an explicit nil

### UnsetBannedAt
`func (o *GroupLimitedMember) UnsetBannedAt()`

UnsetBannedAt ensures that no value is present for BannedAt, not even an explicit nil
### GetCreatedAt

`func (o *GroupLimitedMember) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GroupLimitedMember) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GroupLimitedMember) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GroupLimitedMember) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GroupLimitedMember) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GroupLimitedMember) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetGroupId

`func (o *GroupLimitedMember) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupLimitedMember) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupLimitedMember) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *GroupLimitedMember) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetHasJoinedFromPurchase

`func (o *GroupLimitedMember) GetHasJoinedFromPurchase() bool`

GetHasJoinedFromPurchase returns the HasJoinedFromPurchase field if non-nil, zero value otherwise.

### GetHasJoinedFromPurchaseOk

`func (o *GroupLimitedMember) GetHasJoinedFromPurchaseOk() (*bool, bool)`

GetHasJoinedFromPurchaseOk returns a tuple with the HasJoinedFromPurchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasJoinedFromPurchase

`func (o *GroupLimitedMember) SetHasJoinedFromPurchase(v bool)`

SetHasJoinedFromPurchase sets HasJoinedFromPurchase field to given value.

### HasHasJoinedFromPurchase

`func (o *GroupLimitedMember) HasHasJoinedFromPurchase() bool`

HasHasJoinedFromPurchase returns a boolean if a field has been set.

### GetId

`func (o *GroupLimitedMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupLimitedMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupLimitedMember) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupLimitedMember) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsRepresenting

`func (o *GroupLimitedMember) GetIsRepresenting() bool`

GetIsRepresenting returns the IsRepresenting field if non-nil, zero value otherwise.

### GetIsRepresentingOk

`func (o *GroupLimitedMember) GetIsRepresentingOk() (*bool, bool)`

GetIsRepresentingOk returns a tuple with the IsRepresenting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRepresenting

`func (o *GroupLimitedMember) SetIsRepresenting(v bool)`

SetIsRepresenting sets IsRepresenting field to given value.

### HasIsRepresenting

`func (o *GroupLimitedMember) HasIsRepresenting() bool`

HasIsRepresenting returns a boolean if a field has been set.

### GetIsSubscribedToAnnouncements

`func (o *GroupLimitedMember) GetIsSubscribedToAnnouncements() bool`

GetIsSubscribedToAnnouncements returns the IsSubscribedToAnnouncements field if non-nil, zero value otherwise.

### GetIsSubscribedToAnnouncementsOk

`func (o *GroupLimitedMember) GetIsSubscribedToAnnouncementsOk() (*bool, bool)`

GetIsSubscribedToAnnouncementsOk returns a tuple with the IsSubscribedToAnnouncements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscribedToAnnouncements

`func (o *GroupLimitedMember) SetIsSubscribedToAnnouncements(v bool)`

SetIsSubscribedToAnnouncements sets IsSubscribedToAnnouncements field to given value.

### HasIsSubscribedToAnnouncements

`func (o *GroupLimitedMember) HasIsSubscribedToAnnouncements() bool`

HasIsSubscribedToAnnouncements returns a boolean if a field has been set.

### GetIsSubscribedToEventAnnouncements

`func (o *GroupLimitedMember) GetIsSubscribedToEventAnnouncements() bool`

GetIsSubscribedToEventAnnouncements returns the IsSubscribedToEventAnnouncements field if non-nil, zero value otherwise.

### GetIsSubscribedToEventAnnouncementsOk

`func (o *GroupLimitedMember) GetIsSubscribedToEventAnnouncementsOk() (*bool, bool)`

GetIsSubscribedToEventAnnouncementsOk returns a tuple with the IsSubscribedToEventAnnouncements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscribedToEventAnnouncements

`func (o *GroupLimitedMember) SetIsSubscribedToEventAnnouncements(v bool)`

SetIsSubscribedToEventAnnouncements sets IsSubscribedToEventAnnouncements field to given value.

### HasIsSubscribedToEventAnnouncements

`func (o *GroupLimitedMember) HasIsSubscribedToEventAnnouncements() bool`

HasIsSubscribedToEventAnnouncements returns a boolean if a field has been set.

### GetJoinedAt

`func (o *GroupLimitedMember) GetJoinedAt() time.Time`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *GroupLimitedMember) GetJoinedAtOk() (*time.Time, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *GroupLimitedMember) SetJoinedAt(v time.Time)`

SetJoinedAt sets JoinedAt field to given value.

### HasJoinedAt

`func (o *GroupLimitedMember) HasJoinedAt() bool`

HasJoinedAt returns a boolean if a field has been set.

### SetJoinedAtNil

`func (o *GroupLimitedMember) SetJoinedAtNil(b bool)`

 SetJoinedAtNil sets the value for JoinedAt to be an explicit nil

### UnsetJoinedAt
`func (o *GroupLimitedMember) UnsetJoinedAt()`

UnsetJoinedAt ensures that no value is present for JoinedAt, not even an explicit nil
### GetLastPostReadAt

`func (o *GroupLimitedMember) GetLastPostReadAt() time.Time`

GetLastPostReadAt returns the LastPostReadAt field if non-nil, zero value otherwise.

### GetLastPostReadAtOk

`func (o *GroupLimitedMember) GetLastPostReadAtOk() (*time.Time, bool)`

GetLastPostReadAtOk returns a tuple with the LastPostReadAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPostReadAt

`func (o *GroupLimitedMember) SetLastPostReadAt(v time.Time)`

SetLastPostReadAt sets LastPostReadAt field to given value.

### HasLastPostReadAt

`func (o *GroupLimitedMember) HasLastPostReadAt() bool`

HasLastPostReadAt returns a boolean if a field has been set.

### SetLastPostReadAtNil

`func (o *GroupLimitedMember) SetLastPostReadAtNil(b bool)`

 SetLastPostReadAtNil sets the value for LastPostReadAt to be an explicit nil

### UnsetLastPostReadAt
`func (o *GroupLimitedMember) UnsetLastPostReadAt()`

UnsetLastPostReadAt ensures that no value is present for LastPostReadAt, not even an explicit nil
### GetMRoleIds

`func (o *GroupLimitedMember) GetMRoleIds() []string`

GetMRoleIds returns the MRoleIds field if non-nil, zero value otherwise.

### GetMRoleIdsOk

`func (o *GroupLimitedMember) GetMRoleIdsOk() (*[]string, bool)`

GetMRoleIdsOk returns a tuple with the MRoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMRoleIds

`func (o *GroupLimitedMember) SetMRoleIds(v []string)`

SetMRoleIds sets MRoleIds field to given value.

### HasMRoleIds

`func (o *GroupLimitedMember) HasMRoleIds() bool`

HasMRoleIds returns a boolean if a field has been set.

### GetManagerNotes

`func (o *GroupLimitedMember) GetManagerNotes() string`

GetManagerNotes returns the ManagerNotes field if non-nil, zero value otherwise.

### GetManagerNotesOk

`func (o *GroupLimitedMember) GetManagerNotesOk() (*string, bool)`

GetManagerNotesOk returns a tuple with the ManagerNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerNotes

`func (o *GroupLimitedMember) SetManagerNotes(v string)`

SetManagerNotes sets ManagerNotes field to given value.

### HasManagerNotes

`func (o *GroupLimitedMember) HasManagerNotes() bool`

HasManagerNotes returns a boolean if a field has been set.

### SetManagerNotesNil

`func (o *GroupLimitedMember) SetManagerNotesNil(b bool)`

 SetManagerNotesNil sets the value for ManagerNotes to be an explicit nil

### UnsetManagerNotes
`func (o *GroupLimitedMember) UnsetManagerNotes()`

UnsetManagerNotes ensures that no value is present for ManagerNotes, not even an explicit nil
### GetMembershipStatus

`func (o *GroupLimitedMember) GetMembershipStatus() GroupMemberStatus`

GetMembershipStatus returns the MembershipStatus field if non-nil, zero value otherwise.

### GetMembershipStatusOk

`func (o *GroupLimitedMember) GetMembershipStatusOk() (*GroupMemberStatus, bool)`

GetMembershipStatusOk returns a tuple with the MembershipStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipStatus

`func (o *GroupLimitedMember) SetMembershipStatus(v GroupMemberStatus)`

SetMembershipStatus sets MembershipStatus field to given value.

### HasMembershipStatus

`func (o *GroupLimitedMember) HasMembershipStatus() bool`

HasMembershipStatus returns a boolean if a field has been set.

### GetRoleIds

`func (o *GroupLimitedMember) GetRoleIds() []string`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *GroupLimitedMember) GetRoleIdsOk() (*[]string, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *GroupLimitedMember) SetRoleIds(v []string)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *GroupLimitedMember) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.

### GetUserId

`func (o *GroupLimitedMember) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GroupLimitedMember) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GroupLimitedMember) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GroupLimitedMember) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVisibility

`func (o *GroupLimitedMember) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GroupLimitedMember) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GroupLimitedMember) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GroupLimitedMember) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


