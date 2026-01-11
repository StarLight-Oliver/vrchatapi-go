# DiscordDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | https://discord.com/developers/docs/reference#snowflakes | [optional] 

## Methods

### NewDiscordDetails

`func NewDiscordDetails() *DiscordDetails`

NewDiscordDetails instantiates a new DiscordDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscordDetailsWithDefaults

`func NewDiscordDetailsWithDefaults() *DiscordDetails`

NewDiscordDetailsWithDefaults instantiates a new DiscordDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalName

`func (o *DiscordDetails) GetGlobalName() string`

GetGlobalName returns the GlobalName field if non-nil, zero value otherwise.

### GetGlobalNameOk

`func (o *DiscordDetails) GetGlobalNameOk() (*string, bool)`

GetGlobalNameOk returns a tuple with the GlobalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalName

`func (o *DiscordDetails) SetGlobalName(v string)`

SetGlobalName sets GlobalName field to given value.

### HasGlobalName

`func (o *DiscordDetails) HasGlobalName() bool`

HasGlobalName returns a boolean if a field has been set.

### GetId

`func (o *DiscordDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiscordDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiscordDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DiscordDetails) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


