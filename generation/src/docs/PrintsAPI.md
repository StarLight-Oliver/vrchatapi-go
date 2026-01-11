# \PrintsAPI

All URIs are relative to *https://api.vrchat.cloud/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePrint**](PrintsAPI.md#DeletePrint) | **Delete** /prints/{printId} | Delete Print
[**EditPrint**](PrintsAPI.md#EditPrint) | **Post** /prints/{printId} | Edit Print
[**GetPrint**](PrintsAPI.md#GetPrint) | **Get** /prints/{printId} | Get Print
[**GetUserPrints**](PrintsAPI.md#GetUserPrints) | **Get** /prints/user/{userId} | Get Own Prints
[**UploadPrint**](PrintsAPI.md#UploadPrint) | **Post** /prints | Upload Print



## DeletePrint

> DeletePrint(ctx, printId).Execute()

Delete Print



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
	printId := "prnt_0a0aa0a0-85ea-42eb-b2f7-4840d7f341fa" // string | Print ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PrintsAPI.DeletePrint(context.Background(), printId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrintsAPI.DeletePrint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printId** | **string** | Print ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrintRequest struct via the builder pattern


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


## EditPrint

> Print EditPrint(ctx, printId).Image(image).Note(note).Execute()

Edit Print



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
	printId := "prnt_0a0aa0a0-85ea-42eb-b2f7-4840d7f341fa" // string | Print ID.
	image := os.NewFile(1234, "some_file") // *os.File | The binary blob of the png file.
	note := "note_example" // string | The caption for the image. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrintsAPI.EditPrint(context.Background(), printId).Image(image).Note(note).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrintsAPI.EditPrint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditPrint`: Print
	fmt.Fprintf(os.Stdout, "Response from `PrintsAPI.EditPrint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printId** | **string** | Print ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditPrintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **image** | ***os.File** | The binary blob of the png file. | 
 **note** | **string** | The caption for the image. | 

### Return type

[**Print**](Print.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrint

> Print GetPrint(ctx, printId).Execute()

Get Print



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
	printId := "prnt_0a0aa0a0-85ea-42eb-b2f7-4840d7f341fa" // string | Print ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrintsAPI.GetPrint(context.Background(), printId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrintsAPI.GetPrint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrint`: Print
	fmt.Fprintf(os.Stdout, "Response from `PrintsAPI.GetPrint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printId** | **string** | Print ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Print**](Print.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserPrints

> []Print GetUserPrints(ctx, userId).Execute()

Get Own Prints



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
	resp, r, err := apiClient.PrintsAPI.GetUserPrints(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrintsAPI.GetUserPrints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserPrints`: []Print
	fmt.Fprintf(os.Stdout, "Response from `PrintsAPI.GetUserPrints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserPrintsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Print**](Print.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadPrint

> Print UploadPrint(ctx).Image(image).Timestamp(timestamp).Note(note).WorldId(worldId).WorldName(worldName).Execute()

Upload Print



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
	image := os.NewFile(1234, "some_file") // *os.File | The binary blob of the png file.
	timestamp := time.Now() // time.Time | The time the image was captured.
	note := "note_example" // string | The caption for the image. (optional)
	worldId := "worldId_example" // string | The id of the world in which the image was captured. (optional)
	worldName := "worldName_example" // string | The name of the world in which the image was captured. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrintsAPI.UploadPrint(context.Background()).Image(image).Timestamp(timestamp).Note(note).WorldId(worldId).WorldName(worldName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrintsAPI.UploadPrint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadPrint`: Print
	fmt.Fprintf(os.Stdout, "Response from `PrintsAPI.UploadPrint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadPrintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **image** | ***os.File** | The binary blob of the png file. | 
 **timestamp** | **time.Time** | The time the image was captured. | 
 **note** | **string** | The caption for the image. | 
 **worldId** | **string** | The id of the world in which the image was captured. | 
 **worldName** | **string** | The name of the world in which the image was captured. | 

### Return type

[**Print**](Print.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

