# AdminUserResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**AdminUserFields**](AdminUserFields.md) |  | [optional] 

## Methods

### NewAdminUserResponseModel

`func NewAdminUserResponseModel() *AdminUserResponseModel`

NewAdminUserResponseModel instantiates a new AdminUserResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminUserResponseModelWithDefaults

`func NewAdminUserResponseModelWithDefaults() *AdminUserResponseModel`

NewAdminUserResponseModelWithDefaults instantiates a new AdminUserResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AdminUserResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminUserResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminUserResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminUserResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *AdminUserResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AdminUserResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AdminUserResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AdminUserResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetUser

`func (o *AdminUserResponseModel) GetUser() AdminUserFields`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AdminUserResponseModel) GetUserOk() (*AdminUserFields, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AdminUserResponseModel) SetUser(v AdminUserFields)`

SetUser sets User field to given value.

### HasUser

`func (o *AdminUserResponseModel) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


