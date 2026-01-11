# InventoryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | **[]string** |  | 
**CreatedAt** | **time.Time** |  | 
**DefaultAttributes** | [**map[string]InventoryDefaultAttributesValue**](InventoryDefaultAttributesValue.md) |  | 
**Description** | **string** |  | 
**EquipSlot** | Pointer to [**InventoryEquipSlot**](InventoryEquipSlot.md) |  | [optional] [default to EMPTY]
**EquipSlots** | Pointer to [**[]InventoryEquipSlot**](InventoryEquipSlot.md) |  | [optional] 
**ExpiryDate** | Pointer to **NullableTime** |  | [optional] 
**Flags** | **[]string** |  | 
**HolderId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**Id** | **string** |  | 
**ImageUrl** | **string** |  | 
**IsArchived** | **bool** |  | 
**IsSeen** | **bool** |  | 
**ItemType** | [**InventoryItemType**](InventoryItemType.md) |  | [default to BUNDLE]
**ItemTypeLabel** | **string** |  | 
**Metadata** | [**InventoryMetadata**](InventoryMetadata.md) |  | 
**Name** | **string** |  | 
**Quantifiable** | **bool** |  | 
**Tags** | **[]string** |  | 
**TemplateId** | **string** |  | 
**TemplateCreatedAt** | **time.Time** |  | 
**TemplateUpdatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**UserAttributes** | [**InventoryUserAttributes**](InventoryUserAttributes.md) |  | 
**ValidateUserAttributes** | **bool** |  | 

## Methods

### NewInventoryItem

`func NewInventoryItem(collections []string, createdAt time.Time, defaultAttributes map[string]InventoryDefaultAttributesValue, description string, flags []string, holderId string, id string, imageUrl string, isArchived bool, isSeen bool, itemType InventoryItemType, itemTypeLabel string, metadata InventoryMetadata, name string, quantifiable bool, tags []string, templateId string, templateCreatedAt time.Time, templateUpdatedAt time.Time, updatedAt time.Time, userAttributes InventoryUserAttributes, validateUserAttributes bool, ) *InventoryItem`

NewInventoryItem instantiates a new InventoryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryItemWithDefaults

`func NewInventoryItemWithDefaults() *InventoryItem`

NewInventoryItemWithDefaults instantiates a new InventoryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *InventoryItem) GetCollections() []string`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *InventoryItem) GetCollectionsOk() (*[]string, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *InventoryItem) SetCollections(v []string)`

SetCollections sets Collections field to given value.


### GetCreatedAt

`func (o *InventoryItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InventoryItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InventoryItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDefaultAttributes

`func (o *InventoryItem) GetDefaultAttributes() map[string]InventoryDefaultAttributesValue`

GetDefaultAttributes returns the DefaultAttributes field if non-nil, zero value otherwise.

### GetDefaultAttributesOk

`func (o *InventoryItem) GetDefaultAttributesOk() (*map[string]InventoryDefaultAttributesValue, bool)`

GetDefaultAttributesOk returns a tuple with the DefaultAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAttributes

`func (o *InventoryItem) SetDefaultAttributes(v map[string]InventoryDefaultAttributesValue)`

SetDefaultAttributes sets DefaultAttributes field to given value.


### GetDescription

`func (o *InventoryItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InventoryItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InventoryItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEquipSlot

`func (o *InventoryItem) GetEquipSlot() InventoryEquipSlot`

GetEquipSlot returns the EquipSlot field if non-nil, zero value otherwise.

### GetEquipSlotOk

`func (o *InventoryItem) GetEquipSlotOk() (*InventoryEquipSlot, bool)`

GetEquipSlotOk returns a tuple with the EquipSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipSlot

`func (o *InventoryItem) SetEquipSlot(v InventoryEquipSlot)`

SetEquipSlot sets EquipSlot field to given value.

### HasEquipSlot

`func (o *InventoryItem) HasEquipSlot() bool`

HasEquipSlot returns a boolean if a field has been set.

### GetEquipSlots

`func (o *InventoryItem) GetEquipSlots() []InventoryEquipSlot`

GetEquipSlots returns the EquipSlots field if non-nil, zero value otherwise.

### GetEquipSlotsOk

`func (o *InventoryItem) GetEquipSlotsOk() (*[]InventoryEquipSlot, bool)`

GetEquipSlotsOk returns a tuple with the EquipSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipSlots

`func (o *InventoryItem) SetEquipSlots(v []InventoryEquipSlot)`

SetEquipSlots sets EquipSlots field to given value.

### HasEquipSlots

`func (o *InventoryItem) HasEquipSlots() bool`

HasEquipSlots returns a boolean if a field has been set.

### GetExpiryDate

`func (o *InventoryItem) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *InventoryItem) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *InventoryItem) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *InventoryItem) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDateNil

`func (o *InventoryItem) SetExpiryDateNil(b bool)`

 SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil

### UnsetExpiryDate
`func (o *InventoryItem) UnsetExpiryDate()`

UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
### GetFlags

`func (o *InventoryItem) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *InventoryItem) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *InventoryItem) SetFlags(v []string)`

SetFlags sets Flags field to given value.


### GetHolderId

`func (o *InventoryItem) GetHolderId() string`

GetHolderId returns the HolderId field if non-nil, zero value otherwise.

### GetHolderIdOk

`func (o *InventoryItem) GetHolderIdOk() (*string, bool)`

GetHolderIdOk returns a tuple with the HolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderId

`func (o *InventoryItem) SetHolderId(v string)`

SetHolderId sets HolderId field to given value.


### GetId

`func (o *InventoryItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InventoryItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InventoryItem) SetId(v string)`

SetId sets Id field to given value.


### GetImageUrl

`func (o *InventoryItem) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *InventoryItem) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *InventoryItem) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.


### GetIsArchived

`func (o *InventoryItem) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *InventoryItem) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *InventoryItem) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.


### GetIsSeen

`func (o *InventoryItem) GetIsSeen() bool`

GetIsSeen returns the IsSeen field if non-nil, zero value otherwise.

### GetIsSeenOk

`func (o *InventoryItem) GetIsSeenOk() (*bool, bool)`

GetIsSeenOk returns a tuple with the IsSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSeen

`func (o *InventoryItem) SetIsSeen(v bool)`

SetIsSeen sets IsSeen field to given value.


### GetItemType

`func (o *InventoryItem) GetItemType() InventoryItemType`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *InventoryItem) GetItemTypeOk() (*InventoryItemType, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *InventoryItem) SetItemType(v InventoryItemType)`

SetItemType sets ItemType field to given value.


### GetItemTypeLabel

`func (o *InventoryItem) GetItemTypeLabel() string`

GetItemTypeLabel returns the ItemTypeLabel field if non-nil, zero value otherwise.

### GetItemTypeLabelOk

`func (o *InventoryItem) GetItemTypeLabelOk() (*string, bool)`

GetItemTypeLabelOk returns a tuple with the ItemTypeLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTypeLabel

`func (o *InventoryItem) SetItemTypeLabel(v string)`

SetItemTypeLabel sets ItemTypeLabel field to given value.


### GetMetadata

`func (o *InventoryItem) GetMetadata() InventoryMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InventoryItem) GetMetadataOk() (*InventoryMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InventoryItem) SetMetadata(v InventoryMetadata)`

SetMetadata sets Metadata field to given value.


### GetName

`func (o *InventoryItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InventoryItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InventoryItem) SetName(v string)`

SetName sets Name field to given value.


### GetQuantifiable

`func (o *InventoryItem) GetQuantifiable() bool`

GetQuantifiable returns the Quantifiable field if non-nil, zero value otherwise.

### GetQuantifiableOk

`func (o *InventoryItem) GetQuantifiableOk() (*bool, bool)`

GetQuantifiableOk returns a tuple with the Quantifiable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantifiable

`func (o *InventoryItem) SetQuantifiable(v bool)`

SetQuantifiable sets Quantifiable field to given value.


### GetTags

`func (o *InventoryItem) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InventoryItem) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InventoryItem) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetTemplateId

`func (o *InventoryItem) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *InventoryItem) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *InventoryItem) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.


### GetTemplateCreatedAt

`func (o *InventoryItem) GetTemplateCreatedAt() time.Time`

GetTemplateCreatedAt returns the TemplateCreatedAt field if non-nil, zero value otherwise.

### GetTemplateCreatedAtOk

`func (o *InventoryItem) GetTemplateCreatedAtOk() (*time.Time, bool)`

GetTemplateCreatedAtOk returns a tuple with the TemplateCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateCreatedAt

`func (o *InventoryItem) SetTemplateCreatedAt(v time.Time)`

SetTemplateCreatedAt sets TemplateCreatedAt field to given value.


### GetTemplateUpdatedAt

`func (o *InventoryItem) GetTemplateUpdatedAt() time.Time`

GetTemplateUpdatedAt returns the TemplateUpdatedAt field if non-nil, zero value otherwise.

### GetTemplateUpdatedAtOk

`func (o *InventoryItem) GetTemplateUpdatedAtOk() (*time.Time, bool)`

GetTemplateUpdatedAtOk returns a tuple with the TemplateUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateUpdatedAt

`func (o *InventoryItem) SetTemplateUpdatedAt(v time.Time)`

SetTemplateUpdatedAt sets TemplateUpdatedAt field to given value.


### GetUpdatedAt

`func (o *InventoryItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InventoryItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InventoryItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserAttributes

`func (o *InventoryItem) GetUserAttributes() InventoryUserAttributes`

GetUserAttributes returns the UserAttributes field if non-nil, zero value otherwise.

### GetUserAttributesOk

`func (o *InventoryItem) GetUserAttributesOk() (*InventoryUserAttributes, bool)`

GetUserAttributesOk returns a tuple with the UserAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttributes

`func (o *InventoryItem) SetUserAttributes(v InventoryUserAttributes)`

SetUserAttributes sets UserAttributes field to given value.


### GetValidateUserAttributes

`func (o *InventoryItem) GetValidateUserAttributes() bool`

GetValidateUserAttributes returns the ValidateUserAttributes field if non-nil, zero value otherwise.

### GetValidateUserAttributesOk

`func (o *InventoryItem) GetValidateUserAttributesOk() (*bool, bool)`

GetValidateUserAttributesOk returns a tuple with the ValidateUserAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateUserAttributes

`func (o *InventoryItem) SetValidateUserAttributes(v bool)`

SetValidateUserAttributes sets ValidateUserAttributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


