# \EconomyAPI

All URIs are relative to *https://api.vrchat.cloud/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetActiveLicenses**](EconomyAPI.md#GetActiveLicenses) | **Get** /economy/licenses/active | Get Active Licenses
[**GetBalance**](EconomyAPI.md#GetBalance) | **Get** /user/{userId}/balance | Get Balance
[**GetBalanceEarnings**](EconomyAPI.md#GetBalanceEarnings) | **Get** /user/{userId}/balance/earnings | Get Balance Earnings
[**GetBulkGiftPurchases**](EconomyAPI.md#GetBulkGiftPurchases) | **Get** /user/bulk/gift/purchases | Get Bulk Gift Purchases
[**GetCurrentSubscriptions**](EconomyAPI.md#GetCurrentSubscriptions) | **Get** /auth/user/subscription | Get Current Subscriptions
[**GetEconomyAccount**](EconomyAPI.md#GetEconomyAccount) | **Get** /user/{userId}/economy/account | Get Economy Account
[**GetLicenseGroup**](EconomyAPI.md#GetLicenseGroup) | **Get** /licenseGroups/{licenseGroupId} | Get License Group
[**GetProductListing**](EconomyAPI.md#GetProductListing) | **Get** /listing/{productId} | Get Product Listing
[**GetProductListingAlternate**](EconomyAPI.md#GetProductListingAlternate) | **Get** /products/{productId} | Get Product Listing (alternate)
[**GetProductListings**](EconomyAPI.md#GetProductListings) | **Get** /user/{userId}/listings | Get User Product Listings
[**GetProductPurchases**](EconomyAPI.md#GetProductPurchases) | **Get** /economy/purchases | Get Product Purchases
[**GetRecentSubscription**](EconomyAPI.md#GetRecentSubscription) | **Get** /user/subscription/recent | Get Recent Subscription
[**GetSteamTransaction**](EconomyAPI.md#GetSteamTransaction) | **Get** /Steam/transactions/{transactionId} | Get Steam Transaction
[**GetSteamTransactions**](EconomyAPI.md#GetSteamTransactions) | **Get** /Steam/transactions | List Steam Transactions
[**GetStore**](EconomyAPI.md#GetStore) | **Get** /economy/store | Get Store
[**GetStoreShelves**](EconomyAPI.md#GetStoreShelves) | **Get** /economy/store/shelves | Get Store Shelves
[**GetSubscriptions**](EconomyAPI.md#GetSubscriptions) | **Get** /subscriptions | List Subscriptions
[**GetTiliaStatus**](EconomyAPI.md#GetTiliaStatus) | **Get** /tilia/status | Get Tilia Status
[**GetTiliaTos**](EconomyAPI.md#GetTiliaTos) | **Get** /user/{userId}/tilia/tos | Get Tilia TOS Agreement Status
[**GetTokenBundles**](EconomyAPI.md#GetTokenBundles) | **Get** /tokenBundles | List Token Bundles
[**GetUserCreditsEligible**](EconomyAPI.md#GetUserCreditsEligible) | **Get** /users/{userId}/credits/eligible | Get User Credits Eligiblity
[**GetUserSubscriptionEligible**](EconomyAPI.md#GetUserSubscriptionEligible) | **Get** /users/{userId}/subscription/eligible | Get User Subscription Eligiblity
[**PurchaseProductListing**](EconomyAPI.md#PurchaseProductListing) | **Post** /economy/purchase/listing | Purchase Product Listing
[**UpdateTiliaTos**](EconomyAPI.md#UpdateTiliaTos) | **Put** /user/{userId}/tilia/tos | Update Tilia TOS Agreement Status



## GetActiveLicenses

> []License GetActiveLicenses(ctx).Execute()

Get Active Licenses



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
	resp, r, err := apiClient.EconomyAPI.GetActiveLicenses(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EconomyAPI.GetActiveLicenses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActiveLicenses`: []License
	fmt.Fprintf(os.Stdout, "Response from `EconomyAPI.GetActiveLicenses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveLicensesRequest struct via the builder pattern


### Return type

[**[]License**](License.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBalance

> Balance GetBalance(ctx, userId).Execute()

Get Balance



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
	resp, r, err := apiClient.EconomyAPI.GetBalance(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EconomyAPI.GetBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBalance`: Balance
	fmt.Fprintf(os.Stdout, "Response from `EconomyAPI.GetBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Balance**](Balance.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBalanceEarnings

> Balance GetBalanceEarnings(ctx, userId).Execute()

Get Balance Earnings



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
	resp, r, err := apiClient.EconomyAPI.GetBalanceEarnings(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EconomyAPI.GetBalanceEarnings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBalanceEarnings`: Balance
	fmt.Fprintf(os.Stdout, "Response from `EconomyAPI.GetBalanceEarnings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBalanceEarningsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Balance**](Balance.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBulkGiftPurchases

> []map[string]interface{} GetBulkGiftPurchases(ctx).MostRecent(mostRecent).Execute()

Get Bulk Gift Purchases



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
	mostRecent := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EconomyAPI.GetBulkGiftPurchases(context.Background()).MostRecent(mostRecent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EconomyAPI.GetBulkGiftPurchases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBulkGiftPurchases`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EconomyAPI.GetBulkGiftPurchases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBulkGiftPurchasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mostRecent** | **bool** |  | 

### Return type

**[]map[string]interface{}**

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentSubscriptions

> []UserSubscription GetCurrentSubscriptions(ctx).Execute()

Get Current Subscriptions



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
	resp, r, err := apiClient.EconomyAPI.GetCurrentSubscriptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EconomyAPI.GetCurrentSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentSubscriptions`: []UserSubscription
	fmt.Fprintf(os.Stdout, "Response from `EconomyAPI.GetCurrentSubscriptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentSubscriptionsRequest struct via the builder pattern


### Return type

[**[]UserSubscription**](UserSubscription.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEconomyAccount

> EconomyAccount GetEconomyAccount(ctx, userId).Execute()

Get Economy Account



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
	resp, r, err := apiClient.EconomyAPI.GetEconomyAccount(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EconomyAPI.GetEconomyAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEconomyAccount`: EconomyAccount
	fmt.Fprintf(os.Stdout, "Response from `EconomyAPI.GetEconomyAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEconomyAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EconomyAccount**](EconomyAccount.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseGroup

> LicenseGroup GetLicenseGroup(ctx, licenseGroupId).Execute()

Get License Group



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
	licenseGroupId := "licenseGroupId_example" // string | Must be a valid license group ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EconomyAPI.GetLicenseGroup(context.Background(), licenseGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EconomyAPI.GetLicenseGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLicenseGroup`: LicenseGroup
	fmt.Fprintf(os.Stdout, "Response from `EconomyAPI.GetLicenseGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**licenseGroupId** | **string** | Must be a valid license group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LicenseGroup**](LicenseGroup.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductListing

> ProductListing GetProductListing(ctx, productId).Hydrate(hydrate).Execute()

Get Product Listing



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
	productId := "productId_example" // string | Must be a valid product ID.
	hydrate := true // bool | Populates some fields and changes types of others for certain objects. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EconomyAPI.GetProductListing(context.Background(), productId).Hydrate(hydrate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EconomyAPI.GetProductListing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProductListing`: ProductListing
	fmt.Fprintf(os.Stdout, "Response from `EconomyAPI.GetProductListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | Must be a valid product ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hydrate** | **bool** | Populates some fields and changes types of others for certain objects. | 

### Return type

[**ProductListing**](ProductListing.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductListingAlternate

> ProductListing GetProductListingAlternate(ctx, productId).Execute()

Get Product Listing (alternate)



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
	productId := "productId_example" // string | Must be a valid product ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EconomyAPI.GetProductListingAlternate(context.Background(), productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EconomyAPI.GetProductListingAlternate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProductListingAlternate`: ProductListing
	fmt.Fprintf(os.Stdout, "Response from `EconomyAPI.GetProductListingAlternate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | Must be a valid product ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductListingAlternateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProductListing**](ProductListing.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductListings

> []ProductListing GetProductListings(ctx, userId).N(n).Offset(offset).Hydrate(hydrate).GroupId(groupId).Active(active).Execute()

Get User Product Listings



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
	hydrate := true // bool | Populates some fields and changes types of others for certain objects. (optional)
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID. (optional)
	active := true // bool | Filter for users' listings and inventory bundles. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EconomyAPI.GetProductListings(context.Background(), userId).N(n).Offset(offset).Hydrate(hydrate).GroupId(groupId).Active(active).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EconomyAPI.GetProductListings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProductListings`: []ProductListing
	fmt.Fprintf(os.Stdout, "Response from `EconomyAPI.GetProductListings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductListingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 
 **hydrate** | **bool** | Populates some fields and changes types of others for certain objects. | 
 **groupId** | **string** | Must be a valid group ID. | 
 **active** | **bool** | Filter for users&#39; listings and inventory bundles. | 

### Return type

[**[]ProductListing**](ProductListing.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductPurchases

> []ProductPurchase GetProductPurchases(ctx).BuyerId(buyerId).N(n).Offset(offset).MostRecent(mostRecent).Sort(sort).Order(order).Execute()

Get Product Purchases



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
	buyerId := "buyerId_example" // string | Must be a valid user ID.
	n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
	offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)
	mostRecent := true // bool |  (optional)
	sort := openapiclient.SortOptionProductPurchase("purchaseDate") // SortOptionProductPurchase | The sort order of the results. (optional) (default to "purchaseDate")
	order := openapiclient.OrderOptionShort("asc") // OrderOptionShort | Result ordering (optional) (default to "desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EconomyAPI.GetProductPurchases(context.Background()).BuyerId(buyerId).N(n).Offset(offset).MostRecent(mostRecent).Sort(sort).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EconomyAPI.GetProductPurchases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProductPurchases`: []ProductPurchase
	fmt.Fprintf(os.Stdout, "Response from `EconomyAPI.GetProductPurchases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProductPurchasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buyerId** | **string** | Must be a valid user ID. | 
 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 
 **mostRecent** | **bool** |  | 
 **sort** | [**SortOptionProductPurchase**](SortOptionProductPurchase.md) | The sort order of the results. | [default to &quot;purchaseDate&quot;]
 **order** | [**OrderOptionShort**](OrderOptionShort.md) | Result ordering | [default to &quot;desc&quot;]

### Return type

[**[]ProductPurchase**](ProductPurchase.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecentSubscription

> UserSubscription GetRecentSubscription(ctx).Execute()

Get Recent Subscription



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
	resp, r, err := apiClient.EconomyAPI.GetRecentSubscription(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EconomyAPI.GetRecentSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecentSubscription`: UserSubscription
	fmt.Fprintf(os.Stdout, "Response from `EconomyAPI.GetRecentSubscription`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecentSubscriptionRequest struct via the builder pattern


### Return type

[**UserSubscription**](UserSubscription.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSteamTransaction

> Transaction GetSteamTransaction(ctx, transactionId).Execute()

Get Steam Transaction



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
	transactionId := "transactionId_example" // string | Must be a valid transaction ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EconomyAPI.GetSteamTransaction(context.Background(), transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EconomyAPI.GetSteamTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSteamTransaction`: Transaction
	fmt.Fprintf(os.Stdout, "Response from `EconomyAPI.GetSteamTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | Must be a valid transaction ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSteamTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Transaction**](Transaction.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSteamTransactions

> []Transaction GetSteamTransactions(ctx).Execute()

List Steam Transactions



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
	resp, r, err := apiClient.EconomyAPI.GetSteamTransactions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EconomyAPI.GetSteamTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSteamTransactions`: []Transaction
	fmt.Fprintf(os.Stdout, "Response from `EconomyAPI.GetSteamTransactions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSteamTransactionsRequest struct via the builder pattern


### Return type

[**[]Transaction**](Transaction.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStore

> Store GetStore(ctx).StoreId(storeId).HydrateListings(hydrateListings).HydrateProducts(hydrateProducts).Execute()

Get Store



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
	storeId := "storeId_example" // string | 
	hydrateListings := true // bool | Listings fields will be populated. (optional)
	hydrateProducts := true // bool | Products fields will be populated. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EconomyAPI.GetStore(context.Background()).StoreId(storeId).HydrateListings(hydrateListings).HydrateProducts(hydrateProducts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EconomyAPI.GetStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStore`: Store
	fmt.Fprintf(os.Stdout, "Response from `EconomyAPI.GetStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeId** | **string** |  | 
 **hydrateListings** | **bool** | Listings fields will be populated. | 
 **hydrateProducts** | **bool** | Products fields will be populated. | 

### Return type

[**Store**](Store.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStoreShelves

> []StoreShelf GetStoreShelves(ctx).StoreId(storeId).HydrateListings(hydrateListings).Fetch(fetch).Execute()

Get Store Shelves



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
	storeId := "storeId_example" // string | 
	hydrateListings := true // bool | Listings fields will be populated. (optional)
	fetch := openapiclient.StoreView("all") // StoreView |  (optional) (default to "public")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EconomyAPI.GetStoreShelves(context.Background()).StoreId(storeId).HydrateListings(hydrateListings).Fetch(fetch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EconomyAPI.GetStoreShelves``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStoreShelves`: []StoreShelf
	fmt.Fprintf(os.Stdout, "Response from `EconomyAPI.GetStoreShelves`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStoreShelvesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeId** | **string** |  | 
 **hydrateListings** | **bool** | Listings fields will be populated. | 
 **fetch** | [**StoreView**](StoreView.md) |  | [default to &quot;public&quot;]

### Return type

[**[]StoreShelf**](StoreShelf.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriptions

> []Subscription GetSubscriptions(ctx).Execute()

List Subscriptions



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
	resp, r, err := apiClient.EconomyAPI.GetSubscriptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EconomyAPI.GetSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubscriptions`: []Subscription
	fmt.Fprintf(os.Stdout, "Response from `EconomyAPI.GetSubscriptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionsRequest struct via the builder pattern


### Return type

[**[]Subscription**](Subscription.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTiliaStatus

> TiliaStatus GetTiliaStatus(ctx).Execute()

Get Tilia Status



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
	resp, r, err := apiClient.EconomyAPI.GetTiliaStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EconomyAPI.GetTiliaStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTiliaStatus`: TiliaStatus
	fmt.Fprintf(os.Stdout, "Response from `EconomyAPI.GetTiliaStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTiliaStatusRequest struct via the builder pattern


### Return type

[**TiliaStatus**](TiliaStatus.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTiliaTos

> TiliaTOS GetTiliaTos(ctx, userId).Execute()

Get Tilia TOS Agreement Status



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
	resp, r, err := apiClient.EconomyAPI.GetTiliaTos(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EconomyAPI.GetTiliaTos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTiliaTos`: TiliaTOS
	fmt.Fprintf(os.Stdout, "Response from `EconomyAPI.GetTiliaTos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTiliaTosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TiliaTOS**](TiliaTOS.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenBundles

> []TokenBundle GetTokenBundles(ctx).Execute()

List Token Bundles



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
	resp, r, err := apiClient.EconomyAPI.GetTokenBundles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EconomyAPI.GetTokenBundles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokenBundles`: []TokenBundle
	fmt.Fprintf(os.Stdout, "Response from `EconomyAPI.GetTokenBundles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenBundlesRequest struct via the builder pattern


### Return type

[**[]TokenBundle**](TokenBundle.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserCreditsEligible

> UserCreditsEligible GetUserCreditsEligible(ctx, userId).SubscriptionId(subscriptionId).Execute()

Get User Credits Eligiblity



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
	subscriptionId := "subscriptionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EconomyAPI.GetUserCreditsEligible(context.Background(), userId).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EconomyAPI.GetUserCreditsEligible``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserCreditsEligible`: UserCreditsEligible
	fmt.Fprintf(os.Stdout, "Response from `EconomyAPI.GetUserCreditsEligible`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserCreditsEligibleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionId** | **string** |  | 

### Return type

[**UserCreditsEligible**](UserCreditsEligible.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserSubscriptionEligible

> UserSubscriptionEligible GetUserSubscriptionEligible(ctx, userId).SteamId(steamId).Execute()

Get User Subscription Eligiblity



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
	steamId := "game night" // string | The Steam ID of the user. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EconomyAPI.GetUserSubscriptionEligible(context.Background(), userId).SteamId(steamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EconomyAPI.GetUserSubscriptionEligible``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserSubscriptionEligible`: UserSubscriptionEligible
	fmt.Fprintf(os.Stdout, "Response from `EconomyAPI.GetUserSubscriptionEligible`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSubscriptionEligibleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **steamId** | **string** | The Steam ID of the user. | 

### Return type

[**UserSubscriptionEligible**](UserSubscriptionEligible.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PurchaseProductListing

> ProductPurchase PurchaseProductListing(ctx).PurchaseProductListingRequest(purchaseProductListingRequest).Execute()

Purchase Product Listing



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
	purchaseProductListingRequest := *openapiclient.NewPurchaseProductListingRequest("prod_bfbc2315-247a-44d7-bfea-5237f8d56cb4", int32(123), int32(123)) // PurchaseProductListingRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EconomyAPI.PurchaseProductListing(context.Background()).PurchaseProductListingRequest(purchaseProductListingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EconomyAPI.PurchaseProductListing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PurchaseProductListing`: ProductPurchase
	fmt.Fprintf(os.Stdout, "Response from `EconomyAPI.PurchaseProductListing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPurchaseProductListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **purchaseProductListingRequest** | [**PurchaseProductListingRequest**](PurchaseProductListingRequest.md) |  | 

### Return type

[**ProductPurchase**](ProductPurchase.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTiliaTos

> map[string]interface{} UpdateTiliaTos(ctx, userId).UpdateTiliaTOSRequest(updateTiliaTOSRequest).Execute()

Update Tilia TOS Agreement Status



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
	updateTiliaTOSRequest := *openapiclient.NewUpdateTiliaTOSRequest(false) // UpdateTiliaTOSRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EconomyAPI.UpdateTiliaTos(context.Background(), userId).UpdateTiliaTOSRequest(updateTiliaTOSRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EconomyAPI.UpdateTiliaTos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTiliaTos`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EconomyAPI.UpdateTiliaTos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Must be a valid user ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTiliaTosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTiliaTOSRequest** | [**UpdateTiliaTOSRequest**](UpdateTiliaTOSRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

