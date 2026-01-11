# InventoryTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**Collections** | **[]string** |  | 
**CreatedAt** | **time.Time** |  | 
**DefaultAttributes** | **map[string]interface{}** |  | 
**Description** | **string** |  | 
**EquipSlots** | **[]string** |  | 
**Flags** | **[]string** |  | 
**Id** | **string** |  | 
**ImageUrl** | **string** |  | 
**ItemType** | [**InventoryItemType**](InventoryItemType.md) |  | [default to BUNDLE]
**ItemTypeLabel** | **string** |  | 
**Metadata** | Pointer to [**InventoryMetadata**](InventoryMetadata.md) |  | [optional] 
**Name** | **string** |  | 
**NotificationDetails** | Pointer to [**InventoryNotificationDetails**](InventoryNotificationDetails.md) |  | [optional] 
**Status** | **string** |  | 
**Tags** | **[]string** |  | 
**UpdatedAt** | **time.Time** |  | 
**ValidateUserAttributes** | **bool** |  | 

## Methods

### NewInventoryTemplate

`func NewInventoryTemplate(authorId string, collections []string, createdAt time.Time, defaultAttributes map[string]interface{}, description string, equipSlots []string, flags []string, id string, imageUrl string, itemType InventoryItemType, itemTypeLabel string, name string, status string, tags []string, updatedAt time.Time, validateUserAttributes bool, ) *InventoryTemplate`

NewInventoryTemplate instantiates a new InventoryTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryTemplateWithDefaults

`func NewInventoryTemplateWithDefaults() *InventoryTemplate`

NewInventoryTemplateWithDefaults instantiates a new InventoryTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorId

`func (o *InventoryTemplate) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *InventoryTemplate) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *InventoryTemplate) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.


### GetCollections

`func (o *InventoryTemplate) GetCollections() []string`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *InventoryTemplate) GetCollectionsOk() (*[]string, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *InventoryTemplate) SetCollections(v []string)`

SetCollections sets Collections field to given value.


### GetCreatedAt

`func (o *InventoryTemplate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InventoryTemplate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InventoryTemplate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDefaultAttributes

`func (o *InventoryTemplate) GetDefaultAttributes() map[string]interface{}`

GetDefaultAttributes returns the DefaultAttributes field if non-nil, zero value otherwise.

### GetDefaultAttributesOk

`func (o *InventoryTemplate) GetDefaultAttributesOk() (*map[string]interface{}, bool)`

GetDefaultAttributesOk returns a tuple with the DefaultAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAttributes

`func (o *InventoryTemplate) SetDefaultAttributes(v map[string]interface{})`

SetDefaultAttributes sets DefaultAttributes field to given value.


### GetDescription

`func (o *InventoryTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InventoryTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InventoryTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEquipSlots

`func (o *InventoryTemplate) GetEquipSlots() []string`

GetEquipSlots returns the EquipSlots field if non-nil, zero value otherwise.

### GetEquipSlotsOk

`func (o *InventoryTemplate) GetEquipSlotsOk() (*[]string, bool)`

GetEquipSlotsOk returns a tuple with the EquipSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipSlots

`func (o *InventoryTemplate) SetEquipSlots(v []string)`

SetEquipSlots sets EquipSlots field to given value.


### GetFlags

`func (o *InventoryTemplate) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *InventoryTemplate) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *InventoryTemplate) SetFlags(v []string)`

SetFlags sets Flags field to given value.


### GetId

`func (o *InventoryTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InventoryTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InventoryTemplate) SetId(v string)`

SetId sets Id field to given value.


### GetImageUrl

`func (o *InventoryTemplate) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *InventoryTemplate) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *InventoryTemplate) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.


### GetItemType

`func (o *InventoryTemplate) GetItemType() InventoryItemType`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *InventoryTemplate) GetItemTypeOk() (*InventoryItemType, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *InventoryTemplate) SetItemType(v InventoryItemType)`

SetItemType sets ItemType field to given value.


### GetItemTypeLabel

`func (o *InventoryTemplate) GetItemTypeLabel() string`

GetItemTypeLabel returns the ItemTypeLabel field if non-nil, zero value otherwise.

### GetItemTypeLabelOk

`func (o *InventoryTemplate) GetItemTypeLabelOk() (*string, bool)`

GetItemTypeLabelOk returns a tuple with the ItemTypeLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTypeLabel

`func (o *InventoryTemplate) SetItemTypeLabel(v string)`

SetItemTypeLabel sets ItemTypeLabel field to given value.


### GetMetadata

`func (o *InventoryTemplate) GetMetadata() InventoryMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InventoryTemplate) GetMetadataOk() (*InventoryMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InventoryTemplate) SetMetadata(v InventoryMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InventoryTemplate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *InventoryTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InventoryTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InventoryTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetNotificationDetails

`func (o *InventoryTemplate) GetNotificationDetails() InventoryNotificationDetails`

GetNotificationDetails returns the NotificationDetails field if non-nil, zero value otherwise.

### GetNotificationDetailsOk

`func (o *InventoryTemplate) GetNotificationDetailsOk() (*InventoryNotificationDetails, bool)`

GetNotificationDetailsOk returns a tuple with the NotificationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDetails

`func (o *InventoryTemplate) SetNotificationDetails(v InventoryNotificationDetails)`

SetNotificationDetails sets NotificationDetails field to given value.

### HasNotificationDetails

`func (o *InventoryTemplate) HasNotificationDetails() bool`

HasNotificationDetails returns a boolean if a field has been set.

### GetStatus

`func (o *InventoryTemplate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InventoryTemplate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InventoryTemplate) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTags

`func (o *InventoryTemplate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InventoryTemplate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InventoryTemplate) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetUpdatedAt

`func (o *InventoryTemplate) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InventoryTemplate) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InventoryTemplate) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetValidateUserAttributes

`func (o *InventoryTemplate) GetValidateUserAttributes() bool`

GetValidateUserAttributes returns the ValidateUserAttributes field if non-nil, zero value otherwise.

### GetValidateUserAttributesOk

`func (o *InventoryTemplate) GetValidateUserAttributesOk() (*bool, bool)`

GetValidateUserAttributesOk returns a tuple with the ValidateUserAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateUserAttributes

`func (o *InventoryTemplate) SetValidateUserAttributes(v bool)`

SetValidateUserAttributes sets ValidateUserAttributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


