# \CalendarAPI

All URIs are relative to *https://api.vrchat.cloud/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroupCalendarEvent**](CalendarAPI.md#CreateGroupCalendarEvent) | **Post** /calendar/{groupId}/event | Create a calendar event
[**DeleteGroupCalendarEvent**](CalendarAPI.md#DeleteGroupCalendarEvent) | **Delete** /calendar/{groupId}/{calendarId} | Delete a calendar event
[**DiscoverCalendarEvents**](CalendarAPI.md#DiscoverCalendarEvents) | **Get** /calendar/discover | Discover calendar events
[**FollowGroupCalendarEvent**](CalendarAPI.md#FollowGroupCalendarEvent) | **Post** /calendar/{groupId}/{calendarId}/follow | Follow a calendar event
[**GetCalendarEvents**](CalendarAPI.md#GetCalendarEvents) | **Get** /calendar | List calendar events
[**GetFeaturedCalendarEvents**](CalendarAPI.md#GetFeaturedCalendarEvents) | **Get** /calendar/featured | List featured calendar events
[**GetFollowedCalendarEvents**](CalendarAPI.md#GetFollowedCalendarEvents) | **Get** /calendar/following | List followed calendar events
[**GetGroupCalendarEvent**](CalendarAPI.md#GetGroupCalendarEvent) | **Get** /calendar/{groupId}/{calendarId} | Get a calendar event
[**GetGroupCalendarEventICS**](CalendarAPI.md#GetGroupCalendarEventICS) | **Get** /calendar/{groupId}/{calendarId}.ics | Download calendar event as ICS
[**GetGroupCalendarEvents**](CalendarAPI.md#GetGroupCalendarEvents) | **Get** /calendar/{groupId} | List a group&#39;s calendar events
[**GetGroupNextCalendarEvent**](CalendarAPI.md#GetGroupNextCalendarEvent) | **Get** /calendar/{groupId}/next | Get next calendar event
[**SearchCalendarEvents**](CalendarAPI.md#SearchCalendarEvents) | **Get** /calendar/search | Search for calendar events
[**UpdateGroupCalendarEvent**](CalendarAPI.md#UpdateGroupCalendarEvent) | **Put** /calendar/{groupId}/{calendarId}/event | Update a calendar event



## CreateGroupCalendarEvent

> CalendarEvent CreateGroupCalendarEvent(ctx, groupId).CreateCalendarEventRequest(createCalendarEventRequest).Execute()

Create a calendar event



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
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	createCalendarEventRequest := *openapiclient.NewCreateCalendarEventRequest(openapiclient.CalendarEventAccess("group"), openapiclient.CalendarEventCategory("arts"), "Description_example", time.Now(), false, time.Now(), "Performance Event!") // CreateCalendarEventRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CalendarAPI.CreateGroupCalendarEvent(context.Background(), groupId).CreateCalendarEventRequest(createCalendarEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CalendarAPI.CreateGroupCalendarEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGroupCalendarEvent`: CalendarEvent
	fmt.Fprintf(os.Stdout, "Response from `CalendarAPI.CreateGroupCalendarEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupCalendarEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createCalendarEventRequest** | [**CreateCalendarEventRequest**](CreateCalendarEventRequest.md) |  | 

### Return type

[**CalendarEvent**](CalendarEvent.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroupCalendarEvent

> Success DeleteGroupCalendarEvent(ctx, groupId, calendarId).Execute()

Delete a calendar event



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
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	calendarId := "cal_00000000-0000-0000-0000-000000000000" // string | Must be a valid calendar ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CalendarAPI.DeleteGroupCalendarEvent(context.Background(), groupId, calendarId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CalendarAPI.DeleteGroupCalendarEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGroupCalendarEvent`: Success
	fmt.Fprintf(os.Stdout, "Response from `CalendarAPI.DeleteGroupCalendarEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**calendarId** | **string** | Must be a valid calendar ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupCalendarEventRequest struct via the builder pattern


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


## DiscoverCalendarEvents

> CalendarEventDiscovery DiscoverCalendarEvents(ctx).Scope(scope).Categories(categories).Tags(tags).FeaturedResults(featuredResults).NonFeaturedResults(nonFeaturedResults).PersonalizedResults(personalizedResults).MinimumInterestCount(minimumInterestCount).MinimumRemainingMinutes(minimumRemainingMinutes).UpcomingOffsetMinutes(upcomingOffsetMinutes).N(n).NextCursor(nextCursor).Execute()

Discover calendar events



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
	scope := openapiclient.CalendarEventDiscoveryScope("all") // CalendarEventDiscoveryScope | Scope for calendar event discovery. (optional) (default to "upcoming")
	categories := "avatars,exploration,gaming,roleplaying,music,performance" // string | Filter for calendar event discovery. (optional)
	tags := "vrc_event_group_fair" // string | Filter for calendar event discovery. (optional)
	featuredResults := openapiclient.CalendarEventDiscoveryInclusion("exclude") // CalendarEventDiscoveryInclusion | Filter for calendar event discovery. (optional) (default to "include")
	nonFeaturedResults := openapiclient.CalendarEventDiscoveryInclusion("exclude") // CalendarEventDiscoveryInclusion | Filter for calendar event discovery. (optional) (default to "include")
	personalizedResults := openapiclient.CalendarEventDiscoveryInclusion("exclude") // CalendarEventDiscoveryInclusion | Filter for calendar event discovery. (optional) (default to "include")
	minimumInterestCount := int32(5) // int32 | Filter for calendar event discovery. (optional)
	minimumRemainingMinutes := int32(10) // int32 | Filter for calendar event discovery. (optional)
	upcomingOffsetMinutes := int32(10080) // int32 | Filter for calendar event discovery. (optional)
	n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
	nextCursor := "nextCursor_example" // string | Cursor returned from previous calendar discovery queries (see nextCursor property of the schema CalendarEventDiscovery). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CalendarAPI.DiscoverCalendarEvents(context.Background()).Scope(scope).Categories(categories).Tags(tags).FeaturedResults(featuredResults).NonFeaturedResults(nonFeaturedResults).PersonalizedResults(personalizedResults).MinimumInterestCount(minimumInterestCount).MinimumRemainingMinutes(minimumRemainingMinutes).UpcomingOffsetMinutes(upcomingOffsetMinutes).N(n).NextCursor(nextCursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CalendarAPI.DiscoverCalendarEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverCalendarEvents`: CalendarEventDiscovery
	fmt.Fprintf(os.Stdout, "Response from `CalendarAPI.DiscoverCalendarEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscoverCalendarEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | [**CalendarEventDiscoveryScope**](CalendarEventDiscoveryScope.md) | Scope for calendar event discovery. | [default to &quot;upcoming&quot;]
 **categories** | **string** | Filter for calendar event discovery. | 
 **tags** | **string** | Filter for calendar event discovery. | 
 **featuredResults** | [**CalendarEventDiscoveryInclusion**](CalendarEventDiscoveryInclusion.md) | Filter for calendar event discovery. | [default to &quot;include&quot;]
 **nonFeaturedResults** | [**CalendarEventDiscoveryInclusion**](CalendarEventDiscoveryInclusion.md) | Filter for calendar event discovery. | [default to &quot;include&quot;]
 **personalizedResults** | [**CalendarEventDiscoveryInclusion**](CalendarEventDiscoveryInclusion.md) | Filter for calendar event discovery. | [default to &quot;include&quot;]
 **minimumInterestCount** | **int32** | Filter for calendar event discovery. | 
 **minimumRemainingMinutes** | **int32** | Filter for calendar event discovery. | 
 **upcomingOffsetMinutes** | **int32** | Filter for calendar event discovery. | 
 **n** | **int32** | The number of objects to return. | [default to 60]
 **nextCursor** | **string** | Cursor returned from previous calendar discovery queries (see nextCursor property of the schema CalendarEventDiscovery). | 

### Return type

[**CalendarEventDiscovery**](CalendarEventDiscovery.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FollowGroupCalendarEvent

> CalendarEvent FollowGroupCalendarEvent(ctx, groupId, calendarId).FollowCalendarEventRequest(followCalendarEventRequest).Execute()

Follow a calendar event



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
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	calendarId := "cal_00000000-0000-0000-0000-000000000000" // string | Must be a valid calendar ID.
	followCalendarEventRequest := *openapiclient.NewFollowCalendarEventRequest(true) // FollowCalendarEventRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CalendarAPI.FollowGroupCalendarEvent(context.Background(), groupId, calendarId).FollowCalendarEventRequest(followCalendarEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CalendarAPI.FollowGroupCalendarEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FollowGroupCalendarEvent`: CalendarEvent
	fmt.Fprintf(os.Stdout, "Response from `CalendarAPI.FollowGroupCalendarEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**calendarId** | **string** | Must be a valid calendar ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFollowGroupCalendarEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **followCalendarEventRequest** | [**FollowCalendarEventRequest**](FollowCalendarEventRequest.md) |  | 

### Return type

[**CalendarEvent**](CalendarEvent.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCalendarEvents

> PaginatedCalendarEventList GetCalendarEvents(ctx).Date(date).N(n).Offset(offset).Execute()

List calendar events



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
	date := time.Now() // time.Time | The month to search in. (optional)
	n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
	offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CalendarAPI.GetCalendarEvents(context.Background()).Date(date).N(n).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CalendarAPI.GetCalendarEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCalendarEvents`: PaginatedCalendarEventList
	fmt.Fprintf(os.Stdout, "Response from `CalendarAPI.GetCalendarEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCalendarEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **time.Time** | The month to search in. | 
 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 

### Return type

[**PaginatedCalendarEventList**](PaginatedCalendarEventList.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeaturedCalendarEvents

> PaginatedCalendarEventList GetFeaturedCalendarEvents(ctx).Date(date).N(n).Offset(offset).Execute()

List featured calendar events



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
	date := time.Now() // time.Time | The month to search in. (optional)
	n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
	offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CalendarAPI.GetFeaturedCalendarEvents(context.Background()).Date(date).N(n).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CalendarAPI.GetFeaturedCalendarEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeaturedCalendarEvents`: PaginatedCalendarEventList
	fmt.Fprintf(os.Stdout, "Response from `CalendarAPI.GetFeaturedCalendarEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFeaturedCalendarEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **time.Time** | The month to search in. | 
 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 

### Return type

[**PaginatedCalendarEventList**](PaginatedCalendarEventList.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFollowedCalendarEvents

> PaginatedCalendarEventList GetFollowedCalendarEvents(ctx).Date(date).N(n).Offset(offset).Execute()

List followed calendar events



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
	date := time.Now() // time.Time | The month to search in. (optional)
	n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
	offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CalendarAPI.GetFollowedCalendarEvents(context.Background()).Date(date).N(n).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CalendarAPI.GetFollowedCalendarEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFollowedCalendarEvents`: PaginatedCalendarEventList
	fmt.Fprintf(os.Stdout, "Response from `CalendarAPI.GetFollowedCalendarEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFollowedCalendarEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **date** | **time.Time** | The month to search in. | 
 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 

### Return type

[**PaginatedCalendarEventList**](PaginatedCalendarEventList.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupCalendarEvent

> CalendarEvent GetGroupCalendarEvent(ctx, groupId, calendarId).Execute()

Get a calendar event



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
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	calendarId := "cal_00000000-0000-0000-0000-000000000000" // string | Must be a valid calendar ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CalendarAPI.GetGroupCalendarEvent(context.Background(), groupId, calendarId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CalendarAPI.GetGroupCalendarEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupCalendarEvent`: CalendarEvent
	fmt.Fprintf(os.Stdout, "Response from `CalendarAPI.GetGroupCalendarEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**calendarId** | **string** | Must be a valid calendar ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupCalendarEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CalendarEvent**](CalendarEvent.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupCalendarEventICS

> *os.File GetGroupCalendarEventICS(ctx, groupId, calendarId).Execute()

Download calendar event as ICS



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
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	calendarId := "cal_00000000-0000-0000-0000-000000000000" // string | Must be a valid calendar ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CalendarAPI.GetGroupCalendarEventICS(context.Background(), groupId, calendarId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CalendarAPI.GetGroupCalendarEventICS``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupCalendarEventICS`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `CalendarAPI.GetGroupCalendarEventICS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**calendarId** | **string** | Must be a valid calendar ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupCalendarEventICSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/calendar, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupCalendarEvents

> PaginatedCalendarEventList GetGroupCalendarEvents(ctx, groupId).Date(date).N(n).Offset(offset).Execute()

List a group's calendar events



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
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	date := time.Now() // time.Time | The month to search in. (optional)
	n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
	offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CalendarAPI.GetGroupCalendarEvents(context.Background(), groupId).Date(date).N(n).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CalendarAPI.GetGroupCalendarEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupCalendarEvents`: PaginatedCalendarEventList
	fmt.Fprintf(os.Stdout, "Response from `CalendarAPI.GetGroupCalendarEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupCalendarEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **date** | **time.Time** | The month to search in. | 
 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 

### Return type

[**PaginatedCalendarEventList**](PaginatedCalendarEventList.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupNextCalendarEvent

> CalendarEvent GetGroupNextCalendarEvent(ctx, groupId).Execute()

Get next calendar event



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
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CalendarAPI.GetGroupNextCalendarEvent(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CalendarAPI.GetGroupNextCalendarEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupNextCalendarEvent`: CalendarEvent
	fmt.Fprintf(os.Stdout, "Response from `CalendarAPI.GetGroupNextCalendarEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupNextCalendarEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CalendarEvent**](CalendarEvent.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCalendarEvents

> PaginatedCalendarEventList SearchCalendarEvents(ctx).SearchTerm(searchTerm).UtcOffset(utcOffset).N(n).Offset(offset).IsInternalVariant(isInternalVariant).Execute()

Search for calendar events



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
	searchTerm := "game night" // string | Search term for calendar events.
	utcOffset := int32(56) // int32 | The offset from UTC in hours of the client or authenticated user. (optional)
	n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
	offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)
	isInternalVariant := false // bool | Not quite sure what this actually does (exists on the website but doesn't seem to be used) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CalendarAPI.SearchCalendarEvents(context.Background()).SearchTerm(searchTerm).UtcOffset(utcOffset).N(n).Offset(offset).IsInternalVariant(isInternalVariant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CalendarAPI.SearchCalendarEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCalendarEvents`: PaginatedCalendarEventList
	fmt.Fprintf(os.Stdout, "Response from `CalendarAPI.SearchCalendarEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCalendarEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchTerm** | **string** | Search term for calendar events. | 
 **utcOffset** | **int32** | The offset from UTC in hours of the client or authenticated user. | 
 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 
 **isInternalVariant** | **bool** | Not quite sure what this actually does (exists on the website but doesn&#39;t seem to be used) | 

### Return type

[**PaginatedCalendarEventList**](PaginatedCalendarEventList.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupCalendarEvent

> CalendarEvent UpdateGroupCalendarEvent(ctx, groupId, calendarId).UpdateCalendarEventRequest(updateCalendarEventRequest).Execute()

Update a calendar event



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
	groupId := "grp_00000000-0000-0000-0000-000000000000" // string | Must be a valid group ID.
	calendarId := "cal_00000000-0000-0000-0000-000000000000" // string | Must be a valid calendar ID.
	updateCalendarEventRequest := *openapiclient.NewUpdateCalendarEventRequest() // UpdateCalendarEventRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CalendarAPI.UpdateGroupCalendarEvent(context.Background(), groupId, calendarId).UpdateCalendarEventRequest(updateCalendarEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CalendarAPI.UpdateGroupCalendarEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGroupCalendarEvent`: CalendarEvent
	fmt.Fprintf(os.Stdout, "Response from `CalendarAPI.UpdateGroupCalendarEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Must be a valid group ID. | 
**calendarId** | **string** | Must be a valid calendar ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupCalendarEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateCalendarEventRequest** | [**UpdateCalendarEventRequest**](UpdateCalendarEventRequest.md) |  | 

### Return type

[**CalendarEvent**](CalendarEvent.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

