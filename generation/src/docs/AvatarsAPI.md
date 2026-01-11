# \AvatarsAPI

All URIs are relative to *https://api.vrchat.cloud/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAvatar**](AvatarsAPI.md#CreateAvatar) | **Post** /avatars | Create Avatar
[**DeleteAvatar**](AvatarsAPI.md#DeleteAvatar) | **Delete** /avatars/{avatarId} | Delete Avatar
[**DeleteImpostor**](AvatarsAPI.md#DeleteImpostor) | **Delete** /avatars/{avatarId}/impostor | Delete generated Impostor
[**EnqueueImpostor**](AvatarsAPI.md#EnqueueImpostor) | **Post** /avatars/{avatarId}/impostor/enqueue | Enqueue Impostor generation
[**GetAvatar**](AvatarsAPI.md#GetAvatar) | **Get** /avatars/{avatarId} | Get Avatar
[**GetAvatarStyles**](AvatarsAPI.md#GetAvatarStyles) | **Get** /avatarStyles | Get Avatar Styles
[**GetFavoritedAvatars**](AvatarsAPI.md#GetFavoritedAvatars) | **Get** /avatars/favorites | List Favorited Avatars
[**GetImpostorQueueStats**](AvatarsAPI.md#GetImpostorQueueStats) | **Get** /avatars/impostor/queue/stats | Get Impostor Queue Stats
[**GetLicensedAvatars**](AvatarsAPI.md#GetLicensedAvatars) | **Get** /avatars/licensed | List Licensed Avatars
[**GetOwnAvatar**](AvatarsAPI.md#GetOwnAvatar) | **Get** /users/{userId}/avatar | Get Own Avatar
[**SearchAvatars**](AvatarsAPI.md#SearchAvatars) | **Get** /avatars | Search Avatars
[**SelectAvatar**](AvatarsAPI.md#SelectAvatar) | **Put** /avatars/{avatarId}/select | Select Avatar
[**SelectFallbackAvatar**](AvatarsAPI.md#SelectFallbackAvatar) | **Put** /avatars/{avatarId}/selectFallback | Select Fallback Avatar
[**UpdateAvatar**](AvatarsAPI.md#UpdateAvatar) | **Put** /avatars/{avatarId} | Update Avatar



## CreateAvatar

> Avatar CreateAvatar(ctx).CreateAvatarRequest(createAvatarRequest).Execute()

Create Avatar



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
	createAvatarRequest := *openapiclient.NewCreateAvatarRequest("ImageUrl_example", "Name_example") // CreateAvatarRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AvatarsAPI.CreateAvatar(context.Background()).CreateAvatarRequest(createAvatarRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvatarsAPI.CreateAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAvatar`: Avatar
	fmt.Fprintf(os.Stdout, "Response from `AvatarsAPI.CreateAvatar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAvatarRequest** | [**CreateAvatarRequest**](CreateAvatarRequest.md) |  | 

### Return type

[**Avatar**](Avatar.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAvatar

> Avatar DeleteAvatar(ctx, avatarId).Execute()

Delete Avatar



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
	avatarId := "avatarId_example" // string | Must be a valid avatar ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AvatarsAPI.DeleteAvatar(context.Background(), avatarId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvatarsAPI.DeleteAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAvatar`: Avatar
	fmt.Fprintf(os.Stdout, "Response from `AvatarsAPI.DeleteAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**avatarId** | **string** | Must be a valid avatar ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Avatar**](Avatar.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteImpostor

> DeleteImpostor(ctx, avatarId).Execute()

Delete generated Impostor



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
	avatarId := "avatarId_example" // string | Must be a valid avatar ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AvatarsAPI.DeleteImpostor(context.Background(), avatarId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvatarsAPI.DeleteImpostor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**avatarId** | **string** | Must be a valid avatar ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteImpostorRequest struct via the builder pattern


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


## EnqueueImpostor

> ServiceStatus EnqueueImpostor(ctx, avatarId).Execute()

Enqueue Impostor generation



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
	avatarId := "avatarId_example" // string | Must be a valid avatar ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AvatarsAPI.EnqueueImpostor(context.Background(), avatarId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvatarsAPI.EnqueueImpostor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnqueueImpostor`: ServiceStatus
	fmt.Fprintf(os.Stdout, "Response from `AvatarsAPI.EnqueueImpostor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**avatarId** | **string** | Must be a valid avatar ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnqueueImpostorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceStatus**](ServiceStatus.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvatar

> Avatar GetAvatar(ctx, avatarId).Execute()

Get Avatar



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
	avatarId := "avatarId_example" // string | Must be a valid avatar ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AvatarsAPI.GetAvatar(context.Background(), avatarId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvatarsAPI.GetAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvatar`: Avatar
	fmt.Fprintf(os.Stdout, "Response from `AvatarsAPI.GetAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**avatarId** | **string** | Must be a valid avatar ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Avatar**](Avatar.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvatarStyles

> []AvatarStyle GetAvatarStyles(ctx).Execute()

Get Avatar Styles



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
	resp, r, err := apiClient.AvatarsAPI.GetAvatarStyles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvatarsAPI.GetAvatarStyles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvatarStyles`: []AvatarStyle
	fmt.Fprintf(os.Stdout, "Response from `AvatarsAPI.GetAvatarStyles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvatarStylesRequest struct via the builder pattern


### Return type

[**[]AvatarStyle**](AvatarStyle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFavoritedAvatars

> []Avatar GetFavoritedAvatars(ctx).Featured(featured).Sort(sort).N(n).Order(order).Offset(offset).Search(search).Tag(tag).Notag(notag).ReleaseStatus(releaseStatus).MaxUnityVersion(maxUnityVersion).MinUnityVersion(minUnityVersion).Platform(platform).UserId(userId).Execute()

List Favorited Avatars



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
	resp, r, err := apiClient.AvatarsAPI.GetFavoritedAvatars(context.Background()).Featured(featured).Sort(sort).N(n).Order(order).Offset(offset).Search(search).Tag(tag).Notag(notag).ReleaseStatus(releaseStatus).MaxUnityVersion(maxUnityVersion).MinUnityVersion(minUnityVersion).Platform(platform).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvatarsAPI.GetFavoritedAvatars``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFavoritedAvatars`: []Avatar
	fmt.Fprintf(os.Stdout, "Response from `AvatarsAPI.GetFavoritedAvatars`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFavoritedAvatarsRequest struct via the builder pattern


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

[**[]Avatar**](Avatar.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImpostorQueueStats

> ServiceQueueStats GetImpostorQueueStats(ctx).Execute()

Get Impostor Queue Stats



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
	resp, r, err := apiClient.AvatarsAPI.GetImpostorQueueStats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvatarsAPI.GetImpostorQueueStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImpostorQueueStats`: ServiceQueueStats
	fmt.Fprintf(os.Stdout, "Response from `AvatarsAPI.GetImpostorQueueStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetImpostorQueueStatsRequest struct via the builder pattern


### Return type

[**ServiceQueueStats**](ServiceQueueStats.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicensedAvatars

> []Avatar GetLicensedAvatars(ctx).N(n).Offset(offset).Execute()

List Licensed Avatars



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
	resp, r, err := apiClient.AvatarsAPI.GetLicensedAvatars(context.Background()).N(n).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvatarsAPI.GetLicensedAvatars``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLicensedAvatars`: []Avatar
	fmt.Fprintf(os.Stdout, "Response from `AvatarsAPI.GetLicensedAvatars`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLicensedAvatarsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 

### Return type

[**[]Avatar**](Avatar.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOwnAvatar

> Avatar GetOwnAvatar(ctx, userId).Execute()

Get Own Avatar



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
	resp, r, err := apiClient.AvatarsAPI.GetOwnAvatar(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvatarsAPI.GetOwnAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOwnAvatar`: Avatar
	fmt.Fprintf(os.Stdout, "Response from `AvatarsAPI.GetOwnAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOwnAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Avatar**](Avatar.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAvatars

> []Avatar SearchAvatars(ctx).Featured(featured).Sort(sort).User(user).UserId(userId).N(n).Order(order).Offset(offset).Tag(tag).Notag(notag).ReleaseStatus(releaseStatus).MaxUnityVersion(maxUnityVersion).MinUnityVersion(minUnityVersion).Platform(platform).IsInternalVariant(isInternalVariant).Execute()

Search Avatars



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
	user := "user_example" // string | Set to `me` for searching own avatars. (optional)
	userId := "userId_example" // string | Filter by UserID. (optional)
	n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
	order := openapiclient.OrderOption("ascending") // OrderOption | Result ordering (optional) (default to "descending")
	offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)
	tag := "tag_example" // string | Tags to include (comma-separated). Any of the tags needs to be present. (optional)
	notag := "notag_example" // string | Tags to exclude (comma-separated). (optional)
	releaseStatus := openapiclient.ReleaseStatus("all") // ReleaseStatus | Filter by ReleaseStatus. (optional) (default to "public")
	maxUnityVersion := "maxUnityVersion_example" // string | The maximum Unity version supported by the asset. (optional)
	minUnityVersion := "minUnityVersion_example" // string | The minimum Unity version supported by the asset. (optional)
	platform := "platform_example" // string | The platform the asset supports. (optional)
	isInternalVariant := false // bool | Not quite sure what this actually does (exists on the website but doesn't seem to be used) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AvatarsAPI.SearchAvatars(context.Background()).Featured(featured).Sort(sort).User(user).UserId(userId).N(n).Order(order).Offset(offset).Tag(tag).Notag(notag).ReleaseStatus(releaseStatus).MaxUnityVersion(maxUnityVersion).MinUnityVersion(minUnityVersion).Platform(platform).IsInternalVariant(isInternalVariant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvatarsAPI.SearchAvatars``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAvatars`: []Avatar
	fmt.Fprintf(os.Stdout, "Response from `AvatarsAPI.SearchAvatars`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAvatarsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **featured** | **bool** | Filters on featured results. | 
 **sort** | [**SortOption**](SortOption.md) | The sort order of the results. | [default to &quot;popularity&quot;]
 **user** | **string** | Set to &#x60;me&#x60; for searching own avatars. | 
 **userId** | **string** | Filter by UserID. | 
 **n** | **int32** | The number of objects to return. | [default to 60]
 **order** | [**OrderOption**](OrderOption.md) | Result ordering | [default to &quot;descending&quot;]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 
 **tag** | **string** | Tags to include (comma-separated). Any of the tags needs to be present. | 
 **notag** | **string** | Tags to exclude (comma-separated). | 
 **releaseStatus** | [**ReleaseStatus**](ReleaseStatus.md) | Filter by ReleaseStatus. | [default to &quot;public&quot;]
 **maxUnityVersion** | **string** | The maximum Unity version supported by the asset. | 
 **minUnityVersion** | **string** | The minimum Unity version supported by the asset. | 
 **platform** | **string** | The platform the asset supports. | 
 **isInternalVariant** | **bool** | Not quite sure what this actually does (exists on the website but doesn&#39;t seem to be used) | 

### Return type

[**[]Avatar**](Avatar.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SelectAvatar

> CurrentUser SelectAvatar(ctx, avatarId).Execute()

Select Avatar



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
	avatarId := "avatarId_example" // string | Must be a valid avatar ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AvatarsAPI.SelectAvatar(context.Background(), avatarId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvatarsAPI.SelectAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SelectAvatar`: CurrentUser
	fmt.Fprintf(os.Stdout, "Response from `AvatarsAPI.SelectAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**avatarId** | **string** | Must be a valid avatar ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSelectAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CurrentUser**](CurrentUser.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SelectFallbackAvatar

> CurrentUser SelectFallbackAvatar(ctx, avatarId).Execute()

Select Fallback Avatar



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
	avatarId := "avatarId_example" // string | Must be a valid avatar ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AvatarsAPI.SelectFallbackAvatar(context.Background(), avatarId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvatarsAPI.SelectFallbackAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SelectFallbackAvatar`: CurrentUser
	fmt.Fprintf(os.Stdout, "Response from `AvatarsAPI.SelectFallbackAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**avatarId** | **string** | Must be a valid avatar ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSelectFallbackAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CurrentUser**](CurrentUser.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAvatar

> Avatar UpdateAvatar(ctx, avatarId).UpdateAvatarRequest(updateAvatarRequest).Execute()

Update Avatar



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
	avatarId := "avatarId_example" // string | Must be a valid avatar ID.
	updateAvatarRequest := *openapiclient.NewUpdateAvatarRequest() // UpdateAvatarRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AvatarsAPI.UpdateAvatar(context.Background(), avatarId).UpdateAvatarRequest(updateAvatarRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvatarsAPI.UpdateAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAvatar`: Avatar
	fmt.Fprintf(os.Stdout, "Response from `AvatarsAPI.UpdateAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**avatarId** | **string** | Must be a valid avatar ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAvatarRequest** | [**UpdateAvatarRequest**](UpdateAvatarRequest.md) |  | 

### Return type

[**Avatar**](Avatar.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

