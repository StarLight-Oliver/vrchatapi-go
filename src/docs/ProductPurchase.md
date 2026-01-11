# ProductPurchase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyerDisplayName** | **string** |  | 
**BuyerId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**FirstParty** | Pointer to **bool** |  | [optional] 
**IsBuyer** | **bool** |  | 
**IsGift** | **bool** |  | 
**IsReceiver** | **bool** |  | 
**IsSeller** | **bool** |  | 
**ListingCurrentlyAvailable** | **bool** |  | 
**ListingDisplayName** | **string** |  | 
**ListingId** | **string** |  | 
**ListingImageId** | **string** |  | 
**ListingSubtitle** | **string** |  | 
**ListingType** | [**ProductListingType**](ProductListingType.md) |  | [default to SUBSCRIPTION]
**Products** | **[]map[string]interface{}** |  | 
**PurchaseActive** | **bool** |  | 
**PurchaseContext** | [**ProductPurchasePurchaseContext**](ProductPurchasePurchaseContext.md) |  | 
**PurchaseCurrentStatus** | **string** |  | 
**PurchaseDate** | **time.Time** |  | 
**PurchaseDuration** | Pointer to **int32** |  | [optional] 
**PurchaseDurationType** | Pointer to **string** |  | [optional] 
**PurchaseEndDate** | **time.Time** |  | 
**PurchaseId** | **string** |  | 
**PurchaseLatest** | **bool** |  | 
**PurchasePrice** | **int32** |  | 
**PurchaseQuantity** | **int32** |  | 
**PurchaseStartDate** | **time.Time** |  | 
**PurchaseToken** | **map[string]interface{}** |  | 
**PurchaseType** | **string** |  | 
**PurchaseUnitPrice** | **int32** |  | 
**ReceiverDisplayName** | **string** |  | 
**ReceiverId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**Recurrable** | **bool** |  | 
**Refundable** | **bool** |  | 
**SellerDisplayName** | **string** |  | 
**SellerId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**Stackable** | **bool** |  | 
**WillRecur** | **bool** |  | 

## Methods

### NewProductPurchase

`func NewProductPurchase(buyerDisplayName string, buyerId string, isBuyer bool, isGift bool, isReceiver bool, isSeller bool, listingCurrentlyAvailable bool, listingDisplayName string, listingId string, listingImageId string, listingSubtitle string, listingType ProductListingType, products []map[string]interface{}, purchaseActive bool, purchaseContext ProductPurchasePurchaseContext, purchaseCurrentStatus string, purchaseDate time.Time, purchaseEndDate time.Time, purchaseId string, purchaseLatest bool, purchasePrice int32, purchaseQuantity int32, purchaseStartDate time.Time, purchaseToken map[string]interface{}, purchaseType string, purchaseUnitPrice int32, receiverDisplayName string, receiverId string, recurrable bool, refundable bool, sellerDisplayName string, sellerId string, stackable bool, willRecur bool, ) *ProductPurchase`

NewProductPurchase instantiates a new ProductPurchase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductPurchaseWithDefaults

`func NewProductPurchaseWithDefaults() *ProductPurchase`

NewProductPurchaseWithDefaults instantiates a new ProductPurchase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyerDisplayName

`func (o *ProductPurchase) GetBuyerDisplayName() string`

GetBuyerDisplayName returns the BuyerDisplayName field if non-nil, zero value otherwise.

### GetBuyerDisplayNameOk

`func (o *ProductPurchase) GetBuyerDisplayNameOk() (*string, bool)`

GetBuyerDisplayNameOk returns a tuple with the BuyerDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerDisplayName

`func (o *ProductPurchase) SetBuyerDisplayName(v string)`

SetBuyerDisplayName sets BuyerDisplayName field to given value.


### GetBuyerId

`func (o *ProductPurchase) GetBuyerId() string`

GetBuyerId returns the BuyerId field if non-nil, zero value otherwise.

### GetBuyerIdOk

`func (o *ProductPurchase) GetBuyerIdOk() (*string, bool)`

GetBuyerIdOk returns a tuple with the BuyerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerId

`func (o *ProductPurchase) SetBuyerId(v string)`

SetBuyerId sets BuyerId field to given value.


### GetFirstParty

`func (o *ProductPurchase) GetFirstParty() bool`

GetFirstParty returns the FirstParty field if non-nil, zero value otherwise.

### GetFirstPartyOk

`func (o *ProductPurchase) GetFirstPartyOk() (*bool, bool)`

GetFirstPartyOk returns a tuple with the FirstParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstParty

`func (o *ProductPurchase) SetFirstParty(v bool)`

SetFirstParty sets FirstParty field to given value.

### HasFirstParty

`func (o *ProductPurchase) HasFirstParty() bool`

HasFirstParty returns a boolean if a field has been set.

### GetIsBuyer

`func (o *ProductPurchase) GetIsBuyer() bool`

GetIsBuyer returns the IsBuyer field if non-nil, zero value otherwise.

### GetIsBuyerOk

`func (o *ProductPurchase) GetIsBuyerOk() (*bool, bool)`

GetIsBuyerOk returns a tuple with the IsBuyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuyer

`func (o *ProductPurchase) SetIsBuyer(v bool)`

SetIsBuyer sets IsBuyer field to given value.


### GetIsGift

`func (o *ProductPurchase) GetIsGift() bool`

GetIsGift returns the IsGift field if non-nil, zero value otherwise.

### GetIsGiftOk

`func (o *ProductPurchase) GetIsGiftOk() (*bool, bool)`

GetIsGiftOk returns a tuple with the IsGift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGift

`func (o *ProductPurchase) SetIsGift(v bool)`

SetIsGift sets IsGift field to given value.


### GetIsReceiver

`func (o *ProductPurchase) GetIsReceiver() bool`

GetIsReceiver returns the IsReceiver field if non-nil, zero value otherwise.

### GetIsReceiverOk

`func (o *ProductPurchase) GetIsReceiverOk() (*bool, bool)`

GetIsReceiverOk returns a tuple with the IsReceiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReceiver

`func (o *ProductPurchase) SetIsReceiver(v bool)`

SetIsReceiver sets IsReceiver field to given value.


### GetIsSeller

`func (o *ProductPurchase) GetIsSeller() bool`

GetIsSeller returns the IsSeller field if non-nil, zero value otherwise.

### GetIsSellerOk

`func (o *ProductPurchase) GetIsSellerOk() (*bool, bool)`

GetIsSellerOk returns a tuple with the IsSeller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSeller

`func (o *ProductPurchase) SetIsSeller(v bool)`

SetIsSeller sets IsSeller field to given value.


### GetListingCurrentlyAvailable

`func (o *ProductPurchase) GetListingCurrentlyAvailable() bool`

GetListingCurrentlyAvailable returns the ListingCurrentlyAvailable field if non-nil, zero value otherwise.

### GetListingCurrentlyAvailableOk

`func (o *ProductPurchase) GetListingCurrentlyAvailableOk() (*bool, bool)`

GetListingCurrentlyAvailableOk returns a tuple with the ListingCurrentlyAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingCurrentlyAvailable

`func (o *ProductPurchase) SetListingCurrentlyAvailable(v bool)`

SetListingCurrentlyAvailable sets ListingCurrentlyAvailable field to given value.


### GetListingDisplayName

`func (o *ProductPurchase) GetListingDisplayName() string`

GetListingDisplayName returns the ListingDisplayName field if non-nil, zero value otherwise.

### GetListingDisplayNameOk

`func (o *ProductPurchase) GetListingDisplayNameOk() (*string, bool)`

GetListingDisplayNameOk returns a tuple with the ListingDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingDisplayName

`func (o *ProductPurchase) SetListingDisplayName(v string)`

SetListingDisplayName sets ListingDisplayName field to given value.


### GetListingId

`func (o *ProductPurchase) GetListingId() string`

GetListingId returns the ListingId field if non-nil, zero value otherwise.

### GetListingIdOk

`func (o *ProductPurchase) GetListingIdOk() (*string, bool)`

GetListingIdOk returns a tuple with the ListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingId

`func (o *ProductPurchase) SetListingId(v string)`

SetListingId sets ListingId field to given value.


### GetListingImageId

`func (o *ProductPurchase) GetListingImageId() string`

GetListingImageId returns the ListingImageId field if non-nil, zero value otherwise.

### GetListingImageIdOk

`func (o *ProductPurchase) GetListingImageIdOk() (*string, bool)`

GetListingImageIdOk returns a tuple with the ListingImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingImageId

`func (o *ProductPurchase) SetListingImageId(v string)`

SetListingImageId sets ListingImageId field to given value.


### GetListingSubtitle

`func (o *ProductPurchase) GetListingSubtitle() string`

GetListingSubtitle returns the ListingSubtitle field if non-nil, zero value otherwise.

### GetListingSubtitleOk

`func (o *ProductPurchase) GetListingSubtitleOk() (*string, bool)`

GetListingSubtitleOk returns a tuple with the ListingSubtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingSubtitle

`func (o *ProductPurchase) SetListingSubtitle(v string)`

SetListingSubtitle sets ListingSubtitle field to given value.


### GetListingType

`func (o *ProductPurchase) GetListingType() ProductListingType`

GetListingType returns the ListingType field if non-nil, zero value otherwise.

### GetListingTypeOk

`func (o *ProductPurchase) GetListingTypeOk() (*ProductListingType, bool)`

GetListingTypeOk returns a tuple with the ListingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingType

`func (o *ProductPurchase) SetListingType(v ProductListingType)`

SetListingType sets ListingType field to given value.


### GetProducts

`func (o *ProductPurchase) GetProducts() []map[string]interface{}`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *ProductPurchase) GetProductsOk() (*[]map[string]interface{}, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *ProductPurchase) SetProducts(v []map[string]interface{})`

SetProducts sets Products field to given value.


### GetPurchaseActive

`func (o *ProductPurchase) GetPurchaseActive() bool`

GetPurchaseActive returns the PurchaseActive field if non-nil, zero value otherwise.

### GetPurchaseActiveOk

`func (o *ProductPurchase) GetPurchaseActiveOk() (*bool, bool)`

GetPurchaseActiveOk returns a tuple with the PurchaseActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseActive

`func (o *ProductPurchase) SetPurchaseActive(v bool)`

SetPurchaseActive sets PurchaseActive field to given value.


### GetPurchaseContext

`func (o *ProductPurchase) GetPurchaseContext() ProductPurchasePurchaseContext`

GetPurchaseContext returns the PurchaseContext field if non-nil, zero value otherwise.

### GetPurchaseContextOk

`func (o *ProductPurchase) GetPurchaseContextOk() (*ProductPurchasePurchaseContext, bool)`

GetPurchaseContextOk returns a tuple with the PurchaseContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseContext

`func (o *ProductPurchase) SetPurchaseContext(v ProductPurchasePurchaseContext)`

SetPurchaseContext sets PurchaseContext field to given value.


### GetPurchaseCurrentStatus

`func (o *ProductPurchase) GetPurchaseCurrentStatus() string`

GetPurchaseCurrentStatus returns the PurchaseCurrentStatus field if non-nil, zero value otherwise.

### GetPurchaseCurrentStatusOk

`func (o *ProductPurchase) GetPurchaseCurrentStatusOk() (*string, bool)`

GetPurchaseCurrentStatusOk returns a tuple with the PurchaseCurrentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseCurrentStatus

`func (o *ProductPurchase) SetPurchaseCurrentStatus(v string)`

SetPurchaseCurrentStatus sets PurchaseCurrentStatus field to given value.


### GetPurchaseDate

`func (o *ProductPurchase) GetPurchaseDate() time.Time`

GetPurchaseDate returns the PurchaseDate field if non-nil, zero value otherwise.

### GetPurchaseDateOk

`func (o *ProductPurchase) GetPurchaseDateOk() (*time.Time, bool)`

GetPurchaseDateOk returns a tuple with the PurchaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDate

`func (o *ProductPurchase) SetPurchaseDate(v time.Time)`

SetPurchaseDate sets PurchaseDate field to given value.


### GetPurchaseDuration

`func (o *ProductPurchase) GetPurchaseDuration() int32`

GetPurchaseDuration returns the PurchaseDuration field if non-nil, zero value otherwise.

### GetPurchaseDurationOk

`func (o *ProductPurchase) GetPurchaseDurationOk() (*int32, bool)`

GetPurchaseDurationOk returns a tuple with the PurchaseDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDuration

`func (o *ProductPurchase) SetPurchaseDuration(v int32)`

SetPurchaseDuration sets PurchaseDuration field to given value.

### HasPurchaseDuration

`func (o *ProductPurchase) HasPurchaseDuration() bool`

HasPurchaseDuration returns a boolean if a field has been set.

### GetPurchaseDurationType

`func (o *ProductPurchase) GetPurchaseDurationType() string`

GetPurchaseDurationType returns the PurchaseDurationType field if non-nil, zero value otherwise.

### GetPurchaseDurationTypeOk

`func (o *ProductPurchase) GetPurchaseDurationTypeOk() (*string, bool)`

GetPurchaseDurationTypeOk returns a tuple with the PurchaseDurationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDurationType

`func (o *ProductPurchase) SetPurchaseDurationType(v string)`

SetPurchaseDurationType sets PurchaseDurationType field to given value.

### HasPurchaseDurationType

`func (o *ProductPurchase) HasPurchaseDurationType() bool`

HasPurchaseDurationType returns a boolean if a field has been set.

### GetPurchaseEndDate

`func (o *ProductPurchase) GetPurchaseEndDate() time.Time`

GetPurchaseEndDate returns the PurchaseEndDate field if non-nil, zero value otherwise.

### GetPurchaseEndDateOk

`func (o *ProductPurchase) GetPurchaseEndDateOk() (*time.Time, bool)`

GetPurchaseEndDateOk returns a tuple with the PurchaseEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseEndDate

`func (o *ProductPurchase) SetPurchaseEndDate(v time.Time)`

SetPurchaseEndDate sets PurchaseEndDate field to given value.


### GetPurchaseId

`func (o *ProductPurchase) GetPurchaseId() string`

GetPurchaseId returns the PurchaseId field if non-nil, zero value otherwise.

### GetPurchaseIdOk

`func (o *ProductPurchase) GetPurchaseIdOk() (*string, bool)`

GetPurchaseIdOk returns a tuple with the PurchaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseId

`func (o *ProductPurchase) SetPurchaseId(v string)`

SetPurchaseId sets PurchaseId field to given value.


### GetPurchaseLatest

`func (o *ProductPurchase) GetPurchaseLatest() bool`

GetPurchaseLatest returns the PurchaseLatest field if non-nil, zero value otherwise.

### GetPurchaseLatestOk

`func (o *ProductPurchase) GetPurchaseLatestOk() (*bool, bool)`

GetPurchaseLatestOk returns a tuple with the PurchaseLatest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseLatest

`func (o *ProductPurchase) SetPurchaseLatest(v bool)`

SetPurchaseLatest sets PurchaseLatest field to given value.


### GetPurchasePrice

`func (o *ProductPurchase) GetPurchasePrice() int32`

GetPurchasePrice returns the PurchasePrice field if non-nil, zero value otherwise.

### GetPurchasePriceOk

`func (o *ProductPurchase) GetPurchasePriceOk() (*int32, bool)`

GetPurchasePriceOk returns a tuple with the PurchasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasePrice

`func (o *ProductPurchase) SetPurchasePrice(v int32)`

SetPurchasePrice sets PurchasePrice field to given value.


### GetPurchaseQuantity

`func (o *ProductPurchase) GetPurchaseQuantity() int32`

GetPurchaseQuantity returns the PurchaseQuantity field if non-nil, zero value otherwise.

### GetPurchaseQuantityOk

`func (o *ProductPurchase) GetPurchaseQuantityOk() (*int32, bool)`

GetPurchaseQuantityOk returns a tuple with the PurchaseQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseQuantity

`func (o *ProductPurchase) SetPurchaseQuantity(v int32)`

SetPurchaseQuantity sets PurchaseQuantity field to given value.


### GetPurchaseStartDate

`func (o *ProductPurchase) GetPurchaseStartDate() time.Time`

GetPurchaseStartDate returns the PurchaseStartDate field if non-nil, zero value otherwise.

### GetPurchaseStartDateOk

`func (o *ProductPurchase) GetPurchaseStartDateOk() (*time.Time, bool)`

GetPurchaseStartDateOk returns a tuple with the PurchaseStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseStartDate

`func (o *ProductPurchase) SetPurchaseStartDate(v time.Time)`

SetPurchaseStartDate sets PurchaseStartDate field to given value.


### GetPurchaseToken

`func (o *ProductPurchase) GetPurchaseToken() map[string]interface{}`

GetPurchaseToken returns the PurchaseToken field if non-nil, zero value otherwise.

### GetPurchaseTokenOk

`func (o *ProductPurchase) GetPurchaseTokenOk() (*map[string]interface{}, bool)`

GetPurchaseTokenOk returns a tuple with the PurchaseToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseToken

`func (o *ProductPurchase) SetPurchaseToken(v map[string]interface{})`

SetPurchaseToken sets PurchaseToken field to given value.


### SetPurchaseTokenNil

`func (o *ProductPurchase) SetPurchaseTokenNil(b bool)`

 SetPurchaseTokenNil sets the value for PurchaseToken to be an explicit nil

### UnsetPurchaseToken
`func (o *ProductPurchase) UnsetPurchaseToken()`

UnsetPurchaseToken ensures that no value is present for PurchaseToken, not even an explicit nil
### GetPurchaseType

`func (o *ProductPurchase) GetPurchaseType() string`

GetPurchaseType returns the PurchaseType field if non-nil, zero value otherwise.

### GetPurchaseTypeOk

`func (o *ProductPurchase) GetPurchaseTypeOk() (*string, bool)`

GetPurchaseTypeOk returns a tuple with the PurchaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseType

`func (o *ProductPurchase) SetPurchaseType(v string)`

SetPurchaseType sets PurchaseType field to given value.


### GetPurchaseUnitPrice

`func (o *ProductPurchase) GetPurchaseUnitPrice() int32`

GetPurchaseUnitPrice returns the PurchaseUnitPrice field if non-nil, zero value otherwise.

### GetPurchaseUnitPriceOk

`func (o *ProductPurchase) GetPurchaseUnitPriceOk() (*int32, bool)`

GetPurchaseUnitPriceOk returns a tuple with the PurchaseUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseUnitPrice

`func (o *ProductPurchase) SetPurchaseUnitPrice(v int32)`

SetPurchaseUnitPrice sets PurchaseUnitPrice field to given value.


### GetReceiverDisplayName

`func (o *ProductPurchase) GetReceiverDisplayName() string`

GetReceiverDisplayName returns the ReceiverDisplayName field if non-nil, zero value otherwise.

### GetReceiverDisplayNameOk

`func (o *ProductPurchase) GetReceiverDisplayNameOk() (*string, bool)`

GetReceiverDisplayNameOk returns a tuple with the ReceiverDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverDisplayName

`func (o *ProductPurchase) SetReceiverDisplayName(v string)`

SetReceiverDisplayName sets ReceiverDisplayName field to given value.


### GetReceiverId

`func (o *ProductPurchase) GetReceiverId() string`

GetReceiverId returns the ReceiverId field if non-nil, zero value otherwise.

### GetReceiverIdOk

`func (o *ProductPurchase) GetReceiverIdOk() (*string, bool)`

GetReceiverIdOk returns a tuple with the ReceiverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverId

`func (o *ProductPurchase) SetReceiverId(v string)`

SetReceiverId sets ReceiverId field to given value.


### GetRecurrable

`func (o *ProductPurchase) GetRecurrable() bool`

GetRecurrable returns the Recurrable field if non-nil, zero value otherwise.

### GetRecurrableOk

`func (o *ProductPurchase) GetRecurrableOk() (*bool, bool)`

GetRecurrableOk returns a tuple with the Recurrable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrable

`func (o *ProductPurchase) SetRecurrable(v bool)`

SetRecurrable sets Recurrable field to given value.


### GetRefundable

`func (o *ProductPurchase) GetRefundable() bool`

GetRefundable returns the Refundable field if non-nil, zero value otherwise.

### GetRefundableOk

`func (o *ProductPurchase) GetRefundableOk() (*bool, bool)`

GetRefundableOk returns a tuple with the Refundable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundable

`func (o *ProductPurchase) SetRefundable(v bool)`

SetRefundable sets Refundable field to given value.


### GetSellerDisplayName

`func (o *ProductPurchase) GetSellerDisplayName() string`

GetSellerDisplayName returns the SellerDisplayName field if non-nil, zero value otherwise.

### GetSellerDisplayNameOk

`func (o *ProductPurchase) GetSellerDisplayNameOk() (*string, bool)`

GetSellerDisplayNameOk returns a tuple with the SellerDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerDisplayName

`func (o *ProductPurchase) SetSellerDisplayName(v string)`

SetSellerDisplayName sets SellerDisplayName field to given value.


### GetSellerId

`func (o *ProductPurchase) GetSellerId() string`

GetSellerId returns the SellerId field if non-nil, zero value otherwise.

### GetSellerIdOk

`func (o *ProductPurchase) GetSellerIdOk() (*string, bool)`

GetSellerIdOk returns a tuple with the SellerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerId

`func (o *ProductPurchase) SetSellerId(v string)`

SetSellerId sets SellerId field to given value.


### GetStackable

`func (o *ProductPurchase) GetStackable() bool`

GetStackable returns the Stackable field if non-nil, zero value otherwise.

### GetStackableOk

`func (o *ProductPurchase) GetStackableOk() (*bool, bool)`

GetStackableOk returns a tuple with the Stackable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackable

`func (o *ProductPurchase) SetStackable(v bool)`

SetStackable sets Stackable field to given value.


### GetWillRecur

`func (o *ProductPurchase) GetWillRecur() bool`

GetWillRecur returns the WillRecur field if non-nil, zero value otherwise.

### GetWillRecurOk

`func (o *ProductPurchase) GetWillRecurOk() (*bool, bool)`

GetWillRecurOk returns a tuple with the WillRecur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWillRecur

`func (o *ProductPurchase) SetWillRecur(v bool)`

SetWillRecur sets WillRecur field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


