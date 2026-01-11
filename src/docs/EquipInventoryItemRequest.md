# EquipInventoryItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EquipSlot** | [**InventoryEquipSlot**](InventoryEquipSlot.md) |  | [default to EMPTY]

## Methods

### NewEquipInventoryItemRequest

`func NewEquipInventoryItemRequest(equipSlot InventoryEquipSlot, ) *EquipInventoryItemRequest`

NewEquipInventoryItemRequest instantiates a new EquipInventoryItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipInventoryItemRequestWithDefaults

`func NewEquipInventoryItemRequestWithDefaults() *EquipInventoryItemRequest`

NewEquipInventoryItemRequestWithDefaults instantiates a new EquipInventoryItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEquipSlot

`func (o *EquipInventoryItemRequest) GetEquipSlot() InventoryEquipSlot`

GetEquipSlot returns the EquipSlot field if non-nil, zero value otherwise.

### GetEquipSlotOk

`func (o *EquipInventoryItemRequest) GetEquipSlotOk() (*InventoryEquipSlot, bool)`

GetEquipSlotOk returns a tuple with the EquipSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipSlot

`func (o *EquipInventoryItemRequest) SetEquipSlot(v InventoryEquipSlot)`

SetEquipSlot sets EquipSlot field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


