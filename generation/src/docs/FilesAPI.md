# \FilesAPI

All URIs are relative to *https://api.vrchat.cloud/api/1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFile**](FilesAPI.md#CreateFile) | **Post** /file | Create File
[**CreateFileVersion**](FilesAPI.md#CreateFileVersion) | **Post** /file/{fileId} | Create File Version
[**DeleteFile**](FilesAPI.md#DeleteFile) | **Delete** /file/{fileId} | Delete File
[**DeleteFileVersion**](FilesAPI.md#DeleteFileVersion) | **Delete** /file/{fileId}/{versionId} | Delete File Version
[**DownloadFileVersion**](FilesAPI.md#DownloadFileVersion) | **Get** /file/{fileId}/{versionId} | Download File Version
[**FinishFileDataUpload**](FilesAPI.md#FinishFileDataUpload) | **Put** /file/{fileId}/{versionId}/{fileType}/finish | Finish FileData Upload
[**GetAdminAssetBundle**](FilesAPI.md#GetAdminAssetBundle) | **Get** /adminassetbundles/{adminAssetBundleId} | Get AdminAssetBundle
[**GetContentAgreementStatus**](FilesAPI.md#GetContentAgreementStatus) | **Get** /agreement | Get Content Agreement Status
[**GetFile**](FilesAPI.md#GetFile) | **Get** /file/{fileId} | Show File
[**GetFileAnalysis**](FilesAPI.md#GetFileAnalysis) | **Get** /analysis/{fileId}/{versionId} | Get File Version Analysis
[**GetFileAnalysisSecurity**](FilesAPI.md#GetFileAnalysisSecurity) | **Get** /analysis/{fileId}/{versionId}/security | Get File Version Analysis Security
[**GetFileAnalysisStandard**](FilesAPI.md#GetFileAnalysisStandard) | **Get** /analysis/{fileId}/{versionId}/standard | Get File Version Analysis Standard
[**GetFileDataUploadStatus**](FilesAPI.md#GetFileDataUploadStatus) | **Get** /file/{fileId}/{versionId}/{fileType}/status | Check FileData Upload Status
[**GetFiles**](FilesAPI.md#GetFiles) | **Get** /files | List Files
[**SetGroupGalleryFileOrder**](FilesAPI.md#SetGroupGalleryFileOrder) | **Put** /files/order | Set Group Gallery File Order
[**StartFileDataUpload**](FilesAPI.md#StartFileDataUpload) | **Put** /file/{fileId}/{versionId}/{fileType}/start | Start FileData Upload
[**SubmitContentAgreement**](FilesAPI.md#SubmitContentAgreement) | **Post** /agreement | Submit Content Agreement
[**UpdateAssetReviewNotes**](FilesAPI.md#UpdateAssetReviewNotes) | **Put** /assetReview/{assetReviewId}/notes | Update Asset Review Notes
[**UploadGalleryImage**](FilesAPI.md#UploadGalleryImage) | **Post** /gallery | Upload gallery image
[**UploadIcon**](FilesAPI.md#UploadIcon) | **Post** /icon | Upload icon
[**UploadImage**](FilesAPI.md#UploadImage) | **Post** /file/image | Upload gallery image, icon, emoji or sticker



## CreateFile

> File CreateFile(ctx).CreateFileRequest(createFileRequest).Execute()

Create File



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
	createFileRequest := *openapiclient.NewCreateFileRequest("Extension_example", openapiclient.MIMEType("application/gzip"), "Name_example") // CreateFileRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.CreateFile(context.Background()).CreateFileRequest(createFileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.CreateFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFile`: File
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.CreateFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFileRequest** | [**CreateFileRequest**](CreateFileRequest.md) |  | 

### Return type

[**File**](File.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFileVersion

> File CreateFileVersion(ctx, fileId).CreateFileVersionRequest(createFileVersionRequest).Execute()

Create File Version



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
	fileId := "file_00000000-0000-0000-0000-000000000000" // string | Must be a valid file ID.
	createFileVersionRequest := *openapiclient.NewCreateFileVersionRequest("SignatureMd5_example", int32(123)) // CreateFileVersionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.CreateFileVersion(context.Background(), fileId).CreateFileVersionRequest(createFileVersionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.CreateFileVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFileVersion`: File
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.CreateFileVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | Must be a valid file ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFileVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createFileVersionRequest** | [**CreateFileVersionRequest**](CreateFileVersionRequest.md) |  | 

### Return type

[**File**](File.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFile

> File DeleteFile(ctx, fileId).Execute()

Delete File



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
	fileId := "file_00000000-0000-0000-0000-000000000000" // string | Must be a valid file ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.DeleteFile(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.DeleteFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFile`: File
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.DeleteFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | Must be a valid file ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**File**](File.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFileVersion

> File DeleteFileVersion(ctx, fileId, versionId).Execute()

Delete File Version



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
	fileId := "file_00000000-0000-0000-0000-000000000000" // string | Must be a valid file ID.
	versionId := int32(1) // int32 | Version ID of the asset.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.DeleteFileVersion(context.Background(), fileId, versionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.DeleteFileVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFileVersion`: File
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.DeleteFileVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | Must be a valid file ID. | 
**versionId** | **int32** | Version ID of the asset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFileVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**File**](File.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadFileVersion

> *os.File DownloadFileVersion(ctx, fileId, versionId).Execute()

Download File Version



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
	fileId := "file_00000000-0000-0000-0000-000000000000" // string | Must be a valid file ID.
	versionId := int32(1) // int32 | Version ID of the asset.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.DownloadFileVersion(context.Background(), fileId, versionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.DownloadFileVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadFileVersion`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.DownloadFileVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | Must be a valid file ID. | 
**versionId** | **int32** | Version ID of the asset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadFileVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FinishFileDataUpload

> File FinishFileDataUpload(ctx, fileId, versionId, fileType).FinishFileDataUploadRequest(finishFileDataUploadRequest).Execute()

Finish FileData Upload



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
	fileId := "file_00000000-0000-0000-0000-000000000000" // string | Must be a valid file ID.
	versionId := int32(1) // int32 | Version ID of the asset.
	fileType := "file" // string | Type of file.
	finishFileDataUploadRequest := *openapiclient.NewFinishFileDataUploadRequest("0", "0") // FinishFileDataUploadRequest | Please see documentation on ETag's: [https://teppen.io/2018/06/23/aws_s3_etags/](https://teppen.io/2018/06/23/aws_s3_etags/)  ETag's should NOT be present when uploading a `signature`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.FinishFileDataUpload(context.Background(), fileId, versionId, fileType).FinishFileDataUploadRequest(finishFileDataUploadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.FinishFileDataUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FinishFileDataUpload`: File
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.FinishFileDataUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | Must be a valid file ID. | 
**versionId** | **int32** | Version ID of the asset. | 
**fileType** | **string** | Type of file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFinishFileDataUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **finishFileDataUploadRequest** | [**FinishFileDataUploadRequest**](FinishFileDataUploadRequest.md) | Please see documentation on ETag&#39;s: [https://teppen.io/2018/06/23/aws_s3_etags/](https://teppen.io/2018/06/23/aws_s3_etags/)  ETag&#39;s should NOT be present when uploading a &#x60;signature&#x60;. | 

### Return type

[**File**](File.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminAssetBundle

> AdminAssetBundle GetAdminAssetBundle(ctx, adminAssetBundleId).Execute()

Get AdminAssetBundle



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
	adminAssetBundleId := "aab_e159e72c-ce54-4fbe-8c37-96af02f6d18d" // string | Must be a valid admin asset bundle ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.GetAdminAssetBundle(context.Background(), adminAssetBundleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetAdminAssetBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdminAssetBundle`: AdminAssetBundle
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetAdminAssetBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adminAssetBundleId** | **string** | Must be a valid admin asset bundle ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminAssetBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdminAssetBundle**](AdminAssetBundle.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContentAgreementStatus

> AgreementStatus GetContentAgreementStatus(ctx).AgreementCode(agreementCode).ContentId(contentId).Version(version).Execute()

Get Content Agreement Status



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
	agreementCode := openapiclient.AgreementCode("content.copyright.owned") // AgreementCode | The type of agreement (currently content.copyright.owned) (default to "content.copyright.owned")
	contentId := "avtr_c38a1615-5bf5-42b4-84eb-a8b6c37cbd11" // string | The id of the content being uploaded, such as a WorldID, AvatarID, or PropID
	version := int32(1) // int32 | The version of the agreement (currently 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.GetContentAgreementStatus(context.Background()).AgreementCode(agreementCode).ContentId(contentId).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetContentAgreementStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContentAgreementStatus`: AgreementStatus
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetContentAgreementStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContentAgreementStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agreementCode** | [**AgreementCode**](AgreementCode.md) | The type of agreement (currently content.copyright.owned) | [default to &quot;content.copyright.owned&quot;]
 **contentId** | **string** | The id of the content being uploaded, such as a WorldID, AvatarID, or PropID | 
 **version** | **int32** | The version of the agreement (currently 1) | 

### Return type

[**AgreementStatus**](AgreementStatus.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFile

> File GetFile(ctx, fileId).Execute()

Show File



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
	fileId := "file_00000000-0000-0000-0000-000000000000" // string | Must be a valid file ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.GetFile(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFile`: File
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | Must be a valid file ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**File**](File.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileAnalysis

> FileAnalysis GetFileAnalysis(ctx, fileId, versionId).Execute()

Get File Version Analysis



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
	fileId := "file_00000000-0000-0000-0000-000000000000" // string | Must be a valid file ID.
	versionId := int32(1) // int32 | Version ID of the asset.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.GetFileAnalysis(context.Background(), fileId, versionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFileAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileAnalysis`: FileAnalysis
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFileAnalysis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | Must be a valid file ID. | 
**versionId** | **int32** | Version ID of the asset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FileAnalysis**](FileAnalysis.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileAnalysisSecurity

> FileAnalysis GetFileAnalysisSecurity(ctx, fileId, versionId).Execute()

Get File Version Analysis Security



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
	fileId := "file_00000000-0000-0000-0000-000000000000" // string | Must be a valid file ID.
	versionId := int32(1) // int32 | Version ID of the asset.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.GetFileAnalysisSecurity(context.Background(), fileId, versionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFileAnalysisSecurity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileAnalysisSecurity`: FileAnalysis
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFileAnalysisSecurity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | Must be a valid file ID. | 
**versionId** | **int32** | Version ID of the asset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileAnalysisSecurityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FileAnalysis**](FileAnalysis.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileAnalysisStandard

> FileAnalysis GetFileAnalysisStandard(ctx, fileId, versionId).Execute()

Get File Version Analysis Standard



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
	fileId := "file_00000000-0000-0000-0000-000000000000" // string | Must be a valid file ID.
	versionId := int32(1) // int32 | Version ID of the asset.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.GetFileAnalysisStandard(context.Background(), fileId, versionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFileAnalysisStandard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileAnalysisStandard`: FileAnalysis
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFileAnalysisStandard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | Must be a valid file ID. | 
**versionId** | **int32** | Version ID of the asset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileAnalysisStandardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FileAnalysis**](FileAnalysis.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileDataUploadStatus

> FileVersionUploadStatus GetFileDataUploadStatus(ctx, fileId, versionId, fileType).Execute()

Check FileData Upload Status



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
	fileId := "file_00000000-0000-0000-0000-000000000000" // string | Must be a valid file ID.
	versionId := int32(1) // int32 | Version ID of the asset.
	fileType := "file" // string | Type of file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.GetFileDataUploadStatus(context.Background(), fileId, versionId, fileType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFileDataUploadStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileDataUploadStatus`: FileVersionUploadStatus
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFileDataUploadStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | Must be a valid file ID. | 
**versionId** | **int32** | Version ID of the asset. | 
**fileType** | **string** | Type of file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileDataUploadStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FileVersionUploadStatus**](FileVersionUploadStatus.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFiles

> []File GetFiles(ctx).Tag(tag).UserId(userId).N(n).Offset(offset).Execute()

List Files



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
	tag := "tag_example" // string | Tag, for example \"icon\" or \"gallery\", not included by default. (optional)
	userId := "userId_example" // string | UserID, will always generate a 500 permission error. (optional)
	n := int32(56) // int32 | The number of objects to return. (optional) (default to 60)
	offset := int32(56) // int32 | A zero-based offset from the default object sorting from where search results start. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.GetFiles(context.Background()).Tag(tag).UserId(userId).N(n).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFiles`: []File
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | **string** | Tag, for example \&quot;icon\&quot; or \&quot;gallery\&quot;, not included by default. | 
 **userId** | **string** | UserID, will always generate a 500 permission error. | 
 **n** | **int32** | The number of objects to return. | [default to 60]
 **offset** | **int32** | A zero-based offset from the default object sorting from where search results start. | 

### Return type

[**[]File**](File.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetGroupGalleryFileOrder

> GroupGalleryFileOrder SetGroupGalleryFileOrder(ctx).GroupGalleryFileOrderRequest(groupGalleryFileOrderRequest).Execute()

Set Group Gallery File Order



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
	groupGalleryFileOrderRequest := *openapiclient.NewGroupGalleryFileOrderRequest("ggal_a03a4b55-4ca6-4490-9519-40ba6351a233", []string{"file_ce35d830-e20a-4df0-a6d4-5aaef4508044"}) // GroupGalleryFileOrderRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.SetGroupGalleryFileOrder(context.Background()).GroupGalleryFileOrderRequest(groupGalleryFileOrderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.SetGroupGalleryFileOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetGroupGalleryFileOrder`: GroupGalleryFileOrder
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.SetGroupGalleryFileOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetGroupGalleryFileOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupGalleryFileOrderRequest** | [**GroupGalleryFileOrderRequest**](GroupGalleryFileOrderRequest.md) |  | 

### Return type

[**GroupGalleryFileOrder**](GroupGalleryFileOrder.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartFileDataUpload

> FileUploadURL StartFileDataUpload(ctx, fileId, versionId, fileType).PartNumber(partNumber).Execute()

Start FileData Upload



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
	fileId := "file_00000000-0000-0000-0000-000000000000" // string | Must be a valid file ID.
	versionId := int32(1) // int32 | Version ID of the asset.
	fileType := "file" // string | Type of file.
	partNumber := int32(1) // int32 | The part number to start uploading. If not provided, the first part will be started. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.StartFileDataUpload(context.Background(), fileId, versionId, fileType).PartNumber(partNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.StartFileDataUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartFileDataUpload`: FileUploadURL
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.StartFileDataUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | Must be a valid file ID. | 
**versionId** | **int32** | Version ID of the asset. | 
**fileType** | **string** | Type of file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartFileDataUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **partNumber** | **int32** | The part number to start uploading. If not provided, the first part will be started. | 

### Return type

[**FileUploadURL**](FileUploadURL.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitContentAgreement

> Agreement SubmitContentAgreement(ctx).AgreementRequest(agreementRequest).Execute()

Submit Content Agreement



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
	agreementRequest := *openapiclient.NewAgreementRequest(openapiclient.AgreementCode("content.copyright.owned"), "By clicking OK, I certify that I have the necessary rights to upload this content and that it will not infringe on any third-party legal or intellectual property rights.", "avtr_c38a1615-5bf5-42b4-84eb-a8b6c37cbd11", int32(1)) // AgreementRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.SubmitContentAgreement(context.Background()).AgreementRequest(agreementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.SubmitContentAgreement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitContentAgreement`: Agreement
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.SubmitContentAgreement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitContentAgreementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agreementRequest** | [**AgreementRequest**](AgreementRequest.md) |  | 

### Return type

[**Agreement**](Agreement.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssetReviewNotes

> UpdateAssetReviewNotes(ctx, assetReviewId).UpdateAssetReviewNotesRequest(updateAssetReviewNotesRequest).Execute()

Update Asset Review Notes



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
	assetReviewId := "assetReviewId_example" // string | Must be an valid asset review ID.
	updateAssetReviewNotesRequest := *openapiclient.NewUpdateAssetReviewNotesRequest("ReviewNotes_example") // UpdateAssetReviewNotesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilesAPI.UpdateAssetReviewNotes(context.Background(), assetReviewId).UpdateAssetReviewNotesRequest(updateAssetReviewNotesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.UpdateAssetReviewNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetReviewId** | **string** | Must be an valid asset review ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAssetReviewNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAssetReviewNotesRequest** | [**UpdateAssetReviewNotesRequest**](UpdateAssetReviewNotesRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadGalleryImage

> File UploadGalleryImage(ctx).File(file).Execute()

Upload gallery image



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
	file := os.NewFile(1234, "some_file") // *os.File | The binary blob of the png file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.UploadGalleryImage(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.UploadGalleryImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadGalleryImage`: File
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.UploadGalleryImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadGalleryImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The binary blob of the png file. | 

### Return type

[**File**](File.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadIcon

> File UploadIcon(ctx).File(file).Execute()

Upload icon



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
	file := os.NewFile(1234, "some_file") // *os.File | The binary blob of the png file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.UploadIcon(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.UploadIcon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadIcon`: File
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.UploadIcon`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadIconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The binary blob of the png file. | 

### Return type

[**File**](File.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadImage

> File UploadImage(ctx).File(file).Tag(tag).AnimationStyle(animationStyle).Frames(frames).FramesOverTime(framesOverTime).LoopStyle(loopStyle).MaskTag(maskTag).Execute()

Upload gallery image, icon, emoji or sticker



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
	file := os.NewFile(1234, "some_file") // *os.File | The binary blob of the png file.
	tag := openapiclient.ImagePurpose("admin") // ImagePurpose |  (default to "gallery")
	animationStyle := openapiclient.ImageAnimationStyle("aura") // ImageAnimationStyle |  (optional)
	frames := int32(56) // int32 | Required for animated images. Total number of frames of the spritesheet to be animated. (optional)
	framesOverTime := int32(56) // int32 | Required for animated images. Animation frames per second. (optional)
	loopStyle := openapiclient.ImageLoopStyle("linear") // ImageLoopStyle |  (optional) (default to "linear")
	maskTag := openapiclient.ImageMask("circle") // ImageMask |  (optional) (default to "square")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.UploadImage(context.Background()).File(file).Tag(tag).AnimationStyle(animationStyle).Frames(frames).FramesOverTime(framesOverTime).LoopStyle(loopStyle).MaskTag(maskTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.UploadImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadImage`: File
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.UploadImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The binary blob of the png file. | 
 **tag** | [**ImagePurpose**](ImagePurpose.md) |  | [default to &quot;gallery&quot;]
 **animationStyle** | [**ImageAnimationStyle**](ImageAnimationStyle.md) |  | 
 **frames** | **int32** | Required for animated images. Total number of frames of the spritesheet to be animated. | 
 **framesOverTime** | **int32** | Required for animated images. Animation frames per second. | 
 **loopStyle** | [**ImageLoopStyle**](ImageLoopStyle.md) |  | [default to &quot;linear&quot;]
 **maskTag** | [**ImageMask**](ImageMask.md) |  | [default to &quot;square&quot;]

### Return type

[**File**](File.md)

### Authorization

[authCookie](../README.md#authCookie)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

