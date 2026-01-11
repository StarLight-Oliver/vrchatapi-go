# \JamsAPI

All URIs are relative to *https://api.vrchat.cloud/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetJam**](JamsAPI.md#GetJam) | **Get** /jams/{jamId} | Show jam information
[**GetJamSubmissions**](JamsAPI.md#GetJamSubmissions) | **Get** /jams/{jamId}/submissions | Show jam submissions
[**GetJams**](JamsAPI.md#GetJams) | **Get** /jams | Show jams list



## GetJam

> Jam GetJam(ctx, jamId).Execute()

Show jam information



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
	jamId := "jam_0b7e3f6d-4647-4648-b2a1-1431e76906d9" // string | Must be a valid query ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JamsAPI.GetJam(context.Background(), jamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamsAPI.GetJam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJam`: Jam
	fmt.Fprintf(os.Stdout, "Response from `JamsAPI.GetJam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jamId** | **string** | Must be a valid query ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Jam**](Jam.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJamSubmissions

> []Submission GetJamSubmissions(ctx, jamId).Execute()

Show jam submissions



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
	jamId := "jam_0b7e3f6d-4647-4648-b2a1-1431e76906d9" // string | Must be a valid query ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JamsAPI.GetJamSubmissions(context.Background(), jamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamsAPI.GetJamSubmissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJamSubmissions`: []Submission
	fmt.Fprintf(os.Stdout, "Response from `JamsAPI.GetJamSubmissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jamId** | **string** | Must be a valid query ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJamSubmissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Submission**](Submission.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJams

> []Jam GetJams(ctx).Type_(type_).Execute()

Show jams list



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
	type_ := "avatar" // string | Only show jams of this type (`avatar` or `world`). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JamsAPI.GetJams(context.Background()).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JamsAPI.GetJams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJams`: []Jam
	fmt.Fprintf(os.Stdout, "Response from `JamsAPI.GetJams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Only show jams of this type (&#x60;avatar&#x60; or &#x60;world&#x60;). | 

### Return type

[**[]Jam**](Jam.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

