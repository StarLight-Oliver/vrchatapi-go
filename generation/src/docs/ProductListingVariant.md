# ProductListingVariant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EffectiveFrom** | Pointer to **time.Time** |  | [optional] 
**ListingVariantId** | **string** |  | 
**NonRefundable** | **bool** |  | 
**Quantity** | **int32** |  | 
**SellerVariant** | **bool** |  | 
**UnitPriceTokens** | **int32** |  | 

## Methods

### NewProductListingVariant

`func NewProductListingVariant(listingVariantId string, nonRefundable bool, quantity int32, sellerVariant bool, unitPriceTokens int32, ) *ProductListingVariant`

NewProductListingVariant instantiates a new ProductListingVariant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductListingVariantWithDefaults

`func NewProductListingVariantWithDefaults() *ProductListingVariant`

NewProductListingVariantWithDefaults instantiates a new ProductListingVariant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffectiveFrom

`func (o *ProductListingVariant) GetEffectiveFrom() time.Time`

GetEffectiveFrom returns the EffectiveFrom field if non-nil, zero value otherwise.

### GetEffectiveFromOk

`func (o *ProductListingVariant) GetEffectiveFromOk() (*time.Time, bool)`

GetEffectiveFromOk returns a tuple with the EffectiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveFrom

`func (o *ProductListingVariant) SetEffectiveFrom(v time.Time)`

SetEffectiveFrom sets EffectiveFrom field to given value.

### HasEffectiveFrom

`func (o *ProductListingVariant) HasEffectiveFrom() bool`

HasEffectiveFrom returns a boolean if a field has been set.

### GetListingVariantId

`func (o *ProductListingVariant) GetListingVariantId() string`

GetListingVariantId returns the ListingVariantId field if non-nil, zero value otherwise.

### GetListingVariantIdOk

`func (o *ProductListingVariant) GetListingVariantIdOk() (*string, bool)`

GetListingVariantIdOk returns a tuple with the ListingVariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingVariantId

`func (o *ProductListingVariant) SetListingVariantId(v string)`

SetListingVariantId sets ListingVariantId field to given value.


### GetNonRefundable

`func (o *ProductListingVariant) GetNonRefundable() bool`

GetNonRefundable returns the NonRefundable field if non-nil, zero value otherwise.

### GetNonRefundableOk

`func (o *ProductListingVariant) GetNonRefundableOk() (*bool, bool)`

GetNonRefundableOk returns a tuple with the NonRefundable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonRefundable

`func (o *ProductListingVariant) SetNonRefundable(v bool)`

SetNonRefundable sets NonRefundable field to given value.


### GetQuantity

`func (o *ProductListingVariant) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ProductListingVariant) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ProductListingVariant) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetSellerVariant

`func (o *ProductListingVariant) GetSellerVariant() bool`

GetSellerVariant returns the SellerVariant field if non-nil, zero value otherwise.

### GetSellerVariantOk

`func (o *ProductListingVariant) GetSellerVariantOk() (*bool, bool)`

GetSellerVariantOk returns a tuple with the SellerVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerVariant

`func (o *ProductListingVariant) SetSellerVariant(v bool)`

SetSellerVariant sets SellerVariant field to given value.


### GetUnitPriceTokens

`func (o *ProductListingVariant) GetUnitPriceTokens() int32`

GetUnitPriceTokens returns the UnitPriceTokens field if non-nil, zero value otherwise.

### GetUnitPriceTokensOk

`func (o *ProductListingVariant) GetUnitPriceTokensOk() (*int32, bool)`

GetUnitPriceTokensOk returns a tuple with the UnitPriceTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPriceTokens

`func (o *ProductListingVariant) SetUnitPriceTokens(v int32)`

SetUnitPriceTokens sets UnitPriceTokens field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


