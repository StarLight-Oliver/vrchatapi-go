# Product

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | Pointer to **bool** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Description** | **string** |  | 
**DisplayName** | **string** |  | 
**GroupAccess** | Pointer to **bool** |  | [optional] [default to false]
**GroupAccessRemove** | Pointer to **bool** |  | [optional] [default to false]
**GroupId** | Pointer to **string** |  | [optional] 
**GroupRoleId** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**ImageId** | **string** |  | 
**ParentListings** | **[]string** |  | 
**ProductType** | [**ProductType**](ProductType.md) |  | [default to UDON]
**SellerDisplayName** | **string** |  | 
**SellerId** | **string** |  | 
**Tags** | **[]string** |  | 
**Updated** | Pointer to **NullableTime** |  | [optional] 
**UseForSubscriberList** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewProduct

`func NewProduct(description string, displayName string, id string, imageId string, parentListings []string, productType ProductType, sellerDisplayName string, sellerId string, tags []string, ) *Product`

NewProduct instantiates a new Product object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductWithDefaults

`func NewProductWithDefaults() *Product`

NewProductWithDefaults instantiates a new Product object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *Product) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *Product) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *Product) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *Product) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetCreated

`func (o *Product) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Product) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Product) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Product) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *Product) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Product) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Product) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDisplayName

`func (o *Product) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Product) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Product) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetGroupAccess

`func (o *Product) GetGroupAccess() bool`

GetGroupAccess returns the GroupAccess field if non-nil, zero value otherwise.

### GetGroupAccessOk

`func (o *Product) GetGroupAccessOk() (*bool, bool)`

GetGroupAccessOk returns a tuple with the GroupAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAccess

`func (o *Product) SetGroupAccess(v bool)`

SetGroupAccess sets GroupAccess field to given value.

### HasGroupAccess

`func (o *Product) HasGroupAccess() bool`

HasGroupAccess returns a boolean if a field has been set.

### GetGroupAccessRemove

`func (o *Product) GetGroupAccessRemove() bool`

GetGroupAccessRemove returns the GroupAccessRemove field if non-nil, zero value otherwise.

### GetGroupAccessRemoveOk

`func (o *Product) GetGroupAccessRemoveOk() (*bool, bool)`

GetGroupAccessRemoveOk returns a tuple with the GroupAccessRemove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAccessRemove

`func (o *Product) SetGroupAccessRemove(v bool)`

SetGroupAccessRemove sets GroupAccessRemove field to given value.

### HasGroupAccessRemove

`func (o *Product) HasGroupAccessRemove() bool`

HasGroupAccessRemove returns a boolean if a field has been set.

### GetGroupId

`func (o *Product) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Product) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Product) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *Product) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupRoleId

`func (o *Product) GetGroupRoleId() string`

GetGroupRoleId returns the GroupRoleId field if non-nil, zero value otherwise.

### GetGroupRoleIdOk

`func (o *Product) GetGroupRoleIdOk() (*string, bool)`

GetGroupRoleIdOk returns a tuple with the GroupRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupRoleId

`func (o *Product) SetGroupRoleId(v string)`

SetGroupRoleId sets GroupRoleId field to given value.

### HasGroupRoleId

`func (o *Product) HasGroupRoleId() bool`

HasGroupRoleId returns a boolean if a field has been set.

### GetId

`func (o *Product) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Product) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Product) SetId(v string)`

SetId sets Id field to given value.


### GetImageId

`func (o *Product) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *Product) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *Product) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetParentListings

`func (o *Product) GetParentListings() []string`

GetParentListings returns the ParentListings field if non-nil, zero value otherwise.

### GetParentListingsOk

`func (o *Product) GetParentListingsOk() (*[]string, bool)`

GetParentListingsOk returns a tuple with the ParentListings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentListings

`func (o *Product) SetParentListings(v []string)`

SetParentListings sets ParentListings field to given value.


### GetProductType

`func (o *Product) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *Product) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *Product) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.


### GetSellerDisplayName

`func (o *Product) GetSellerDisplayName() string`

GetSellerDisplayName returns the SellerDisplayName field if non-nil, zero value otherwise.

### GetSellerDisplayNameOk

`func (o *Product) GetSellerDisplayNameOk() (*string, bool)`

GetSellerDisplayNameOk returns a tuple with the SellerDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerDisplayName

`func (o *Product) SetSellerDisplayName(v string)`

SetSellerDisplayName sets SellerDisplayName field to given value.


### GetSellerId

`func (o *Product) GetSellerId() string`

GetSellerId returns the SellerId field if non-nil, zero value otherwise.

### GetSellerIdOk

`func (o *Product) GetSellerIdOk() (*string, bool)`

GetSellerIdOk returns a tuple with the SellerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerId

`func (o *Product) SetSellerId(v string)`

SetSellerId sets SellerId field to given value.


### GetTags

`func (o *Product) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Product) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Product) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetUpdated

`func (o *Product) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Product) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Product) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Product) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdatedNil

`func (o *Product) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *Product) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil
### GetUseForSubscriberList

`func (o *Product) GetUseForSubscriberList() bool`

GetUseForSubscriberList returns the UseForSubscriberList field if non-nil, zero value otherwise.

### GetUseForSubscriberListOk

`func (o *Product) GetUseForSubscriberListOk() (*bool, bool)`

GetUseForSubscriberListOk returns a tuple with the UseForSubscriberList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForSubscriberList

`func (o *Product) SetUseForSubscriberList(v bool)`

SetUseForSubscriberList sets UseForSubscriberList field to given value.

### HasUseForSubscriberList

`func (o *Product) HasUseForSubscriberList() bool`

HasUseForSubscriberList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


