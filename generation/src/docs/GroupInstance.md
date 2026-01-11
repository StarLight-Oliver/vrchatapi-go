# GroupInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | **string** | InstanceID can be \&quot;offline\&quot; on User profiles if you are not friends with that user and \&quot;private\&quot; if you are friends and user is in private instance. | 
**Location** | **string** | Represents a unique location, consisting of a world identifier and an instance identifier, or \&quot;offline\&quot; if the user is not on your friends list. | 
**MemberCount** | **int32** |  | 
**World** | [**World**](World.md) |  | 

## Methods

### NewGroupInstance

`func NewGroupInstance(instanceId string, location string, memberCount int32, world World, ) *GroupInstance`

NewGroupInstance instantiates a new GroupInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupInstanceWithDefaults

`func NewGroupInstanceWithDefaults() *GroupInstance`

NewGroupInstanceWithDefaults instantiates a new GroupInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *GroupInstance) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *GroupInstance) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *GroupInstance) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetLocation

`func (o *GroupInstance) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GroupInstance) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GroupInstance) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetMemberCount

`func (o *GroupInstance) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *GroupInstance) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *GroupInstance) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.


### GetWorld

`func (o *GroupInstance) GetWorld() World`

GetWorld returns the World field if non-nil, zero value otherwise.

### GetWorldOk

`func (o *GroupInstance) GetWorldOk() (*World, bool)`

GetWorldOk returns a tuple with the World field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorld

`func (o *GroupInstance) SetWorld(v World)`

SetWorld sets World field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


