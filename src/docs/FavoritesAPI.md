# \FavoritesAPI

All URIs are relative to *https://api.vrchat.cloud/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFavorite**](FavoritesAPI.md#AddFavorite) | **Post** /favorites | Add Favorite
[**ClearFavoriteGroup**](FavoritesAPI.md#ClearFavoriteGroup) | **Delete** /favorite/group/{favoriteGroupType}/{favoriteGroupName}/{userId} | Clear Favorite Group
[**GetFavoriteGroup**](FavoritesAPI.md#GetFavoriteGroup) | **Get** /favorite/group/{favoriteGroupType}/{favoriteGroupName}/{userId} | Show Favorite Group
[**GetFavoriteGroups**](FavoritesAPI.md#GetFavoriteGroups) | **Get** /favorite/groups | List Favorite Groups
[**GetFavoriteLimits**](FavoritesAPI.md#GetFavoriteLimits) | **Get** /auth/user/favoritelimits | Get Favorite Limits
[**GetFavorites**](FavoritesAPI.md#GetFavorites) | **Get** /favorites | List Favorites
[**RemoveFavorite**](FavoritesAPI.md#RemoveFavorite) | **Delete** /favorites/{favoriteId} | Remove Favorite
[**UpdateFavoriteGroup**](FavoritesAPI.md#UpdateFavoriteGroup) | **Put** /favorite/group/{favoriteGroupType}/{favoriteGroupName}/{userId} | Update Favorite Group



## AddFavorite

> Favorite AddFavorite(ctx).AddFavoriteRequest(addFavoriteRequest).Execute()

Add Favorite



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
	addFavoriteRequest := *openapiclient.NewAddFavoriteRequest("FavoriteId_example", []string{"Tags_example"}, openapiclient.FavoriteType("avatar")) // AddFavoriteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FavoritesAPI.AddFavorite(context.Background()).AddFavoriteRequest(addFavoriteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FavoritesAPI.AddFavorite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddFavorite`: Favorite
	fmt.Fprintf(os.Stdout, "Response from `FavoritesAPI.AddFavorite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddFavoriteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addFavoriteRequest** | [**AddFavoriteRequest**](AddFavoriteRequest.md) |  | 

### Return type

[**Favorite**](Favorite.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClearFavoriteGroup

> Success ClearFavoriteGroup(ctx, favoriteGroupType, favoriteGroupName, userId).Execute()

Clear Favorite Group



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
	favoriteGroupType := "favoriteGroupType_example" // string | The type of group to fetch, must be a valid FavoriteType.
	favoriteGroupName := "favoriteGroupName_example" // string | The name of the group to fetch, must be a name of a FavoriteGroup.
	userId := "userId_example" // string | Must be a valid user ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FavoritesAPI.ClearFavoriteGroup(context.Background(), favoriteGroupType, favoriteGroupName, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FavoritesAPI.ClearFavoriteGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClearFavoriteGroup`: Success
	fmt.Fprintf(os.Stdout, "Response from `FavoritesAPI.ClearFavoriteGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**favoriteGroupType** | **string** | The type of group to fetch, must be a valid FavoriteType. | 
**favoriteGroupName** | **string** | The name of the group to fetch, must be a name of a FavoriteGroup. | 
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearFavoriteGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Success**](Success.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFavoriteGroup

> FavoriteGroup GetFavoriteGroup(ctx, favoriteGroupType, favoriteGroupName, userId).Execute()

Show Favorite Group



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
	favoriteGroupType := "favoriteGroupType_example" // string | The type of group to fetch, must be a valid FavoriteType.
	favoriteGroupName := "favoriteGroupName_example" // string | The name of the group to fetch, must be a name of a FavoriteGroup.
	userId := "userId_example" // string | Must be a valid user ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FavoritesAPI.GetFavoriteGroup(context.Background(), favoriteGroupType, favoriteGroupName, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FavoritesAPI.GetFavoriteGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFavoriteGroup`: FavoriteGroup
	fmt.Fprintf(os.Stdout, "Response from `FavoritesAPI.GetFavoriteGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**favoriteGroupType** | **string** | The type of group to fetch, must be a valid FavoriteType. | 
**favoriteGroupName** | **string** | The name of the group to fetch, must be a name of a FavoriteGroup. | 
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFavoriteGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FavoriteGroup**](FavoriteGroup.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFavoriteGroups

> []FavoriteGroup GetFavoriteGroups(ctx).N(n).Offset(offset).UserId(userId).OwnerId(ownerId).Execute()

List Favorite Groups



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
	userId := "userId_example" // string | Target user to see information on, admin-only. (optional)
	ownerId := "ownerId_example" // string | The owner of whoms favorite groups to return. Must be a UserID. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FavoritesAPI.GetFavoriteGroups(context.Background()).N(n).Offset(offset).UserId(userId).OwnerId(ownerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FavoritesAPI.GetFavoriteGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFavoriteGroups`: []FavoriteGroup
	fmt.Fprintf(os.Stdout, "Response from `FavoritesAPI.GetFavoriteGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFavoriteGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 
 **userId** | **string** | Target user to see information on, admin-only. | 
 **ownerId** | **string** | The owner of whoms favorite groups to return. Must be a UserID. | 

### Return type

[**[]FavoriteGroup**](FavoriteGroup.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFavoriteLimits

> FavoriteLimits GetFavoriteLimits(ctx).Execute()

Get Favorite Limits



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
	resp, r, err := apiClient.FavoritesAPI.GetFavoriteLimits(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FavoritesAPI.GetFavoriteLimits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFavoriteLimits`: FavoriteLimits
	fmt.Fprintf(os.Stdout, "Response from `FavoritesAPI.GetFavoriteLimits`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFavoriteLimitsRequest struct via the builder pattern


### Return type

[**FavoriteLimits**](FavoriteLimits.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFavorites

> []Favorite GetFavorites(ctx).N(n).Offset(offset).Type_(type_).Tag(tag).Execute()

List Favorites



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
	type_ := "type__example" // string | The type of favorites to return, FavoriteType. (optional)
	tag := "tag_example" // string | Tags to include (comma-separated). Any of the tags needs to be present. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FavoritesAPI.GetFavorites(context.Background()).N(n).Offset(offset).Type_(type_).Tag(tag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FavoritesAPI.GetFavorites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFavorites`: []Favorite
	fmt.Fprintf(os.Stdout, "Response from `FavoritesAPI.GetFavorites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFavoritesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 
 **type_** | **string** | The type of favorites to return, FavoriteType. | 
 **tag** | **string** | Tags to include (comma-separated). Any of the tags needs to be present. | 

### Return type

[**[]Favorite**](Favorite.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveFavorite

> Success RemoveFavorite(ctx, favoriteId).Execute()

Remove Favorite



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
	favoriteId := "favoriteId_example" // string | Must be a valid favorite ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FavoritesAPI.RemoveFavorite(context.Background(), favoriteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FavoritesAPI.RemoveFavorite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveFavorite`: Success
	fmt.Fprintf(os.Stdout, "Response from `FavoritesAPI.RemoveFavorite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**favoriteId** | **string** | Must be a valid favorite ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveFavoriteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Success**](Success.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFavoriteGroup

> UpdateFavoriteGroup(ctx, favoriteGroupType, favoriteGroupName, userId).UpdateFavoriteGroupRequest(updateFavoriteGroupRequest).Execute()

Update Favorite Group



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
	favoriteGroupType := "favoriteGroupType_example" // string | The type of group to fetch, must be a valid FavoriteType.
	favoriteGroupName := "favoriteGroupName_example" // string | The name of the group to fetch, must be a name of a FavoriteGroup.
	userId := "userId_example" // string | Must be a valid user ID.
	updateFavoriteGroupRequest := *openapiclient.NewUpdateFavoriteGroupRequest() // UpdateFavoriteGroupRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FavoritesAPI.UpdateFavoriteGroup(context.Background(), favoriteGroupType, favoriteGroupName, userId).UpdateFavoriteGroupRequest(updateFavoriteGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FavoritesAPI.UpdateFavoriteGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**favoriteGroupType** | **string** | The type of group to fetch, must be a valid FavoriteType. | 
**favoriteGroupName** | **string** | The name of the group to fetch, must be a name of a FavoriteGroup. | 
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFavoriteGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateFavoriteGroupRequest** | [**UpdateFavoriteGroupRequest**](UpdateFavoriteGroupRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

