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

// checks if the ProductListing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductListing{}

// ProductListing struct for ProductListing
type ProductListing struct {
	Active               bool                     `json:"active"`
	BuyerRefundable      bool                     `json:"buyerRefundable"`
	Description          string                   `json:"description"`
	DisplayName          string                   `json:"displayName"`
	Duration             NullableInt32            `json:"duration,omitempty"`
	DurationType         NullableString           `json:"durationType,omitempty"`
	GroupIcon            *string                  `json:"groupIcon,omitempty"`
	GroupId              *string                  `json:"groupId,omitempty"`
	GroupName            NullableString           `json:"groupName,omitempty"`
	HasAvatar            bool                     `json:"hasAvatar"`
	HasUdon              bool                     `json:"hasUdon"`
	HydratedProducts     []Product                `json:"hydratedProducts,omitempty"`
	Id                   string                   `json:"id"`
	ImageId              *string                  `json:"imageId,omitempty"`
	ImageUrl             NullableString           `json:"imageUrl,omitempty"`
	ListingType          ProductListingType       `json:"listingType"`
	ListingVariants      []ProductListingVariant  `json:"listingVariants,omitempty"`
	Permanent            *bool                    `json:"permanent,omitempty"`
	PriceTokens          int32                    `json:"priceTokens"`
	ProductIds           []string                 `json:"productIds"`
	ProductType          ProductType              `json:"productType"`
	Products             []map[string]interface{} `json:"products"`
	Quantifiable         *bool                    `json:"quantifiable,omitempty"`
	Recurrable           bool                     `json:"recurrable"`
	Refundable           bool                     `json:"refundable"`
	SellerDisplayName    string                   `json:"sellerDisplayName"`
	SellerId             string                   `json:"sellerId"`
	SoldByVrc            *bool                    `json:"soldByVrc,omitempty"`
	Stackable            bool                     `json:"stackable"`
	StoreIds             []string                 `json:"storeIds"`
	Subtitle             *string                  `json:"subtitle,omitempty"`
	Tags                 []string                 `json:"tags,omitempty"`
	VrcPlusDiscountPrice *int32                   `json:"vrcPlusDiscountPrice,omitempty"`
	WhenToExpire         NullableTime             `json:"whenToExpire,omitempty"`
}

type _ProductListing ProductListing

// NewProductListing instantiates a new ProductListing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductListing(active bool, buyerRefundable bool, description string, displayName string, hasAvatar bool, hasUdon bool, id string, listingType ProductListingType, priceTokens int32, productIds []string, productType ProductType, products []map[string]interface{}, recurrable bool, refundable bool, sellerDisplayName string, sellerId string, stackable bool, storeIds []string) *ProductListing {
	this := ProductListing{}
	this.Active = active
	this.BuyerRefundable = buyerRefundable
	this.Description = description
	this.DisplayName = displayName
	this.HasAvatar = hasAvatar
	this.HasUdon = hasUdon
	this.Id = id
	this.ListingType = listingType
	this.PriceTokens = priceTokens
	this.ProductIds = productIds
	this.ProductType = productType
	this.Products = products
	this.Recurrable = recurrable
	this.Refundable = refundable
	this.SellerDisplayName = sellerDisplayName
	this.SellerId = sellerId
	this.Stackable = stackable
	this.StoreIds = storeIds
	return &this
}

// NewProductListingWithDefaults instantiates a new ProductListing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductListingWithDefaults() *ProductListing {
	this := ProductListing{}
	var listingType ProductListingType = ProductListingType_SUBSCRIPTION
	this.ListingType = listingType
	var productType ProductType = ProductType_UDON
	this.ProductType = productType
	return &this
}

// GetActive returns the Active field value
func (o *ProductListing) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *ProductListing) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *ProductListing) SetActive(v bool) {
	o.Active = v
}

// GetBuyerRefundable returns the BuyerRefundable field value
func (o *ProductListing) GetBuyerRefundable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BuyerRefundable
}

// GetBuyerRefundableOk returns a tuple with the BuyerRefundable field value
// and a boolean to check if the value has been set.
func (o *ProductListing) GetBuyerRefundableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuyerRefundable, true
}

// SetBuyerRefundable sets field value
func (o *ProductListing) SetBuyerRefundable(v bool) {
	o.BuyerRefundable = v
}

// GetDescription returns the Description field value
func (o *ProductListing) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ProductListing) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ProductListing) SetDescription(v string) {
	o.Description = v
}

// GetDisplayName returns the DisplayName field value
func (o *ProductListing) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *ProductListing) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *ProductListing) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductListing) GetDuration() int32 {
	if o == nil || IsNil(o.Duration.Get()) {
		var ret int32
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductListing) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *ProductListing) HasDuration() bool {
	if o != nil && o.Duration.IsSet() {
		return true
	}

	return false
}

// SetDuration gets a reference to the given NullableInt32 and assigns it to the Duration field.
func (o *ProductListing) SetDuration(v int32) {
	o.Duration.Set(&v)
}

// SetDurationNil sets the value for Duration to be an explicit nil
func (o *ProductListing) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil
func (o *ProductListing) UnsetDuration() {
	o.Duration.Unset()
}

// GetDurationType returns the DurationType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductListing) GetDurationType() string {
	if o == nil || IsNil(o.DurationType.Get()) {
		var ret string
		return ret
	}
	return *o.DurationType.Get()
}

// GetDurationTypeOk returns a tuple with the DurationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductListing) GetDurationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DurationType.Get(), o.DurationType.IsSet()
}

// HasDurationType returns a boolean if a field has been set.
func (o *ProductListing) HasDurationType() bool {
	if o != nil && o.DurationType.IsSet() {
		return true
	}

	return false
}

// SetDurationType gets a reference to the given NullableString and assigns it to the DurationType field.
func (o *ProductListing) SetDurationType(v string) {
	o.DurationType.Set(&v)
}

// SetDurationTypeNil sets the value for DurationType to be an explicit nil
func (o *ProductListing) SetDurationTypeNil() {
	o.DurationType.Set(nil)
}

// UnsetDurationType ensures that no value is present for DurationType, not even an explicit nil
func (o *ProductListing) UnsetDurationType() {
	o.DurationType.Unset()
}

// GetGroupIcon returns the GroupIcon field value if set, zero value otherwise.
func (o *ProductListing) GetGroupIcon() string {
	if o == nil || IsNil(o.GroupIcon) {
		var ret string
		return ret
	}
	return *o.GroupIcon
}

// GetGroupIconOk returns a tuple with the GroupIcon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductListing) GetGroupIconOk() (*string, bool) {
	if o == nil || IsNil(o.GroupIcon) {
		return nil, false
	}
	return o.GroupIcon, true
}

// HasGroupIcon returns a boolean if a field has been set.
func (o *ProductListing) HasGroupIcon() bool {
	if o != nil && !IsNil(o.GroupIcon) {
		return true
	}

	return false
}

// SetGroupIcon gets a reference to the given string and assigns it to the GroupIcon field.
func (o *ProductListing) SetGroupIcon(v string) {
	o.GroupIcon = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ProductListing) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductListing) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ProductListing) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ProductListing) SetGroupId(v string) {
	o.GroupId = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductListing) GetGroupName() string {
	if o == nil || IsNil(o.GroupName.Get()) {
		var ret string
		return ret
	}
	return *o.GroupName.Get()
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductListing) GetGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupName.Get(), o.GroupName.IsSet()
}

// HasGroupName returns a boolean if a field has been set.
func (o *ProductListing) HasGroupName() bool {
	if o != nil && o.GroupName.IsSet() {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given NullableString and assigns it to the GroupName field.
func (o *ProductListing) SetGroupName(v string) {
	o.GroupName.Set(&v)
}

// SetGroupNameNil sets the value for GroupName to be an explicit nil
func (o *ProductListing) SetGroupNameNil() {
	o.GroupName.Set(nil)
}

// UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
func (o *ProductListing) UnsetGroupName() {
	o.GroupName.Unset()
}

// GetHasAvatar returns the HasAvatar field value
func (o *ProductListing) GetHasAvatar() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasAvatar
}

// GetHasAvatarOk returns a tuple with the HasAvatar field value
// and a boolean to check if the value has been set.
func (o *ProductListing) GetHasAvatarOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasAvatar, true
}

// SetHasAvatar sets field value
func (o *ProductListing) SetHasAvatar(v bool) {
	o.HasAvatar = v
}

// GetHasUdon returns the HasUdon field value
func (o *ProductListing) GetHasUdon() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasUdon
}

// GetHasUdonOk returns a tuple with the HasUdon field value
// and a boolean to check if the value has been set.
func (o *ProductListing) GetHasUdonOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasUdon, true
}

// SetHasUdon sets field value
func (o *ProductListing) SetHasUdon(v bool) {
	o.HasUdon = v
}

// GetHydratedProducts returns the HydratedProducts field value if set, zero value otherwise.
func (o *ProductListing) GetHydratedProducts() []Product {
	if o == nil || IsNil(o.HydratedProducts) {
		var ret []Product
		return ret
	}
	return o.HydratedProducts
}

// GetHydratedProductsOk returns a tuple with the HydratedProducts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductListing) GetHydratedProductsOk() ([]Product, bool) {
	if o == nil || IsNil(o.HydratedProducts) {
		return nil, false
	}
	return o.HydratedProducts, true
}

// HasHydratedProducts returns a boolean if a field has been set.
func (o *ProductListing) HasHydratedProducts() bool {
	if o != nil && !IsNil(o.HydratedProducts) {
		return true
	}

	return false
}

// SetHydratedProducts gets a reference to the given []Product and assigns it to the HydratedProducts field.
func (o *ProductListing) SetHydratedProducts(v []Product) {
	o.HydratedProducts = v
}

// GetId returns the Id field value
func (o *ProductListing) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProductListing) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProductListing) SetId(v string) {
	o.Id = v
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *ProductListing) GetImageId() string {
	if o == nil || IsNil(o.ImageId) {
		var ret string
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductListing) GetImageIdOk() (*string, bool) {
	if o == nil || IsNil(o.ImageId) {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *ProductListing) HasImageId() bool {
	if o != nil && !IsNil(o.ImageId) {
		return true
	}

	return false
}

// SetImageId gets a reference to the given string and assigns it to the ImageId field.
func (o *ProductListing) SetImageId(v string) {
	o.ImageId = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductListing) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ImageUrl.Get()
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductListing) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageUrl.Get(), o.ImageUrl.IsSet()
}

// HasImageUrl returns a boolean if a field has been set.
func (o *ProductListing) HasImageUrl() bool {
	if o != nil && o.ImageUrl.IsSet() {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given NullableString and assigns it to the ImageUrl field.
func (o *ProductListing) SetImageUrl(v string) {
	o.ImageUrl.Set(&v)
}

// SetImageUrlNil sets the value for ImageUrl to be an explicit nil
func (o *ProductListing) SetImageUrlNil() {
	o.ImageUrl.Set(nil)
}

// UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
func (o *ProductListing) UnsetImageUrl() {
	o.ImageUrl.Unset()
}

// GetListingType returns the ListingType field value
func (o *ProductListing) GetListingType() ProductListingType {
	if o == nil {
		var ret ProductListingType
		return ret
	}

	return o.ListingType
}

// GetListingTypeOk returns a tuple with the ListingType field value
// and a boolean to check if the value has been set.
func (o *ProductListing) GetListingTypeOk() (*ProductListingType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListingType, true
}

// SetListingType sets field value
func (o *ProductListing) SetListingType(v ProductListingType) {
	o.ListingType = v
}

// GetListingVariants returns the ListingVariants field value if set, zero value otherwise.
func (o *ProductListing) GetListingVariants() []ProductListingVariant {
	if o == nil || IsNil(o.ListingVariants) {
		var ret []ProductListingVariant
		return ret
	}
	return o.ListingVariants
}

// GetListingVariantsOk returns a tuple with the ListingVariants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductListing) GetListingVariantsOk() ([]ProductListingVariant, bool) {
	if o == nil || IsNil(o.ListingVariants) {
		return nil, false
	}
	return o.ListingVariants, true
}

// HasListingVariants returns a boolean if a field has been set.
func (o *ProductListing) HasListingVariants() bool {
	if o != nil && !IsNil(o.ListingVariants) {
		return true
	}

	return false
}

// SetListingVariants gets a reference to the given []ProductListingVariant and assigns it to the ListingVariants field.
func (o *ProductListing) SetListingVariants(v []ProductListingVariant) {
	o.ListingVariants = v
}

// GetPermanent returns the Permanent field value if set, zero value otherwise.
func (o *ProductListing) GetPermanent() bool {
	if o == nil || IsNil(o.Permanent) {
		var ret bool
		return ret
	}
	return *o.Permanent
}

// GetPermanentOk returns a tuple with the Permanent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductListing) GetPermanentOk() (*bool, bool) {
	if o == nil || IsNil(o.Permanent) {
		return nil, false
	}
	return o.Permanent, true
}

// HasPermanent returns a boolean if a field has been set.
func (o *ProductListing) HasPermanent() bool {
	if o != nil && !IsNil(o.Permanent) {
		return true
	}

	return false
}

// SetPermanent gets a reference to the given bool and assigns it to the Permanent field.
func (o *ProductListing) SetPermanent(v bool) {
	o.Permanent = &v
}

// GetPriceTokens returns the PriceTokens field value
func (o *ProductListing) GetPriceTokens() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PriceTokens
}

// GetPriceTokensOk returns a tuple with the PriceTokens field value
// and a boolean to check if the value has been set.
func (o *ProductListing) GetPriceTokensOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceTokens, true
}

// SetPriceTokens sets field value
func (o *ProductListing) SetPriceTokens(v int32) {
	o.PriceTokens = v
}

// GetProductIds returns the ProductIds field value
func (o *ProductListing) GetProductIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ProductIds
}

// GetProductIdsOk returns a tuple with the ProductIds field value
// and a boolean to check if the value has been set.
func (o *ProductListing) GetProductIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductIds, true
}

// SetProductIds sets field value
func (o *ProductListing) SetProductIds(v []string) {
	o.ProductIds = v
}

// GetProductType returns the ProductType field value
func (o *ProductListing) GetProductType() ProductType {
	if o == nil {
		var ret ProductType
		return ret
	}

	return o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value
// and a boolean to check if the value has been set.
func (o *ProductListing) GetProductTypeOk() (*ProductType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductType, true
}

// SetProductType sets field value
func (o *ProductListing) SetProductType(v ProductType) {
	o.ProductType = v
}

// GetProducts returns the Products field value
func (o *ProductListing) GetProducts() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Products
}

// GetProductsOk returns a tuple with the Products field value
// and a boolean to check if the value has been set.
func (o *ProductListing) GetProductsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Products, true
}

// SetProducts sets field value
func (o *ProductListing) SetProducts(v []map[string]interface{}) {
	o.Products = v
}

// GetQuantifiable returns the Quantifiable field value if set, zero value otherwise.
func (o *ProductListing) GetQuantifiable() bool {
	if o == nil || IsNil(o.Quantifiable) {
		var ret bool
		return ret
	}
	return *o.Quantifiable
}

// GetQuantifiableOk returns a tuple with the Quantifiable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductListing) GetQuantifiableOk() (*bool, bool) {
	if o == nil || IsNil(o.Quantifiable) {
		return nil, false
	}
	return o.Quantifiable, true
}

// HasQuantifiable returns a boolean if a field has been set.
func (o *ProductListing) HasQuantifiable() bool {
	if o != nil && !IsNil(o.Quantifiable) {
		return true
	}

	return false
}

// SetQuantifiable gets a reference to the given bool and assigns it to the Quantifiable field.
func (o *ProductListing) SetQuantifiable(v bool) {
	o.Quantifiable = &v
}

// GetRecurrable returns the Recurrable field value
func (o *ProductListing) GetRecurrable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Recurrable
}

// GetRecurrableOk returns a tuple with the Recurrable field value
// and a boolean to check if the value has been set.
func (o *ProductListing) GetRecurrableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recurrable, true
}

// SetRecurrable sets field value
func (o *ProductListing) SetRecurrable(v bool) {
	o.Recurrable = v
}

// GetRefundable returns the Refundable field value
func (o *ProductListing) GetRefundable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Refundable
}

// GetRefundableOk returns a tuple with the Refundable field value
// and a boolean to check if the value has been set.
func (o *ProductListing) GetRefundableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Refundable, true
}

// SetRefundable sets field value
func (o *ProductListing) SetRefundable(v bool) {
	o.Refundable = v
}

// GetSellerDisplayName returns the SellerDisplayName field value
func (o *ProductListing) GetSellerDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SellerDisplayName
}

// GetSellerDisplayNameOk returns a tuple with the SellerDisplayName field value
// and a boolean to check if the value has been set.
func (o *ProductListing) GetSellerDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellerDisplayName, true
}

// SetSellerDisplayName sets field value
func (o *ProductListing) SetSellerDisplayName(v string) {
	o.SellerDisplayName = v
}

// GetSellerId returns the SellerId field value
func (o *ProductListing) GetSellerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SellerId
}

// GetSellerIdOk returns a tuple with the SellerId field value
// and a boolean to check if the value has been set.
func (o *ProductListing) GetSellerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellerId, true
}

// SetSellerId sets field value
func (o *ProductListing) SetSellerId(v string) {
	o.SellerId = v
}

// GetSoldByVrc returns the SoldByVrc field value if set, zero value otherwise.
func (o *ProductListing) GetSoldByVrc() bool {
	if o == nil || IsNil(o.SoldByVrc) {
		var ret bool
		return ret
	}
	return *o.SoldByVrc
}

// GetSoldByVrcOk returns a tuple with the SoldByVrc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductListing) GetSoldByVrcOk() (*bool, bool) {
	if o == nil || IsNil(o.SoldByVrc) {
		return nil, false
	}
	return o.SoldByVrc, true
}

// HasSoldByVrc returns a boolean if a field has been set.
func (o *ProductListing) HasSoldByVrc() bool {
	if o != nil && !IsNil(o.SoldByVrc) {
		return true
	}

	return false
}

// SetSoldByVrc gets a reference to the given bool and assigns it to the SoldByVrc field.
func (o *ProductListing) SetSoldByVrc(v bool) {
	o.SoldByVrc = &v
}

// GetStackable returns the Stackable field value
func (o *ProductListing) GetStackable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Stackable
}

// GetStackableOk returns a tuple with the Stackable field value
// and a boolean to check if the value has been set.
func (o *ProductListing) GetStackableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stackable, true
}

// SetStackable sets field value
func (o *ProductListing) SetStackable(v bool) {
	o.Stackable = v
}

// GetStoreIds returns the StoreIds field value
func (o *ProductListing) GetStoreIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.StoreIds
}

// GetStoreIdsOk returns a tuple with the StoreIds field value
// and a boolean to check if the value has been set.
func (o *ProductListing) GetStoreIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StoreIds, true
}

// SetStoreIds sets field value
func (o *ProductListing) SetStoreIds(v []string) {
	o.StoreIds = v
}

// GetSubtitle returns the Subtitle field value if set, zero value otherwise.
func (o *ProductListing) GetSubtitle() string {
	if o == nil || IsNil(o.Subtitle) {
		var ret string
		return ret
	}
	return *o.Subtitle
}

// GetSubtitleOk returns a tuple with the Subtitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductListing) GetSubtitleOk() (*string, bool) {
	if o == nil || IsNil(o.Subtitle) {
		return nil, false
	}
	return o.Subtitle, true
}

// HasSubtitle returns a boolean if a field has been set.
func (o *ProductListing) HasSubtitle() bool {
	if o != nil && !IsNil(o.Subtitle) {
		return true
	}

	return false
}

// SetSubtitle gets a reference to the given string and assigns it to the Subtitle field.
func (o *ProductListing) SetSubtitle(v string) {
	o.Subtitle = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ProductListing) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductListing) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ProductListing) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ProductListing) SetTags(v []string) {
	o.Tags = v
}

// GetVrcPlusDiscountPrice returns the VrcPlusDiscountPrice field value if set, zero value otherwise.
func (o *ProductListing) GetVrcPlusDiscountPrice() int32 {
	if o == nil || IsNil(o.VrcPlusDiscountPrice) {
		var ret int32
		return ret
	}
	return *o.VrcPlusDiscountPrice
}

// GetVrcPlusDiscountPriceOk returns a tuple with the VrcPlusDiscountPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductListing) GetVrcPlusDiscountPriceOk() (*int32, bool) {
	if o == nil || IsNil(o.VrcPlusDiscountPrice) {
		return nil, false
	}
	return o.VrcPlusDiscountPrice, true
}

// HasVrcPlusDiscountPrice returns a boolean if a field has been set.
func (o *ProductListing) HasVrcPlusDiscountPrice() bool {
	if o != nil && !IsNil(o.VrcPlusDiscountPrice) {
		return true
	}

	return false
}

// SetVrcPlusDiscountPrice gets a reference to the given int32 and assigns it to the VrcPlusDiscountPrice field.
func (o *ProductListing) SetVrcPlusDiscountPrice(v int32) {
	o.VrcPlusDiscountPrice = &v
}

// GetWhenToExpire returns the WhenToExpire field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductListing) GetWhenToExpire() time.Time {
	if o == nil || IsNil(o.WhenToExpire.Get()) {
		var ret time.Time
		return ret
	}
	return *o.WhenToExpire.Get()
}

// GetWhenToExpireOk returns a tuple with the WhenToExpire field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductListing) GetWhenToExpireOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.WhenToExpire.Get(), o.WhenToExpire.IsSet()
}

// HasWhenToExpire returns a boolean if a field has been set.
func (o *ProductListing) HasWhenToExpire() bool {
	if o != nil && o.WhenToExpire.IsSet() {
		return true
	}

	return false
}

// SetWhenToExpire gets a reference to the given NullableTime and assigns it to the WhenToExpire field.
func (o *ProductListing) SetWhenToExpire(v time.Time) {
	o.WhenToExpire.Set(&v)
}

// SetWhenToExpireNil sets the value for WhenToExpire to be an explicit nil
func (o *ProductListing) SetWhenToExpireNil() {
	o.WhenToExpire.Set(nil)
}

// UnsetWhenToExpire ensures that no value is present for WhenToExpire, not even an explicit nil
func (o *ProductListing) UnsetWhenToExpire() {
	o.WhenToExpire.Unset()
}

func (o ProductListing) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductListing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active"] = o.Active
	toSerialize["buyerRefundable"] = o.BuyerRefundable
	toSerialize["description"] = o.Description
	toSerialize["displayName"] = o.DisplayName
	if o.Duration.IsSet() {
		toSerialize["duration"] = o.Duration.Get()
	}
	if o.DurationType.IsSet() {
		toSerialize["durationType"] = o.DurationType.Get()
	}
	if !IsNil(o.GroupIcon) {
		toSerialize["groupIcon"] = o.GroupIcon
	}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if o.GroupName.IsSet() {
		toSerialize["groupName"] = o.GroupName.Get()
	}
	toSerialize["hasAvatar"] = o.HasAvatar
	toSerialize["hasUdon"] = o.HasUdon
	if !IsNil(o.HydratedProducts) {
		toSerialize["hydratedProducts"] = o.HydratedProducts
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.ImageId) {
		toSerialize["imageId"] = o.ImageId
	}
	if o.ImageUrl.IsSet() {
		toSerialize["imageUrl"] = o.ImageUrl.Get()
	}
	toSerialize["listingType"] = o.ListingType
	if !IsNil(o.ListingVariants) {
		toSerialize["listingVariants"] = o.ListingVariants
	}
	if !IsNil(o.Permanent) {
		toSerialize["permanent"] = o.Permanent
	}
	toSerialize["priceTokens"] = o.PriceTokens
	toSerialize["productIds"] = o.ProductIds
	toSerialize["productType"] = o.ProductType
	toSerialize["products"] = o.Products
	if !IsNil(o.Quantifiable) {
		toSerialize["quantifiable"] = o.Quantifiable
	}
	toSerialize["recurrable"] = o.Recurrable
	toSerialize["refundable"] = o.Refundable
	toSerialize["sellerDisplayName"] = o.SellerDisplayName
	toSerialize["sellerId"] = o.SellerId
	if !IsNil(o.SoldByVrc) {
		toSerialize["soldByVrc"] = o.SoldByVrc
	}
	toSerialize["stackable"] = o.Stackable
	toSerialize["storeIds"] = o.StoreIds
	if !IsNil(o.Subtitle) {
		toSerialize["subtitle"] = o.Subtitle
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.VrcPlusDiscountPrice) {
		toSerialize["vrcPlusDiscountPrice"] = o.VrcPlusDiscountPrice
	}
	if o.WhenToExpire.IsSet() {
		toSerialize["whenToExpire"] = o.WhenToExpire.Get()
	}
	return toSerialize, nil
}

func (o *ProductListing) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"active",
		"buyerRefundable",
		"description",
		"displayName",
		"hasAvatar",
		"hasUdon",
		"id",
		"listingType",
		"priceTokens",
		"productIds",
		"productType",
		"products",
		"recurrable",
		"refundable",
		"sellerDisplayName",
		"sellerId",
		"stackable",
		"storeIds",
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

	varProductListing := _ProductListing{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProductListing)

	if err != nil {
		return err
	}

	*o = ProductListing(varProductListing)

	return err
}

type NullableProductListing struct {
	value *ProductListing
	isSet bool
}

func (v NullableProductListing) Get() *ProductListing {
	return v.value
}

func (v *NullableProductListing) Set(val *ProductListing) {
	v.value = val
	v.isSet = true
}

func (v NullableProductListing) IsSet() bool {
	return v.isSet
}

func (v *NullableProductListing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductListing(val *ProductListing) *NullableProductListing {
	return &NullableProductListing{value: val, isSet: true}
}

func (v NullableProductListing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductListing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
