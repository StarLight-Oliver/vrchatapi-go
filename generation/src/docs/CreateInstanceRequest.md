# CreateInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgeGate** | Pointer to **bool** |  | [optional] [default to false]
**CalendarEntryId** | Pointer to **string** |  | [optional] 
**CanRequestInvite** | Pointer to **bool** | Only applies to invite type instances to make them invite+ | [optional] [default to false]
**ClosedAt** | Pointer to **time.Time** | The time after which users won&#39;t be allowed to join the instance. This doesn&#39;t work for public instances. | [optional] 
**ContentSettings** | Pointer to [**InstanceContentSettings**](InstanceContentSettings.md) |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**GroupAccessType** | Pointer to [**GroupAccessType**](GroupAccessType.md) |  | [optional] [default to MEMBERS]
**HardClose** | Pointer to **bool** | Currently unused, but will eventually be a flag to set if the closing of the instance should kick people. | [optional] [default to false]
**InstancePersistenceEnabled** | Pointer to **NullableBool** |  | [optional] 
**InviteOnly** | Pointer to **bool** |  | [optional] [default to false]
**OwnerId** | Pointer to **NullableString** | A groupId if the instance type is \&quot;group\&quot;, null if instance type is public, or a userId otherwise | [optional] 
**PlayerPersistenceEnabled** | Pointer to **NullableBool** |  | [optional] 
**QueueEnabled** | Pointer to **bool** |  | [optional] [default to false]
**Region** | [**InstanceRegion**](InstanceRegion.md) |  | [default to US]
**RoleIds** | Pointer to **[]string** | Group roleIds that are allowed to join if the type is \&quot;group\&quot; and groupAccessType is \&quot;member\&quot; | [optional] 
**Type** | [**InstanceType**](InstanceType.md) |  | 
**WorldId** | **string** | WorldID be \&quot;offline\&quot; on User profiles if you are not friends with that user. | 

## Methods

### NewCreateInstanceRequest

`func NewCreateInstanceRequest(region InstanceRegion, type_ InstanceType, worldId string, ) *CreateInstanceRequest`

NewCreateInstanceRequest instantiates a new CreateInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInstanceRequestWithDefaults

`func NewCreateInstanceRequestWithDefaults() *CreateInstanceRequest`

NewCreateInstanceRequestWithDefaults instantiates a new CreateInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgeGate

`func (o *CreateInstanceRequest) GetAgeGate() bool`

GetAgeGate returns the AgeGate field if non-nil, zero value otherwise.

### GetAgeGateOk

`func (o *CreateInstanceRequest) GetAgeGateOk() (*bool, bool)`

GetAgeGateOk returns a tuple with the AgeGate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeGate

`func (o *CreateInstanceRequest) SetAgeGate(v bool)`

SetAgeGate sets AgeGate field to given value.

### HasAgeGate

`func (o *CreateInstanceRequest) HasAgeGate() bool`

HasAgeGate returns a boolean if a field has been set.

### GetCalendarEntryId

`func (o *CreateInstanceRequest) GetCalendarEntryId() string`

GetCalendarEntryId returns the CalendarEntryId field if non-nil, zero value otherwise.

### GetCalendarEntryIdOk

`func (o *CreateInstanceRequest) GetCalendarEntryIdOk() (*string, bool)`

GetCalendarEntryIdOk returns a tuple with the CalendarEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarEntryId

`func (o *CreateInstanceRequest) SetCalendarEntryId(v string)`

SetCalendarEntryId sets CalendarEntryId field to given value.

### HasCalendarEntryId

`func (o *CreateInstanceRequest) HasCalendarEntryId() bool`

HasCalendarEntryId returns a boolean if a field has been set.

### GetCanRequestInvite

`func (o *CreateInstanceRequest) GetCanRequestInvite() bool`

GetCanRequestInvite returns the CanRequestInvite field if non-nil, zero value otherwise.

### GetCanRequestInviteOk

`func (o *CreateInstanceRequest) GetCanRequestInviteOk() (*bool, bool)`

GetCanRequestInviteOk returns a tuple with the CanRequestInvite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRequestInvite

`func (o *CreateInstanceRequest) SetCanRequestInvite(v bool)`

SetCanRequestInvite sets CanRequestInvite field to given value.

### HasCanRequestInvite

`func (o *CreateInstanceRequest) HasCanRequestInvite() bool`

HasCanRequestInvite returns a boolean if a field has been set.

### GetClosedAt

`func (o *CreateInstanceRequest) GetClosedAt() time.Time`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *CreateInstanceRequest) GetClosedAtOk() (*time.Time, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *CreateInstanceRequest) SetClosedAt(v time.Time)`

SetClosedAt sets ClosedAt field to given value.

### HasClosedAt

`func (o *CreateInstanceRequest) HasClosedAt() bool`

HasClosedAt returns a boolean if a field has been set.

### GetContentSettings

`func (o *CreateInstanceRequest) GetContentSettings() InstanceContentSettings`

GetContentSettings returns the ContentSettings field if non-nil, zero value otherwise.

### GetContentSettingsOk

`func (o *CreateInstanceRequest) GetContentSettingsOk() (*InstanceContentSettings, bool)`

GetContentSettingsOk returns a tuple with the ContentSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSettings

`func (o *CreateInstanceRequest) SetContentSettings(v InstanceContentSettings)`

SetContentSettings sets ContentSettings field to given value.

### HasContentSettings

`func (o *CreateInstanceRequest) HasContentSettings() bool`

HasContentSettings returns a boolean if a field has been set.

### GetDisplayName

`func (o *CreateInstanceRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateInstanceRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateInstanceRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateInstanceRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *CreateInstanceRequest) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CreateInstanceRequest) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetGroupAccessType

`func (o *CreateInstanceRequest) GetGroupAccessType() GroupAccessType`

GetGroupAccessType returns the GroupAccessType field if non-nil, zero value otherwise.

### GetGroupAccessTypeOk

`func (o *CreateInstanceRequest) GetGroupAccessTypeOk() (*GroupAccessType, bool)`

GetGroupAccessTypeOk returns a tuple with the GroupAccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAccessType

`func (o *CreateInstanceRequest) SetGroupAccessType(v GroupAccessType)`

SetGroupAccessType sets GroupAccessType field to given value.

### HasGroupAccessType

`func (o *CreateInstanceRequest) HasGroupAccessType() bool`

HasGroupAccessType returns a boolean if a field has been set.

### GetHardClose

`func (o *CreateInstanceRequest) GetHardClose() bool`

GetHardClose returns the HardClose field if non-nil, zero value otherwise.

### GetHardCloseOk

`func (o *CreateInstanceRequest) GetHardCloseOk() (*bool, bool)`

GetHardCloseOk returns a tuple with the HardClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardClose

`func (o *CreateInstanceRequest) SetHardClose(v bool)`

SetHardClose sets HardClose field to given value.

### HasHardClose

`func (o *CreateInstanceRequest) HasHardClose() bool`

HasHardClose returns a boolean if a field has been set.

### GetInstancePersistenceEnabled

`func (o *CreateInstanceRequest) GetInstancePersistenceEnabled() bool`

GetInstancePersistenceEnabled returns the InstancePersistenceEnabled field if non-nil, zero value otherwise.

### GetInstancePersistenceEnabledOk

`func (o *CreateInstanceRequest) GetInstancePersistenceEnabledOk() (*bool, bool)`

GetInstancePersistenceEnabledOk returns a tuple with the InstancePersistenceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancePersistenceEnabled

`func (o *CreateInstanceRequest) SetInstancePersistenceEnabled(v bool)`

SetInstancePersistenceEnabled sets InstancePersistenceEnabled field to given value.

### HasInstancePersistenceEnabled

`func (o *CreateInstanceRequest) HasInstancePersistenceEnabled() bool`

HasInstancePersistenceEnabled returns a boolean if a field has been set.

### SetInstancePersistenceEnabledNil

`func (o *CreateInstanceRequest) SetInstancePersistenceEnabledNil(b bool)`

 SetInstancePersistenceEnabledNil sets the value for InstancePersistenceEnabled to be an explicit nil

### UnsetInstancePersistenceEnabled
`func (o *CreateInstanceRequest) UnsetInstancePersistenceEnabled()`

UnsetInstancePersistenceEnabled ensures that no value is present for InstancePersistenceEnabled, not even an explicit nil
### GetInviteOnly

`func (o *CreateInstanceRequest) GetInviteOnly() bool`

GetInviteOnly returns the InviteOnly field if non-nil, zero value otherwise.

### GetInviteOnlyOk

`func (o *CreateInstanceRequest) GetInviteOnlyOk() (*bool, bool)`

GetInviteOnlyOk returns a tuple with the InviteOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteOnly

`func (o *CreateInstanceRequest) SetInviteOnly(v bool)`

SetInviteOnly sets InviteOnly field to given value.

### HasInviteOnly

`func (o *CreateInstanceRequest) HasInviteOnly() bool`

HasInviteOnly returns a boolean if a field has been set.

### GetOwnerId

`func (o *CreateInstanceRequest) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *CreateInstanceRequest) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *CreateInstanceRequest) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *CreateInstanceRequest) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *CreateInstanceRequest) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *CreateInstanceRequest) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetPlayerPersistenceEnabled

`func (o *CreateInstanceRequest) GetPlayerPersistenceEnabled() bool`

GetPlayerPersistenceEnabled returns the PlayerPersistenceEnabled field if non-nil, zero value otherwise.

### GetPlayerPersistenceEnabledOk

`func (o *CreateInstanceRequest) GetPlayerPersistenceEnabledOk() (*bool, bool)`

GetPlayerPersistenceEnabledOk returns a tuple with the PlayerPersistenceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerPersistenceEnabled

`func (o *CreateInstanceRequest) SetPlayerPersistenceEnabled(v bool)`

SetPlayerPersistenceEnabled sets PlayerPersistenceEnabled field to given value.

### HasPlayerPersistenceEnabled

`func (o *CreateInstanceRequest) HasPlayerPersistenceEnabled() bool`

HasPlayerPersistenceEnabled returns a boolean if a field has been set.

### SetPlayerPersistenceEnabledNil

`func (o *CreateInstanceRequest) SetPlayerPersistenceEnabledNil(b bool)`

 SetPlayerPersistenceEnabledNil sets the value for PlayerPersistenceEnabled to be an explicit nil

### UnsetPlayerPersistenceEnabled
`func (o *CreateInstanceRequest) UnsetPlayerPersistenceEnabled()`

UnsetPlayerPersistenceEnabled ensures that no value is present for PlayerPersistenceEnabled, not even an explicit nil
### GetQueueEnabled

`func (o *CreateInstanceRequest) GetQueueEnabled() bool`

GetQueueEnabled returns the QueueEnabled field if non-nil, zero value otherwise.

### GetQueueEnabledOk

`func (o *CreateInstanceRequest) GetQueueEnabledOk() (*bool, bool)`

GetQueueEnabledOk returns a tuple with the QueueEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueEnabled

`func (o *CreateInstanceRequest) SetQueueEnabled(v bool)`

SetQueueEnabled sets QueueEnabled field to given value.

### HasQueueEnabled

`func (o *CreateInstanceRequest) HasQueueEnabled() bool`

HasQueueEnabled returns a boolean if a field has been set.

### GetRegion

`func (o *CreateInstanceRequest) GetRegion() InstanceRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateInstanceRequest) GetRegionOk() (*InstanceRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateInstanceRequest) SetRegion(v InstanceRegion)`

SetRegion sets Region field to given value.


### GetRoleIds

`func (o *CreateInstanceRequest) GetRoleIds() []string`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *CreateInstanceRequest) GetRoleIdsOk() (*[]string, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *CreateInstanceRequest) SetRoleIds(v []string)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *CreateInstanceRequest) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.

### GetType

`func (o *CreateInstanceRequest) GetType() InstanceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateInstanceRequest) GetTypeOk() (*InstanceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateInstanceRequest) SetType(v InstanceType)`

SetType sets Type field to given value.


### GetWorldId

`func (o *CreateInstanceRequest) GetWorldId() string`

GetWorldId returns the WorldId field if non-nil, zero value otherwise.

### GetWorldIdOk

`func (o *CreateInstanceRequest) GetWorldIdOk() (*string, bool)`

GetWorldIdOk returns a tuple with the WorldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorldId

`func (o *CreateInstanceRequest) SetWorldId(v string)`

SetWorldId sets WorldId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


