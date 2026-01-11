# \AuthenticationAPI

All URIs are relative to *https://api.vrchat.cloud/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelPending2FA**](AuthenticationAPI.md#CancelPending2FA) | **Delete** /auth/twofactorauth/totp/pending | Cancel pending enabling of time-based 2FA codes
[**CheckUserExists**](AuthenticationAPI.md#CheckUserExists) | **Get** /auth/exists | Check User Exists
[**ConfirmEmail**](AuthenticationAPI.md#ConfirmEmail) | **Get** /auth/confirmEmail | Confirm Email
[**CreateGlobalAvatarModeration**](AuthenticationAPI.md#CreateGlobalAvatarModeration) | **Post** /auth/user/avatarmoderations | Create Global Avatar Moderation
[**DeleteGlobalAvatarModeration**](AuthenticationAPI.md#DeleteGlobalAvatarModeration) | **Delete** /auth/user/avatarmoderations | Delete Global Avatar Moderation
[**DeleteModerationReport**](AuthenticationAPI.md#DeleteModerationReport) | **Delete** /moderationReports/{moderationReportId} | Delete Moderation Report
[**DeleteUser**](AuthenticationAPI.md#DeleteUser) | **Put** /users/{userId}/delete | Delete User
[**Disable2FA**](AuthenticationAPI.md#Disable2FA) | **Delete** /auth/twofactorauth | Disable 2FA
[**Enable2FA**](AuthenticationAPI.md#Enable2FA) | **Post** /auth/twofactorauth/totp/pending | Enable time-based 2FA codes
[**GetCurrentUser**](AuthenticationAPI.md#GetCurrentUser) | **Get** /auth/user | Login and/or Get Current User Info
[**GetGlobalAvatarModerations**](AuthenticationAPI.md#GetGlobalAvatarModerations) | **Get** /auth/user/avatarmoderations | Get Global Avatar Moderations
[**GetModerationReports**](AuthenticationAPI.md#GetModerationReports) | **Get** /moderationReports | Get Moderation Reports
[**GetRecoveryCodes**](AuthenticationAPI.md#GetRecoveryCodes) | **Get** /auth/user/twofactorauth/otp | Get 2FA Recovery codes
[**Logout**](AuthenticationAPI.md#Logout) | **Put** /logout | Logout
[**RegisterUserAccount**](AuthenticationAPI.md#RegisterUserAccount) | **Post** /auth/register | Register User Account
[**ResendEmailConfirmation**](AuthenticationAPI.md#ResendEmailConfirmation) | **Post** /auth/user/resendEmail | Resend Email Confirmation
[**SubmitModerationReport**](AuthenticationAPI.md#SubmitModerationReport) | **Post** /moderationReports | Submit Moderation Report
[**Verify2FA**](AuthenticationAPI.md#Verify2FA) | **Post** /auth/twofactorauth/totp/verify | Verify 2FA code
[**Verify2FAEmailCode**](AuthenticationAPI.md#Verify2FAEmailCode) | **Post** /auth/twofactorauth/emailotp/verify | Verify 2FA email code
[**VerifyAuthToken**](AuthenticationAPI.md#VerifyAuthToken) | **Get** /auth | Verify Auth Token
[**VerifyLoginPlace**](AuthenticationAPI.md#VerifyLoginPlace) | **Get** /auth/verifyLoginPlace | Verify Login Place
[**VerifyPending2FA**](AuthenticationAPI.md#VerifyPending2FA) | **Post** /auth/twofactorauth/totp/pending/verify | Verify Pending 2FA code
[**VerifyRecoveryCode**](AuthenticationAPI.md#VerifyRecoveryCode) | **Post** /auth/twofactorauth/otp/verify | Verify 2FA code with Recovery code



## CancelPending2FA

> Disable2FAResult CancelPending2FA(ctx).Execute()

Cancel pending enabling of time-based 2FA codes



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
	resp, r, err := apiClient.AuthenticationAPI.CancelPending2FA(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.CancelPending2FA``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelPending2FA`: Disable2FAResult
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.CancelPending2FA`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCancelPending2FARequest struct via the builder pattern


### Return type

[**Disable2FAResult**](Disable2FAResult.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckUserExists

> UserExists CheckUserExists(ctx).Email(email).DisplayName(displayName).Username(username).ExcludeUserId(excludeUserId).Execute()

Check User Exists



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
	email := "email_example" // string | Filter by email. (optional)
	displayName := "displayName_example" // string | Filter by displayName. (optional)
	username := "username_example" // string | Filter by Username. (optional)
	excludeUserId := "excludeUserId_example" // string | Exclude by UserID. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.CheckUserExists(context.Background()).Email(email).DisplayName(displayName).Username(username).ExcludeUserId(excludeUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.CheckUserExists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckUserExists`: UserExists
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.CheckUserExists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckUserExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Filter by email. | 
 **displayName** | **string** | Filter by displayName. | 
 **username** | **string** | Filter by Username. | 
 **excludeUserId** | **string** | Exclude by UserID. | 

### Return type

[**UserExists**](UserExists.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfirmEmail

> ConfirmEmail(ctx).Id(id).VerifyEmail(verifyEmail).Execute()

Confirm Email



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
	id := "usr_00000000-0000-0000-0000-000000000000" // string | Target user for which to verify email.
	verifyEmail := "eml_00000000-0000-0000-0000-000000000000" // string | Token to verify email.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.ConfirmEmail(context.Background()).Id(id).VerifyEmail(verifyEmail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.ConfirmEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfirmEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Target user for which to verify email. | 
 **verifyEmail** | **string** | Token to verify email. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGlobalAvatarModeration

> AvatarModerationCreated CreateGlobalAvatarModeration(ctx).CreateAvatarModerationRequest(createAvatarModerationRequest).Execute()

Create Global Avatar Moderation



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
	createAvatarModerationRequest := *openapiclient.NewCreateAvatarModerationRequest(openapiclient.AvatarModerationType("block"), "avtr_912d66a4-4714-43b8-8407-7de2cafbf55b") // CreateAvatarModerationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.CreateGlobalAvatarModeration(context.Background()).CreateAvatarModerationRequest(createAvatarModerationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.CreateGlobalAvatarModeration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGlobalAvatarModeration`: AvatarModerationCreated
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.CreateGlobalAvatarModeration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGlobalAvatarModerationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAvatarModerationRequest** | [**CreateAvatarModerationRequest**](CreateAvatarModerationRequest.md) |  | 

### Return type

[**AvatarModerationCreated**](AvatarModerationCreated.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGlobalAvatarModeration

> OkStatus2 DeleteGlobalAvatarModeration(ctx).TargetAvatarId(targetAvatarId).AvatarModerationType(avatarModerationType).Execute()

Delete Global Avatar Moderation



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
	targetAvatarId := "targetAvatarId_example" // string | Must be a valid avatar ID.
	avatarModerationType := openapiclient.AvatarModerationType("block") // AvatarModerationType | The avatar moderation type associated with the avatar.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.DeleteGlobalAvatarModeration(context.Background()).TargetAvatarId(targetAvatarId).AvatarModerationType(avatarModerationType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.DeleteGlobalAvatarModeration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGlobalAvatarModeration`: OkStatus2
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.DeleteGlobalAvatarModeration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGlobalAvatarModerationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **targetAvatarId** | **string** | Must be a valid avatar ID. | 
 **avatarModerationType** | [**AvatarModerationType**](AvatarModerationType.md) | The avatar moderation type associated with the avatar. | 

### Return type

[**OkStatus2**](OkStatus2.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteModerationReport

> SuccessFlag DeleteModerationReport(ctx, moderationReportId).Execute()

Delete Moderation Report



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
	moderationReportId := "moderationReportId_example" // string | The moderation report id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.DeleteModerationReport(context.Background(), moderationReportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.DeleteModerationReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteModerationReport`: SuccessFlag
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.DeleteModerationReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moderationReportId** | **string** | The moderation report id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteModerationReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SuccessFlag**](SuccessFlag.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> CurrentUser DeleteUser(ctx, userId).Execute()

Delete User



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
	resp, r, err := apiClient.AuthenticationAPI.DeleteUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUser`: CurrentUser
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.DeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


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


## Disable2FA

> Disable2FAResult Disable2FA(ctx).Execute()

Disable 2FA



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
	resp, r, err := apiClient.AuthenticationAPI.Disable2FA(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.Disable2FA``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Disable2FA`: Disable2FAResult
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.Disable2FA`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDisable2FARequest struct via the builder pattern


### Return type

[**Disable2FAResult**](Disable2FAResult.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Enable2FA

> Pending2FAResult Enable2FA(ctx).Execute()

Enable time-based 2FA codes



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
	resp, r, err := apiClient.AuthenticationAPI.Enable2FA(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.Enable2FA``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Enable2FA`: Pending2FAResult
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.Enable2FA`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEnable2FARequest struct via the builder pattern


### Return type

[**Pending2FAResult**](Pending2FAResult.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUser

> RegisterUserAccount200Response GetCurrentUser(ctx).Execute()

Login and/or Get Current User Info



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
	resp, r, err := apiClient.AuthenticationAPI.GetCurrentUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetCurrentUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentUser`: RegisterUserAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetCurrentUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserRequest struct via the builder pattern


### Return type

[**RegisterUserAccount200Response**](RegisterUserAccount200Response.md)

### Authorization

[authHeader](../README.md#authHeader), [twoFactorAuthCookie](../README.md#twoFactorAuthCookie), [authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalAvatarModerations

> []AvatarModeration GetGlobalAvatarModerations(ctx).Execute()

Get Global Avatar Moderations



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
	resp, r, err := apiClient.AuthenticationAPI.GetGlobalAvatarModerations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetGlobalAvatarModerations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGlobalAvatarModerations`: []AvatarModeration
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetGlobalAvatarModerations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalAvatarModerationsRequest struct via the builder pattern


### Return type

[**[]AvatarModeration**](AvatarModeration.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModerationReports

> PaginatedModerationReportList GetModerationReports(ctx).Offset(offset).N(n).ReportingUserId(reportingUserId).Status(status).Type_(type_).Execute()

Get Moderation Reports



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
	offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)
	n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
	reportingUserId := "reportingUserId_example" // string | Filter for moderation reports. (optional)
	status := "status_example" // string | Filter for moderation reports. One of: `closed`... (optional)
	type_ := "type__example" // string | Filter for moderation reports. One of: `avatar`, `group`, `user`, `world`... (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.GetModerationReports(context.Background()).Offset(offset).N(n).ReportingUserId(reportingUserId).Status(status).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetModerationReports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModerationReports`: PaginatedModerationReportList
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetModerationReports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetModerationReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 
 **n** | **int32** | The number of objects to return. | [default to 60]
 **reportingUserId** | **string** | Filter for moderation reports. | 
 **status** | **string** | Filter for moderation reports. One of: &#x60;closed&#x60;... | 
 **type_** | **string** | Filter for moderation reports. One of: &#x60;avatar&#x60;, &#x60;group&#x60;, &#x60;user&#x60;, &#x60;world&#x60;... | 

### Return type

[**PaginatedModerationReportList**](PaginatedModerationReportList.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecoveryCodes

> TwoFactorRecoveryCodes GetRecoveryCodes(ctx).Execute()

Get 2FA Recovery codes



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
	resp, r, err := apiClient.AuthenticationAPI.GetRecoveryCodes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.GetRecoveryCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecoveryCodes`: TwoFactorRecoveryCodes
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.GetRecoveryCodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecoveryCodesRequest struct via the builder pattern


### Return type

[**TwoFactorRecoveryCodes**](TwoFactorRecoveryCodes.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Logout

> Success Logout(ctx).Execute()

Logout



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
	resp, r, err := apiClient.AuthenticationAPI.Logout(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.Logout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Logout`: Success
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.Logout`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLogoutRequest struct via the builder pattern


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


## RegisterUserAccount

> RegisterUserAccount200Response RegisterUserAccount(ctx).RegisterUserAccountRequest(registerUserAccountRequest).Execute()

Register User Account



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
	registerUserAccountRequest := *openapiclient.NewRegisterUserAccountRequest(int32(123), "CaptchaCode_example", "Day_example", "Email_example", "Month_example", "Password_example", false, "Username_example", "Year_example") // RegisterUserAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.RegisterUserAccount(context.Background()).RegisterUserAccountRequest(registerUserAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.RegisterUserAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterUserAccount`: RegisterUserAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.RegisterUserAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterUserAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerUserAccountRequest** | [**RegisterUserAccountRequest**](RegisterUserAccountRequest.md) |  | 

### Return type

[**RegisterUserAccount200Response**](RegisterUserAccount200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResendEmailConfirmation

> Success ResendEmailConfirmation(ctx).Execute()

Resend Email Confirmation



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
	resp, r, err := apiClient.AuthenticationAPI.ResendEmailConfirmation(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.ResendEmailConfirmation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResendEmailConfirmation`: Success
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.ResendEmailConfirmation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiResendEmailConfirmationRequest struct via the builder pattern


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


## SubmitModerationReport

> ModerationReport SubmitModerationReport(ctx).SubmitModerationReportRequest(submitModerationReportRequest).Execute()

Submit Moderation Report



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
	submitModerationReportRequest := *openapiclient.NewSubmitModerationReportRequest("behavior", "usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469", "child", "user") // SubmitModerationReportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.SubmitModerationReport(context.Background()).SubmitModerationReportRequest(submitModerationReportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.SubmitModerationReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitModerationReport`: ModerationReport
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.SubmitModerationReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitModerationReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitModerationReportRequest** | [**SubmitModerationReportRequest**](SubmitModerationReportRequest.md) |  | 

### Return type

[**ModerationReport**](ModerationReport.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Verify2FA

> Verify2FAResult Verify2FA(ctx).TwoFactorAuthCode(twoFactorAuthCode).Execute()

Verify 2FA code



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
	twoFactorAuthCode := *openapiclient.NewTwoFactorAuthCode("Code_example") // TwoFactorAuthCode | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.Verify2FA(context.Background()).TwoFactorAuthCode(twoFactorAuthCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.Verify2FA``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Verify2FA`: Verify2FAResult
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.Verify2FA`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerify2FARequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **twoFactorAuthCode** | [**TwoFactorAuthCode**](TwoFactorAuthCode.md) |  | 

### Return type

[**Verify2FAResult**](Verify2FAResult.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Verify2FAEmailCode

> Verify2FAEmailCodeResult Verify2FAEmailCode(ctx).TwoFactorEmailCode(twoFactorEmailCode).Execute()

Verify 2FA email code



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
	twoFactorEmailCode := *openapiclient.NewTwoFactorEmailCode("Code_example") // TwoFactorEmailCode | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.Verify2FAEmailCode(context.Background()).TwoFactorEmailCode(twoFactorEmailCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.Verify2FAEmailCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Verify2FAEmailCode`: Verify2FAEmailCodeResult
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.Verify2FAEmailCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerify2FAEmailCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **twoFactorEmailCode** | [**TwoFactorEmailCode**](TwoFactorEmailCode.md) |  | 

### Return type

[**Verify2FAEmailCodeResult**](Verify2FAEmailCodeResult.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyAuthToken

> VerifyAuthTokenResult VerifyAuthToken(ctx).Execute()

Verify Auth Token



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
	resp, r, err := apiClient.AuthenticationAPI.VerifyAuthToken(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.VerifyAuthToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyAuthToken`: VerifyAuthTokenResult
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.VerifyAuthToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyAuthTokenRequest struct via the builder pattern


### Return type

[**VerifyAuthTokenResult**](VerifyAuthTokenResult.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyLoginPlace

> VerifyLoginPlace(ctx).Token(token).UserId(userId).Execute()

Verify Login Place



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
	token := "token_example" // string | Token to verify login attempt.
	userId := "userId_example" // string | Filter by UserID. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPI.VerifyLoginPlace(context.Background()).Token(token).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.VerifyLoginPlace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyLoginPlaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | Token to verify login attempt. | 
 **userId** | **string** | Filter by UserID. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyPending2FA

> Verify2FAResult VerifyPending2FA(ctx).TwoFactorAuthCode(twoFactorAuthCode).Execute()

Verify Pending 2FA code



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
	twoFactorAuthCode := *openapiclient.NewTwoFactorAuthCode("Code_example") // TwoFactorAuthCode | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.VerifyPending2FA(context.Background()).TwoFactorAuthCode(twoFactorAuthCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.VerifyPending2FA``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyPending2FA`: Verify2FAResult
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.VerifyPending2FA`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyPending2FARequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **twoFactorAuthCode** | [**TwoFactorAuthCode**](TwoFactorAuthCode.md) |  | 

### Return type

[**Verify2FAResult**](Verify2FAResult.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyRecoveryCode

> Verify2FAResult VerifyRecoveryCode(ctx).TwoFactorAuthCode(twoFactorAuthCode).Execute()

Verify 2FA code with Recovery code



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
	twoFactorAuthCode := *openapiclient.NewTwoFactorAuthCode("Code_example") // TwoFactorAuthCode | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPI.VerifyRecoveryCode(context.Background()).TwoFactorAuthCode(twoFactorAuthCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.VerifyRecoveryCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyRecoveryCode`: Verify2FAResult
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.VerifyRecoveryCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyRecoveryCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **twoFactorAuthCode** | [**TwoFactorAuthCode**](TwoFactorAuthCode.md) |  | 

### Return type

[**Verify2FAResult**](Verify2FAResult.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

