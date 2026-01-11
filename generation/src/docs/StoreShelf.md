# StoreShelf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HighlightListing** | Pointer to [**ProductListing**](ProductListing.md) |  | [optional] 
**HighlightListingId** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**ListingIds** | **[]string** |  | 
**Listings** | Pointer to [**[]ProductListing**](ProductListing.md) |  | [optional] 
**ShelfDescription** | **string** |  | 
**ShelfLayout** | **string** |  | 
**ShelfTitle** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewStoreShelf

`func NewStoreShelf(id string, listingIds []string, shelfDescription string, shelfLayout string, shelfTitle string, updatedAt time.Time, ) *StoreShelf`

NewStoreShelf instantiates a new StoreShelf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreShelfWithDefaults

`func NewStoreShelfWithDefaults() *StoreShelf`

NewStoreShelfWithDefaults instantiates a new StoreShelf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHighlightListing

`func (o *StoreShelf) GetHighlightListing() ProductListing`

GetHighlightListing returns the HighlightListing field if non-nil, zero value otherwise.

### GetHighlightListingOk

`func (o *StoreShelf) GetHighlightListingOk() (*ProductListing, bool)`

GetHighlightListingOk returns a tuple with the HighlightListing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightListing

`func (o *StoreShelf) SetHighlightListing(v ProductListing)`

SetHighlightListing sets HighlightListing field to given value.

### HasHighlightListing

`func (o *StoreShelf) HasHighlightListing() bool`

HasHighlightListing returns a boolean if a field has been set.

### GetHighlightListingId

`func (o *StoreShelf) GetHighlightListingId() string`

GetHighlightListingId returns the HighlightListingId field if non-nil, zero value otherwise.

### GetHighlightListingIdOk

`func (o *StoreShelf) GetHighlightListingIdOk() (*string, bool)`

GetHighlightListingIdOk returns a tuple with the HighlightListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightListingId

`func (o *StoreShelf) SetHighlightListingId(v string)`

SetHighlightListingId sets HighlightListingId field to given value.

### HasHighlightListingId

`func (o *StoreShelf) HasHighlightListingId() bool`

HasHighlightListingId returns a boolean if a field has been set.

### GetId

`func (o *StoreShelf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StoreShelf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StoreShelf) SetId(v string)`

SetId sets Id field to given value.


### GetListingIds

`func (o *StoreShelf) GetListingIds() []string`

GetListingIds returns the ListingIds field if non-nil, zero value otherwise.

### GetListingIdsOk

`func (o *StoreShelf) GetListingIdsOk() (*[]string, bool)`

GetListingIdsOk returns a tuple with the ListingIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingIds

`func (o *StoreShelf) SetListingIds(v []string)`

SetListingIds sets ListingIds field to given value.


### GetListings

`func (o *StoreShelf) GetListings() []ProductListing`

GetListings returns the Listings field if non-nil, zero value otherwise.

### GetListingsOk

`func (o *StoreShelf) GetListingsOk() (*[]ProductListing, bool)`

GetListingsOk returns a tuple with the Listings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListings

`func (o *StoreShelf) SetListings(v []ProductListing)`

SetListings sets Listings field to given value.

### HasListings

`func (o *StoreShelf) HasListings() bool`

HasListings returns a boolean if a field has been set.

### GetShelfDescription

`func (o *StoreShelf) GetShelfDescription() string`

GetShelfDescription returns the ShelfDescription field if non-nil, zero value otherwise.

### GetShelfDescriptionOk

`func (o *StoreShelf) GetShelfDescriptionOk() (*string, bool)`

GetShelfDescriptionOk returns a tuple with the ShelfDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShelfDescription

`func (o *StoreShelf) SetShelfDescription(v string)`

SetShelfDescription sets ShelfDescription field to given value.


### GetShelfLayout

`func (o *StoreShelf) GetShelfLayout() string`

GetShelfLayout returns the ShelfLayout field if non-nil, zero value otherwise.

### GetShelfLayoutOk

`func (o *StoreShelf) GetShelfLayoutOk() (*string, bool)`

GetShelfLayoutOk returns a tuple with the ShelfLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShelfLayout

`func (o *StoreShelf) SetShelfLayout(v string)`

SetShelfLayout sets ShelfLayout field to given value.


### GetShelfTitle

`func (o *StoreShelf) GetShelfTitle() string`

GetShelfTitle returns the ShelfTitle field if non-nil, zero value otherwise.

### GetShelfTitleOk

`func (o *StoreShelf) GetShelfTitleOk() (*string, bool)`

GetShelfTitleOk returns a tuple with the ShelfTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShelfTitle

`func (o *StoreShelf) SetShelfTitle(v string)`

SetShelfTitle sets ShelfTitle field to given value.


### GetUpdatedAt

`func (o *StoreShelf) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *StoreShelf) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *StoreShelf) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


