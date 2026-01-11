# ProductListing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** |  | 
**BuyerRefundable** | **bool** |  | 
**Description** | **string** |  | 
**DisplayName** | **string** |  | 
**Duration** | Pointer to **NullableInt32** |  | [optional] 
**DurationType** | Pointer to **NullableString** |  | [optional] 
**GroupIcon** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**GroupName** | Pointer to **NullableString** |  | [optional] 
**HasAvatar** | **bool** |  | 
**HasUdon** | **bool** |  | 
**HydratedProducts** | Pointer to [**[]Product**](Product.md) |  | [optional] 
**Id** | **string** |  | 
**ImageId** | Pointer to **string** |  | [optional] 
**ImageUrl** | Pointer to **NullableString** |  | [optional] 
**ListingType** | [**ProductListingType**](ProductListingType.md) |  | [default to SUBSCRIPTION]
**ListingVariants** | Pointer to [**[]ProductListingVariant**](ProductListingVariant.md) |  | [optional] 
**Permanent** | Pointer to **bool** |  | [optional] 
**PriceTokens** | **int32** |  | 
**ProductIds** | **[]string** |  | 
**ProductType** | [**ProductType**](ProductType.md) |  | [default to UDON]
**Products** | **[]map[string]interface{}** |  | 
**Quantifiable** | Pointer to **bool** |  | [optional] 
**Recurrable** | **bool** |  | 
**Refundable** | **bool** |  | 
**SellerDisplayName** | **string** |  | 
**SellerId** | **string** |  | 
**SoldByVrc** | Pointer to **bool** |  | [optional] 
**Stackable** | **bool** |  | 
**StoreIds** | **[]string** |  | 
**Subtitle** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**VrcPlusDiscountPrice** | Pointer to **int32** |  | [optional] 
**WhenToExpire** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewProductListing

`func NewProductListing(active bool, buyerRefundable bool, description string, displayName string, hasAvatar bool, hasUdon bool, id string, listingType ProductListingType, priceTokens int32, productIds []string, productType ProductType, products []map[string]interface{}, recurrable bool, refundable bool, sellerDisplayName string, sellerId string, stackable bool, storeIds []string, ) *ProductListing`

NewProductListing instantiates a new ProductListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductListingWithDefaults

`func NewProductListingWithDefaults() *ProductListing`

NewProductListingWithDefaults instantiates a new ProductListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ProductListing) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ProductListing) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ProductListing) SetActive(v bool)`

SetActive sets Active field to given value.


### GetBuyerRefundable

`func (o *ProductListing) GetBuyerRefundable() bool`

GetBuyerRefundable returns the BuyerRefundable field if non-nil, zero value otherwise.

### GetBuyerRefundableOk

`func (o *ProductListing) GetBuyerRefundableOk() (*bool, bool)`

GetBuyerRefundableOk returns a tuple with the BuyerRefundable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerRefundable

`func (o *ProductListing) SetBuyerRefundable(v bool)`

SetBuyerRefundable sets BuyerRefundable field to given value.


### GetDescription

`func (o *ProductListing) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductListing) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductListing) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDisplayName

`func (o *ProductListing) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ProductListing) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ProductListing) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDuration

`func (o *ProductListing) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ProductListing) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ProductListing) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ProductListing) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *ProductListing) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *ProductListing) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetDurationType

`func (o *ProductListing) GetDurationType() string`

GetDurationType returns the DurationType field if non-nil, zero value otherwise.

### GetDurationTypeOk

`func (o *ProductListing) GetDurationTypeOk() (*string, bool)`

GetDurationTypeOk returns a tuple with the DurationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationType

`func (o *ProductListing) SetDurationType(v string)`

SetDurationType sets DurationType field to given value.

### HasDurationType

`func (o *ProductListing) HasDurationType() bool`

HasDurationType returns a boolean if a field has been set.

### SetDurationTypeNil

`func (o *ProductListing) SetDurationTypeNil(b bool)`

 SetDurationTypeNil sets the value for DurationType to be an explicit nil

### UnsetDurationType
`func (o *ProductListing) UnsetDurationType()`

UnsetDurationType ensures that no value is present for DurationType, not even an explicit nil
### GetGroupIcon

`func (o *ProductListing) GetGroupIcon() string`

GetGroupIcon returns the GroupIcon field if non-nil, zero value otherwise.

### GetGroupIconOk

`func (o *ProductListing) GetGroupIconOk() (*string, bool)`

GetGroupIconOk returns a tuple with the GroupIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIcon

`func (o *ProductListing) SetGroupIcon(v string)`

SetGroupIcon sets GroupIcon field to given value.

### HasGroupIcon

`func (o *ProductListing) HasGroupIcon() bool`

HasGroupIcon returns a boolean if a field has been set.

### GetGroupId

`func (o *ProductListing) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ProductListing) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ProductListing) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ProductListing) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupName

`func (o *ProductListing) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ProductListing) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ProductListing) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *ProductListing) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *ProductListing) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *ProductListing) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetHasAvatar

`func (o *ProductListing) GetHasAvatar() bool`

GetHasAvatar returns the HasAvatar field if non-nil, zero value otherwise.

### GetHasAvatarOk

`func (o *ProductListing) GetHasAvatarOk() (*bool, bool)`

GetHasAvatarOk returns a tuple with the HasAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAvatar

`func (o *ProductListing) SetHasAvatar(v bool)`

SetHasAvatar sets HasAvatar field to given value.


### GetHasUdon

`func (o *ProductListing) GetHasUdon() bool`

GetHasUdon returns the HasUdon field if non-nil, zero value otherwise.

### GetHasUdonOk

`func (o *ProductListing) GetHasUdonOk() (*bool, bool)`

GetHasUdonOk returns a tuple with the HasUdon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUdon

`func (o *ProductListing) SetHasUdon(v bool)`

SetHasUdon sets HasUdon field to given value.


### GetHydratedProducts

`func (o *ProductListing) GetHydratedProducts() []Product`

GetHydratedProducts returns the HydratedProducts field if non-nil, zero value otherwise.

### GetHydratedProductsOk

`func (o *ProductListing) GetHydratedProductsOk() (*[]Product, bool)`

GetHydratedProductsOk returns a tuple with the HydratedProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHydratedProducts

`func (o *ProductListing) SetHydratedProducts(v []Product)`

SetHydratedProducts sets HydratedProducts field to given value.

### HasHydratedProducts

`func (o *ProductListing) HasHydratedProducts() bool`

HasHydratedProducts returns a boolean if a field has been set.

### GetId

`func (o *ProductListing) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductListing) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductListing) SetId(v string)`

SetId sets Id field to given value.


### GetImageId

`func (o *ProductListing) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *ProductListing) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *ProductListing) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *ProductListing) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetImageUrl

`func (o *ProductListing) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *ProductListing) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *ProductListing) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *ProductListing) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *ProductListing) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *ProductListing) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetListingType

`func (o *ProductListing) GetListingType() ProductListingType`

GetListingType returns the ListingType field if non-nil, zero value otherwise.

### GetListingTypeOk

`func (o *ProductListing) GetListingTypeOk() (*ProductListingType, bool)`

GetListingTypeOk returns a tuple with the ListingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingType

`func (o *ProductListing) SetListingType(v ProductListingType)`

SetListingType sets ListingType field to given value.


### GetListingVariants

`func (o *ProductListing) GetListingVariants() []ProductListingVariant`

GetListingVariants returns the ListingVariants field if non-nil, zero value otherwise.

### GetListingVariantsOk

`func (o *ProductListing) GetListingVariantsOk() (*[]ProductListingVariant, bool)`

GetListingVariantsOk returns a tuple with the ListingVariants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingVariants

`func (o *ProductListing) SetListingVariants(v []ProductListingVariant)`

SetListingVariants sets ListingVariants field to given value.

### HasListingVariants

`func (o *ProductListing) HasListingVariants() bool`

HasListingVariants returns a boolean if a field has been set.

### GetPermanent

`func (o *ProductListing) GetPermanent() bool`

GetPermanent returns the Permanent field if non-nil, zero value otherwise.

### GetPermanentOk

`func (o *ProductListing) GetPermanentOk() (*bool, bool)`

GetPermanentOk returns a tuple with the Permanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanent

`func (o *ProductListing) SetPermanent(v bool)`

SetPermanent sets Permanent field to given value.

### HasPermanent

`func (o *ProductListing) HasPermanent() bool`

HasPermanent returns a boolean if a field has been set.

### GetPriceTokens

`func (o *ProductListing) GetPriceTokens() int32`

GetPriceTokens returns the PriceTokens field if non-nil, zero value otherwise.

### GetPriceTokensOk

`func (o *ProductListing) GetPriceTokensOk() (*int32, bool)`

GetPriceTokensOk returns a tuple with the PriceTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTokens

`func (o *ProductListing) SetPriceTokens(v int32)`

SetPriceTokens sets PriceTokens field to given value.


### GetProductIds

`func (o *ProductListing) GetProductIds() []string`

GetProductIds returns the ProductIds field if non-nil, zero value otherwise.

### GetProductIdsOk

`func (o *ProductListing) GetProductIdsOk() (*[]string, bool)`

GetProductIdsOk returns a tuple with the ProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIds

`func (o *ProductListing) SetProductIds(v []string)`

SetProductIds sets ProductIds field to given value.


### GetProductType

`func (o *ProductListing) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *ProductListing) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *ProductListing) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.


### GetProducts

`func (o *ProductListing) GetProducts() []map[string]interface{}`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *ProductListing) GetProductsOk() (*[]map[string]interface{}, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *ProductListing) SetProducts(v []map[string]interface{})`

SetProducts sets Products field to given value.


### GetQuantifiable

`func (o *ProductListing) GetQuantifiable() bool`

GetQuantifiable returns the Quantifiable field if non-nil, zero value otherwise.

### GetQuantifiableOk

`func (o *ProductListing) GetQuantifiableOk() (*bool, bool)`

GetQuantifiableOk returns a tuple with the Quantifiable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantifiable

`func (o *ProductListing) SetQuantifiable(v bool)`

SetQuantifiable sets Quantifiable field to given value.

### HasQuantifiable

`func (o *ProductListing) HasQuantifiable() bool`

HasQuantifiable returns a boolean if a field has been set.

### GetRecurrable

`func (o *ProductListing) GetRecurrable() bool`

GetRecurrable returns the Recurrable field if non-nil, zero value otherwise.

### GetRecurrableOk

`func (o *ProductListing) GetRecurrableOk() (*bool, bool)`

GetRecurrableOk returns a tuple with the Recurrable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrable

`func (o *ProductListing) SetRecurrable(v bool)`

SetRecurrable sets Recurrable field to given value.


### GetRefundable

`func (o *ProductListing) GetRefundable() bool`

GetRefundable returns the Refundable field if non-nil, zero value otherwise.

### GetRefundableOk

`func (o *ProductListing) GetRefundableOk() (*bool, bool)`

GetRefundableOk returns a tuple with the Refundable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundable

`func (o *ProductListing) SetRefundable(v bool)`

SetRefundable sets Refundable field to given value.


### GetSellerDisplayName

`func (o *ProductListing) GetSellerDisplayName() string`

GetSellerDisplayName returns the SellerDisplayName field if non-nil, zero value otherwise.

### GetSellerDisplayNameOk

`func (o *ProductListing) GetSellerDisplayNameOk() (*string, bool)`

GetSellerDisplayNameOk returns a tuple with the SellerDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerDisplayName

`func (o *ProductListing) SetSellerDisplayName(v string)`

SetSellerDisplayName sets SellerDisplayName field to given value.


### GetSellerId

`func (o *ProductListing) GetSellerId() string`

GetSellerId returns the SellerId field if non-nil, zero value otherwise.

### GetSellerIdOk

`func (o *ProductListing) GetSellerIdOk() (*string, bool)`

GetSellerIdOk returns a tuple with the SellerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerId

`func (o *ProductListing) SetSellerId(v string)`

SetSellerId sets SellerId field to given value.


### GetSoldByVrc

`func (o *ProductListing) GetSoldByVrc() bool`

GetSoldByVrc returns the SoldByVrc field if non-nil, zero value otherwise.

### GetSoldByVrcOk

`func (o *ProductListing) GetSoldByVrcOk() (*bool, bool)`

GetSoldByVrcOk returns a tuple with the SoldByVrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoldByVrc

`func (o *ProductListing) SetSoldByVrc(v bool)`

SetSoldByVrc sets SoldByVrc field to given value.

### HasSoldByVrc

`func (o *ProductListing) HasSoldByVrc() bool`

HasSoldByVrc returns a boolean if a field has been set.

### GetStackable

`func (o *ProductListing) GetStackable() bool`

GetStackable returns the Stackable field if non-nil, zero value otherwise.

### GetStackableOk

`func (o *ProductListing) GetStackableOk() (*bool, bool)`

GetStackableOk returns a tuple with the Stackable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackable

`func (o *ProductListing) SetStackable(v bool)`

SetStackable sets Stackable field to given value.


### GetStoreIds

`func (o *ProductListing) GetStoreIds() []string`

GetStoreIds returns the StoreIds field if non-nil, zero value otherwise.

### GetStoreIdsOk

`func (o *ProductListing) GetStoreIdsOk() (*[]string, bool)`

GetStoreIdsOk returns a tuple with the StoreIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreIds

`func (o *ProductListing) SetStoreIds(v []string)`

SetStoreIds sets StoreIds field to given value.


### GetSubtitle

`func (o *ProductListing) GetSubtitle() string`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *ProductListing) GetSubtitleOk() (*string, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *ProductListing) SetSubtitle(v string)`

SetSubtitle sets Subtitle field to given value.

### HasSubtitle

`func (o *ProductListing) HasSubtitle() bool`

HasSubtitle returns a boolean if a field has been set.

### GetTags

`func (o *ProductListing) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProductListing) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProductListing) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProductListing) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVrcPlusDiscountPrice

`func (o *ProductListing) GetVrcPlusDiscountPrice() int32`

GetVrcPlusDiscountPrice returns the VrcPlusDiscountPrice field if non-nil, zero value otherwise.

### GetVrcPlusDiscountPriceOk

`func (o *ProductListing) GetVrcPlusDiscountPriceOk() (*int32, bool)`

GetVrcPlusDiscountPriceOk returns a tuple with the VrcPlusDiscountPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrcPlusDiscountPrice

`func (o *ProductListing) SetVrcPlusDiscountPrice(v int32)`

SetVrcPlusDiscountPrice sets VrcPlusDiscountPrice field to given value.

### HasVrcPlusDiscountPrice

`func (o *ProductListing) HasVrcPlusDiscountPrice() bool`

HasVrcPlusDiscountPrice returns a boolean if a field has been set.

### GetWhenToExpire

`func (o *ProductListing) GetWhenToExpire() time.Time`

GetWhenToExpire returns the WhenToExpire field if non-nil, zero value otherwise.

### GetWhenToExpireOk

`func (o *ProductListing) GetWhenToExpireOk() (*time.Time, bool)`

GetWhenToExpireOk returns a tuple with the WhenToExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhenToExpire

`func (o *ProductListing) SetWhenToExpire(v time.Time)`

SetWhenToExpire sets WhenToExpire field to given value.

### HasWhenToExpire

`func (o *ProductListing) HasWhenToExpire() bool`

HasWhenToExpire returns a boolean if a field has been set.

### SetWhenToExpireNil

`func (o *ProductListing) SetWhenToExpireNil(b bool)`

 SetWhenToExpireNil sets the value for WhenToExpire to be an explicit nil

### UnsetWhenToExpire
`func (o *ProductListing) UnsetWhenToExpire()`

UnsetWhenToExpire ensures that no value is present for WhenToExpire, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


