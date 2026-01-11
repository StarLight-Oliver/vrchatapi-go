# Instance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] [default to true]
**AgeGate** | Pointer to **NullableBool** |  | [optional] 
**CalendarEntryId** | Pointer to **NullableString** |  | [optional] 
**CanRequestInvite** | Pointer to **bool** |  | [optional] [default to true]
**Capacity** | Pointer to **int32** |  | [optional] 
**ClientNumber** | **string** | Always returns \&quot;unknown\&quot;. | 
**ClosedAt** | Pointer to **NullableTime** |  | [optional] 
**ContentSettings** | Pointer to [**InstanceContentSettings**](InstanceContentSettings.md) |  | [optional] 
**CreatorId** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Friends** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 
**Full** | **bool** |  | [default to false]
**GameServerVersion** | Pointer to **NullableInt32** |  | [optional] 
**GroupAccessType** | Pointer to [**GroupAccessType**](GroupAccessType.md) |  | [optional] [default to MEMBERS]
**HardClose** | Pointer to **NullableBool** |  | [optional] 
**HasCapacityForYou** | Pointer to **bool** |  | [optional] 
**Hidden** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 
**Id** | **string** | InstanceID can be \&quot;offline\&quot; on User profiles if you are not friends with that user and \&quot;private\&quot; if you are friends and user is in private instance. | 
**InstanceId** | **string** | InstanceID can be \&quot;offline\&quot; on User profiles if you are not friends with that user and \&quot;private\&quot; if you are friends and user is in private instance. | 
**InstancePersistenceEnabled** | Pointer to **NullableBool** |  | [optional] 
**Location** | **string** | Represents a unique location, consisting of a world identifier and an instance identifier, or \&quot;offline\&quot; if the user is not on your friends list. | 
**NUsers** | **int32** |  | 
**Name** | **string** |  | 
**Nonce** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **NullableString** | A groupId if the instance type is \&quot;group\&quot;, null if instance type is public, or a userId otherwise | [optional] 
**Permanent** | **bool** |  | [default to false]
**PhotonRegion** | [**Region**](Region.md) |  | [default to US]
**Platforms** | [**InstancePlatforms**](InstancePlatforms.md) |  | 
**PlayerPersistenceEnabled** | Pointer to **NullableBool** |  | [optional] 
**Private** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 
**QueueEnabled** | **bool** |  | 
**QueueSize** | **int32** |  | 
**RecommendedCapacity** | **int32** |  | 
**Region** | [**InstanceRegion**](InstanceRegion.md) |  | [default to US]
**RoleRestricted** | Pointer to **bool** |  | [optional] 
**SecureName** | **string** |  | 
**ShortName** | Pointer to **NullableString** |  | [optional] 
**Strict** | **bool** |  | 
**Tags** | **[]string** | The tags array on Instances usually contain the language tags of the people in the instance.  | 
**Type** | [**InstanceType**](InstanceType.md) |  | 
**UserCount** | **int32** |  | 
**Users** | Pointer to [**[]LimitedUserInstance**](LimitedUserInstance.md) | The users field is present on instances created by the requesting user. | [optional] 
**World** | [**World**](World.md) |  | 
**WorldId** | **string** | WorldID be \&quot;offline\&quot; on User profiles if you are not friends with that user. | 

## Methods

### NewInstance

`func NewInstance(clientNumber string, full bool, id string, instanceId string, location string, nUsers int32, name string, permanent bool, photonRegion Region, platforms InstancePlatforms, queueEnabled bool, queueSize int32, recommendedCapacity int32, region InstanceRegion, secureName string, strict bool, tags []string, type_ InstanceType, userCount int32, world World, worldId string, ) *Instance`

NewInstance instantiates a new Instance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceWithDefaults

`func NewInstanceWithDefaults() *Instance`

NewInstanceWithDefaults instantiates a new Instance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *Instance) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Instance) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Instance) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Instance) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAgeGate

`func (o *Instance) GetAgeGate() bool`

GetAgeGate returns the AgeGate field if non-nil, zero value otherwise.

### GetAgeGateOk

`func (o *Instance) GetAgeGateOk() (*bool, bool)`

GetAgeGateOk returns a tuple with the AgeGate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeGate

`func (o *Instance) SetAgeGate(v bool)`

SetAgeGate sets AgeGate field to given value.

### HasAgeGate

`func (o *Instance) HasAgeGate() bool`

HasAgeGate returns a boolean if a field has been set.

### SetAgeGateNil

`func (o *Instance) SetAgeGateNil(b bool)`

 SetAgeGateNil sets the value for AgeGate to be an explicit nil

### UnsetAgeGate
`func (o *Instance) UnsetAgeGate()`

UnsetAgeGate ensures that no value is present for AgeGate, not even an explicit nil
### GetCalendarEntryId

`func (o *Instance) GetCalendarEntryId() string`

GetCalendarEntryId returns the CalendarEntryId field if non-nil, zero value otherwise.

### GetCalendarEntryIdOk

`func (o *Instance) GetCalendarEntryIdOk() (*string, bool)`

GetCalendarEntryIdOk returns a tuple with the CalendarEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendarEntryId

`func (o *Instance) SetCalendarEntryId(v string)`

SetCalendarEntryId sets CalendarEntryId field to given value.

### HasCalendarEntryId

`func (o *Instance) HasCalendarEntryId() bool`

HasCalendarEntryId returns a boolean if a field has been set.

### SetCalendarEntryIdNil

`func (o *Instance) SetCalendarEntryIdNil(b bool)`

 SetCalendarEntryIdNil sets the value for CalendarEntryId to be an explicit nil

### UnsetCalendarEntryId
`func (o *Instance) UnsetCalendarEntryId()`

UnsetCalendarEntryId ensures that no value is present for CalendarEntryId, not even an explicit nil
### GetCanRequestInvite

`func (o *Instance) GetCanRequestInvite() bool`

GetCanRequestInvite returns the CanRequestInvite field if non-nil, zero value otherwise.

### GetCanRequestInviteOk

`func (o *Instance) GetCanRequestInviteOk() (*bool, bool)`

GetCanRequestInviteOk returns a tuple with the CanRequestInvite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRequestInvite

`func (o *Instance) SetCanRequestInvite(v bool)`

SetCanRequestInvite sets CanRequestInvite field to given value.

### HasCanRequestInvite

`func (o *Instance) HasCanRequestInvite() bool`

HasCanRequestInvite returns a boolean if a field has been set.

### GetCapacity

`func (o *Instance) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *Instance) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *Instance) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *Instance) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetClientNumber

`func (o *Instance) GetClientNumber() string`

GetClientNumber returns the ClientNumber field if non-nil, zero value otherwise.

### GetClientNumberOk

`func (o *Instance) GetClientNumberOk() (*string, bool)`

GetClientNumberOk returns a tuple with the ClientNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNumber

`func (o *Instance) SetClientNumber(v string)`

SetClientNumber sets ClientNumber field to given value.


### GetClosedAt

`func (o *Instance) GetClosedAt() time.Time`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *Instance) GetClosedAtOk() (*time.Time, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *Instance) SetClosedAt(v time.Time)`

SetClosedAt sets ClosedAt field to given value.

### HasClosedAt

`func (o *Instance) HasClosedAt() bool`

HasClosedAt returns a boolean if a field has been set.

### SetClosedAtNil

`func (o *Instance) SetClosedAtNil(b bool)`

 SetClosedAtNil sets the value for ClosedAt to be an explicit nil

### UnsetClosedAt
`func (o *Instance) UnsetClosedAt()`

UnsetClosedAt ensures that no value is present for ClosedAt, not even an explicit nil
### GetContentSettings

`func (o *Instance) GetContentSettings() InstanceContentSettings`

GetContentSettings returns the ContentSettings field if non-nil, zero value otherwise.

### GetContentSettingsOk

`func (o *Instance) GetContentSettingsOk() (*InstanceContentSettings, bool)`

GetContentSettingsOk returns a tuple with the ContentSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSettings

`func (o *Instance) SetContentSettings(v InstanceContentSettings)`

SetContentSettings sets ContentSettings field to given value.

### HasContentSettings

`func (o *Instance) HasContentSettings() bool`

HasContentSettings returns a boolean if a field has been set.

### GetCreatorId

`func (o *Instance) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *Instance) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *Instance) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *Instance) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetDisplayName

`func (o *Instance) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Instance) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Instance) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Instance) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *Instance) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *Instance) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetFriends

`func (o *Instance) GetFriends() string`

GetFriends returns the Friends field if non-nil, zero value otherwise.

### GetFriendsOk

`func (o *Instance) GetFriendsOk() (*string, bool)`

GetFriendsOk returns a tuple with the Friends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriends

`func (o *Instance) SetFriends(v string)`

SetFriends sets Friends field to given value.

### HasFriends

`func (o *Instance) HasFriends() bool`

HasFriends returns a boolean if a field has been set.

### GetFull

`func (o *Instance) GetFull() bool`

GetFull returns the Full field if non-nil, zero value otherwise.

### GetFullOk

`func (o *Instance) GetFullOk() (*bool, bool)`

GetFullOk returns a tuple with the Full field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFull

`func (o *Instance) SetFull(v bool)`

SetFull sets Full field to given value.


### GetGameServerVersion

`func (o *Instance) GetGameServerVersion() int32`

GetGameServerVersion returns the GameServerVersion field if non-nil, zero value otherwise.

### GetGameServerVersionOk

`func (o *Instance) GetGameServerVersionOk() (*int32, bool)`

GetGameServerVersionOk returns a tuple with the GameServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameServerVersion

`func (o *Instance) SetGameServerVersion(v int32)`

SetGameServerVersion sets GameServerVersion field to given value.

### HasGameServerVersion

`func (o *Instance) HasGameServerVersion() bool`

HasGameServerVersion returns a boolean if a field has been set.

### SetGameServerVersionNil

`func (o *Instance) SetGameServerVersionNil(b bool)`

 SetGameServerVersionNil sets the value for GameServerVersion to be an explicit nil

### UnsetGameServerVersion
`func (o *Instance) UnsetGameServerVersion()`

UnsetGameServerVersion ensures that no value is present for GameServerVersion, not even an explicit nil
### GetGroupAccessType

`func (o *Instance) GetGroupAccessType() GroupAccessType`

GetGroupAccessType returns the GroupAccessType field if non-nil, zero value otherwise.

### GetGroupAccessTypeOk

`func (o *Instance) GetGroupAccessTypeOk() (*GroupAccessType, bool)`

GetGroupAccessTypeOk returns a tuple with the GroupAccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAccessType

`func (o *Instance) SetGroupAccessType(v GroupAccessType)`

SetGroupAccessType sets GroupAccessType field to given value.

### HasGroupAccessType

`func (o *Instance) HasGroupAccessType() bool`

HasGroupAccessType returns a boolean if a field has been set.

### GetHardClose

`func (o *Instance) GetHardClose() bool`

GetHardClose returns the HardClose field if non-nil, zero value otherwise.

### GetHardCloseOk

`func (o *Instance) GetHardCloseOk() (*bool, bool)`

GetHardCloseOk returns a tuple with the HardClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardClose

`func (o *Instance) SetHardClose(v bool)`

SetHardClose sets HardClose field to given value.

### HasHardClose

`func (o *Instance) HasHardClose() bool`

HasHardClose returns a boolean if a field has been set.

### SetHardCloseNil

`func (o *Instance) SetHardCloseNil(b bool)`

 SetHardCloseNil sets the value for HardClose to be an explicit nil

### UnsetHardClose
`func (o *Instance) UnsetHardClose()`

UnsetHardClose ensures that no value is present for HardClose, not even an explicit nil
### GetHasCapacityForYou

`func (o *Instance) GetHasCapacityForYou() bool`

GetHasCapacityForYou returns the HasCapacityForYou field if non-nil, zero value otherwise.

### GetHasCapacityForYouOk

`func (o *Instance) GetHasCapacityForYouOk() (*bool, bool)`

GetHasCapacityForYouOk returns a tuple with the HasCapacityForYou field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCapacityForYou

`func (o *Instance) SetHasCapacityForYou(v bool)`

SetHasCapacityForYou sets HasCapacityForYou field to given value.

### HasHasCapacityForYou

`func (o *Instance) HasHasCapacityForYou() bool`

HasHasCapacityForYou returns a boolean if a field has been set.

### GetHidden

`func (o *Instance) GetHidden() string`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *Instance) GetHiddenOk() (*string, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *Instance) SetHidden(v string)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *Instance) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetId

`func (o *Instance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Instance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Instance) SetId(v string)`

SetId sets Id field to given value.


### GetInstanceId

`func (o *Instance) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *Instance) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *Instance) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetInstancePersistenceEnabled

`func (o *Instance) GetInstancePersistenceEnabled() bool`

GetInstancePersistenceEnabled returns the InstancePersistenceEnabled field if non-nil, zero value otherwise.

### GetInstancePersistenceEnabledOk

`func (o *Instance) GetInstancePersistenceEnabledOk() (*bool, bool)`

GetInstancePersistenceEnabledOk returns a tuple with the InstancePersistenceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancePersistenceEnabled

`func (o *Instance) SetInstancePersistenceEnabled(v bool)`

SetInstancePersistenceEnabled sets InstancePersistenceEnabled field to given value.

### HasInstancePersistenceEnabled

`func (o *Instance) HasInstancePersistenceEnabled() bool`

HasInstancePersistenceEnabled returns a boolean if a field has been set.

### SetInstancePersistenceEnabledNil

`func (o *Instance) SetInstancePersistenceEnabledNil(b bool)`

 SetInstancePersistenceEnabledNil sets the value for InstancePersistenceEnabled to be an explicit nil

### UnsetInstancePersistenceEnabled
`func (o *Instance) UnsetInstancePersistenceEnabled()`

UnsetInstancePersistenceEnabled ensures that no value is present for InstancePersistenceEnabled, not even an explicit nil
### GetLocation

`func (o *Instance) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Instance) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Instance) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetNUsers

`func (o *Instance) GetNUsers() int32`

GetNUsers returns the NUsers field if non-nil, zero value otherwise.

### GetNUsersOk

`func (o *Instance) GetNUsersOk() (*int32, bool)`

GetNUsersOk returns a tuple with the NUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNUsers

`func (o *Instance) SetNUsers(v int32)`

SetNUsers sets NUsers field to given value.


### GetName

`func (o *Instance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Instance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Instance) SetName(v string)`

SetName sets Name field to given value.


### GetNonce

`func (o *Instance) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *Instance) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *Instance) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *Instance) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetOwnerId

`func (o *Instance) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Instance) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Instance) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *Instance) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *Instance) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *Instance) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetPermanent

`func (o *Instance) GetPermanent() bool`

GetPermanent returns the Permanent field if non-nil, zero value otherwise.

### GetPermanentOk

`func (o *Instance) GetPermanentOk() (*bool, bool)`

GetPermanentOk returns a tuple with the Permanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanent

`func (o *Instance) SetPermanent(v bool)`

SetPermanent sets Permanent field to given value.


### GetPhotonRegion

`func (o *Instance) GetPhotonRegion() Region`

GetPhotonRegion returns the PhotonRegion field if non-nil, zero value otherwise.

### GetPhotonRegionOk

`func (o *Instance) GetPhotonRegionOk() (*Region, bool)`

GetPhotonRegionOk returns a tuple with the PhotonRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotonRegion

`func (o *Instance) SetPhotonRegion(v Region)`

SetPhotonRegion sets PhotonRegion field to given value.


### GetPlatforms

`func (o *Instance) GetPlatforms() InstancePlatforms`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *Instance) GetPlatformsOk() (*InstancePlatforms, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *Instance) SetPlatforms(v InstancePlatforms)`

SetPlatforms sets Platforms field to given value.


### GetPlayerPersistenceEnabled

`func (o *Instance) GetPlayerPersistenceEnabled() bool`

GetPlayerPersistenceEnabled returns the PlayerPersistenceEnabled field if non-nil, zero value otherwise.

### GetPlayerPersistenceEnabledOk

`func (o *Instance) GetPlayerPersistenceEnabledOk() (*bool, bool)`

GetPlayerPersistenceEnabledOk returns a tuple with the PlayerPersistenceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerPersistenceEnabled

`func (o *Instance) SetPlayerPersistenceEnabled(v bool)`

SetPlayerPersistenceEnabled sets PlayerPersistenceEnabled field to given value.

### HasPlayerPersistenceEnabled

`func (o *Instance) HasPlayerPersistenceEnabled() bool`

HasPlayerPersistenceEnabled returns a boolean if a field has been set.

### SetPlayerPersistenceEnabledNil

`func (o *Instance) SetPlayerPersistenceEnabledNil(b bool)`

 SetPlayerPersistenceEnabledNil sets the value for PlayerPersistenceEnabled to be an explicit nil

### UnsetPlayerPersistenceEnabled
`func (o *Instance) UnsetPlayerPersistenceEnabled()`

UnsetPlayerPersistenceEnabled ensures that no value is present for PlayerPersistenceEnabled, not even an explicit nil
### GetPrivate

`func (o *Instance) GetPrivate() string`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *Instance) GetPrivateOk() (*string, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *Instance) SetPrivate(v string)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *Instance) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetQueueEnabled

`func (o *Instance) GetQueueEnabled() bool`

GetQueueEnabled returns the QueueEnabled field if non-nil, zero value otherwise.

### GetQueueEnabledOk

`func (o *Instance) GetQueueEnabledOk() (*bool, bool)`

GetQueueEnabledOk returns a tuple with the QueueEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueEnabled

`func (o *Instance) SetQueueEnabled(v bool)`

SetQueueEnabled sets QueueEnabled field to given value.


### GetQueueSize

`func (o *Instance) GetQueueSize() int32`

GetQueueSize returns the QueueSize field if non-nil, zero value otherwise.

### GetQueueSizeOk

`func (o *Instance) GetQueueSizeOk() (*int32, bool)`

GetQueueSizeOk returns a tuple with the QueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueSize

`func (o *Instance) SetQueueSize(v int32)`

SetQueueSize sets QueueSize field to given value.


### GetRecommendedCapacity

`func (o *Instance) GetRecommendedCapacity() int32`

GetRecommendedCapacity returns the RecommendedCapacity field if non-nil, zero value otherwise.

### GetRecommendedCapacityOk

`func (o *Instance) GetRecommendedCapacityOk() (*int32, bool)`

GetRecommendedCapacityOk returns a tuple with the RecommendedCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedCapacity

`func (o *Instance) SetRecommendedCapacity(v int32)`

SetRecommendedCapacity sets RecommendedCapacity field to given value.


### GetRegion

`func (o *Instance) GetRegion() InstanceRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Instance) GetRegionOk() (*InstanceRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Instance) SetRegion(v InstanceRegion)`

SetRegion sets Region field to given value.


### GetRoleRestricted

`func (o *Instance) GetRoleRestricted() bool`

GetRoleRestricted returns the RoleRestricted field if non-nil, zero value otherwise.

### GetRoleRestrictedOk

`func (o *Instance) GetRoleRestrictedOk() (*bool, bool)`

GetRoleRestrictedOk returns a tuple with the RoleRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleRestricted

`func (o *Instance) SetRoleRestricted(v bool)`

SetRoleRestricted sets RoleRestricted field to given value.

### HasRoleRestricted

`func (o *Instance) HasRoleRestricted() bool`

HasRoleRestricted returns a boolean if a field has been set.

### GetSecureName

`func (o *Instance) GetSecureName() string`

GetSecureName returns the SecureName field if non-nil, zero value otherwise.

### GetSecureNameOk

`func (o *Instance) GetSecureNameOk() (*string, bool)`

GetSecureNameOk returns a tuple with the SecureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureName

`func (o *Instance) SetSecureName(v string)`

SetSecureName sets SecureName field to given value.


### GetShortName

`func (o *Instance) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *Instance) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *Instance) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *Instance) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### SetShortNameNil

`func (o *Instance) SetShortNameNil(b bool)`

 SetShortNameNil sets the value for ShortName to be an explicit nil

### UnsetShortName
`func (o *Instance) UnsetShortName()`

UnsetShortName ensures that no value is present for ShortName, not even an explicit nil
### GetStrict

`func (o *Instance) GetStrict() bool`

GetStrict returns the Strict field if non-nil, zero value otherwise.

### GetStrictOk

`func (o *Instance) GetStrictOk() (*bool, bool)`

GetStrictOk returns a tuple with the Strict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrict

`func (o *Instance) SetStrict(v bool)`

SetStrict sets Strict field to given value.


### GetTags

`func (o *Instance) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Instance) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Instance) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetType

`func (o *Instance) GetType() InstanceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Instance) GetTypeOk() (*InstanceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Instance) SetType(v InstanceType)`

SetType sets Type field to given value.


### GetUserCount

`func (o *Instance) GetUserCount() int32`

GetUserCount returns the UserCount field if non-nil, zero value otherwise.

### GetUserCountOk

`func (o *Instance) GetUserCountOk() (*int32, bool)`

GetUserCountOk returns a tuple with the UserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCount

`func (o *Instance) SetUserCount(v int32)`

SetUserCount sets UserCount field to given value.


### GetUsers

`func (o *Instance) GetUsers() []LimitedUserInstance`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Instance) GetUsersOk() (*[]LimitedUserInstance, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Instance) SetUsers(v []LimitedUserInstance)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *Instance) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetWorld

`func (o *Instance) GetWorld() World`

GetWorld returns the World field if non-nil, zero value otherwise.

### GetWorldOk

`func (o *Instance) GetWorldOk() (*World, bool)`

GetWorldOk returns a tuple with the World field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorld

`func (o *Instance) SetWorld(v World)`

SetWorld sets World field to given value.


### GetWorldId

`func (o *Instance) GetWorldId() string`

GetWorldId returns the WorldId field if non-nil, zero value otherwise.

### GetWorldIdOk

`func (o *Instance) GetWorldIdOk() (*string, bool)`

GetWorldIdOk returns a tuple with the WorldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorldId

`func (o *Instance) SetWorldId(v string)`

SetWorldId sets WorldId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


