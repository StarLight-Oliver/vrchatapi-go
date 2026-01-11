# GroupRoleTemplateValuesRoles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**BasePermissions** | Pointer to [**[]GroupPermissions**](GroupPermissions.md) |  | [optional] 
**IsAddedOnJoin** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewGroupRoleTemplateValuesRoles

`func NewGroupRoleTemplateValuesRoles() *GroupRoleTemplateValuesRoles`

NewGroupRoleTemplateValuesRoles instantiates a new GroupRoleTemplateValuesRoles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupRoleTemplateValuesRolesWithDefaults

`func NewGroupRoleTemplateValuesRolesWithDefaults() *GroupRoleTemplateValuesRoles`

NewGroupRoleTemplateValuesRolesWithDefaults instantiates a new GroupRoleTemplateValuesRoles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GroupRoleTemplateValuesRoles) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupRoleTemplateValuesRoles) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupRoleTemplateValuesRoles) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupRoleTemplateValuesRoles) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GroupRoleTemplateValuesRoles) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupRoleTemplateValuesRoles) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupRoleTemplateValuesRoles) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupRoleTemplateValuesRoles) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBasePermissions

`func (o *GroupRoleTemplateValuesRoles) GetBasePermissions() []GroupPermissions`

GetBasePermissions returns the BasePermissions field if non-nil, zero value otherwise.

### GetBasePermissionsOk

`func (o *GroupRoleTemplateValuesRoles) GetBasePermissionsOk() (*[]GroupPermissions, bool)`

GetBasePermissionsOk returns a tuple with the BasePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePermissions

`func (o *GroupRoleTemplateValuesRoles) SetBasePermissions(v []GroupPermissions)`

SetBasePermissions sets BasePermissions field to given value.

### HasBasePermissions

`func (o *GroupRoleTemplateValuesRoles) HasBasePermissions() bool`

HasBasePermissions returns a boolean if a field has been set.

### GetIsAddedOnJoin

`func (o *GroupRoleTemplateValuesRoles) GetIsAddedOnJoin() bool`

GetIsAddedOnJoin returns the IsAddedOnJoin field if non-nil, zero value otherwise.

### GetIsAddedOnJoinOk

`func (o *GroupRoleTemplateValuesRoles) GetIsAddedOnJoinOk() (*bool, bool)`

GetIsAddedOnJoinOk returns a tuple with the IsAddedOnJoin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAddedOnJoin

`func (o *GroupRoleTemplateValuesRoles) SetIsAddedOnJoin(v bool)`

SetIsAddedOnJoin sets IsAddedOnJoin field to given value.

### HasIsAddedOnJoin

`func (o *GroupRoleTemplateValuesRoles) HasIsAddedOnJoin() bool`

HasIsAddedOnJoin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


