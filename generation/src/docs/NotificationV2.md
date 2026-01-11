# NotificationV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanDelete** | **bool** |  | 
**Category** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Data** | [**NotificationV2Data**](NotificationV2Data.md) |  | 
**Details** | Pointer to [**NotificationV2DetailsBoop**](NotificationV2DetailsBoop.md) |  | [optional] 
**ExpiresAt** | **time.Time** |  | 
**ExpiryAfterSeen** | **NullableInt32** |  | 
**Id** | **string** |  | 
**IgnoreDND** | **bool** |  | 
**ImageUrl** | **NullableString** |  | 
**IsSystem** | **bool** |  | 
**Link** | **string** |  | 
**LinkText** | **string** |  | 
**LinkTextKey** | **NullableString** |  | 
**Message** | **string** |  | 
**MessageKey** | Pointer to **NullableString** |  | [optional] 
**ReceiverUserId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**RelatedNotificationsId** | **NullableString** |  | 
**RequireSeen** | **bool** |  | 
**Responses** | [**[]NotificationV2Response**](NotificationV2Response.md) |  | 
**Seen** | **bool** |  | 
**SenderUserId** | **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | 
**SenderUsername** | **NullableString** |  | 
**Title** | **string** |  | 
**TitleKey** | **NullableString** |  | 
**Type** | [**NotificationV2Type**](NotificationV2Type.md) |  | [default to GROUP_ANNOUNCEMENT]
**UpdatedAt** | **time.Time** |  | 
**Version** | **int32** |  | [default to 2]

## Methods

### NewNotificationV2

`func NewNotificationV2(canDelete bool, category string, createdAt time.Time, data NotificationV2Data, expiresAt time.Time, expiryAfterSeen NullableInt32, id string, ignoreDND bool, imageUrl NullableString, isSystem bool, link string, linkText string, linkTextKey NullableString, message string, receiverUserId string, relatedNotificationsId NullableString, requireSeen bool, responses []NotificationV2Response, seen bool, senderUserId string, senderUsername NullableString, title string, titleKey NullableString, type_ NotificationV2Type, updatedAt time.Time, version int32, ) *NotificationV2`

NewNotificationV2 instantiates a new NotificationV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationV2WithDefaults

`func NewNotificationV2WithDefaults() *NotificationV2`

NewNotificationV2WithDefaults instantiates a new NotificationV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanDelete

`func (o *NotificationV2) GetCanDelete() bool`

GetCanDelete returns the CanDelete field if non-nil, zero value otherwise.

### GetCanDeleteOk

`func (o *NotificationV2) GetCanDeleteOk() (*bool, bool)`

GetCanDeleteOk returns a tuple with the CanDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDelete

`func (o *NotificationV2) SetCanDelete(v bool)`

SetCanDelete sets CanDelete field to given value.


### GetCategory

`func (o *NotificationV2) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *NotificationV2) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *NotificationV2) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetCreatedAt

`func (o *NotificationV2) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationV2) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationV2) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetData

`func (o *NotificationV2) GetData() NotificationV2Data`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NotificationV2) GetDataOk() (*NotificationV2Data, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NotificationV2) SetData(v NotificationV2Data)`

SetData sets Data field to given value.


### GetDetails

`func (o *NotificationV2) GetDetails() NotificationV2DetailsBoop`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *NotificationV2) GetDetailsOk() (*NotificationV2DetailsBoop, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *NotificationV2) SetDetails(v NotificationV2DetailsBoop)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *NotificationV2) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetExpiresAt

`func (o *NotificationV2) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *NotificationV2) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *NotificationV2) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetExpiryAfterSeen

`func (o *NotificationV2) GetExpiryAfterSeen() int32`

GetExpiryAfterSeen returns the ExpiryAfterSeen field if non-nil, zero value otherwise.

### GetExpiryAfterSeenOk

`func (o *NotificationV2) GetExpiryAfterSeenOk() (*int32, bool)`

GetExpiryAfterSeenOk returns a tuple with the ExpiryAfterSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryAfterSeen

`func (o *NotificationV2) SetExpiryAfterSeen(v int32)`

SetExpiryAfterSeen sets ExpiryAfterSeen field to given value.


### SetExpiryAfterSeenNil

`func (o *NotificationV2) SetExpiryAfterSeenNil(b bool)`

 SetExpiryAfterSeenNil sets the value for ExpiryAfterSeen to be an explicit nil

### UnsetExpiryAfterSeen
`func (o *NotificationV2) UnsetExpiryAfterSeen()`

UnsetExpiryAfterSeen ensures that no value is present for ExpiryAfterSeen, not even an explicit nil
### GetId

`func (o *NotificationV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationV2) SetId(v string)`

SetId sets Id field to given value.


### GetIgnoreDND

`func (o *NotificationV2) GetIgnoreDND() bool`

GetIgnoreDND returns the IgnoreDND field if non-nil, zero value otherwise.

### GetIgnoreDNDOk

`func (o *NotificationV2) GetIgnoreDNDOk() (*bool, bool)`

GetIgnoreDNDOk returns a tuple with the IgnoreDND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreDND

`func (o *NotificationV2) SetIgnoreDND(v bool)`

SetIgnoreDND sets IgnoreDND field to given value.


### GetImageUrl

`func (o *NotificationV2) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *NotificationV2) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *NotificationV2) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.


### SetImageUrlNil

`func (o *NotificationV2) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *NotificationV2) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetIsSystem

`func (o *NotificationV2) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *NotificationV2) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *NotificationV2) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.


### GetLink

`func (o *NotificationV2) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *NotificationV2) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *NotificationV2) SetLink(v string)`

SetLink sets Link field to given value.


### GetLinkText

`func (o *NotificationV2) GetLinkText() string`

GetLinkText returns the LinkText field if non-nil, zero value otherwise.

### GetLinkTextOk

`func (o *NotificationV2) GetLinkTextOk() (*string, bool)`

GetLinkTextOk returns a tuple with the LinkText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkText

`func (o *NotificationV2) SetLinkText(v string)`

SetLinkText sets LinkText field to given value.


### GetLinkTextKey

`func (o *NotificationV2) GetLinkTextKey() string`

GetLinkTextKey returns the LinkTextKey field if non-nil, zero value otherwise.

### GetLinkTextKeyOk

`func (o *NotificationV2) GetLinkTextKeyOk() (*string, bool)`

GetLinkTextKeyOk returns a tuple with the LinkTextKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkTextKey

`func (o *NotificationV2) SetLinkTextKey(v string)`

SetLinkTextKey sets LinkTextKey field to given value.


### SetLinkTextKeyNil

`func (o *NotificationV2) SetLinkTextKeyNil(b bool)`

 SetLinkTextKeyNil sets the value for LinkTextKey to be an explicit nil

### UnsetLinkTextKey
`func (o *NotificationV2) UnsetLinkTextKey()`

UnsetLinkTextKey ensures that no value is present for LinkTextKey, not even an explicit nil
### GetMessage

`func (o *NotificationV2) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NotificationV2) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NotificationV2) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetMessageKey

`func (o *NotificationV2) GetMessageKey() string`

GetMessageKey returns the MessageKey field if non-nil, zero value otherwise.

### GetMessageKeyOk

`func (o *NotificationV2) GetMessageKeyOk() (*string, bool)`

GetMessageKeyOk returns a tuple with the MessageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageKey

`func (o *NotificationV2) SetMessageKey(v string)`

SetMessageKey sets MessageKey field to given value.

### HasMessageKey

`func (o *NotificationV2) HasMessageKey() bool`

HasMessageKey returns a boolean if a field has been set.

### SetMessageKeyNil

`func (o *NotificationV2) SetMessageKeyNil(b bool)`

 SetMessageKeyNil sets the value for MessageKey to be an explicit nil

### UnsetMessageKey
`func (o *NotificationV2) UnsetMessageKey()`

UnsetMessageKey ensures that no value is present for MessageKey, not even an explicit nil
### GetReceiverUserId

`func (o *NotificationV2) GetReceiverUserId() string`

GetReceiverUserId returns the ReceiverUserId field if non-nil, zero value otherwise.

### GetReceiverUserIdOk

`func (o *NotificationV2) GetReceiverUserIdOk() (*string, bool)`

GetReceiverUserIdOk returns a tuple with the ReceiverUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverUserId

`func (o *NotificationV2) SetReceiverUserId(v string)`

SetReceiverUserId sets ReceiverUserId field to given value.


### GetRelatedNotificationsId

`func (o *NotificationV2) GetRelatedNotificationsId() string`

GetRelatedNotificationsId returns the RelatedNotificationsId field if non-nil, zero value otherwise.

### GetRelatedNotificationsIdOk

`func (o *NotificationV2) GetRelatedNotificationsIdOk() (*string, bool)`

GetRelatedNotificationsIdOk returns a tuple with the RelatedNotificationsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedNotificationsId

`func (o *NotificationV2) SetRelatedNotificationsId(v string)`

SetRelatedNotificationsId sets RelatedNotificationsId field to given value.


### SetRelatedNotificationsIdNil

`func (o *NotificationV2) SetRelatedNotificationsIdNil(b bool)`

 SetRelatedNotificationsIdNil sets the value for RelatedNotificationsId to be an explicit nil

### UnsetRelatedNotificationsId
`func (o *NotificationV2) UnsetRelatedNotificationsId()`

UnsetRelatedNotificationsId ensures that no value is present for RelatedNotificationsId, not even an explicit nil
### GetRequireSeen

`func (o *NotificationV2) GetRequireSeen() bool`

GetRequireSeen returns the RequireSeen field if non-nil, zero value otherwise.

### GetRequireSeenOk

`func (o *NotificationV2) GetRequireSeenOk() (*bool, bool)`

GetRequireSeenOk returns a tuple with the RequireSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSeen

`func (o *NotificationV2) SetRequireSeen(v bool)`

SetRequireSeen sets RequireSeen field to given value.


### GetResponses

`func (o *NotificationV2) GetResponses() []NotificationV2Response`

GetResponses returns the Responses field if non-nil, zero value otherwise.

### GetResponsesOk

`func (o *NotificationV2) GetResponsesOk() (*[]NotificationV2Response, bool)`

GetResponsesOk returns a tuple with the Responses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponses

`func (o *NotificationV2) SetResponses(v []NotificationV2Response)`

SetResponses sets Responses field to given value.


### GetSeen

`func (o *NotificationV2) GetSeen() bool`

GetSeen returns the Seen field if non-nil, zero value otherwise.

### GetSeenOk

`func (o *NotificationV2) GetSeenOk() (*bool, bool)`

GetSeenOk returns a tuple with the Seen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeen

`func (o *NotificationV2) SetSeen(v bool)`

SetSeen sets Seen field to given value.


### GetSenderUserId

`func (o *NotificationV2) GetSenderUserId() string`

GetSenderUserId returns the SenderUserId field if non-nil, zero value otherwise.

### GetSenderUserIdOk

`func (o *NotificationV2) GetSenderUserIdOk() (*string, bool)`

GetSenderUserIdOk returns a tuple with the SenderUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderUserId

`func (o *NotificationV2) SetSenderUserId(v string)`

SetSenderUserId sets SenderUserId field to given value.


### GetSenderUsername

`func (o *NotificationV2) GetSenderUsername() string`

GetSenderUsername returns the SenderUsername field if non-nil, zero value otherwise.

### GetSenderUsernameOk

`func (o *NotificationV2) GetSenderUsernameOk() (*string, bool)`

GetSenderUsernameOk returns a tuple with the SenderUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderUsername

`func (o *NotificationV2) SetSenderUsername(v string)`

SetSenderUsername sets SenderUsername field to given value.


### SetSenderUsernameNil

`func (o *NotificationV2) SetSenderUsernameNil(b bool)`

 SetSenderUsernameNil sets the value for SenderUsername to be an explicit nil

### UnsetSenderUsername
`func (o *NotificationV2) UnsetSenderUsername()`

UnsetSenderUsername ensures that no value is present for SenderUsername, not even an explicit nil
### GetTitle

`func (o *NotificationV2) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NotificationV2) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NotificationV2) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetTitleKey

`func (o *NotificationV2) GetTitleKey() string`

GetTitleKey returns the TitleKey field if non-nil, zero value otherwise.

### GetTitleKeyOk

`func (o *NotificationV2) GetTitleKeyOk() (*string, bool)`

GetTitleKeyOk returns a tuple with the TitleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleKey

`func (o *NotificationV2) SetTitleKey(v string)`

SetTitleKey sets TitleKey field to given value.


### SetTitleKeyNil

`func (o *NotificationV2) SetTitleKeyNil(b bool)`

 SetTitleKeyNil sets the value for TitleKey to be an explicit nil

### UnsetTitleKey
`func (o *NotificationV2) UnsetTitleKey()`

UnsetTitleKey ensures that no value is present for TitleKey, not even an explicit nil
### GetType

`func (o *NotificationV2) GetType() NotificationV2Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationV2) GetTypeOk() (*NotificationV2Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationV2) SetType(v NotificationV2Type)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *NotificationV2) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NotificationV2) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NotificationV2) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVersion

`func (o *NotificationV2) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NotificationV2) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NotificationV2) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


