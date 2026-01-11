# Store

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**DisplayName** | **string** |  | 
**GroupId** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**ListingIds** | Pointer to **[]string** | Only for store type world and group | [optional] 
**Listings** | Pointer to [**[]ProductListing**](ProductListing.md) | Only for store type world and group | [optional] 
**SellerDisplayName** | **string** |  | 
**SellerId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**ShelfIds** | Pointer to **[]string** | Only for store type house | [optional] 
**Shelves** | Pointer to [**[]StoreShelf**](StoreShelf.md) | Only for store type house | [optional] 
**StoreId** | **string** |  | 
**StoreType** | [**StoreType**](StoreType.md) |  | [default to GROUP]
**Tags** | **[]string** |  | 
**WorldId** | Pointer to **string** | WorldID be \&quot;offline\&quot; on User profiles if you are not friends with that user. | [optional] 

## Methods

### NewStore

`func NewStore(description string, displayName string, id string, sellerDisplayName string, sellerId string, storeId string, storeType StoreType, tags []string, ) *Store`

NewStore instantiates a new Store object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreWithDefaults

`func NewStoreWithDefaults() *Store`

NewStoreWithDefaults instantiates a new Store object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Store) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Store) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Store) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDisplayName

`func (o *Store) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Store) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Store) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetGroupId

`func (o *Store) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Store) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Store) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *Store) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *Store) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Store) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Store) SetId(v string)`

SetId sets Id field to given value.


### GetListingIds

`func (o *Store) GetListingIds() []string`

GetListingIds returns the ListingIds field if non-nil, zero value otherwise.

### GetListingIdsOk

`func (o *Store) GetListingIdsOk() (*[]string, bool)`

GetListingIdsOk returns a tuple with the ListingIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingIds

`func (o *Store) SetListingIds(v []string)`

SetListingIds sets ListingIds field to given value.

### HasListingIds

`func (o *Store) HasListingIds() bool`

HasListingIds returns a boolean if a field has been set.

### GetListings

`func (o *Store) GetListings() []ProductListing`

GetListings returns the Listings field if non-nil, zero value otherwise.

### GetListingsOk

`func (o *Store) GetListingsOk() (*[]ProductListing, bool)`

GetListingsOk returns a tuple with the Listings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListings

`func (o *Store) SetListings(v []ProductListing)`

SetListings sets Listings field to given value.

### HasListings

`func (o *Store) HasListings() bool`

HasListings returns a boolean if a field has been set.

### GetSellerDisplayName

`func (o *Store) GetSellerDisplayName() string`

GetSellerDisplayName returns the SellerDisplayName field if non-nil, zero value otherwise.

### GetSellerDisplayNameOk

`func (o *Store) GetSellerDisplayNameOk() (*string, bool)`

GetSellerDisplayNameOk returns a tuple with the SellerDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerDisplayName

`func (o *Store) SetSellerDisplayName(v string)`

SetSellerDisplayName sets SellerDisplayName field to given value.


### GetSellerId

`func (o *Store) GetSellerId() string`

GetSellerId returns the SellerId field if non-nil, zero value otherwise.

### GetSellerIdOk

`func (o *Store) GetSellerIdOk() (*string, bool)`

GetSellerIdOk returns a tuple with the SellerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerId

`func (o *Store) SetSellerId(v string)`

SetSellerId sets SellerId field to given value.


### GetShelfIds

`func (o *Store) GetShelfIds() []string`

GetShelfIds returns the ShelfIds field if non-nil, zero value otherwise.

### GetShelfIdsOk

`func (o *Store) GetShelfIdsOk() (*[]string, bool)`

GetShelfIdsOk returns a tuple with the ShelfIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShelfIds

`func (o *Store) SetShelfIds(v []string)`

SetShelfIds sets ShelfIds field to given value.

### HasShelfIds

`func (o *Store) HasShelfIds() bool`

HasShelfIds returns a boolean if a field has been set.

### GetShelves

`func (o *Store) GetShelves() []StoreShelf`

GetShelves returns the Shelves field if non-nil, zero value otherwise.

### GetShelvesOk

`func (o *Store) GetShelvesOk() (*[]StoreShelf, bool)`

GetShelvesOk returns a tuple with the Shelves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShelves

`func (o *Store) SetShelves(v []StoreShelf)`

SetShelves sets Shelves field to given value.

### HasShelves

`func (o *Store) HasShelves() bool`

HasShelves returns a boolean if a field has been set.

### GetStoreId

`func (o *Store) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *Store) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *Store) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.


### GetStoreType

`func (o *Store) GetStoreType() StoreType`

GetStoreType returns the StoreType field if non-nil, zero value otherwise.

### GetStoreTypeOk

`func (o *Store) GetStoreTypeOk() (*StoreType, bool)`

GetStoreTypeOk returns a tuple with the StoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreType

`func (o *Store) SetStoreType(v StoreType)`

SetStoreType sets StoreType field to given value.


### GetTags

`func (o *Store) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Store) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Store) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetWorldId

`func (o *Store) GetWorldId() string`

GetWorldId returns the WorldId field if non-nil, zero value otherwise.

### GetWorldIdOk

`func (o *Store) GetWorldIdOk() (*string, bool)`

GetWorldIdOk returns a tuple with the WorldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorldId

`func (o *Store) SetWorldId(v string)`

SetWorldId sets WorldId field to given value.

### HasWorldId

`func (o *Store) HasWorldId() bool`

HasWorldId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


