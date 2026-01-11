# InventoryConsumptionResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | **[]map[string]interface{}** |  | 
**InventoryItems** | [**[]InventoryItem**](InventoryItem.md) |  | 
**InventoryItemsCreated** | **int32** |  | 

## Methods

### NewInventoryConsumptionResults

`func NewInventoryConsumptionResults(errors []map[string]interface{}, inventoryItems []InventoryItem, inventoryItemsCreated int32, ) *InventoryConsumptionResults`

NewInventoryConsumptionResults instantiates a new InventoryConsumptionResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryConsumptionResultsWithDefaults

`func NewInventoryConsumptionResultsWithDefaults() *InventoryConsumptionResults`

NewInventoryConsumptionResultsWithDefaults instantiates a new InventoryConsumptionResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *InventoryConsumptionResults) GetErrors() []map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InventoryConsumptionResults) GetErrorsOk() (*[]map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InventoryConsumptionResults) SetErrors(v []map[string]interface{})`

SetErrors sets Errors field to given value.


### GetInventoryItems

`func (o *InventoryConsumptionResults) GetInventoryItems() []InventoryItem`

GetInventoryItems returns the InventoryItems field if non-nil, zero value otherwise.

### GetInventoryItemsOk

`func (o *InventoryConsumptionResults) GetInventoryItemsOk() (*[]InventoryItem, bool)`

GetInventoryItemsOk returns a tuple with the InventoryItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryItems

`func (o *InventoryConsumptionResults) SetInventoryItems(v []InventoryItem)`

SetInventoryItems sets InventoryItems field to given value.


### GetInventoryItemsCreated

`func (o *InventoryConsumptionResults) GetInventoryItemsCreated() int32`

GetInventoryItemsCreated returns the InventoryItemsCreated field if non-nil, zero value otherwise.

### GetInventoryItemsCreatedOk

`func (o *InventoryConsumptionResults) GetInventoryItemsCreatedOk() (*int32, bool)`

GetInventoryItemsCreatedOk returns a tuple with the InventoryItemsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryItemsCreated

`func (o *InventoryConsumptionResults) SetInventoryItemsCreated(v int32)`

SetInventoryItemsCreated sets InventoryItemsCreated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


