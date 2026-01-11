# \MiscellaneousAPI

All URIs are relative to *https://api.vrchat.cloud/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAssignedPermissions**](MiscellaneousAPI.md#GetAssignedPermissions) | **Get** /auth/permissions | Get Assigned Permissions
[**GetCSS**](MiscellaneousAPI.md#GetCSS) | **Get** /css/app.css | Download CSS
[**GetConfig**](MiscellaneousAPI.md#GetConfig) | **Get** /config | Fetch API Config
[**GetCurrentOnlineUsers**](MiscellaneousAPI.md#GetCurrentOnlineUsers) | **Get** /visits | Current Online Users
[**GetHealth**](MiscellaneousAPI.md#GetHealth) | **Get** /health | Check API Health
[**GetInfoPush**](MiscellaneousAPI.md#GetInfoPush) | **Get** /infoPush | Show Information Notices
[**GetJavaScript**](MiscellaneousAPI.md#GetJavaScript) | **Get** /js/app.js | Download JavaScript
[**GetPermission**](MiscellaneousAPI.md#GetPermission) | **Get** /permissions/{permissionId} | Get Permission
[**GetSystemTime**](MiscellaneousAPI.md#GetSystemTime) | **Get** /time | Current System Time



## GetAssignedPermissions

> []Permission GetAssignedPermissions(ctx).Execute()

Get Assigned Permissions



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
	resp, r, err := apiClient.MiscellaneousAPI.GetAssignedPermissions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiscellaneousAPI.GetAssignedPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssignedPermissions`: []Permission
	fmt.Fprintf(os.Stdout, "Response from `MiscellaneousAPI.GetAssignedPermissions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssignedPermissionsRequest struct via the builder pattern


### Return type

[**[]Permission**](Permission.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCSS

> string GetCSS(ctx).Variant(variant).Branch(branch).Execute()

Download CSS



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
	variant := "variant_example" // string | Specifies which `variant` of the site. Public is the end-user site, while `internal` is the staff-only site with special pages for moderation and management. (optional) (default to "public")
	branch := "branch_example" // string | Specifies which git branch the site should load frontend source code from. (optional) (default to "main")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiscellaneousAPI.GetCSS(context.Background()).Variant(variant).Branch(branch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiscellaneousAPI.GetCSS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCSS`: string
	fmt.Fprintf(os.Stdout, "Response from `MiscellaneousAPI.GetCSS`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCSSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variant** | **string** | Specifies which &#x60;variant&#x60; of the site. Public is the end-user site, while &#x60;internal&#x60; is the staff-only site with special pages for moderation and management. | [default to &quot;public&quot;]
 **branch** | **string** | Specifies which git branch the site should load frontend source code from. | [default to &quot;main&quot;]

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/css, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfig

> APIConfig GetConfig(ctx).Execute()

Fetch API Config



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
	resp, r, err := apiClient.MiscellaneousAPI.GetConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiscellaneousAPI.GetConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfig`: APIConfig
	fmt.Fprintf(os.Stdout, "Response from `MiscellaneousAPI.GetConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigRequest struct via the builder pattern


### Return type

[**APIConfig**](APIConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentOnlineUsers

> int32 GetCurrentOnlineUsers(ctx).Execute()

Current Online Users



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
	resp, r, err := apiClient.MiscellaneousAPI.GetCurrentOnlineUsers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiscellaneousAPI.GetCurrentOnlineUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentOnlineUsers`: int32
	fmt.Fprintf(os.Stdout, "Response from `MiscellaneousAPI.GetCurrentOnlineUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentOnlineUsersRequest struct via the builder pattern


### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHealth

> APIHealth GetHealth(ctx).Execute()

Check API Health



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
	resp, r, err := apiClient.MiscellaneousAPI.GetHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiscellaneousAPI.GetHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHealth`: APIHealth
	fmt.Fprintf(os.Stdout, "Response from `MiscellaneousAPI.GetHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHealthRequest struct via the builder pattern


### Return type

[**APIHealth**](APIHealth.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfoPush

> []InfoPush GetInfoPush(ctx).Require(require).Include(include).Execute()

Show Information Notices



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
	require := "require_example" // string | Tags to include (comma-separated). All of the tags needs to be present. (optional)
	include := "include_example" // string | Tags to include (comma-separated). Any of the tags needs to be present. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiscellaneousAPI.GetInfoPush(context.Background()).Require(require).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiscellaneousAPI.GetInfoPush``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfoPush`: []InfoPush
	fmt.Fprintf(os.Stdout, "Response from `MiscellaneousAPI.GetInfoPush`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInfoPushRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **require** | **string** | Tags to include (comma-separated). All of the tags needs to be present. | 
 **include** | **string** | Tags to include (comma-separated). Any of the tags needs to be present. | 

### Return type

[**[]InfoPush**](InfoPush.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJavaScript

> string GetJavaScript(ctx).Variant(variant).Branch(branch).Execute()

Download JavaScript



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
	variant := "variant_example" // string | Specifies which `variant` of the site. Public is the end-user site, while `internal` is the staff-only site with special pages for moderation and management. (optional) (default to "public")
	branch := "branch_example" // string | Specifies which git branch the site should load frontend source code from. (optional) (default to "main")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiscellaneousAPI.GetJavaScript(context.Background()).Variant(variant).Branch(branch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiscellaneousAPI.GetJavaScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJavaScript`: string
	fmt.Fprintf(os.Stdout, "Response from `MiscellaneousAPI.GetJavaScript`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJavaScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variant** | **string** | Specifies which &#x60;variant&#x60; of the site. Public is the end-user site, while &#x60;internal&#x60; is the staff-only site with special pages for moderation and management. | [default to &quot;public&quot;]
 **branch** | **string** | Specifies which git branch the site should load frontend source code from. | [default to &quot;main&quot;]

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/javascript, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPermission

> Permission GetPermission(ctx, permissionId).Execute()

Get Permission



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
	permissionId := "permissionId_example" // string | Must be a valid permission ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiscellaneousAPI.GetPermission(context.Background(), permissionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiscellaneousAPI.GetPermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPermission`: Permission
	fmt.Fprintf(os.Stdout, "Response from `MiscellaneousAPI.GetPermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionId** | **string** | Must be a valid permission ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Permission**](Permission.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemTime

> time.Time GetSystemTime(ctx).Execute()

Current System Time



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
	resp, r, err := apiClient.MiscellaneousAPI.GetSystemTime(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiscellaneousAPI.GetSystemTime``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemTime`: time.Time
	fmt.Fprintf(os.Stdout, "Response from `MiscellaneousAPI.GetSystemTime`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemTimeRequest struct via the builder pattern


### Return type

[**time.Time**](time.Time.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

