# UserExists

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NameOk** | Pointer to **bool** | Is the username valid? | [optional] [default to false]
**UserExists** | **bool** | Status if a user exist with that username or userId. | [default to false]

## Methods

### NewUserExists

`func NewUserExists(userExists bool, ) *UserExists`

NewUserExists instantiates a new UserExists object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserExistsWithDefaults

`func NewUserExistsWithDefaults() *UserExists`

NewUserExistsWithDefaults instantiates a new UserExists object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNameOk

`func (o *UserExists) GetNameOk() bool`

GetNameOk returns the NameOk field if non-nil, zero value otherwise.

### GetNameOkOk

`func (o *UserExists) GetNameOkOk() (*bool, bool)`

GetNameOkOk returns a tuple with the NameOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOk

`func (o *UserExists) SetNameOk(v bool)`

SetNameOk sets NameOk field to given value.

### HasNameOk

`func (o *UserExists) HasNameOk() bool`

HasNameOk returns a boolean if a field has been set.

### GetUserExists

`func (o *UserExists) GetUserExists() bool`

GetUserExists returns the UserExists field if non-nil, zero value otherwise.

### GetUserExistsOk

`func (o *UserExists) GetUserExistsOk() (*bool, bool)`

GetUserExistsOk returns a tuple with the UserExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserExists

`func (o *UserExists) SetUserExists(v bool)`

SetUserExists sets UserExists field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


