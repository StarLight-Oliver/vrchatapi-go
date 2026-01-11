# GroupRoleTemplateValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasePermissions** | [**[]GroupPermissions**](GroupPermissions.md) |  | 
**Description** | **string** |  | 
**Name** | **string** |  | 
**Roles** | [**GroupRoleTemplateValuesRoles**](GroupRoleTemplateValuesRoles.md) |  | 

## Methods

### NewGroupRoleTemplateValues

`func NewGroupRoleTemplateValues(basePermissions []GroupPermissions, description string, name string, roles GroupRoleTemplateValuesRoles, ) *GroupRoleTemplateValues`

NewGroupRoleTemplateValues instantiates a new GroupRoleTemplateValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupRoleTemplateValuesWithDefaults

`func NewGroupRoleTemplateValuesWithDefaults() *GroupRoleTemplateValues`

NewGroupRoleTemplateValuesWithDefaults instantiates a new GroupRoleTemplateValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasePermissions

`func (o *GroupRoleTemplateValues) GetBasePermissions() []GroupPermissions`

GetBasePermissions returns the BasePermissions field if non-nil, zero value otherwise.

### GetBasePermissionsOk

`func (o *GroupRoleTemplateValues) GetBasePermissionsOk() (*[]GroupPermissions, bool)`

GetBasePermissionsOk returns a tuple with the BasePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePermissions

`func (o *GroupRoleTemplateValues) SetBasePermissions(v []GroupPermissions)`

SetBasePermissions sets BasePermissions field to given value.


### GetDescription

`func (o *GroupRoleTemplateValues) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupRoleTemplateValues) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupRoleTemplateValues) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *GroupRoleTemplateValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupRoleTemplateValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupRoleTemplateValues) SetName(v string)`

SetName sets Name field to given value.


### GetRoles

`func (o *GroupRoleTemplateValues) GetRoles() GroupRoleTemplateValuesRoles`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GroupRoleTemplateValues) GetRolesOk() (*GroupRoleTemplateValuesRoles, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GroupRoleTemplateValues) SetRoles(v GroupRoleTemplateValuesRoles)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


