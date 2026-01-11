# NotificationV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **string** |  | 
**Icon** | **string** |  | 
**Text** | **string** |  | 
**TextKey** | **NullableString** |  | 
**Type** | **string** |  | 

## Methods

### NewNotificationV2Response

`func NewNotificationV2Response(data string, icon string, text string, textKey NullableString, type_ string, ) *NotificationV2Response`

NewNotificationV2Response instantiates a new NotificationV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationV2ResponseWithDefaults

`func NewNotificationV2ResponseWithDefaults() *NotificationV2Response`

NewNotificationV2ResponseWithDefaults instantiates a new NotificationV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *NotificationV2Response) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NotificationV2Response) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NotificationV2Response) SetData(v string)`

SetData sets Data field to given value.


### GetIcon

`func (o *NotificationV2Response) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *NotificationV2Response) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *NotificationV2Response) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetText

`func (o *NotificationV2Response) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *NotificationV2Response) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *NotificationV2Response) SetText(v string)`

SetText sets Text field to given value.


### GetTextKey

`func (o *NotificationV2Response) GetTextKey() string`

GetTextKey returns the TextKey field if non-nil, zero value otherwise.

### GetTextKeyOk

`func (o *NotificationV2Response) GetTextKeyOk() (*string, bool)`

GetTextKeyOk returns a tuple with the TextKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextKey

`func (o *NotificationV2Response) SetTextKey(v string)`

SetTextKey sets TextKey field to given value.


### SetTextKeyNil

`func (o *NotificationV2Response) SetTextKeyNil(b bool)`

 SetTextKeyNil sets the value for TextKey to be an explicit nil

### UnsetTextKey
`func (o *NotificationV2Response) UnsetTextKey()`

UnsetTextKey ensures that no value is present for TextKey, not even an explicit nil
### GetType

`func (o *NotificationV2Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationV2Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationV2Response) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


