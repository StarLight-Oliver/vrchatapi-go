# \InventoryAPI

All URIs are relative to *https://api.vrchat.cloud/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsumeOwnInventoryItem**](InventoryAPI.md#ConsumeOwnInventoryItem) | **Put** /inventory/{inventoryItemId}/consume | Consume Own Inventory Item
[**DeleteOwnInventoryItem**](InventoryAPI.md#DeleteOwnInventoryItem) | **Delete** /inventory/{inventoryItemId} | Delete Own Inventory Item
[**EquipOwnInventoryItem**](InventoryAPI.md#EquipOwnInventoryItem) | **Put** /inventory/{inventoryItemId}/equip | Equip Own Inventory Item
[**GetInventory**](InventoryAPI.md#GetInventory) | **Get** /inventory | Get Inventory
[**GetInventoryCollections**](InventoryAPI.md#GetInventoryCollections) | **Get** /inventory/collections | List Inventory Collections
[**GetInventoryDrops**](InventoryAPI.md#GetInventoryDrops) | **Get** /inventory/drops | List Inventory Drops
[**GetInventoryTemplate**](InventoryAPI.md#GetInventoryTemplate) | **Get** /inventory/template/{inventoryTemplateId} | Get Inventory Template
[**GetOwnInventoryItem**](InventoryAPI.md#GetOwnInventoryItem) | **Get** /inventory/{inventoryItemId} | Get Own Inventory Item
[**GetUserInventoryItem**](InventoryAPI.md#GetUserInventoryItem) | **Get** /user/{userId}/inventory/{inventoryItemId} | Get User Inventory Item
[**ShareInventoryItemDirect**](InventoryAPI.md#ShareInventoryItemDirect) | **Post** /inventory/cloning/direct | Share Inventory Item Direct
[**ShareInventoryItemPedestal**](InventoryAPI.md#ShareInventoryItemPedestal) | **Get** /inventory/cloning/pedestal | Share Inventory Item by Pedestal
[**SpawnInventoryItem**](InventoryAPI.md#SpawnInventoryItem) | **Get** /inventory/spawn | Spawn Inventory Item
[**UnequipOwnInventorySlot**](InventoryAPI.md#UnequipOwnInventorySlot) | **Delete** /inventory/{inventoryItemId}/equip | Unequip Own Inventory Slot
[**UpdateOwnInventoryItem**](InventoryAPI.md#UpdateOwnInventoryItem) | **Put** /inventory/{inventoryItemId} | Update Own Inventory Item



## ConsumeOwnInventoryItem

> InventoryConsumptionResults ConsumeOwnInventoryItem(ctx, inventoryItemId).Execute()

Consume Own Inventory Item



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	inventoryItemId := "inv_00000000-0000-0000-0000-000000000000" // string | Must be a valid inventory item ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.ConsumeOwnInventoryItem(context.Background(), inventoryItemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.ConsumeOwnInventoryItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsumeOwnInventoryItem`: InventoryConsumptionResults
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.ConsumeOwnInventoryItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inventoryItemId** | **string** | Must be a valid inventory item ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsumeOwnInventoryItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InventoryConsumptionResults**](InventoryConsumptionResults.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOwnInventoryItem

> SuccessFlag DeleteOwnInventoryItem(ctx, inventoryItemId).Execute()

Delete Own Inventory Item



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	inventoryItemId := "inv_00000000-0000-0000-0000-000000000000" // string | Must be a valid inventory item ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.DeleteOwnInventoryItem(context.Background(), inventoryItemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.DeleteOwnInventoryItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOwnInventoryItem`: SuccessFlag
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.DeleteOwnInventoryItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inventoryItemId** | **string** | Must be a valid inventory item ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOwnInventoryItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SuccessFlag**](SuccessFlag.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EquipOwnInventoryItem

> InventoryItem EquipOwnInventoryItem(ctx, inventoryItemId).EquipInventoryItemRequest(equipInventoryItemRequest).Execute()

Equip Own Inventory Item



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	inventoryItemId := "inv_00000000-0000-0000-0000-000000000000" // string | Must be a valid inventory item ID.
	equipInventoryItemRequest := *openapiclient.NewEquipInventoryItemRequest(openapiclient.InventoryEquipSlot("")) // EquipInventoryItemRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.EquipOwnInventoryItem(context.Background(), inventoryItemId).EquipInventoryItemRequest(equipInventoryItemRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.EquipOwnInventoryItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EquipOwnInventoryItem`: InventoryItem
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.EquipOwnInventoryItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inventoryItemId** | **string** | Must be a valid inventory item ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEquipOwnInventoryItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **equipInventoryItemRequest** | [**EquipInventoryItemRequest**](EquipInventoryItemRequest.md) |  | 

### Return type

[**InventoryItem**](InventoryItem.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInventory

> Inventory GetInventory(ctx).N(n).Offset(offset).HolderId(holderId).EquipSlot(equipSlot).Order(order).Tags(tags).Types(types).Flags(flags).NotTypes(notTypes).NotFlags(notFlags).Archived(archived).Execute()

Get Inventory



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
	offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)
	holderId := "holderId_example" // string | The UserID of the owner of the inventory; defaults to the currently authenticated user. (optional)
	equipSlot := openapiclient.InventoryEquipSlot("") // InventoryEquipSlot | Filter for inventory retrieval. (optional) (default to "")
	order := "order_example" // string | Sort order for inventory retrieval. (optional)
	tags := "tags_example" // string | Filter tags for inventory retrieval (comma-separated). (optional)
	types := openapiclient.InventoryItemType("bundle") // InventoryItemType | Filter for inventory retrieval. (optional) (default to "bundle")
	flags := openapiclient.InventoryFlag("archivable") // InventoryFlag | Filter flags for inventory retrieval (comma-separated). (optional) (default to "instantiatable")
	notTypes := openapiclient.InventoryItemType("bundle") // InventoryItemType | Filter out types for inventory retrieval (comma-separated). (optional) (default to "bundle")
	notFlags := openapiclient.InventoryFlag("archivable") // InventoryFlag | Filter out flags for inventory retrieval (comma-separated). (optional) (default to "instantiatable")
	archived := true // bool | Filter archived status for inventory retrieval. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.GetInventory(context.Background()).N(n).Offset(offset).HolderId(holderId).EquipSlot(equipSlot).Order(order).Tags(tags).Types(types).Flags(flags).NotTypes(notTypes).NotFlags(notFlags).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.GetInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInventory`: Inventory
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.GetInventory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 
 **holderId** | **string** | The UserID of the owner of the inventory; defaults to the currently authenticated user. | 
 **equipSlot** | [**InventoryEquipSlot**](InventoryEquipSlot.md) | Filter for inventory retrieval. | [default to &quot;&quot;]
 **order** | **string** | Sort order for inventory retrieval. | 
 **tags** | **string** | Filter tags for inventory retrieval (comma-separated). | 
 **types** | [**InventoryItemType**](InventoryItemType.md) | Filter for inventory retrieval. | [default to &quot;bundle&quot;]
 **flags** | [**InventoryFlag**](InventoryFlag.md) | Filter flags for inventory retrieval (comma-separated). | [default to &quot;instantiatable&quot;]
 **notTypes** | [**InventoryItemType**](InventoryItemType.md) | Filter out types for inventory retrieval (comma-separated). | [default to &quot;bundle&quot;]
 **notFlags** | [**InventoryFlag**](InventoryFlag.md) | Filter out flags for inventory retrieval (comma-separated). | [default to &quot;instantiatable&quot;]
 **archived** | **bool** | Filter archived status for inventory retrieval. | 

### Return type

[**Inventory**](Inventory.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInventoryCollections

> []string GetInventoryCollections(ctx).Execute()

List Inventory Collections



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.GetInventoryCollections(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.GetInventoryCollections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInventoryCollections`: []string
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.GetInventoryCollections`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInventoryCollectionsRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInventoryDrops

> []InventoryDrop GetInventoryDrops(ctx).Active(active).Execute()

List Inventory Drops



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	active := true // bool | Filter for users' listings and inventory bundles. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.GetInventoryDrops(context.Background()).Active(active).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.GetInventoryDrops``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInventoryDrops`: []InventoryDrop
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.GetInventoryDrops`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInventoryDropsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **active** | **bool** | Filter for users&#39; listings and inventory bundles. | 

### Return type

[**[]InventoryDrop**](InventoryDrop.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInventoryTemplate

> InventoryTemplate GetInventoryTemplate(ctx, inventoryTemplateId).Execute()

Get Inventory Template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	inventoryTemplateId := "invt_00000000-0000-0000-0000-000000000000" // string | Must be a valid inventory template ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.GetInventoryTemplate(context.Background(), inventoryTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.GetInventoryTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInventoryTemplate`: InventoryTemplate
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.GetInventoryTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inventoryTemplateId** | **string** | Must be a valid inventory template ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInventoryTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InventoryTemplate**](InventoryTemplate.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOwnInventoryItem

> InventoryItem GetOwnInventoryItem(ctx, inventoryItemId).Execute()

Get Own Inventory Item



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	inventoryItemId := "inv_00000000-0000-0000-0000-000000000000" // string | Must be a valid inventory item ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.GetOwnInventoryItem(context.Background(), inventoryItemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.GetOwnInventoryItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOwnInventoryItem`: InventoryItem
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.GetOwnInventoryItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inventoryItemId** | **string** | Must be a valid inventory item ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOwnInventoryItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InventoryItem**](InventoryItem.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserInventoryItem

> InventoryItem GetUserInventoryItem(ctx, userId, inventoryItemId).Execute()

Get User Inventory Item



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	userId := "userId_example" // string | Must be a valid user ID.
	inventoryItemId := "inv_00000000-0000-0000-0000-000000000000" // string | Must be a valid inventory item ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.GetUserInventoryItem(context.Background(), userId, inventoryItemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.GetUserInventoryItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserInventoryItem`: InventoryItem
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.GetUserInventoryItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 
**inventoryItemId** | **string** | Must be a valid inventory item ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserInventoryItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InventoryItem**](InventoryItem.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShareInventoryItemDirect

> OkStatus ShareInventoryItemDirect(ctx).ItemId(itemId).Duration(duration).ShareInventoryItemDirectRequest(shareInventoryItemDirectRequest).Execute()

Share Inventory Item Direct



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	itemId := "itemId_example" // string | Id for inventory item sharing.
	duration := int32(56) // int32 | The duration before the sharing pedestal despawns. (default to 90)
	shareInventoryItemDirectRequest := *openapiclient.NewShareInventoryItemDirectRequest("inv_10bce5b0-2d2b-44e0-900d-db6534615162", []string{"usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469"}) // ShareInventoryItemDirectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.ShareInventoryItemDirect(context.Background()).ItemId(itemId).Duration(duration).ShareInventoryItemDirectRequest(shareInventoryItemDirectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.ShareInventoryItemDirect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShareInventoryItemDirect`: OkStatus
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.ShareInventoryItemDirect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShareInventoryItemDirectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemId** | **string** | Id for inventory item sharing. | 
 **duration** | **int32** | The duration before the sharing pedestal despawns. | [default to 90]
 **shareInventoryItemDirectRequest** | [**ShareInventoryItemDirectRequest**](ShareInventoryItemDirectRequest.md) |  | 

### Return type

[**OkStatus**](OkStatus.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShareInventoryItemPedestal

> InventorySpawn ShareInventoryItemPedestal(ctx).ItemId(itemId).Duration(duration).Execute()

Share Inventory Item by Pedestal



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	itemId := "itemId_example" // string | Id for inventory item sharing.
	duration := int32(56) // int32 | The duration before the sharing pedestal despawns. (default to 90)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.ShareInventoryItemPedestal(context.Background()).ItemId(itemId).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.ShareInventoryItemPedestal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShareInventoryItemPedestal`: InventorySpawn
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.ShareInventoryItemPedestal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShareInventoryItemPedestalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemId** | **string** | Id for inventory item sharing. | 
 **duration** | **int32** | The duration before the sharing pedestal despawns. | [default to 90]

### Return type

[**InventorySpawn**](InventorySpawn.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpawnInventoryItem

> InventorySpawn SpawnInventoryItem(ctx).Id(id).Execute()

Spawn Inventory Item



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	id := "id_example" // string | Id for inventory item spawning.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.SpawnInventoryItem(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.SpawnInventoryItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpawnInventoryItem`: InventorySpawn
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.SpawnInventoryItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpawnInventoryItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Id for inventory item spawning. | 

### Return type

[**InventorySpawn**](InventorySpawn.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnequipOwnInventorySlot

> string UnequipOwnInventorySlot(ctx, inventoryItemId).Execute()

Unequip Own Inventory Slot



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	inventoryItemId := openapiclient.InventoryEquipSlot("") // InventoryEquipSlot | Selector for inventory slot management. (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.UnequipOwnInventorySlot(context.Background(), inventoryItemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.UnequipOwnInventorySlot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnequipOwnInventorySlot`: string
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.UnequipOwnInventorySlot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inventoryItemId** | [**InventoryEquipSlot**](.md) | Selector for inventory slot management. | [default to &quot;&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiUnequipOwnInventorySlotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOwnInventoryItem

> InventoryItem UpdateOwnInventoryItem(ctx, inventoryItemId).UpdateInventoryItemRequest(updateInventoryItemRequest).Execute()

Update Own Inventory Item



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	inventoryItemId := "inv_00000000-0000-0000-0000-000000000000" // string | Must be a valid inventory item ID.
	updateInventoryItemRequest := *openapiclient.NewUpdateInventoryItemRequest() // UpdateInventoryItemRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.UpdateOwnInventoryItem(context.Background(), inventoryItemId).UpdateInventoryItemRequest(updateInventoryItemRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.UpdateOwnInventoryItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOwnInventoryItem`: InventoryItem
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.UpdateOwnInventoryItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inventoryItemId** | **string** | Must be a valid inventory item ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOwnInventoryItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateInventoryItemRequest** | [**UpdateInventoryItemRequest**](UpdateInventoryItemRequest.md) |  | 

### Return type

[**InventoryItem**](InventoryItem.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

