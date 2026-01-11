# \InstancesAPI

All URIs are relative to *https://api.vrchat.cloud/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloseInstance**](InstancesAPI.md#CloseInstance) | **Delete** /instances/{worldId}:{instanceId} | Close Instance
[**CreateInstance**](InstancesAPI.md#CreateInstance) | **Post** /instances | Create Instance
[**GetInstance**](InstancesAPI.md#GetInstance) | **Get** /instances/{worldId}:{instanceId} | Get Instance
[**GetInstanceByShortName**](InstancesAPI.md#GetInstanceByShortName) | **Get** /instances/s/{shortName} | Get Instance By Short Name
[**GetRecentLocations**](InstancesAPI.md#GetRecentLocations) | **Get** /instances/recent | List Recent Locations
[**GetShortName**](InstancesAPI.md#GetShortName) | **Get** /instances/{worldId}:{instanceId}/shortName | Get Instance Short Name



## CloseInstance

> Instance CloseInstance(ctx, worldId, instanceId).HardClose(hardClose).ClosedAt(closedAt).Execute()

Close Instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/StarLight-Oliver/vrchatapi-go"
)

func main() {
	worldId := "worldId_example" // string | Must be a valid world ID.
	instanceId := "instanceId_example" // string | Must be a valid instance ID.
	hardClose := true // bool | Whether to hard close the instance. Defaults to false. (optional)
	closedAt := time.Now() // time.Time | The time after which users won't be allowed to join the instances. If omitted, the instance will be closed immediately. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.CloseInstance(context.Background(), worldId, instanceId).HardClose(hardClose).ClosedAt(closedAt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.CloseInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloseInstance`: Instance
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.CloseInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**worldId** | **string** | Must be a valid world ID. | 
**instanceId** | **string** | Must be a valid instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloseInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hardClose** | **bool** | Whether to hard close the instance. Defaults to false. | 
 **closedAt** | **time.Time** | The time after which users won&#39;t be allowed to join the instances. If omitted, the instance will be closed immediately. | 

### Return type

[**Instance**](Instance.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInstance

> Instance CreateInstance(ctx).CreateInstanceRequest(createInstanceRequest).Execute()

Create Instance



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
	createInstanceRequest := *openapiclient.NewCreateInstanceRequest(openapiclient.InstanceRegion("eu"), openapiclient.InstanceType("friends"), "wrld_4432ea9b-729c-46e3-8eaf-846aa0a37fdd") // CreateInstanceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.CreateInstance(context.Background()).CreateInstanceRequest(createInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.CreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInstance`: Instance
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.CreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createInstanceRequest** | [**CreateInstanceRequest**](CreateInstanceRequest.md) |  | 

### Return type

[**Instance**](Instance.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstance

> Instance GetInstance(ctx, worldId, instanceId).Execute()

Get Instance



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
	worldId := "worldId_example" // string | Must be a valid world ID.
	instanceId := "instanceId_example" // string | Must be a valid instance ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.GetInstance(context.Background(), worldId, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.GetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstance`: Instance
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.GetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**worldId** | **string** | Must be a valid world ID. | 
**instanceId** | **string** | Must be a valid instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Instance**](Instance.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstanceByShortName

> Instance GetInstanceByShortName(ctx, shortName).Execute()

Get Instance By Short Name



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
	shortName := "shortName_example" // string | Must be a valid instance short name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.GetInstanceByShortName(context.Background(), shortName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.GetInstanceByShortName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstanceByShortName`: Instance
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.GetInstanceByShortName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shortName** | **string** | Must be a valid instance short name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceByShortNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Instance**](Instance.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecentLocations

> []string GetRecentLocations(ctx).N(n).Offset(offset).Execute()

List Recent Locations



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.GetRecentLocations(context.Background()).N(n).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.GetRecentLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecentLocations`: []string
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.GetRecentLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecentLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 

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


## GetShortName

> InstanceShortNameResponse GetShortName(ctx, worldId, instanceId).Execute()

Get Instance Short Name



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
	worldId := "worldId_example" // string | Must be a valid world ID.
	instanceId := "instanceId_example" // string | Must be a valid instance ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.GetShortName(context.Background(), worldId, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.GetShortName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShortName`: InstanceShortNameResponse
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.GetShortName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**worldId** | **string** | Must be a valid world ID. | 
**instanceId** | **string** | Must be a valid instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShortNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InstanceShortNameResponse**](InstanceShortNameResponse.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

