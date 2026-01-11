# InventoryDrop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**CreatedAt** | **time.Time** |  | 
**DropExpiryDate** | **NullableTime** |  | 
**EndDropDate** | **time.Time** |  | 
**Id** | **string** |  | 
**IsDisabled** | **bool** |  | 
**Name** | **string** |  | 
**NotificationDetails** | [**InventoryNotificationDetails**](InventoryNotificationDetails.md) |  | 
**StartDropDate** | **time.Time** |  | 
**Status** | **string** |  | 
**Tags** | **[]string** |  | 
**TargetGroup** | **string** |  | 
**TemplateIds** | **[]string** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewInventoryDrop

`func NewInventoryDrop(authorId string, createdAt time.Time, dropExpiryDate NullableTime, endDropDate time.Time, id string, isDisabled bool, name string, notificationDetails InventoryNotificationDetails, startDropDate time.Time, status string, tags []string, targetGroup string, templateIds []string, updatedAt time.Time, ) *InventoryDrop`

NewInventoryDrop instantiates a new InventoryDrop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryDropWithDefaults

`func NewInventoryDropWithDefaults() *InventoryDrop`

NewInventoryDropWithDefaults instantiates a new InventoryDrop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorId

`func (o *InventoryDrop) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *InventoryDrop) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *InventoryDrop) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.


### GetCreatedAt

`func (o *InventoryDrop) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InventoryDrop) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InventoryDrop) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDropExpiryDate

`func (o *InventoryDrop) GetDropExpiryDate() time.Time`

GetDropExpiryDate returns the DropExpiryDate field if non-nil, zero value otherwise.

### GetDropExpiryDateOk

`func (o *InventoryDrop) GetDropExpiryDateOk() (*time.Time, bool)`

GetDropExpiryDateOk returns a tuple with the DropExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropExpiryDate

`func (o *InventoryDrop) SetDropExpiryDate(v time.Time)`

SetDropExpiryDate sets DropExpiryDate field to given value.


### SetDropExpiryDateNil

`func (o *InventoryDrop) SetDropExpiryDateNil(b bool)`

 SetDropExpiryDateNil sets the value for DropExpiryDate to be an explicit nil

### UnsetDropExpiryDate
`func (o *InventoryDrop) UnsetDropExpiryDate()`

UnsetDropExpiryDate ensures that no value is present for DropExpiryDate, not even an explicit nil
### GetEndDropDate

`func (o *InventoryDrop) GetEndDropDate() time.Time`

GetEndDropDate returns the EndDropDate field if non-nil, zero value otherwise.

### GetEndDropDateOk

`func (o *InventoryDrop) GetEndDropDateOk() (*time.Time, bool)`

GetEndDropDateOk returns a tuple with the EndDropDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDropDate

`func (o *InventoryDrop) SetEndDropDate(v time.Time)`

SetEndDropDate sets EndDropDate field to given value.


### GetId

`func (o *InventoryDrop) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InventoryDrop) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InventoryDrop) SetId(v string)`

SetId sets Id field to given value.


### GetIsDisabled

`func (o *InventoryDrop) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *InventoryDrop) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *InventoryDrop) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.


### GetName

`func (o *InventoryDrop) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InventoryDrop) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InventoryDrop) SetName(v string)`

SetName sets Name field to given value.


### GetNotificationDetails

`func (o *InventoryDrop) GetNotificationDetails() InventoryNotificationDetails`

GetNotificationDetails returns the NotificationDetails field if non-nil, zero value otherwise.

### GetNotificationDetailsOk

`func (o *InventoryDrop) GetNotificationDetailsOk() (*InventoryNotificationDetails, bool)`

GetNotificationDetailsOk returns a tuple with the NotificationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDetails

`func (o *InventoryDrop) SetNotificationDetails(v InventoryNotificationDetails)`

SetNotificationDetails sets NotificationDetails field to given value.


### GetStartDropDate

`func (o *InventoryDrop) GetStartDropDate() time.Time`

GetStartDropDate returns the StartDropDate field if non-nil, zero value otherwise.

### GetStartDropDateOk

`func (o *InventoryDrop) GetStartDropDateOk() (*time.Time, bool)`

GetStartDropDateOk returns a tuple with the StartDropDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDropDate

`func (o *InventoryDrop) SetStartDropDate(v time.Time)`

SetStartDropDate sets StartDropDate field to given value.


### GetStatus

`func (o *InventoryDrop) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InventoryDrop) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InventoryDrop) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTags

`func (o *InventoryDrop) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InventoryDrop) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InventoryDrop) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetTargetGroup

`func (o *InventoryDrop) GetTargetGroup() string`

GetTargetGroup returns the TargetGroup field if non-nil, zero value otherwise.

### GetTargetGroupOk

`func (o *InventoryDrop) GetTargetGroupOk() (*string, bool)`

GetTargetGroupOk returns a tuple with the TargetGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroup

`func (o *InventoryDrop) SetTargetGroup(v string)`

SetTargetGroup sets TargetGroup field to given value.


### GetTemplateIds

`func (o *InventoryDrop) GetTemplateIds() []string`

GetTemplateIds returns the TemplateIds field if non-nil, zero value otherwise.

### GetTemplateIdsOk

`func (o *InventoryDrop) GetTemplateIdsOk() (*[]string, bool)`

GetTemplateIdsOk returns a tuple with the TemplateIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateIds

`func (o *InventoryDrop) SetTemplateIds(v []string)`

SetTemplateIds sets TemplateIds field to given value.


### GetUpdatedAt

`func (o *InventoryDrop) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InventoryDrop) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InventoryDrop) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


