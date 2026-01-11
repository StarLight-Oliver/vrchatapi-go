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
	"encoding/json"
	"fmt"
	"time"
)

// checks if the ProductPurchase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductPurchase{}

// ProductPurchase struct for ProductPurchase
type ProductPurchase struct {
	BuyerDisplayName string `json:"buyerDisplayName"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	BuyerId                   string                         `json:"buyerId"`
	FirstParty                *bool                          `json:"firstParty,omitempty"`
	IsBuyer                   bool                           `json:"isBuyer"`
	IsGift                    bool                           `json:"isGift"`
	IsReceiver                bool                           `json:"isReceiver"`
	IsSeller                  bool                           `json:"isSeller"`
	ListingCurrentlyAvailable bool                           `json:"listingCurrentlyAvailable"`
	ListingDisplayName        string                         `json:"listingDisplayName"`
	ListingId                 string                         `json:"listingId"`
	ListingImageId            string                         `json:"listingImageId"`
	ListingSubtitle           string                         `json:"listingSubtitle"`
	ListingType               ProductListingType             `json:"listingType"`
	Products                  []map[string]interface{}       `json:"products"`
	PurchaseActive            bool                           `json:"purchaseActive"`
	PurchaseContext           ProductPurchasePurchaseContext `json:"purchaseContext"`
	PurchaseCurrentStatus     string                         `json:"purchaseCurrentStatus"`
	PurchaseDate              time.Time                      `json:"purchaseDate"`
	PurchaseDuration          *int32                         `json:"purchaseDuration,omitempty"`
	PurchaseDurationType      *string                        `json:"purchaseDurationType,omitempty"`
	PurchaseEndDate           time.Time                      `json:"purchaseEndDate"`
	PurchaseId                string                         `json:"purchaseId"`
	PurchaseLatest            bool                           `json:"purchaseLatest"`
	PurchasePrice             int32                          `json:"purchasePrice"`
	PurchaseQuantity          int32                          `json:"purchaseQuantity"`
	PurchaseStartDate         time.Time                      `json:"purchaseStartDate"`
	PurchaseToken             map[string]interface{}         `json:"purchaseToken"`
	PurchaseType              string                         `json:"purchaseType"`
	PurchaseUnitPrice         int32                          `json:"purchaseUnitPrice"`
	ReceiverDisplayName       string                         `json:"receiverDisplayName"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	ReceiverId        string `json:"receiverId"`
	Recurrable        bool   `json:"recurrable"`
	Refundable        bool   `json:"refundable"`
	SellerDisplayName string `json:"sellerDisplayName"`
	// A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	SellerId  string `json:"sellerId"`
	Stackable bool   `json:"stackable"`
	WillRecur bool   `json:"willRecur"`
}

type _ProductPurchase ProductPurchase

// NewProductPurchase instantiates a new ProductPurchase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductPurchase(buyerDisplayName string, buyerId string, isBuyer bool, isGift bool, isReceiver bool, isSeller bool, listingCurrentlyAvailable bool, listingDisplayName string, listingId string, listingImageId string, listingSubtitle string, listingType ProductListingType, products []map[string]interface{}, purchaseActive bool, purchaseContext ProductPurchasePurchaseContext, purchaseCurrentStatus string, purchaseDate time.Time, purchaseEndDate time.Time, purchaseId string, purchaseLatest bool, purchasePrice int32, purchaseQuantity int32, purchaseStartDate time.Time, purchaseToken map[string]interface{}, purchaseType string, purchaseUnitPrice int32, receiverDisplayName string, receiverId string, recurrable bool, refundable bool, sellerDisplayName string, sellerId string, stackable bool, willRecur bool) *ProductPurchase {
	this := ProductPurchase{}
	this.BuyerDisplayName = buyerDisplayName
	this.BuyerId = buyerId
	this.IsBuyer = isBuyer
	this.IsGift = isGift
	this.IsReceiver = isReceiver
	this.IsSeller = isSeller
	this.ListingCurrentlyAvailable = listingCurrentlyAvailable
	this.ListingDisplayName = listingDisplayName
	this.ListingId = listingId
	this.ListingImageId = listingImageId
	this.ListingSubtitle = listingSubtitle
	this.ListingType = listingType
	this.Products = products
	this.PurchaseActive = purchaseActive
	this.PurchaseContext = purchaseContext
	this.PurchaseCurrentStatus = purchaseCurrentStatus
	this.PurchaseDate = purchaseDate
	this.PurchaseEndDate = purchaseEndDate
	this.PurchaseId = purchaseId
	this.PurchaseLatest = purchaseLatest
	this.PurchasePrice = purchasePrice
	this.PurchaseQuantity = purchaseQuantity
	this.PurchaseStartDate = purchaseStartDate
	this.PurchaseToken = purchaseToken
	this.PurchaseType = purchaseType
	this.PurchaseUnitPrice = purchaseUnitPrice
	this.ReceiverDisplayName = receiverDisplayName
	this.ReceiverId = receiverId
	this.Recurrable = recurrable
	this.Refundable = refundable
	this.SellerDisplayName = sellerDisplayName
	this.SellerId = sellerId
	this.Stackable = stackable
	this.WillRecur = willRecur
	return &this
}

// NewProductPurchaseWithDefaults instantiates a new ProductPurchase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductPurchaseWithDefaults() *ProductPurchase {
	this := ProductPurchase{}
	var listingType ProductListingType = ProductListingType_SUBSCRIPTION
	this.ListingType = listingType
	return &this
}

// GetBuyerDisplayName returns the BuyerDisplayName field value
func (o *ProductPurchase) GetBuyerDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BuyerDisplayName
}

// GetBuyerDisplayNameOk returns a tuple with the BuyerDisplayName field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetBuyerDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuyerDisplayName, true
}

// SetBuyerDisplayName sets field value
func (o *ProductPurchase) SetBuyerDisplayName(v string) {
	o.BuyerDisplayName = v
}

// GetBuyerId returns the BuyerId field value
func (o *ProductPurchase) GetBuyerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BuyerId
}

// GetBuyerIdOk returns a tuple with the BuyerId field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetBuyerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuyerId, true
}

// SetBuyerId sets field value
func (o *ProductPurchase) SetBuyerId(v string) {
	o.BuyerId = v
}

// GetFirstParty returns the FirstParty field value if set, zero value otherwise.
func (o *ProductPurchase) GetFirstParty() bool {
	if o == nil || IsNil(o.FirstParty) {
		var ret bool
		return ret
	}
	return *o.FirstParty
}

// GetFirstPartyOk returns a tuple with the FirstParty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetFirstPartyOk() (*bool, bool) {
	if o == nil || IsNil(o.FirstParty) {
		return nil, false
	}
	return o.FirstParty, true
}

// HasFirstParty returns a boolean if a field has been set.
func (o *ProductPurchase) HasFirstParty() bool {
	if o != nil && !IsNil(o.FirstParty) {
		return true
	}

	return false
}

// SetFirstParty gets a reference to the given bool and assigns it to the FirstParty field.
func (o *ProductPurchase) SetFirstParty(v bool) {
	o.FirstParty = &v
}

// GetIsBuyer returns the IsBuyer field value
func (o *ProductPurchase) GetIsBuyer() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsBuyer
}

// GetIsBuyerOk returns a tuple with the IsBuyer field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetIsBuyerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsBuyer, true
}

// SetIsBuyer sets field value
func (o *ProductPurchase) SetIsBuyer(v bool) {
	o.IsBuyer = v
}

// GetIsGift returns the IsGift field value
func (o *ProductPurchase) GetIsGift() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsGift
}

// GetIsGiftOk returns a tuple with the IsGift field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetIsGiftOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsGift, true
}

// SetIsGift sets field value
func (o *ProductPurchase) SetIsGift(v bool) {
	o.IsGift = v
}

// GetIsReceiver returns the IsReceiver field value
func (o *ProductPurchase) GetIsReceiver() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsReceiver
}

// GetIsReceiverOk returns a tuple with the IsReceiver field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetIsReceiverOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsReceiver, true
}

// SetIsReceiver sets field value
func (o *ProductPurchase) SetIsReceiver(v bool) {
	o.IsReceiver = v
}

// GetIsSeller returns the IsSeller field value
func (o *ProductPurchase) GetIsSeller() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSeller
}

// GetIsSellerOk returns a tuple with the IsSeller field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetIsSellerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSeller, true
}

// SetIsSeller sets field value
func (o *ProductPurchase) SetIsSeller(v bool) {
	o.IsSeller = v
}

// GetListingCurrentlyAvailable returns the ListingCurrentlyAvailable field value
func (o *ProductPurchase) GetListingCurrentlyAvailable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ListingCurrentlyAvailable
}

// GetListingCurrentlyAvailableOk returns a tuple with the ListingCurrentlyAvailable field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetListingCurrentlyAvailableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListingCurrentlyAvailable, true
}

// SetListingCurrentlyAvailable sets field value
func (o *ProductPurchase) SetListingCurrentlyAvailable(v bool) {
	o.ListingCurrentlyAvailable = v
}

// GetListingDisplayName returns the ListingDisplayName field value
func (o *ProductPurchase) GetListingDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ListingDisplayName
}

// GetListingDisplayNameOk returns a tuple with the ListingDisplayName field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetListingDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListingDisplayName, true
}

// SetListingDisplayName sets field value
func (o *ProductPurchase) SetListingDisplayName(v string) {
	o.ListingDisplayName = v
}

// GetListingId returns the ListingId field value
func (o *ProductPurchase) GetListingId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ListingId
}

// GetListingIdOk returns a tuple with the ListingId field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetListingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListingId, true
}

// SetListingId sets field value
func (o *ProductPurchase) SetListingId(v string) {
	o.ListingId = v
}

// GetListingImageId returns the ListingImageId field value
func (o *ProductPurchase) GetListingImageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ListingImageId
}

// GetListingImageIdOk returns a tuple with the ListingImageId field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetListingImageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListingImageId, true
}

// SetListingImageId sets field value
func (o *ProductPurchase) SetListingImageId(v string) {
	o.ListingImageId = v
}

// GetListingSubtitle returns the ListingSubtitle field value
func (o *ProductPurchase) GetListingSubtitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ListingSubtitle
}

// GetListingSubtitleOk returns a tuple with the ListingSubtitle field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetListingSubtitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListingSubtitle, true
}

// SetListingSubtitle sets field value
func (o *ProductPurchase) SetListingSubtitle(v string) {
	o.ListingSubtitle = v
}

// GetListingType returns the ListingType field value
func (o *ProductPurchase) GetListingType() ProductListingType {
	if o == nil {
		var ret ProductListingType
		return ret
	}

	return o.ListingType
}

// GetListingTypeOk returns a tuple with the ListingType field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetListingTypeOk() (*ProductListingType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListingType, true
}

// SetListingType sets field value
func (o *ProductPurchase) SetListingType(v ProductListingType) {
	o.ListingType = v
}

// GetProducts returns the Products field value
func (o *ProductPurchase) GetProducts() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Products
}

// GetProductsOk returns a tuple with the Products field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetProductsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Products, true
}

// SetProducts sets field value
func (o *ProductPurchase) SetProducts(v []map[string]interface{}) {
	o.Products = v
}

// GetPurchaseActive returns the PurchaseActive field value
func (o *ProductPurchase) GetPurchaseActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PurchaseActive
}

// GetPurchaseActiveOk returns a tuple with the PurchaseActive field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetPurchaseActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchaseActive, true
}

// SetPurchaseActive sets field value
func (o *ProductPurchase) SetPurchaseActive(v bool) {
	o.PurchaseActive = v
}

// GetPurchaseContext returns the PurchaseContext field value
func (o *ProductPurchase) GetPurchaseContext() ProductPurchasePurchaseContext {
	if o == nil {
		var ret ProductPurchasePurchaseContext
		return ret
	}

	return o.PurchaseContext
}

// GetPurchaseContextOk returns a tuple with the PurchaseContext field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetPurchaseContextOk() (*ProductPurchasePurchaseContext, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchaseContext, true
}

// SetPurchaseContext sets field value
func (o *ProductPurchase) SetPurchaseContext(v ProductPurchasePurchaseContext) {
	o.PurchaseContext = v
}

// GetPurchaseCurrentStatus returns the PurchaseCurrentStatus field value
func (o *ProductPurchase) GetPurchaseCurrentStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PurchaseCurrentStatus
}

// GetPurchaseCurrentStatusOk returns a tuple with the PurchaseCurrentStatus field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetPurchaseCurrentStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchaseCurrentStatus, true
}

// SetPurchaseCurrentStatus sets field value
func (o *ProductPurchase) SetPurchaseCurrentStatus(v string) {
	o.PurchaseCurrentStatus = v
}

// GetPurchaseDate returns the PurchaseDate field value
func (o *ProductPurchase) GetPurchaseDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.PurchaseDate
}

// GetPurchaseDateOk returns a tuple with the PurchaseDate field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetPurchaseDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchaseDate, true
}

// SetPurchaseDate sets field value
func (o *ProductPurchase) SetPurchaseDate(v time.Time) {
	o.PurchaseDate = v
}

// GetPurchaseDuration returns the PurchaseDuration field value if set, zero value otherwise.
func (o *ProductPurchase) GetPurchaseDuration() int32 {
	if o == nil || IsNil(o.PurchaseDuration) {
		var ret int32
		return ret
	}
	return *o.PurchaseDuration
}

// GetPurchaseDurationOk returns a tuple with the PurchaseDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetPurchaseDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.PurchaseDuration) {
		return nil, false
	}
	return o.PurchaseDuration, true
}

// HasPurchaseDuration returns a boolean if a field has been set.
func (o *ProductPurchase) HasPurchaseDuration() bool {
	if o != nil && !IsNil(o.PurchaseDuration) {
		return true
	}

	return false
}

// SetPurchaseDuration gets a reference to the given int32 and assigns it to the PurchaseDuration field.
func (o *ProductPurchase) SetPurchaseDuration(v int32) {
	o.PurchaseDuration = &v
}

// GetPurchaseDurationType returns the PurchaseDurationType field value if set, zero value otherwise.
func (o *ProductPurchase) GetPurchaseDurationType() string {
	if o == nil || IsNil(o.PurchaseDurationType) {
		var ret string
		return ret
	}
	return *o.PurchaseDurationType
}

// GetPurchaseDurationTypeOk returns a tuple with the PurchaseDurationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetPurchaseDurationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PurchaseDurationType) {
		return nil, false
	}
	return o.PurchaseDurationType, true
}

// HasPurchaseDurationType returns a boolean if a field has been set.
func (o *ProductPurchase) HasPurchaseDurationType() bool {
	if o != nil && !IsNil(o.PurchaseDurationType) {
		return true
	}

	return false
}

// SetPurchaseDurationType gets a reference to the given string and assigns it to the PurchaseDurationType field.
func (o *ProductPurchase) SetPurchaseDurationType(v string) {
	o.PurchaseDurationType = &v
}

// GetPurchaseEndDate returns the PurchaseEndDate field value
func (o *ProductPurchase) GetPurchaseEndDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.PurchaseEndDate
}

// GetPurchaseEndDateOk returns a tuple with the PurchaseEndDate field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetPurchaseEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchaseEndDate, true
}

// SetPurchaseEndDate sets field value
func (o *ProductPurchase) SetPurchaseEndDate(v time.Time) {
	o.PurchaseEndDate = v
}

// GetPurchaseId returns the PurchaseId field value
func (o *ProductPurchase) GetPurchaseId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PurchaseId
}

// GetPurchaseIdOk returns a tuple with the PurchaseId field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetPurchaseIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchaseId, true
}

// SetPurchaseId sets field value
func (o *ProductPurchase) SetPurchaseId(v string) {
	o.PurchaseId = v
}

// GetPurchaseLatest returns the PurchaseLatest field value
func (o *ProductPurchase) GetPurchaseLatest() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PurchaseLatest
}

// GetPurchaseLatestOk returns a tuple with the PurchaseLatest field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetPurchaseLatestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchaseLatest, true
}

// SetPurchaseLatest sets field value
func (o *ProductPurchase) SetPurchaseLatest(v bool) {
	o.PurchaseLatest = v
}

// GetPurchasePrice returns the PurchasePrice field value
func (o *ProductPurchase) GetPurchasePrice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PurchasePrice
}

// GetPurchasePriceOk returns a tuple with the PurchasePrice field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetPurchasePriceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchasePrice, true
}

// SetPurchasePrice sets field value
func (o *ProductPurchase) SetPurchasePrice(v int32) {
	o.PurchasePrice = v
}

// GetPurchaseQuantity returns the PurchaseQuantity field value
func (o *ProductPurchase) GetPurchaseQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PurchaseQuantity
}

// GetPurchaseQuantityOk returns a tuple with the PurchaseQuantity field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetPurchaseQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchaseQuantity, true
}

// SetPurchaseQuantity sets field value
func (o *ProductPurchase) SetPurchaseQuantity(v int32) {
	o.PurchaseQuantity = v
}

// GetPurchaseStartDate returns the PurchaseStartDate field value
func (o *ProductPurchase) GetPurchaseStartDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.PurchaseStartDate
}

// GetPurchaseStartDateOk returns a tuple with the PurchaseStartDate field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetPurchaseStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchaseStartDate, true
}

// SetPurchaseStartDate sets field value
func (o *ProductPurchase) SetPurchaseStartDate(v time.Time) {
	o.PurchaseStartDate = v
}

// GetPurchaseToken returns the PurchaseToken field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *ProductPurchase) GetPurchaseToken() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.PurchaseToken
}

// GetPurchaseTokenOk returns a tuple with the PurchaseToken field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductPurchase) GetPurchaseTokenOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.PurchaseToken) {
		return map[string]interface{}{}, false
	}
	return o.PurchaseToken, true
}

// SetPurchaseToken sets field value
func (o *ProductPurchase) SetPurchaseToken(v map[string]interface{}) {
	o.PurchaseToken = v
}

// GetPurchaseType returns the PurchaseType field value
func (o *ProductPurchase) GetPurchaseType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PurchaseType
}

// GetPurchaseTypeOk returns a tuple with the PurchaseType field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetPurchaseTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchaseType, true
}

// SetPurchaseType sets field value
func (o *ProductPurchase) SetPurchaseType(v string) {
	o.PurchaseType = v
}

// GetPurchaseUnitPrice returns the PurchaseUnitPrice field value
func (o *ProductPurchase) GetPurchaseUnitPrice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PurchaseUnitPrice
}

// GetPurchaseUnitPriceOk returns a tuple with the PurchaseUnitPrice field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetPurchaseUnitPriceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchaseUnitPrice, true
}

// SetPurchaseUnitPrice sets field value
func (o *ProductPurchase) SetPurchaseUnitPrice(v int32) {
	o.PurchaseUnitPrice = v
}

// GetReceiverDisplayName returns the ReceiverDisplayName field value
func (o *ProductPurchase) GetReceiverDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReceiverDisplayName
}

// GetReceiverDisplayNameOk returns a tuple with the ReceiverDisplayName field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetReceiverDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceiverDisplayName, true
}

// SetReceiverDisplayName sets field value
func (o *ProductPurchase) SetReceiverDisplayName(v string) {
	o.ReceiverDisplayName = v
}

// GetReceiverId returns the ReceiverId field value
func (o *ProductPurchase) GetReceiverId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReceiverId
}

// GetReceiverIdOk returns a tuple with the ReceiverId field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetReceiverIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceiverId, true
}

// SetReceiverId sets field value
func (o *ProductPurchase) SetReceiverId(v string) {
	o.ReceiverId = v
}

// GetRecurrable returns the Recurrable field value
func (o *ProductPurchase) GetRecurrable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Recurrable
}

// GetRecurrableOk returns a tuple with the Recurrable field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetRecurrableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recurrable, true
}

// SetRecurrable sets field value
func (o *ProductPurchase) SetRecurrable(v bool) {
	o.Recurrable = v
}

// GetRefundable returns the Refundable field value
func (o *ProductPurchase) GetRefundable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Refundable
}

// GetRefundableOk returns a tuple with the Refundable field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetRefundableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Refundable, true
}

// SetRefundable sets field value
func (o *ProductPurchase) SetRefundable(v bool) {
	o.Refundable = v
}

// GetSellerDisplayName returns the SellerDisplayName field value
func (o *ProductPurchase) GetSellerDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SellerDisplayName
}

// GetSellerDisplayNameOk returns a tuple with the SellerDisplayName field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetSellerDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellerDisplayName, true
}

// SetSellerDisplayName sets field value
func (o *ProductPurchase) SetSellerDisplayName(v string) {
	o.SellerDisplayName = v
}

// GetSellerId returns the SellerId field value
func (o *ProductPurchase) GetSellerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SellerId
}

// GetSellerIdOk returns a tuple with the SellerId field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetSellerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellerId, true
}

// SetSellerId sets field value
func (o *ProductPurchase) SetSellerId(v string) {
	o.SellerId = v
}

// GetStackable returns the Stackable field value
func (o *ProductPurchase) GetStackable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Stackable
}

// GetStackableOk returns a tuple with the Stackable field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetStackableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stackable, true
}

// SetStackable sets field value
func (o *ProductPurchase) SetStackable(v bool) {
	o.Stackable = v
}

// GetWillRecur returns the WillRecur field value
func (o *ProductPurchase) GetWillRecur() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.WillRecur
}

// GetWillRecurOk returns a tuple with the WillRecur field value
// and a boolean to check if the value has been set.
func (o *ProductPurchase) GetWillRecurOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WillRecur, true
}

// SetWillRecur sets field value
func (o *ProductPurchase) SetWillRecur(v bool) {
	o.WillRecur = v
}

func (o ProductPurchase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductPurchase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["buyerDisplayName"] = o.BuyerDisplayName
	toSerialize["buyerId"] = o.BuyerId
	if !IsNil(o.FirstParty) {
		toSerialize["firstParty"] = o.FirstParty
	}
	toSerialize["isBuyer"] = o.IsBuyer
	toSerialize["isGift"] = o.IsGift
	toSerialize["isReceiver"] = o.IsReceiver
	toSerialize["isSeller"] = o.IsSeller
	toSerialize["listingCurrentlyAvailable"] = o.ListingCurrentlyAvailable
	toSerialize["listingDisplayName"] = o.ListingDisplayName
	toSerialize["listingId"] = o.ListingId
	toSerialize["listingImageId"] = o.ListingImageId
	toSerialize["listingSubtitle"] = o.ListingSubtitle
	toSerialize["listingType"] = o.ListingType
	toSerialize["products"] = o.Products
	toSerialize["purchaseActive"] = o.PurchaseActive
	toSerialize["purchaseContext"] = o.PurchaseContext
	toSerialize["purchaseCurrentStatus"] = o.PurchaseCurrentStatus
	toSerialize["purchaseDate"] = o.PurchaseDate
	if !IsNil(o.PurchaseDuration) {
		toSerialize["purchaseDuration"] = o.PurchaseDuration
	}
	if !IsNil(o.PurchaseDurationType) {
		toSerialize["purchaseDurationType"] = o.PurchaseDurationType
	}
	toSerialize["purchaseEndDate"] = o.PurchaseEndDate
	toSerialize["purchaseId"] = o.PurchaseId
	toSerialize["purchaseLatest"] = o.PurchaseLatest
	toSerialize["purchasePrice"] = o.PurchasePrice
	toSerialize["purchaseQuantity"] = o.PurchaseQuantity
	toSerialize["purchaseStartDate"] = o.PurchaseStartDate
	if o.PurchaseToken != nil {
		toSerialize["purchaseToken"] = o.PurchaseToken
	}
	toSerialize["purchaseType"] = o.PurchaseType
	toSerialize["purchaseUnitPrice"] = o.PurchaseUnitPrice
	toSerialize["receiverDisplayName"] = o.ReceiverDisplayName
	toSerialize["receiverId"] = o.ReceiverId
	toSerialize["recurrable"] = o.Recurrable
	toSerialize["refundable"] = o.Refundable
	toSerialize["sellerDisplayName"] = o.SellerDisplayName
	toSerialize["sellerId"] = o.SellerId
	toSerialize["stackable"] = o.Stackable
	toSerialize["willRecur"] = o.WillRecur
	return toSerialize, nil
}

func (o *ProductPurchase) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"buyerDisplayName",
		"buyerId",
		"isBuyer",
		"isGift",
		"isReceiver",
		"isSeller",
		"listingCurrentlyAvailable",
		"listingDisplayName",
		"listingId",
		"listingImageId",
		"listingSubtitle",
		"listingType",
		"products",
		"purchaseActive",
		"purchaseContext",
		"purchaseCurrentStatus",
		"purchaseDate",
		"purchaseEndDate",
		"purchaseId",
		"purchaseLatest",
		"purchasePrice",
		"purchaseQuantity",
		"purchaseStartDate",
		"purchaseToken",
		"purchaseType",
		"purchaseUnitPrice",
		"receiverDisplayName",
		"receiverId",
		"recurrable",
		"refundable",
		"sellerDisplayName",
		"sellerId",
		"stackable",
		"willRecur",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varProductPurchase := _ProductPurchase{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProductPurchase)

	if err != nil {
		return err
	}

	*o = ProductPurchase(varProductPurchase)

	return err
}

type NullableProductPurchase struct {
	value *ProductPurchase
	isSet bool
}

func (v NullableProductPurchase) Get() *ProductPurchase {
	return v.value
}

func (v *NullableProductPurchase) Set(val *ProductPurchase) {
	v.value = val
	v.isSet = true
}

func (v NullableProductPurchase) IsSet() bool {
	return v.isSet
}

func (v *NullableProductPurchase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductPurchase(val *ProductPurchase) *NullableProductPurchase {
	return &NullableProductPurchase{value: val, isSet: true}
}

func (v NullableProductPurchase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductPurchase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
