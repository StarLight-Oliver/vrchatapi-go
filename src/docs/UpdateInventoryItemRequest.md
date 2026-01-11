# UpdateInventoryItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsArchived** | Pointer to **bool** |  | [optional] 
**IsSeen** | Pointer to **bool** |  | [optional] 
**UserAttributes** | Pointer to [**InventoryUserAttributes**](InventoryUserAttributes.md) |  | [optional] 

## Methods

### NewUpdateInventoryItemRequest

`func NewUpdateInventoryItemRequest() *UpdateInventoryItemRequest`

NewUpdateInventoryItemRequest instantiates a new UpdateInventoryItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInventoryItemRequestWithDefaults

`func NewUpdateInventoryItemRequestWithDefaults() *UpdateInventoryItemRequest`

NewUpdateInventoryItemRequestWithDefaults instantiates a new UpdateInventoryItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsArchived

`func (o *UpdateInventoryItemRequest) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *UpdateInventoryItemRequest) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *UpdateInventoryItemRequest) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *UpdateInventoryItemRequest) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### GetIsSeen

`func (o *UpdateInventoryItemRequest) GetIsSeen() bool`

GetIsSeen returns the IsSeen field if non-nil, zero value otherwise.

### GetIsSeenOk

`func (o *UpdateInventoryItemRequest) GetIsSeenOk() (*bool, bool)`

GetIsSeenOk returns a tuple with the IsSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSeen

`func (o *UpdateInventoryItemRequest) SetIsSeen(v bool)`

SetIsSeen sets IsSeen field to given value.

### HasIsSeen

`func (o *UpdateInventoryItemRequest) HasIsSeen() bool`

HasIsSeen returns a boolean if a field has been set.

### GetUserAttributes

`func (o *UpdateInventoryItemRequest) GetUserAttributes() InventoryUserAttributes`

GetUserAttributes returns the UserAttributes field if non-nil, zero value otherwise.

### GetUserAttributesOk

`func (o *UpdateInventoryItemRequest) GetUserAttributesOk() (*InventoryUserAttributes, bool)`

GetUserAttributesOk returns a tuple with the UserAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttributes

`func (o *UpdateInventoryItemRequest) SetUserAttributes(v InventoryUserAttributes)`

SetUserAttributes sets UserAttributes field to given value.

### HasUserAttributes

`func (o *UpdateInventoryItemRequest) HasUserAttributes() bool`

HasUserAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


