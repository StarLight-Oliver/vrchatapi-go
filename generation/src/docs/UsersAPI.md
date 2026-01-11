# \UsersAPI

All URIs are relative to *https://api.vrchat.cloud/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTags**](UsersAPI.md#AddTags) | **Post** /users/{userId}/addTags | Add User Tags
[**CheckUserPersistenceExists**](UsersAPI.md#CheckUserPersistenceExists) | **Get** /users/{userId}/{worldId}/persist/exists | Check User Persistence Exists
[**DeleteAllUserPersistenceData**](UsersAPI.md#DeleteAllUserPersistenceData) | **Delete** /users/{userId}/persist | Delete All User Persistence Data
[**DeleteUserPersistence**](UsersAPI.md#DeleteUserPersistence) | **Delete** /users/{userId}/{worldId}/persist | Delete User Persistence
[**GetBlockedGroups**](UsersAPI.md#GetBlockedGroups) | **Get** /users/{userId}/groups/userblocked | Get User Group Blocks
[**GetInvitedGroups**](UsersAPI.md#GetInvitedGroups) | **Get** /users/{userId}/groups/invited | Get User Group Invited
[**GetMutualFriends**](UsersAPI.md#GetMutualFriends) | **Get** /users/{userId}/mutuals/friends | Get User Mutual Friends
[**GetMutualGroups**](UsersAPI.md#GetMutualGroups) | **Get** /users/{userId}/mutuals/groups | Get User Mutual Groups
[**GetMutuals**](UsersAPI.md#GetMutuals) | **Get** /users/{userId}/mutuals | Get User Mutuals
[**GetUser**](UsersAPI.md#GetUser) | **Get** /users/{userId} | Get User by ID
[**GetUserAllGroupPermissions**](UsersAPI.md#GetUserAllGroupPermissions) | **Get** /users/{userId}/groups/permissions | Get user&#39;s permissions for all joined groups.
[**GetUserByName**](UsersAPI.md#GetUserByName) | **Get** /users/{username}/name | Get User by Username
[**GetUserFeedback**](UsersAPI.md#GetUserFeedback) | **Get** /users/{userId}/feedback | Get User Feedback
[**GetUserGroupInstances**](UsersAPI.md#GetUserGroupInstances) | **Get** /users/{userId}/instances/groups | Get User Group Instances
[**GetUserGroupInstancesForGroup**](UsersAPI.md#GetUserGroupInstancesForGroup) | **Get** /users/{userId}/instances/groups/{groupId} | Get User Group Instances for a specific Group
[**GetUserGroupRequests**](UsersAPI.md#GetUserGroupRequests) | **Get** /users/{userId}/groups/requested | Get User Group Requests
[**GetUserGroups**](UsersAPI.md#GetUserGroups) | **Get** /users/{userId}/groups | Get User Groups
[**GetUserNote**](UsersAPI.md#GetUserNote) | **Get** /userNotes/{userNoteId} | Get User Note
[**GetUserNotes**](UsersAPI.md#GetUserNotes) | **Get** /userNotes | Get User Notes
[**GetUserRepresentedGroup**](UsersAPI.md#GetUserRepresentedGroup) | **Get** /users/{userId}/groups/represented | Get user&#39;s current represented group
[**RemoveTags**](UsersAPI.md#RemoveTags) | **Post** /users/{userId}/removeTags | Remove User Tags
[**SearchUsers**](UsersAPI.md#SearchUsers) | **Get** /users | Search All Users
[**UpdateBadge**](UsersAPI.md#UpdateBadge) | **Put** /users/{userId}/badges/{badgeId} | Update User Badge
[**UpdateUser**](UsersAPI.md#UpdateUser) | **Put** /users/{userId} | Update User Info
[**UpdateUserNote**](UsersAPI.md#UpdateUserNote) | **Post** /userNotes | Update User Note



## AddTags

> CurrentUser AddTags(ctx, userId).ChangeUserTagsRequest(changeUserTagsRequest).Execute()

Add User Tags



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
	changeUserTagsRequest := *openapiclient.NewChangeUserTagsRequest([]string{"Tags_example"}) // ChangeUserTagsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AddTags(context.Background(), userId).ChangeUserTagsRequest(changeUserTagsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AddTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddTags`: CurrentUser
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AddTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changeUserTagsRequest** | [**ChangeUserTagsRequest**](ChangeUserTagsRequest.md) |  | 

### Return type

[**CurrentUser**](CurrentUser.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	r, err := apiClient.UsersAPI.CheckUserPersistenceExists(context.Background(), userId, worldId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CheckUserPersistenceExists``: %v\n", err)
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
	r, err := apiClient.UsersAPI.DeleteAllUserPersistenceData(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteAllUserPersistenceData``: %v\n", err)
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
	r, err := apiClient.UsersAPI.DeleteUserPersistence(context.Background(), userId, worldId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteUserPersistence``: %v\n", err)
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


## GetBlockedGroups

> []Group GetBlockedGroups(ctx, userId).Execute()

Get User Group Blocks



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
	resp, r, err := apiClient.UsersAPI.GetBlockedGroups(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetBlockedGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockedGroups`: []Group
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetBlockedGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockedGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Group**](Group.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvitedGroups

> []Group GetInvitedGroups(ctx, userId).Execute()

Get User Group Invited



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
	resp, r, err := apiClient.UsersAPI.GetInvitedGroups(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetInvitedGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvitedGroups`: []Group
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetInvitedGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvitedGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Group**](Group.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMutualFriends

> []MutualFriend GetMutualFriends(ctx, userId).N(n).Offset(offset).Execute()

Get User Mutual Friends



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
	n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
	offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetMutualFriends(context.Background(), userId).N(n).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetMutualFriends``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMutualFriends`: []MutualFriend
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetMutualFriends`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMutualFriendsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 

### Return type

[**[]MutualFriend**](MutualFriend.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMutualGroups

> []LimitedUserGroups GetMutualGroups(ctx, userId).N(n).Offset(offset).Execute()

Get User Mutual Groups



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
	n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
	offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetMutualGroups(context.Background(), userId).N(n).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetMutualGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMutualGroups`: []LimitedUserGroups
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetMutualGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMutualGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 

### Return type

[**[]LimitedUserGroups**](LimitedUserGroups.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMutuals

> Mutuals GetMutuals(ctx, userId).Execute()

Get User Mutuals



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
	resp, r, err := apiClient.UsersAPI.GetMutuals(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetMutuals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMutuals`: Mutuals
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetMutuals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMutualsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Mutuals**](Mutuals.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> User GetUser(ctx, userId).Execute()

Get User by ID



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
	resp, r, err := apiClient.UsersAPI.GetUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserAllGroupPermissions

> map[string][]GroupPermissions GetUserAllGroupPermissions(ctx, userId).GroupIds(groupIds).Execute()

Get user's permissions for all joined groups.



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
	groupIds := "grp_00000000-0000-0000-0000-000000000000,grp_11111111-1111-1111-1111-111111111111" // string | Comma-separated (no spaces!) list of GroupIDs to retrieve permissions for. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUserAllGroupPermissions(context.Background(), userId).GroupIds(groupIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserAllGroupPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserAllGroupPermissions`: map[string][]GroupPermissions
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserAllGroupPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAllGroupPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupIds** | **string** | Comma-separated (no spaces!) list of GroupIDs to retrieve permissions for. | 

### Return type

[**map[string][]GroupPermissions**](array.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserByName

> User GetUserByName(ctx, username).Execute()

Get User by Username



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
	username := "username_example" // string | Username of the user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUserByName(context.Background(), username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserByName`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserFeedback

> []Feedback GetUserFeedback(ctx, userId).ContentId(contentId).N(n).Offset(offset).Execute()

Get User Feedback



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
	contentId := true // bool | Filter for users' previously submitted feedback, e.g., a groupId, userId, avatarId, etc. (optional)
	n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
	offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUserFeedback(context.Background(), userId).ContentId(contentId).N(n).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserFeedback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserFeedback`: []Feedback
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserFeedback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentId** | **bool** | Filter for users&#39; previously submitted feedback, e.g., a groupId, userId, avatarId, etc. | 
 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 

### Return type

[**[]Feedback**](Feedback.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserGroupInstances

> GetUserGroupInstances200Response GetUserGroupInstances(ctx, userId).Execute()

Get User Group Instances



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
	resp, r, err := apiClient.UsersAPI.GetUserGroupInstances(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserGroupInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserGroupInstances`: GetUserGroupInstances200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserGroupInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserGroupInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetUserGroupInstances200Response**](GetUserGroupInstances200Response.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserGroupInstancesForGroup

> GetUserGroupInstances200Response GetUserGroupInstancesForGroup(ctx, userId, groupId).Execute()

Get User Group Instances for a specific Group



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
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUserGroupInstancesForGroup(context.Background(), userId, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserGroupInstancesForGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserGroupInstancesForGroup`: GetUserGroupInstances200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserGroupInstancesForGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserGroupInstancesForGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetUserGroupInstances200Response**](GetUserGroupInstances200Response.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserGroupRequests

> []Group GetUserGroupRequests(ctx, userId).Execute()

Get User Group Requests



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
	resp, r, err := apiClient.UsersAPI.GetUserGroupRequests(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserGroupRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserGroupRequests`: []Group
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserGroupRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserGroupRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Group**](Group.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserGroups

> []LimitedUserGroups GetUserGroups(ctx, userId).Execute()

Get User Groups



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
	resp, r, err := apiClient.UsersAPI.GetUserGroups(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserGroups`: []LimitedUserGroups
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]LimitedUserGroups**](LimitedUserGroups.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserNote

> UserNote GetUserNote(ctx, userNoteId).Execute()

Get User Note



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
	userNoteId := "userNoteId_example" // string | Must be a valid user note ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUserNote(context.Background(), userNoteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserNote`: UserNote
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userNoteId** | **string** | Must be a valid user note ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserNote**](UserNote.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserNotes

> []UserNote GetUserNotes(ctx).N(n).Offset(offset).Execute()

Get User Notes



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
	resp, r, err := apiClient.UsersAPI.GetUserNotes(context.Background()).N(n).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserNotes`: []UserNote
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserNotes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 

### Return type

[**[]UserNote**](UserNote.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserRepresentedGroup

> RepresentedGroup GetUserRepresentedGroup(ctx, userId).Execute()

Get user's current represented group



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
	resp, r, err := apiClient.UsersAPI.GetUserRepresentedGroup(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserRepresentedGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserRepresentedGroup`: RepresentedGroup
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserRepresentedGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRepresentedGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RepresentedGroup**](RepresentedGroup.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveTags

> CurrentUser RemoveTags(ctx, userId).ChangeUserTagsRequest(changeUserTagsRequest).Execute()

Remove User Tags



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
	changeUserTagsRequest := *openapiclient.NewChangeUserTagsRequest([]string{"Tags_example"}) // ChangeUserTagsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.RemoveTags(context.Background(), userId).ChangeUserTagsRequest(changeUserTagsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.RemoveTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveTags`: CurrentUser
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.RemoveTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changeUserTagsRequest** | [**ChangeUserTagsRequest**](ChangeUserTagsRequest.md) |  | 

### Return type

[**CurrentUser**](CurrentUser.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchUsers

> []LimitedUserSearch SearchUsers(ctx).Search(search).DeveloperType(developerType).N(n).Offset(offset).IsInternalVariant(isInternalVariant).Execute()

Search All Users



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
	search := "search_example" // string | Searches by `displayName`. Will return empty array if search query is empty or missing. (optional)
	developerType := "developerType_example" // string | Active user by developer type, none for normal users and internal for moderators (optional)
	n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
	offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)
	isInternalVariant := false // bool | Not quite sure what this actually does (exists on the website but doesn't seem to be used) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.SearchUsers(context.Background()).Search(search).DeveloperType(developerType).N(n).Offset(offset).IsInternalVariant(isInternalVariant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.SearchUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchUsers`: []LimitedUserSearch
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.SearchUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Searches by &#x60;displayName&#x60;. Will return empty array if search query is empty or missing. | 
 **developerType** | **string** | Active user by developer type, none for normal users and internal for moderators | 
 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 
 **isInternalVariant** | **bool** | Not quite sure what this actually does (exists on the website but doesn&#39;t seem to be used) | 

### Return type

[**[]LimitedUserSearch**](LimitedUserSearch.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBadge

> UpdateBadge(ctx, userId, badgeId).UpdateUserBadgeRequest(updateUserBadgeRequest).Execute()

Update User Badge



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
	badgeId := "badgeId_example" // string | Must be a valid badge ID.
	updateUserBadgeRequest := *openapiclient.NewUpdateUserBadgeRequest() // UpdateUserBadgeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UpdateBadge(context.Background(), userId, badgeId).UpdateUserBadgeRequest(updateUserBadgeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateBadge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 
**badgeId** | **string** | Must be a valid badge ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBadgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateUserBadgeRequest** | [**UpdateUserBadgeRequest**](UpdateUserBadgeRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> CurrentUser UpdateUser(ctx, userId).UpdateUserRequest(updateUserRequest).Execute()

Update User Info



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
	updateUserRequest := *openapiclient.NewUpdateUserRequest() // UpdateUserRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateUser(context.Background(), userId).UpdateUserRequest(updateUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUser`: CurrentUser
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateUserRequest** | [**UpdateUserRequest**](UpdateUserRequest.md) |  | 

### Return type

[**CurrentUser**](CurrentUser.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserNote

> UserNote UpdateUserNote(ctx).UpdateUserNoteRequest(updateUserNoteRequest).Execute()

Update User Note



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
	updateUserNoteRequest := *openapiclient.NewUpdateUserNoteRequest("Note_example", "usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469") // UpdateUserNoteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UpdateUserNote(context.Background()).UpdateUserNoteRequest(updateUserNoteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserNote`: UserNote
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UpdateUserNote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateUserNoteRequest** | [**UpdateUserNoteRequest**](UpdateUserNoteRequest.md) |  | 

### Return type

[**UserNote**](UserNote.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

