# \WorldsAPI

All URIs are relative to *https://api.vrchat.cloud/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckUserPersistenceExists**](WorldsAPI.md#CheckUserPersistenceExists) | **Get** /users/{userId}/{worldId}/persist/exists | Check User Persistence Exists
[**CreateWorld**](WorldsAPI.md#CreateWorld) | **Post** /worlds | Create World
[**DeleteAllUserPersistenceData**](WorldsAPI.md#DeleteAllUserPersistenceData) | **Delete** /users/{userId}/persist | Delete All User Persistence Data
[**DeleteUserPersistence**](WorldsAPI.md#DeleteUserPersistence) | **Delete** /users/{userId}/{worldId}/persist | Delete User Persistence
[**DeleteWorld**](WorldsAPI.md#DeleteWorld) | **Delete** /worlds/{worldId} | Delete World
[**GetActiveWorlds**](WorldsAPI.md#GetActiveWorlds) | **Get** /worlds/active | List Active Worlds
[**GetFavoritedWorlds**](WorldsAPI.md#GetFavoritedWorlds) | **Get** /worlds/favorites | List Favorited Worlds
[**GetRecentWorlds**](WorldsAPI.md#GetRecentWorlds) | **Get** /worlds/recent | List Recent Worlds
[**GetWorld**](WorldsAPI.md#GetWorld) | **Get** /worlds/{worldId} | Get World by ID
[**GetWorldInstance**](WorldsAPI.md#GetWorldInstance) | **Get** /worlds/{worldId}/{instanceId} | Get World Instance
[**GetWorldMetadata**](WorldsAPI.md#GetWorldMetadata) | **Get** /worlds/{worldId}/metadata | Get World Metadata
[**GetWorldPublishStatus**](WorldsAPI.md#GetWorldPublishStatus) | **Get** /worlds/{worldId}/publish | Get World Publish Status
[**PublishWorld**](WorldsAPI.md#PublishWorld) | **Put** /worlds/{worldId}/publish | Publish World
[**SearchWorlds**](WorldsAPI.md#SearchWorlds) | **Get** /worlds | Search All Worlds
[**UnpublishWorld**](WorldsAPI.md#UnpublishWorld) | **Delete** /worlds/{worldId}/publish | Unpublish World
[**UpdateWorld**](WorldsAPI.md#UpdateWorld) | **Put** /worlds/{worldId} | Update World



## CheckUserPersistenceExists

> CheckUserPersistenceExists(ctx, userId, worldId).Execute()

Check User Persistence Exists



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
	worldId := "worldId_example" // string | Must be a valid world ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorldsAPI.CheckUserPersistenceExists(context.Background(), userId, worldId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorldsAPI.CheckUserPersistenceExists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 
**worldId** | **string** | Must be a valid world ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckUserPersistenceExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorld

> World CreateWorld(ctx).CreateWorldRequest(createWorldRequest).Execute()

Create World



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
	createWorldRequest := *openapiclient.NewCreateWorldRequest("AssetUrl_example", "ImageUrl_example", "Name_example") // CreateWorldRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorldsAPI.CreateWorld(context.Background()).CreateWorldRequest(createWorldRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorldsAPI.CreateWorld``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorld`: World
	fmt.Fprintf(os.Stdout, "Response from `WorldsAPI.CreateWorld`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWorldRequest** | [**CreateWorldRequest**](CreateWorldRequest.md) |  | 

### Return type

[**World**](World.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllUserPersistenceData

> DeleteAllUserPersistenceData(ctx, userId).Execute()

Delete All User Persistence Data



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorldsAPI.DeleteAllUserPersistenceData(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorldsAPI.DeleteAllUserPersistenceData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllUserPersistenceDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserPersistence

> DeleteUserPersistence(ctx, userId, worldId).Execute()

Delete User Persistence



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
	worldId := "worldId_example" // string | Must be a valid world ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorldsAPI.DeleteUserPersistence(context.Background(), userId, worldId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorldsAPI.DeleteUserPersistence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 
**worldId** | **string** | Must be a valid world ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserPersistenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorld

> DeleteWorld(ctx, worldId).Execute()

Delete World



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorldsAPI.DeleteWorld(context.Background(), worldId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorldsAPI.DeleteWorld``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**worldId** | **string** | Must be a valid world ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActiveWorlds

> []LimitedWorld GetActiveWorlds(ctx).Featured(featured).Sort(sort).N(n).Order(order).Offset(offset).Search(search).Tag(tag).Notag(notag).ReleaseStatus(releaseStatus).MaxUnityVersion(maxUnityVersion).MinUnityVersion(minUnityVersion).Platform(platform).Noplatform(noplatform).Execute()

List Active Worlds



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
	featured := true // bool | Filters on featured results. (optional)
	sort := openapiclient.SortOption("_created_at") // SortOption | The sort order of the results. (optional) (default to "popularity")
	n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
	order := openapiclient.OrderOption("ascending") // OrderOption | Result ordering (optional) (default to "descending")
	offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)
	search := "search_example" // string | Filters by world name. (optional)
	tag := "tag_example" // string | Tags to include (comma-separated). Any of the tags needs to be present. (optional)
	notag := "notag_example" // string | Tags to exclude (comma-separated). (optional)
	releaseStatus := openapiclient.ReleaseStatus("all") // ReleaseStatus | Filter by ReleaseStatus. (optional) (default to "public")
	maxUnityVersion := "maxUnityVersion_example" // string | The maximum Unity version supported by the asset. (optional)
	minUnityVersion := "minUnityVersion_example" // string | The minimum Unity version supported by the asset. (optional)
	platform := "platform_example" // string | The platform the asset supports. (optional)
	noplatform := "noplatform_example" // string | The platform the asset does not support. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorldsAPI.GetActiveWorlds(context.Background()).Featured(featured).Sort(sort).N(n).Order(order).Offset(offset).Search(search).Tag(tag).Notag(notag).ReleaseStatus(releaseStatus).MaxUnityVersion(maxUnityVersion).MinUnityVersion(minUnityVersion).Platform(platform).Noplatform(noplatform).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorldsAPI.GetActiveWorlds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActiveWorlds`: []LimitedWorld
	fmt.Fprintf(os.Stdout, "Response from `WorldsAPI.GetActiveWorlds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveWorldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **featured** | **bool** | Filters on featured results. | 
 **sort** | [**SortOption**](SortOption.md) | The sort order of the results. | [default to &quot;popularity&quot;]
 **n** | **int32** | The number of objects to return. | [default to 60]
 **order** | [**OrderOption**](OrderOption.md) | Result ordering | [default to &quot;descending&quot;]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 
 **search** | **string** | Filters by world name. | 
 **tag** | **string** | Tags to include (comma-separated). Any of the tags needs to be present. | 
 **notag** | **string** | Tags to exclude (comma-separated). | 
 **releaseStatus** | [**ReleaseStatus**](ReleaseStatus.md) | Filter by ReleaseStatus. | [default to &quot;public&quot;]
 **maxUnityVersion** | **string** | The maximum Unity version supported by the asset. | 
 **minUnityVersion** | **string** | The minimum Unity version supported by the asset. | 
 **platform** | **string** | The platform the asset supports. | 
 **noplatform** | **string** | The platform the asset does not support. | 

### Return type

[**[]LimitedWorld**](LimitedWorld.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFavoritedWorlds

> []FavoritedWorld GetFavoritedWorlds(ctx).Featured(featured).Sort(sort).N(n).Order(order).Offset(offset).Search(search).Tag(tag).Notag(notag).ReleaseStatus(releaseStatus).MaxUnityVersion(maxUnityVersion).MinUnityVersion(minUnityVersion).Platform(platform).UserId(userId).Execute()

List Favorited Worlds



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
	featured := true // bool | Filters on featured results. (optional)
	sort := openapiclient.SortOption("_created_at") // SortOption | The sort order of the results. (optional) (default to "popularity")
	n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
	order := openapiclient.OrderOption("ascending") // OrderOption | Result ordering (optional) (default to "descending")
	offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)
	search := "search_example" // string | Filters by world name. (optional)
	tag := "tag_example" // string | Tags to include (comma-separated). Any of the tags needs to be present. (optional)
	notag := "notag_example" // string | Tags to exclude (comma-separated). (optional)
	releaseStatus := openapiclient.ReleaseStatus("all") // ReleaseStatus | Filter by ReleaseStatus. (optional) (default to "public")
	maxUnityVersion := "maxUnityVersion_example" // string | The maximum Unity version supported by the asset. (optional)
	minUnityVersion := "minUnityVersion_example" // string | The minimum Unity version supported by the asset. (optional)
	platform := "platform_example" // string | The platform the asset supports. (optional)
	userId := "userId_example" // string | Target user to see information on, admin-only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorldsAPI.GetFavoritedWorlds(context.Background()).Featured(featured).Sort(sort).N(n).Order(order).Offset(offset).Search(search).Tag(tag).Notag(notag).ReleaseStatus(releaseStatus).MaxUnityVersion(maxUnityVersion).MinUnityVersion(minUnityVersion).Platform(platform).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorldsAPI.GetFavoritedWorlds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFavoritedWorlds`: []FavoritedWorld
	fmt.Fprintf(os.Stdout, "Response from `WorldsAPI.GetFavoritedWorlds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFavoritedWorldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **featured** | **bool** | Filters on featured results. | 
 **sort** | [**SortOption**](SortOption.md) | The sort order of the results. | [default to &quot;popularity&quot;]
 **n** | **int32** | The number of objects to return. | [default to 60]
 **order** | [**OrderOption**](OrderOption.md) | Result ordering | [default to &quot;descending&quot;]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 
 **search** | **string** | Filters by world name. | 
 **tag** | **string** | Tags to include (comma-separated). Any of the tags needs to be present. | 
 **notag** | **string** | Tags to exclude (comma-separated). | 
 **releaseStatus** | [**ReleaseStatus**](ReleaseStatus.md) | Filter by ReleaseStatus. | [default to &quot;public&quot;]
 **maxUnityVersion** | **string** | The maximum Unity version supported by the asset. | 
 **minUnityVersion** | **string** | The minimum Unity version supported by the asset. | 
 **platform** | **string** | The platform the asset supports. | 
 **userId** | **string** | Target user to see information on, admin-only. | 

### Return type

[**[]FavoritedWorld**](FavoritedWorld.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecentWorlds

> []LimitedWorld GetRecentWorlds(ctx).Featured(featured).Sort(sort).N(n).Order(order).Offset(offset).Search(search).Tag(tag).Notag(notag).ReleaseStatus(releaseStatus).MaxUnityVersion(maxUnityVersion).MinUnityVersion(minUnityVersion).Platform(platform).UserId(userId).Execute()

List Recent Worlds



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
	featured := true // bool | Filters on featured results. (optional)
	sort := openapiclient.SortOption("_created_at") // SortOption | The sort order of the results. (optional) (default to "popularity")
	n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
	order := openapiclient.OrderOption("ascending") // OrderOption | Result ordering (optional) (default to "descending")
	offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)
	search := "search_example" // string | Filters by world name. (optional)
	tag := "tag_example" // string | Tags to include (comma-separated). Any of the tags needs to be present. (optional)
	notag := "notag_example" // string | Tags to exclude (comma-separated). (optional)
	releaseStatus := openapiclient.ReleaseStatus("all") // ReleaseStatus | Filter by ReleaseStatus. (optional) (default to "public")
	maxUnityVersion := "maxUnityVersion_example" // string | The maximum Unity version supported by the asset. (optional)
	minUnityVersion := "minUnityVersion_example" // string | The minimum Unity version supported by the asset. (optional)
	platform := "platform_example" // string | The platform the asset supports. (optional)
	userId := "userId_example" // string | Target user to see information on, admin-only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorldsAPI.GetRecentWorlds(context.Background()).Featured(featured).Sort(sort).N(n).Order(order).Offset(offset).Search(search).Tag(tag).Notag(notag).ReleaseStatus(releaseStatus).MaxUnityVersion(maxUnityVersion).MinUnityVersion(minUnityVersion).Platform(platform).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorldsAPI.GetRecentWorlds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecentWorlds`: []LimitedWorld
	fmt.Fprintf(os.Stdout, "Response from `WorldsAPI.GetRecentWorlds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecentWorldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **featured** | **bool** | Filters on featured results. | 
 **sort** | [**SortOption**](SortOption.md) | The sort order of the results. | [default to &quot;popularity&quot;]
 **n** | **int32** | The number of objects to return. | [default to 60]
 **order** | [**OrderOption**](OrderOption.md) | Result ordering | [default to &quot;descending&quot;]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 
 **search** | **string** | Filters by world name. | 
 **tag** | **string** | Tags to include (comma-separated). Any of the tags needs to be present. | 
 **notag** | **string** | Tags to exclude (comma-separated). | 
 **releaseStatus** | [**ReleaseStatus**](ReleaseStatus.md) | Filter by ReleaseStatus. | [default to &quot;public&quot;]
 **maxUnityVersion** | **string** | The maximum Unity version supported by the asset. | 
 **minUnityVersion** | **string** | The minimum Unity version supported by the asset. | 
 **platform** | **string** | The platform the asset supports. | 
 **userId** | **string** | Target user to see information on, admin-only. | 

### Return type

[**[]LimitedWorld**](LimitedWorld.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorld

> World GetWorld(ctx, worldId).Execute()

Get World by ID



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorldsAPI.GetWorld(context.Background(), worldId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorldsAPI.GetWorld``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorld`: World
	fmt.Fprintf(os.Stdout, "Response from `WorldsAPI.GetWorld`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**worldId** | **string** | Must be a valid world ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**World**](World.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorldInstance

> Instance GetWorldInstance(ctx, worldId, instanceId).Execute()

Get World Instance



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
	resp, r, err := apiClient.WorldsAPI.GetWorldInstance(context.Background(), worldId, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorldsAPI.GetWorldInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorldInstance`: Instance
	fmt.Fprintf(os.Stdout, "Response from `WorldsAPI.GetWorldInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**worldId** | **string** | Must be a valid world ID. | 
**instanceId** | **string** | Must be a valid instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldInstanceRequest struct via the builder pattern


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


## GetWorldMetadata

> WorldMetadata GetWorldMetadata(ctx, worldId).Execute()

Get World Metadata



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorldsAPI.GetWorldMetadata(context.Background(), worldId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorldsAPI.GetWorldMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorldMetadata`: WorldMetadata
	fmt.Fprintf(os.Stdout, "Response from `WorldsAPI.GetWorldMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**worldId** | **string** | Must be a valid world ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorldMetadata**](WorldMetadata.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorldPublishStatus

> WorldPublishStatus GetWorldPublishStatus(ctx, worldId).Execute()

Get World Publish Status



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorldsAPI.GetWorldPublishStatus(context.Background(), worldId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorldsAPI.GetWorldPublishStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorldPublishStatus`: WorldPublishStatus
	fmt.Fprintf(os.Stdout, "Response from `WorldsAPI.GetWorldPublishStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**worldId** | **string** | Must be a valid world ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldPublishStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorldPublishStatus**](WorldPublishStatus.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishWorld

> PublishWorld(ctx, worldId).Execute()

Publish World



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorldsAPI.PublishWorld(context.Background(), worldId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorldsAPI.PublishWorld``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**worldId** | **string** | Must be a valid world ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishWorldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchWorlds

> []LimitedWorld SearchWorlds(ctx).Featured(featured).Sort(sort).User(user).UserId(userId).N(n).Order(order).Offset(offset).Search(search).Tag(tag).Notag(notag).ReleaseStatus(releaseStatus).MaxUnityVersion(maxUnityVersion).MinUnityVersion(minUnityVersion).Platform(platform).Noplatform(noplatform).Fuzzy(fuzzy).AvatarSpecific(avatarSpecific).Execute()

Search All Worlds



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
	featured := true // bool | Filters on featured results. (optional)
	sort := openapiclient.SortOption("_created_at") // SortOption | The sort order of the results. (optional) (default to "popularity")
	user := "user_example" // string | Set to `me` for searching own worlds. (optional)
	userId := "userId_example" // string | Filter by UserID. (optional)
	n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
	order := openapiclient.OrderOption("ascending") // OrderOption | Result ordering (optional) (default to "descending")
	offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)
	search := "search_example" // string | Filters by world name. (optional)
	tag := "tag_example" // string | Tags to include (comma-separated). Any of the tags needs to be present. (optional)
	notag := "notag_example" // string | Tags to exclude (comma-separated). (optional)
	releaseStatus := openapiclient.ReleaseStatus("all") // ReleaseStatus | Filter by ReleaseStatus. (optional) (default to "public")
	maxUnityVersion := "maxUnityVersion_example" // string | The maximum Unity version supported by the asset. (optional)
	minUnityVersion := "minUnityVersion_example" // string | The minimum Unity version supported by the asset. (optional)
	platform := "platform_example" // string | The platform the asset supports. (optional)
	noplatform := "noplatform_example" // string | The platform the asset does not support. (optional)
	fuzzy := true // bool |  (optional)
	avatarSpecific := true // bool | Only search for avatar worlds. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorldsAPI.SearchWorlds(context.Background()).Featured(featured).Sort(sort).User(user).UserId(userId).N(n).Order(order).Offset(offset).Search(search).Tag(tag).Notag(notag).ReleaseStatus(releaseStatus).MaxUnityVersion(maxUnityVersion).MinUnityVersion(minUnityVersion).Platform(platform).Noplatform(noplatform).Fuzzy(fuzzy).AvatarSpecific(avatarSpecific).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorldsAPI.SearchWorlds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchWorlds`: []LimitedWorld
	fmt.Fprintf(os.Stdout, "Response from `WorldsAPI.SearchWorlds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchWorldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **featured** | **bool** | Filters on featured results. | 
 **sort** | [**SortOption**](SortOption.md) | The sort order of the results. | [default to &quot;popularity&quot;]
 **user** | **string** | Set to &#x60;me&#x60; for searching own worlds. | 
 **userId** | **string** | Filter by UserID. | 
 **n** | **int32** | The number of objects to return. | [default to 60]
 **order** | [**OrderOption**](OrderOption.md) | Result ordering | [default to &quot;descending&quot;]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 
 **search** | **string** | Filters by world name. | 
 **tag** | **string** | Tags to include (comma-separated). Any of the tags needs to be present. | 
 **notag** | **string** | Tags to exclude (comma-separated). | 
 **releaseStatus** | [**ReleaseStatus**](ReleaseStatus.md) | Filter by ReleaseStatus. | [default to &quot;public&quot;]
 **maxUnityVersion** | **string** | The maximum Unity version supported by the asset. | 
 **minUnityVersion** | **string** | The minimum Unity version supported by the asset. | 
 **platform** | **string** | The platform the asset supports. | 
 **noplatform** | **string** | The platform the asset does not support. | 
 **fuzzy** | **bool** |  | 
 **avatarSpecific** | **bool** | Only search for avatar worlds. | 

### Return type

[**[]LimitedWorld**](LimitedWorld.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnpublishWorld

> UnpublishWorld(ctx, worldId).Execute()

Unpublish World



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorldsAPI.UnpublishWorld(context.Background(), worldId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorldsAPI.UnpublishWorld``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**worldId** | **string** | Must be a valid world ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnpublishWorldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorld

> World UpdateWorld(ctx, worldId).UpdateWorldRequest(updateWorldRequest).Execute()

Update World



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
	updateWorldRequest := *openapiclient.NewUpdateWorldRequest() // UpdateWorldRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorldsAPI.UpdateWorld(context.Background(), worldId).UpdateWorldRequest(updateWorldRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorldsAPI.UpdateWorld``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWorld`: World
	fmt.Fprintf(os.Stdout, "Response from `WorldsAPI.UpdateWorld`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**worldId** | **string** | Must be a valid world ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWorldRequest** | [**UpdateWorldRequest**](UpdateWorldRequest.md) |  | 

### Return type

[**World**](World.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

