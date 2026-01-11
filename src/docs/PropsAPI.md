# \PropsAPI

All URIs are relative to *https://api.vrchat.cloud/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProp**](PropsAPI.md#CreateProp) | **Post** /props | Create Prop
[**DeleteProp**](PropsAPI.md#DeleteProp) | **Delete** /props/{propId} | Delete Prop
[**GetProp**](PropsAPI.md#GetProp) | **Get** /props/{propId} | Get Prop
[**GetPropPublishStatus**](PropsAPI.md#GetPropPublishStatus) | **Get** /props/{propId}/publish | Get Prop Publish Status
[**ListProps**](PropsAPI.md#ListProps) | **Get** /props | List Props
[**PublishProp**](PropsAPI.md#PublishProp) | **Put** /props/{propId}/publish | Publish Prop
[**UnpublishProp**](PropsAPI.md#UnpublishProp) | **Delete** /props/{propId}/publish | Unpublish Prop
[**UpdateProp**](PropsAPI.md#UpdateProp) | **Put** /props/{propId} | Update Prop



## CreateProp

> Prop CreateProp(ctx).CreatePropRequest(createPropRequest).Execute()

Create Prop



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
	createPropRequest := *openapiclient.NewCreatePropRequest("AssetUrl_example", int32(123), "Description_example", "prop_829ba6f6-b837-49d9-b9a9-056b82103b58", "ImageUrl_example", "Name_example", "standalonewindows", int32(123), []string{"Tags_example"}, "UnityVersion_example", int32(123)) // CreatePropRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PropsAPI.CreateProp(context.Background()).CreatePropRequest(createPropRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PropsAPI.CreateProp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProp`: Prop
	fmt.Fprintf(os.Stdout, "Response from `PropsAPI.CreateProp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePropRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPropRequest** | [**CreatePropRequest**](CreatePropRequest.md) |  | 

### Return type

[**Prop**](Prop.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProp

> DeleteProp(ctx, propId).Execute()

Delete Prop



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
	propId := "prop_829ba6f6-b837-49d9-b9a9-056b82103b58" // string | Prop ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PropsAPI.DeleteProp(context.Background(), propId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PropsAPI.DeleteProp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**propId** | **string** | Prop ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePropRequest struct via the builder pattern


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


## GetProp

> Prop GetProp(ctx, propId).Execute()

Get Prop



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
	propId := "prop_829ba6f6-b837-49d9-b9a9-056b82103b58" // string | Prop ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PropsAPI.GetProp(context.Background(), propId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PropsAPI.GetProp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProp`: Prop
	fmt.Fprintf(os.Stdout, "Response from `PropsAPI.GetProp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**propId** | **string** | Prop ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPropRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Prop**](Prop.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPropPublishStatus

> PropPublishStatus GetPropPublishStatus(ctx, propId).Execute()

Get Prop Publish Status



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
	propId := "prop_829ba6f6-b837-49d9-b9a9-056b82103b58" // string | Prop ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PropsAPI.GetPropPublishStatus(context.Background(), propId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PropsAPI.GetPropPublishStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPropPublishStatus`: PropPublishStatus
	fmt.Fprintf(os.Stdout, "Response from `PropsAPI.GetPropPublishStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**propId** | **string** | Prop ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPropPublishStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PropPublishStatus**](PropPublishStatus.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProps

> []Prop ListProps(ctx).AuthorId(authorId).N(n).Offset(offset).Execute()

List Props



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
	authorId := "authorId_example" // string | Must be a valid user ID.
	n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
	offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PropsAPI.ListProps(context.Background()).AuthorId(authorId).N(n).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PropsAPI.ListProps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListProps`: []Prop
	fmt.Fprintf(os.Stdout, "Response from `PropsAPI.ListProps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPropsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorId** | **string** | Must be a valid user ID. | 
 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 

### Return type

[**[]Prop**](Prop.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishProp

> PropPublishStatus PublishProp(ctx, propId).Execute()

Publish Prop



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
	propId := "prop_829ba6f6-b837-49d9-b9a9-056b82103b58" // string | Prop ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PropsAPI.PublishProp(context.Background(), propId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PropsAPI.PublishProp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishProp`: PropPublishStatus
	fmt.Fprintf(os.Stdout, "Response from `PropsAPI.PublishProp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**propId** | **string** | Prop ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishPropRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PropPublishStatus**](PropPublishStatus.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnpublishProp

> PropPublishStatus UnpublishProp(ctx, propId).Execute()

Unpublish Prop



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
	propId := "prop_829ba6f6-b837-49d9-b9a9-056b82103b58" // string | Prop ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PropsAPI.UnpublishProp(context.Background(), propId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PropsAPI.UnpublishProp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnpublishProp`: PropPublishStatus
	fmt.Fprintf(os.Stdout, "Response from `PropsAPI.UnpublishProp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**propId** | **string** | Prop ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnpublishPropRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PropPublishStatus**](PropPublishStatus.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProp

> Prop UpdateProp(ctx, propId).UpdatePropRequest(updatePropRequest).Execute()

Update Prop



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
	propId := "prop_829ba6f6-b837-49d9-b9a9-056b82103b58" // string | Prop ID.
	updatePropRequest := *openapiclient.NewUpdatePropRequest() // UpdatePropRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PropsAPI.UpdateProp(context.Background(), propId).UpdatePropRequest(updatePropRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PropsAPI.UpdateProp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProp`: Prop
	fmt.Fprintf(os.Stdout, "Response from `PropsAPI.UpdateProp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**propId** | **string** | Prop ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePropRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePropRequest** | [**UpdatePropRequest**](UpdatePropRequest.md) |  | 

### Return type

[**Prop**](Prop.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

