# Permission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**PermissionData**](PermissionData.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**Name** | **string** |  | 
**OwnerDisplayName** | **string** |  | 
**OwnerId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewPermission

`func NewPermission(id string, name string, ownerDisplayName string, ownerId string, ) *Permission`

NewPermission instantiates a new Permission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionWithDefaults

`func NewPermissionWithDefaults() *Permission`

NewPermissionWithDefaults instantiates a new Permission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *Permission) GetData() PermissionData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Permission) GetDataOk() (*PermissionData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Permission) SetData(v PermissionData)`

SetData sets Data field to given value.

### HasData

`func (o *Permission) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDescription

`func (o *Permission) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Permission) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Permission) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Permission) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *Permission) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Permission) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Permission) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Permission) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *Permission) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Permission) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Permission) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Permission) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Permission) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Permission) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerDisplayName

`func (o *Permission) GetOwnerDisplayName() string`

GetOwnerDisplayName returns the OwnerDisplayName field if non-nil, zero value otherwise.

### GetOwnerDisplayNameOk

`func (o *Permission) GetOwnerDisplayNameOk() (*string, bool)`

GetOwnerDisplayNameOk returns a tuple with the OwnerDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerDisplayName

`func (o *Permission) SetOwnerDisplayName(v string)`

SetOwnerDisplayName sets OwnerDisplayName field to given value.


### GetOwnerId

`func (o *Permission) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Permission) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Permission) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetType

`func (o *Permission) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Permission) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Permission) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Permission) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


