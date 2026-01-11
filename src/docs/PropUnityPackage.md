# PropUnityPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetUrl** | **string** |  | 
**AssetVersion** | **int32** |  | 
**Platform** | **string** | This is normally &#x60;android&#x60;, &#x60;ios&#x60;, &#x60;standalonewindows&#x60;, &#x60;web&#x60;, or the empty value &#x60;&#x60;, but also supposedly can be any random Unity version such as &#x60;2019.2.4-801-Release&#x60; or &#x60;2019.2.2-772-Release&#x60; or even &#x60;unknownplatform&#x60;. | 
**PropSignature** | **string** |  | 
**UnityVersion** | **string** |  | [default to "2022.3.22f1"]
**Variant** | **string** |  | 

## Methods

### NewPropUnityPackage

`func NewPropUnityPackage(assetUrl string, assetVersion int32, platform string, propSignature string, unityVersion string, variant string, ) *PropUnityPackage`

NewPropUnityPackage instantiates a new PropUnityPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropUnityPackageWithDefaults

`func NewPropUnityPackageWithDefaults() *PropUnityPackage`

NewPropUnityPackageWithDefaults instantiates a new PropUnityPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetUrl

`func (o *PropUnityPackage) GetAssetUrl() string`

GetAssetUrl returns the AssetUrl field if non-nil, zero value otherwise.

### GetAssetUrlOk

`func (o *PropUnityPackage) GetAssetUrlOk() (*string, bool)`

GetAssetUrlOk returns a tuple with the AssetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetUrl

`func (o *PropUnityPackage) SetAssetUrl(v string)`

SetAssetUrl sets AssetUrl field to given value.


### GetAssetVersion

`func (o *PropUnityPackage) GetAssetVersion() int32`

GetAssetVersion returns the AssetVersion field if non-nil, zero value otherwise.

### GetAssetVersionOk

`func (o *PropUnityPackage) GetAssetVersionOk() (*int32, bool)`

GetAssetVersionOk returns a tuple with the AssetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetVersion

`func (o *PropUnityPackage) SetAssetVersion(v int32)`

SetAssetVersion sets AssetVersion field to given value.


### GetPlatform

`func (o *PropUnityPackage) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PropUnityPackage) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PropUnityPackage) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetPropSignature

`func (o *PropUnityPackage) GetPropSignature() string`

GetPropSignature returns the PropSignature field if non-nil, zero value otherwise.

### GetPropSignatureOk

`func (o *PropUnityPackage) GetPropSignatureOk() (*string, bool)`

GetPropSignatureOk returns a tuple with the PropSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropSignature

`func (o *PropUnityPackage) SetPropSignature(v string)`

SetPropSignature sets PropSignature field to given value.


### GetUnityVersion

`func (o *PropUnityPackage) GetUnityVersion() string`

GetUnityVersion returns the UnityVersion field if non-nil, zero value otherwise.

### GetUnityVersionOk

`func (o *PropUnityPackage) GetUnityVersionOk() (*string, bool)`

GetUnityVersionOk returns a tuple with the UnityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnityVersion

`func (o *PropUnityPackage) SetUnityVersion(v string)`

SetUnityVersion sets UnityVersion field to given value.


### GetVariant

`func (o *PropUnityPackage) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *PropUnityPackage) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *PropUnityPackage) SetVariant(v string)`

SetVariant sets Variant field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


