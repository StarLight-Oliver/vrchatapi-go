# APIConfigConstantsGROUPS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CAPACITY** | Pointer to **int32** | Maximum group capacity | [optional] [default to 100000]
**GROUP_TRANSFER_REQUIREMENTS** | Pointer to **[]string** | Requirements for transferring group ownership | [optional] 
**MAX_INVITES_REQUESTS** | Pointer to **int32** | Maximum number of invite requests | [optional] [default to 50]
**MAX_JOINED** | Pointer to **int32** | Maximum number of joined groups | [optional] [default to 100]
**MAX_JOINED_PLUS** | Pointer to **int32** | Maximum number of joined groups for VRChat Plus members | [optional] [default to 200]
**MAX_LANGUAGES** | Pointer to **int32** | Maximum number of supported languages | [optional] [default to 10]
**MAX_LINKS** | Pointer to **int32** | Maximum number of group links | [optional] [default to 3]
**MAX_MANAGEMENT_ROLES** | Pointer to **int32** | Maximum number of management roles in a group | [optional] [default to 5]
**MAX_OWNED** | Pointer to **int32** | Maximum number of groups a user can own | [optional] [default to 5]
**MAX_ROLES** | Pointer to **int32** | Maximum number of roles in a group | [optional] [default to 50]

## Methods

### NewAPIConfigConstantsGROUPS

`func NewAPIConfigConstantsGROUPS() *APIConfigConstantsGROUPS`

NewAPIConfigConstantsGROUPS instantiates a new APIConfigConstantsGROUPS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIConfigConstantsGROUPSWithDefaults

`func NewAPIConfigConstantsGROUPSWithDefaults() *APIConfigConstantsGROUPS`

NewAPIConfigConstantsGROUPSWithDefaults instantiates a new APIConfigConstantsGROUPS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCAPACITY

`func (o *APIConfigConstantsGROUPS) GetCAPACITY() int32`

GetCAPACITY returns the CAPACITY field if non-nil, zero value otherwise.

### GetCAPACITYOk

`func (o *APIConfigConstantsGROUPS) GetCAPACITYOk() (*int32, bool)`

GetCAPACITYOk returns a tuple with the CAPACITY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCAPACITY

`func (o *APIConfigConstantsGROUPS) SetCAPACITY(v int32)`

SetCAPACITY sets CAPACITY field to given value.

### HasCAPACITY

`func (o *APIConfigConstantsGROUPS) HasCAPACITY() bool`

HasCAPACITY returns a boolean if a field has been set.

### GetGROUP_TRANSFER_REQUIREMENTS

`func (o *APIConfigConstantsGROUPS) GetGROUP_TRANSFER_REQUIREMENTS() []string`

GetGROUP_TRANSFER_REQUIREMENTS returns the GROUP_TRANSFER_REQUIREMENTS field if non-nil, zero value otherwise.

### GetGROUP_TRANSFER_REQUIREMENTSOk

`func (o *APIConfigConstantsGROUPS) GetGROUP_TRANSFER_REQUIREMENTSOk() (*[]string, bool)`

GetGROUP_TRANSFER_REQUIREMENTSOk returns a tuple with the GROUP_TRANSFER_REQUIREMENTS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGROUP_TRANSFER_REQUIREMENTS

`func (o *APIConfigConstantsGROUPS) SetGROUP_TRANSFER_REQUIREMENTS(v []string)`

SetGROUP_TRANSFER_REQUIREMENTS sets GROUP_TRANSFER_REQUIREMENTS field to given value.

### HasGROUP_TRANSFER_REQUIREMENTS

`func (o *APIConfigConstantsGROUPS) HasGROUP_TRANSFER_REQUIREMENTS() bool`

HasGROUP_TRANSFER_REQUIREMENTS returns a boolean if a field has been set.

### GetMAX_INVITES_REQUESTS

`func (o *APIConfigConstantsGROUPS) GetMAX_INVITES_REQUESTS() int32`

GetMAX_INVITES_REQUESTS returns the MAX_INVITES_REQUESTS field if non-nil, zero value otherwise.

### GetMAX_INVITES_REQUESTSOk

`func (o *APIConfigConstantsGROUPS) GetMAX_INVITES_REQUESTSOk() (*int32, bool)`

GetMAX_INVITES_REQUESTSOk returns a tuple with the MAX_INVITES_REQUESTS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMAX_INVITES_REQUESTS

`func (o *APIConfigConstantsGROUPS) SetMAX_INVITES_REQUESTS(v int32)`

SetMAX_INVITES_REQUESTS sets MAX_INVITES_REQUESTS field to given value.

### HasMAX_INVITES_REQUESTS

`func (o *APIConfigConstantsGROUPS) HasMAX_INVITES_REQUESTS() bool`

HasMAX_INVITES_REQUESTS returns a boolean if a field has been set.

### GetMAX_JOINED

`func (o *APIConfigConstantsGROUPS) GetMAX_JOINED() int32`

GetMAX_JOINED returns the MAX_JOINED field if non-nil, zero value otherwise.

### GetMAX_JOINEDOk

`func (o *APIConfigConstantsGROUPS) GetMAX_JOINEDOk() (*int32, bool)`

GetMAX_JOINEDOk returns a tuple with the MAX_JOINED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMAX_JOINED

`func (o *APIConfigConstantsGROUPS) SetMAX_JOINED(v int32)`

SetMAX_JOINED sets MAX_JOINED field to given value.

### HasMAX_JOINED

`func (o *APIConfigConstantsGROUPS) HasMAX_JOINED() bool`

HasMAX_JOINED returns a boolean if a field has been set.

### GetMAX_JOINED_PLUS

`func (o *APIConfigConstantsGROUPS) GetMAX_JOINED_PLUS() int32`

GetMAX_JOINED_PLUS returns the MAX_JOINED_PLUS field if non-nil, zero value otherwise.

### GetMAX_JOINED_PLUSOk

`func (o *APIConfigConstantsGROUPS) GetMAX_JOINED_PLUSOk() (*int32, bool)`

GetMAX_JOINED_PLUSOk returns a tuple with the MAX_JOINED_PLUS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMAX_JOINED_PLUS

`func (o *APIConfigConstantsGROUPS) SetMAX_JOINED_PLUS(v int32)`

SetMAX_JOINED_PLUS sets MAX_JOINED_PLUS field to given value.

### HasMAX_JOINED_PLUS

`func (o *APIConfigConstantsGROUPS) HasMAX_JOINED_PLUS() bool`

HasMAX_JOINED_PLUS returns a boolean if a field has been set.

### GetMAX_LANGUAGES

`func (o *APIConfigConstantsGROUPS) GetMAX_LANGUAGES() int32`

GetMAX_LANGUAGES returns the MAX_LANGUAGES field if non-nil, zero value otherwise.

### GetMAX_LANGUAGESOk

`func (o *APIConfigConstantsGROUPS) GetMAX_LANGUAGESOk() (*int32, bool)`

GetMAX_LANGUAGESOk returns a tuple with the MAX_LANGUAGES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMAX_LANGUAGES

`func (o *APIConfigConstantsGROUPS) SetMAX_LANGUAGES(v int32)`

SetMAX_LANGUAGES sets MAX_LANGUAGES field to given value.

### HasMAX_LANGUAGES

`func (o *APIConfigConstantsGROUPS) HasMAX_LANGUAGES() bool`

HasMAX_LANGUAGES returns a boolean if a field has been set.

### GetMAX_LINKS

`func (o *APIConfigConstantsGROUPS) GetMAX_LINKS() int32`

GetMAX_LINKS returns the MAX_LINKS field if non-nil, zero value otherwise.

### GetMAX_LINKSOk

`func (o *APIConfigConstantsGROUPS) GetMAX_LINKSOk() (*int32, bool)`

GetMAX_LINKSOk returns a tuple with the MAX_LINKS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMAX_LINKS

`func (o *APIConfigConstantsGROUPS) SetMAX_LINKS(v int32)`

SetMAX_LINKS sets MAX_LINKS field to given value.

### HasMAX_LINKS

`func (o *APIConfigConstantsGROUPS) HasMAX_LINKS() bool`

HasMAX_LINKS returns a boolean if a field has been set.

### GetMAX_MANAGEMENT_ROLES

`func (o *APIConfigConstantsGROUPS) GetMAX_MANAGEMENT_ROLES() int32`

GetMAX_MANAGEMENT_ROLES returns the MAX_MANAGEMENT_ROLES field if non-nil, zero value otherwise.

### GetMAX_MANAGEMENT_ROLESOk

`func (o *APIConfigConstantsGROUPS) GetMAX_MANAGEMENT_ROLESOk() (*int32, bool)`

GetMAX_MANAGEMENT_ROLESOk returns a tuple with the MAX_MANAGEMENT_ROLES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMAX_MANAGEMENT_ROLES

`func (o *APIConfigConstantsGROUPS) SetMAX_MANAGEMENT_ROLES(v int32)`

SetMAX_MANAGEMENT_ROLES sets MAX_MANAGEMENT_ROLES field to given value.

### HasMAX_MANAGEMENT_ROLES

`func (o *APIConfigConstantsGROUPS) HasMAX_MANAGEMENT_ROLES() bool`

HasMAX_MANAGEMENT_ROLES returns a boolean if a field has been set.

### GetMAX_OWNED

`func (o *APIConfigConstantsGROUPS) GetMAX_OWNED() int32`

GetMAX_OWNED returns the MAX_OWNED field if non-nil, zero value otherwise.

### GetMAX_OWNEDOk

`func (o *APIConfigConstantsGROUPS) GetMAX_OWNEDOk() (*int32, bool)`

GetMAX_OWNEDOk returns a tuple with the MAX_OWNED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMAX_OWNED

`func (o *APIConfigConstantsGROUPS) SetMAX_OWNED(v int32)`

SetMAX_OWNED sets MAX_OWNED field to given value.

### HasMAX_OWNED

`func (o *APIConfigConstantsGROUPS) HasMAX_OWNED() bool`

HasMAX_OWNED returns a boolean if a field has been set.

### GetMAX_ROLES

`func (o *APIConfigConstantsGROUPS) GetMAX_ROLES() int32`

GetMAX_ROLES returns the MAX_ROLES field if non-nil, zero value otherwise.

### GetMAX_ROLESOk

`func (o *APIConfigConstantsGROUPS) GetMAX_ROLESOk() (*int32, bool)`

GetMAX_ROLESOk returns a tuple with the MAX_ROLES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMAX_ROLES

`func (o *APIConfigConstantsGROUPS) SetMAX_ROLES(v int32)`

SetMAX_ROLES sets MAX_ROLES field to given value.

### HasMAX_ROLES

`func (o *APIConfigConstantsGROUPS) HasMAX_ROLES() bool`

HasMAX_ROLES returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


