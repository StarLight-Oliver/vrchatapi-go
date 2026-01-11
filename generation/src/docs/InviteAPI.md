# \InviteAPI

All URIs are relative to *https://api.vrchat.cloud/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInviteMessage**](InviteAPI.md#GetInviteMessage) | **Get** /message/{userId}/{messageType}/{slot} | Get Invite Message
[**GetInviteMessages**](InviteAPI.md#GetInviteMessages) | **Get** /message/{userId}/{messageType} | List Invite Messages
[**InviteMyselfTo**](InviteAPI.md#InviteMyselfTo) | **Post** /invite/myself/to/{worldId}:{instanceId} | Invite Myself To Instance
[**InviteUser**](InviteAPI.md#InviteUser) | **Post** /invite/{userId} | Invite User
[**InviteUserWithPhoto**](InviteAPI.md#InviteUserWithPhoto) | **Post** /invite/{userId}/photo | Invite User with photo
[**RequestInvite**](InviteAPI.md#RequestInvite) | **Post** /requestInvite/{userId} | Request Invite
[**RequestInviteWithPhoto**](InviteAPI.md#RequestInviteWithPhoto) | **Post** /requestInvite/{userId}/photo | Request Invite with photo
[**ResetInviteMessage**](InviteAPI.md#ResetInviteMessage) | **Delete** /message/{userId}/{messageType}/{slot} | Reset Invite Message
[**RespondInvite**](InviteAPI.md#RespondInvite) | **Post** /invite/{notificationId}/response | Respond Invite
[**RespondInviteWithPhoto**](InviteAPI.md#RespondInviteWithPhoto) | **Post** /invite/{notificationId}/response/photo | Respond Invite with photo
[**UpdateInviteMessage**](InviteAPI.md#UpdateInviteMessage) | **Put** /message/{userId}/{messageType}/{slot} | Update Invite Message



## GetInviteMessage

> InviteMessage GetInviteMessage(ctx, userId, messageType, slot).Execute()

Get Invite Message



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
	messageType := openapiclient.InviteMessageType("message") // InviteMessageType | The type of message to fetch, must be a valid InviteMessageType. (default to "message")
	slot := int32(56) // int32 | The message slot to fetch of a given message type.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InviteAPI.GetInviteMessage(context.Background(), userId, messageType, slot).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InviteAPI.GetInviteMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInviteMessage`: InviteMessage
	fmt.Fprintf(os.Stdout, "Response from `InviteAPI.GetInviteMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 
**messageType** | [**InviteMessageType**](.md) | The type of message to fetch, must be a valid InviteMessageType. | [default to &quot;message&quot;]
**slot** | **int32** | The message slot to fetch of a given message type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInviteMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**InviteMessage**](InviteMessage.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInviteMessages

> []InviteMessage GetInviteMessages(ctx, userId, messageType).Execute()

List Invite Messages



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
	messageType := openapiclient.InviteMessageType("message") // InviteMessageType | The type of message to fetch, must be a valid InviteMessageType. (default to "message")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InviteAPI.GetInviteMessages(context.Background(), userId, messageType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InviteAPI.GetInviteMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInviteMessages`: []InviteMessage
	fmt.Fprintf(os.Stdout, "Response from `InviteAPI.GetInviteMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 
**messageType** | [**InviteMessageType**](.md) | The type of message to fetch, must be a valid InviteMessageType. | [default to &quot;message&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetInviteMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]InviteMessage**](InviteMessage.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InviteMyselfTo

> SentNotification InviteMyselfTo(ctx, worldId, instanceId).Execute()

Invite Myself To Instance



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
	resp, r, err := apiClient.InviteAPI.InviteMyselfTo(context.Background(), worldId, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InviteAPI.InviteMyselfTo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InviteMyselfTo`: SentNotification
	fmt.Fprintf(os.Stdout, "Response from `InviteAPI.InviteMyselfTo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**worldId** | **string** | Must be a valid world ID. | 
**instanceId** | **string** | Must be a valid instance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInviteMyselfToRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SentNotification**](SentNotification.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InviteUser

> SentNotification InviteUser(ctx, userId).InviteRequest(inviteRequest).Execute()

Invite User



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
	inviteRequest := *openapiclient.NewInviteRequest("12345~hidden(usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469)~region(eu)~nonce(27e8414a-59a0-4f3d-af1f-f27557eb49a2)") // InviteRequest | Slot number of the Invite Message to use when inviting a user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InviteAPI.InviteUser(context.Background(), userId).InviteRequest(inviteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InviteAPI.InviteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InviteUser`: SentNotification
	fmt.Fprintf(os.Stdout, "Response from `InviteAPI.InviteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInviteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inviteRequest** | [**InviteRequest**](InviteRequest.md) | Slot number of the Invite Message to use when inviting a user. | 

### Return type

[**SentNotification**](SentNotification.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InviteUserWithPhoto

> SentNotification InviteUserWithPhoto(ctx, userId).Data(data).Image(image).Execute()

Invite User with photo



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
	data := *openapiclient.NewInviteRequest("12345~hidden(usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469)~region(eu)~nonce(27e8414a-59a0-4f3d-af1f-f27557eb49a2)") // InviteRequest | 
	image := os.NewFile(1234, "some_file") // *os.File | The binary blob of the png file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InviteAPI.InviteUserWithPhoto(context.Background(), userId).Data(data).Image(image).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InviteAPI.InviteUserWithPhoto``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InviteUserWithPhoto`: SentNotification
	fmt.Fprintf(os.Stdout, "Response from `InviteAPI.InviteUserWithPhoto`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInviteUserWithPhotoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**InviteRequest**](InviteRequest.md) |  | 
 **image** | ***os.File** | The binary blob of the png file. | 

### Return type

[**SentNotification**](SentNotification.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestInvite

> Notification RequestInvite(ctx, userId).RequestInviteRequest(requestInviteRequest).Execute()

Request Invite



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
	requestInviteRequest := *openapiclient.NewRequestInviteRequest() // RequestInviteRequest | Slot number of the Request Message to use when request an invite. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InviteAPI.RequestInvite(context.Background(), userId).RequestInviteRequest(requestInviteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InviteAPI.RequestInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestInvite`: Notification
	fmt.Fprintf(os.Stdout, "Response from `InviteAPI.RequestInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestInviteRequest** | [**RequestInviteRequest**](RequestInviteRequest.md) | Slot number of the Request Message to use when request an invite. | 

### Return type

[**Notification**](Notification.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestInviteWithPhoto

> Notification RequestInviteWithPhoto(ctx, userId).Data(data).Image(image).Execute()

Request Invite with photo



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
	data := *openapiclient.NewRequestInviteRequest() // RequestInviteRequest | 
	image := os.NewFile(1234, "some_file") // *os.File | The binary blob of the png file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InviteAPI.RequestInviteWithPhoto(context.Background(), userId).Data(data).Image(image).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InviteAPI.RequestInviteWithPhoto``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestInviteWithPhoto`: Notification
	fmt.Fprintf(os.Stdout, "Response from `InviteAPI.RequestInviteWithPhoto`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestInviteWithPhotoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**RequestInviteRequest**](RequestInviteRequest.md) |  | 
 **image** | ***os.File** | The binary blob of the png file. | 

### Return type

[**Notification**](Notification.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetInviteMessage

> []InviteMessage ResetInviteMessage(ctx, userId, messageType, slot).Execute()

Reset Invite Message



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
	messageType := openapiclient.InviteMessageType("message") // InviteMessageType | The type of message to fetch, must be a valid InviteMessageType. (default to "message")
	slot := int32(56) // int32 | The message slot to fetch of a given message type.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InviteAPI.ResetInviteMessage(context.Background(), userId, messageType, slot).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InviteAPI.ResetInviteMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetInviteMessage`: []InviteMessage
	fmt.Fprintf(os.Stdout, "Response from `InviteAPI.ResetInviteMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 
**messageType** | [**InviteMessageType**](.md) | The type of message to fetch, must be a valid InviteMessageType. | [default to &quot;message&quot;]
**slot** | **int32** | The message slot to fetch of a given message type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetInviteMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]InviteMessage**](InviteMessage.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RespondInvite

> Notification RespondInvite(ctx, notificationId).InviteResponse(inviteResponse).Execute()

Respond Invite



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
	inviteResponse := *openapiclient.NewInviteResponse(int32(123)) // InviteResponse | Slot number of the Response Message to use when responding to a user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InviteAPI.RespondInvite(context.Background(), notificationId).InviteResponse(inviteResponse).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InviteAPI.RespondInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RespondInvite`: Notification
	fmt.Fprintf(os.Stdout, "Response from `InviteAPI.RespondInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** | Must be a valid notification ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRespondInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inviteResponse** | [**InviteResponse**](InviteResponse.md) | Slot number of the Response Message to use when responding to a user. | 

### Return type

[**Notification**](Notification.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RespondInviteWithPhoto

> Notification RespondInviteWithPhoto(ctx, notificationId).Data(data).Image(image).Execute()

Respond Invite with photo



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
	data := *openapiclient.NewInviteResponse(int32(123)) // InviteResponse | 
	image := os.NewFile(1234, "some_file") // *os.File | The binary blob of the png file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InviteAPI.RespondInviteWithPhoto(context.Background(), notificationId).Data(data).Image(image).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InviteAPI.RespondInviteWithPhoto``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RespondInviteWithPhoto`: Notification
	fmt.Fprintf(os.Stdout, "Response from `InviteAPI.RespondInviteWithPhoto`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** | Must be a valid notification ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRespondInviteWithPhotoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**InviteResponse**](InviteResponse.md) |  | 
 **image** | ***os.File** | The binary blob of the png file. | 

### Return type

[**Notification**](Notification.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInviteMessage

> []InviteMessage UpdateInviteMessage(ctx, userId, messageType, slot).UpdateInviteMessageRequest(updateInviteMessageRequest).Execute()

Update Invite Message



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
	messageType := openapiclient.InviteMessageType("message") // InviteMessageType | The type of message to fetch, must be a valid InviteMessageType. (default to "message")
	slot := int32(56) // int32 | The message slot to fetch of a given message type.
	updateInviteMessageRequest := *openapiclient.NewUpdateInviteMessageRequest("Message_example") // UpdateInviteMessageRequest | Message of what to set the invite message to. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InviteAPI.UpdateInviteMessage(context.Background(), userId, messageType, slot).UpdateInviteMessageRequest(updateInviteMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InviteAPI.UpdateInviteMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInviteMessage`: []InviteMessage
	fmt.Fprintf(os.Stdout, "Response from `InviteAPI.UpdateInviteMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 
**messageType** | [**InviteMessageType**](.md) | The type of message to fetch, must be a valid InviteMessageType. | [default to &quot;message&quot;]
**slot** | **int32** | The message slot to fetch of a given message type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInviteMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateInviteMessageRequest** | [**UpdateInviteMessageRequest**](UpdateInviteMessageRequest.md) | Message of what to set the invite message to. | 

### Return type

[**[]InviteMessage**](InviteMessage.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

