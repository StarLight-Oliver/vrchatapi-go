# \NotificationsAPI

All URIs are relative to *https://api.vrchat.cloud/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptFriendRequest**](NotificationsAPI.md#AcceptFriendRequest) | **Put** /auth/user/notifications/{notificationId}/accept | Accept Friend Request
[**AcknowledgeNotificationV2**](NotificationsAPI.md#AcknowledgeNotificationV2) | **Post** /notifications/{notificationId}/see | Acknowledge NotificationV2
[**ClearNotifications**](NotificationsAPI.md#ClearNotifications) | **Put** /auth/user/notifications/clear | Clear All Notifications
[**DeleteAllNotificationV2s**](NotificationsAPI.md#DeleteAllNotificationV2s) | **Delete** /notifications | Delete All NotificationV2s
[**DeleteNotification**](NotificationsAPI.md#DeleteNotification) | **Put** /auth/user/notifications/{notificationId}/hide | Delete Notification
[**DeleteNotificationV2**](NotificationsAPI.md#DeleteNotificationV2) | **Delete** /notifications/{notificationId} | Delete NotificationV2
[**GetNotification**](NotificationsAPI.md#GetNotification) | **Get** /auth/user/notifications/{notificationId} | Show notification
[**GetNotificationV2**](NotificationsAPI.md#GetNotificationV2) | **Get** /notifications/{notificationId} | Get NotificationV2
[**GetNotificationV2s**](NotificationsAPI.md#GetNotificationV2s) | **Get** /notifications | List NotificationV2s
[**GetNotifications**](NotificationsAPI.md#GetNotifications) | **Get** /auth/user/notifications | List Notifications
[**MarkNotificationAsRead**](NotificationsAPI.md#MarkNotificationAsRead) | **Put** /auth/user/notifications/{notificationId}/see | Mark Notification As Read
[**ReplyNotificationV2**](NotificationsAPI.md#ReplyNotificationV2) | **Post** /notifications/{notificationId}/reply | Reply NotificationV2
[**RespondNotificationV2**](NotificationsAPI.md#RespondNotificationV2) | **Post** /notifications/{notificationId}/respond | Respond NotificationV2



## AcceptFriendRequest

> Success AcceptFriendRequest(ctx, notificationId).Execute()

Accept Friend Request



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
	notificationId := "notificationId_example" // string | Must be a valid notification ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.AcceptFriendRequest(context.Background(), notificationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.AcceptFriendRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcceptFriendRequest`: Success
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.AcceptFriendRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** | Must be a valid notification ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptFriendRequestRequest struct via the builder pattern


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


## AcknowledgeNotificationV2

> NotificationV2 AcknowledgeNotificationV2(ctx, notificationId).Execute()

Acknowledge NotificationV2



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
	notificationId := "notificationId_example" // string | Must be a valid notification ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.AcknowledgeNotificationV2(context.Background(), notificationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.AcknowledgeNotificationV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcknowledgeNotificationV2`: NotificationV2
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.AcknowledgeNotificationV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** | Must be a valid notification ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcknowledgeNotificationV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationV2**](NotificationV2.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClearNotifications

> Success ClearNotifications(ctx).Execute()

Clear All Notifications



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
	resp, r, err := apiClient.NotificationsAPI.ClearNotifications(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.ClearNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClearNotifications`: Success
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.ClearNotifications`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiClearNotificationsRequest struct via the builder pattern


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


## DeleteAllNotificationV2s

> Success DeleteAllNotificationV2s(ctx).Execute()

Delete All NotificationV2s



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
	resp, r, err := apiClient.NotificationsAPI.DeleteAllNotificationV2s(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.DeleteAllNotificationV2s``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAllNotificationV2s`: Success
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.DeleteAllNotificationV2s`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllNotificationV2sRequest struct via the builder pattern


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


## DeleteNotification

> Notification DeleteNotification(ctx, notificationId).Execute()

Delete Notification



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
	notificationId := "notificationId_example" // string | Must be a valid notification ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.DeleteNotification(context.Background(), notificationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.DeleteNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNotification`: Notification
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.DeleteNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** | Must be a valid notification ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Notification**](Notification.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotificationV2

> Success DeleteNotificationV2(ctx, notificationId).Execute()

Delete NotificationV2



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
	notificationId := "notificationId_example" // string | Must be a valid notification ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.DeleteNotificationV2(context.Background(), notificationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.DeleteNotificationV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNotificationV2`: Success
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.DeleteNotificationV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** | Must be a valid notification ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationV2Request struct via the builder pattern


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


## GetNotification

> Notification GetNotification(ctx, notificationId).Execute()

Show notification



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
	notificationId := "notificationId_example" // string | Must be a valid notification ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.GetNotification(context.Background(), notificationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotification`: Notification
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** | Must be a valid notification ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Notification**](Notification.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationV2

> NotificationV2 GetNotificationV2(ctx, notificationId).Execute()

Get NotificationV2



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
	notificationId := "notificationId_example" // string | Must be a valid notification ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.GetNotificationV2(context.Background(), notificationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetNotificationV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationV2`: NotificationV2
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetNotificationV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** | Must be a valid notification ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationV2**](NotificationV2.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationV2s

> []NotificationV2 GetNotificationV2s(ctx).Limit(limit).Execute()

List NotificationV2s



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
	limit := int32(100) // int32 | The maximum number of entries to get. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.GetNotificationV2s(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetNotificationV2s``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationV2s`: []NotificationV2
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetNotificationV2s`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationV2sRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of entries to get. | 

### Return type

[**[]NotificationV2**](NotificationV2.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotifications

> []Notification GetNotifications(ctx).Type_(type_).Sent(sent).Hidden(hidden).After(after).N(n).Offset(offset).Execute()

List Notifications



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
	type_ := "all" // string | Only send notifications of this type (can use `all` for all). This parameter no longer does anything, and is deprecated. (optional)
	sent := true // bool | Return notifications sent by the user. Must be false or omitted. (optional)
	hidden := true // bool | Whether to return hidden or non-hidden notifications. True only allowed on type `friendRequest`. (optional)
	after := "five_minutes_ago" // string | Only return notifications sent after this Date. Ignored if type is `friendRequest`. (optional)
	n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
	offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.GetNotifications(context.Background()).Type_(type_).Sent(sent).Hidden(hidden).After(after).N(n).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotifications`: []Notification
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetNotifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Only send notifications of this type (can use &#x60;all&#x60; for all). This parameter no longer does anything, and is deprecated. | 
 **sent** | **bool** | Return notifications sent by the user. Must be false or omitted. | 
 **hidden** | **bool** | Whether to return hidden or non-hidden notifications. True only allowed on type &#x60;friendRequest&#x60;. | 
 **after** | **string** | Only return notifications sent after this Date. Ignored if type is &#x60;friendRequest&#x60;. | 
 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 

### Return type

[**[]Notification**](Notification.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkNotificationAsRead

> Notification MarkNotificationAsRead(ctx, notificationId).Execute()

Mark Notification As Read



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
	notificationId := "notificationId_example" // string | Must be a valid notification ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.MarkNotificationAsRead(context.Background(), notificationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.MarkNotificationAsRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkNotificationAsRead`: Notification
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.MarkNotificationAsRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** | Must be a valid notification ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkNotificationAsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Notification**](Notification.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplyNotificationV2

> NotificationV2 ReplyNotificationV2(ctx, notificationId).Body(body).Execute()

Reply NotificationV2



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
	notificationId := "notificationId_example" // string | Must be a valid notification ID.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.ReplyNotificationV2(context.Background(), notificationId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.ReplyNotificationV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplyNotificationV2`: NotificationV2
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.ReplyNotificationV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** | Must be a valid notification ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplyNotificationV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**NotificationV2**](NotificationV2.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RespondNotificationV2

> NotificationV2 RespondNotificationV2(ctx, notificationId).RespondNotificationV2Request(respondNotificationV2Request).Execute()

Respond NotificationV2



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
	notificationId := "notificationId_example" // string | Must be a valid notification ID.
	respondNotificationV2Request := *openapiclient.NewRespondNotificationV2Request("Accept:
  value: accept
Boop:
  value: boop
Decline:
  value: decline
Delete:
  value: delete
Unsubscribe:
  value: unsubscribe
") // RespondNotificationV2Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.RespondNotificationV2(context.Background(), notificationId).RespondNotificationV2Request(respondNotificationV2Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.RespondNotificationV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RespondNotificationV2`: NotificationV2
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.RespondNotificationV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** | Must be a valid notification ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRespondNotificationV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **respondNotificationV2Request** | [**RespondNotificationV2Request**](RespondNotificationV2Request.md) |  | 

### Return type

[**NotificationV2**](NotificationV2.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

