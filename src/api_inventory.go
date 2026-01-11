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
)


// InventoryAPIService InventoryAPI service
type InventoryAPIService service

type ApiConsumeOwnInventoryItemRequest struct {
	ctx context.Context
	ApiService *InventoryAPIService
	inventoryItemId string
}

func (r ApiConsumeOwnInventoryItemRequest) Execute() (*InventoryConsumptionResults, *http.Response, error) {
	return r.ApiService.ConsumeOwnInventoryItemExecute(r)
}

/*
ConsumeOwnInventoryItem Consume Own Inventory Item

Returns the modified InventoryItem object as held by the currently logged in user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param inventoryItemId Must be a valid inventory item ID.
 @return ApiConsumeOwnInventoryItemRequest
*/
func (a *InventoryAPIService) ConsumeOwnInventoryItem(ctx context.Context, inventoryItemId string) ApiConsumeOwnInventoryItemRequest {
	return ApiConsumeOwnInventoryItemRequest{
		ApiService: a,
		ctx: ctx,
		inventoryItemId: inventoryItemId,
	}
}

// Execute executes the request
//  @return InventoryConsumptionResults
func (a *InventoryAPIService) ConsumeOwnInventoryItemExecute(r ApiConsumeOwnInventoryItemRequest) (*InventoryConsumptionResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InventoryConsumptionResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryAPIService.ConsumeOwnInventoryItem")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory/{inventoryItemId}/consume"
	localVarPath = strings.Replace(localVarPath, "{"+"inventoryItemId"+"}", url.PathEscape(parameterValueToString(r.inventoryItemId, "inventoryItemId")), -1)

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

type ApiDeleteOwnInventoryItemRequest struct {
	ctx context.Context
	ApiService *InventoryAPIService
	inventoryItemId string
}

func (r ApiDeleteOwnInventoryItemRequest) Execute() (*SuccessFlag, *http.Response, error) {
	return r.ApiService.DeleteOwnInventoryItemExecute(r)
}

/*
DeleteOwnInventoryItem Delete Own Inventory Item

Deletes an InventoryItem from the inventory of the currently logged in user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param inventoryItemId Must be a valid inventory item ID.
 @return ApiDeleteOwnInventoryItemRequest
*/
func (a *InventoryAPIService) DeleteOwnInventoryItem(ctx context.Context, inventoryItemId string) ApiDeleteOwnInventoryItemRequest {
	return ApiDeleteOwnInventoryItemRequest{
		ApiService: a,
		ctx: ctx,
		inventoryItemId: inventoryItemId,
	}
}

// Execute executes the request
//  @return SuccessFlag
func (a *InventoryAPIService) DeleteOwnInventoryItemExecute(r ApiDeleteOwnInventoryItemRequest) (*SuccessFlag, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SuccessFlag
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryAPIService.DeleteOwnInventoryItem")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory/{inventoryItemId}"
	localVarPath = strings.Replace(localVarPath, "{"+"inventoryItemId"+"}", url.PathEscape(parameterValueToString(r.inventoryItemId, "inventoryItemId")), -1)

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

type ApiEquipOwnInventoryItemRequest struct {
	ctx context.Context
	ApiService *InventoryAPIService
	inventoryItemId string
	equipInventoryItemRequest *EquipInventoryItemRequest
}

func (r ApiEquipOwnInventoryItemRequest) EquipInventoryItemRequest(equipInventoryItemRequest EquipInventoryItemRequest) ApiEquipOwnInventoryItemRequest {
	r.equipInventoryItemRequest = &equipInventoryItemRequest
	return r
}

func (r ApiEquipOwnInventoryItemRequest) Execute() (*InventoryItem, *http.Response, error) {
	return r.ApiService.EquipOwnInventoryItemExecute(r)
}

/*
EquipOwnInventoryItem Equip Own Inventory Item

Returns the modified InventoryItem object as held by the currently logged in user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param inventoryItemId Must be a valid inventory item ID.
 @return ApiEquipOwnInventoryItemRequest
*/
func (a *InventoryAPIService) EquipOwnInventoryItem(ctx context.Context, inventoryItemId string) ApiEquipOwnInventoryItemRequest {
	return ApiEquipOwnInventoryItemRequest{
		ApiService: a,
		ctx: ctx,
		inventoryItemId: inventoryItemId,
	}
}

// Execute executes the request
//  @return InventoryItem
func (a *InventoryAPIService) EquipOwnInventoryItemExecute(r ApiEquipOwnInventoryItemRequest) (*InventoryItem, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InventoryItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryAPIService.EquipOwnInventoryItem")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory/{inventoryItemId}/equip"
	localVarPath = strings.Replace(localVarPath, "{"+"inventoryItemId"+"}", url.PathEscape(parameterValueToString(r.inventoryItemId, "inventoryItemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.equipInventoryItemRequest
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

type ApiGetInventoryRequest struct {
	ctx context.Context
	ApiService *InventoryAPIService
	n *int32
	offset *int32
	holderId *string
	equipSlot *InventoryEquipSlot
	order *string
	tags *string
	types *InventoryItemType
	flags *InventoryFlag
	notTypes *InventoryItemType
	notFlags *InventoryFlag
	archived *bool
}

// The number of objects to return.
func (r ApiGetInventoryRequest) N(n int32) ApiGetInventoryRequest {
	r.n = &n
	return r
}

// A zero-based offset from the default object sorting from where search results start.
func (r ApiGetInventoryRequest) Offset(offset int32) ApiGetInventoryRequest {
	r.offset = &offset
	return r
}

// The UserID of the owner of the inventory; defaults to the currently authenticated user.
func (r ApiGetInventoryRequest) HolderId(holderId string) ApiGetInventoryRequest {
	r.holderId = &holderId
	return r
}

// Filter for inventory retrieval.
func (r ApiGetInventoryRequest) EquipSlot(equipSlot InventoryEquipSlot) ApiGetInventoryRequest {
	r.equipSlot = &equipSlot
	return r
}

// Sort order for inventory retrieval.
func (r ApiGetInventoryRequest) Order(order string) ApiGetInventoryRequest {
	r.order = &order
	return r
}

// Filter tags for inventory retrieval (comma-separated).
func (r ApiGetInventoryRequest) Tags(tags string) ApiGetInventoryRequest {
	r.tags = &tags
	return r
}

// Filter for inventory retrieval.
func (r ApiGetInventoryRequest) Types(types InventoryItemType) ApiGetInventoryRequest {
	r.types = &types
	return r
}

// Filter flags for inventory retrieval (comma-separated).
func (r ApiGetInventoryRequest) Flags(flags InventoryFlag) ApiGetInventoryRequest {
	r.flags = &flags
	return r
}

// Filter out types for inventory retrieval (comma-separated).
func (r ApiGetInventoryRequest) NotTypes(notTypes InventoryItemType) ApiGetInventoryRequest {
	r.notTypes = &notTypes
	return r
}

// Filter out flags for inventory retrieval (comma-separated).
func (r ApiGetInventoryRequest) NotFlags(notFlags InventoryFlag) ApiGetInventoryRequest {
	r.notFlags = &notFlags
	return r
}

// Filter archived status for inventory retrieval.
func (r ApiGetInventoryRequest) Archived(archived bool) ApiGetInventoryRequest {
	r.archived = &archived
	return r
}

func (r ApiGetInventoryRequest) Execute() (*Inventory, *http.Response, error) {
	return r.ApiService.GetInventoryExecute(r)
}

/*
GetInventory Get Inventory

Returns an Inventory object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetInventoryRequest
*/
func (a *InventoryAPIService) GetInventory(ctx context.Context) ApiGetInventoryRequest {
	return ApiGetInventoryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Inventory
func (a *InventoryAPIService) GetInventoryExecute(r ApiGetInventoryRequest) (*Inventory, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Inventory
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryAPIService.GetInventory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	if r.holderId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "holderId", r.holderId, "form", "")
	}
	if r.equipSlot != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "equipSlot", r.equipSlot, "form", "")
	} else {
        var defaultValue InventoryEquipSlot = ""
        parameterAddToHeaderOrQuery(localVarQueryParams, "equipSlot", defaultValue, "form", "")
        r.equipSlot = &defaultValue
	}
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "form", "")
	}
	if r.tags != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tags", r.tags, "form", "")
	}
	if r.types != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "types", r.types, "form", "")
	} else {
        var defaultValue InventoryItemType = "bundle"
        parameterAddToHeaderOrQuery(localVarQueryParams, "types", defaultValue, "form", "")
        r.types = &defaultValue
	}
	if r.flags != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "flags", r.flags, "form", "")
	} else {
        var defaultValue InventoryFlag = "instantiatable"
        parameterAddToHeaderOrQuery(localVarQueryParams, "flags", defaultValue, "form", "")
        r.flags = &defaultValue
	}
	if r.notTypes != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "notTypes", r.notTypes, "form", "")
	} else {
        var defaultValue InventoryItemType = "bundle"
        parameterAddToHeaderOrQuery(localVarQueryParams, "notTypes", defaultValue, "form", "")
        r.notTypes = &defaultValue
	}
	if r.notFlags != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "notFlags", r.notFlags, "form", "")
	} else {
        var defaultValue InventoryFlag = "instantiatable"
        parameterAddToHeaderOrQuery(localVarQueryParams, "notFlags", defaultValue, "form", "")
        r.notFlags = &defaultValue
	}
	if r.archived != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "archived", r.archived, "form", "")
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

type ApiGetInventoryCollectionsRequest struct {
	ctx context.Context
	ApiService *InventoryAPIService
}

func (r ApiGetInventoryCollectionsRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.GetInventoryCollectionsExecute(r)
}

/*
GetInventoryCollections List Inventory Collections

Returns a list of collection names.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetInventoryCollectionsRequest
*/
func (a *InventoryAPIService) GetInventoryCollections(ctx context.Context) ApiGetInventoryCollectionsRequest {
	return ApiGetInventoryCollectionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []string
func (a *InventoryAPIService) GetInventoryCollectionsExecute(r ApiGetInventoryCollectionsRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryAPIService.GetInventoryCollections")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory/collections"

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

type ApiGetInventoryDropsRequest struct {
	ctx context.Context
	ApiService *InventoryAPIService
	active *bool
}

// Filter for users&#39; listings and inventory bundles.
func (r ApiGetInventoryDropsRequest) Active(active bool) ApiGetInventoryDropsRequest {
	r.active = &active
	return r
}

func (r ApiGetInventoryDropsRequest) Execute() ([]InventoryDrop, *http.Response, error) {
	return r.ApiService.GetInventoryDropsExecute(r)
}

/*
GetInventoryDrops List Inventory Drops

Returns a list of InventoryDrop objects.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetInventoryDropsRequest
*/
func (a *InventoryAPIService) GetInventoryDrops(ctx context.Context) ApiGetInventoryDropsRequest {
	return ApiGetInventoryDropsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []InventoryDrop
func (a *InventoryAPIService) GetInventoryDropsExecute(r ApiGetInventoryDropsRequest) ([]InventoryDrop, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []InventoryDrop
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryAPIService.GetInventoryDrops")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory/drops"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.active != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "active", r.active, "form", "")
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

type ApiGetInventoryTemplateRequest struct {
	ctx context.Context
	ApiService *InventoryAPIService
	inventoryTemplateId string
}

func (r ApiGetInventoryTemplateRequest) Execute() (*InventoryTemplate, *http.Response, error) {
	return r.ApiService.GetInventoryTemplateExecute(r)
}

/*
GetInventoryTemplate Get Inventory Template

Returns an InventoryTemplate object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param inventoryTemplateId Must be a valid inventory template ID.
 @return ApiGetInventoryTemplateRequest
*/
func (a *InventoryAPIService) GetInventoryTemplate(ctx context.Context, inventoryTemplateId string) ApiGetInventoryTemplateRequest {
	return ApiGetInventoryTemplateRequest{
		ApiService: a,
		ctx: ctx,
		inventoryTemplateId: inventoryTemplateId,
	}
}

// Execute executes the request
//  @return InventoryTemplate
func (a *InventoryAPIService) GetInventoryTemplateExecute(r ApiGetInventoryTemplateRequest) (*InventoryTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InventoryTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryAPIService.GetInventoryTemplate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory/template/{inventoryTemplateId}"
	localVarPath = strings.Replace(localVarPath, "{"+"inventoryTemplateId"+"}", url.PathEscape(parameterValueToString(r.inventoryTemplateId, "inventoryTemplateId")), -1)

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

type ApiGetOwnInventoryItemRequest struct {
	ctx context.Context
	ApiService *InventoryAPIService
	inventoryItemId string
}

func (r ApiGetOwnInventoryItemRequest) Execute() (*InventoryItem, *http.Response, error) {
	return r.ApiService.GetOwnInventoryItemExecute(r)
}

/*
GetOwnInventoryItem Get Own Inventory Item

Returns an InventoryItem object held by the currently logged in user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param inventoryItemId Must be a valid inventory item ID.
 @return ApiGetOwnInventoryItemRequest
*/
func (a *InventoryAPIService) GetOwnInventoryItem(ctx context.Context, inventoryItemId string) ApiGetOwnInventoryItemRequest {
	return ApiGetOwnInventoryItemRequest{
		ApiService: a,
		ctx: ctx,
		inventoryItemId: inventoryItemId,
	}
}

// Execute executes the request
//  @return InventoryItem
func (a *InventoryAPIService) GetOwnInventoryItemExecute(r ApiGetOwnInventoryItemRequest) (*InventoryItem, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InventoryItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryAPIService.GetOwnInventoryItem")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory/{inventoryItemId}"
	localVarPath = strings.Replace(localVarPath, "{"+"inventoryItemId"+"}", url.PathEscape(parameterValueToString(r.inventoryItemId, "inventoryItemId")), -1)

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

type ApiGetUserInventoryItemRequest struct {
	ctx context.Context
	ApiService *InventoryAPIService
	userId string
	inventoryItemId string
}

func (r ApiGetUserInventoryItemRequest) Execute() (*InventoryItem, *http.Response, error) {
	return r.ApiService.GetUserInventoryItemExecute(r)
}

/*
GetUserInventoryItem Get User Inventory Item

Returns an InventoryItem object held by the given user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId Must be a valid user ID.
 @param inventoryItemId Must be a valid inventory item ID.
 @return ApiGetUserInventoryItemRequest
*/
func (a *InventoryAPIService) GetUserInventoryItem(ctx context.Context, userId string, inventoryItemId string) ApiGetUserInventoryItemRequest {
	return ApiGetUserInventoryItemRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		inventoryItemId: inventoryItemId,
	}
}

// Execute executes the request
//  @return InventoryItem
func (a *InventoryAPIService) GetUserInventoryItemExecute(r ApiGetUserInventoryItemRequest) (*InventoryItem, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InventoryItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryAPIService.GetUserInventoryItem")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user/{userId}/inventory/{inventoryItemId}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"inventoryItemId"+"}", url.PathEscape(parameterValueToString(r.inventoryItemId, "inventoryItemId")), -1)

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

type ApiShareInventoryItemDirectRequest struct {
	ctx context.Context
	ApiService *InventoryAPIService
	itemId *string
	duration *int32
	shareInventoryItemDirectRequest *ShareInventoryItemDirectRequest
}

// Id for inventory item sharing.
func (r ApiShareInventoryItemDirectRequest) ItemId(itemId string) ApiShareInventoryItemDirectRequest {
	r.itemId = &itemId
	return r
}

// The duration before the sharing pedestal despawns.
func (r ApiShareInventoryItemDirectRequest) Duration(duration int32) ApiShareInventoryItemDirectRequest {
	r.duration = &duration
	return r
}

func (r ApiShareInventoryItemDirectRequest) ShareInventoryItemDirectRequest(shareInventoryItemDirectRequest ShareInventoryItemDirectRequest) ApiShareInventoryItemDirectRequest {
	r.shareInventoryItemDirectRequest = &shareInventoryItemDirectRequest
	return r
}

func (r ApiShareInventoryItemDirectRequest) Execute() (*OkStatus, *http.Response, error) {
	return r.ApiService.ShareInventoryItemDirectExecute(r)
}

/*
ShareInventoryItemDirect Share Inventory Item Direct

Share content directly with other users.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiShareInventoryItemDirectRequest
*/
func (a *InventoryAPIService) ShareInventoryItemDirect(ctx context.Context) ApiShareInventoryItemDirectRequest {
	return ApiShareInventoryItemDirectRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OkStatus
func (a *InventoryAPIService) ShareInventoryItemDirectExecute(r ApiShareInventoryItemDirectRequest) (*OkStatus, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OkStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryAPIService.ShareInventoryItemDirect")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory/cloning/direct"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.itemId == nil {
		return localVarReturnValue, nil, reportError("itemId is required and must be specified")
	}
	if r.duration == nil {
		return localVarReturnValue, nil, reportError("duration is required and must be specified")
	}
	if *r.duration < 30 {
		return localVarReturnValue, nil, reportError("duration must be greater than 30")
	}
	if *r.duration > 86400 {
		return localVarReturnValue, nil, reportError("duration must be less than 86400")
	}
	if r.shareInventoryItemDirectRequest == nil {
		return localVarReturnValue, nil, reportError("shareInventoryItemDirectRequest is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "itemId", r.itemId, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "duration", r.duration, "form", "")
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
	localVarPostBody = r.shareInventoryItemDirectRequest
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

type ApiShareInventoryItemPedestalRequest struct {
	ctx context.Context
	ApiService *InventoryAPIService
	itemId *string
	duration *int32
}

// Id for inventory item sharing.
func (r ApiShareInventoryItemPedestalRequest) ItemId(itemId string) ApiShareInventoryItemPedestalRequest {
	r.itemId = &itemId
	return r
}

// The duration before the sharing pedestal despawns.
func (r ApiShareInventoryItemPedestalRequest) Duration(duration int32) ApiShareInventoryItemPedestalRequest {
	r.duration = &duration
	return r
}

func (r ApiShareInventoryItemPedestalRequest) Execute() (*InventorySpawn, *http.Response, error) {
	return r.ApiService.ShareInventoryItemPedestalExecute(r)
}

/*
ShareInventoryItemPedestal Share Inventory Item by Pedestal

Returns an InventorySpawn object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiShareInventoryItemPedestalRequest
*/
func (a *InventoryAPIService) ShareInventoryItemPedestal(ctx context.Context) ApiShareInventoryItemPedestalRequest {
	return ApiShareInventoryItemPedestalRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return InventorySpawn
func (a *InventoryAPIService) ShareInventoryItemPedestalExecute(r ApiShareInventoryItemPedestalRequest) (*InventorySpawn, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InventorySpawn
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryAPIService.ShareInventoryItemPedestal")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory/cloning/pedestal"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.itemId == nil {
		return localVarReturnValue, nil, reportError("itemId is required and must be specified")
	}
	if r.duration == nil {
		return localVarReturnValue, nil, reportError("duration is required and must be specified")
	}
	if *r.duration < 30 {
		return localVarReturnValue, nil, reportError("duration must be greater than 30")
	}
	if *r.duration > 86400 {
		return localVarReturnValue, nil, reportError("duration must be less than 86400")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "itemId", r.itemId, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "duration", r.duration, "form", "")
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

type ApiSpawnInventoryItemRequest struct {
	ctx context.Context
	ApiService *InventoryAPIService
	id *string
}

// Id for inventory item spawning.
func (r ApiSpawnInventoryItemRequest) Id(id string) ApiSpawnInventoryItemRequest {
	r.id = &id
	return r
}

func (r ApiSpawnInventoryItemRequest) Execute() (*InventorySpawn, *http.Response, error) {
	return r.ApiService.SpawnInventoryItemExecute(r)
}

/*
SpawnInventoryItem Spawn Inventory Item

Returns an InventorySpawn object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSpawnInventoryItemRequest
*/
func (a *InventoryAPIService) SpawnInventoryItem(ctx context.Context) ApiSpawnInventoryItemRequest {
	return ApiSpawnInventoryItemRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return InventorySpawn
func (a *InventoryAPIService) SpawnInventoryItemExecute(r ApiSpawnInventoryItemRequest) (*InventorySpawn, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InventorySpawn
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryAPIService.SpawnInventoryItem")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory/spawn"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.id == nil {
		return localVarReturnValue, nil, reportError("id is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "form", "")
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

type ApiUnequipOwnInventorySlotRequest struct {
	ctx context.Context
	ApiService *InventoryAPIService
	inventoryItemId InventoryEquipSlot
}

func (r ApiUnequipOwnInventorySlotRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.UnequipOwnInventorySlotExecute(r)
}

/*
UnequipOwnInventorySlot Unequip Own Inventory Slot

Unequips the InventoryItem in the given slot of the inventory of the currently logged in user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param inventoryItemId Selector for inventory slot management.
 @return ApiUnequipOwnInventorySlotRequest
*/
func (a *InventoryAPIService) UnequipOwnInventorySlot(ctx context.Context, inventoryItemId InventoryEquipSlot) ApiUnequipOwnInventorySlotRequest {
	return ApiUnequipOwnInventorySlotRequest{
		ApiService: a,
		ctx: ctx,
		inventoryItemId: inventoryItemId,
	}
}

// Execute executes the request
//  @return string
func (a *InventoryAPIService) UnequipOwnInventorySlotExecute(r ApiUnequipOwnInventorySlotRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryAPIService.UnequipOwnInventorySlot")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory/{inventoryItemId}/equip"
	localVarPath = strings.Replace(localVarPath, "{"+"inventoryItemId"+"}", url.PathEscape(parameterValueToString(r.inventoryItemId, "inventoryItemId")), -1)

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

type ApiUpdateOwnInventoryItemRequest struct {
	ctx context.Context
	ApiService *InventoryAPIService
	inventoryItemId string
	updateInventoryItemRequest *UpdateInventoryItemRequest
}

func (r ApiUpdateOwnInventoryItemRequest) UpdateInventoryItemRequest(updateInventoryItemRequest UpdateInventoryItemRequest) ApiUpdateOwnInventoryItemRequest {
	r.updateInventoryItemRequest = &updateInventoryItemRequest
	return r
}

func (r ApiUpdateOwnInventoryItemRequest) Execute() (*InventoryItem, *http.Response, error) {
	return r.ApiService.UpdateOwnInventoryItemExecute(r)
}

/*
UpdateOwnInventoryItem Update Own Inventory Item

Returns the modified InventoryItem object as held by the currently logged in user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param inventoryItemId Must be a valid inventory item ID.
 @return ApiUpdateOwnInventoryItemRequest
*/
func (a *InventoryAPIService) UpdateOwnInventoryItem(ctx context.Context, inventoryItemId string) ApiUpdateOwnInventoryItemRequest {
	return ApiUpdateOwnInventoryItemRequest{
		ApiService: a,
		ctx: ctx,
		inventoryItemId: inventoryItemId,
	}
}

// Execute executes the request
//  @return InventoryItem
func (a *InventoryAPIService) UpdateOwnInventoryItemExecute(r ApiUpdateOwnInventoryItemRequest) (*InventoryItem, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InventoryItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InventoryAPIService.UpdateOwnInventoryItem")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inventory/{inventoryItemId}"
	localVarPath = strings.Replace(localVarPath, "{"+"inventoryItemId"+"}", url.PathEscape(parameterValueToString(r.inventoryItemId, "inventoryItemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.updateInventoryItemRequest
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
