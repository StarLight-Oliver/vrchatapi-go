# \PlayermoderationAPI

All URIs are relative to *https://api.vrchat.cloud/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearAllPlayerModerations**](PlayermoderationAPI.md#ClearAllPlayerModerations) | **Delete** /auth/user/playermoderations | Clear All Player Moderations
[**GetPlayerModerations**](PlayermoderationAPI.md#GetPlayerModerations) | **Get** /auth/user/playermoderations | Search Player Moderations
[**ModerateUser**](PlayermoderationAPI.md#ModerateUser) | **Post** /auth/user/playermoderations | Moderate User
[**UnmoderateUser**](PlayermoderationAPI.md#UnmoderateUser) | **Put** /auth/user/unplayermoderate | Unmoderate User



## ClearAllPlayerModerations

> Success ClearAllPlayerModerations(ctx).Execute()

Clear All Player Moderations



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
	resp, r, err := apiClient.PlayermoderationAPI.ClearAllPlayerModerations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayermoderationAPI.ClearAllPlayerModerations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClearAllPlayerModerations`: Success
	fmt.Fprintf(os.Stdout, "Response from `PlayermoderationAPI.ClearAllPlayerModerations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiClearAllPlayerModerationsRequest struct via the builder pattern


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


## GetPlayerModerations

> []PlayerModeration GetPlayerModerations(ctx).Type_(type_).TargetUserId(targetUserId).Execute()

Search Player Moderations



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
	type_ := openapiclient.PlayerModerationType("block") // PlayerModerationType | Must be one of PlayerModerationType. (optional) (default to "unmute")
	targetUserId := "targetUserId_example" // string | Must be valid UserID. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayermoderationAPI.GetPlayerModerations(context.Background()).Type_(type_).TargetUserId(targetUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayermoderationAPI.GetPlayerModerations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlayerModerations`: []PlayerModeration
	fmt.Fprintf(os.Stdout, "Response from `PlayermoderationAPI.GetPlayerModerations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayerModerationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**PlayerModerationType**](PlayerModerationType.md) | Must be one of PlayerModerationType. | [default to &quot;unmute&quot;]
 **targetUserId** | **string** | Must be valid UserID. | 

### Return type

[**[]PlayerModeration**](PlayerModeration.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModerateUser

> PlayerModeration ModerateUser(ctx).ModerateUserRequest(moderateUserRequest).Execute()

Moderate User



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
	moderateUserRequest := *openapiclient.NewModerateUserRequest("usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469", openapiclient.PlayerModerationType("block")) // ModerateUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayermoderationAPI.ModerateUser(context.Background()).ModerateUserRequest(moderateUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayermoderationAPI.ModerateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModerateUser`: PlayerModeration
	fmt.Fprintf(os.Stdout, "Response from `PlayermoderationAPI.ModerateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModerateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **moderateUserRequest** | [**ModerateUserRequest**](ModerateUserRequest.md) |  | 

### Return type

[**PlayerModeration**](PlayerModeration.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnmoderateUser

> Success UnmoderateUser(ctx).ModerateUserRequest(moderateUserRequest).Execute()

Unmoderate User



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
	moderateUserRequest := *openapiclient.NewModerateUserRequest("usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469", openapiclient.PlayerModerationType("block")) // ModerateUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlayermoderationAPI.UnmoderateUser(context.Background()).ModerateUserRequest(moderateUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlayermoderationAPI.UnmoderateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnmoderateUser`: Success
	fmt.Fprintf(os.Stdout, "Response from `PlayermoderationAPI.UnmoderateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnmoderateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **moderateUserRequest** | [**ModerateUserRequest**](ModerateUserRequest.md) |  | 

### Return type

[**Success**](Success.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

