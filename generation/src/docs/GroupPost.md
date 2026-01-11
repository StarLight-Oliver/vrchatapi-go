# GroupPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorId** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**EditorId** | Pointer to **string** | A users unique ID, usually in the form of &#x60;usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469&#x60;. Legacy players can have old IDs in the form of &#x60;8JoV9XEdpo&#x60;. The ID can never be changed. | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ImageId** | Pointer to **string** |  | [optional] 
**ImageUrl** | Pointer to **NullableString** |  | [optional] 
**RoleId** | Pointer to **[]string** |   | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Visibility** | Pointer to [**GroupPostVisibility**](GroupPostVisibility.md) |  | [optional] 

## Methods

### NewGroupPost

`func NewGroupPost() *GroupPost`

NewGroupPost instantiates a new GroupPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupPostWithDefaults

`func NewGroupPostWithDefaults() *GroupPost`

NewGroupPostWithDefaults instantiates a new GroupPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorId

`func (o *GroupPost) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *GroupPost) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *GroupPost) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *GroupPost) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GroupPost) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GroupPost) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GroupPost) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GroupPost) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEditorId

`func (o *GroupPost) GetEditorId() string`

GetEditorId returns the EditorId field if non-nil, zero value otherwise.

### GetEditorIdOk

`func (o *GroupPost) GetEditorIdOk() (*string, bool)`

GetEditorIdOk returns a tuple with the EditorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorId

`func (o *GroupPost) SetEditorId(v string)`

SetEditorId sets EditorId field to given value.

### HasEditorId

`func (o *GroupPost) HasEditorId() bool`

HasEditorId returns a boolean if a field has been set.

### GetGroupId

`func (o *GroupPost) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupPost) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupPost) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *GroupPost) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *GroupPost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupPost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupPost) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupPost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageId

`func (o *GroupPost) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *GroupPost) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *GroupPost) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *GroupPost) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetImageUrl

`func (o *GroupPost) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *GroupPost) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *GroupPost) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *GroupPost) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *GroupPost) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *GroupPost) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetRoleId

`func (o *GroupPost) GetRoleId() []string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *GroupPost) GetRoleIdOk() (*[]string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *GroupPost) SetRoleId(v []string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *GroupPost) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetText

`func (o *GroupPost) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *GroupPost) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *GroupPost) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *GroupPost) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTitle

`func (o *GroupPost) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GroupPost) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GroupPost) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GroupPost) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GroupPost) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GroupPost) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GroupPost) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GroupPost) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVisibility

`func (o *GroupPost) GetVisibility() GroupPostVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GroupPost) GetVisibilityOk() (*GroupPostVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GroupPost) SetVisibility(v GroupPostVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GroupPost) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


