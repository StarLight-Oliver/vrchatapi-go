# PurchaseProductListingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListingId** | **string** |  | 
**Quantity** | **int32** |  | [default to 1]
**TotalPrice** | **int32** |  | 

## Methods

### NewPurchaseProductListingRequest

`func NewPurchaseProductListingRequest(listingId string, quantity int32, totalPrice int32, ) *PurchaseProductListingRequest`

NewPurchaseProductListingRequest instantiates a new PurchaseProductListingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseProductListingRequestWithDefaults

`func NewPurchaseProductListingRequestWithDefaults() *PurchaseProductListingRequest`

NewPurchaseProductListingRequestWithDefaults instantiates a new PurchaseProductListingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListingId

`func (o *PurchaseProductListingRequest) GetListingId() string`

GetListingId returns the ListingId field if non-nil, zero value otherwise.

### GetListingIdOk

`func (o *PurchaseProductListingRequest) GetListingIdOk() (*string, bool)`

GetListingIdOk returns a tuple with the ListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingId

`func (o *PurchaseProductListingRequest) SetListingId(v string)`

SetListingId sets ListingId field to given value.


### GetQuantity

`func (o *PurchaseProductListingRequest) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PurchaseProductListingRequest) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PurchaseProductListingRequest) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetTotalPrice

`func (o *PurchaseProductListingRequest) GetTotalPrice() int32`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *PurchaseProductListingRequest) GetTotalPriceOk() (*int32, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *PurchaseProductListingRequest) SetTotalPrice(v int32)`

SetTotalPrice sets TotalPrice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


