/*
VRChat API Documentation

![VRChat API Banner](https://vrchatapi.github.io/assets/img/api_banner_1500x400.png)  # Welcome to the VRChat API  Before we begin, we would like to state this is a **COMMUNITY DRIVEN PROJECT**. This means that everything you read on here was written by the community itself and is **not** officially supported by VRChat. The documentation is provided \"AS IS\", and any action you take towards VRChat is completely your own responsibility.  The documentation and additional libraries SHALL ONLY be used for applications interacting with VRChat's API in accordance with their [Terms of Service](https://hello.vrchat.com/legal) and [Community Guidelines](https://hello.vrchat.com/community-guidelines), and MUST NOT be used for modifying the client, \"avatar ripping\", or other illegal activities. Malicious usage or spamming the API may result in account termination. Certain parts of the API are also more sensitive than others, for example moderation, so please tread extra carefully and read the warnings when present.  ![Tupper Policy on API](https://i.imgur.com/yLlW7Ok.png)  Finally, use of the API using applications other than the approved methods (website, VRChat application, Unity SDK) is not officially supported. VRChat provides no guarantee or support for external applications using the API. Access to API endpoints may break **at any time, without notice**. Therefore, please **do not ping** VRChat Staff in the VRChat Discord if you are having API problems, as they do not provide API support. We will make a best effort in keeping this documentation and associated language libraries up to date, but things might be outdated or missing. If you find that something is no longer valid, please contact us on Discord or [create an issue](https://github.com/vrchatapi/specification/issues) and tell us so we can fix it.  # Getting Started  The VRChat API can be used to programmatically retrieve or update information regarding your profile, friends, avatars, worlds and more. The API consists of two parts, \"Photon\" which is only used in-game, and the \"Web API\" which is used by both the game and the website. This documentation focuses only on the Web API.  The API is designed around the REST ideology, providing semi-simple and usually predictable URIs to access and modify objects. Requests support standard HTTP methods like GET, PUT, POST, and DELETE and standard status codes. Response bodies are always UTF-8 encoded JSON objects, unless explicitly documented otherwise.  <div class=\"callout callout-error\">   <strong>🛑 Warning! Do not touch Photon!</strong><br>   Photon is only used by the in-game client and should <b>not</b> be touched. Doing so may result in permanent account termination. </div>  <div class=\"callout callout-info\">   <strong>ℹ️ Authentication</strong><br>   Read <a href=\"#tag--authentication\">Authentication</a> for how to log in. </div>  # Using the API  For simply exploring what the API can do it is strongly recommended to download [Insomnia](https://insomnia.rest/download), a free and open-source API client that's great for sending requests to the API in an orderly fashion. Insomnia allows you to send data in the format that's required for VRChat's API. It is also possible to try out the API in your browser, by first logging in at [vrchat.com/home](https://vrchat.com/home/) and then going to [vrchat.com/api/1/auth/user](https://vrchat.com/api/1/auth/user), but the information will be much harder to work with.  For more permanent operation such as software development it is instead recommended to use one of the existing language SDKs. This community project maintains API libraries in several languages, which allows you to interact with the API with simple function calls rather than having to implement the HTTP protocol yourself. Most of these libraries are automatically generated from the API specification, sometimes with additional helpful wrapper code to make usage easier. This allows them to be almost automatically updated and expanded upon as soon as a new feature is introduced in the specification itself. The libraries can be found on [GitHub](https://github.com/vrchatapi) or following:  * [NodeJS (JavaScript)](https://www.npmjs.com/package/vrchat) * [Dart](https://pub.dev/packages/vrchat_dart) * [Rust](https://crates.io/crates/vrchatapi) * [C#](https://github.com/vrchatapi/vrchatapi-csharp) * [Python](https://github.com/vrchatapi/vrchatapi-python)  # Pagination  Most endpoints enforce pagination, meaning they will only return 10 entries by default, and never more than 100.<br> Using both the limit and offset parameters allows you to easily paginate through a large number of objects.  | Query Parameter | Type | Description | | ----------|--|------- | | `n` | integer  | The number of objects to return. This value often defaults to 10. Highest limit is always 100.| | `offset` | integer  | A zero-based offset from the default object sorting.|  If a request returns fewer objects than the `limit` parameter, there are no more items available to return.  # Contribution  Do you want to get involved in the documentation effort? Do you want to help improve one of the language API libraries? This project is an [OPEN Open Source Project](https://openopensource.org)! This means that individuals making significant and valuable contributions are given commit-access to the project. It also means we are very open and welcoming of new people making contributions, unlike some more guarded open-source projects.  [![Discord](https://img.shields.io/static/v1?label=vrchatapi&message=discord&color=blueviolet&style=for-the-badge)](https://discord.gg/qjZE9C9fkB)

API version: 1.20.7
Contact: vrchatapi.lpv0t@aries.fyi
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vrchatapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
	"os"
)


// CalendarAPIService CalendarAPI service
type CalendarAPIService service

type ApiCreateGroupCalendarEventRequest struct {
	ctx context.Context
	ApiService *CalendarAPIService
	groupId string
	createCalendarEventRequest *CreateCalendarEventRequest
}

func (r ApiCreateGroupCalendarEventRequest) CreateCalendarEventRequest(createCalendarEventRequest CreateCalendarEventRequest) ApiCreateGroupCalendarEventRequest {
	r.createCalendarEventRequest = &createCalendarEventRequest
	return r
}

func (r ApiCreateGroupCalendarEventRequest) Execute() (*CalendarEvent, *http.Response, error) {
	return r.ApiService.CreateGroupCalendarEventExecute(r)
}

/*
CreateGroupCalendarEvent Create a calendar event

Creates an event for a group on the calendar

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Must be a valid group ID.
 @return ApiCreateGroupCalendarEventRequest
*/
func (a *CalendarAPIService) CreateGroupCalendarEvent(ctx context.Context, groupId string) ApiCreateGroupCalendarEventRequest {
	return ApiCreateGroupCalendarEventRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return CalendarEvent
func (a *CalendarAPIService) CreateGroupCalendarEventExecute(r ApiCreateGroupCalendarEventRequest) (*CalendarEvent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CalendarEvent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarAPIService.CreateGroupCalendarEvent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/calendar/{groupId}/event"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createCalendarEventRequest == nil {
		return localVarReturnValue, nil, reportError("createCalendarEventRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createCalendarEventRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteGroupCalendarEventRequest struct {
	ctx context.Context
	ApiService *CalendarAPIService
	groupId string
	calendarId string
}

func (r ApiDeleteGroupCalendarEventRequest) Execute() (*Success, *http.Response, error) {
	return r.ApiService.DeleteGroupCalendarEventExecute(r)
}

/*
DeleteGroupCalendarEvent Delete a calendar event

Delete a group calendar event

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Must be a valid group ID.
 @param calendarId Must be a valid calendar ID.
 @return ApiDeleteGroupCalendarEventRequest
*/
func (a *CalendarAPIService) DeleteGroupCalendarEvent(ctx context.Context, groupId string, calendarId string) ApiDeleteGroupCalendarEventRequest {
	return ApiDeleteGroupCalendarEventRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		calendarId: calendarId,
	}
}

// Execute executes the request
//  @return Success
func (a *CalendarAPIService) DeleteGroupCalendarEventExecute(r ApiDeleteGroupCalendarEventRequest) (*Success, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Success
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarAPIService.DeleteGroupCalendarEvent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/calendar/{groupId}/{calendarId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"calendarId"+"}", url.PathEscape(parameterValueToString(r.calendarId, "calendarId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDiscoverCalendarEventsRequest struct {
	ctx context.Context
	ApiService *CalendarAPIService
	scope *CalendarEventDiscoveryScope
	categories *string
	tags *string
	featuredResults *CalendarEventDiscoveryInclusion
	nonFeaturedResults *CalendarEventDiscoveryInclusion
	personalizedResults *CalendarEventDiscoveryInclusion
	minimumInterestCount *int32
	minimumRemainingMinutes *int32
	upcomingOffsetMinutes *int32
	n *int32
	nextCursor *string
}

// Scope for calendar event discovery.
func (r ApiDiscoverCalendarEventsRequest) Scope(scope CalendarEventDiscoveryScope) ApiDiscoverCalendarEventsRequest {
	r.scope = &scope
	return r
}

// Filter for calendar event discovery.
func (r ApiDiscoverCalendarEventsRequest) Categories(categories string) ApiDiscoverCalendarEventsRequest {
	r.categories = &categories
	return r
}

// Filter for calendar event discovery.
func (r ApiDiscoverCalendarEventsRequest) Tags(tags string) ApiDiscoverCalendarEventsRequest {
	r.tags = &tags
	return r
}

// Filter for calendar event discovery.
func (r ApiDiscoverCalendarEventsRequest) FeaturedResults(featuredResults CalendarEventDiscoveryInclusion) ApiDiscoverCalendarEventsRequest {
	r.featuredResults = &featuredResults
	return r
}

// Filter for calendar event discovery.
func (r ApiDiscoverCalendarEventsRequest) NonFeaturedResults(nonFeaturedResults CalendarEventDiscoveryInclusion) ApiDiscoverCalendarEventsRequest {
	r.nonFeaturedResults = &nonFeaturedResults
	return r
}

// Filter for calendar event discovery.
func (r ApiDiscoverCalendarEventsRequest) PersonalizedResults(personalizedResults CalendarEventDiscoveryInclusion) ApiDiscoverCalendarEventsRequest {
	r.personalizedResults = &personalizedResults
	return r
}

// Filter for calendar event discovery.
func (r ApiDiscoverCalendarEventsRequest) MinimumInterestCount(minimumInterestCount int32) ApiDiscoverCalendarEventsRequest {
	r.minimumInterestCount = &minimumInterestCount
	return r
}

// Filter for calendar event discovery.
func (r ApiDiscoverCalendarEventsRequest) MinimumRemainingMinutes(minimumRemainingMinutes int32) ApiDiscoverCalendarEventsRequest {
	r.minimumRemainingMinutes = &minimumRemainingMinutes
	return r
}

// Filter for calendar event discovery.
func (r ApiDiscoverCalendarEventsRequest) UpcomingOffsetMinutes(upcomingOffsetMinutes int32) ApiDiscoverCalendarEventsRequest {
	r.upcomingOffsetMinutes = &upcomingOffsetMinutes
	return r
}

// The number of objects to return.
func (r ApiDiscoverCalendarEventsRequest) N(n int32) ApiDiscoverCalendarEventsRequest {
	r.n = &n
	return r
}

// Cursor returned from previous calendar discovery queries (see nextCursor property of the schema CalendarEventDiscovery).
func (r ApiDiscoverCalendarEventsRequest) NextCursor(nextCursor string) ApiDiscoverCalendarEventsRequest {
	r.nextCursor = &nextCursor
	return r
}

func (r ApiDiscoverCalendarEventsRequest) Execute() (*CalendarEventDiscovery, *http.Response, error) {
	return r.ApiService.DiscoverCalendarEventsExecute(r)
}

/*
DiscoverCalendarEvents Discover calendar events

Get a list of calendar events
Initially, call without a `nextCursor` parameter
For every successive call, use the `nextCursor` property returned in the previous call & the `number` of entries desired for this call
The `nextCursor` internally keeps track of the `offset` of the results, the initial request parameters, and accounts for discrepancies that might arise from time elapsed between calls

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDiscoverCalendarEventsRequest
*/
func (a *CalendarAPIService) DiscoverCalendarEvents(ctx context.Context) ApiDiscoverCalendarEventsRequest {
	return ApiDiscoverCalendarEventsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CalendarEventDiscovery
func (a *CalendarAPIService) DiscoverCalendarEventsExecute(r ApiDiscoverCalendarEventsRequest) (*CalendarEventDiscovery, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CalendarEventDiscovery
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarAPIService.DiscoverCalendarEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/calendar/discover"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.scope != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "scope", r.scope, "form", "")
	} else {
        var defaultValue CalendarEventDiscoveryScope = "upcoming"
        parameterAddToHeaderOrQuery(localVarQueryParams, "scope", defaultValue, "form", "")
        r.scope = &defaultValue
	}
	if r.categories != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "categories", r.categories, "form", "")
	}
	if r.tags != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tags", r.tags, "form", "")
	}
	if r.featuredResults != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "featuredResults", r.featuredResults, "form", "")
	} else {
        var defaultValue CalendarEventDiscoveryInclusion = "include"
        parameterAddToHeaderOrQuery(localVarQueryParams, "featuredResults", defaultValue, "form", "")
        r.featuredResults = &defaultValue
	}
	if r.nonFeaturedResults != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nonFeaturedResults", r.nonFeaturedResults, "form", "")
	} else {
        var defaultValue CalendarEventDiscoveryInclusion = "include"
        parameterAddToHeaderOrQuery(localVarQueryParams, "nonFeaturedResults", defaultValue, "form", "")
        r.nonFeaturedResults = &defaultValue
	}
	if r.personalizedResults != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "personalizedResults", r.personalizedResults, "form", "")
	} else {
        var defaultValue CalendarEventDiscoveryInclusion = "include"
        parameterAddToHeaderOrQuery(localVarQueryParams, "personalizedResults", defaultValue, "form", "")
        r.personalizedResults = &defaultValue
	}
	if r.minimumInterestCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "minimumInterestCount", r.minimumInterestCount, "form", "")
	}
	if r.minimumRemainingMinutes != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "minimumRemainingMinutes", r.minimumRemainingMinutes, "form", "")
	}
	if r.upcomingOffsetMinutes != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "upcomingOffsetMinutes", r.upcomingOffsetMinutes, "form", "")
	}
	if r.n != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "n", r.n, "form", "")
	} else {
        var defaultValue int32 = 60
        parameterAddToHeaderOrQuery(localVarQueryParams, "n", defaultValue, "form", "")
        r.n = &defaultValue
	}
	if r.nextCursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nextCursor", r.nextCursor, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiFollowGroupCalendarEventRequest struct {
	ctx context.Context
	ApiService *CalendarAPIService
	groupId string
	calendarId string
	followCalendarEventRequest *FollowCalendarEventRequest
}

func (r ApiFollowGroupCalendarEventRequest) FollowCalendarEventRequest(followCalendarEventRequest FollowCalendarEventRequest) ApiFollowGroupCalendarEventRequest {
	r.followCalendarEventRequest = &followCalendarEventRequest
	return r
}

func (r ApiFollowGroupCalendarEventRequest) Execute() (*CalendarEvent, *http.Response, error) {
	return r.ApiService.FollowGroupCalendarEventExecute(r)
}

/*
FollowGroupCalendarEvent Follow a calendar event

Follow or unfollow an event on a group's calendar

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Must be a valid group ID.
 @param calendarId Must be a valid calendar ID.
 @return ApiFollowGroupCalendarEventRequest
*/
func (a *CalendarAPIService) FollowGroupCalendarEvent(ctx context.Context, groupId string, calendarId string) ApiFollowGroupCalendarEventRequest {
	return ApiFollowGroupCalendarEventRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		calendarId: calendarId,
	}
}

// Execute executes the request
//  @return CalendarEvent
func (a *CalendarAPIService) FollowGroupCalendarEventExecute(r ApiFollowGroupCalendarEventRequest) (*CalendarEvent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CalendarEvent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarAPIService.FollowGroupCalendarEvent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/calendar/{groupId}/{calendarId}/follow"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"calendarId"+"}", url.PathEscape(parameterValueToString(r.calendarId, "calendarId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.followCalendarEventRequest == nil {
		return localVarReturnValue, nil, reportError("followCalendarEventRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.followCalendarEventRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCalendarEventsRequest struct {
	ctx context.Context
	ApiService *CalendarAPIService
	date *time.Time
	n *int32
	offset *int32
}

// The month to search in.
func (r ApiGetCalendarEventsRequest) Date(date time.Time) ApiGetCalendarEventsRequest {
	r.date = &date
	return r
}

// The number of objects to return.
func (r ApiGetCalendarEventsRequest) N(n int32) ApiGetCalendarEventsRequest {
	r.n = &n
	return r
}

// A zero-based offset from the default object sorting from where search results start.
func (r ApiGetCalendarEventsRequest) Offset(offset int32) ApiGetCalendarEventsRequest {
	r.offset = &offset
	return r
}

func (r ApiGetCalendarEventsRequest) Execute() (*PaginatedCalendarEventList, *http.Response, error) {
	return r.ApiService.GetCalendarEventsExecute(r)
}

/*
GetCalendarEvents List calendar events

Get a list of a user's calendar events for the month in ?date

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCalendarEventsRequest
*/
func (a *CalendarAPIService) GetCalendarEvents(ctx context.Context) ApiGetCalendarEventsRequest {
	return ApiGetCalendarEventsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedCalendarEventList
func (a *CalendarAPIService) GetCalendarEventsExecute(r ApiGetCalendarEventsRequest) (*PaginatedCalendarEventList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedCalendarEventList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarAPIService.GetCalendarEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/calendar"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.date != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "date", r.date, "form", "")
	}
	if r.n != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "n", r.n, "form", "")
	} else {
        var defaultValue int32 = 60
        parameterAddToHeaderOrQuery(localVarQueryParams, "n", defaultValue, "form", "")
        r.n = &defaultValue
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetFeaturedCalendarEventsRequest struct {
	ctx context.Context
	ApiService *CalendarAPIService
	date *time.Time
	n *int32
	offset *int32
}

// The month to search in.
func (r ApiGetFeaturedCalendarEventsRequest) Date(date time.Time) ApiGetFeaturedCalendarEventsRequest {
	r.date = &date
	return r
}

// The number of objects to return.
func (r ApiGetFeaturedCalendarEventsRequest) N(n int32) ApiGetFeaturedCalendarEventsRequest {
	r.n = &n
	return r
}

// A zero-based offset from the default object sorting from where search results start.
func (r ApiGetFeaturedCalendarEventsRequest) Offset(offset int32) ApiGetFeaturedCalendarEventsRequest {
	r.offset = &offset
	return r
}

func (r ApiGetFeaturedCalendarEventsRequest) Execute() (*PaginatedCalendarEventList, *http.Response, error) {
	return r.ApiService.GetFeaturedCalendarEventsExecute(r)
}

/*
GetFeaturedCalendarEvents List featured calendar events

Get a list of a featured calendar events for the month in ?date

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFeaturedCalendarEventsRequest
*/
func (a *CalendarAPIService) GetFeaturedCalendarEvents(ctx context.Context) ApiGetFeaturedCalendarEventsRequest {
	return ApiGetFeaturedCalendarEventsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedCalendarEventList
func (a *CalendarAPIService) GetFeaturedCalendarEventsExecute(r ApiGetFeaturedCalendarEventsRequest) (*PaginatedCalendarEventList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedCalendarEventList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarAPIService.GetFeaturedCalendarEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/calendar/featured"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.date != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "date", r.date, "form", "")
	}
	if r.n != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "n", r.n, "form", "")
	} else {
        var defaultValue int32 = 60
        parameterAddToHeaderOrQuery(localVarQueryParams, "n", defaultValue, "form", "")
        r.n = &defaultValue
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetFollowedCalendarEventsRequest struct {
	ctx context.Context
	ApiService *CalendarAPIService
	date *time.Time
	n *int32
	offset *int32
}

// The month to search in.
func (r ApiGetFollowedCalendarEventsRequest) Date(date time.Time) ApiGetFollowedCalendarEventsRequest {
	r.date = &date
	return r
}

// The number of objects to return.
func (r ApiGetFollowedCalendarEventsRequest) N(n int32) ApiGetFollowedCalendarEventsRequest {
	r.n = &n
	return r
}

// A zero-based offset from the default object sorting from where search results start.
func (r ApiGetFollowedCalendarEventsRequest) Offset(offset int32) ApiGetFollowedCalendarEventsRequest {
	r.offset = &offset
	return r
}

func (r ApiGetFollowedCalendarEventsRequest) Execute() (*PaginatedCalendarEventList, *http.Response, error) {
	return r.ApiService.GetFollowedCalendarEventsExecute(r)
}

/*
GetFollowedCalendarEvents List followed calendar events

Get a list of a followed calendar events for the month in ?date

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFollowedCalendarEventsRequest
*/
func (a *CalendarAPIService) GetFollowedCalendarEvents(ctx context.Context) ApiGetFollowedCalendarEventsRequest {
	return ApiGetFollowedCalendarEventsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedCalendarEventList
func (a *CalendarAPIService) GetFollowedCalendarEventsExecute(r ApiGetFollowedCalendarEventsRequest) (*PaginatedCalendarEventList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedCalendarEventList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarAPIService.GetFollowedCalendarEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/calendar/following"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.date != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "date", r.date, "form", "")
	}
	if r.n != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "n", r.n, "form", "")
	} else {
        var defaultValue int32 = 60
        parameterAddToHeaderOrQuery(localVarQueryParams, "n", defaultValue, "form", "")
        r.n = &defaultValue
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetGroupCalendarEventRequest struct {
	ctx context.Context
	ApiService *CalendarAPIService
	groupId string
	calendarId string
}

func (r ApiGetGroupCalendarEventRequest) Execute() (*CalendarEvent, *http.Response, error) {
	return r.ApiService.GetGroupCalendarEventExecute(r)
}

/*
GetGroupCalendarEvent Get a calendar event

Get a group calendar event

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Must be a valid group ID.
 @param calendarId Must be a valid calendar ID.
 @return ApiGetGroupCalendarEventRequest
*/
func (a *CalendarAPIService) GetGroupCalendarEvent(ctx context.Context, groupId string, calendarId string) ApiGetGroupCalendarEventRequest {
	return ApiGetGroupCalendarEventRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		calendarId: calendarId,
	}
}

// Execute executes the request
//  @return CalendarEvent
func (a *CalendarAPIService) GetGroupCalendarEventExecute(r ApiGetGroupCalendarEventRequest) (*CalendarEvent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CalendarEvent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarAPIService.GetGroupCalendarEvent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/calendar/{groupId}/{calendarId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"calendarId"+"}", url.PathEscape(parameterValueToString(r.calendarId, "calendarId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetGroupCalendarEventICSRequest struct {
	ctx context.Context
	ApiService *CalendarAPIService
	groupId string
	calendarId string
}

func (r ApiGetGroupCalendarEventICSRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.GetGroupCalendarEventICSExecute(r)
}

/*
GetGroupCalendarEventICS Download calendar event as ICS

Returns the specified calendar in iCalendar (ICS) format.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Must be a valid group ID.
 @param calendarId Must be a valid calendar ID.
 @return ApiGetGroupCalendarEventICSRequest
*/
func (a *CalendarAPIService) GetGroupCalendarEventICS(ctx context.Context, groupId string, calendarId string) ApiGetGroupCalendarEventICSRequest {
	return ApiGetGroupCalendarEventICSRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		calendarId: calendarId,
	}
}

// Execute executes the request
//  @return *os.File
func (a *CalendarAPIService) GetGroupCalendarEventICSExecute(r ApiGetGroupCalendarEventICSRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarAPIService.GetGroupCalendarEventICS")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/calendar/{groupId}/{calendarId}.ics"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"calendarId"+"}", url.PathEscape(parameterValueToString(r.calendarId, "calendarId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/calendar", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetGroupCalendarEventsRequest struct {
	ctx context.Context
	ApiService *CalendarAPIService
	groupId string
	date *time.Time
	n *int32
	offset *int32
}

// The month to search in.
func (r ApiGetGroupCalendarEventsRequest) Date(date time.Time) ApiGetGroupCalendarEventsRequest {
	r.date = &date
	return r
}

// The number of objects to return.
func (r ApiGetGroupCalendarEventsRequest) N(n int32) ApiGetGroupCalendarEventsRequest {
	r.n = &n
	return r
}

// A zero-based offset from the default object sorting from where search results start.
func (r ApiGetGroupCalendarEventsRequest) Offset(offset int32) ApiGetGroupCalendarEventsRequest {
	r.offset = &offset
	return r
}

func (r ApiGetGroupCalendarEventsRequest) Execute() (*PaginatedCalendarEventList, *http.Response, error) {
	return r.ApiService.GetGroupCalendarEventsExecute(r)
}

/*
GetGroupCalendarEvents List a group's calendar events

Get a list of a group's calendar events

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Must be a valid group ID.
 @return ApiGetGroupCalendarEventsRequest
*/
func (a *CalendarAPIService) GetGroupCalendarEvents(ctx context.Context, groupId string) ApiGetGroupCalendarEventsRequest {
	return ApiGetGroupCalendarEventsRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return PaginatedCalendarEventList
func (a *CalendarAPIService) GetGroupCalendarEventsExecute(r ApiGetGroupCalendarEventsRequest) (*PaginatedCalendarEventList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedCalendarEventList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarAPIService.GetGroupCalendarEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/calendar/{groupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.date != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "date", r.date, "form", "")
	}
	if r.n != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "n", r.n, "form", "")
	} else {
        var defaultValue int32 = 60
        parameterAddToHeaderOrQuery(localVarQueryParams, "n", defaultValue, "form", "")
        r.n = &defaultValue
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetGroupNextCalendarEventRequest struct {
	ctx context.Context
	ApiService *CalendarAPIService
	groupId string
}

func (r ApiGetGroupNextCalendarEventRequest) Execute() (*CalendarEvent, *http.Response, error) {
	return r.ApiService.GetGroupNextCalendarEventExecute(r)
}

/*
GetGroupNextCalendarEvent Get next calendar event

Get the closest future calendar event scheduled for a group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Must be a valid group ID.
 @return ApiGetGroupNextCalendarEventRequest
*/
func (a *CalendarAPIService) GetGroupNextCalendarEvent(ctx context.Context, groupId string) ApiGetGroupNextCalendarEventRequest {
	return ApiGetGroupNextCalendarEventRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return CalendarEvent
func (a *CalendarAPIService) GetGroupNextCalendarEventExecute(r ApiGetGroupNextCalendarEventRequest) (*CalendarEvent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CalendarEvent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarAPIService.GetGroupNextCalendarEvent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/calendar/{groupId}/next"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSearchCalendarEventsRequest struct {
	ctx context.Context
	ApiService *CalendarAPIService
	searchTerm *string
	utcOffset *int32
	n *int32
	offset *int32
	isInternalVariant *bool
}

// Search term for calendar events.
func (r ApiSearchCalendarEventsRequest) SearchTerm(searchTerm string) ApiSearchCalendarEventsRequest {
	r.searchTerm = &searchTerm
	return r
}

// The offset from UTC in hours of the client or authenticated user.
func (r ApiSearchCalendarEventsRequest) UtcOffset(utcOffset int32) ApiSearchCalendarEventsRequest {
	r.utcOffset = &utcOffset
	return r
}

// The number of objects to return.
func (r ApiSearchCalendarEventsRequest) N(n int32) ApiSearchCalendarEventsRequest {
	r.n = &n
	return r
}

// A zero-based offset from the default object sorting from where search results start.
func (r ApiSearchCalendarEventsRequest) Offset(offset int32) ApiSearchCalendarEventsRequest {
	r.offset = &offset
	return r
}

// Not quite sure what this actually does (exists on the website but doesn&#39;t seem to be used)
func (r ApiSearchCalendarEventsRequest) IsInternalVariant(isInternalVariant bool) ApiSearchCalendarEventsRequest {
	r.isInternalVariant = &isInternalVariant
	return r
}

func (r ApiSearchCalendarEventsRequest) Execute() (*PaginatedCalendarEventList, *http.Response, error) {
	return r.ApiService.SearchCalendarEventsExecute(r)
}

/*
SearchCalendarEvents Search for calendar events

Get a list of calendar events by search terms

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSearchCalendarEventsRequest
*/
func (a *CalendarAPIService) SearchCalendarEvents(ctx context.Context) ApiSearchCalendarEventsRequest {
	return ApiSearchCalendarEventsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedCalendarEventList
func (a *CalendarAPIService) SearchCalendarEventsExecute(r ApiSearchCalendarEventsRequest) (*PaginatedCalendarEventList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedCalendarEventList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarAPIService.SearchCalendarEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/calendar/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.searchTerm == nil {
		return localVarReturnValue, nil, reportError("searchTerm is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "searchTerm", r.searchTerm, "form", "")
	if r.utcOffset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "utcOffset", r.utcOffset, "form", "")
	}
	if r.n != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "n", r.n, "form", "")
	} else {
        var defaultValue int32 = 60
        parameterAddToHeaderOrQuery(localVarQueryParams, "n", defaultValue, "form", "")
        r.n = &defaultValue
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.isInternalVariant != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isInternalVariant", r.isInternalVariant, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateGroupCalendarEventRequest struct {
	ctx context.Context
	ApiService *CalendarAPIService
	groupId string
	calendarId string
	updateCalendarEventRequest *UpdateCalendarEventRequest
}

func (r ApiUpdateGroupCalendarEventRequest) UpdateCalendarEventRequest(updateCalendarEventRequest UpdateCalendarEventRequest) ApiUpdateGroupCalendarEventRequest {
	r.updateCalendarEventRequest = &updateCalendarEventRequest
	return r
}

func (r ApiUpdateGroupCalendarEventRequest) Execute() (*CalendarEvent, *http.Response, error) {
	return r.ApiService.UpdateGroupCalendarEventExecute(r)
}

/*
UpdateGroupCalendarEvent Update a calendar event

Updates an event for a group on the calendar

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Must be a valid group ID.
 @param calendarId Must be a valid calendar ID.
 @return ApiUpdateGroupCalendarEventRequest
*/
func (a *CalendarAPIService) UpdateGroupCalendarEvent(ctx context.Context, groupId string, calendarId string) ApiUpdateGroupCalendarEventRequest {
	return ApiUpdateGroupCalendarEventRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		calendarId: calendarId,
	}
}

// Execute executes the request
//  @return CalendarEvent
func (a *CalendarAPIService) UpdateGroupCalendarEventExecute(r ApiUpdateGroupCalendarEventRequest) (*CalendarEvent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CalendarEvent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalendarAPIService.UpdateGroupCalendarEvent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/calendar/{groupId}/{calendarId}/event"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"calendarId"+"}", url.PathEscape(parameterValueToString(r.calendarId, "calendarId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateCalendarEventRequest == nil {
		return localVarReturnValue, nil, reportError("updateCalendarEventRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateCalendarEventRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
